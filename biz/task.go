package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/pagination"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
)

// TaskType 任务类型
type TaskType string

// Values provides list valid values for Enum.
func (TaskType) Values() []string {
	return []string{
		string(TaskTypeUnknown),
		string(TaskTypeFrame),
		string(TaskTypeStream),
		string(TaskTypeAudio),
	}
}

const (
	// TaskTypeUnknown 未知
	TaskTypeUnknown TaskType = "unknown"
	// TaskTypeFrame 图片帧
	TaskTypeFrame TaskType = "frame"
	// TaskTypeStream 视频流
	TaskTypeStream TaskType = "stream"
	// TaskTypeAudio 音频流
	TaskTypeAudio TaskType = "audio"

	ServerDevice = "node"
)

// TaskStatus 任务状态
type TaskStatus string

// Values provides list valid values for Enum.
func (TaskStatus) Values() []string {
	return []string{
		string(TaskStatusUnknown),
		string(TaskStatusRunning),
		string(TaskStatusStarting),
		string(TaskStatusInitializing),
		string(TaskStatusFailed),
		string(TaskStatusStopping),
		string(TaskStatusStopped),
	}
}

const (
	TaskStatusUnknown      TaskStatus = "unknown"      // 未知
	TaskStatusRunning      TaskStatus = "running"      // 运行中
	TaskStatusStarting     TaskStatus = "starting"     // 启动中
	TaskStatusInitializing TaskStatus = "initializing" // 正在初始化
	TaskStatusFailed       TaskStatus = "failed"       // 失败
	TaskStatusStopping     TaskStatus = "stopping"     // 停止中
	TaskStatusStopped      TaskStatus = "stopped"      // 已停止
)

type TaskIdStatus struct {
	TaskId uint       `json:"taskId"`
	Status TaskStatus `json:"status"`
	Reason string     `json:"reason"`
}

// TaskAlgo 任务算法关联数据
type TaskAlgo struct {
	Id        uint64  // 算法id
	Interval  float64 // 算法识别频率
	Extra     string  // 算法运行参数
	Algorithm *Algorithm
}

// TaskCamera 任务关联的摄像头
type TaskCamera struct {
	Id          uint64
	Camera      *Camera
	MultiImgBox string `json:"multi_img_box"`
	ModeName    string
}

// Task 任务
type Task struct {
	Id            uint64       // id
	Name          string       // 名称
	Type          TaskType     // 类型
	Algo          *TaskAlgo    // 算法
	Device        *Device      // 任务关联的设备
	Cameras       []TaskCamera // 任务关联的摄像头
	Status        TaskStatus   // 状态
	Reason        string       // 状态失败原因
	Extend        string       // 扩展字段
	ParentId      string       // 父ID, 作为任务下发的ID
	IsWarn        uint32       // 首页是否告警
	AlgoConfig    string       // 算法特有配置
	LastStartTime *time.Time   // 最后启动时间
	CreatedAt     time.Time    // 创建时间
	UpdatedAt     time.Time    // 更新时间
	DeletedAt     time.Time    // 删除时间
	TenantID      string       // 租户id
	AccessOrgList string       // 授权的组织 id 列表，#分隔
}

// TaskListFilter 批量查询过滤条件
type TaskListFilter struct {
	/*
		模糊查询条件
	*/
	NameContain     string
	IdOrNameContain string
	/*
		精确查询条件
	*/
	Status      TaskStatus // 状态
	StatusNotEq TaskStatus // 不等于状态
	DeviceId    uint64     // 设备
	TaskExtra   string     // 算法扩展数据
	ParentId    string     // 任务组ID
	ExtAlgoIdEq string     // 算法extID, 算法表的algoId
	/*
		范围查询条件
	*/
	IdIn          []uint64 // 包含任务ID
	AlgorithmIdIn []uint64 // 包含算法
	CameraIdIn    []uint64 // 包含摄像机
	DeviceIdIn    []uint64 // 包含设备
	/*
		分页
	*/
	Pagination *pagination.Pagination // 分页
}

// CheckTaskExistsRequest 任务字段值重复校验请求参数
type CheckTaskExistsRequest struct {
	Name      string
	ExcludeId uint64
	DeviceID  uint64
}

// TaskOption 查询选项
type TaskOption struct {
	PreloadAlgorithm  bool // 加载关联的任务数据
	PreloadFace       bool // 加载关联的人脸数据
	PreloadVehicle    bool // 加载关联的车辆数据
	PreloadCamera     bool // 加载关联的摄像头数据
	PreloadTaskCamera bool // 加载关联的任务摄像头
	PreloadDevice     bool // 加载关联的任务设备信息

	IncludeDeleted bool // 是否查询包含已删除数据 s
}

type MessageData struct {
	MsgID   string  `json:"msgId"`
	Title   string  `json:"title"`
	Content Content `json:"content"`
}

type Content struct {
	TaskID     string `json:"taskId"`
	TaskStatus string `json:"taskStatus"`
}
type UpdateTaskStatus struct {
	PlatformID  string      `json:"platformId" desc:"对方携带的平台信息"`
	MessageData MessageData `json:"messageData" desc:"消息内容"`
}

type TaskRepo interface {
	// Save 创建任务
	Save(ctx context.Context, ta *Task) (uint64, error)

	// Update 更新任务
	Update(ctx context.Context, id uint64, ta *Task) (bool, error)

	// UpdateStatus 更新任务状态
	UpdateStatus(ctx context.Context, id uint64, status TaskStatus) error

	// Delete 删除任务
	Delete(ctx context.Context, id uint64) error

	// Find 查询任务
	Find(ctx context.Context, id uint64, option *TaskOption) (*Task, error)

	// List 查询任务列表
	List(ctx context.Context, filter *TaskListFilter, option *TaskOption) ([]*Task, int, error)

	// CountByTaskStatus 根据状态统计任务数量
	CountByTaskStatus(ctx context.Context) (map[TaskStatus]uint64, error)

	// CheckExists 检查字段值是否重复
	CheckExists(ctx context.Context, req *CheckTaskExistsRequest) (bool, error)

	// ListByCameraId 根据摄像头过滤任务
	ListByCameraId(ctx context.Context, cameraId uint64, option *TaskOption) ([]*Task, error)

	// CountTaskCameraByStatus 根据状态统计任务摄像头数量
	CountTaskCameraByStatus(ctx context.Context, status []TaskStatus) (map[TaskStatus]uint64, error)

	// FindTaskCameraById 获取任务摄像头
	FindTaskCameraById(ctx context.Context, taskId uint64, cameraId uint64) (*TaskCamera, error)

	// CountAlgoInUse 统计使用中的算法数量
	CountAlgoInUse(ctx context.Context) (int, error)

	GetTaskStartTimes(ctx context.Context, parentId string) (int, error)

	// 获取任务状态变更原因
	GetReason(ctx context.Context, id uint64) (string, error)

	// 更新任务状态变更原因
	UpdateReason(ctx context.Context, id uint64, reason string) error

	// UpdateStatusAndReason 更新任务状态及原因
	UpdateStatusAndReason(ctx context.Context, id uint64, status TaskStatus, reason string) error

	// Events 获取任务事件
	// 当 channel 空间不足时（事件数 >= size），会自动丢弃最早的事件
	Events(ctx context.Context, size int) <-chan TaskEvent
}

type SchedulerRepo interface {
	// Start 启动任务
	Start(ctx context.Context, ta *Task) error

	// Stop 停止任务
	Stop(ctx context.Context, ta *Task) error

	// GetLog 查询任务日志
	GetLog(ctx context.Context, ta *Task, conn *websocket.Conn) error

	// GetStatus 查询任务状态
	GetStatuses(ctx context.Context, taskIdList []uint) ([]TaskIdStatus, error)

	// Close 关闭
	Close() error
}

type SchedulerVSSRepo interface {
	// StartVSS 启动任务
	StartVSS(ctx context.Context, tasks []*Task) error

	// StopVSS 停止任务
	StopVSS(ctx context.Context, ta *Task) error

	// GetLogVSS 查询任务日志
	GetLogVSS(ctx context.Context, ta *Task, conn *websocket.Conn) error
}

type SchedulerRepoSelector func(mode string) SchedulerRepo

type TaskUsecase struct {
	taskRepo         TaskRepo
	selector         SchedulerRepoSelector
	schedulerVSSRepo SchedulerVSSRepo
	conf             *conf.Scheduler

	log *log.Helper
}

func NewTaskUsecase(taskRepo TaskRepo, selector SchedulerRepoSelector, schedulerVSSRepo SchedulerVSSRepo, logger log.Logger) *TaskUsecase {
	return &TaskUsecase{
		taskRepo:         taskRepo,
		selector:         selector,
		schedulerVSSRepo: schedulerVSSRepo,

		log: log.NewHelper(logger),
	}
}

// Create 创建
func (uc *TaskUsecase) Create(ctx context.Context, ca *Task) (uint64, error) {
	if exists, err := uc.taskRepo.CheckExists(ctx, &CheckTaskExistsRequest{
		Name: ca.Name,
	}); err != nil {
		return 0, err
	} else if exists {
		return 0, errors.ErrorInvalidParam("任务名重复")
	}
	return uc.taskRepo.Save(ctx, ca)
}

// Update 更新
func (uc *TaskUsecase) Update(ctx context.Context, id uint64, ca *Task) (bool, error) {
	if exists, err := uc.taskRepo.CheckExists(ctx, &CheckTaskExistsRequest{
		Name:      ca.Name,
		ExcludeId: id,
	}); err != nil {
		return false, err
	} else if exists {
		return false, errors.ErrorInvalidParam("任务名重复")
	}

	return uc.taskRepo.Update(ctx, id, ca)
}

// Delete 删除
func (uc *TaskUsecase) Delete(ctx context.Context, id uint64) error {
	return uc.taskRepo.Delete(ctx, id)
}

// Find 查询
func (uc *TaskUsecase) Find(ctx context.Context, id uint64, option *TaskOption) (*Task, error) {
	return uc.taskRepo.Find(ctx, id, option)
}

// List 批量查询
func (uc *TaskUsecase) List(ctx context.Context, filter *TaskListFilter, option *TaskOption) ([]*Task, int, error) {
	return uc.taskRepo.List(ctx, filter, option)
}

// ListByCameraId 根据摄像机id查询绑定的任务
func (uc *TaskUsecase) ListByCameraId(ctx context.Context, cameraId uint64, option *TaskOption) ([]*Task, error) {
	return uc.taskRepo.ListByCameraId(ctx, cameraId, option)
}

// CountByTaskStatus 根据任务状态查询任务数量
func (uc *TaskUsecase) CountByTaskStatus(ctx context.Context) (map[TaskStatus]uint64, error) {
	return uc.taskRepo.CountByTaskStatus(ctx)
}

// Start 启动任务
func (uc *TaskUsecase) Start(ctx context.Context, id uint64) error {
	return uc.StartWithStatus(ctx, id, TaskStatusRunning)
}

// StartByTask 启动任务
func (uc *TaskUsecase) StartByTask(ctx context.Context, ta *Task) error {
	if ta == nil {
		uc.log.Errorf("[StartByTask] task is nil")
		return errors.ErrorInvalidParam("task is nil")
	}
	id := ta.Id

	// 退出时更新任务状态
	newStatus := TaskStatusRunning
	newReason := "任务已启动"
	defer func() {
		if err := uc.taskRepo.UpdateStatusAndReason(ctx, id, newStatus, newReason); err != nil {
			uc.log.Errorf("update task status err: %v", err)
		}
	}()

	// 修改状态为启动中
	if err := uc.taskRepo.UpdateStatus(ctx, id, TaskStatusStarting); err != nil {
		newStatus = TaskStatusFailed
		newReason = "任务启动失败"
		return err
	}

	// 区分模式 分别处理
	schedulerRepo := uc.selector(uc.getMode(ta))
	if err := schedulerRepo.Start(ctx, ta); err != nil {
		newStatus = TaskStatusFailed
		newReason = "任务启动失败"
		return err
	}
	return nil
}

// StartWithStatus 启动任务
func (uc *TaskUsecase) StartWithStatus(ctx context.Context, id uint64, taskStatus TaskStatus) error {
	ta, err := uc.Find(ctx, id, &TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
		PreloadDevice:     true,
	})
	if err != nil {
		return err
	}

	// 退出时更新任务状态
	newStatus := TaskStatusRunning
	newReason := "任务已启动"
	defer func() {
		if err := uc.taskRepo.UpdateStatusAndReason(ctx, id, newStatus, newReason); err != nil {
			uc.log.Errorf("update task status err: %v", err)
		}
	}()

	// 修改状态为启动中
	if err := uc.taskRepo.UpdateStatus(ctx, id, TaskStatusStarting); err != nil {
		newStatus = TaskStatusFailed
		newReason = "任务启动失败"
		return err
	}

	// 区分模式 分别处理
	schedulerRepo := uc.selector(uc.getMode(ta))
	if err := schedulerRepo.Start(ctx, ta); err != nil {
		newStatus = TaskStatusFailed
		newReason = "任务启动失败"
		return err
	}
	return nil
}

// Stop 停止任务
func (uc *TaskUsecase) Stop(ctx context.Context, id uint64) error {
	return uc.StopWithStatus(ctx, id, TaskStatusStopped)
}

// StopWithStatus 停止任务
func (uc *TaskUsecase) StopWithStatus(ctx context.Context, id uint64, taskStatus TaskStatus) error {
	ta, err := uc.Find(ctx, id, &TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
	})
	if err != nil {
		return err
	}

	// 退出时更新任务状态
	newStatus := TaskStatusStopped
	newReason := "任务已停止"
	defer func() {
		if err := uc.taskRepo.UpdateStatusAndReason(ctx, id, newStatus, newReason); err != nil {
			uc.log.Errorf("update task status err: %v", err)
		}
	}()

	// 修改状态为停止中
	if err := uc.taskRepo.UpdateStatus(ctx, id, TaskStatusStopping); err != nil {
		return err
	}

	// 区分模式 分别处理
	schedulerRepo := uc.selector(uc.getMode(ta))
	if err := schedulerRepo.Stop(ctx, ta); err != nil {
		return err
	}
	return nil
}

// GetStatuses 获取任务状态
func (uc *TaskUsecase) GetStatuses(ctx context.Context, taskIdList []uint) ([]TaskIdStatus, error) {
	// 其他服务
	taskIdStatusList := make([]TaskIdStatus, 0)

	filterIdIn := make([]uint64, 0)
	for _, taskId := range taskIdList {
		filterIdIn = append(filterIdIn, uint64(taskId))
	}
	filter := &TaskListFilter{
		IdIn: filterIdIn,
	}
	taskOption := &TaskOption{
		PreloadDevice: true,
	}
	tasks, _, err := uc.taskRepo.List(ctx, filter, taskOption)
	if err != nil {
		return taskIdStatusList, err
	}
	var ServerTaskIdList []uint
	for _, task := range tasks {
		if task.Device.Type == ServerDevice && task.Type != TaskTypeAudio {
			// 放入k8s服务器列表
			ServerTaskIdList = append(ServerTaskIdList, uint(task.Id))
			continue
		}
		taskIdStatusList = append(taskIdStatusList, TaskIdStatus{
			TaskId: uint(task.Id),
			Status: task.Status,
			Reason: task.Reason,
		})
	}
	if len(ServerTaskIdList) > 0 {
		k8sScheduler := uc.selector("k8s")
		taskStatus, err := k8sScheduler.GetStatuses(ctx, ServerTaskIdList)
		if err != nil {
			uc.log.Errorf("查询pod状态失败 error: ", err)
		} else {
			taskIdStatusList = append(taskIdStatusList, taskStatus...)
		}
	}
	return taskIdStatusList, nil
}

// UpdateStatus 更新任务状态
func (uc *TaskUsecase) UpdateStatus(ctx context.Context, id uint64, status TaskStatus) error {
	return uc.taskRepo.UpdateStatus(ctx, id, status)
}

// CountTaskCameraByStatus 根据状态统计任务摄像头数量
func (uc *TaskUsecase) CountTaskCameraByStatus(ctx context.Context, status []TaskStatus) (map[TaskStatus]uint64, error) {
	return uc.taskRepo.CountTaskCameraByStatus(ctx, status)
}

// FindTaskCameraById 获取任务摄像头
func (uc *TaskUsecase) FindTaskCameraById(ctx context.Context, taskId uint64, cameraId uint64) (*TaskCamera, error) {
	return uc.taskRepo.FindTaskCameraById(ctx, taskId, cameraId)
}

// CountAlgoInUse 查询使用中的算法数量
func (uc *TaskUsecase) CountAlgoInUse(ctx context.Context) (int, error) {
	return uc.taskRepo.CountAlgoInUse(ctx)
}

// CheckExists 检查字段值是否重复
func (uc *TaskUsecase) CheckExists(ctx context.Context, req *CheckTaskExistsRequest) (bool, error) {
	return uc.taskRepo.CheckExists(ctx, req)
}

// CheckAndUpdateTaskStatus 查询所有摄像机的在线状态并更新数据库
func (uc *TaskUsecase) CheckAndUpdateTaskStatus(ctx context.Context) error {
	start := time.Now()
	arr, total, err := uc.taskRepo.List(ctx, nil, nil)
	if err != nil {
		uc.log.Infof("CheckAndUpdateTaskStatus task List, taskList: %+v, err: %+v", arr, err)
		return err
	}

	taskIds := make([]uint, 0)
	for _, task := range arr {
		taskIds = append(taskIds, uint(task.Id))
	}
	k8sScheduler := uc.selector("k8s")
	taskIdStatusList, err := k8sScheduler.GetStatuses(ctx, taskIds)
	uc.log.Infof("CheckAndUpdateTaskStatus k8s GetStatuses, taskIdStatusList: %+v, err: %+v", taskIdStatusList, err)
	if err != nil {
		return err
	}
	taskIdStatusMap := make(map[uint]TaskStatus)
	for _, taskIdStatus := range taskIdStatusList {
		taskIdStatusMap[taskIdStatus.TaskId] = taskIdStatus.Status
	}

	var wg sync.WaitGroup
	for _, task := range arr {
		status, ok := taskIdStatusMap[uint(task.Id)]
		if !ok || task.Status == status {
			continue
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, taskId uint64, status TaskStatus) {
			defer wg.Done()
			if err := uc.taskRepo.UpdateStatus(ctx, taskId, status); err != nil {
				uc.log.Errorf("update task status err: %v", err)
			}
		}(&wg, task.Id, status)
	}
	wg.Wait()
	uc.log.Infof("CheckAndUpdateTaskStatus done, total: %d, cost: %v", total, time.Since(start))
	return nil
}

// StartVSS 启动任务
func (uc *TaskUsecase) StartVSS(ctx context.Context, tasks []*Task) error {
	return uc.schedulerVSSRepo.StartVSS(ctx, tasks)
}

// StopVSS 停止任务
func (uc *TaskUsecase) StopVSS(ctx context.Context, id uint64) error {
	ta, err := uc.Find(ctx, id, &TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
	})
	if err != nil {
		return err
	}
	return uc.schedulerVSSRepo.StopVSS(ctx, ta)
}

// UpdateEcxVSSStatus 提供给ecx使用
func (uc *TaskUsecase) UpdateEcxVSSStatus(ctx context.Context, params UpdateTaskStatus) error {
	parentId := params.MessageData.Content.TaskID
	ecxTaskStatus := params.MessageData.Content.TaskStatus

	// 检查当前任务所在组的所有任务，如果启动中，停止中，失败的任务，就不再启动
	filter := TaskListFilter{
		ParentId: parentId,
	}
	allTasks, _, err := uc.taskRepo.List(ctx, &filter, &TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
		PreloadDevice:     true,
	})
	if err != nil {
		uc.log.Errorf("UpdateEcxVSSStatus 获取组对应的任务列表失败: %v", err)
		return err
	}
	if len(allTasks) == 0 {
		uc.log.Errorf("UpdateEcxVSSStatus 获取组对应的任务列表为空")
		return nil
	}

	if allTasks[0].Device.State == DeviceStateOffline {
		uc.log.Error("关联设备下线")
		return nil
	}

	notStoppedTasks := make([]*Task, 0)
	runningTasks := make([]*Task, 0)
	startingTasks := make([]*Task, 0)
	stoppingTasks := make([]*Task, 0)
	failedTasks := make([]*Task, 0)

	for _, task := range allTasks {
		if task.Status != TaskStatusStopped {
			notStoppedTasks = append(notStoppedTasks, task)
		}
		switch task.Status {
		case TaskStatusStarting:
			startingTasks = append(startingTasks, task)
		case TaskStatusStopping:
			stoppingTasks = append(stoppingTasks, task)
		case TaskStatusRunning:
			runningTasks = append(runningTasks, task)
		case TaskStatusFailed:
			failedTasks = append(failedTasks, task)
		}
	}

	// 已经标记为失败的任务，不再接受ecx更新状态
	if len(failedTasks) > 0 {
		return nil
	}

	switch ecxTaskStatus {
	case "creating", "cancel":
		return nil
	case "offline":
		var wg sync.WaitGroup
		for _, task := range notStoppedTasks {
			wg.Add(1)
			go func(wg *sync.WaitGroup, taskId uint64, status TaskStatus) {
				defer wg.Done()
				if err := uc.taskRepo.UpdateStatus(ctx, taskId, status); err != nil {
					uc.log.Errorf("update task status err: %v", err)
				}
			}(&wg, task.Id, TaskStatusFailed)
		}
		wg.Wait()
		return nil
	case "running":
		var wg sync.WaitGroup
		for _, task := range startingTasks {
			wg.Add(1)
			go func(wg *sync.WaitGroup, taskId uint64, status TaskStatus) {
				defer wg.Done()
				if err := uc.taskRepo.UpdateStatus(ctx, taskId, status); err != nil {
					uc.log.Errorf("update task status err: %v", err)
				}
			}(&wg, task.Id, TaskStatusRunning)
		}
		wg.Wait()
		return nil
	case "online":
		times, err := uc.taskRepo.GetTaskStartTimes(ctx, parentId)
		if err != nil {
			uc.log.Errorf("GetTaskStartTimes err: %v", err)
		}
		if times == 1 {
			return nil
		}
		if len(stoppingTasks) > 0 {
			var wg sync.WaitGroup
			for _, task := range stoppingTasks {
				wg.Add(1)
				go func(wg *sync.WaitGroup, taskId uint64, status TaskStatus) {
					defer wg.Done()
					if err := uc.taskRepo.UpdateStatus(ctx, taskId, status); err != nil {
						uc.log.Errorf("update task status err: %v", err)
					}
				}(&wg, task.Id, TaskStatusStopped)
			}
			wg.Wait()
			err = uc.StartVSS(ctx, runningTasks)
			if err != nil {
				uc.log.Errorf("UpdateEcxVSSStatus Start tasks err: %+v", err)
				return err
			}
		}
		if len(startingTasks) > 0 {
			tmp := make([]*Task, 0)
			tmp = append(tmp, runningTasks...)
			tmp = append(tmp, startingTasks...)
			err = uc.StartVSS(ctx, tmp)
			if err != nil {
				uc.log.Errorf("UpdateEcxVSSStatus Start tasks err: %+v", err)
				return err
			}
		}
		if len(failedTasks) > 0 {
			var wg sync.WaitGroup
			for _, task := range notStoppedTasks {
				wg.Add(1)
				go func(wg *sync.WaitGroup, taskId uint64, status TaskStatus) {
					defer wg.Done()
					if err := uc.taskRepo.UpdateStatus(ctx, taskId, status); err != nil {
						uc.log.Errorf("update task status err: %v", err)
					}
				}(&wg, task.Id, TaskStatusStopped)
			}
			wg.Wait()
			err = uc.StopVSS(ctx, failedTasks[0].Id)
			if err != nil {
				uc.log.Errorf("UpdateEcxVSSStatus Stop tasks err: %+v", err)
				return err
			}
		}
	}
	return nil
}

// GetLogVSS 查询任务日志
func (uc *TaskUsecase) GetLogVSS(ctx context.Context, task *Task, conn *websocket.Conn) error {
	return uc.schedulerVSSRepo.GetLogVSS(ctx, task, conn)
}

// GetLog 查询任务日志
func (uc *TaskUsecase) GetLog(ctx context.Context, task *Task, conn *websocket.Conn) error {
	mode := uc.getMode(task)
	schedulerRepo := uc.selector(mode)
	return schedulerRepo.GetLog(ctx, task, conn)
}

// UpdateEcxStatus 提供给ecx使用
func (uc *TaskUsecase) UpdateEcxStatus(ctx context.Context, params UpdateTaskStatus) error {
	mode := "ecx"

	uc.log.Debugf("ecx input is : %+v", params)

	taskIDStr := params.MessageData.Content.TaskID
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		uc.log.Errorf("更新的任务id不是数字: %v", err)
		return err
	}

	ta, err := uc.Find(ctx, uint64(taskID), &TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
		PreloadDevice:     true,
	})
	if err != nil {
		return err
	}

	if ta.Device.State == "offline" {
		uc.log.Error("关联设备下线")
		return nil
	}

	// 已经标记为未运行和失败的任务，不再接受ecx更新状态
	if ta.Status == TaskStatusStopped || ta.Status == TaskStatusFailed {
		if params.MessageData.Content.TaskStatus != "running" {
			return nil
		}
	}

	var taskStatus TaskStatus
	switch params.MessageData.Content.TaskStatus {
	case "online":
		if ta.Status == TaskStatusStarting || ta.Status == TaskStatusRunning { // 如果是从creating状态变成online状态，说明是停止了
			taskStatus = TaskStatusStopped
		} else {
			taskStatus = TaskStatusInitializing
		}
	case "offline":
		taskStatus = TaskStatusFailed
	case "running":
		taskStatus = TaskStatusRunning
	case "creating":
		taskStatus = TaskStatusStarting
	case "stopped":
		taskStatus = TaskStatusStopped
	case "cancel":
		return nil
	default:
		uc.log.Error("更新的任务状态不在已知的枚举值中")
		return err
	}

	if taskStatus == TaskStatusFailed {
		// 如果是在5秒之中就失败,增加一次重试,应对盒子重启的情况
		if time.Now().Unix()-ta.LastStartTime.Unix() < 5 {
			schedulerRepo := uc.selector(mode)
			err = schedulerRepo.Stop(ctx, ta)
			if err != nil {
				uc.log.Errorf("停止ecx任务失败：%v", err)
				return err
			}
			time.Sleep(5 * time.Second)
			err = schedulerRepo.Start(ctx, ta)
			if err != nil {
				uc.log.Errorf("启动ecx任务失败：%v", err)
				return err
			}
			return nil
		}
	}

	return uc.taskRepo.UpdateStatus(ctx, uint64(taskID), taskStatus)
}

// GetReason 获取任务状态变更原因
func (uc *TaskUsecase) GetReason(ctx context.Context, id uint64) (string, error) {
	return uc.taskRepo.GetReason(ctx, id)
}

// UpdateReason 更新任务状态变更原因
func (uc *TaskUsecase) UpdateReason(ctx context.Context, id uint64, reason string) error {
	return uc.taskRepo.UpdateReason(ctx, id, reason)
}

// UpdateStatusAndReason 更新任务状态和原因
func (uc *TaskUsecase) UpdateStatusAndReason(ctx context.Context, id uint64, status TaskStatus, reason string) error {
	return uc.taskRepo.UpdateStatusAndReason(ctx, id, status, reason)
}

// 是否在string 数组中
func IsInArrayString(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

// ParseMode 解析算法模式
func (uc *TaskUsecase) ParseMode(algo *Algorithm) string {
	if algo == nil {
		return ""
	}
	if algo.Platform == "node" {
		return "k8s"
	} else {
		// 默认走api模式 如果是docker的算法 返回docker
		provider := algo.Provider
		// apiProvider := []string{"sophgo_park", "sophgo_city"}
		dockerProvider := []string{"ctyun_ias"}
		if IsInArrayString(provider, dockerProvider) {
			return "docker"
		}
		return "api"
	}
}

func (uc *TaskUsecase) getMode(ta *Task) string {
	var mode string
	if ta != nil && ta.Algo != nil && ta.Algo.Algorithm != nil {
		mode = uc.ParseMode(ta.Algo.Algorithm)
	} else {
		mode = ""
	}
	return mode
}

type TaskEventType string

const (
	TaskUpdate TaskEventType = "task_update"
)

// TaskEvent 任务事件
type TaskEvent struct {
	Type TaskEventType
	Task Task
	Meta map[string]any
}

func (uc *TaskUsecase) Events(ctx context.Context, size int) <-chan TaskEvent {
	return uc.taskRepo.Events(ctx, size)
}

// TaskCameraRemoveDuplicates 摄像机去重
func TaskCameraRemoveDuplicates(taskCamera, subTaskCamera []TaskCamera) []TaskCamera {
	cameraIdMap := make(map[uint64]struct{})
	for _, camera := range taskCamera {
		cameraIdMap[camera.Id] = struct{}{}
	}
	for _, camera := range subTaskCamera {
		if _, ok := cameraIdMap[camera.Id]; !ok {
			taskCamera = append(taskCamera, camera)
			cameraIdMap[camera.Id] = struct{}{}
		}
	}
	return taskCamera
}

// 解析算法动态参数-布控周期
func ParseControlTime(algoConfig string) (controlPeriod string, err error) {
	config := make(map[string]interface{})
	if err := json.Unmarshal([]byte(algoConfig), &config); err != nil {
		errMsg := fmt.Errorf("解析算法配置失败: %v", err)
		return "", errMsg
	}
	if _, ok := config["control_period"]; !ok {
		return "", nil
	} else {
		controlPeriod = config["control_period"].(string)
	}
	return controlPeriod, err
}

// 解析算法动态参数-告警间隔
func ParsePeriod(algoConfig string) (period uint32, err error) {
	var pVal, pUnitVal float64 // 周期值, 周期单位值
	var pUnit string           // 周期单位

	config := make(map[string]interface{})
	if err := json.Unmarshal([]byte(algoConfig), &config); err != nil {
		return 0, fmt.Errorf("JSON解析失败: err: %v, algoConfig: %v", err, algoConfig)
	}
	// 获取周期值和周期单位
	if _, ok := config["periodVal"]; !ok {
		return 0, fmt.Errorf("algoConfig中没有periodVal字段: algoConfig: %v", algoConfig)
	} else {
		pVal, ok = config["periodVal"].(float64)
		if !ok {
			return 0, fmt.Errorf("periodVal字段不支持转换为float64: algoConfig: %v", algoConfig)
		}
	}
	if _, ok := config["periodValUnit"]; !ok {
		return 0, fmt.Errorf("algoConfig中没有periodValUnit字段: algoConfig: %v", algoConfig)
	} else {
		pUnit, ok = config["periodValUnit"].(string)
		if !ok {
			return 0, fmt.Errorf("periodValUnit字段不支持转换为string: algoConfig: %v", algoConfig)
		}
	}

	// 转换周期单位
	switch pUnit {
	case "s":
		pUnitVal = 1.0
	case "m":
		pUnitVal = 60.0
	case "h":
		pUnitVal = 3600.0
	default:
		return 0, fmt.Errorf("周期单位不支持: %v", pUnit)
	}

	period = uint32(pVal * pUnitVal)
	return period, nil
}
