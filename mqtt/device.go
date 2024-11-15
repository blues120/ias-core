package mqtt

import (
	"fmt"
)

const (
	topicDeviceStatus        = "server/device/%d/status"    // 设备状态
	topicLicenseAuthDeadline = "server/device/authDeadline" // licence授权时间
	topicDeviceSyncAlgoList  = "server/device/syncAlgoList" // 算法列表同步
	topicDeviceUpdateName    = "device/%d/updateName"       // 更新盒子设备名
)

type LicenseAuthDeadline struct {
	DeviceID     uint64
	AuthDeadline string
}

type DeviceAlgoList struct {
	DeviceID uint64
	AlgoList []*DeviceAlgoInfo // 算法列表
}

type DeviceAlgoInfo struct {
	AlgoGroupId      uint   // 算法组ID
	AlgoGroupName    string // 算法组name
	AlgoGroupVersion string // 算法组版本号，用于算法组整体更新场景
	Name             string // 算法名称
	Version          string // 算法版本号
	InstallTime      string // 安装时间
}

type DeviceNameUpdate struct {
	Name string // 设备名字
}

func GetTopicDeviceStatus(deviceID uint64) string {
	return fmt.Sprintf(topicDeviceStatus, deviceID)
}

func GetTopicLicenseAuthDeadline() string {
	return topicLicenseAuthDeadline
}

func GetTopicDeviceSyncAlgoList() string {
	return topicDeviceSyncAlgoList
}

func GetTopicDeviceNameUpdate(deviceID uint64) string {
	return fmt.Sprintf(topicDeviceUpdateName, deviceID)
}
