package biz

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewCameraUsecase,
	NewTaskUsecase,
	NewAuthUsecase,
	NewPushUsecase,
	NewInformUsecase,
	NewOssUsecase,
	NewSmsNotifyUsecase,
	NewWarningAlertUsecase,
	NewSystemUsecase,
	NewEventSubsUsecase,
	NewSignatureUsecase,
	NewDeviceUsecase,
	NewMqttUsecase,
	NewUpPlatformUseCase,
	NewFileUploadUsecase,
	NewAlgoUsecase,
	NewUsecase,
	NewAlgoManagementUsecase,
	NewWarnTypeUsecase,
	NewVssUsecase,
	NewDeviceAlgoUsecase,
	NewOrganizationUsecase,
	NewAreaUsecase,
	NewTasklimitsUsecase,
	NewK8sMetricsUsecase,
)

// Transaction 事务封装
type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type Usecase struct {
	oss          *OssUsecase
	fileUpload   *FileUploadUsecase
	camera       *CameraUsecase
	task         *TaskUsecase
	algo         *AlgoUsecase
	auth         *AuthUsecase
	warningAlert *WarningAlertUsecase
}

func NewUsecase(
	oss *OssUsecase,
	fileUpload *FileUploadUsecase,
	camera *CameraUsecase,
	task *TaskUsecase,
	algo *AlgoUsecase,
	auth *AuthUsecase,
	warningAlert *WarningAlertUsecase,
) *Usecase {
	uc := &Usecase{
		oss:          oss,
		fileUpload:   fileUpload,
		camera:       camera,
		task:         task,
		algo:         algo,
		auth:         auth,
		warningAlert: warningAlert,
	}

	// tricks
	// 手动注入uc，允许uc间相互调用
	// camera.uc = uc

	return uc
}
