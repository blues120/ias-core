package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type WarnType struct {
	ID        uint32     //告警ID
	AppName   string     // 应用名称
	AlarmType string     // 告警类型
	AlarmName string     // 告警名称
	SubAlarms []SubAlarm //子警告名称

}
type WarnTypeFilter struct {
	NameEq string // 名称
}

type SubAlarm struct {
	ID          uint32
	AlarmName   string
	AlarmTypeID uint32 //类型id
}

type WarnTypeRepo interface {
	List(ctx context.Context, filter *WarnTypeFilter) ([]*SubAlarm, error)
}

type WarnTypeUsecase struct {
	repo WarnTypeRepo
	log  *log.Helper
}

func NewWarnTypeUsecase(repo WarnTypeRepo, logger log.Logger) *WarnTypeUsecase {
	return &WarnTypeUsecase{repo: repo, log: log.NewHelper(logger)}
}

// List 查询
func (uc *WarnTypeUsecase) List(ctx context.Context, filter *WarnTypeFilter) ([]*SubAlarm, error) {
	return uc.repo.List(ctx, filter)
}
