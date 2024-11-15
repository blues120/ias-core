package data

import (
	"context"

	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

type upPlatformRepo struct {
	data *Data

	log *log.Helper
}

func (r *upPlatformRepo) Get(ctx context.Context, defaultGbId string) (*biz.UpPlatform, error) {
	platform, err := r.getOrCreateDefaultUpPlatform(ctx, defaultGbId)
	if err != nil {
		return nil, err
	}
	return ConvertToBiz(platform), nil
}

func (r *upPlatformRepo) Reset(ctx context.Context, defaultGbId string) error {
	_, err := r.getOrCreateDefaultUpPlatform(ctx, defaultGbId)
	return err
}

func (r *upPlatformRepo) Update(ctx context.Context, platform *biz.UpPlatform) error {
	// 查找第一条记录
	upPlatform, err := r.data.db.UpPlatform(ctx).Query().First(context.Background())
	if err != nil {
		return err // 如果出现错误，如没有找到记录，则返回错误
	}

	// 使用biz结构更新找到的记录
	_, err = r.data.db.UpPlatform(ctx).
		UpdateOne(upPlatform).
		SetSipID(platform.SipID).
		SetSipDomain(platform.SipDomain).
		SetSipIP(platform.SipIP).
		SetSipPort(platform.SipPort).
		SetSipUser(platform.SipUser).
		SetSipPassword(platform.SipPassword).
		SetDescription(platform.Description).
		SetHeartbeatInterval(platform.HeartbeatInterval).
		SetRegisterInterval(platform.RegisterInterval).
		SetTransType(platform.TransType).
		SetGBID(platform.GbID).
		SetCascadestatus(platform.CascadeStatus).
		SetRegistrationStatus(platform.RegistrationStatus). //这里单独Update 默认修改后都是offline 等待重新更新状态
		Save(ctx)

	return err
}

func (r *upPlatformRepo) UpdateRegistrationStatus(ctx context.Context, registrationStatus string) error {
	// 查找第一条记录
	upPlatform, err := r.data.db.UpPlatform(ctx).Query().First(context.Background())
	if err != nil {
		return err // 如果出现错误，如没有找到记录，则返回错误
	}

	// 使用biz结构更新找到的记录
	_, err = r.data.db.UpPlatform(ctx).
		UpdateOne(upPlatform).
		SetRegistrationStatus(registrationStatus). //这里单独Update
		Save(ctx)

	return err
}

func ConvertToBiz(upPlatform *ent.UpPlatform) *biz.UpPlatform {
	return &biz.UpPlatform{
		SipID:              upPlatform.SipID,
		SipDomain:          upPlatform.SipDomain,
		SipIP:              upPlatform.SipIP,
		SipPort:            upPlatform.SipPort,
		SipUser:            upPlatform.SipUser,
		SipPassword:        upPlatform.SipPassword,
		Description:        upPlatform.Description,
		HeartbeatInterval:  upPlatform.HeartbeatInterval,
		RegisterInterval:   upPlatform.RegisterInterval,
		TransType:          upPlatform.TransType,
		GbID:               upPlatform.GBID,
		CascadeStatus:      upPlatform.Cascadestatus,
		RegistrationStatus: upPlatform.RegistrationStatus,
	}
}

func NewUpPlatformRepo(data *Data, logger log.Logger) biz.UpPlatformRepo {
	return &upPlatformRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//func (r *upPlatformRepo) resetOrCreateDefaultUpPlatform(ctx context.Context) (*ent.UpPlatform, error) {
//	// 查询第一条记录
//	upPlatform, err := r.data.db.UpPlatform(ctx).Query().First(ctx)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, err
//	}
//
//	// 更新或创建默认记录
//	if upPlatform != nil {
//		// 更新现有记录为默认值
//		_, err = r.data.db.UpPlatform(ctx).UpdateOne(upPlatform).
//			SetSipID("34020000002000000001").
//			SetSipDomain("3402000000").
//			SetSipIP("").
//			SetSipPort(5060).
//			SetGBID("34020000002000000001").
//			SetSipUser("").
//			SetSipPassword("").
//			SetRegisterInterval(3600).
//			SetHeartbeatInterval(60).
//			SetTransType("UDP").
//			SetDescription("").
//			SetRegistrationStatus("offline").
//			SetCascadestatus("disable").
//			Save(ctx)
//	} else {
//		// 创建默认记录
//		upPlatform, err = r.data.db.UpPlatform(ctx).Create().
//			SetSipID("34020000002000000001").
//			SetSipDomain("3402000000").
//			SetSipIP("").
//			SetSipPort(5060).
//			SetGBID("34020000002000000001").
//			SetSipUser("").
//			SetSipPassword("").
//			SetRegisterInterval(3600).
//			SetHeartbeatInterval(60).
//			SetTransType("UDP").
//			SetDescription("").
//			SetRegistrationStatus("offline").
//			SetCascadestatus("disable").
//			Save(ctx)
//	}
//
//	return upPlatform, err
//}

func (r *upPlatformRepo) getOrCreateDefaultUpPlatform(ctx context.Context, defaultGbId string) (*ent.UpPlatform, error) {
	// 查询第一条记录
	upPlatform, err := r.data.db.UpPlatform(ctx).Query().First(ctx)
	if err == nil {
		return upPlatform, nil // 返回找到的记录
	}

	if !ent.IsNotFound(err) {
		return nil, err // 返回查询过程中的其他错误
	}

	// 不存在记录，创建默认记录
	upPlatform, err = r.data.db.UpPlatform(ctx).Create().
		SetSipID("34020000002000000001").
		SetSipDomain("3402000000").
		SetSipIP("192.168.1.66").
		SetSipPort(5060).
		SetGBID(defaultGbId).
		SetSipUser("admin").
		SetSipPassword("").
		SetRegisterInterval(3600).
		SetHeartbeatInterval(60).
		SetTransType("UDP").
		SetDescription("").
		SetRegistrationStatus("offline").
		SetCascadestatus("disable").
		Save(ctx)

	return upPlatform, err // 返回创建的默认记录或错误
}