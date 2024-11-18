package biz

import (
	"context"
	"errors"
	"time"

	"github.com/blues120/ias-core/pkg/pagination"
)

type DeviceAlgo struct {
	ID               uint64    `json:"id"`                 // id
	DeviceId         uint64    `json:"device_id"`          // 设备id
	AlgoGroupID      uint      `json:"algo_group_id"`      // 算法组ID
	AlgoGroupName    string    `json:"algo_group_name"`    // 算法组name(如十二合一算法)，前端显示用
	AlgoGroupVersion string    `json:"algo_group_version"` // 算法组版本号
	Name             string    `json:"name"`               // 算法名称
	InstallTime      time.Time `json:"install_time"`       // 安装时间
	Version          string    `json:"version"`            // 算法版本
}

type DeviceAlgoRepo interface {
	List(ctx context.Context, filter *DeviceAlgoFilter) ([]*DeviceAlgo, int, error)
	Save(ctx context.Context, algo *DeviceAlgo) (uint64, error)
	Update(ctx context.Context, algo *DeviceAlgo) error
	Delete(ctx context.Context, id uint64) error
	Find(ctx context.Context, id uint64) (*DeviceAlgo, error)
	FindByCondition(ctx context.Context, algo *DeviceAlgo) (bool, error)
	Count(ctx context.Context) (int, error)
	BatchUpdateDeviceAlgo(ctx context.Context, deviceId uint64, algoList []*DeviceAlgo) error
}

// DeviceAlgoFilter 批量查询过滤条件
type DeviceAlgoFilter struct {
	DeviceId    uint64 		// 设备id
	AlgoGroupId uint   		// 算法组id
	DeviceIds   []uint64 	// 设备ID列表
	/*
		分页
	*/
	Pagination *pagination.Pagination
}

type DeviceAlgoUsecase struct {
	repo DeviceAlgoRepo
}

func NewDeviceAlgoUsecase(repo DeviceAlgoRepo) *DeviceAlgoUsecase {
	return &DeviceAlgoUsecase{repo: repo}
}

// List 根据条件查询
func (uc *DeviceAlgoUsecase) List(ctx context.Context, filter *DeviceAlgoFilter) ([]*DeviceAlgo, int, error) {
	return uc.repo.List(ctx, filter)
}

func (uc *DeviceAlgoUsecase) Create(ctx context.Context, deviceAlgoReq *DeviceAlgo) (uint64, error) {
	// 算法名称重名校验
	isExist, err := uc.repo.FindByCondition(ctx, deviceAlgoReq)
	if err != nil {
		return 0, err
	}
	if isExist {
		return 0, errors.New("deviceAlgo  already exists")
	}
	return uc.repo.Save(ctx, deviceAlgoReq)
}

func (uc *DeviceAlgoUsecase) Update(ctx context.Context, id uint64, AlgoReq *DeviceAlgo) error {
	update, err := uc.repo.Find(ctx, id)
	if err != nil {
		return err
	}
	AlgoReq.ID = update.ID
	return uc.repo.Update(ctx, AlgoReq)
}

// Delete 根据id删除
func (uc *DeviceAlgoUsecase) Delete(ctx context.Context, id uint64) error {
	info, err := uc.repo.Find(ctx, id)
	if err != nil || info == nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}

func (uc *DeviceAlgoUsecase) Find(ctx context.Context, id uint64) (*DeviceAlgo, error) {
	return uc.repo.Find(ctx, id)
}

// Count 查询算法总数
func (uc *DeviceAlgoUsecase) Count(ctx context.Context) (int, error) {
	return uc.repo.Count(ctx)
}

func (uc *DeviceAlgoUsecase) BatchUpdateDeviceAlgo(ctx context.Context, deviceId uint64, algoList []*DeviceAlgo) error {
	return uc.repo.BatchUpdateDeviceAlgo(ctx, deviceId, algoList)
}
