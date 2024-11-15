package data

import (
	"context"

	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/devicecamera"
)

// 汇总摄像机ID列表
func collectCameraIDs(cameras []*biz.Camera) []uint64 {
	var ids []uint64
	for _, ca := range cameras {
		ids = append(ids, ca.Id)
	}
	return ids
}

// 关联设备与摄像头(多对多)
func (r *deviceRepo) bindCamerasToDevice(ctx context.Context, deviceID uint64, cameras []*biz.Camera) error {
	cameraIDs := collectCameraIDs(cameras)
	if len(cameraIDs) == 0 {
		return nil
	}

	var deviceCameras []*ent.DeviceCameraCreate
	for _, ca := range cameras {
		deviceCamera := r.data.db.DeviceCamera(ctx).Create()
		deviceCamera.SetCameraID(ca.Id).
			SetDeviceID(deviceID)
		deviceCameras = append(deviceCameras, deviceCamera)
	}

	return r.data.db.DeviceCamera(ctx).CreateBulk(deviceCameras...).Exec(ctx)
}

// 解除关联设备与摄像头(多对多)
func (r *deviceRepo) unbindCamerasFromDevice(ctx context.Context, deviceID uint64) error {
	_, err := r.data.db.DeviceCamera(ctx).Delete().Where(devicecamera.DeviceID(deviceID)).Exec(ctx)
	return err
}
