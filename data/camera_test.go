//go:build integration
// +build integration

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/errors"
	"testing"
)

func TestCameraRepo(t *testing.T) {
	data, cleanup := NewTestData(t)
	defer cleanup()

	ctx := context.Background()
	repo := NewCameraRepo(data, nil, log.DefaultLogger)

	// 创建摄像机
	ca := &biz.Camera{
		Name:              "test",
		StreamingProtocol: NewTestStreamingProtocol(),
		Status:            biz.CameraStatusOnline,
	}
	id, err := repo.Save(ctx, ca)
	require.NoError(t, err)
	require.Equal(t, uint64(1), id)

	// 查询摄像机
	ca, err = repo.Find(ctx, id, &biz.CameraOption{})
	require.NoError(t, err)
	require.Equal(t, ca.Id, id)
	require.Equal(t, ca.Status, biz.CameraStatusOnline)

	// 范围查询摄像机
	_, cnt, err := repo.List(ctx, &biz.CameraListFilter{NameContain: ca.Name}, &biz.CameraOption{})
	require.NoError(t, err)
	require.Equal(t, cnt, 1)

	// 删除摄像机
	err = repo.Delete(ctx, id)
	require.NoError(t, err)

	// 查询摄像机是否存在
	_, err = repo.Exist(ctx, &biz.CameraExistField{Name: ca.Name}, &biz.CameraOption{})
	require.True(t, errors.IsCameraNotFound(err))
}
