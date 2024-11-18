package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/device"
	"github.com/blues120/ias-core/data/ent/devicetoken"
	"github.com/blues120/ias-core/data/ent/mixin"
	"github.com/blues120/ias-core/pkg/convert"
)

type deviceRepo struct {
	data *Data
	tx   biz.Transaction
	log  *log.Helper
}

// NewDeviceRepo 实例化deviceRepo
func NewDeviceRepo(data *Data, trans biz.Transaction, logger log.Logger) biz.DeviceRepo {
	return &deviceRepo{
		data: data,
		tx:   trans,
		log:  log.NewHelper(logger),
	}
}

// CheckExists 检查设备字段值是否重复
func (r *deviceRepo) CheckExists(ctx context.Context, req *biz.DeviceUniqueFilter) (bool, error) {
	query := r.data.db.Device(ctx).Query()
	if req.DisplayName != "" {
		query = query.Where(device.DisplayName(req.DisplayName))
	}
	if req.SerialNo != "" {
		query = query.Where(device.SerialNo(req.SerialNo))
	}
	if req.ExcludeId > 0 {
		query = query.Where(device.IDNEQ(req.ExcludeId))
	}
	return query.Exist(ctx)
}

// List 批量查询
func (r *deviceRepo) List(ctx context.Context, filter *biz.DeviceListFilter, option *biz.DeviceOption) ([]*biz.Device, int, error) {
	query := r.data.db.Device(ctx).Query()
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)

		// 判断是否包含已删除记录
		if filter.IncludeDeleted {
			ctx = mixin.SkipSoftDelete(ctx)
		}
	}

	if option != nil {
		query = r.buildQueryByOption(query, option)
	}

	// 查询总数
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	// 分页查询
	arr, err := r.buildQueryPagination(query, filter).Clone().Order(
		ent.Asc(device.FieldState),
		ent.Desc(device.FieldActivatedAt),
	).All(ctx)

	if err != nil {
		return nil, 0, err
	}

	return DeviceEntArrToBiz(arr), total, nil
}

// Find 查询单个设备详情
func (r *deviceRepo) Find(ctx context.Context, dvcID uint64, option *biz.DeviceOption) (*biz.Device, error) {
	query := r.data.db.Device(ctx).Query().
		Where(device.ID(dvcID))
	if option != nil {
		if option.PreloadCamera {
			query = query.WithCamera()
		}
		// 判断是否包含已删除记录
		if option.IncludeDeleted {
			ctx = mixin.SkipSoftDelete(ctx)
		}
	}
	dvc, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return entToBiz(dvc), nil
}

// FindByExtID 根据device_ext_id查询设备详情
func (r *deviceRepo) FindByExtID(ctx context.Context, extID string) (bool, *biz.Device, error) {
	dvc, err := r.data.db.Device(ctx).Query().
		Where(device.ExtID(extID)).
		First(mixin.SkipSoftDelete(ctx))
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, entToBiz(dvc), nil
}

// FindByMac 根据mac查询设备详情
func (r *deviceRepo) FindByMac(ctx context.Context, mac string) (bool, *biz.Device, error) {
	dvc, err := r.data.db.Device(ctx).Query().
		Where(device.MACEQ(mac)).
		First(mixin.SkipSoftDelete(ctx))
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, entToBiz(dvc), nil
}

// FindByModel 根据型号查询设备详情
func (r *deviceRepo) FindByModel(ctx context.Context, model string) (*biz.Device, error) {
	dvc, err := r.data.db.Device(ctx).Query().
		Where(device.ModelEQ(model)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return entToBiz(dvc), nil
}

// Save 新增设备
func (r *deviceRepo) Save(ctx context.Context, device *biz.Device) (uint64, error) {
	var dvcID uint64
	err := r.tx.InTx(ctx, func(ctx context.Context) error {
		dvc := bizToEnt(device)
		entDvc, err := r.data.db.Device(ctx).Create().
			SetDevice(dvc).
			SetActivatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return err
		}
		dvcID = entDvc.ID

		// 关联camera  新增设备时不需要关联摄像头
		// r.bindCamerasToDevice(ctx, dvcID, device.Cameras)

		return nil
	})

	return dvcID, err
}

// Update 更新设备
func (r *deviceRepo) Update(ctx context.Context, dvcID uint64, device *biz.Device, option *biz.UpdateDeviceOption) error {
	return r.tx.InTx(ctx, func(ctx context.Context) error {
		dvc := bizToEnt(device)
		cmd := r.data.db.Device(ctx).UpdateOneID(dvcID).SetDevice(dvc)
		if option != nil && option.ReActivate {
			cmd = cmd.ClearDeletedAt()           // 恢复为未删除
			cmd = cmd.SetActivatedAt(time.Now()) // 重置激活时间
		}
		_, err := cmd.Save(ctx)
		if err != nil {
			return err
		}

		if option != nil && option.SkipCamera { // 跳过摄像头关联
			return nil
		}

		// 删除关联的摄像头
		if err := r.unbindCamerasFromDevice(ctx, dvcID); err != nil {
			return err
		}

		// 重设关联的摄像头
		return r.bindCamerasToDevice(ctx, dvcID, device.Cameras)
	})
}

// Delete 删除设备
func (r *deviceRepo) Delete(ctx context.Context, dvcID uint64) error {
	return r.tx.InTx(ctx, func(ctx context.Context) error {
		// 删除关联的摄像头
		if err := r.unbindCamerasFromDevice(ctx, dvcID); err != nil {
			return err
		}

		// 删除设备
		return r.data.db.Device(ctx).DeleteOneID(dvcID).Exec(ctx)
	})
}

// GetAvailableToken 获取可用设备注册码/接入码
func (r *deviceRepo) GetAvailableToken(ctx context.Context) (*biz.DeviceToken, error) {
	token, err := r.data.db.DeviceToken(ctx).Query().
		Where(devicetoken.CreatedAtGT(time.Now().Add(-24 * time.Hour))).
		Where(devicetoken.DeviceExtID("")).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.DeviceToken{
		DeviceExtId: token.DeviceExtID,
		Token:       token.Token,
	}, nil
}

// UpdateState 更新设备状态
func (r *deviceRepo) UpdateState(ctx context.Context, dvcID uint64, state biz.DeviceState) error {
	_, err := r.data.db.Device(ctx).UpdateOneID(dvcID).
		SetState(state).
		Save(ctx)
	return err
}

// FindToken 查询设备注册码/接入码
func (r *deviceRepo) FindToken(ctx context.Context, token string) (*biz.DeviceToken, error) {
	t, err := r.data.db.DeviceToken(ctx).Query().
		Where(devicetoken.Token(token)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.DeviceToken{
		DeviceExtId: t.DeviceExtID,
		Token:       t.Token,
	}, nil
}

// SaveToken 保存设备注册码/接入码
func (r *deviceRepo) SaveToken(ctx context.Context, token *biz.DeviceToken) error {
	_, err := r.data.db.DeviceToken(ctx).Create().
		SetDeviceExtID(token.DeviceExtId).
		SetToken(token.Token).
		Save(ctx)
	return err
}

// BindTokenToDevice 更新设备注册码/接入码,绑定设备
func (r *deviceRepo) BindTokenToDevice(ctx context.Context, token *biz.DeviceToken) error {
	_, err := r.data.db.DeviceToken(ctx).Update().
		Where(devicetoken.Token(token.Token)).
		Where(devicetoken.DeviceExtID("")). // 仅限未绑定设备时更新
		SetDeviceExtID(token.DeviceExtId).
		Save(ctx)
	return err
}

// 分页
func (r *deviceRepo) buildQueryPagination(query *ent.DeviceQuery, filter *biz.DeviceListFilter) *ent.DeviceQuery {
	if filter == nil || filter.Pagination == nil {
		return query
	}

	pagination := filter.Pagination
	return query.Offset(pagination.Offset()).Limit(pagination.PageSize)
}

// 查询条件
func (r *deviceRepo) buildQueryByFilter(query *ent.DeviceQuery, filter *biz.DeviceListFilter) *ent.DeviceQuery {
	// 模糊查询条件
	if filter.Name != "" {
		query = query.Where(device.DisplayNameContains(filter.Name))
	}
	if filter.SerialNo != "" {
		query = query.Where(device.SerialNoContains(filter.SerialNo))
	}

	// 精确查询条件
	if filter.State != "" {
		query = query.Where(device.StateEQ(filter.State))
	}
	if filter.Type != "" {
		query = query.Where(device.TypeEQ(filter.Type))
	}
	// 范围查询条件
	if len(filter.DeviceIds) > 0 {
		query = query.Where(device.IDIn(filter.DeviceIds...))
	}
	if len(filter.Models) > 0 {
		query = query.Where(device.ModelIn(filter.Models...))
	}
	return query
}

// 预加载查询
func (r *deviceRepo) buildQueryByOption(query *ent.DeviceQuery, option *biz.DeviceOption) *ent.DeviceQuery {
	if option == nil {
		return query
	}

	if option.PreloadCamera {
		query = query.WithCamera()
	}
	return query
}

// DeviceEntArrToBiz 将device数组从ent层转化为biz层结构
func DeviceEntArrToBiz(arr []*ent.Device) []*biz.Device {
	return convert.ToArr[ent.Device, biz.Device](entToBiz, arr)
}

// entToBiz 将device从ent层转化为biz层结构
func entToBiz(dvc *ent.Device) *biz.Device {
	deviceAttr := getDeviceAttr(dvc.DeviceInfo)
	bizDevice := &biz.Device{
		Id:             dvc.ID,
		Name:           dvc.Name,
		DisplayName:    dvc.DisplayName,
		SerialNo:       dvc.SerialNo,
		Mac:            dvc.MAC,
		ZoneName:       dvc.ZoneName,
		Type:           dvc.Type,
		ZoneId:         dvc.ZoneID,
		WorkspaceId:    dvc.WorkspaceID,
		ExtId:          dvc.ExtID,
		State:          dvc.State,
		CreatedAt:      dvc.CreatedAt,
		UpdatedAt:      dvc.UpdatedAt,
		DeletedAt:      dvc.DeletedAt,
		ActivatedAt:    dvc.ActivatedAt,
		EquipId:        dvc.EquipID,
		EquipPassword:  dvc.EquipPassword,
		DeviceInfo:     dvc.DeviceInfo,
		DeviceAttrInfo: deviceAttr,
		AuthDeadline:   dvc.AuthDeadline,
		Model:          dvc.Model,
		TenantId:       dvc.TenantID,
		AccessOrgList:  dvc.AccessOrgList,
	}

	for _, camera := range dvc.Edges.Camera {
		bizDevice.Cameras = append(bizDevice.Cameras, CameraEntToBiz(camera))
	}
	return bizDevice
}

// biz设备转为ent设备
func bizToEnt(dvc *biz.Device) *ent.Device {
	deviceAttr := getDeviceAttr(dvc.DeviceInfo)
	return &ent.Device{
		ID:            dvc.Id,
		Name:          dvc.Name,
		DisplayName:   dvc.DisplayName,
		SerialNo:      dvc.SerialNo,
		MAC:           dvc.Mac,
		ZoneName:      dvc.ZoneName,
		Type:          dvc.Type,
		ZoneID:        dvc.ZoneId,
		WorkspaceID:   dvc.WorkspaceId,
		ExtID:         dvc.ExtId,
		State:         dvc.State,
		EquipID:       dvc.EquipId,
		EquipPassword: dvc.EquipPassword,
		DeviceInfo:    dvc.DeviceInfo,
		ActivatedAt:   dvc.ActivatedAt,
		AuthDeadline:  dvc.AuthDeadline,
		Model:         deviceAttr.Model,
		TenantID:      dvc.TenantId,
		AccessOrgList: dvc.AccessOrgList,
	}
}

func getDeviceAttr(deviceInfo string) *biz.DeviceAttr {
	deviceAttr := new(biz.DeviceAttr)
	json.Unmarshal([]byte(deviceInfo), &deviceAttr)
	return deviceAttr
}

func (r *deviceRepo) CountByStatus(ctx context.Context) ([]*biz.DeviceStatusCount, error) {
	// 数据返回
	var deviceData []*biz.DeviceStatusCount

	// 数据渲染
	err := r.data.db.Device(ctx).
		Query().
		GroupBy(device.FieldState).
		Aggregate(ent.Count()).
		Scan(ctx, &deviceData)
	if err != nil {
		return nil, err
	}
	return deviceData, nil
}
