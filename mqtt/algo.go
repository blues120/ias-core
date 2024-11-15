package mqtt

import (
	"fmt"
)

const (
	topicAlgoInstall   = "server/algo/%d/install"   // 算法包安装
	topicAlgoUninstall = "server/algo/%d/uninstall" // 算法包卸载
	topicAlgoProcess   = "server/algo/process"      //算法包进度
)

const (
	statusSuccess    = "success"
	statusFailure    = "fail"
	statusInProgress = "processing"
)

func GetStatusSuccess() string {
	return fmt.Sprintf(statusSuccess)
}
func GetStatusFailure() string {
	return fmt.Sprintf(statusFailure)
}
func GetStatusInProcess() string {
	return fmt.Sprintf(statusInProgress)
}

func GetTopicAlgoInstall(deviceID uint64) string {
	return fmt.Sprintf(topicAlgoInstall, deviceID)
}

func GetTopicAlgoProgress() string {
	return fmt.Sprintf(topicAlgoProcess)
}

type AlgoInstallRequest struct {
	FileUrl string `json:"fileUrl"`
	Image   string `json:"image"`
	FileId  uint64 `json:"fileId"`
	Meta    string `json:"meta"`
}

func GetTopicAlgoUninstall(deviceID uint64) string {
	return fmt.Sprintf(topicAlgoUninstall, deviceID)
}

type AlgoUninstallRequest struct {
	AlgoGroupId uint64     `json:"algo_group_id"`
	AlgoItems   []AlgoItem `json:"algo_items"`
}

type AlgoItem struct {
	AlgoName string `json:"algo_name"`
	Provider string `json:"provider"`
}

type InstallProcessInfo struct {
	DeviceID uint64 `json:"deviceID"`
	Progress uint64 `json:"progress"`
	FileID   uint64 `json:"fileID"`
	Status   string `json:"status"`
	Reason   string `json:"reason"`
}
