package mqtt

import (
	"fmt"

	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz/streaming"
)

const (
	topicCameraPreviewOpen  = "device/%d/camera/preview/open"  // 开启摄像机预览推流
	topicCameraPreviewClose = "device/%d/camera/preview/close" // 关闭摄像机预览推流

	topicCameraSnapshot         = "device/%d/camera/snapshot" // 生成摄像头截图
	topicCameraSnapshotCallback = "server/camera/snapshot"    // 获取摄像头截图

	topicCameraInfoSync = "device/%d/camera/info" // 摄像机状态同步
	topicCameraInfo     = "server/camera/info"    // 摄像机状态

	topicOnvifSearch = "device/%d/onvif/search" // 搜索onvif摄像机
	topicOnvifInfo   = "server/onvif/info"      // onvif摄像机信息

	topicStreamStatus     = "device/%d/stream/status"   // 测试摄像机流连通性
	topicStreamStatusInfo = "server/stream/status/info" // 摄像机流连通性信息
)

func GetTopicCameraPreviewOpen(deviceId uint64) string {
	return fmt.Sprintf(topicCameraPreviewOpen, deviceId)
}

func GetTopicCameraPreviewClose(deviceId uint64) string {
	return fmt.Sprintf(topicCameraPreviewClose, deviceId)
}

func GetTopicCameraSnapshot(deviceId uint64) string {
	return fmt.Sprintf(topicCameraSnapshot, deviceId)
}

func GetTopicCameraSnapshotCallback() string {
	return topicCameraSnapshotCallback
}

func GetTopicCameraInfoSync(deviceID uint64) string {
	return fmt.Sprintf(topicCameraInfoSync, deviceID)
}

func GetTopicCameraInfo() string {
	return topicCameraInfo
}

func GetTopicOnvifSearch(deviceID uint64) string {
	return fmt.Sprintf(topicOnvifSearch, deviceID)
}

func GetTopicOnvifInfo() string {
	return topicOnvifInfo
}

func GetTopicStreamStatus(deviceID uint64) string {
	return fmt.Sprintf(topicStreamStatus, deviceID)
}

func GetTopicStreamStatusInfo() string {
	return topicStreamStatusInfo
}

type CameraSnapshot struct {
	DeviceID   uint64
	CameraID   uint64
	RtspSource string
}

type CameraSnapshotCallback struct {
	DeviceID uint64
	CameraID uint64
	Image    string
	Width    int32
	Height   int32
}

type CameraInfoSync struct {
	CameraID uint64
	Source   string
	Type     streaming.ProtocolType
	State    biz.CameraStatus
	WithInfo bool
}

type CameraInfo struct {
	CameraID uint64
	State    biz.CameraStatus
	Info     *streaming.Info
}

type CameraOnvifSearchCallback struct {
	DeviceID     uint64            `json:"deviceID,omitempty"`
	OnvifCameras []biz.OnvifCamera `json:"onvifCameras,omitempty"`
	Err          string            `json:"err,omitempty"`
}

type CameraStreamStatus struct {
	DeviceID     uint64 `json:"deviceID,omitempty"`
	CameraID     uint64 `json:"cameraID,omitempty"`
	CameraSource string `json:"cameraSource,omitempty"`
	IsOnline     bool   `json:"isOnline,omitempty"`
	Err          string `json:"err,omitempty"`
}
