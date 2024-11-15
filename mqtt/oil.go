package mqtt

import "fmt"

const (
	topicImageCache  = "device/%d/image/cache"  // 开启摄像机预览推流
	topicImageUpload = "device/%d/image/upload" // 关闭摄像机预览推流
	topicImageClear  = "device/%d/image/clear"  // 关闭摄像机预览推流
)

func GetTopicImageCache(deviceId uint64) string {
	return fmt.Sprintf(topicImageCache, deviceId)
}

func GetTopicImageUpload(deviceId uint64) string {
	return fmt.Sprintf(topicImageUpload, deviceId)
}

func GetTopicImageClear(deviceId uint64) string {
	return fmt.Sprintf(topicImageClear, deviceId)
}

func GetTopicImageClearFinish(deviceID uint64) string {
	return fmt.Sprintf("device/%d/image/clear/finish", deviceID)
}
