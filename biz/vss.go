package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/vss"
	standardProto "github.com/blues120/ias-proto/api/standard/v1"
	warningProto "github.com/blues120/ias-proto/api/warning/v1"
	"strconv"
)

type VssUsecase struct {
	deviceRepo    DeviceRepo
	cameraRepo    CameraRepo
	eventSubsRepo EventSubsRepo
	service       *conf.Service
	vssSign       *conf.VssSign
	vssClient     *vss.VssCallbackClient

	log *log.Helper
}

type VssMessage struct {
	PlatformID  string         `json:"platformId"`
	MessageData VssMessageData `json:"messageData"`
}

type VssMessageData struct {
	Base  BaseData              `json:"base"`
	MsgID string                `json:"msg_id"`
	Title string                `json:"title"`
	Event *standardProto.Events `json:"event"`
}

type BaseData struct {
	BoxID       string       `json:"box_id"`
	WorkspaceID string       `json:"workspace_id"`
	ZoneID      string       `json:"zone_id"`
	BoxAttr     BoxAttribute `json:"box_attr"`
	ModelName   string       `json:"model_name"`
	ObjectID    string       `json:"object_id"`
}

type BoxAttribute struct {
	DeviceID  string `json:"device_id"`
	ChannelID string `json:"channel_id"`
}

func NewVssUsecase(deviceRepo DeviceRepo, cameraRepo CameraRepo, eventSubsRepo EventSubsRepo, service *conf.Service, vssSign *conf.VssSign, logger log.Logger) *VssUsecase {
	vssClient := vss.NewVssCallbackClient(vssSign, logger)
	return &VssUsecase{
		deviceRepo:    deviceRepo,
		cameraRepo:    cameraRepo,
		eventSubsRepo: eventSubsRepo,
		service:       service,
		vssSign:       vssSign,
		vssClient:     vssClient,
		log:           log.NewHelper(logger)}
}

// SendWarning2Vss 推送告警消息给数生，适用任何应用
func (s *VssUsecase) SendWarning2Vss(ctx context.Context, pushUrl string, deviceId, cameraId, msgId uint64, title, modelName string, msg *warningProto.Frame, warnMsg *VssMessage) error {
	// 有些VssMessage的属性core不好处理，在应用处理完传入
	if warnMsg == nil {
		// 查询摄像头
		camera, err := s.cameraRepo.Find(ctx, cameraId, nil)
		if err != nil {
			return err
		}
		// 防止event_subs没数据的情况
		if len(camera.ChannelId) != 0 {
			if pushUrl == "" {
				// 查询设备
				var device *Device
				device, err = s.deviceRepo.Find(ctx, deviceId, nil)
				if err != nil {
					return err
				}
				// 查询订阅消息
				var eventSubs *EventSubscription
				eventSubs, err = s.eventSubsRepo.FindByFilter(ctx, &EventSubsFindFilter{
					Status:    EventSubStatusEnable,
					BoxId:     device.ExtId,
					ChannelId: camera.ChannelId,
				})
				if err != nil {
					return err
				}
				if eventSubs == nil {
					return nil
				}
				pushUrl = eventSubs.Callback
				warnMsg = s.ToVssMessage(device, camera, msgId, title, modelName, msg)
			}
		} else {
			s.log.Infof("camera.ChannelId empty,stop pushing vss，deviceId %v, cameraId: %v", deviceId, cameraId)
			return nil
		}
	}
	param, err := json.Marshal(warnMsg)
	if err != nil {
		return err
	}
	res, err := s.vssClient.Post(pushUrl, "aialarmMsgPush", param)
	s.log.Infof("send warn message %s, response: %s", string(param), string(res))
	if err != nil {
		return err
	}
	return nil
}

// ToVssMessage 组装上报结构体
func (s *VssUsecase) ToVssMessage(device *Device, camera *Camera, msgId uint64, title, modelName string, msg *warningProto.Frame) *VssMessage {
	// 发送告警， 组装告警的结果
	return &VssMessage{
		PlatformID: s.vssSign.Id,
		MessageData: VssMessageData{
			Base: BaseData{
				BoxID:       device.ExtId,
				WorkspaceID: device.WorkspaceId,
				ZoneID:      device.ZoneId,
				BoxAttr: BoxAttribute{
					DeviceID:  device.EquipId,
					ChannelID: camera.ChannelId,
				},
				ModelName: modelName,
				ObjectID:  msg.ObjectId,
			},
			MsgID: strconv.FormatUint(msgId, 10),
			Title: title,
			Event: msg.AlgoOutput.Events,
		},
	}
}

// CheckIsVssService 判断服务是否是数生服务
func (s *VssUsecase) CheckIsVssService() bool {
	if s.service != nil {
		return s.service.Name == conf.Service_VSS_SERVICE
	}
	return false
}
