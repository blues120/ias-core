package biz

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redsync/redsync/v4"
	"github.com/blues120/ias-kit/time_helper"
)

// 告警消息
type WarningInfo struct {
	AppName   string // 告警所属应用名, 如warning
	ID        uint64
	IsWarn    bool
	TaskID    uint64
	CameraID  uint64
	WarnTime  uint64
	AlarmName string
}

// 短信错误记录
type SmsErr struct {
	AppName  string // 应用名称
	RecordID uint64 // 应用内告警ID
	ErrorMsg string // sms服务错误信息
}

type SmsNotifyRepo interface {
	// RecordSmsErr 短信发送错误记录
	RecordSmsErr(ctx context.Context, data *SmsErr) error

	// FilterUnwarnedPhoneNumberList 过滤出近期未发送过警告信息的手机号码
	FilterUnwarnedPhoneNumberList(ctx context.Context, taskID uint64, phoneNumberList []string) []string

	// MarkPhoneAsWarned 标记手机号码已发送警告信息
	MarkPhoneAsWarned(ctx context.Context, taskID uint64, phoneNumber string)

	// EopPost 发送告警短信
	EopPost(body []byte) ([]byte, error)

	// AddLock 申请加锁
	AddLock(taskID uint64) (*redsync.Mutex, error)

	// ReleaseLock 释放锁
	ReleaseLock(*redsync.Mutex) error
}

type SmsNotifyUsecase struct {
	informRepo    InformRepo
	smsNotifyRepo SmsNotifyRepo
	cameraRepo    CameraRepo

	log *log.Helper
}

func NewSmsNotifyUsecase(
	informRepo InformRepo, smsNotifyRepo SmsNotifyRepo, cameraRepo CameraRepo,
	logger log.Logger,
) *SmsNotifyUsecase {
	return &SmsNotifyUsecase{
		informRepo:    informRepo,
		smsNotifyRepo: smsNotifyRepo,
		cameraRepo:    cameraRepo,
		log:           log.NewHelper(log.With(logger, "clazz", "SmsNotifyUsecase")),
	}
}

// Process 处理告警短信发送
func (uc *SmsNotifyUsecase) Process(ctx context.Context, record *WarningInfo) (err error) {
	needed, informData, phoneNumberList := uc.isNotifyNeeded(ctx, record)
	if !needed {
		return
	}

	// 拼接短信内容信息
	templateParams := uc.getTemplateParams(ctx, record)

	// 加锁
	var lk *redsync.Mutex
	lk, err = uc.smsNotifyRepo.AddLock(record.TaskID)
	if err != nil {
		uc.log.Errorf("send-sms lock err: %s", err.Error())
		return
	}
	defer uc.smsNotifyRepo.ReleaseLock(lk)

	// 发送频率检查，过滤得到需要发送报警的手机号码
	phoneNumberListToNotify := uc.smsNotifyRepo.FilterUnwarnedPhoneNumberList(ctx, record.TaskID, phoneNumberList)
	if len(phoneNumberListToNotify) == 0 {
		uc.log.Debugf("pass (phoneNumberListToNotify empty) record: %v task: %v", record.ID, record.TaskID)
		return
	}

	var smsBody []byte
	for _, phoneNum := range phoneNumberListToNotify {
		smsBody, err = json.Marshal(map[string]string{
			"action":        "SendSms",
			"signName":      informData.SignName,
			"phoneNumber":   phoneNum,
			"templateCode":  informData.TemplateCode,
			"templateParam": string(templateParams),
		})
		if err != nil {
			uc.log.Errorf("sms params encode err: %s", err.Error())
			return
		}
		uc.sendSms(ctx, smsBody, record)
		uc.smsNotifyRepo.MarkPhoneAsWarned(ctx, record.TaskID, phoneNum)
	}

	return
}

// 是否需要发送短信
func (uc *SmsNotifyUsecase) isNotifyNeeded(ctx context.Context, record *WarningInfo) (bool, *Inform, []string) {
	if !record.IsWarn { // 非告警
		uc.log.Debugf("pass (IsWarn: %v) record: %v task: %v", record.IsWarn, record.ID, record.TaskID)
		return false, nil, nil
	}
	// 读取报警信息：模板及手机号码等
	informData, err := uc.informRepo.FindByTaskID(ctx, record.TaskID)
	if err != nil {
		uc.log.Errorf("pass (error) get inform err: %v", err)
		return false, nil, nil
	}
	// 未开启通知，终止报警流程
	if informData.NotifySwitch != "on" {
		uc.log.Debugf("pass (NotifySwitch: %v) record: %v task: %v", informData.NotifySwitch, record.ID, record.TaskID)
		return false, nil, nil
	}

	phoneNumbers := make([]string, 0)
	if err = json.Unmarshal([]byte(informData.PhoneNumbers), &phoneNumbers); err != nil {
		uc.log.Errorf("pass (Exception) parse phone numbers err: %v", err)
		return false, informData, nil
	}
	// 通知手机号码为空，终止报警流程
	if len(phoneNumbers) == 0 {
		uc.log.Debugf("pass (phoneNumberList empty) record: %v task: %v", record.ID, record.TaskID)
		return false, informData, phoneNumbers
	}

	return true, informData, phoneNumbers
}

// 获取模板参数
func (uc *SmsNotifyUsecase) getTemplateParams(ctx context.Context, record *WarningInfo) string {
	camera, err := uc.cameraRepo.Find(ctx, record.CameraID, nil)
	if err != nil {
		uc.log.Errorf("get camera err: %v", err)
		return ""
	}
	templateParams, _ := json.Marshal(map[string]string{
		"ctyun_time":       time_helper.TimestampToDateTime(record.WarnTime),
		"ctyun_alert_name": record.AlarmName,
		"ctyun_location":   camera.Position,
	}) // 报警模板参数

	return string(templateParams)
}

// SendSms 发送短信
func (uc *SmsNotifyUsecase) sendSms(ctx context.Context, smsBody []byte, record *WarningInfo) {
	sendRsp, err := uc.smsNotifyRepo.EopPost(smsBody)
	if err != nil {
		_ = uc.smsNotifyRepo.RecordSmsErr(ctx, &SmsErr{
			AppName:  record.AppName,
			RecordID: record.ID,
			ErrorMsg: err.Error(),
		})

		uc.log.Errorf("send sms err: %s", err.Error())
		return
	}
	uc.log.Infof("sms notify response: %s", string(sendRsp))

	// 判断是否发送成功
	parsedSendRsp := make(map[string]string)
	_ = json.Unmarshal(sendRsp, &parsedSendRsp)
	value, exists := parsedSendRsp["Code"]
	if !exists {
		value = parsedSendRsp["code"]
	}
	if value != "OK" {
		if err := uc.smsNotifyRepo.RecordSmsErr(ctx, &SmsErr{
			AppName:  record.AppName,
			RecordID: record.ID,
			ErrorMsg: string(sendRsp),
		}); err != nil {
			uc.log.Warnf("record sms err failed: %s", err.Error())
		}
	}
}
