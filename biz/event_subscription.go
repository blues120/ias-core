package biz

import (
	"context"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/pagination"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// EventSubStatus 订阅状态
type EventSubStatus string

// Values provides list valid values for Enum.
func (EventSubStatus) Values() []string {
	return []string{
		string(EventSubStatusUnknown),
		string(EventSubStatusEnable),
		string(EventSubStatusDisable),
	}
}

const (
	// 全部状态
	EventSubStatusUnknown EventSubStatus = "unknown"
	// EventSubStatusEnable  开启订阅
	EventSubStatusEnable EventSubStatus = "enable"
	// EventSubStatusDisable 停止订阅
	EventSubStatusDisable EventSubStatus = "disable"
)

// EventSubscription 订阅告警
type EventSubscription struct {
	ID         uint64         // id
	BoxId      string         // 盒子id
	ChannelId  string         // 通道id
	Callback   string         // 回调地址
	TemplateId string         // 模版id
	Status     EventSubStatus // 订阅状态
	CreatedAt  time.Time      // 创建时间
	UpdatedAt  time.Time      // 更新时间
}

type EventSubsRepo interface {
	List(ctx context.Context, filter *EventSubsFilter) ([]*EventSubscription, int, error)
	Save(ctx context.Context, eventSub *EventSubscription) (uint64, error)
	Update(ctx context.Context, id uint64, eventSub *EventSubscription) error
	Delete(ctx context.Context, id uint64) error
	Find(ctx context.Context, id uint64) (*EventSubscription, error)
	FindByFilter(ctx context.Context, filter *EventSubsFindFilter) (*EventSubscription, error)
	CheckExists(ctx context.Context, req *CheckEventSubsExistsRequest) (bool, error)
	BatchUpdatestatus(ctx context.Context, filter *EventSubsFilter) error
}

// EventSubsFilter 批量查询过滤条件
type EventSubsFilter struct {
	/*
		精确查询条件
	*/
	Status     EventSubStatus // 状态
	BoxId      string         // 盒子id
	ChannelId  string         // 通道id
	TemplateId string         // 模版id
	/*
		范围查询条件
	*/
	ChannelIds []string // 包含通道id
	/*
		分页
	*/
	Pagination *pagination.Pagination // 分页
}

// EventSubsFindFilter 订阅查询字段参数
type EventSubsFindFilter struct {
	BoxId      string         // 盒子id
	ChannelId  string         // 通道id
	TemplateId string         // 模版id
	Status     EventSubStatus // 状态
}

// CheckEventSubsExistsRequest 订阅字段值重复校验请求参数
type CheckEventSubsExistsRequest struct {
	BoxId     string // 盒子id
	ChannelId string // 通道id
}

type EventSubsUsecase struct {
	eventSubsRepo EventSubsRepo

	log *log.Helper
}

func NewEventSubsUsecase(repo EventSubsRepo, logger log.Logger) *EventSubsUsecase {
	return &EventSubsUsecase{eventSubsRepo: repo, log: log.NewHelper(logger)}
}

// Create 创建
func (uc *EventSubsUsecase) Create(ctx context.Context, ca *EventSubscription) (uint64, error) {
	return uc.eventSubsRepo.Save(ctx, ca)
}

// Update 更新
func (uc *EventSubsUsecase) Update(ctx context.Context, id uint64, ca *EventSubscription) error {
	return uc.eventSubsRepo.Update(ctx, id, ca)
}

// Delete 删除
func (uc *EventSubsUsecase) Delete(ctx context.Context, id uint64) error {
	_, err := uc.eventSubsRepo.Find(ctx, id)
	if err != nil {
		return err
	}
	return uc.eventSubsRepo.Delete(ctx, id)
}

// Find 查询
func (uc *EventSubsUsecase) Find(ctx context.Context, id uint64) (*EventSubscription, error) {
	return uc.eventSubsRepo.Find(ctx, id)
}

// FindByFilter 查询
func (uc *EventSubsUsecase) FindByFilter(ctx context.Context, filter *EventSubsFindFilter) (*EventSubscription, error) {
	return uc.eventSubsRepo.FindByFilter(ctx, filter)
}

// List 批量查询
func (uc *EventSubsUsecase) List(ctx context.Context, filter *EventSubsFilter) ([]*EventSubscription, int, error) {
	return uc.eventSubsRepo.List(ctx, filter)
}

// CheckExists 检查的字段值是否存在
func (uc *EventSubsUsecase) CheckExists(ctx context.Context, req *CheckEventSubsExistsRequest) (bool, error) {
	return uc.eventSubsRepo.CheckExists(ctx, req)
}

// Disablestatus 将推送状态置为disable
func (r *EventSubsUsecase) BatchUpdatestatus(ctx context.Context, filter *EventSubsFilter) error {
	return r.eventSubsRepo.BatchUpdatestatus(ctx, filter)
}
