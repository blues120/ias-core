package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/pagination"
)

// EdgeDeviceType 设备类型  DeviceType被占用
type EdgeDeviceType string

const (
	EdgeDeviceTypeBox  EdgeDeviceType = "box"  // 边缘盒子
	EdgeDeviceTypeNode EdgeDeviceType = "node" // 边缘节点
)

func (EdgeDeviceType) Values() []string {
	return []string{string(EdgeDeviceTypeBox), string(EdgeDeviceTypeNode)}
}

// DeviceState 设备状态
type DeviceState string

const (
	DeviceStateOnline  DeviceState = "online"  // 在线
	DeviceStateOffline DeviceState = "offline" // 离线
)

func (DeviceState) Values() []string {
	return []string{string(DeviceStateOnline), string(DeviceStateOffline)}
}

// 设备信息
type Device struct {
	Id             uint64         // id
	Name           string         // 边缘设备名称
	DisplayName    string         // 边缘设备展示名称
	Type           EdgeDeviceType // 边缘设备类型
	ExtId          string         // 边缘设备ID
	SerialNo       string         // 序列号
	State          DeviceState    // 状态
	Mac            string         // MAC地址
	ZoneName       string         // 区域名称
	ZoneId         string         // 区域ID
	WorkspaceId    string         // 工作空间ID
	CreatedAt      time.Time      // 创建时间
	UpdatedAt      time.Time      // 更新时间
	DeletedAt      time.Time      // 删除时间
	ActivatedAt    time.Time      // 重新激活时间
	Cameras        []*Camera      // 设备关联的摄像头
	EquipId        string         // 数生使用的设备ID
	EquipPassword  string         // 数生使用的设备密码
	DeviceInfo     string         // 设备信息
	DeviceAttrInfo *DeviceAttr    // 设备信息的映射
	AuthDeadline   int64          // 授权有效期
	Model          string         // 设备型号
	TenantId       string         // 租户id
	AccessOrgList  string         // 授权的组织id列表
}

type DeviceAttr struct {
	Version string `json:"version"` // 系统版本号
	Model   string `json:"model"`   // 设备型号
}

// 设备接入码/激活码
type DeviceToken struct {
	DeviceExtId string // 所绑定的设备ID
	Token       string // 接入码
}

// DeviceListFilter 批量查询过滤条件
type DeviceListFilter struct {
	/*
		模糊查询条件
	*/
	Name     string
	SerialNo string
	/*
		精确查询条件
	*/
	State DeviceState    // 状态
	Type  EdgeDeviceType // 类型
	/*
		分页
	*/
	Pagination *pagination.Pagination // 分页

	// 特殊查询项
	IncludeDeleted bool // 查询结果是否包含已删除的设备

	// 批量查询
	DeviceIds []uint64 // 设备ID列表
	Models    []string // 设备型号列表
}

type DeviceOption struct {
	PreloadCamera  bool // 加载关联的摄像头数据
	IncludeDeleted bool // 查询结果是否包含已删除的设备
}

// 更新设备选项
type UpdateDeviceOption struct {
	ReActivate bool // 是否重新激活: 恢复设备为未删除并重置激活时间
	SkipCamera bool // 是否跳过更新摄像头
}

// 重复性检查参数
type DeviceUniqueFilter struct {
	SerialNo    string // 序列号
	DisplayName string // 展示名称
	ExcludeId   uint64 // 排除的ID
}

type DeviceRepo interface {
	// List 查询节点列表
	List(ctx context.Context, filter *DeviceListFilter, option *DeviceOption) ([]*Device, int, error)
	// Find 查询设备详情
	Find(ctx context.Context, deviceId uint64, option *DeviceOption) (*Device, error)
	// FindByExtID 根据ext id查询设备详情
	FindByExtID(ctx context.Context, deviceExtId string) (bool, *Device, error)
	// FindByMac 根据mac地址查询设备详情
	FindByMac(ctx context.Context, mac string) (bool, *Device, error)
	// FindByModel 根据型号查询设备详情
	FindByModel(ctx context.Context, model string) (*Device, error)
	// CheckExists 检查设备是否存在
	CheckExists(ctx context.Context, filter *DeviceUniqueFilter) (bool, error)
	// Save 添加设备
	Save(ctx context.Context, device *Device) (uint64, error)
	// Update 更新设备
	Update(ctx context.Context, deviceId uint64, device *Device, option *UpdateDeviceOption) error
	// Delete 删除设备
	Delete(ctx context.Context, deviceId uint64) error
	// UpdateState 更新设备状态
	UpdateState(ctx context.Context, deviceId uint64, state DeviceState) error

	// GetAvailableToken 获取可用设备接入码/激活码
	GetAvailableToken(ctx context.Context) (*DeviceToken, error)
	// FindToken 查询设备接入码/激活码
	FindToken(ctx context.Context, token string) (*DeviceToken, error)
	// SaveToken 保存设备接入码/激活码
	SaveToken(ctx context.Context, token *DeviceToken) error
	// BindTokenToDevice 绑定设备接入码/激活码
	BindTokenToDevice(ctx context.Context, token *DeviceToken) error

	// CountByStatus 查询边缘设备接入统计数据
	CountByStatus(ctx context.Context) ([]*DeviceStatusCount, error)
}

type DeviceUsecase struct {
	repo DeviceRepo
	log  *log.Helper
}

func NewDeviceUsecase(repo DeviceRepo, logger log.Logger) *DeviceUsecase {
	return &DeviceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// List 查询节点列表
func (uc *DeviceUsecase) List(ctx context.Context, filter *DeviceListFilter, option *DeviceOption) ([]*Device, int, error) {
	return uc.repo.List(ctx, filter, option)
}

// Find 查询设备详情
func (uc *DeviceUsecase) Find(ctx context.Context, deviceId uint64, option *DeviceOption) (*Device, error) {
	return uc.repo.Find(ctx, deviceId, option)
}

// FindByExtID 根据ext id查询设备详情
func (uc *DeviceUsecase) FindByExtID(ctx context.Context, deviceExtId string) (exists bool, device *Device, err error) {
	return uc.repo.FindByExtID(ctx, deviceExtId)
}

// FindByMac 根据mac地址查询设备详情
func (uc *DeviceUsecase) FindByMac(ctx context.Context, mac string) (exists bool, device *Device, err error) {
	return uc.repo.FindByMac(ctx, mac)
}

// FindByMode 根据型号查询设备详情
func (uc *DeviceUsecase) FindByModel(ctx context.Context, model string) (device *Device, err error) {
	return uc.repo.FindByModel(ctx, model)
}

// Create 添加设备
func (uc *DeviceUsecase) Create(ctx context.Context, device *Device) (uint64, error) {
	return uc.repo.Save(ctx, device)
}

// CheckExists 检查设备是否存在
func (uc *DeviceUsecase) CheckExists(ctx context.Context, filter *DeviceUniqueFilter) (bool, error) {
	return uc.repo.CheckExists(ctx, filter)
}

// Update 更新设备
func (uc *DeviceUsecase) Update(ctx context.Context, deviceId uint64, device *Device, option *UpdateDeviceOption) error {
	return uc.repo.Update(ctx, deviceId, device, option)
}

// Delete 删除设备
func (uc *DeviceUsecase) Delete(ctx context.Context, deviceId uint64) error {
	return uc.repo.Delete(ctx, deviceId)
}

// UpdateState 更新设备状态
func (uc *DeviceUsecase) UpdateState(ctx context.Context, deviceId uint64, state DeviceState) error {
	return uc.repo.UpdateState(ctx, deviceId, state)
}

// GetAvailableToken 获取可用设备接入码/激活码
func (uc *DeviceUsecase) GetAvailableToken(ctx context.Context) (*DeviceToken, error) {
	return uc.repo.GetAvailableToken(ctx)
}

// FindToken 查询设备接入码/激活码
func (uc *DeviceUsecase) FindToken(ctx context.Context, token string) (*DeviceToken, error) {
	return uc.repo.FindToken(ctx, token)
}

// SaveToken 保存设备接入码/激活码
func (uc *DeviceUsecase) SaveToken(ctx context.Context, token *DeviceToken) error {
	return uc.repo.SaveToken(ctx, token)
}

// BindTokenToDevice 绑定设备接入码/激活码
func (uc *DeviceUsecase) BindTokenToDevice(ctx context.Context, token *DeviceToken) error {
	return uc.repo.BindTokenToDevice(ctx, token)
}

type DeviceStatusCount struct {
	Status DeviceState `json:"state"` // 是否在线
	Count  int32       `json:"count"` // 总数
}

// CountByStatus 统计每种状态包含的设备数量
func (uc *DeviceUsecase) CountByStatus(ctx context.Context) ([]*DeviceStatusCount, error) {
	return uc.repo.CountByStatus(ctx)
}
