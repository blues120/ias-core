// Code generated by MockGen. DO NOT EDIT.
// Source: biz/camera.go
//
// Generated by this command:
//
//	mockgen -source=biz/camera.go -destination=biz/mock/camera_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	bytes "bytes"
	context "context"
	reflect "reflect"

	biz "gitlab.ctyuncdn.cn/ias/ias-core/biz"
	streaming "gitlab.ctyuncdn.cn/ias/ias-core/biz/streaming"
	gomock "go.uber.org/mock/gomock"
)

// MockCameraRepo is a mock of CameraRepo interface.
type MockCameraRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCameraRepoMockRecorder
}

// MockCameraRepoMockRecorder is the mock recorder for MockCameraRepo.
type MockCameraRepoMockRecorder struct {
	mock *MockCameraRepo
}

// NewMockCameraRepo creates a new mock instance.
func NewMockCameraRepo(ctrl *gomock.Controller) *MockCameraRepo {
	mock := &MockCameraRepo{ctrl: ctrl}
	mock.recorder = &MockCameraRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCameraRepo) EXPECT() *MockCameraRepoMockRecorder {
	return m.recorder
}

// BatchDelete mocks base method.
func (m *MockCameraRepo) BatchDelete(ctx context.Context, ids []uint64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchDelete", ctx, ids)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchDelete indicates an expected call of BatchDelete.
func (mr *MockCameraRepoMockRecorder) BatchDelete(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchDelete", reflect.TypeOf((*MockCameraRepo)(nil).BatchDelete), ctx, ids)
}

// BatchUpdateChannel mocks base method.
func (m *MockCameraRepo) BatchUpdateChannel(ctx context.Context, channels []*biz.GBChannel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateChannel", ctx, channels)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchUpdateChannel indicates an expected call of BatchUpdateChannel.
func (mr *MockCameraRepoMockRecorder) BatchUpdateChannel(ctx, channels any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateChannel", reflect.TypeOf((*MockCameraRepo)(nil).BatchUpdateChannel), ctx, channels)
}

// CountByBindTask mocks base method.
func (m *MockCameraRepo) CountByBindTask(ctx context.Context, ids []uint64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByBindTask", ctx, ids)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByBindTask indicates an expected call of CountByBindTask.
func (mr *MockCameraRepoMockRecorder) CountByBindTask(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByBindTask", reflect.TypeOf((*MockCameraRepo)(nil).CountByBindTask), ctx, ids)
}

// CountByStatus mocks base method.
func (m *MockCameraRepo) CountByStatus(ctx context.Context) ([]*biz.CameraStatusCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByStatus", ctx)
	ret0, _ := ret[0].([]*biz.CameraStatusCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByStatus indicates an expected call of CountByStatus.
func (mr *MockCameraRepoMockRecorder) CountByStatus(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByStatus", reflect.TypeOf((*MockCameraRepo)(nil).CountByStatus), ctx)
}

// CountRegion mocks base method.
func (m *MockCameraRepo) CountRegion(ctx context.Context) ([]*biz.CameraRegionCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountRegion", ctx)
	ret0, _ := ret[0].([]*biz.CameraRegionCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountRegion indicates an expected call of CountRegion.
func (mr *MockCameraRepoMockRecorder) CountRegion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountRegion", reflect.TypeOf((*MockCameraRepo)(nil).CountRegion), ctx)
}

// Delete mocks base method.
func (m *MockCameraRepo) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCameraRepoMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCameraRepo)(nil).Delete), ctx, id)
}

// Exist mocks base method.
func (m *MockCameraRepo) Exist(ctx context.Context, field *biz.CameraExistField, option *biz.CameraOption) (*biz.Camera, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", ctx, field, option)
	ret0, _ := ret[0].(*biz.Camera)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exist indicates an expected call of Exist.
func (mr *MockCameraRepoMockRecorder) Exist(ctx, field, option any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockCameraRepo)(nil).Exist), ctx, field, option)
}

// Find mocks base method.
func (m *MockCameraRepo) Find(ctx context.Context, id uint64, option *biz.CameraOption) (*biz.Camera, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id, option)
	ret0, _ := ret[0].(*biz.Camera)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCameraRepoMockRecorder) Find(ctx, id, option any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCameraRepo)(nil).Find), ctx, id, option)
}

// GetChannelList mocks base method.
func (m *MockCameraRepo) GetChannelList(ctx context.Context) ([]*biz.GBChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelList", ctx)
	ret0, _ := ret[0].([]*biz.GBChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelList indicates an expected call of GetChannelList.
func (mr *MockCameraRepoMockRecorder) GetChannelList(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelList", reflect.TypeOf((*MockCameraRepo)(nil).GetChannelList), ctx)
}

// Import mocks base method.
func (m *MockCameraRepo) Import(ctx context.Context, importId string, importer biz.CameraImporter, cas any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, importId, importer, cas)
	ret0, _ := ret[0].(error)
	return ret0
}

// Import indicates an expected call of Import.
func (mr *MockCameraRepoMockRecorder) Import(ctx, importId, importer, cas any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockCameraRepo)(nil).Import), ctx, importId, importer, cas)
}

// List mocks base method.
func (m *MockCameraRepo) List(ctx context.Context, filter *biz.CameraListFilter, option *biz.CameraOption) ([]*biz.Camera, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter, option)
	ret0, _ := ret[0].([]*biz.Camera)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockCameraRepoMockRecorder) List(ctx, filter, option any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCameraRepo)(nil).List), ctx, filter, option)
}

// QueryImportProgress mocks base method.
func (m *MockCameraRepo) QueryImportProgress(ctx context.Context, importId string) (*biz.ImportProgress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryImportProgress", ctx, importId)
	ret0, _ := ret[0].(*biz.ImportProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryImportProgress indicates an expected call of QueryImportProgress.
func (mr *MockCameraRepoMockRecorder) QueryImportProgress(ctx, importId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryImportProgress", reflect.TypeOf((*MockCameraRepo)(nil).QueryImportProgress), ctx, importId)
}

// Save mocks base method.
func (m *MockCameraRepo) Save(ctx context.Context, ca *biz.Camera) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, ca)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCameraRepoMockRecorder) Save(ctx, ca any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCameraRepo)(nil).Save), ctx, ca)
}

// Update mocks base method.
func (m *MockCameraRepo) Update(ctx context.Context, id uint64, ca *biz.Camera) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, ca)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCameraRepoMockRecorder) Update(ctx, id, ca any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCameraRepo)(nil).Update), ctx, id, ca)
}

// UpdateStatus mocks base method.
func (m *MockCameraRepo) UpdateStatus(ctx context.Context, id uint64, status biz.CameraStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockCameraRepoMockRecorder) UpdateStatus(ctx, id, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockCameraRepo)(nil).UpdateStatus), ctx, id, status)
}

// UpdateStreamInfo mocks base method.
func (m *MockCameraRepo) UpdateStreamInfo(ctx context.Context, id uint64, info *streaming.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStreamInfo", ctx, id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStreamInfo indicates an expected call of UpdateStreamInfo.
func (mr *MockCameraRepoMockRecorder) UpdateStreamInfo(ctx, id, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStreamInfo", reflect.TypeOf((*MockCameraRepo)(nil).UpdateStreamInfo), ctx, id, info)
}

// MockCameraImporter is a mock of CameraImporter interface.
type MockCameraImporter struct {
	ctrl     *gomock.Controller
	recorder *MockCameraImporterMockRecorder
}

// MockCameraImporterMockRecorder is the mock recorder for MockCameraImporter.
type MockCameraImporterMockRecorder struct {
	mock *MockCameraImporter
}

// NewMockCameraImporter creates a new mock instance.
func NewMockCameraImporter(ctrl *gomock.Controller) *MockCameraImporter {
	mock := &MockCameraImporter{ctrl: ctrl}
	mock.recorder = &MockCameraImporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCameraImporter) EXPECT() *MockCameraImporterMockRecorder {
	return m.recorder
}

// ErrRecord mocks base method.
func (m *MockCameraImporter) ErrRecord(arg0 error, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ErrRecord", arg0, arg1)
}

// ErrRecord indicates an expected call of ErrRecord.
func (mr *MockCameraImporterMockRecorder) ErrRecord(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrRecord", reflect.TypeOf((*MockCameraImporter)(nil).ErrRecord), arg0, arg1)
}

// GetRecord mocks base method.
func (m *MockCameraImporter) GetRecord() *bytes.Buffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord")
	ret0, _ := ret[0].(*bytes.Buffer)
	return ret0
}

// GetRecord indicates an expected call of GetRecord.
func (mr *MockCameraImporterMockRecorder) GetRecord() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockCameraImporter)(nil).GetRecord))
}

// Handler mocks base method.
func (m *MockCameraImporter) Handler(arg0 any, arg1 int, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handler indicates an expected call of Handler.
func (mr *MockCameraImporterMockRecorder) Handler(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockCameraImporter)(nil).Handler), arg0, arg1, arg2)
}