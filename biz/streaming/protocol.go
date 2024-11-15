package streaming

import (
	"context"
	"fmt"
	"time"
)

// Protocol 流媒体协议
type Protocol interface {
	// Source 流媒体协议地址
	Source() string

	// Type 流媒体协议类型
	Type() ProtocolType

	// IsOnline 是否在线
	IsOnline(context.Context, time.Duration) (bool, error)

	// GetSnapshot 获取快照
	GetSnapshot(context.Context) ([]byte, error)

	// GetInfo 获取流信息
	// Deprecated: use GetInfoWithTimeout instead
	GetInfo(ctx context.Context) (*Info, error)

	// GetInfoWithTimeout 获取流信息，超时时间为 timeout
	GetInfoWithTimeout(ctx context.Context, timeout time.Duration) (*Info, error)
}

func NewProtocol(addr string, pt ProtocolType) (Protocol, error) {
	switch pt {
	case ProtocolTypeRtsp:
		return NewRtsp(addr)
	case ProtocolTypeOnvif:
		return NewOnvif(addr)
	case ProtocolTypeOfflineVideo:
		return NewOfflineVideo(addr)
	case ProtocolTypeGB28181:
		return NewGb28181(addr)
	default:
		return nil, fmt.Errorf("invalid protocol type: %s", pt)
	}
}

// ProtocolType 流媒体协议类型
type ProtocolType string

const (
	ProtocolTypeRtsp         ProtocolType = "rtsp"
	ProtocolTypeRtmp         ProtocolType = "rtmp"
	ProtocolTypeOnvif        ProtocolType = "onvif"
	ProtocolTypeOfflineVideo ProtocolType = "offline-video" // 离线视频
	ProtocolTypeGB28181      ProtocolType = "gb28181"       // 国标GB28181视频
)

// Values .
// 注意：Type 类型的常量必须在此处定义
func (ProtocolType) Values() []string {
	return []string{
		string(ProtocolTypeRtsp),
		string(ProtocolTypeRtmp),
		string(ProtocolTypeOnvif),
		string(ProtocolTypeOfflineVideo),
		string(ProtocolTypeGB28181),
	}
}
