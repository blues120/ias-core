package biz

import (
	"context"
	"time"

	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/pagination"

	"github.com/go-kratos/kratos/v2/log"
)

// WarnPushType 推送类型
type WarnPushType string

// Values provides list valid values for Enum.
func (WarnPushType) Values() []string {
	return []string{
		string(WarnPushTypeWebhook),
	}
}

const (
	// WarnPushTypeWebhook webhook
	WarnPushTypeWebhook WarnPushType = "webhook"
)

// WarnPushStatus 推送状态
type WarnPushStatus string

// Values provides list valid values for Enum.
func (WarnPushStatus) Values() []string {
	return []string{
		string(WarnPushStatusUnknown),
		string(WarnPushStatusEnabled),
		string(WarnPushStatusDisabled),
	}
}

const (
	WarnPushStatusUnknown  WarnPushStatus = "unknown"  // 全部状态
	WarnPushStatusEnabled  WarnPushStatus = "enabled"  // 已启用
	WarnPushStatusDisabled WarnPushStatus = "disabled" // 已禁用
)

// WarnPushMode 推送状态
type WarnPushMode string

// Values provides list valid values for Enum.
func (WarnPushMode) Values() []string {
	return []string{
		string(WarnPushModeRealTime),
	}
}

const (
	WarnPushModeRealTime WarnPushMode = "realtime" // 已启用
)

// WarnPushLogStatus 推送状态
type WarnPushLogStatus string

// Values provides list valid values for Enum.
func (WarnPushLogStatus) Values() []string {
	return []string{
		string(WarnPushLogStatusFailed),
		string(WarnPushLogStatusSuccess),
	}
}

const (
	WarnPushLogStatusFailed  WarnPushLogStatus = "failed"  // 失败
	WarnPushLogStatusSuccess WarnPushLogStatus = "success" // 成功
)

// WarnPush 推送
type WarnPush struct {
	Id            uint64         // 推送id
	Name          string         // 推送名称
	Type          WarnPushType   // 推送类型
	Url           string         // 推送地址
	Remark        string         // 备注
	Mode          WarnPushMode   // 推送模式
	Status        WarnPushStatus // 推送状态
	TenantID      string         // 租户id
	AccessOrgList string         // 授权的组织 id 列表，#分隔
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// PushListFilter 批量查询过滤条件
type PushListFilter struct {
	/*
		模糊查询条件
	*/
	NameContain string
	/*
		精确查询条件
	*/
	Status WarnPushStatus // 状态

	/*
		分页
	*/
	Pagination *pagination.Pagination // 分页
}

type WarnPushLog struct {
	Id            uint64 // 推送id
	PushId        uint64
	Push          *WarnPush // 推送设置
	Param         string
	Remark        string            // 备注
	Status        WarnPushLogStatus // 推送状态
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TenantID      string
	AccessOrgList string
}

type PushLogListFilter struct {
	ParamContain string            // 参数
	Status       WarnPushLogStatus // 状态
	BeginTime    *time.Time
	EndTime      *time.Time
	PushIds      []uint64

	/*
		分页
	*/
	Pagination *pagination.Pagination // 分页
}

// CheckWarnPushExistsRequest 任务字段值重复校验请求参数
type CheckWarnPushExistsRequest struct {
	Name      string
	Url       string
	ExcludeId uint64
}

type WarnPushRepo interface {
	// Save 创建任务
	Save(ctx context.Context, ta *WarnPush) (uint64, error)

	// Update 更新任务
	Update(ctx context.Context, id uint64, ta *WarnPush) (bool, error)

	// UpdateStatus 更新状态
	UpdateStatus(ctx context.Context, id uint64, status WarnPushStatus) error

	// Delete 删除任务
	Delete(ctx context.Context, id uint64) error

	// Find 查询任务
	Find(ctx context.Context, id uint64) (*WarnPush, error)

	// CheckExists 检查 WarnPush 的字段值是否重复
	CheckExists(ctx context.Context, req *CheckWarnPushExistsRequest) (bool, error)

	// List 查询任务列表
	List(ctx context.Context, filter *PushListFilter) ([]*WarnPush, int, error)

	// CreateLog 生成推送日志
	CreateLog(ctx context.Context, pushId uint64, log *WarnPushLog) error

	// ListLog 推送日志列表
	ListLog(ctx context.Context, filter *PushLogListFilter) ([]*WarnPushLog, int, error)
}

type WarnPushUsecase struct {
	repo WarnPushRepo

	log *log.Helper
}

func NewPushUsecase(repo WarnPushRepo, logger log.Logger) *WarnPushUsecase {
	return &WarnPushUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create 创建
func (uc *WarnPushUsecase) Create(ctx context.Context, ca *WarnPush) (uint64, error) {
	return uc.repo.Save(ctx, ca)
}

// Update 更新
func (uc *WarnPushUsecase) Update(ctx context.Context, id uint64, ca *WarnPush) (bool, error) {
	return uc.repo.Update(ctx, id, ca)
}

// UpdateStatus 更新状态
func (uc *WarnPushUsecase) UpdateStatus(ctx context.Context, id uint64, status WarnPushStatus) error {
	return uc.repo.UpdateStatus(ctx, id, status)
}

// Delete 删除
func (uc *WarnPushUsecase) Delete(ctx context.Context, id uint64) error {
	return uc.repo.Delete(ctx, id)
}

// Find 查询
func (uc *WarnPushUsecase) Find(ctx context.Context, id uint64) (*WarnPush, error) {
	return uc.repo.Find(ctx, id)
}

// List 批量查询
func (uc *WarnPushUsecase) List(ctx context.Context, filter *PushListFilter) ([]*WarnPush, int, error) {
	return uc.repo.List(ctx, filter)
}

// CheckExists 检查 WarnPush 的字段值是否重复
func (uc *WarnPushUsecase) CheckExists(ctx context.Context, req *CheckWarnPushExistsRequest) (bool, error) {
	return uc.repo.CheckExists(ctx, req)
}

// CreateLog 生成推送日志
func (uc *WarnPushUsecase) CreateLog(ctx context.Context, pushId uint64, log *WarnPushLog) error {
	return uc.repo.CreateLog(ctx, pushId, log)
}

// ListLog 推送日志列表
func (uc *WarnPushUsecase) ListLog(ctx context.Context, filter *PushLogListFilter) ([]*WarnPushLog, int, error) {
	return uc.repo.ListLog(ctx, filter)
}
