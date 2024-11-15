package agent

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
)

func TaskToAgentConfig(bc *conf.Bootstrap, ta *biz.Task, agentConf *conf.AgentExtend) *Config {
	mq := bc.Data.Mq
	taskId := GetTaskId(ta)

	config := &Config{
		Version: "v1.1",
		Algorithm: &Algorithm{
			Confidence: 0, // unused
		},
		Task: &Task{
			TaskId:     taskId,
			AlgoConfig: ta.AlgoConfig,
		},
		Agent: &Agent{
			Server: &AgentServer{
				Http: &AgentServerHttp{
					Addr:    "0.0.0.0:8082",
					Timeout: "5s",
				},
				Grpc: &AgentServerGrpc{
					Addr:    "0.0.0.0:9000",
					Timeout: "1s",
				},
				File: &AgentServerFile{
					OutputDir: "/data/meta",
					Timeout:   "6s",
				},
				RabbitMQ: &RabbitMQ{
					Addr:       mq.Url,
					Exchange:   "amq.topic",
					RoutingKey: fmt.Sprintf("%d", ta.Id),
					QueueName:  fmt.Sprintf("%d", ta.Id),
				},
			},
			Data: &AgentData{
				RabbitMQ: &RabbitMQ{
					Addr:         mq.Url,
					Exchange:     mq.ExchangeName,
					ExchangeType: mq.ExchangeType,
					RoutingKey:   ta.Algo.Algorithm.AppName,
					QueueName:    mq.QueueName,
				},
			},
		},
	}
	if agentConf != nil && len(agentConf.FileTimeOut) > 0 {
		config.Agent.Server.File.Timeout = agentConf.FileTimeOut
	}
	cameraMultiImgBox := make(map[uint64]string, 0)
	for _, ca := range ta.Cameras {
		if ca.MultiImgBox != "" && ca.MultiImgBox != "[]" { // 多区域场景
			cameraMultiImgBox[ca.Camera.Id] = ca.MultiImgBox
		}
	}

	sources := make([]*Source, 0, len(ta.Cameras))
	for _, tc := range ta.Cameras {
		camera := tc.Camera
		item := &Source{
			CameraID:          camera.Id,
			Source:            camera.StreamingProtocol.Source(),
			DropFrameInterval: float32(ta.Algo.Interval),
			Codec:             camera.StreamingInfo.CodecName,
			Width:             uint(camera.StreamingInfo.Width),
			Height:            uint(camera.StreamingInfo.Height),
		}
		if multiImgBox, has := cameraMultiImgBox[camera.Id]; has {
			if multiImgBox != "" {
				item.Events = ParseEvents(multiImgBox, item.Width, item.Height)
			}
		}
		sources = append(sources, item)
	}
	config.Sources = sources

	config.Agent.Data.Oss = &Oss{}
	switch cfg := bc.Data.Oss.Oss.(type) {
	case *conf.Data_Oss_Local_:
		config.Agent.Data.Oss.Local = &OssLocal{StorePath: cfg.Local.StorePath}
	case *conf.Data_Oss_AwsS3_:
		config.Agent.Data.Oss.S3 = &OssS3{
			AK:              cfg.AwsS3.Ak,
			SK:              cfg.AwsS3.Sk,
			Bucket:          cfg.AwsS3.Bucket,
			Endpoint:        cfg.AwsS3.Endpoint,
			Region:          cfg.AwsS3.Region,
			TrainDataBucket: cfg.AwsS3.TrainDataBucket,
		}
	}

	return config
}

// 取taskId
func GetTaskId(task *biz.Task) string {
	if task == nil {
		return ""
	}
	if task.ParentId != "" {
		return task.ParentId
	}
	return strconv.FormatUint(task.Id, 10)
}

func ParseEvents(eventStr string, width, height uint) []Event {
	if eventStr == "" {
		return nil
	}
	// 解析原始输入数据结构
	var input map[string][][]struct {
		X float64
		Y float64
	}
	if err := jsoniter.Unmarshal([]byte(eventStr), &input); err != nil {
		return nil
	}

	ret := make([]Event, 0)
	var detectDirections []int
	var targetLine []int
	var targetDirection []int

	for k, v := range input {
		if k == "targetArea" || k == "excludeArea" {
			var detectAreas []int
			detectAreasShape := []int{0}
			detectAreasShapeItemNum := 0
			for _, item := range v {
				for _, c := range item {
					detectAreas = append(detectAreas, int(c.X*float64(width)))
					detectAreas = append(detectAreas, int(c.Y*float64(height)))
				}
				detectAreasShapeItemNum++
				detectAreasShape = append(detectAreasShape, len(item))
			}
			detectAreasShape[0] = detectAreasShapeItemNum
			event := Event{
				DetectAreas:      detectAreas,
				DetectAreasShape: detectAreasShape,
			}
			if k == "targetArea" {
				event.Name = "in-filter"
			} else {
				event.Name = "out-filter"
			}
			ret = append(ret, event)
		} else if k == "targetLine" {
			for _, item := range v {
				for _, c := range item {
					targetLine = append(targetLine, int(c.X*float64(width)))
					targetLine = append(targetLine, int(c.Y*float64(height)))
				}
			}
		} else if k == "targetDirection" {
			for _, item := range v {
				for _, c := range item {
					targetDirection = append(targetDirection, int(c.X*float64(width)))
					targetDirection = append(targetDirection, int(c.Y*float64(height)))
				}
			}
		}
	}

	// 说明是画线和方向的人流量检测算法
	if len(targetLine) > 0 {
		detectDirections = append(detectDirections, targetDirection...)
		detectDirections = append(detectDirections, targetLine...)
		event := Event{
			DetectDirections: detectDirections,
		}
		ret = append(ret, event)
	}

	return ret
}

func ParseEventsVSS(modeName string, eventStr string, width, height uint) []Event {
	if eventStr == "" {
		return []Event{
			{ModelName: modeName},
		}
	}
	// 解析原始输入数据结构
	var input map[string][][]struct {
		X float64
		Y float64
	}
	if err := jsoniter.Unmarshal([]byte(eventStr), &input); err != nil {
		return nil
	}

	ret := make([]Event, 0)
	var detectDirections []int
	var targetLine []int
	var targetDirection []int

	for k, v := range input {
		if k == "targetArea" || k == "excludeArea" {
			var detectAreas []int
			detectAreasShape := []int{0}
			detectAreasShapeItemNum := 0
			for _, item := range v {
				for _, c := range item {
					detectAreas = append(detectAreas, int(c.X*float64(width)))
					detectAreas = append(detectAreas, int(c.Y*float64(height)))
				}
				detectAreasShapeItemNum++
				detectAreasShape = append(detectAreasShape, len(item))
			}
			detectAreasShape[0] = detectAreasShapeItemNum
			event := Event{
				DetectAreas:      detectAreas,
				DetectAreasShape: detectAreasShape,
			}
			if k == "targetArea" {
				event.Name = "in-filter"
			} else {
				event.Name = "out-filter"
			}
			event.ModelName = modeName
			ret = append(ret, event)
		} else if k == "targetLine" {
			for _, item := range v {
				for _, c := range item {
					targetLine = append(targetLine, int(c.X*float64(width)))
					targetLine = append(targetLine, int(c.Y*float64(height)))
				}
			}
		} else if k == "targetDirection" {
			for _, item := range v {
				for _, c := range item {
					targetDirection = append(targetDirection, int(c.X*float64(width)))
					targetDirection = append(targetDirection, int(c.Y*float64(height)))
				}
			}
		}
	}

	// 说明是画线和方向的人流量检测算法
	if len(targetLine) > 0 {
		detectDirections = append(detectDirections, targetDirection...)
		detectDirections = append(detectDirections, targetLine...)
		event := Event{
			DetectDirections: detectDirections,
		}
		event.ModelName = modeName
		ret = append(ret, event)
	}

	return ret
}

// Config .
type Config struct {
	Version   string     `json:"version" yaml:"version"`     // 版本号
	Sources   []*Source  `json:"sources" yaml:"sources"`     // 流地址配置
	Algorithm *Algorithm `json:"algorithm" yaml:"algorithm"` // 算法配置
	Task      *Task      `json:"task" yaml:"task"`           // 任务配置
	Agent     *Agent     `json:"agent" yaml:"agent"`         // agent配置
}

// Source 流地址配置
type Source struct {
	Source            string   `json:"source" yaml:"source"`                           // 流地址
	CameraID          uint64   `json:"camera_id" yaml:"camera_id"`                     // 摄像机ID
	DropFrameInterval float32  `json:"drop_frame_interval" yaml:"drop_frame_interval"` // 跳帧频率 0~30
	Codec             string   `json:"codec" yaml:"codec"`                             // 视频流编解码器
	Width             uint     `json:"width" yaml:"width"`                             // 摄像机分辨率-宽
	Height            uint     `json:"height" yaml:"height"`                           // 摄像机分辨率-高
	Coords            []*Coord `json:"coord" yaml:"coord"`                             // roi坐标信息；方向信息
	Events            []Event  `json:"events" yaml:"events"`                           // 事件信息
}

type Coord [2]float64

type Event struct {
	Name                  string `json:"name" yaml:"name"`                                       // 任务ID
	DetectAreas           []int  `json:"detect_areas" yaml:"detect_areas,flow"`                  // 检测区域，将多个区域按顺序变成一维数组，第二个区域接在第一个区域后面
	DetectAreasShape      []int  `json:"detect_areas_shape" yaml:"detect_areas_shape,flow"`      // 检测区域形状表示方法，第一个数2表示有2个检测区域，第2个数字5表示第1个区域有5个点，第3个数字6表示有6个点
	DetectDirections      []int  `json:"detect_directions" yaml:"detect_directions,flow"`        // 检测区域的方向，一个方向为[x1,y1,x2,y2]，这里有2个区域对应2个方向
	DetectDirectionsCount int    `json:"detect_directions_count" yaml:"detect_directions_count"` // 检测区域的方向数量
	ModelName             string `json:"model_name" yaml:"model_name"`                           // 算法名
}

// Algorithm 算法配置
type Algorithm struct {
	Confidence float64 `json:"confidence" yaml:"confidence"` // 置信度阈值
}

// Task 任务配置
type Task struct {
	TaskId     string `json:"task_id" yaml:"task_id"`           // 任务Id
	AlgoConfig string `json:"algo_config"   yaml:"algo_config"` // 算法个性化配置
}

// Agent agent配置
type Agent struct {
	Server   *AgentServer `json:"server" yaml:"server"`
	Data     *AgentData   `json:"data" yaml:"data"`
	Srsurl   *AgentSrsurl `json:"srsurl" yaml:"srsurl"`
	Rtsp     *AgentRtsp   `json:"rtsp" yaml:"rtsp"`
	Udp      *AgentRtsp   `json:"udp" yaml:"udp"`
	Callback *Callback    `json:"callback" yaml:"callback"`
}

type AgentServer struct {
	Http     *AgentServerHttp `json:"http" yaml:"http"`
	Grpc     *AgentServerGrpc `json:"grpc" yaml:"grpc"`
	File     *AgentServerFile `json:"file" yaml:"file"`
	RabbitMQ *RabbitMQ        `json:"rabbitmq" yaml:"rabbitmq"`
}

type AgentServerHttp struct {
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

type AgentServerGrpc struct {
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

type AgentServerFile struct {
	OutputDir string `json:"output_dir" yaml:"output_dir"`
	Timeout   string `json:"timeout" yaml:"timeout"`
}

type AgentData struct {
	RabbitMQ *RabbitMQ `json:"rabbitmq" yaml:"rabbitmq"`
	Oss      *Oss      `json:"oss" yaml:"oss"` // OSS配置 one of (OssLocal, OssS3)
}

type RabbitMQ struct {
	Addr         string `json:"addr" yaml:"addr"`
	Exchange     string `json:"exchange" yaml:"exchange"`
	ExchangeType string `json:"exchange_type" yaml:"exchange_type"`
	RoutingKey   string `json:"routing_key" yaml:"routing_key"`
	QueueName    string `json:"queue_name" yaml:"queue_name"`
}

type Oss struct {
	Local *OssLocal `json:"local" yaml:"local"`
	S3    *OssS3    `json:"s3" yaml:"s3"`
}

// OssLocal 本地存储配置
type OssLocal struct {
	StorePath string `json:"store_path" yaml:"store_path"` // 本地存储地址
}

// OssS3 S3存储配置
type OssS3 struct {
	AK              string `json:"ak" yaml:"ak"`
	SK              string `json:"sk" yaml:"sk"`
	Bucket          string `json:"bucket" yaml:"bucket"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	Region          string `json:"region" yaml:"region"`
	TrainDataBucket string `json:"trainDataBucket" yaml:"trainDataBucket"`
}

type Callback struct {
	Addr               string            `json:"addr" yaml:"addr"`
	Signature          *Signature        `json:"signature" yaml:"signature"`
	Urls               map[string]string `json:"urls" yaml:"urls"`
	AuthEnable         bool              `json:"auth_enable" yaml:"auth_enable"`
	InsecureSkipVerify bool              `json:"insecure_skip_verify" yaml:"insecure_skip_verify"`
}

type Signature struct {
	AppId     string `json:"app_id" yaml:"app_id"`
	AppSecret string `json:"app_secret" yaml:"app_secret"`
}

type AgentSrsurl struct {
	RtmpUrl   string `json:"rtmp_url" yaml:"rtmp_url"`
	FlvUrl    string `json:"flv_url" yaml:"flv_url"`
	StreamUrl string `json:"stream_url" yaml:"stream_url"`
}

type AgentRtsp struct {
	Enable int32  `json:"enable" yaml:"enable"`
	Port   string `json:"port" yaml:"port"`
}
