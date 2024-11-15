package biz

import (
	"context"
	"errors"
	"strings"
	"time"

	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/pagination"
)

type UpdateAlgoKey struct{}

type AlgoOption struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type AlgoCondition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type Linkage struct {
	Target     string          `json:"target"`
	Action     string          `json:"action"`
	Options    []AlgoOption    `json:"options"`
	Value      interface{}     `json:"value"`
	Conditions []AlgoCondition `json:"conditions"`
}
type InputParamsInfo struct {
	Name         string       `json:"name"`         // 参数名
	Label        string       `json:"label"`        // 参数中文名
	Details      string       `json:"details"`      // 参数描述
	Type         string       `json:"type"`         // 参数值类型：string/int/float/enumValue/enumValueArray/bool/line//polygon /worktime其中之一
	Sence        string       `json:"sence"`        // 参数场景
	DefaultValue interface{}  `json:"defaultValue"` // 默认值
	Required     bool         `json:"required"`     // 是否为算法必须
	Min          int          `json:"min"`          // 最小值
	Max          int          `json:"max"`          // 最大值
	Step         int          `json:"step"`         // ?
	Group        string       `json:"group"`
	Options      []AlgoOption `json:"options"`
	Linkages     []Linkage    `json:"linkage"`
}

type AlgoConfigJsons struct {
	AlgoJsons []InputParamsInfo
}

type Algorithm struct {
	ID               uint64    `json:"id"`                 // id
	Name             string    `json:"name"`               // 算法名称
	Type             string    `json:"type"`               // 算法类型（图片帧image/视频流video）
	Desc             string    `json:"description"`        // 算法描述
	Version          string    `json:"version"`            // 算法版本
	AppName          string    `json:"app_name"`           // 应用名称
	AlarmType        string    `json:"alarm_type"`         // 告警类型
	AlarmName        string    `json:"alarm_name"`         // 告警名称
	Notify           uint      `json:"notify"`             // 是否告警
	Extend           string    `json:"extend"`             // 非通用属性
	Detail           string    `json:"detail"`             // 详情json（图片帧）
	FileName         string    `json:"file_name"`          // 算法文件名称（视频流）
	ConfigName       string    `json:"config_name"`        // 配置文件名称（视频流）
	CreatedAt        time.Time `json:"created_at"`         // 创建时间
	UpdatedAt        time.Time `json:"updated_at"`         // 更新时间
	DrawType         uint      `json:"draw_type"`          // 绘制区域类型 多边形区域/流量方向及界线 1/2
	BaseType         uint      `json:"base_type"`          // 底库类型 无/人员/车辆 0/1/2
	Image            string    `json:"image"`              // 算法镜像地址
	LabelMap         string    `json:"label_map"`          // 中英label映射
	Target           string    `json:"target"`             // 检测目标
	AlgoNameEn       string    `json:"algo_name_en"`       // 算法英文名，下发给agent配置需要
	AlgoGroupID      uint      `json:"algo_group_id"`      // 算法组ID
	AlgoGroupName    string    `json:"algo_group_name"`    // 算法组name(如十二合一算法)，前端显示用
	AlgoGroupVersion string    `json:"algo_group_version"` // 算法组版本号
	Config           string    `json:"config"`             // 算法特有配置，如算能算法的单独配置
	Provider         string    `json:"provider"`           // 算法供应商 ctyun_ias/sophgo_park/sophgo_city
	AlgoID           string    `json:"algo_id"`            // 算法 id，非自增 id，用于填充算能园区算法 alg_flag 字段（使用时需转换为 Uint）
	Available        uint      `json:"available"`          // 是否可用 默认为1,不可用为0
	Platform         string    `json:"platform"`           // 平台类型,服务器/边缘设备
	DeviceModel      string    `json:"device_model"`       // 设备型号
	IsGroupType      uint      `json:"is_group_type"`      // 是否是多合一算法组类型
	Prefix           string    `json:"prefix"`             // 算法包服务启动api前缀 ip:port
}

// AlgoFrameExtend 图片帧算法原json结构
type AlgoFrameExtend struct {
	AlertRules [][]struct {
		Compare string      `json:"compare"`
		Key     string      `json:"key"`
		Value   interface{} `json:"value"`
	} `json:"alert_rules"`
	Desc        string `json:"desc" binding:"lte=200"`
	Name        string `json:"name" binding:"required,match1"`
	RequestBody struct {
		Action    string `json:"Action"`
		AppKey    string `json:"AppKey"`
		ImageData string `json:"ImageData"`
		Token     string `json:"Token"`
		Cls       string `json:"Cls,omitempty"`
	} `json:"request_body"`
	ResultType string `json:"result_type" binding:"required"`
	URL        string `json:"url" binding:"required"`
	Version    string `json:"version" binding:"required,match2"`
	Notify     uint   `json:"notify" binding:"oneof=0 1"`
	AlarmType  string `json:"alarm_type" binding:"required"`
	AlarmName  string `json:"alarm_name" binding:"required,match1"`
}

type AlgoStreamExtend struct {
	FileName   string `json:"fileName"`
	ConfigName string `json:"configName"`
}

type CanAlgoOperation struct {
	CanOperation bool   `json:"canOperation"` // 是否可以进行操作
	RelTaskMsg   string `json:"relTaskMsg"`   // 绑定算法的任务信息
}

// AlgoFilter 批量查询过滤条件
type AlgoFilter struct {
	NameEq             string // 算法名称
	AvailableEq        *uint
	AlgoNameEn         string // 英文算法名
	ProviderEq         string // 供应商
	AlgoGroupIDEq      uint   // 算法组id
	AlgoGroupNameEq    string // 算法组名称
	AlgoGroupVersionEq string // 算法组版本号
	PlatformEq         string // 平台
	DeviceModelEq      string // 设备型号
	/*
		模糊查询条件
	*/
	NameContains          string // 名称包含
	AlgoGroupNameContains string // 算法包名称包含

	/*
		范围查询条件
	*/
	IDIn []uint64 // 包含算法ID
	/*
		分页
	*/
	Pagination *pagination.Pagination
}

type AlgoRepo interface {
	List(ctx context.Context, filter *AlgoFilter) ([]*Algorithm, int, error)
	Save(ctx context.Context, algo *Algorithm) (uint64, error)
	Update(ctx context.Context, algo *Algorithm) error
	Delete(ctx context.Context, id uint64) error
	FindByCondition(ctx context.Context, algo *Algorithm) (bool, error)
	Find(ctx context.Context, id uint64) (*Algorithm, error)
	FindOneByAlgoNameEn(ctx context.Context, algoNameEn string) (*Algorithm, error)
	FindAlgorithms(ctx context.Context, filter *AlgoFilter) ([]*Algorithm, error)
	FindTasksByAlgorithmID(ctx context.Context, id uint64) ([]*Task, error)
	ResetAvailableAlgo(ctx context.Context, IDList []uint64) error
	Count(ctx context.Context) (int, error)
	UpdateAlgoGroupVersion(ctx context.Context, name, version string) error
}

type AlgoUsecase struct {
	repo AlgoRepo
}

func NewAlgoUsecase(repo AlgoRepo) *AlgoUsecase {
	return &AlgoUsecase{repo: repo}
}

// List 根据条件查询
func (uc *AlgoUsecase) List(ctx context.Context, filter *AlgoFilter) ([]*Algorithm, int, error) {
	return uc.repo.List(ctx, filter)
}

// Create 校验是否重名并入库
// ale 为历史遗留字段，各服务版本统一后考虑删除
func (uc *AlgoUsecase) Create(ctx context.Context, AlgoReq *Algorithm, ale *AlgoFrameExtend) (uint64, error) {
	// 算法名称重名校验
	isExist, err := uc.repo.FindByCondition(ctx, AlgoReq)
	if err != nil {
		return 0, err
	}
	if isExist {
		return 0, errors.New("algorithm name already exists")
	}

	// 告警名称重名校验
	if AlgoReq.AlarmName != "" {
		isAlarmNameExist, err := uc.repo.FindByCondition(ctx, AlgoReq)
		if err != nil {
			return 0, err
		}
		if isAlarmNameExist {
			return 0, errors.New("algorithm alarm name already exists")
		}
	}

	return uc.repo.Save(ctx, AlgoReq)
}

// Update 校验是否重名并update库
// ale 为历史遗留字段，各服务版本统一后考虑删除
func (uc *AlgoUsecase) Update(ctx context.Context, id uint64, AlgoReq *Algorithm, ale *AlgoFrameExtend) error {
	// 算法ID是否存在
	update, err := uc.repo.Find(ctx, id)
	if err != nil {
		return err
	}

	// 算法名称重名校验
	if update.Name != AlgoReq.Name {
		isExist, err := uc.repo.FindByCondition(ctx, AlgoReq)
		if err != nil {
			return err
		}
		if isExist {
			return errors.New("algorithm name already exists")
		}
	}

	// 告警名称重名校验
	if update.AlarmName != AlgoReq.AlarmName {
		// 告警名称重名校验
		isExist, err := uc.repo.FindByCondition(ctx, AlgoReq)
		if err != nil {
			return err
		}
		if isExist {
			return errors.New("algorithm alarm name already exists")
		}
	}
	// 判断是否有正在运行的任务
	updateSenario, ok := ctx.Value(UpdateAlgoKey{}).(string)
	if !ok || updateSenario != "install" { // 如果是安装算法包场景，跳过对正在运行的任务的判断
		tasks, err := uc.repo.FindTasksByAlgorithmID(ctx, id)
		if err != nil {
			return err
		}
		for _, task := range tasks {
			if task.Status != TaskStatusStopped {
				return errors.New("algorithm has running task")
			}
		}
	}

	AlgoReq.ID = update.ID
	return uc.repo.Update(ctx, AlgoReq)
}

// Delete 根据id删除
func (uc *AlgoUsecase) Delete(ctx context.Context, id uint64) error {
	tasks, err := uc.repo.FindTasksByAlgorithmID(ctx, id)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.Status != TaskStatusStopped {
			return errors.New("algorithm has running task")
		}
	}
	return uc.repo.Delete(ctx, id)
}

// CanDo 进行编辑或删除操作前判断是否可行
func (uc *AlgoUsecase) CanDo(ctx context.Context, id uint64, operation string) (result CanAlgoOperation, err error) {
	result.CanOperation = true
	// 根据算法ID查询任务
	tasks, err := uc.repo.FindTasksByAlgorithmID(ctx, id)
	if err != nil {
		return result, err
	}

	// 获取所有任务名称/运行中的任务名称
	var runningTaskNames, TaskNames []string
	for _, task := range tasks {
		if task.Status != TaskStatusStopped {
			runningTaskNames = append(runningTaskNames, task.Name)
		}
		TaskNames = append(TaskNames, task.Name)
	}
	if operation == "update" {
		if len(runningTaskNames) != 0 {
			result.CanOperation = false
			result.RelTaskMsg = "该算法已在如下任务中运行:" + strings.Join(runningTaskNames, ",") + "。请停止相应任务后重试"
		}
	} else {
		if len(TaskNames) != 0 {
			result.CanOperation = false
			result.RelTaskMsg = "该算法已在如下任务中绑定:" + strings.Join(TaskNames, ",") + "。请取消绑定后重试"
		}
	}
	return
}

func (uc *AlgoUsecase) ResetAvailableAlgo(ctx context.Context, IDList []uint64) error {
	return uc.repo.ResetAvailableAlgo(ctx, IDList)
}

func (uc *AlgoUsecase) Find(ctx context.Context, id uint64) (*Algorithm, error) {
	return uc.repo.Find(ctx, id)
}

// FindOneByAlgoNameEn 根据算法英文名称查询
func (uc *AlgoUsecase) FindOneByAlgoNameEn(ctx context.Context, algoNameEn string) (*Algorithm, error) {
	return uc.repo.FindOneByAlgoNameEn(ctx, algoNameEn)
}

// Count 查询算法总数
func (uc *AlgoUsecase) Count(ctx context.Context) (int, error) {
	return uc.repo.Count(ctx)
}

// UpdateAlgoGroupVersion 更新算法组版本号
func (uc *AlgoUsecase) UpdateAlgoGroupVersion(ctx context.Context, name, version string) error {
	return uc.repo.UpdateAlgoGroupVersion(ctx, name, version)
}

// UpdateAlgoGroupVersion 更新算法组版本号
func (uc *AlgoUsecase) FindAlgorithms(ctx context.Context, filter *AlgoFilter) ([]*Algorithm, error) {
	return uc.repo.FindAlgorithms(ctx, filter)
}
