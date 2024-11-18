package biz

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/beevik/etree"
	wsdiscovery "github.com/use-go/onvif/ws-discovery"
	"github.com/blues120/ias-core/biz/streaming"
	"github.com/blues120/ias-core/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/access"
	"github.com/blues120/ias-core/pkg/pagination"
)

// CameraStatus 摄像机状态
type CameraStatus string

const (
	CameraStatusOnline  CameraStatus = "online"
	CameraStatusOffline CameraStatus = "offline"
)

// Values provides list valid values for Enum.
func (CameraStatus) Values() []string {
	return []string{string(CameraStatusOnline), string(CameraStatusOffline)}
}

// MediaType 多媒体设备类型
type MediaType string

const (
	MediaTypeCamera MediaType = "camera"
	MediaTypeAudio  MediaType = "audio"
)

// Values provides list valid values for Enum.
func (MediaType) Values() []string {
	return []string{string(MediaTypeCamera), string(MediaTypeAudio)}
}

// 多媒体设备类型对应的中文名称
type MediaTypeCnName string

const (
	MediaTypeCnCamera MediaTypeCnName = "视频设备"
	MediaTypeCnAudio  MediaTypeCnName = "音频设备"
)

// Values provides list valid values for Enum.
func (MediaTypeCnName) Values() []string {
	return []string{string(MediaTypeCnCamera), string(MediaTypeCnAudio)}
}

type CameraDevice struct {
	DeviceId uint64
	CameraId uint64
}

// Camera 摄像机
type Camera struct {
	Id                uint64    // id
	Name              string    // 名称
	Position          string    // 点位
	Region            string    // 区域
	RegionStr         string    // 区域展示
	Longitude         float64   // 经度
	Latitude          float64   // 纬度
	CustomNumber      int       // 自定义编号
	ChannelId         string    // 通道id
	SerialNumber      string    // 设备序列号
	PoleNumber        string    // 杆号
	DeviceDescription string    // 设备描述
	Scene             string    // 适用场景
	Place             string    // 所属场所
	Type              MediaType // 多媒体设备类型
	TenantID          string    // 租户id
	AccessOrgList     string    // 授权的组织 id
	IsStatusOffline   bool      // 状态是否离线 摄像机离线提示用

	// ** 国标添加字段
	TransType    string // 传输协议UDP/TCP
	DeviceIP     string // IP地址
	DevicePort   int32  // 端口号
	GbId         string // 国标ID
	SipUser      string // 国标信令SIP认证用户名
	SipPassword  string // 国标信令SIP认证密码
	ChannelGbId  string // 国标通道ID
	GBDeviceType string // 国标设备类型 ipc nvr platform
	// **

	CreatedAt         time.Time          // 创建时间
	UpdatedAt         time.Time          // 更新时间
	Status            CameraStatus       // 状态
	StreamingProtocol streaming.Protocol // 流媒体协议
	StreamingInfo     *streaming.Info    // 流媒体信息
	Task              []*Task            // 关联任务
	CameraDevice      []*CameraDevice    // CameraDevice关联表数据
	Device            []*Device          // 关联设备
}

// Refresh 刷新摄像机信息
func (ca *Camera) Refresh(ctx context.Context) {
	updateFuncs := []func(context.Context){
		ca.refreshStatus,
		ca.refreshStreamingInfo,
	}
	var wg sync.WaitGroup
	for _, updateFunc := range updateFuncs {
		wg.Add(1)
		go func(fn func(context.Context)) {
			defer wg.Done()
			fn(ctx)
		}(updateFunc)
	}
	wg.Wait()
}

// refreshStatus 刷新摄像机状态
func (ca *Camera) refreshStatus(ctx context.Context) {
	ca.Status = CameraStatusOffline
	if isOnline, _ := ca.StreamingProtocol.IsOnline(ctx, time.Second*20); isOnline {
		ca.Status = CameraStatusOnline
	}
}

// refreshStreamingInfo 刷新摄像机流媒体信息
func (ca *Camera) refreshStreamingInfo(ctx context.Context) {
	info, _ := ca.StreamingProtocol.GetInfo(ctx)
	ca.StreamingInfo = info
}

type CameraRepo interface {
	// Save 创建摄像机
	Save(ctx context.Context, ca *Camera) (uint64, error)

	// Update 更新摄像机
	Update(ctx context.Context, id uint64, ca *Camera) error

	// UpdateStatus 更新摄像机状态
	UpdateStatus(ctx context.Context, id uint64, status CameraStatus) error

	// UpdateStreamInfo 更新摄像机流信息
	UpdateStreamInfo(ctx context.Context, id uint64, info *streaming.Info) error

	// Delete 删除摄像机
	Delete(ctx context.Context, id uint64) error

	// BatchDelete 批量删除摄像机
	BatchDelete(ctx context.Context, ids []uint64) (int, error)

	// Find 查询摄像机
	Find(ctx context.Context, id uint64, option *CameraOption) (*Camera, error)

	// Exist 摄像机是否存在
	Exist(ctx context.Context, field *CameraExistField, option *CameraOption) (*Camera, error)

	// List 查询摄像机列表
	List(ctx context.Context, filter *CameraListFilter, option *CameraOption) ([]*Camera, int, error)

	// Import 导入摄像机 异步执行
	Import(ctx context.Context, importId string, importer CameraImporter, cas interface{}) error

	// QueryImportProgress 查询导入进度
	QueryImportProgress(ctx context.Context, importId string) (*ImportProgress, error)

	// CountRegion 统计每个区域包含的摄像机数量
	CountRegion(ctx context.Context) ([]*CameraRegionCount, error)

	// CountByBindTask 统计绑定了任务的摄像机数量
	CountByBindTask(ctx context.Context, ids []uint64) (int, error)

	// CountByStatus 统计每种状态包含的摄像机数量
	CountByStatus(ctx context.Context) ([]*CameraStatusCount, error)

	// GetChannelList 获取本机国标通道列表
	GetChannelList(ctx context.Context) ([]*GBChannel, error)

	// BatchUpdateChannel 批量更新国标通道向上级联号
	BatchUpdateChannel(ctx context.Context, channels []*GBChannel) error
}

type CameraUsecase struct {
	cameraRepo CameraRepo
	data       *conf.Data

	log *log.Helper
}

func NewCameraUsecase(repo CameraRepo, data *conf.Data, logger log.Logger) *CameraUsecase {
	return &CameraUsecase{cameraRepo: repo, data: data, log: log.NewHelper(logger)}
}

// Create 创建摄像机
func (uc *CameraUsecase) Create(ctx context.Context, ca *Camera) (uint64, error) {
	if ca.StreamingProtocol == nil {
		return 0, errors.ErrorEmptyStreamingProtocol("摄像机流媒体协议为空")
	}
	return uc.cameraRepo.Save(ctx, ca)
}

// Update 更新摄像机
func (uc *CameraUsecase) Update(ctx context.Context, id uint64, ca *Camera) error {
	if ca.StreamingProtocol == nil {
		return errors.ErrorEmptyStreamingProtocol("摄像机流媒体协议为空")
	}
	return uc.cameraRepo.Update(ctx, id, ca)
}

// UpdateStatus 更新摄像机状态
func (uc *CameraUsecase) UpdateStatus(ctx context.Context, id uint64, status CameraStatus) error {
	return uc.cameraRepo.UpdateStatus(ctx, id, status)
}

// UpdateStreamInfo 更新摄像机流信息
func (uc *CameraUsecase) UpdateStreamInfo(ctx context.Context, id uint64, info *streaming.Info) error {
	return uc.cameraRepo.UpdateStreamInfo(ctx, id, info)
}

// Delete 删除摄像机
func (uc *CameraUsecase) Delete(ctx context.Context, id uint64) error {
	_, err := uc.cameraRepo.Find(ctx, id, nil)
	if err != nil {
		return err
	}
	err = uc.cameraRepo.Delete(ctx, id)
	if err == nil {
		_ = uc.StopGbPreview(ctx, id)
	}
	return err
}

// BatchDelete 批量删除摄像机
func (uc *CameraUsecase) BatchDelete(ctx context.Context, ids []uint64) (int, error) {
	return uc.cameraRepo.BatchDelete(ctx, ids)
}

// CameraOption 查询选项
type CameraOption struct {
	PreloadTask         bool // 加载关联的任务数据
	PreloadDeviceCamera bool // 加载关联的deviceCamera数据
	PreloadDevice       bool // 加载关联的device数据
}

// Find 查询摄像机
func (uc *CameraUsecase) Find(ctx context.Context, id uint64, option *CameraOption) (*Camera, error) {
	return uc.cameraRepo.Find(ctx, id, option)
}

// CameraExistField .
type CameraExistField struct {
	Name              string             // 名称
	CustomNumber      int                // 自定义编号
	SerialNumber      string             // 设备序列号
	StreamingProtocol streaming.Protocol // 流媒体协议
	ChannelId         string             // 通道ID
	GbId              string             // 国标ID
	GBDeviceType      string             // 国标设备类型
	ChannelGbId       string             // 通道国标ID
}

// Exist 摄像机是否存在
// field 中任意一个字段存在，则摄像机存在
// 不存在时 err == errors.ErrorCameraNotFound
func (uc *CameraUsecase) Exist(ctx context.Context, field *CameraExistField, option *CameraOption) (*Camera, error) {
	return uc.cameraRepo.Exist(ctx, field, option)
}

// CameraListFilter 批量查询过滤条件
type CameraListFilter struct {
	/*
		模糊查询条件
	*/
	NameContain     string   // 名称
	PositionContain string   // 点位
	SceneContain    []string // 适用场景
	Places          string   // 模糊查询所属场所 搜索项查询用
	/*
		精确查询条件
	*/
	DeviceID              uint64                 // 设备ID
	Status                CameraStatus           // 状态
	Region                string                 // 区域
	SerialNumber          string                 // 设备序列号
	StreamingProtocolType streaming.ProtocolType // 流媒体协议类型
	Place                 string                 // 精确查询所属场所 查询下属摄像头用
	Type                  MediaType              // 多媒体设备类型
	/*
		范围查询条件
	*/
	AlgorithmIdIn []uint64 // 算法ID列表
	CameraIds     []uint64 // 摄像机ID列表
	/*
		特殊查询条件
	*/
	// CustomNumberIn 按编号范围搜索，例如"1,4,6-12"，表示搜索编号为1,4和6至12的所有摄像机
	// eg:
	// 1 -> [][]int64{{1}}
	// 1-13 -> [][]int64{{1,13}}
	// 1,3,4 -> [][]int64{{1}, {3}, {4}}}
	// 1,1-13 -> [][]int64{{1},{1,13}}
	// 1-12,17-20 -> [][]int64{{1,12},{17,20}}
	CustomNumberIn [][]int
	/*
		分页
	*/
	Pagination *pagination.Pagination

	// 特殊查询项
	IncludeDeleted bool // 查询结果是否包含已删除的摄像机
}

// List 查询摄像机列表
func (uc *CameraUsecase) List(ctx context.Context, filter *CameraListFilter, option *CameraOption) ([]*Camera, int, error) {
	return uc.cameraRepo.List(ctx, filter, option)
}

// CameraStatusCount 摄像机状态统计
type CameraStatusCount struct {
	Status CameraStatus `json:"status"`
	Count  int          `json:"count"`
}

// CountByStatus 统计每种状态包含的摄像机数量
func (uc *CameraUsecase) CountByStatus(ctx context.Context) ([]*CameraStatusCount, error) {
	return uc.cameraRepo.CountByStatus(ctx)
}

type CameraRegionCount struct {
	Region string `json:"region"`
	Count  int    `json:"count"`
}

// CountByRegion 统计每个区域包含的摄像机数量
func (uc *CameraUsecase) CountByRegion(ctx context.Context) ([]*CameraRegionCount, error) {
	return uc.cameraRepo.CountRegion(ctx)
}

// CountByBindTask 统计绑定了任务的摄像机数量
func (uc *CameraUsecase) CountByBindTask(ctx context.Context, ids []uint64) (int, error) {
	return uc.cameraRepo.CountByBindTask(ctx, ids)
}

// CameraImporter 摄像机导入器
type CameraImporter interface {
	// Handler 定义如何处理 Camera，失败时返回 error
	Handler(interface{}, int, uint64) error

	// ErrRecord 定义如何记录 error，可通过 GetRecord 获取记录内容
	ErrRecord(error, int)

	// GetRecord 获取记录内容用于保存
	GetRecord() *bytes.Buffer
}

// Import 导入摄像机 异步执行
func (uc *CameraUsecase) Import(ctx context.Context, importId string, importer CameraImporter, cas interface{}) error {
	return uc.cameraRepo.Import(ctx, importId, importer, cas)
}

type ImportProgress struct {
	Total   int     // 总数
	Success int     // 成功个数
	Fail    int     // 失败个数
	Process float64 // 上传进度
	Url     string  // 下载链接（Process = 1 时返回）
}

// QueryImportProgress 查询导入进度
func (uc *CameraUsecase) QueryImportProgress(ctx context.Context, importId string) (*ImportProgress, error) {
	return uc.cameraRepo.QueryImportProgress(ctx, importId)
}

// CheckAndUpdateCameraStatus 查询所有摄像机的在线状态并更新数据库
func (uc *CameraUsecase) CheckAndUpdateCameraStatus(ctx context.Context, poolSize int) error {
	start := time.Now()
	arr, total, err := uc.cameraRepo.List(ctx, nil, nil)
	if err != nil {
		return err
	}
	if err := uc.UpdateCameraStatus(ctx, arr, poolSize); err != nil {
		return err
	}
	uc.log.Infof("CheckAndUpdateCameraStatus done, total: %d, cost: %v", total, time.Since(start))
	return nil
}

func (uc *CameraUsecase) UpdateCameraStatus(ctx context.Context, arr []*Camera, poolSize int) error {
	var wg sync.WaitGroup
	pool, err := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		defer wg.Done()
		ca := i.(*Camera)
		status := ca.Status
		// 如果流媒体协议的状态发生了变化，则更新数据库
		if ca.refreshStatus(ctx); ca.Status != status {
			if ca.Status == CameraStatusOffline {
				ca.IsStatusOffline = true
			}
			if err := uc.cameraRepo.UpdateStatus(ctx, ca.Id, ca.Status); err != nil {
				uc.log.Errorf("update camera status err: %v", err)
			}
		}
		if ca.StreamingInfo != nil && ca.StreamingInfo.CodecName != "" {
			return
		}
		// 获取流媒体信息
		if ca.refreshStreamingInfo(ctx); ca.StreamingInfo != nil && ca.StreamingInfo.CodecName != "" {
			if err := uc.cameraRepo.UpdateStreamInfo(ctx, ca.Id, ca.StreamingInfo); err != nil {
				uc.log.Errorf("update camera stream info err: %v", err)
			}
		}
	})
	if err != nil {
		return err
	}
	defer pool.Release()

	for _, ca := range arr {
		ca := ca
		wg.Add(1)
		if err := pool.Invoke(ca); err != nil {
			uc.log.Errorf("CheckAndUpdateCameraStatus err: %v", err)
		}
	}
	wg.Wait()
	return nil
}

func (uc *CameraUsecase) SearchOnvifCameras(nic string) ([]OnvifCamera, error) {
	resp := make([]OnvifCamera, 0)

	// Call a ws-discovery Probe Message to Discover NVT type Devices
	devices, err := wsdiscovery.SendProbe(nic, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		return resp, err
	}

	nvtDevicesSeen := make(map[string]bool)

	for _, j := range devices {
		uc.log.Infof(fmt.Sprintf("SearchOnvifCameras devices:::%v", j))
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return resp, err
		}

		var onvifCamera OnvifCamera
		for _, xaddr := range doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs") {
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]

			// Remove port from xaddr
			if index := strings.Index(xaddr, ":"); index != -1 {
				xaddr = xaddr[:index]
			}

			if !nvtDevicesSeen[xaddr] {
				onvifCamera = OnvifCamera{
					IP:   xaddr,
					Port: "554",
				}
			}
		}

		// 解析摄像头厂商
		for _, scopes := range doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes") {
			dns := strings.Split(scopes.Text(), " ")
			var hardware string
			prefix := "onvif://www.onvif.org/hardware/"
			for _, dn := range dns {
				// 判断是否是hardware开头
				if strings.HasPrefix(dn, prefix) {
					// 去掉前缀
					hardware = strings.TrimPrefix(dn, prefix)
					break
				}
			}
			if hardware != "" {
				brand := strings.Split(hardware, "-")[0]
				onvifCamera.Brand = brand
			}
		}
		resp = append(resp, onvifCamera)
	}
	return resp, nil
}

type OnvifCamera struct {
	IP    string
	Port  string
	Brand string
}

// DeviceType alias for int
type DeviceType int

// Onvif Device Tyoe
const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

type GBChannel struct {
	ChannelId   string
	ChannelGbId string
}

func (uc *CameraUsecase) GetChannelList(ctx context.Context) ([]*GBChannel, error) {
	return uc.cameraRepo.GetChannelList(ctx)
}

func (uc *CameraUsecase) BatchUpdateChannel(ctx context.Context, channels []*GBChannel) error {
	return uc.cameraRepo.BatchUpdateChannel(ctx, channels)
}

func (uc *CameraUsecase) StartGbPreview(ctx context.Context, deviceId uint64) (string, error) {
	addr := uc.data.Gb28181.SipAddr
	uc.log.Infof("addr:%v", addr)
	url := fmt.Sprintf("http://%s%s", addr, access.StartPlay)
	ca, err := uc.cameraRepo.Find(ctx, deviceId, nil)
	if err != nil {
		return "", fmt.Errorf("find camera by deviceId error: %v", err)
	}
	requestId := access.GetUuid()

	startLiveParams := access.SendRequestParamsToGBSipJson{
		GbDeviceId:    ca.GbId,
		GbCameraId:    ca.GbId,
		StreamName:    ca.ChannelGbId,
		TransPriority: ca.TransType,
		OutProtocol:   "rtsp",
		EventType:     "streamServer",
		UserId:        strconv.Itoa(int(ca.Id)),
		RequestId:     requestId,
	}
	if ca.GBDeviceType == "nvr" || ca.GBDeviceType == "platform" {
		startLiveParams.GbCameraId = ca.ChannelGbId
	}

	PlayUrls, err := access.SendStartDeviceRequest(deviceId, "PUT", url, startLiveParams, access.RequestSipTimeOut, requestId)
	if err != nil {
		return "", fmt.Errorf("failed to send start device request: %w", err)
	}
	sp, err := streaming.NewGb28181(PlayUrls.RtspUrl)
	if err != nil {
		return "", fmt.Errorf("failed to create new GB28181 streaming protocol: %w", err)
	}
	ca.StreamingProtocol = sp
	// refresh 操作可能会导致 context 超时，新建 ctx 避免影响摄像头创建
	ca.Refresh(context.Background())
	if err = uc.cameraRepo.Update(ctx, ca.Id, ca); err != nil {
		return "", fmt.Errorf("failed to update camera repository: %w", err)
	}
	//_ = uc.cameraRepo.UpdateStreamSource(ctx, deviceId, PlayUrls.RtspUrl)

	return PlayUrls.RtspUrl, nil
}

func (uc *CameraUsecase) StartCascadePlay(ctx context.Context, param *access.StartPlatformPushStreamParam) (*access.StartPlatformPushStreamResp, error) {
	requestId := access.GetUuid()
	addr := uc.data.Gb28181.SipAddr
	startUrl := "http://" + addr + "/gb28181/v1/devices?Action=startCascadePlay"
	return access.SendPushStreamRequest(param.GbDeviceId, startUrl, param, requestId)
}

func (uc *CameraUsecase) StopCascadePlay(ctx context.Context, param *access.StopPlatformPushStreamParam) error {
	requestId := access.GetUuid()
	addr := uc.data.Gb28181.SipAddr
	stopUrl := "http://" + addr + "/gb28181/v1/devices?Action=stopCascadePlay"
	_, err := access.SendStopPushStreamRequest(param.GbDeviceId, stopUrl, param, requestId)
	return err
}

func (uc *CameraUsecase) StartGbNotPUllPreview(ctx context.Context) error {
	addr := uc.data.Gb28181.SipAddr
	uc.log.Infof("addr:%v", addr)
	url := fmt.Sprintf("http://%s%s", addr, access.StartPlay)
	filter := CameraListFilter{
		StreamingProtocolType: streaming.ProtocolTypeGB28181,
		Status:                CameraStatusOffline,
	}
	arr, _, err := uc.List(ctx, &filter, nil)
	if err != nil {
		uc.log.Errorf("failed to list cameras: %s", err)
		return err
	}
	for _, ca := range arr {
		deviceId := ca.Id
		requestId := access.GetUuid()

		startLiveParams := access.SendRequestParamsToGBSipJson{
			GbDeviceId:    ca.GbId,
			GbCameraId:    ca.GbId,
			StreamName:    ca.ChannelGbId,
			TransPriority: ca.TransType,
			OutProtocol:   "rtsp",
			EventType:     "streamServer",
			UserId:        strconv.Itoa(int(ca.Id)),
			RequestId:     requestId,
		}
		if ca.GBDeviceType == "nvr" || ca.GBDeviceType == "platform" {
			startLiveParams.GbCameraId = ca.ChannelGbId
		}

		PlayUrls, err := access.SendStartDeviceRequest(deviceId, "PUT", url, startLiveParams, access.RequestSipTimeOut, requestId)
		if err != nil {
			return fmt.Errorf("failed to send start device request: %w", err)
		}
		sp, err := streaming.NewGb28181(PlayUrls.RtspUrl)
		if err != nil {
			return fmt.Errorf("failed to create new GB28181 streaming protocol: %w", err)
		}
		ca.StreamingProtocol = sp
		// refresh 操作可能会导致 context 超时，新建 ctx 避免影响摄像头创建
		ca.Refresh(context.Background())
		if err = uc.cameraRepo.Update(ctx, ca.Id, ca); err != nil {
			return fmt.Errorf("failed to update camera repository: %w", err)
		}
	}
	return nil
}

func (uc *CameraUsecase) StopGbPreview(ctx context.Context, deviceId uint64) error {
	addr := uc.data.Gb28181.SipAddr
	url := fmt.Sprintf("http://%s%s", addr, access.StopPlay)
	ca, _ := uc.cameraRepo.Find(ctx, deviceId, nil)
	requestId := access.GetUuid()
	if ca == nil {
		return nil
	}

	stopLiveParams := access.SendRequestParamsToGBSipJson{
		GbDeviceId: ca.GbId,
		GbCameraId: ca.GbId,
		StreamName: ca.ChannelGbId,
		RequestId:  requestId,
	}

	err := access.SendStopDeviceRequest(deviceId, "PUT", url, stopLiveParams, access.RequestSipTimeOut, requestId)
	if err != nil {
		return err
	}

	return nil
}
