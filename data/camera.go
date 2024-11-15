package data

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz/streaming"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/camera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/device"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/task"
	"gitlab.ctyuncdn.cn/ias/ias-core/errors"
	"gitlab.ctyuncdn.cn/ias/ias-core/middleware"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
	"github.com/blues120/ias-kit/oss"
)

type cameraRepo struct {
	data *Data
	oss  oss.Oss

	log *log.Helper
}

func NewCameraRepo(data *Data, oss oss.Oss, logger log.Logger) biz.CameraRepo {
	return &cameraRepo{
		data: data,
		oss:  oss,
		log:  log.NewHelper(logger),
	}
}

func (r *cameraRepo) Save(ctx context.Context, ca *biz.Camera) (uint64, error) {
	data, err := r.data.db.Camera(ctx).Create().SetCamera(r.bizToEnt(ca)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *cameraRepo) Update(ctx context.Context, id uint64, ca *biz.Camera) error {
	if string(ca.Type) == "" {
		ca.Type = biz.MediaTypeCamera
	}
	query := r.data.db.Camera(ctx).
		UpdateOneID(id).
		SetName(ca.Name).
		SetPosition(ca.Position).
		SetRegion(ca.Region).
		SetRegionStr(ca.RegionStr).
		SetLongitude(ca.Longitude).
		SetLatitude(ca.Latitude).
		SetCustomNumber(ca.CustomNumber).
		SetChannelID(ca.ChannelId).
		SetSerialNumber(ca.SerialNumber).
		SetPoleNumber(ca.PoleNumber).
		SetDeviceDescription(ca.DeviceDescription).
		SetScene(ca.Scene).
		SetPlace(ca.Place).
		SetSpType(ca.StreamingProtocol.Type()).
		SetSpSource(ca.StreamingProtocol.Source()).
		SetTransType(ca.TransType).
		SetDeviceIP(ca.DeviceIP).
		SetDevicePort(ca.DevicePort).
		SetGBID(ca.GbId).
		SetSipUser(ca.SipUser).
		SetSipPassword(ca.SipPassword).
		SetGBDeviceType(ca.GBDeviceType).
		SetGBChannelID(ca.ChannelGbId).
		SetType(ca.Type)
	if ca.Status != "" {
		query = query.SetStatus(ca.Status)
	}
	if ca.StreamingInfo != nil {
		info := ca.StreamingInfo
		query = query.SetSpCodecName(info.CodecName).SetSpWidth(info.Width).SetSpHeight(info.Height)
	}
	return query.Exec(ctx)
}

func (r *cameraRepo) UpdateStatus(ctx context.Context, id uint64, status biz.CameraStatus) error {
	// TODO: 这里最好的方法是跳过更新 updatedAt 字段, 但是暂时没找到方法
	return r.data.db.Camera(ctx).UpdateOneID(id).Where(camera.StatusNEQ(status)).SetStatus(status).Exec(ctx)
}

// UpdateStreamInfo 更新摄像机流信息
func (r *cameraRepo) UpdateStreamInfo(ctx context.Context, id uint64, info *streaming.Info) error {
	if info == nil {
		return nil
	}
	return r.data.db.Camera(ctx).UpdateOneID(id).
		SetSpCodecName(info.CodecName).SetSpHeight(info.Height).SetSpWidth(info.Width).Exec(ctx)
}

func (r *cameraRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.Camera(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *cameraRepo) BatchDelete(ctx context.Context, ids []uint64) (int, error) {
	return r.data.db.Camera(ctx).Delete().Where(camera.IDIn(ids...)).Exec(ctx)
}

func (r *cameraRepo) Find(ctx context.Context, id uint64, option *biz.CameraOption) (*biz.Camera, error) {
	query := r.data.db.Camera(ctx).Query()
	if option != nil {
		query = r.buildQueryByOption(query, option)
	}
	ca, err := query.Where(camera.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorCameraNotFound("摄像机(id=%d)不存在", id)
		}
		return nil, err
	}
	return CameraEntToBiz(ca), nil
}

func (r *cameraRepo) Exist(ctx context.Context, field *biz.CameraExistField, option *biz.CameraOption) (*biz.Camera, error) {
	var subQuery []predicate.Camera
	if field.Name != "" {
		subQuery = append(subQuery, camera.NameEQ(field.Name))
	}
	if field.SerialNumber != "" {
		subQuery = append(subQuery, camera.SerialNumberEQ(field.SerialNumber))
	}
	if field.CustomNumber != 0 {
		subQuery = append(subQuery, camera.CustomNumberEQ(field.CustomNumber))
	}
	if field.StreamingProtocol != nil {
		subQuery = append(subQuery, camera.SpSource(field.StreamingProtocol.Source()))
	}
	if field.ChannelId != "" {
		subQuery = append(subQuery, camera.ChannelIDEQ(field.ChannelId))
	}
	if field.GbId != "" {
		subQuery = append(subQuery, camera.GBIDEQ(field.GbId))
	}
	if field.ChannelId != "" {
		subQuery = append(subQuery, camera.ChannelIDEQ(field.ChannelId))
	}
	if field.ChannelGbId != "" {
		subQuery = append(subQuery, camera.GBChannelIDEQ(field.ChannelGbId))
	}
	if field.GbId != "" {
		subQuery = append(subQuery, camera.GBIDEQ(field.GbId))
	}
	if field.GBDeviceType != "" {
		subQuery = append(subQuery, camera.GBDeviceTypeEQ(field.GBDeviceType))
	}
	query := r.data.db.Camera(ctx).Query()
	if option != nil {
		query = r.buildQueryByOption(query, option)
	}
	ca, err := query.Where(camera.Or(subQuery...)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorCameraNotFound("摄像机不存在")
		}
		return nil, err
	}
	return CameraEntToBiz(ca), nil
}

func (r *cameraRepo) List(ctx context.Context, filter *biz.CameraListFilter, option *biz.CameraOption) ([]*biz.Camera, int, error) {
	query := r.data.db.Camera(ctx).Query()
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

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	arr, err := r.buildQueryPagination(query, filter).Clone().Order(
		ent.Asc(camera.FieldStatus),
		ent.Desc(camera.FieldUpdatedAt, camera.FieldCreatedAt, camera.FieldID),
	).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return CameraEntArrToBiz(arr), total, nil
}

func (r *cameraRepo) CountRegion(ctx context.Context) ([]*biz.CameraRegionCount, error) {
	var v = make([]*biz.CameraRegionCount, 0)
	err := r.data.db.Camera(ctx).
		Query().
		GroupBy(camera.FieldRegion).
		Aggregate(ent.Count()).
		Scan(ctx, &v)
	return v, err
}

func (r *cameraRepo) CountByBindTask(ctx context.Context, ids []uint64) (int, error) {
	return r.data.db.Camera(ctx).
		Query().
		Select(camera.FieldID).
		Where(camera.IDIn(ids...)).
		Where(camera.HasTask()).
		Count(ctx)
}

func (r *cameraRepo) buildQueryByFilter(query *ent.CameraQuery, filter *biz.CameraListFilter) *ent.CameraQuery {
	// 模糊查询条件
	if filter.NameContain != "" {
		query = query.Where(camera.NameContains(filter.NameContain))
	}
	if filter.PositionContain != "" {
		query = query.Where(camera.PositionContains(filter.PositionContain))
	}
	if len(filter.SceneContain) > 0 {
		var subQuery []predicate.Camera
		for _, item := range filter.SceneContain {
			subQuery = append(subQuery, camera.SceneContains(item))
		}
		query = query.Where(camera.Or(subQuery...))
	}
	if filter.Places != "" {
		query = query.Where(camera.Place(filter.Places))
	}
	// 精确查询条件
	if filter.DeviceID > 0 {
		query = query.Where(camera.HasDeviceWith(device.ID(filter.DeviceID)))
	}
	if filter.Status != "" {
		query = query.Where(camera.StatusEQ(filter.Status))
	}
	if filter.Region != "" {
		query = query.Where(camera.RegionEQ(filter.Region))
	}
	if filter.SerialNumber != "" {
		query = query.Where(camera.SerialNumberEQ(filter.SerialNumber))
	}
	if filter.StreamingProtocolType != "" {
		query = query.Where(camera.SpTypeEQ(filter.StreamingProtocolType))
	}
	if filter.Place != "" {
		query = query.Where(camera.PlaceEQ(filter.Place))
	}
	if filter.Type != "" {
		query = query.Where(camera.TypeEQ(filter.Type))
	}
	// 范围查询条件
	if len(filter.CameraIds) > 0 {
		query = query.Where(camera.IDIn(filter.CameraIds...))
	}
	if len(filter.AlgorithmIdIn) > 0 {
		query = query.Where(camera.HasTaskWith(task.AlgoIDIn(filter.AlgorithmIdIn...)))
	}
	// 特殊查询条件
	if len(filter.CustomNumberIn) > 0 {
		// eg: where custom_number = 1 or (custom_number between 1 and 2) or (custom_number between 3 and 4)
		var subQuery []predicate.Camera
		for _, item := range filter.CustomNumberIn {
			if len(item) == 1 {
				subQuery = append(subQuery, camera.CustomNumberEQ(item[0]))
			} else {
				subQuery = append(subQuery, camera.And(camera.CustomNumberGTE(item[0]), camera.CustomNumberLTE(item[1])))
			}
		}
		query = query.Where(camera.Or(subQuery...))
	}

	return query
}

func (r *cameraRepo) buildQueryByOption(query *ent.CameraQuery, option *biz.CameraOption) *ent.CameraQuery {
	if option.PreloadTask {
		query = query.WithTask(func(query *ent.TaskQuery) {
			query.WithAlgorithm()
		})
	}
	if option.PreloadDeviceCamera {
		query = query.WithDeviceCamera(func(query *ent.DeviceCameraQuery) {
			query.WithCamera()
		})
	}
	if option.PreloadDevice {
		query = query.WithDevice(func(query *ent.DeviceQuery) {
			query.WithCamera()
		})
	}

	return query
}

// 分页
func (r *cameraRepo) buildQueryPagination(query *ent.CameraQuery, filter *biz.CameraListFilter) *ent.CameraQuery {
	if filter == nil || filter.Pagination == nil {
		return query
	}

	pagination := filter.Pagination
	return query.Offset(pagination.Offset()).Limit(pagination.PageSize)
}

func (r *cameraRepo) bizToEnt(ca *biz.Camera) *ent.Camera {
	if string(ca.Type) == "" {
		ca.Type = biz.MediaTypeCamera
	}
	entCamera := &ent.Camera{
		ID:                ca.Id,
		CreatedAt:         ca.CreatedAt,
		UpdatedAt:         ca.UpdatedAt,
		Name:              ca.Name,
		Position:          ca.Position,
		Region:            ca.Region,
		RegionStr:         ca.RegionStr,
		Longitude:         ca.Longitude,
		Latitude:          ca.Latitude,
		CustomNumber:      ca.CustomNumber,
		ChannelID:         ca.ChannelId,
		SerialNumber:      ca.SerialNumber,
		PoleNumber:        ca.PoleNumber,
		DeviceDescription: ca.DeviceDescription,
		Scene:             ca.Scene,
		Place:             ca.Place,
		Status:            ca.Status,
		SpType:            ca.StreamingProtocol.Type(),
		SpSource:          ca.StreamingProtocol.Source(),
		TransType:         ca.TransType,
		DeviceIP:          ca.DeviceIP,
		DevicePort:        ca.DevicePort,
		GBID:              ca.GbId,
		SipUser:           ca.SipUser,
		SipPassword:       ca.SipPassword,
		GBDeviceType:      ca.GBDeviceType,
		GBChannelID:       ca.ChannelGbId,
		UpGBChannelID:     ca.ChannelGbId,
		Type:              ca.Type,
	}
	if info := ca.StreamingInfo; info != nil {
		entCamera.SpCodecName = info.CodecName
		entCamera.SpWidth = info.Width
		entCamera.SpHeight = info.Height
	}

	return entCamera
}

func getRedisCameraImportKey(key string) string {
	return "ias:camera:import:" + key + ".xlsx"
}

type CameraImportProgress struct {
	Total   int `redis:"total"`
	Success int `redis:"success"`
	Fail    int `redis:"fail"`
}

func (r *cameraRepo) Import(ctx context.Context, importId string, importer biz.CameraImporter, cameraSlice interface{}) error {
	if reflect.TypeOf(cameraSlice).Kind() != reflect.Slice {
		return errors.ErrorInvalidParam("invalid camera slice")
	}
	cas := reflect.ValueOf(cameraSlice)

	progress := CameraImportProgress{Total: cas.Len()}
	key := getRedisCameraImportKey(importId)
	if err := r.data.rdb.HSet(ctx, key, progress).Err(); err != nil {
		return err
	}
	if err := r.data.rdb.Expire(ctx, key, time.Second*600).Err(); err != nil {
		return err
	}
	go func() {
		for idx := 0; idx < cas.Len(); idx++ {
			ca := cas.Index(idx)
			deviceId, _ := ctx.Value(middleware.DeviceIdKey{}).(uint64)
			if err := importer.Handler(ca, idx, deviceId); err != nil {
				importer.ErrRecord(err, idx)
				progress.Fail++
			} else {
				progress.Success++
			}
			// 处理最后一条记录时，如果 fail 的数量大于0，需要先上传结果再更新进度，以免下载失败
			if idx == progress.Total-1 && progress.Fail > 0 {
				if result := importer.GetRecord(); result != nil {
					if err := r.oss.Upload(ctx, key, result); err != nil {
						r.log.Errorf("upload camera import result error: %v", err)
					}
				}
			}
			r.data.rdb.HSet(ctx, key, progress)
		}
	}()
	return nil
}

func (r *cameraRepo) QueryImportProgress(ctx context.Context, importId string) (*biz.ImportProgress, error) {
	var i = &CameraImportProgress{}
	key := getRedisCameraImportKey(importId)
	if err := r.data.rdb.HGetAll(ctx, key).Scan(i); err != nil {
		return nil, err
	}
	ret := &biz.ImportProgress{
		Total:   i.Total,
		Success: i.Success,
		Fail:    i.Fail,
	}
	ret.Process, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(i.Success+i.Fail)/float64(i.Total)), 64)
	if ret.Process == 1 && ret.Fail > 0 {
		ret.Url, _ = r.oss.GenerateUrl(ctx, key, 3600*time.Second)
	}
	return ret, nil
}

func CameraEntToBiz(ca *ent.Camera) *biz.Camera {
	sp, _ := streaming.NewProtocol(ca.SpSource, ca.SpType)
	info := &streaming.Info{
		CodecName: ca.SpCodecName,
		Width:     ca.SpWidth,
		Height:    ca.SpHeight,
	}

	data := &biz.Camera{
		Id:                ca.ID,
		Name:              ca.Name,
		Position:          ca.Position,
		Region:            ca.Region,
		RegionStr:         ca.RegionStr,
		Longitude:         ca.Longitude,
		Latitude:          ca.Latitude,
		CustomNumber:      ca.CustomNumber,
		ChannelId:         ca.ChannelID,
		SerialNumber:      ca.SerialNumber,
		PoleNumber:        ca.PoleNumber,
		DeviceDescription: ca.DeviceDescription,
		Scene:             ca.Scene,
		Place:             ca.Place,
		CreatedAt:         ca.CreatedAt,
		UpdatedAt:         ca.UpdatedAt,
		Status:            ca.Status,
		StreamingProtocol: sp,
		StreamingInfo:     info,
		TransType:         ca.TransType,
		DeviceIP:          ca.DeviceIP,
		DevicePort:        ca.DevicePort,
		GbId:              ca.GBID,
		SipUser:           ca.SipUser,
		SipPassword:       ca.SipPassword,
		GBDeviceType:      ca.GBDeviceType,
		ChannelGbId:       ca.GBChannelID,
		Type:              ca.Type,
		TenantID:          ca.TenantID,
		AccessOrgList:     ca.AccessOrgList,
	}

	if ca.Edges.Task != nil {
		data.Task = TaskEntArrToBiz(ca.Edges.Task)
	}

	if ca.Edges.DeviceCamera != nil {
		data.CameraDevice = CameraDeviceEntArrToBiz(ca.Edges.DeviceCamera)
	}

	if ca.Edges.Device != nil {
		data.Device = DeviceEntArrToBiz(ca.Edges.Device)
	}

	return data
}

func CameraDeviceEntArrToBiz(arr []*ent.DeviceCamera) []*biz.CameraDevice {
	return convert.ToArr[ent.DeviceCamera, biz.CameraDevice](CameraDeviceEntToBiz, arr)
}

func CameraDeviceEntToBiz(arr *ent.DeviceCamera) *biz.CameraDevice {
	deviceCamera := &biz.CameraDevice{
		DeviceId: arr.DeviceID,
		CameraId: arr.CameraID,
	}
	return deviceCamera
}

func CameraEntArrToBiz(arr []*ent.Camera) []*biz.Camera {
	return convert.ToArr[ent.Camera, biz.Camera](CameraEntToBiz, arr)
}

func (r *cameraRepo) CountByStatus(ctx context.Context) ([]*biz.CameraStatusCount, error) {
	// 数据返回
	var cameraData []*biz.CameraStatusCount
	// 数据渲染
	err := r.data.db.Camera(ctx).
		Query().
		GroupBy(camera.FieldStatus).
		Aggregate(ent.Count()).
		Scan(ctx, &cameraData)
	if err != nil {
		return nil, err
	}
	return cameraData, nil
}

func (r *cameraRepo) GetChannelList(ctx context.Context) ([]*biz.GBChannel, error) {
	cameras, err := r.data.db.Camera(ctx).
		Query().
		Where(
			camera.SpTypeEQ(streaming.ProtocolTypeGB28181),
			camera.ChannelIDNEQ(""), // 添加条件确保channel_id不为空 为空不参与级联
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// 将查询结果转换为GBChannel结构体
	var channels []*biz.GBChannel
	for _, c := range cameras {
		if c.GBDeviceType == "ipc" {
			channels = append(channels, &biz.GBChannel{
				ChannelId:   c.ChannelID,
				ChannelGbId: c.UpGBChannelID,
			})
		} else {
			channels = append(channels, &biz.GBChannel{
				ChannelId:   c.ChannelID,
				ChannelGbId: c.UpGBChannelID,
			})
		}
	}
	return channels, nil
}

func (r *cameraRepo) BatchUpdateChannel(ctx context.Context, channels []*biz.GBChannel) error {
	// 使用事务逐个更新Camera表中的Channel
	for _, ch := range channels {
		err := r.data.db.Camera(ctx).
			Update().
			Where(camera.ChannelIDEQ(ch.ChannelId)).
			SetUpGBChannelID(ch.ChannelGbId).
			Exec(ctx)
		if err != nil {
			// 如果遇到错误，回滚事务 TD
			return err
		}
	}

	return nil
}

func (r *cameraRepo) ListGBNotPUll(ctx context.Context, sourceFilter string) ([]*biz.Camera, error) {
	cameras, err := r.data.db.Camera(ctx).
		Query().
		Where(
			camera.SpTypeEQ(streaming.ProtocolTypeGB28181),
			camera.SpSourceEQ(sourceFilter),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return CameraEntArrToBiz(cameras), err
}
