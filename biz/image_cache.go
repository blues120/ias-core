package biz

import "context"

type ImageCache struct {
	DeviceID  string `json:"device_id"`  // 设备id（云边模式下需要）
	ImageID   string `json:"image_id"`   // 图片id（与imageUrl相同）
	TaskID    string `json:"task_id"`    // 任务id
	CameraID  string `json:"camera_id"`  // 摄像机id
	GroupName string `json:"group_name"` // 分组名称（一个handler为一组）
	Refresh   bool   `json:"refresh"`    // 是否删除旧缓存图片
}

type ImageUpload struct {
	DeviceID  string `json:"device_id"`  // 设备id（云边模式下需要）
	ImageID   string `json:"image_id"`   // 图片id（与imageUrl相同）
	TaskID    string `json:"task_id"`    // 任务id
	CameraID  string `json:"camera_id"`  // 摄像机id
	GroupName string `json:"group_name"` // 分组名称（一个handler为一组）
}

type ImageClear struct {
	DeviceID  string `json:"device_id"`  // 设备id（云边模式下需要）
	TaskID    string `json:"task_id"`    // 任务id
	CameraID  string `json:"camera_id"`  // 摄像机id
	GroupName string `json:"group_name"` // 分组名称（一个handler为一组）
}

type ImageCacheRepo interface {
	// CacheImage 缓存图片
	// refresh 为 true 时，会删除旧缓存图片，只保留当前图片
	CacheImage(ctx context.Context, cache *ImageCache) error

	// UploadImage 上传图片
	UploadImage(ctx context.Context, cache *ImageUpload) error

	// ClearImage 清除图片
	ClearImage(ctx context.Context, cache *ImageClear) error
}
