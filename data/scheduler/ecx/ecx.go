package ecx

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/scheduler/agent"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/openapi"
)

type scheduler struct {
	bc     *conf.Bootstrap
	db     *ent.Database
	signUc *biz.SignatureUsecase
	log    *log.Helper
}

func NewScheduler(bc *conf.Bootstrap, logger log.Logger) biz.SchedulerRepo {
	if bc == nil || logger == nil {
		return nil
	}

	return &scheduler{
		bc:  bc,
		db:  &ent.Database{}, // TODO 是否需要使用db
		log: log.NewHelper(logger),
	}
}

// Close 关闭
func (r *scheduler) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

// Start 启动任务
func (r *scheduler) Start(ctx context.Context, ta *biz.Task) error {
	// 启动任务
	startReqBody := r.GenerateStartBody(ta, r.bc.AgentExtend)
	payload, err := json.Marshal(startReqBody)
	r.log.Debugf("StartVSS:%+v, %s\n", startReqBody, string(payload))
	openClient := r.newOpenApiClient()
	rsp, err := openClient.Post("/api/v1/kube/crd/itask/start/"+strconv.Itoa(int(ta.Id)), bytes.NewBuffer(payload))
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

	return nil
}

// Stop 停止任务
func (r *scheduler) Stop(ctx context.Context, ta *biz.Task) error {
	// 查询盒子状态

	// 停止任务
	openClient := r.newOpenApiClient()
	rsp, err := openClient.Post("/api/v1/kube/crd/itask/stop/"+strconv.Itoa(int(ta.Id)), nil)
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

// GetLog 查询任务日志
func (r *scheduler) GetLog(ctx context.Context, ta *biz.Task, conn *websocket.Conn) error {
	algoInfo := ta.Algo.Algorithm
	modelName := getModelName(algoInfo.Image)
	if modelName == "" {
		return errors.ErrorInvalidParam("algo image required")
	}

	openClient := r.newOpenApiClient()

	// 当任务没有启动时候，额外查看镜像状态
	if ta.Status == biz.TaskStatusInitializing || ta.Status == biz.TaskStatusStarting ||
		ta.Status == biz.TaskStatusFailed {
		resp, err := openClient.Get("/api/v1/kube/deployment/" + modelName + "/" + strconv.Itoa(int(ta.Id)))
		if err == nil {
			err = conn.WriteMessage(websocket.TextMessage, []byte(resp+"\n\n"))
		}
	}

	// 获取日志
	_ = conn.WriteMessage(websocket.TextMessage, []byte(getJobName(ta.Id)+"任务日志:\n"))
	eventStream, err := openClient.SSEGet("/api/v1/kube/crd/itask/logs/" + modelName + "/" + strconv.Itoa(int(ta.Id)))
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

// GetStatus 查询任务状态
func (r *scheduler) GetStatuses(ctx context.Context, taskIdList []uint) ([]biz.TaskIdStatus, error) {
	//
	panic("implement me")
}

func (r *scheduler) newOpenApiClient() *openapi.OpenClient {
	return openapi.NewOpenClient(r.bc.OpenApi.Host, r.bc.OpenApi.AppId, r.bc.OpenApi.AppSecret, r.bc.OpenApi.Ac, "", r.log)
}

// 获取模型名称
func getModelName(fullImageName string) string {
	// registry-ibox.ctcdn.cn:5000/ai-service/ibox-trash:0.1.4-ds6.0-xaviernx
	if fullImageName == "" {
		return ""
	}

	parts := strings.Split(fullImageName, ":")
	var imageName string
	if len(parts) == 1 { // image就是event-analysis的场景
		imageName = parts[0]
	} else {
		imageName = parts[len(parts)-2] // 去掉最后的tag
	}
	parts = strings.Split(imageName, "/")
	return parts[len(parts)-1]
}

// 根据任务id，返回job名称
func getJobName(jobId uint64) string {
	return fmt.Sprintf("job%d", jobId)
}

type Decoder struct {
	src *bufio.Reader
}

func NewDecoder(src io.Reader) *Decoder {
	return &Decoder{src: bufio.NewReader(src)}
}

// EcxApiResp api返回结果解析
type EcxApiResp struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
	Data   struct {
		Status string `json:"status"`
		TaskID string `json:"taskId"`
	} `json:"data"`
}

type StartReqBody struct {
	DATA StartReqBodyData `json:"data" yaml:"data"`
}

type StartReqBodyData struct {
	Sources   []StartReqBodySource    `json:"sources" yaml:"sources"`     // 视频流数组
	Algorithm StartReqBodyAlgorithm   `json:"algorithm" yaml:"algorithm"` // 算法配置
	Task      StartReqBodyTaskProfile `json:"task" yaml:"task"`           // 任务服务信息
	Box       StartReqBodyBox         `json:"box" yaml:"box"`             // 盒子配置
	Agent     agent.Agent             `json:"agent" yaml:"agent"`
}

type StartReqBodySource struct {
	Source string `json:"source" yaml:"source"` // 流地址
	// AccessToken       string  `json:"access_token" yaml:"access_token"`            // 视频流认证(用户名:密码) des3加密,密钥：cameraId
	CameraID          uint    `json:"camera_id" yaml:"camera_id"`                     // 摄像机ID
	DropFrameInterval float32 `json:"drop_frame_interval" yaml:"drop_frame_interval"` // 帧率
	Codec             string  `json:"codec" yaml:"codec"`                             // 视频流编解码器
	Width             int     `json:"width" yaml:"width"`                             // 摄像头分辨率-宽
	Height            int     `json:"height" yaml:"height"`                           // 摄像头分辨率-高
	// Coords            []*Coord `json:"coord" yaml:"coord"`                         // roi坐标信息；方向信息
	Events []agent.Event `json:"events" yaml:"events"` // 事件信息
}

type StartReqBodyAlgorithm struct {
	Confidence float64 `json:"confidence" yaml:"confidence"` // 置信度阈值
}

type StartReqBodyTaskProfile struct {
	TaskID      string `json:"task_id" yaml:"task_id"`           // 任务ID
	BoxID       string `json:"box_id" yaml:"box_id"`             // 盒子id， 盒子纳管生成的boxId
	ZoneID      string `json:"zone_id" yaml:"zone_id"`           // 盒子区域id
	WorkspaceID string `json:"workspace_id" yaml:"workspace_id"` // 工作区id
	EvtInfo     string `json:"evt_info" yaml:"evt_info"`         // 事件消息
	ModelName   string `json:"model_name" yaml:"model_name"`     // 模型名称
	ModelID     string `json:"model_id" yaml:"model_id"`         // 模型id
}

type StartReqBodyBox struct {
	Hostname      string `json:"hostname" yaml:"hostname"`               // //盒子mac地址
	TaskImageName string `json:"task_image_name" yaml:"task_image_name"` // 任务镜像名称
	TaskVersion   string `json:"task_version" yaml:"task_version"`       // 任务镜像tag
}

type Srsurl struct {
	RtmpUrl   string
	FlvUrl    string
	StreamUrl string
}

func (r *scheduler) GenerateStartBody(meta *biz.Task, agentConf *conf.AgentExtend) StartReqBody {
	mq := r.bc.Data.Mq
	var srsurl Srsurl
	err := json.Unmarshal([]byte(meta.Extend), &srsurl)
	if err != nil {
		r.log.Errorf("unmarshal Srsurl failed")
	}
	startReqBody := StartReqBody{
		DATA: StartReqBodyData{
			Sources: make([]StartReqBodySource, 0),
			Algorithm: StartReqBodyAlgorithm{
				Confidence: 0,
			},
			Task:  StartReqBodyTaskProfile{},
			Box:   StartReqBodyBox{},
			Agent: agent.Agent{},
		},
	}

	for _, taskCamera := range meta.Cameras {
		camera := taskCamera.Camera

		width := int(camera.StreamingInfo.Width)
		height := int(camera.StreamingInfo.Height)

		source := StartReqBodySource{
			Source:            camera.StreamingProtocol.Source(),
			CameraID:          uint(camera.Id),
			DropFrameInterval: float32(meta.Algo.Interval),
			Codec:             camera.StreamingInfo.CodecName,
			Width:             width,
			Height:            height,
			Events:            agent.ParseEvents(taskCamera.MultiImgBox, uint(width), uint(height)),
		}
		startReqBody.DATA.Sources = append(startReqBody.DATA.Sources, source)
	}

	algo := meta.Algo.Algorithm
	box := meta.Device
	startReqBody.DATA.Task = StartReqBodyTaskProfile{
		TaskID:      strconv.Itoa(int(meta.Id)),
		BoxID:       box.ExtId,
		ZoneID:      box.ZoneId,
		WorkspaceID: box.WorkspaceId,
		EvtInfo:     "启动任务",
		ModelName:   getModelName(algo.Image),
		ModelID:     strconv.Itoa(int(algo.ID)),
	}

	// 获取镜像名称和版本
	imageName, imageVersion := getImageNameAndVersion(algo.Image)
	startReqBody.DATA.Box = StartReqBodyBox{
		Hostname:      box.Mac + ".ctyun.cn",
		TaskImageName: imageName,
		TaskVersion:   imageVersion,
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
				Timeout:   "6s",
			},
			RabbitMQ: &agent.RabbitMQ{
				Addr:       mq.Url,
				Exchange:   "amq.topic",
				RoutingKey: fmt.Sprintf("%d", meta.Id),
				QueueName:  fmt.Sprintf("%d", meta.Id),
			},
		},
		Callback: &agent.Callback{
			Addr:               r.bc.Callback.Addr,
			Urls:               map[string]string{},
			AuthEnable:         r.bc.Callback.AuthEnable,
			InsecureSkipVerify: true,
		},
		Srsurl: &agent.AgentSrsurl{
			RtmpUrl:   srsurl.RtmpUrl,
			FlvUrl:    srsurl.FlvUrl,
			StreamUrl: srsurl.StreamUrl,
		},
		Rtsp: &agent.AgentRtsp{
			Enable: 1,
			Port:   "8554",
		},
		Udp: &agent.AgentRtsp{
			Enable: 1,
			Port:   fmt.Sprintf("%d", 20000+meta.Id), // 盒子之间udp端口隔离开,  防止广播串流
		},
		Data: &agent.AgentData{
			RabbitMQ: &agent.RabbitMQ{
				Addr:       mq.Url,
				Exchange:   "amq.topic",
				RoutingKey: meta.Algo.Algorithm.AppName,
			},
		},
	}
	if agentConf != nil && len(agentConf.FileTimeOut) > 0 {
		startReqBody.DATA.Agent.Server.File.Timeout = agentConf.FileTimeOut
	}
	startReqBody.DATA.Agent.Data.Oss = &agent.Oss{}
	switch cfg := r.bc.Data.Oss.Oss.(type) {
	case *conf.Data_Oss_Local_:
		startReqBody.DATA.Agent.Data.Oss.Local = &agent.OssLocal{StorePath: cfg.Local.StorePath}
	case *conf.Data_Oss_AwsS3_:
		startReqBody.DATA.Agent.Data.Oss.S3 = &agent.OssS3{
			AK:       cfg.AwsS3.Ak,
			SK:       cfg.AwsS3.Sk,
			Bucket:   cfg.AwsS3.Bucket,
			Endpoint: cfg.AwsS3.Endpoint,
			Region:   cfg.AwsS3.Region,
		}
	}

	for k, v := range r.bc.Callback.Urls {
		startReqBody.DATA.Agent.Callback.Urls[k] = v
	}
	if r.bc.Callback.AuthEnable && box.ExtId != "" {
		// 获取AppId/AppSecret并加密
		sign, err := r.signUc.FindByCondition(context.Background(), box.ExtId, "")
		if err != nil {
			r.log.Error("GenerateStartBody: signUc.FindByBoxId error", err)
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

func getImageNameAndVersion(fullImageName string) (imageName string, imageVersion string) {
	// 101.227.95.9:8025/event-analysis:0.0.1a
	parts := strings.Split(fullImageName, ":")
	if len(parts) < 2 {
		return "", ""
	}
	return strings.Join(parts[0:len(parts)-1], ":"), parts[len(parts)-1]
}
