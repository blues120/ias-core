//go:build integration
// +build integration

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"testing"
	"time"
)

func TestNewTaskRepo(t *testing.T) {
	data, cleanup := NewTestData(t)
	defer cleanup()

	ctx := context.Background()
	taskRepo := NewTaskRepo(data, data.db, log.DefaultLogger)
	cameraRepo := NewCameraRepo(data, nil, log.DefaultLogger)
	algoRepo := NewAlgoRepo(data, log.DefaultLogger)

	// 创建摄像机
	ca := &biz.Camera{
		Name:              "test",
		StreamingProtocol: NewTestStreamingProtocol(),
		Status:            biz.CameraStatusOnline,
	}
	cameraId, err := cameraRepo.Save(ctx, ca)
	require.NoError(t, err)
	require.Equal(t, uint64(1), cameraId)
	ca.Id = cameraId

	// 创建算法
	algo := &biz.Algorithm{
		Name: "test",
		Type: "video",
	}
	algoId, err := algoRepo.Save(ctx, algo)
	require.NoError(t, err)
	require.Equal(t, uint64(1), algoId)
	algo.ID = algoId

	// 创建任务，关联摄像机和算法
	taskId, err := taskRepo.Save(ctx, &biz.Task{
		Name:    "test",
		Type:    "stream",
		Cameras: []biz.TaskCamera{{Camera: ca}},
		Algo:    &biz.TaskAlgo{Id: algoId, Algorithm: algo},
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1), taskId)

	// 查询任务
	taskOption := &biz.TaskOption{
		PreloadAlgorithm:  true,
		PreloadCamera:     true,
		PreloadTaskCamera: true,
		PreloadDevice:     true,
	}
	ta, err := taskRepo.Find(ctx, taskId, taskOption)
	require.NoError(t, err)
	require.Equal(t, ta.Id, taskId)
	require.Equal(t, biz.TaskStatusStopped, ta.Status)

	// 任务关联摄像机正确
	require.Equal(t, cameraId, ta.Cameras[0].Id)

	// 任务关联算法正确
	require.Equal(t, algoId, ta.Algo.Id)

	// 更新任务关联的摄像机和算法
	newCa := ca
	newCameraId, err := cameraRepo.Save(ctx, newCa)
	require.NoError(t, err)
	newCa.Id = newCameraId

	newAlgo := algo
	newAlgoId, err := algoRepo.Save(ctx, newAlgo)
	require.NoError(t, err)
	newAlgo.ID = newAlgoId

	ta.Cameras = []biz.TaskCamera{{Camera: newCa}}
	ta.Algo = &biz.TaskAlgo{Id: newAlgoId, Algorithm: newAlgo}
	ret, err := taskRepo.Update(ctx, taskId, ta)
	require.True(t, ret)
	require.NoError(t, err)

	// 批量查询任务
	tas, cnt, err := taskRepo.List(ctx, &biz.TaskListFilter{}, taskOption)
	require.NoError(t, err)
	require.Equal(t, 1, cnt)

	// 任务关联摄像机正确
	require.Equal(t, newCameraId, tas[0].Cameras[0].Id)

	// 任务关联算法正确
	require.Equal(t, newAlgoId, tas[0].Algo.Id)

	// 根据摄像机查询任务
	tas, err = taskRepo.ListByCameraId(ctx, newCameraId, taskOption)
	require.NoError(t, err)
	require.Equal(t, 1, len(tas))
	require.Equal(t, newAlgoId, tas[0].Algo.Id)
	require.Equal(t, newCameraId, tas[0].Cameras[0].Id)

	// 任务状态更新为运行中
	err = taskRepo.UpdateStatus(ctx, taskId, biz.TaskStatusRunning)
	require.NoError(t, err)

	// 查询任务状态是否正常，最后一次启动时间是否更新
	ta, err = taskRepo.Find(ctx, taskId, &biz.TaskOption{})
	require.NoError(t, err)
	require.Equal(t, ta.Status, biz.TaskStatusRunning)
	require.NotNil(t, ta.LastStartTime)
	require.LessOrEqual(t, *ta.LastStartTime, time.Now())

	// 删除任务
	err = taskRepo.Delete(ctx, taskId)
	require.NoError(t, err)

	// 查询任务列表，任务不存在
	_, cnt, err = taskRepo.List(ctx, &biz.TaskListFilter{}, &biz.TaskOption{})
	require.NoError(t, err)
	require.Equal(t, 0, cnt)
}

func TestTaskRepoEvents(t *testing.T) {
	data, cleanup := NewTestData(t)
	defer cleanup()

	ctx := context.Background()
	taskRepo := NewTaskRepo(data, data.db, log.DefaultLogger)

	// 监听任务事件
	events := taskRepo.Events(ctx, 10)

	// 创建任务
	ta := &biz.Task{
		Name: "test",
		Type: "stream",
	}
	taskId, err := taskRepo.Save(ctx, ta)
	require.NoError(t, err)
	ta.Id = taskId

	// 更新任务
	ta.Name = "test2"
	ret, err := taskRepo.Update(ctx, taskId, ta)
	require.NoError(t, err)
	require.True(t, ret)

	// 接收任务更新事件
	timeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	select {
	case evt := <-events:
		require.Equal(t, biz.TaskUpdate, evt.Type)
		require.Equal(t, taskId, evt.Task.Id)
		require.Equal(t, ta.Name, evt.Task.Name)
	case <-timeout.Done():
		require.Fail(t, "task update event not received")
	}
}
