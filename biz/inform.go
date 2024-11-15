package biz

import (
	"context"
	"time"

	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/pagination"
)

type Inform struct {
	ID             uint64    // id
	AlarmName      string    // 告警名称
	AlarmType      string    // 告警类型
	SignName       string    // 签名
	NotifyTemplate string    // 通知模板内容
	TemplateCode   string    // 通知模板
	PhoneNumbers   string    // 通知号码
	NotifySwitch   string    // 通知开关
	TaskID         uint64    // 任务ID
	TaskName       string    // 任务名称
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间
	TenantID       string    `json:"tenant_id,omitempty"`       // namespace 区间
	AccessOrgList  string    `json:"access_org_list,omitempty"` // 授权的组织 id 列表，#分隔
}

// InformFilter 批量查询过滤条件
type InformFilter struct {
	Pagination *pagination.Pagination // 分页
}

type InformRepo interface {
	List(ctx context.Context, filter *InformFilter) ([]*Inform, int, error)
	Save(ctx context.Context, inform *Inform) (uint64, error)
	Update(ctx context.Context, inform *Inform) error
	Delete(ctx context.Context, id uint64) error
	Find(ctx context.Context, id uint64) (*Inform, error)
	FindByTaskID(ctx context.Context, taskID uint64) (*Inform, error)
	DeleteByTaskID(ctx context.Context, taskID uint64) (int, error)
}

type InformUsecase struct {
	repo InformRepo
}

func NewInformUsecase(repo InformRepo) *InformUsecase {
	return &InformUsecase{repo: repo}
}

// List 根据条件查询
func (uc *InformUsecase) List(ctx context.Context, filter *InformFilter) ([]*Inform, int, error) {
	return uc.repo.List(ctx, filter)
}

// Create 入库
func (uc *InformUsecase) Create(ctx context.Context, InformReq *Inform) (uint64, error) {
	return uc.repo.Save(ctx, InformReq)
}

// Update 校验是否重名并update库
func (uc *InformUsecase) Update(ctx context.Context, id uint64, InformReq *Inform) error {
	return uc.repo.Update(ctx, InformReq)
}

// Delete 根据id删除
func (uc *InformUsecase) Delete(ctx context.Context, id uint64) error {
	return uc.repo.Delete(ctx, id)
}

// Find 根据id查找
func (uc *InformUsecase) Find(ctx context.Context, id uint64) (*Inform, error) {
	return uc.repo.Find(ctx, id)
}

// FindByTaskID 根据任务ID查找
func (uc *InformUsecase) FindByTaskID(ctx context.Context, taskID uint64) (*Inform, error) {
	return uc.repo.FindByTaskID(ctx, taskID)
}

// DeleteByTaskID 根据任务ID删除
func (uc *InformUsecase) DeleteByTaskID(ctx context.Context, taskID uint64) (int, error) {
	return uc.repo.DeleteByTaskID(ctx, taskID)
}
