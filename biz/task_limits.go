package biz

import (
	"context"

	"github.com/blues120/ias-core/pkg/pagination"
)

type TaskLimits struct {
	ID            uint64 // id
	Model         string `json:"model"`         // 盒子设备型号
	MaxCameraNum  uint64 `json:"maxCameraNum"`  // 最大摄像机数
	AlgoNum       uint64 `json:"algoNum"`       // 算法数
	MaxSubTaskNum uint64 `json:"maxSubTaskNum"` // 最大子任务数
}

// InformFilter 批量查询过滤条件
type TasklimitsFilter struct {
	Pagination *pagination.Pagination // 分页
}

type TasklimitsRepo interface {
	List(ctx context.Context, filter *TaskListFilter) ([]*TaskLimits, int, error)
	Save(ctx context.Context, tasklimits *TaskLimits) (uint64, error)
	FindByModel(ctx context.Context, model string) (*TaskLimits, error)
	Update(ctx context.Context, model string, tasklimits *TaskLimits) error
	DeleteByModel(ctx context.Context, model string) (int, error)
}

type TasklimitsUsecase struct {
	repo TasklimitsRepo
}

func NewTasklimitsUsecase(repo TasklimitsRepo) *TasklimitsUsecase {
	return &TasklimitsUsecase{repo: repo}
}

// List 根据条件查询
func (uc *TasklimitsUsecase) List(ctx context.Context, filter *TaskListFilter) ([]*TaskLimits, int, error) {
	return uc.repo.List(ctx, filter)
}

// Create 入库
func (uc *TasklimitsUsecase) Create(ctx context.Context, tasklimitsReq *TaskLimits) (uint64, error) {
	return uc.repo.Save(ctx, tasklimitsReq)
}

// Update 校验是否重名并update库
func (uc *TasklimitsUsecase) Update(ctx context.Context, model string, tasklimitsReq *TaskLimits) error {
	return uc.repo.Update(ctx, model, tasklimitsReq)
}

// FindByTaskID 根据任务ID查找
func (uc *TasklimitsUsecase) FindByModel(ctx context.Context, model string) (*TaskLimits, error) {
	return uc.repo.FindByModel(ctx, model)
}

// DeleteByTaskID 根据任务ID删除
func (uc *TasklimitsUsecase) DeleteByModel(ctx context.Context, model string) (int, error) {
	return uc.repo.DeleteByModel(ctx, model)
}
