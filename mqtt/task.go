package mqtt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

const (
	topicTaskLogFetch = "device/%d/task/log/fetch" // 任务日志查询
	topicTaskLogStop  = "device/%d/task/log/stop"  // 停止任务日志查询
	topicTaskLog      = "server/task/%s/log"       // 任务日志结果

	topicTaskStart = "device/%d/task/start" // 启动任务
	topicTaskStop  = "device/%d/task/stop"  // 停止任务
	topicTaskState = "server/task/state"    // 任务状态结果

	reasonTaskRunningFromRestart = "任务正在运行,服务从重启中恢复"
	reasonSyncTaskFromBox        = "从盒子侧同步最新任务状态"

	IboxTaskRedisPerfix = "ias:IBOX-TASK"
	TaskInfoRedisPerfix = IboxTaskRedisPerfix + "-INFO"
	TaskInfoListPerfix  = IboxTaskRedisPerfix + "-LIST"
)

func GetTopicTaskLogFetch(deviceID uint64) string {
	return fmt.Sprintf(topicTaskLogFetch, deviceID)
}

func GetTopicTaskLogStop(deviceID uint64) string {
	return fmt.Sprintf(topicTaskLogStop, deviceID)
}

func GetTopicTaskLog(taskID string) string {
	return fmt.Sprintf(topicTaskLog, taskID)
}

func GetTopicTaskStart(deviceID uint64) string {
	return fmt.Sprintf(topicTaskStart, deviceID)
}

func GetTopicTaskStop(deviceID uint64) string {
	return fmt.Sprintf(topicTaskStop, deviceID)
}

func GetTopicTaskState() string {
	return topicTaskState
}

func GetReasonTaskRunningFromRestart() string {
	return reasonTaskRunningFromRestart
}

func GetReasonSyncTaskFromBox() string {
	return reasonSyncTaskFromBox
}

// TaskLogFetch 任务日志查询
type TaskLogFetch struct {
	TaskID string `json:"task_id" validate:"required"` // 任务ID
	Mode   string `json:"mode"`                        // 任务模式docker or api
}

// TaskLogStop 停止任务日志查询
type TaskLogStop struct {
	TaskID string `json:"task_id" validate:"required"` // 任务ID
	Mode   string `json:"mode"`                        // 任务模式docker or api
}

// TaskLog 任务日志
type TaskLog string

// TaskStart 启动任务
type TaskStart struct {
	Mode        string    `json:"mode"`                     // 任务模式docker or api
	Agent       Agent     `json:"agent"`                    // agent相关配置
	Algorithm   Algorithm `json:"algorithm"`                // 算法相关配置
	CallbackURL string    `json:"callback_url"`             // 任务状态回调URL
	Sources     []Source  `json:"sources"`                  // 流相关配置
	Task        Task      `json:"task" validate:"required"` // 任务相关配置
	Version     string    `json:"version"`                  // 配置文件版本号
	Mq          RabbitMq  `json:"mq,omitempty"`             // mq配置
	Aws         Aws       `json:"aws,omitempty"`            // aws配置
}

// TaskStop 停止任务
type TaskStop struct {
	TaskID    string    `json:"id" validate:"required"` // 任务ID
	Mode      string    `json:"mode"`                   // 任务模式docker or api
	Algorithm Algorithm `json:"algorithm"`              // 算法相关配置
}

// TaskState 任务状态
type TaskState struct {
	TaskID string         `json:"id" validate:"required"`    // 任务ID
	State  biz.TaskStatus `json:"state" validate:"required"` // 任务状态
	Reason string         `json:"reason"`                    // 任务状态变更原因
}

type Signature struct {
	AppId     string `json:"app_id" yaml:"app_id"`
	AppSecret string `json:"app_secret" yaml:"app_secret"`
}

type AgentCallback struct {
	Addr               string            `json:"addr"`
	Signature          Signature         `json:"signature"`
	Urls               map[string]string `json:"urls"`
	AuthEnable         bool              `json:"auth_enable"`          // aksk/signature鉴权开关
	InsecureSkipVerify bool              `json:"insecure_skip_verify"` // 跳过证书验证
}

func (a AgentCallback) Valid() bool {
	return a.Addr != "" && len(a.Urls) > 0
}

// 图片视频持久化模式
type Media_Persistence_Mode int32

const (
	// 默认方式，表示上传到自有 OSS（local or aws）
	Media_Persistence_Mode_Default Media_Persistence_Mode = 0
	// 数生方式，表示上传到数生资源池
	Media_Persistence_Mode_Vss Media_Persistence_Mode = 1
)

// Agent agent相关配置
type Agent struct {
	Server               Server                 `json:"server"`
	Callback             AgentCallback          `json:"callback"`
	MediaPersistenceMode Media_Persistence_Mode `json:"media_persistence_mode"` // 资源持久化方式，默认 0 表示上传到自有云存储（local or aws）, 1 表示上传到数生资源池
}

// Server agent server相关配置
type Server struct {
	File File `json:"file"`
}

// File agent server file相关配置
type File struct {
	DataOutputDir string `json:"data_output_dir"` // 告警图片视频的目录，默认地址 /data/data
	MetaOutputDir string `json:"meta_output_dir"` // 告警结果json目录，默认地址 /data/meta
}

// Algorithm 算法相关配置
type Algorithm struct {
	Confidence   float64                `json:"confidence"`            // 置信度阈值
	Config       map[string]interface{} `json:"config"`                // 算法个性化配置
	ID           string                 `json:"id"`                    // 算法场景ID
	ModelNames   []string               `json:"model_names,omitempty"` // 算法模型名称
	Mode         string                 `json:"mode"`                  // 取流方式，实时(realtime):当预警任务开始分析时，会持续抓取视频流;截图(capture):根据设置的取流间隔时间，每隔指定时间抓取一次摄像头画面做分析。取流间隔要大于10秒。一般针对无需实时分析，预警触发条件和周期大于10秒以上的场景
	Period       int64                  `json:"period"`                // 预警周期，当触发预警后，特定目标持续在摄像头中出现，会每隔预警周期时间长度产生一次预警
	Trigger      int64                  `json:"trigger"`               // 触发时间，特定目标需要在摄像头中出现的时间到达触发时间后才会产生预警
	Version      string                 `json:"version"`               // 算法场景版本
	Image        string                 `json:"image,omitempty"`       // 算法镜像地址
	AppName      string                 `json:"appName,omitempty"`     // RabbitMQ的RoutingKey
	Prefix       string                 `json:"prefix,omitempty"`      // 算法包服务启动api前缀 ip:port
	CustomConfig map[string]interface{} `json:"custom_config"`         // telestream api算法个性化配置
}

// Source 视频流相关配置
type Source struct {
	CameraID          string  `json:"camera_id"`           // 摄像头ID
	Codec             *string `json:"codec"`               // 视频流编解码器
	DropFrameInterval int64   `json:"drop_frame_interval"` // 跳帧频率 0~30
	Events            []Event `json:"events"`              // 检测区域为空时，默认全屏分析
	Areas             []Area  `json:"areas"`               // telestream api规范
	Lines             []Line  `json:"lines"`               // telestream api规范
	Height            *int64  `json:"height"`              // 分辨率长
	Source            string  `json:"source"`              // 流地址
	Width             *int64  `json:"width"`               // 分辨率宽
	SourceType        string  `json:"source_type"`         // RTSP或VIDEO
}

// Event 检测区域
type Event struct {
	DetectAreas           []int64 `json:"detect_areas"`            // 检测区域，将多个区域按顺序变成一维数组，第二个区域接在第一个区域后面
	DetectAreasShape      []int64 `json:"detect_areas_shape"`      // 检测区域形状表示方法，第一个数2表示有2个检测区域，第2个数字5表示第1个区域有5个点，第3个数字6表示有6个点
	DetectDirections      []int64 `json:"detect_directions"`       // 检测区域的方向，一个方向为 [x1,y1,x2,y2]，这里有2个区域对应2个方向
	DetectDirectionsCount *int64  `json:"detect_directions_count"` // 检测区域方向的个数
	DetectLines           []int64 `json:"detect_lines"`            // 检测线 可选 例如: [x1,y1,x2,y2]，检测线不表示检测方向, 检测方向根据detect_directions字段确定
	DetectLinesCount      *int64  `json:"detect_lines_count"`      // 检测线的个数
	Name                  *string `json:"name"`                    // 检测区域名称
}

type Area struct {
	ID    string    `json:"id"`
	Coord []float64 `json:"coord"`
}

type Line struct {
	ID        string    `json:"id"`
	Coord     []float64 `json:"coord"`
	Direction []float64 `json:"direction"`
}

// Task 任务相关配置
type Task struct {
	TaskID        string       `json:"id"`             // 任务id
	WorkspaceID   string       `json:"workspace_id"`   // 工作区id
	Period        string       `json:"period"`         // 告警周期
	ControlPeriod string       `json:"control_period"` // 布控时段
	Sources       string       `json:"sources"`
	Type          biz.TaskType `json:"type"`         // 任务类型
	AlgoConfig    string       `json:"algo_config"`  // 算法个性化配置
	CallbackURL   string       `json:"callback_url"` // 任务状态回调URL
}

// RabbitMq mq配置
type RabbitMq struct {
	Url string `json:"url,omitempty"` // mq的url
}

func (r RabbitMq) Valid() bool {
	return r.Url != ""
}

// Aws aws配置
type Aws struct {
	Ak       string `json:"ak,omitempty"`       // aws ak
	Sk       string `json:"sk,omitempty"`       // aws sk
	Bucket   string `json:"bucket,omitempty"`   // aws bucket
	Endpoint string `json:"endpoint,omitempty"` // aws endpoint
	Region   string `json:"region,omitempty"`   // aws region
}

func (a Aws) Valid() bool {
	return a.Ak != "" && a.Sk != "" && a.Bucket != "" && a.Endpoint != "" && a.Region != ""
}

// GenerateTaskInfoKey 生成taskinfo的redis key
func GenerateTaskInfoKey(camID, modelName string) string {
	return fmt.Sprintf("%s:%s:%s", TaskInfoRedisPerfix, camID, modelName)
}

// GenerateTaskInfoListKey 生成taskinfo的redis list key
func GenerateTaskInfoListKey(taskID string) string {
	return fmt.Sprintf("%s:%s", TaskInfoListPerfix, taskID)
}

// GetTaskInfoFromRedis 从redis获取taskinfo
func GetTaskInfoFromRedis(rdb *redis.Client, camID, modelName string) (map[string]string, error) {
	key := GenerateTaskInfoKey(camID, modelName)
	taskInfo, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	task := make(map[string]string)
	err = json.Unmarshal([]byte(taskInfo), &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// SaveTaskInfoToRedis 获取taskinfo, 保存到redis
func SaveTaskInfoToRedis(rdb *redis.Client, data TaskStart) error {
	taskInfo, err := json.Marshal(data.Task)
	if err != nil {
		return err
	}

	if len(data.Algorithm.ModelNames) == 0 {
		return fmt.Errorf("modelNames is empty")
	}

	taskInfoKeys := make([]string, 0)
	// 生成redis key, 保存taskinfo到redis
	for _, d := range data.Sources {
		for _, m := range data.Algorithm.ModelNames {
			key := GenerateTaskInfoKey(d.CameraID, m) // 保存每一个camera和model组合的taskinfo
			_, err := rdb.Set(context.Background(), key, taskInfo, 0).Result()
			if err != nil {
				return err
			}
			taskInfoKeys = append(taskInfoKeys, key)
		}
	}

	// 保存taskinfo key list到redis
	return SaveTaskInfoListToRedis(rdb, data.Task.TaskID, taskInfoKeys)
}

// DeleteTaskInfoFromRedis 删除redis中的taskinfo
func DeleteTaskInfoFromRedis(rdb *redis.Client, camID, modelName string) error {
	key := GenerateTaskInfoKey(camID, modelName)
	return rdb.Del(context.Background(), key).Err()
}

// CleanTaskInfoFromRedis 清理redis中的taskinfo缓存
func CleanTaskInfoFromRedis(rdb *redis.Client) error {
	// 从redis中获取符合前缀的所有key
	ctx := context.Background()
	keysToDelete, err := rdb.Keys(ctx, IboxTaskRedisPerfix+"*").Result()
	if err != nil {
		return err
	}
	if len(keysToDelete) == 0 {
		return nil
	}
	// 删除所有符合前缀的key
	return rdb.Del(ctx, keysToDelete...).Err()
}

// SaveTaskInfoListToRedis 保存taskinfo list到redis
func SaveTaskInfoListToRedis(rdb *redis.Client, taskID string, taskInfoKeysList []string) error {
	// taskID => taskinfo redis key list
	listKey := GenerateTaskInfoListKey(taskID)
	value, err := json.Marshal(taskInfoKeysList)
	if err != nil {
		return err
	}
	return rdb.Set(context.Background(), listKey, value, 0).Err()
}

// GetTaskInfoListFromRedis 从redis获取taskinfo list
func GetTaskInfoListFromRedis(rdb *redis.Client, taskID string) ([]string, error) {
	taskInfoKeysList := make([]string, 0)
	key := GenerateTaskInfoListKey(taskID)
	taskInfoKeys, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil { // 未找到key，返回空list
			return taskInfoKeysList, nil
		}
		return nil, err
	}
	// convert string to []string
	err = json.Unmarshal([]byte(taskInfoKeys), &taskInfoKeysList)
	if err != nil {
		return nil, err
	}
	return taskInfoKeysList, nil
}

// DeleteTaskInfoListFromRedis 删除redis中的taskinfo list
func DeleteTaskInfoListFromRedis(rdb *redis.Client, taskID string) error {
	key := GenerateTaskInfoListKey(taskID)
	return rdb.Del(context.Background(), key).Err()
}

// DeleteTaskInfoAndListFromRedis 同时删除任务关联的taskinfo和taskinfo list
func DeleteTaskInfoAndListFromRedis(rdb *redis.Client, taskID string) error {
	// 获取taskinfo list
	taskInfoKeysList, err := GetTaskInfoListFromRedis(rdb, taskID)
	if err != nil {
		return err
	}
	for _, key := range taskInfoKeysList {
		_, err := rdb.Del(context.Background(), key).Result()
		if err != nil {
			return err
		}
	}
	// 最后删除list key
	return DeleteTaskInfoListFromRedis(rdb, taskID)
}
