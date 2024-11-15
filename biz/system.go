package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	EquipAttrKeyDeviceId     = "device_id"
	EquipAttrKeyMqttAddr     = "mqtt_addr"
	EquipAttrKeyMqttUsername = "mqtt_username"
	EquipAttrKeyMqttPassword = "mqtt_password"
)

type EquipAttr struct {
	Id        uint64
	AttrKey   string
	AttrValue string
	Extend    string
}

type ActiveInfo struct {
	Id        uint64
	ProcessId string
	StartTime string
	Result    string
	Msg       string
}

// TaskListFilter 批量查询过滤条件
type EquipAttrtFilter struct {
	AttrKey string
}

type SystemRepo interface {
	InsertEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error)
	ListEquipAttr(ctx context.Context, filter *EquipAttrtFilter) ([]*EquipAttr, error)
	UpdateEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error)
	GetEquipAttr(ctx context.Context, attrKey string) (string, error)
	DeleteEquipAttr(ctx context.Context, attrKey string) (int, error)

	InsertActiveInfo(ctx context.Context, process_id string) (bool, error)
	UpdateActiveInfo(ctx context.Context, process_id string, result string, msg string) (bool, error)
	GetActiveInfo(ctx context.Context, process_id string) (*ActiveInfo, error)

	GetVersion(ctx context.Context) (string, error)
	UpdateVersion(ctx context.Context, newVersion string) error
}

type SystemUsecase struct {
	repo SystemRepo

	log *log.Helper
}

func NewSystemUsecase(repo SystemRepo, logger log.Logger) *SystemUsecase {
	return &SystemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SystemUsecase) InsertEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error) {
	return uc.repo.InsertEquipAttr(ctx, attrKey, attrValue)
}

// List 查询equip_attr
func (uc *SystemUsecase) ListEquipAttr(ctx context.Context, filter *EquipAttrtFilter) ([]*EquipAttr, error) {
	return uc.repo.ListEquipAttr(ctx, filter)
}

// Update 更新equip_attr
func (uc *SystemUsecase) UpdateEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error) {
	return uc.repo.UpdateEquipAttr(ctx, attrKey, attrValue)
}

func (uc *SystemUsecase) DeleteEquipAttr(ctx context.Context, attrKey string) (int, error) {
	return uc.repo.DeleteEquipAttr(ctx, attrKey)
}

func (uc *SystemUsecase) InsertActiveInfo(ctx context.Context, process_id string) (bool, error) {
	return uc.repo.InsertActiveInfo(ctx, process_id)
}

func (uc *SystemUsecase) GetEquipAttr(ctx context.Context, attrKey string) (string, error) {
	return uc.repo.GetEquipAttr(ctx, attrKey)
}

func (uc *SystemUsecase) UpdateActiveInfo(ctx context.Context, process_id string, result string, msg string) (bool, error) {
	return uc.repo.UpdateActiveInfo(ctx, process_id, result, msg)
}

func (uc *SystemUsecase) GetActiveInfo(ctx context.Context, process_id string) (*ActiveInfo, error) {
	return uc.repo.GetActiveInfo(ctx, process_id)
}

func (uc *SystemUsecase) GetVersion(ctx context.Context) (string, error) {
	return uc.repo.GetVersion(ctx)
}

func (uc *SystemUsecase) UpdateVersion(ctx context.Context, newVersion string) error {
	return uc.repo.UpdateVersion(ctx, newVersion)
}
