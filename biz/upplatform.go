package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/access"
)

type UpPlatform struct {
	SipID              string
	SipDomain          string
	SipIP              string
	SipPort            int32
	SipUser            string
	SipPassword        string
	Description        string
	HeartbeatInterval  int32
	RegisterInterval   int32
	TransType          string
	GbID               string
	CascadeStatus      string
	RegistrationStatus string
}

type UpPlatformRepo interface {
	Get(ctx context.Context, defaultGbId string) (*UpPlatform, error)
	Reset(ctx context.Context, defaultGbId string) error
	Update(ctx context.Context, platform *UpPlatform) error
	UpdateRegistrationStatus(ctx context.Context, registrationStatus string) error
}

type UpPlatformUseCase struct {
	upPlatformRepo UpPlatformRepo
	data           *conf.Data

	log *log.Helper
}

func NewUpPlatformUseCase(repo UpPlatformRepo, data *conf.Data, logger log.Logger) *UpPlatformUseCase {
	return &UpPlatformUseCase{upPlatformRepo: repo, data: data, log: log.NewHelper(logger)}
}

func (uc *UpPlatformUseCase) Get(ctx context.Context) (*UpPlatform, error) {
	return uc.upPlatformRepo.Get(ctx, uc.data.Gb28181.LocalGbid)
}

func (uc *UpPlatformUseCase) Reset(ctx context.Context) error {
	return uc.upPlatformRepo.Reset(ctx, uc.data.Gb28181.LocalGbid)
}

func (uc *UpPlatformUseCase) Update(ctx context.Context, platform *UpPlatform) error {
	addr := uc.data.Gb28181.SipAddr
	requestId := access.GetUuid()
	if platform.CascadeStatus == "enable" {
		// 调用信令服务注册
		url := fmt.Sprintf("http://%s%s", addr, access.CascadeRegister)
		//开始调用信令服务器的/gb28181/v1/register/start接口，启动级联
		startUpPlatformParam := &access.StartUpPlatformParam{
			ID:             "1", // 平台ID 多个上级平台时使用
			Name:           platform.SipID,
			CascadeRegion:  platform.SipID,
			LocalId:        platform.GbID,
			ServerId:       platform.SipID,
			ServerRealm:    platform.SipDomain,
			ServerIp:       platform.SipIP,
			ServerPort:     strconv.Itoa(int(platform.SipPort)),
			UserName:       platform.SipUser,
			UserPwd:        platform.SipPassword,
			Expires:        strconv.Itoa(int(platform.RegisterInterval)),
			KeepLive:       strconv.Itoa(int(platform.HeartbeatInterval)),
			Protocol:       platform.TransType,
			CharacterType:  "UTF-8",
			PermissionSet:  "PTZ,RTCP,RECORD",
			CascadeNetWork: "public",
			CascadeType:    1,
		}
		_, err := access.SendStartUpPlatformRequest(platform.SipID, url, startUpPlatformParam, requestId)
		if err != nil {
			return err
		}
	} else {
		// 调用信令服务停止注册
		url := fmt.Sprintf("http://%s%s", addr, access.CascadeUnRegister)
		stopUpPlatformParam := &access.StopUpPlatformParam{
			ID:       strconv.Itoa(int(access.GetRandId())),
			ServerId: platform.SipID,
		}
		_, err := access.SendStopUpPlatformRequest(platform.SipID, url, stopUpPlatformParam, requestId)
		if err != nil {
			return err
		}
	}
	return uc.upPlatformRepo.Update(ctx, platform)
}

func (uc *UpPlatformUseCase) GetLocalGbId(ctx context.Context) (string, error) {
	return uc.data.Gb28181.LocalGbid, nil
}

func (uc *UpPlatformUseCase) UpdateStatus(ctx context.Context, registrationStatus string) error {
	return uc.upPlatformRepo.UpdateRegistrationStatus(ctx, registrationStatus)
}

func (uc *UpPlatformUseCase) CheckUpPlatformStatus(ctx context.Context) error {
	upPlatform, _ := uc.upPlatformRepo.Get(ctx, uc.data.Gb28181.LocalGbid)
	type reqParam struct {
		ID       string
		ServerId string
	}
	// ResponseStruct 定义来自服务器响应的结构体
	type ResponseStruct struct {
		ID             string `json:"ID"`
		ServerId       string `json:"ServerId"`
		StartStatus    string `json:"StartStatus"`
		RegisterStatus string `json:"RegisterStatus"`
		ErrorCode      int    `json:"ErrorCode"`
		ErrorMessage   string `json:"ErrorMessage"`
		RequestId      string `json:"RequestId"`
	}
	param := &reqParam{ID: "1", ServerId: upPlatform.SipID}
	addr := uc.data.Gb28181.SipAddr
	url := fmt.Sprintf("http://%s%s", addr, access.CascadeStatus)

	client := &http.Client{}
	paramBytes, _ := json.Marshal(param)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(paramBytes))
	if err != nil {
		fmt.Println("Error creating CheckUpPlatformStatus request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making CheckUpPlatformStatus request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading CheckUpPlatformStatus response:", err)
		return err
	}

	var responseStruct ResponseStruct
	if err := json.Unmarshal(body, &responseStruct); err != nil {
		fmt.Println("Error unmarshalling CheckUpPlatformStatus response:", err)
		return err
	}

	var status string
	if responseStruct.RegisterStatus == "off" {
		status = "offline"
	} else {
		status = "online"
	}
	return uc.upPlatformRepo.UpdateRegistrationStatus(ctx, status)
}

func (uc *UpPlatformUseCase) CheckCascadeStatus(ctx context.Context) error {
	platform, err := uc.upPlatformRepo.Get(ctx, uc.data.Gb28181.LocalGbid)
	if err != nil {
		return err
	}
	if platform.CascadeStatus == "enable" && platform.RegistrationStatus == "offline" {
		// 如果国标服务开启 且为在线 需要定时任务去开启 可能是刚重启
		// 调用信令服务注册
		addr := uc.data.Gb28181.SipAddr
		requestId := access.GetUuid()
		url := fmt.Sprintf("http://%s%s", addr, access.CascadeRegister)
		//开始调用信令服务器的/gb28181/v1/register/start接口，启动级联
		startUpPlatformParam := &access.StartUpPlatformParam{
			ID:             "1", // 平台ID 多个上级平台时使用
			Name:           platform.SipID,
			CascadeRegion:  platform.SipID,
			LocalId:        platform.GbID,
			ServerId:       platform.SipID,
			ServerRealm:    platform.SipDomain,
			ServerIp:       platform.SipIP,
			ServerPort:     strconv.Itoa(int(platform.SipPort)),
			UserName:       platform.SipUser,
			UserPwd:        platform.SipPassword,
			Expires:        strconv.Itoa(int(platform.RegisterInterval)),
			KeepLive:       strconv.Itoa(int(platform.HeartbeatInterval)),
			Protocol:       platform.TransType,
			CharacterType:  "UTF-8",
			PermissionSet:  "PTZ,RTCP,RECORD",
			CascadeNetWork: "public",
			CascadeType:    1,
		}
		_, err = access.SendStartUpPlatformRequest(platform.SipID, url, startUpPlatformParam, requestId)
		if err != nil {
			return err
		}
	}
	return nil
}
