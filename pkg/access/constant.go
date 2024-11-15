package access

import "time"

const (
	RequestSipTimeOut = 30 * time.Second
	StartPlay         = "/gb28181/v1/devices?Action=startPlay"
	StopPlay          = "/gb28181/v1/devices?Action=stopPlay"
	CascadeStatus     = "/cascadeGb28181/v1/upperPlatforms/status"
	CascadeRegister   = "/cascadeGb28181/v1/upperPlatforms?Action=register"
	CascadeUnRegister = "/cascadeGb28181/v1/upperPlatforms?Action=unregister"
	StartCascadePlay  = "/streamingserver/v1/gb28181/streams?Action=startCascadePlay"
	NonPullStream     = "rtsp://127.0.0.1:554/"
)
