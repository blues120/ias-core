package ecx

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/agent"
	"gitlab.ctyuncdn.cn/ias/ias-core/errors"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/openapi"
)

const (
	redisKeyStartTimesCache = "vss:start_times" // 缓存停止中的任务用于状态显示，Pod 无此状态
)

type schedulerVSS struct {
	bc     *conf.Bootstrap
	rdb    *redis.Client
	signUc *biz.SignatureUsecase

	log *log.Helper
}

func newRedisClient(c *conf.Data_Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           int(c.Db),
		WriteTimeout: c.WriteTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

func NewSchedulerVSS(bc *conf.Bootstrap, signUc *biz.SignatureUsecase, logger log.Logger) biz.SchedulerVSSRepo {
	return &schedulerVSS{
		bc:     bc,
		rdb:    newRedisClient(bc.Data.Redis),
		signUc: signUc,
		log:    log.NewHelper(logger),
	}
}

// StartVSS 启动任务
func (r *schedulerVSS) StartVSS(ctx context.Context, tasks []*biz.Task) error {
	// 启动任务

	for _, task := range tasks {
		r.log.Debugf("StartVSS:task:%+v, device:%+v\n", *task, *task.Device)
	}
	startReqBody := r.GenerateStartBodyVSS(tasks)
	payload, err := json.Marshal(startReqBody)

	r.log.Debugf("StartVSS:%+v, %s\n", startReqBody, string(payload))
	openClient := r.newOpenApiClient()
	rsp, err := openClient.Post("/api/v1/kube/crd/itask/start/"+tasks[0].ParentId, bytes.NewBuffer(payload))
	r.log.Debug("ecx return is : " + string(rsp))
	if err != nil {
		return err
	}

	// api返回结果解析
	parsedRsp := EcxApiResp{}
	if err := json.Unmarshal([]byte(rsp), &parsedRsp); err != nil {
		r.log.Error("GetRegisterToken: json.Unmarshal error", err)
		return err
	}

	if parsedRsp.Code != "core.ok" {
		return errors.ErrorInvalidParam("启动任务时ECX返回错误:" + parsedRsp.Reason)
	}

	startTimesStr, _ := r.rdb.HGet(context.Background(), redisKeyStartTimesCache, tasks[0].ParentId).Result()
	startTimes, err := strconv.Atoi(startTimesStr)
	if err != nil {
		r.rdb.HSet(context.Background(), redisKeyStartTimesCache, tasks[0].ParentId, 1)
	} else {
		r.rdb.HSet(context.Background(), redisKeyStartTimesCache, tasks[0].ParentId, startTimes+1)
	}

	return nil
}

// StopVSS 停止任务
func (r *schedulerVSS) StopVSS(ctx context.Context, ta *biz.Task) error {
	// 停止任务
	openClient := r.newOpenApiClient()
	rsp, err := openClient.Post("/api/v1/kube/crd/itask/stop/"+ta.ParentId, nil)
	r.log.Debug("ecx return is : " + string(rsp))
	if err != nil {
		return err
	}

	// api返回结果解析
	parsedRsp := EcxApiResp{}
	if err := json.Unmarshal([]byte(rsp), &parsedRsp); err != nil {
		r.log.Error("GetRegisterToken: json.Unmarshal error", err)
		return err
	}
	if parsedRsp.Code != "core.ok" {
		return errors.ErrorInvalidParam("停止任务时ECX返回错误:" + parsedRsp.Reason)
	}

	return nil
}

// GetLogVSS 查询任务日志
func (r *schedulerVSS) GetLogVSS(ctx context.Context, ta *biz.Task, conn *websocket.Conn) error {
	algoInfo := ta.Algo.Algorithm
	modelName := getModelName(algoInfo.Image)
	if modelName == "" {
		return errors.ErrorInvalidParam("algo image required")
	}

	openClient := r.newOpenApiClient()

	// 当任务没有启动时候，额外查看镜像状态
	if ta.Status == biz.TaskStatusInitializing || ta.Status == biz.TaskStatusStarting ||
		ta.Status == biz.TaskStatusFailed {
		resp, err := openClient.Get("/api/v1/kube/deployment/" + modelName + "/" + ta.ParentId)
		if err == nil {
			err = conn.WriteMessage(websocket.TextMessage, []byte(resp+"\n\n"))
		}
	}

	// 获取日志
	_ = conn.WriteMessage(websocket.TextMessage, []byte(getJobName(ta.Id)+"任务日志:\n"))
	eventStream, err := openClient.SSEGet("/api/v1/kube/crd/itask/logs/" + modelName + "/" + ta.ParentId)
	if err != nil {
		return err
	}
	defer eventStream.Close()

	// 读取日志
	decoder := NewDecoder(eventStream)
	for {
		line, err := decoder.src.ReadString('\n')
		if err != nil {
			return err
		}
		encodedLine := base64.StdEncoding.EncodeToString([]byte(line))
		if err := conn.WriteMessage(websocket.TextMessage, []byte(encodedLine)); err != nil {
			r.log.Errorf("write data to ws failed %s", err)
			return err
		}
	}
}

func (r *schedulerVSS) newOpenApiClient() *openapi.OpenClient {
	return openapi.NewOpenClient(r.bc.OpenApi.Host, r.bc.OpenApi.AppId, r.bc.OpenApi.AppSecret, r.bc.OpenApi.Ac, "", r.log)
}

type StartReqBodyVSS struct {
	DATA StartReqBodyDataVSS `json:"data" yaml:"data"`
}

type StartReqBodyDataVSS struct {
	Sources   []StartReqBodySourceVSS    `json:"sources" yaml:"sources"`     // 视频流数组
	Algorithm StartReqBodyAlgorithmVSS   `json:"algorithm" yaml:"algorithm"` // 算法配置
	Task      StartReqBodyTaskProfileVSS `json:"task" yaml:"task"`           // 任务服务信息
	Box       StartReqBodyBoxVSS         `json:"box" yaml:"box"`             // 盒子配置
	Agent     agent.Agent                `json:"agent" yaml:"agent"`         // agent服务自用配置
}

type StartReqBodySourceVSS struct {
	Source string `json:"source" yaml:"source"` // 流地址
	// AccessToken       string  `json:"access_token" yaml:"access_token"`               // 视频流认证(用户名:密码) des3加密,密钥：cameraId
	CameraID          uint    `json:"camera_id" yaml:"camera_id"`                     // 摄像机ID
	DropFrameInterval float32 `json:"drop_frame_interval" yaml:"drop_frame_interval"` // 帧率
	Codec             string  `json:"codec" yaml:"codec"`                             // 视频流编解码器
	Width             int     `json:"width" yaml:"width"`                             // 摄像头分辨率-宽
	Height            int     `json:"height" yaml:"height"`                           // 摄像头分辨率-高
	// Coords            []*Coord `json:"coord" yaml:"coord"`                             // roi坐标信息；方向信息
	Events []agent.Event `json:"events" yaml:"events"` // 事件信息
}

type StartReqBodyAlgorithmVSS struct {
	Model []ModelVSS `json:"model" yaml:"model"`
}

type ModelVSS struct {
	Confidence string `json:"confidence" yaml:"confidence"` // 置信度阈值
	Enable     string `json:"enable" yaml:"enable"`         // 开关
	Name       string `json:"name" yaml:"name'"`            // 算法名
	RoutingKey string `json:"routing_key" yaml:"routing_key'"`
}

type StartReqBodyTaskProfileVSS struct {
	TaskID      string `json:"task_id" yaml:"task_id"`           // 任务ID
	BoxID       string `json:"box_id" yaml:"box_id"`             // 盒子id， 盒子纳管生成的boxId
	ZoneID      string `json:"zone_id" yaml:"zone_id"`           // 盒子区域id
	WorkspaceID string `json:"workspace_id" yaml:"workspace_id"` // 工作区id
	ModelName   string `json:"model_name" yaml:"model_name"`     // 模型名称
	ModelID     string `json:"model_id" yaml:"model_id"`         // 模型id
}

type StartReqBodyBoxVSS struct {
	Hostname      string `json:"hostname" yaml:"hostname"`               // //盒子mac地址
	TaskImageName string `json:"task_image_name" yaml:"task_image_name"` // 任务镜像名称
	TaskVersion   string `json:"task_version" yaml:"task_version"`       // 任务镜像tag
}

func (r *schedulerVSS) GenerateStartBodyVSS(meta []*biz.Task) StartReqBodyVSS {
	// 初始化结构体
	startReqBody := StartReqBodyVSS{
		DATA: StartReqBodyDataVSS{
			Sources: make([]StartReqBodySourceVSS, 0),
			Algorithm: StartReqBodyAlgorithmVSS{
				Model: make([]ModelVSS, 0),
			},
			Task:  StartReqBodyTaskProfileVSS{},
			Box:   StartReqBodyBoxVSS{},
			Agent: agent.Agent{},
		},
	}

	if len(meta) == 0 {
		return startReqBody
	}

	startReqBody.DATA.Sources = make([]StartReqBodySourceVSS, 0)
	// task信息
	box := meta[0].Device
	algo := meta[0].Algo.Algorithm
	startReqBody.DATA.Task = StartReqBodyTaskProfileVSS{
		TaskID:      meta[0].ParentId,
		BoxID:       box.ExtId,
		ZoneID:      box.ZoneId,
		WorkspaceID: box.WorkspaceId,
		ModelName:   getModelName(algo.Image),
		ModelID:     strconv.Itoa(int(algo.AlgoGroupID)),
	}
	// algorithm信息
	algorithmMap := make(map[uint64]ModelVSS)
	for _, task := range meta {
		if task.Algo == nil || task.Algo.Algorithm == nil {
			continue
		}
		algorithm := task.Algo.Algorithm
		algorithmMap[algorithm.ID] = ModelVSS{
			Confidence: "0",
			Enable:     "true",
			Name:       algorithm.AlgoNameEn,
			RoutingKey: algorithm.AppName,
		}
	}
	for _, algorithm := range algorithmMap {
		startReqBody.DATA.Algorithm.Model = append(startReqBody.DATA.Algorithm.Model, algorithm)
	}
	// box信息
	imageName, imageVersion := getImageNameAndVersion(algo.Image)
	startReqBody.DATA.Box = StartReqBodyBoxVSS{
		Hostname:      box.Mac + ".ctyun.cn",
		TaskImageName: imageName,
		TaskVersion:   imageVersion,
	}
	// camera信息
	cameraIdCameraMap := make(map[uint64][]biz.TaskCamera)
	cameraIdSourceMap := make(map[uint64]StartReqBodySourceVSS)
	for _, task := range meta {
		for _, taskCamera := range task.Cameras {
			taskCamera.ModeName = task.Algo.Algorithm.AlgoNameEn
			camera := taskCamera.Camera

			width := int(camera.StreamingInfo.Width)
			height := int(camera.StreamingInfo.Height)

			cameraIdCameraMap[taskCamera.Id] = append(cameraIdCameraMap[taskCamera.Id], taskCamera)
			cameraIdSourceMap[taskCamera.Id] = StartReqBodySourceVSS{
				Source:            camera.StreamingProtocol.Source(),
				CameraID:          uint(camera.Id),
				DropFrameInterval: float32(task.Algo.Interval),
				Codec:             camera.StreamingInfo.CodecName,
				Width:             width,
				Height:            height,
				Events:            make([]agent.Event, 0),
			}
		}
	}
	for cameraId, cameraIdSource := range cameraIdSourceMap {
		for _, taskCamera := range cameraIdCameraMap[cameraId] {
			cameraIdSource.Events = append(cameraIdSource.Events, agent.ParseEventsVSS(taskCamera.ModeName, taskCamera.MultiImgBox, uint(cameraIdSource.Width), uint(cameraIdSource.Height))...)
		}
		startReqBody.DATA.Sources = append(startReqBody.DATA.Sources, cameraIdSource)
	}

	// agent信息
	startReqBody.DATA.Agent = agent.Agent{
		Server: &agent.AgentServer{
			Http: &agent.AgentServerHttp{
				Addr:    "0.0.0.0:8000",
				Timeout: "1s",
			},
			Grpc: &agent.AgentServerGrpc{
				Addr:    "0.0.0.0:9000",
				Timeout: "1s",
			},
			File: &agent.AgentServerFile{
				OutputDir: "/data/meta",
			},
		},
		Callback: &agent.Callback{
			Addr:               r.bc.Callback.Addr,
			Urls:               map[string]string{},
			AuthEnable:         r.bc.Callback.AuthEnable,
			InsecureSkipVerify: true,
		},
	}
	for k, v := range r.bc.Callback.Urls {
		startReqBody.DATA.Agent.Callback.Urls[k] = v
	}
	if r.bc.Callback.AuthEnable && box.ExtId != "" {
		// 获取AppId/AppSecret并加密
		sign, err := r.signUc.FindByCondition(context.Background(), box.ExtId, "")
		if err != nil {
			r.log.Error("GenerateStartBodyVSS: signUc.FindByBoxId error", err)
			return startReqBody
		}

		// base64编码 todo优化加密方法
		encryptedAppId := base64.StdEncoding.EncodeToString([]byte(sign.AppId))
		encryptedAppSecret := base64.StdEncoding.EncodeToString([]byte(sign.AppSecret))

		startReqBody.DATA.Agent.Callback.Signature = &agent.Signature{
			AppId:     encryptedAppId,
			AppSecret: encryptedAppSecret,
		}
	}

	return startReqBody
}
