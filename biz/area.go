package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Area struct {
	// id
	ID uint64 `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 层级
	Level uint64 `json:"level,omitempty"`
	// 父id
	Pid int64 `json:"pid,omitempty"`
}

type AreaRepo interface {
	// List 查询
	List(ctx context.Context, level uint32, pid int32) ([]*Area, error)
	QueryNamesByIds(ctx context.Context, id []uint64) ([]string, error)
	QueryByIds(ctx context.Context, id []uint64) ([]*Area, error)
	QueryNameIdPidMap(ctx context.Context) (map[string]string, map[string]string, error)
}

type AreaUsecase struct {
	areaRepo AreaRepo

	log *log.Helper
}

func NewAreaUsecase(repo AreaRepo, logger log.Logger) *AreaUsecase {
	return &AreaUsecase{areaRepo: repo, log: log.NewHelper(logger)}
}

// List 根据条件查询
func (uc *AreaUsecase) List(ctx context.Context, level uint32, pid int32) ([]*Area, error) {
	return uc.areaRepo.List(ctx, level, pid)
}

func (uc *AreaUsecase) QueryNamesByIds(ctx context.Context, id []uint64) ([]string, error) {
	return uc.areaRepo.QueryNamesByIds(ctx, id)
}

func (uc *AreaUsecase) QueryByIds(ctx context.Context, id []uint64) ([]*Area, error) {
	return uc.areaRepo.QueryByIds(ctx, id)
}

func (uc *AreaUsecase) QueryNameIdPidMap(ctx context.Context) (map[string]string, map[string]string, error) {
	return uc.areaRepo.QueryNameIdPidMap(ctx)
}
