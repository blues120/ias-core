package data

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/algorithm"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/camera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/task"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/taskcamera"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
)

const (
	redisKeyStartTimesCache = "vss:start_times" // 缓存停止中的任务用于状态显示，Pod 无此状态
)

const (
	eventChan = "ias-core:task_event"
)

type taskRepo struct {
	data *Data
	tx   biz.Transaction
	log  *log.Helper
}

// NewTaskRepo 实例化taskRepo
func NewTaskRepo(data *Data, trans biz.Transaction, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		tx:   trans,
		log:  log.NewHelper(logger),
	}
}

// TaskEntToBiz 将task从ent层转化为biz层结构
func TaskEntToBiz(ta *ent.Task) *biz.Task {
	bizTask := &biz.Task{
		Id:   ta.ID,
		Name: ta.Name,
		Type: ta.Type,
		Algo: &biz.TaskAlgo{
			Id:       ta.AlgoID,
			Interval: ta.AlgoInterval,
			Extra:    ta.AlgoExtra,
		},
		Device:        &biz.Device{Id: ta.DeviceID},
		Status:        ta.Status,
		Extend:        ta.Extend,
		ParentId:      ta.ParentID,
		IsWarn:        ta.IsWarn,
		AlgoConfig:    ta.AlgoConfig,
		CreatedAt:     ta.CreatedAt,
		UpdatedAt:     ta.UpdatedAt,
		DeletedAt:     ta.DeletedAt,
		Reason:        ta.Reason,
		TenantID:      ta.TenantID,
		AccessOrgList: ta.AccessOrgList,
	}
	if lst := ta.LastStartTime; lst != nil && lst.Valid {
		bizTask.LastStartTime = &lst.Time
	} else {
		bizTask.LastStartTime = nil
	}

	if device := ta.Edges.Device; device != nil {
		bizTask.Device = &biz.Device{
			Id:            device.ID,
			Name:          device.Name,
			DisplayName:   device.DisplayName,
			Type:          device.Type,
			ExtId:         device.ExtID,
			SerialNo:      device.SerialNo,
			State:         device.State,
			Mac:           device.MAC,
			ZoneName:      device.ZoneName,
			ZoneId:        device.ZoneID,
			WorkspaceId:   device.WorkspaceID,
			CreatedAt:     device.CreatedAt,
			UpdatedAt:     device.UpdatedAt,
			DeletedAt:     device.DeletedAt,
			Cameras:       make([]*biz.Camera, 0),
			EquipId:       device.EquipID,
			EquipPassword: device.EquipPassword,
			Model:         device.Model,
		}
	}

	if algo := ta.Edges.Algorithm; algo != nil {
		bizTask.Algo.Algorithm = AlgoEntToBiz(algo)
	}

	cameras := map[uint64]*biz.Camera{}
	for _, c := range ta.Edges.Camera {
		cameras[c.ID] = CameraEntToBiz(c)
	}
	for _, c := range ta.Edges.TaskCamera {
		bizTaskCamera := biz.TaskCamera{
			Id:          c.CameraID,
			Camera:      cameras[c.CameraID],
			MultiImgBox: c.MultiImgBox,
		}

		bizTask.Cameras = append(bizTask.Cameras, bizTaskCamera)
	}

	return bizTask
}

// TaskEntArrToBiz 将task数组从ent层转化为biz层结构
func TaskEntArrToBiz(arr []*ent.Task) []*biz.Task {
	return convert.ToArr[ent.Task, biz.Task](TaskEntToBiz, arr)
}

// TaskBizToEnt 将task从biz层转化为ent层结构
func TaskBizToEnt(bt *biz.Task) *ent.Task {
	task := &ent.Task{
		CreatedAt:  bt.CreatedAt,
		UpdatedAt:  bt.UpdatedAt,
		Name:       bt.Name,
		Type:       bt.Type,
		Extend:     bt.Extend,
		Status:     bt.Status,
		ParentID:   bt.ParentId,
		IsWarn:     bt.IsWarn,
		AlgoConfig: bt.AlgoConfig,
	}

	if algo := bt.Algo; algo != nil {
		task.AlgoID = algo.Id
		task.AlgoInterval = algo.Interval
		task.AlgoExtra = algo.Extra
		if algo.Algorithm != nil {
			task.AlgoGroupID = algo.Algorithm.AlgoGroupID
		}
	}

	if bt.Device != nil {
		task.DeviceID = bt.Device.Id
	}

	if bt.LastStartTime != nil {
		task.LastStartTime = &sql.NullTime{
			Time:  *bt.LastStartTime,
			Valid: true,
		}
	}

	return task
}

// CheckExists 检查 task 的字段值是否重复
func (r *taskRepo) CheckExists(ctx context.Context, req *biz.CheckTaskExistsRequest) (bool, error) {
	query := r.query(ctx).Where(task.NameEQ(req.Name))
	if req.ExcludeId > 0 {
		query = query.Where(task.IDNEQ(req.ExcludeId))
	}
	if req.DeviceID > 0 {
		query = query.Where(task.DeviceIDEQ(req.DeviceID))
	}
	return query.Exist(ctx)
}

// Save 创建 task
func (r *taskRepo) Save(ctx context.Context, ca *biz.Task) (uint64, error) {
	var taskId uint64

	err := r.tx.InTx(ctx, func(ctx context.Context) error {
		task := TaskBizToEnt(ca)
		task.Status = biz.TaskStatusStopped
		dbTask, err := r.data.db.Task(ctx).Create().SetTask(task).Save(ctx)
		if err != nil {
			return err
		}
		taskId = dbTask.ID

		// 创建关联的 task_camera
		return r.createTaskCameras(ctx, taskId, ca.Cameras)
	})

	return taskId, err
}

// createTaskCameras 创建关联的 task_camera
func (r *taskRepo) createTaskCameras(ctx context.Context, taskId uint64, cameras []biz.TaskCamera) error {
	var taskCameras []*ent.TaskCameraCreate
	for _, tc := range cameras {
		taskCamera := r.data.db.TaskCamera(ctx).Create()
		taskCamera.SetCameraID(tc.Camera.Id).
			SetTaskID(taskId).
			SetMultiImgBox(tc.MultiImgBox)
		taskCameras = append(taskCameras, taskCamera)
	}
	return r.data.db.TaskCamera(ctx).CreateBulk(taskCameras...).Exec(ctx)
}

// Update 更新
func (r *taskRepo) Update(ctx context.Context, taskId uint64, ta *biz.Task) (bool, error) {
	_, err := r.get(ctx, taskId, nil)
	if err != nil {
		return false, err
	}

	if err := r.tx.InTx(ctx, func(ctx context.Context) error {
		update := r.data.db.Task(ctx).UpdateOneID(taskId).
			SetName(ta.Name).SetExtend(ta.Extend).SetIsWarn(ta.IsWarn).SetAlgoConfig(ta.AlgoConfig)

		if ta.Algo != nil {
			update = update.SetAlgoID(ta.Algo.Id).
				SetAlgoInterval(ta.Algo.Interval).
				SetAlgoExtra(ta.Algo.Extra)
			if ta.Algo.Algorithm != nil {
				update = update.SetAlgoGroupID(ta.Algo.Algorithm.AlgoGroupID)
			}
		}

		if err = update.Exec(ctx); err != nil {
			return err
		}

		// 在任务管理关联设备表中删除
		if err := r.deleteTaskCameras(ctx, taskId); err != nil {
			return err
		}

		// 创建新的关联关系
		if err := r.createTaskCameras(ctx, taskId, ta.Cameras); err != nil {
			return err
		}

		// 事件通知
		ta.Id = taskId
		event, err := json.Marshal(&biz.TaskEvent{
			Type: biz.TaskUpdate,
			Task: *ta,
		})
		if err != nil {
			return err
		}
		return r.data.rdb.Publish(ctx, eventChan, event).Err()
	}); err != nil {
		r.log.Errorf("update task error: %v", err)
		return false, err
	}
	return true, nil
}

func (r *taskRepo) UpdateStatus(ctx context.Context, id uint64, status biz.TaskStatus) error {
	query := r.data.db.Task(ctx).UpdateOneID(id).SetStatus(status)
	// 任务状态为运行时，更新最后一次启动时间
	if status == biz.TaskStatusRunning {
		query.SetLastStartTime(&sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		})
	}
	return query.Exec(ctx)
}

func (r *taskRepo) UpdateStatusAndReason(ctx context.Context, id uint64, status biz.TaskStatus, reason string) error {
	query := r.data.db.Task(ctx).UpdateOneID(id).Where(task.Or(task.StatusNEQ(status), task.ReasonNEQ(reason))).SetStatus(status).SetReason(reason)
	// 任务状态为运行时，更新最后一次启动时间
	if status == biz.TaskStatusRunning {
		query.SetLastStartTime(&sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		})
	}
	err := query.Exec(ctx)
	if ent.IsNotFound(err) { //当状态不变时，where条件会过滤任务
		return nil

	}
	return err
}

// Delete 删除
func (r *taskRepo) Delete(ctx context.Context, taskId uint64) error {
	err := r.tx.InTx(ctx, func(ctx context.Context) error {
		// 删除任务
		if err := r.data.db.Task(ctx).DeleteOneID(taskId).Exec(ctx); err != nil {
			return err
		}

		// 删除关联的摄像头
		return r.deleteTaskCameras(ctx, taskId)
	})

	return err
}

func (r *taskRepo) deleteTaskCameras(ctx context.Context, taskId uint64) error {
	where := predicate.TaskCamera(sql.FieldEQ(taskcamera.FieldTaskID, taskId))
	_, err := r.data.db.TaskCamera(ctx).Delete().Where(where).Exec(ctx)
	return err
}

// BatchDelete 批量删除
func (r *taskRepo) BatchDelete(ctx context.Context, taskIds []uint64) (int, error) {
	var effected int

	err := r.tx.InTx(ctx, func(ctx context.Context) error {
		// 删除任务
		taskWhere := predicate.Task(sql.FieldIn(taskcamera.FieldTaskID, taskIds))
		total, err := r.data.db.Task(ctx).Delete().Where(taskWhere).Exec(ctx)
		if err != nil {
			return err
		}
		// 成功删除的数量
		effected = total

		// 删除关联的摄像头
		where := predicate.TaskCamera(sql.FieldIn(taskcamera.FieldTaskID, taskIds))
		_, err = r.data.db.TaskCamera(ctx).Delete().Where(where).Exec(ctx)

		return err
	})

	return effected, err
}

// Find 查询
func (r *taskRepo) Find(ctx context.Context, taskId uint64, option *biz.TaskOption) (*biz.Task, error) {
	task, err := r.get(ctx, taskId, option)
	if err != nil {
		return nil, err
	}

	return TaskEntToBiz(task), nil
}

// get 根据ID获取task
func (r *taskRepo) get(ctx context.Context, taskId uint64, option *biz.TaskOption) (*ent.Task, error) {
	query := r.data.db.Task(ctx).Query().Where(task.IDEQ(taskId))

	query = r.buildQueryByOption(query, option)

	// 忽略软删除
	if option != nil && option.IncludeDeleted {
		ctx = mixin.SkipSoftDelete(ctx)
	}

	return query.First(ctx)
}

func (r *taskRepo) ListByCameraId(ctx context.Context, cameraId uint64, option *biz.TaskOption) ([]*biz.Task, error) {
	query := r.query(ctx).Where(task.HasTaskCameraWith(taskcamera.CameraIDEQ(cameraId)))

	query = r.buildQueryByOption(query, option)
	arr, err := query.Clone().All(ctx)
	if err != nil {
		return nil, err
	}

	return TaskEntArrToBiz(arr), nil
}

// List 批量查询
func (r *taskRepo) List(ctx context.Context, filter *biz.TaskListFilter, option *biz.TaskOption) ([]*biz.Task, int, error) {
	query := r.query(ctx)
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)
	}
	if option != nil {
		query = r.buildQueryByOption(query, option)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	arr, err := r.buildQueryPagination(query, filter).Clone().Modify(func(s *sql.Selector) {
		s.OrderBy("`status` ASC", "`last_start_time` IS NULL DESC", "`last_start_time` DESC", "`updated_at` DESC", "`created_at` DESC")
	}).All(ctx)

	if err != nil {
		return nil, 0, err
	}

	return TaskEntArrToBiz(arr), total, nil
}

func (r *taskRepo) buildQueryByFilter(query *ent.TaskQuery, filter *biz.TaskListFilter) *ent.TaskQuery {
	// 模糊查询条件
	if len(filter.IdIn) > 0 {
		query = query.Where(task.IDIn(filter.IdIn...))
	}
	if filter.NameContain != "" {
		query = query.Where(task.NameContains(filter.NameContain))
	}
	if filter.IdOrNameContain != "" {
		var subQuery []predicate.Task = []predicate.Task{
			task.NameContains(filter.IdOrNameContain),
			predicate.Task(sql.FieldContains(task.FieldID, filter.IdOrNameContain)),
		}
		query = query.Where(task.Or(subQuery...))
	}
	if filter.Status != "" {
		query = query.Where(task.StatusEQ(filter.Status))
	}
	if filter.StatusNotEq != "" {
		query = query.Where(task.StatusNEQ(filter.StatusNotEq))
	}
	if filter.DeviceId > 0 {
		query = query.Where(task.DeviceIDEQ(filter.DeviceId))
	}
	if filter.TaskExtra != "" {
		query = query.Where(task.AlgoExtraContains(filter.TaskExtra))
	}
	if filter.ParentId != "" {
		query = query.Where(task.ParentIDEQ(filter.ParentId))
	}
	if filter.ExtAlgoIdEq != "" {
		query = query.Where(task.HasAlgorithmWith(algorithm.AlgoIDEQ(filter.ExtAlgoIdEq)))
	}

	if len(filter.AlgorithmIdIn) > 0 {
		query = query.Where(task.HasAlgorithmWith(algorithm.IDIn(filter.AlgorithmIdIn...)))
	}

	if len(filter.CameraIdIn) > 0 {
		query = query.Where(task.HasCameraWith(camera.IDIn(filter.CameraIdIn...)))
	}

	if len(filter.DeviceIdIn) > 0 {
		query = query.Where(task.DeviceIDIn(filter.DeviceIdIn...))
	}

	return query
}

// buildQueryByOption 添加预加载字段
func (r *taskRepo) buildQueryByOption(query *ent.TaskQuery, option *biz.TaskOption) *ent.TaskQuery {
	if option != nil {
		if option.PreloadAlgorithm {
			query = query.WithAlgorithm()
		}
		if option.PreloadCamera {
			query = query.WithCamera()
		}
		if option.PreloadTaskCamera {
			query = query.WithTaskCamera()
		}
		if option.PreloadDevice {
			query = query.WithDevice()
		}
	}

	return query
}

// 分页
func (r *taskRepo) buildQueryPagination(query *ent.TaskQuery, filter *biz.TaskListFilter) *ent.TaskQuery {
	if filter == nil || filter.Pagination == nil {
		return query
	}

	pagination := filter.Pagination
	return query.Offset(pagination.Offset()).Limit(pagination.PageSize)
}

// CountByTaskStatus 根据任务状态查询任务数量
func (r *taskRepo) CountByTaskStatus(ctx context.Context) (map[biz.TaskStatus]uint64, error) {
	var rows []StatusCount
	err := r.query(ctx).GroupBy(task.FieldStatus).Aggregate(ent.Count()).Scan(ctx, &rows)
	if err != nil {
		return nil, err
	}

	result := statusCountToMap(rows)

	return result, nil
}

// CountTaskCameraByStatus 根据状态统计任务摄像头数量
func (r *taskRepo) CountTaskCameraByStatus(ctx context.Context, statuses []biz.TaskStatus) (map[biz.TaskStatus]uint64, error) {
	var rows []StatusCount
	query := r.query(ctx).Where(task.HasTaskCameraWith())
	if len(statuses) > 0 {
		query = query.Where(task.StatusIn(statuses...))
	}

	err := query.GroupBy(task.FieldStatus).
		Aggregate(ent.Count()).Scan(ctx, &rows)
	if err != nil {
		return nil, err
	}

	result := statusCountToMap(rows)
	return result, nil
}

// FindTaskCameraById 获取任务摄像头
func (r *taskRepo) FindTaskCameraById(ctx context.Context, taskId uint64, cameraId uint64) (*biz.TaskCamera, error) {
	tc, err := r.data.db.TaskCamera(ctx).Query().
		Where(taskcamera.CameraID(cameraId)).WithCamera().
		Where(taskcamera.TaskID(taskId)).First(ctx)
	if err != nil {
		return nil, err
	}

	var bizCamera *biz.Camera
	if tc.Edges.Camera != nil {
		bizCamera = CameraEntToBiz(tc.Edges.Camera)
	}

	return &biz.TaskCamera{
		Id:          0,
		Camera:      bizCamera,
		MultiImgBox: tc.MultiImgBox,
	}, nil
}

// StatusCount 状态数量
type StatusCount struct {
	Status biz.TaskStatus `json:"status"`
	Count  uint64         `json:"count"`
}

func statusCountToMap(rows []StatusCount) map[biz.TaskStatus]uint64 {
	result := make(map[biz.TaskStatus]uint64)
	for _, row := range rows {
		result[row.Status] = row.Count
	}

	return result
}

func (r *taskRepo) query(ctx context.Context) *ent.TaskQuery {
	return r.data.db.Task(ctx).Query()
}

// GetTaskStartTimes 获取任务开始次数
func (r *taskRepo) GetTaskStartTimes(ctx context.Context, parentId string) (int, error) {
	timesStr, err := r.data.rdb.HGet(context.Background(), redisKeyStartTimesCache, parentId).Result()
	if err != nil {
		return 0, err
	}
	times, err := strconv.Atoi(timesStr)
	return times, err
}

// CountAlgoInUse 统计使用中的算法数量
func (r *taskRepo) CountAlgoInUse(ctx context.Context) (int, error) {
	// 数据返回
	var total int
	total, err := r.data.db.Task(ctx).Query().
		Where(task.StatusEQ(biz.TaskStatusRunning)).
		Unique(true).
		Select(task.FieldAlgoID).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// GetReason 获取任务状态变更原因
func (r *taskRepo) GetReason(ctx context.Context, taskId uint64) (string, error) {
	task, err := r.data.db.Task(ctx).Get(ctx, taskId)
	if err != nil {
		return "", err
	}
	return task.Reason, nil
}

// UpdateReason 更新任务状态变更原因
func (r *taskRepo) UpdateReason(ctx context.Context, taskId uint64, reason string) error {
	_, err := r.data.db.Task(ctx).UpdateOneID(taskId).SetReason(reason).Save(ctx)
	return err
}

// Events 获取任务事件
func (r *taskRepo) Events(ctx context.Context, size int) <-chan biz.TaskEvent {
	source := r.data.rdb.Subscribe(ctx, eventChan).Channel()
	target := make(chan biz.TaskEvent, size)
	go func() {
		for {
			msg, ok := <-source
			// 连接断开时退出
			if !ok {
				close(target)
				return
			}
			var taskEvent biz.TaskEvent
			_ = json.Unmarshal([]byte(msg.Payload), &taskEvent)
			select {
			case target <- taskEvent:
			default:
				<-target
				target <- taskEvent
			}
		}
	}()
	return target
}
