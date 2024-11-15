package access

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type StartGbCameraLiveStreamParam struct {
	PeerIp     string `json:"PeerIp"`
	PeerPort   int64  `json:"PeerPort"`
	Ssrc       int64  `json:"Ssrc"`
	TransType  string `json:"TransType"`
	TcpMode    string `json:"TcpMode"`
	AppName    string `json:"AppName"`
	StreamName string `json:"StreamName"`
	PlatformId string `json:"PlatformId"`
	InProtocol string `json:"InProtocol"`
}

type SendRequestParamsToGBSipJson struct {
	GbDeviceId    string
	GbCameraId    string
	GbAudioId     string
	StreamName    string
	StreamType    string
	TransPriority string
	OutProtocol   string
	UserId        string
	PeerIp        string
	PeerPort      string
	Ssrc          string
	TransType     string
	AudioKey      string
	EventType     string
	StartTime     string
	EndTime       string
	RequestId     string
}

type StartPlatformPushStreamParam struct {
	PlatformId     string `json:"PlatformId"`
	GbDeviceId     string `json:"GbDeviceId"`
	GbCameraId     string `json:"GbCameraId"`
	UserId         string `json:"UserId"`
	StreamName     string `json:"StreamName"`
	StreamType     string `json:"StreamType"`
	PeerIp         string `json:"PeerIp"`
	PeerPort       int64  `json:"PeerPort"`
	Ssrc           string `json:"Ssrc"`
	TransType      string `json:"TransType"`
	TransPriority  string `json:"TransPriority"`
	TcpMode        string `json:"TcpMode"`
	StartTime      int64  `json:"StartTime"`
	EndTime        int64  `json:"EndTime"`
	Speed          int64  `json:"Speed"`
	CascadeNetWork string `json:"CascadeNetWork"`
	RequestId      string `json:"RequestId"`
}

type StartPlatformPushStreamResp struct {
	RequestId      string `json:"RequestId"`
	ErrorCode      int64  `json:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage"`
	LocalIp        string `json:"LocalIp"`
	LocalPort      int64  `json:"LocalPort"`
	Ssrc           int64  `json:"Ssrc"`
	StreamName     string `json:"StreamName"`
	CascadeNetWork string `json:"CascadeNetWork"`
}

type StopPlatformPushStreamParam struct {
	PlatformId string `json:"PlatformId"`
	GbDeviceId string `json:"GbDeviceId"`
	GbCameraId string `json:"GbCameraId"`
	StreamName string `json:"StreamName"`
	PeerIp     string `json:"PeerIp"`
	PeerPort   int64  `json:"PeerPort"`
	Ssrc       string `json:"Ssrc"`
	TcpMode    string `json:"TcpMode"`
	RequestId  string `json:"RequestId"`
}

type PlayUrlJson struct {
	RequestId string   `json:"RequestId"`
	PlayUrl   PlayUrls `json:"PlayUrl"`
}

type PlayUrls struct {
	RtmpUrl string `json:"RtmpUrl"`
	FlvUrl  string `json:"FlvUrl"`
	HlsUrl  string `json:"HlsUrl"`
	RtcUrl  string `json:"RtcUrl"`
	RtspUrl string `json:"RtspUrl"`
}

type PlayUrlErr struct {
	ErrorCode    int64  `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

// 都是ipc設備
func dealError(deviceId uint64, errorCode int64) error {
	log.Printf("设备[%v]出现错误", deviceId)
	if errorCode == 10004 {
		log.Printf("更新设备[%v]状态为离线状态", deviceId)
		//_, err := db.Client.Exec("update vss_device set device_status = ?  where device_id  =?  ", "off", deviceId)
		//if err != nil {
		//	log.Println("更新数据库出错", err)
		//	return err
		//}
	}
	return nil
}

func GetUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func SendStartDeviceRequest(deviceId uint64, method, url string, body interface{}, timeout time.Duration, requestId string, ctx ...context.Context) (PlayUrls, error) {
	respBody, err := SendReqWithTimeout(method, url, body, timeout, requestId, ctx...)
	if err != nil {
		return PlayUrls{}, errors.New("信令服务返回错误" + err.Error())
	}
	errorJson := &PlayUrlErr{}
	json.Unmarshal(respBody, &errorJson)
	errorJsonMessage := errorJson.ErrorMessage
	if errorJsonMessage != "" {
		err = dealError(deviceId, errorJson.ErrorCode)
		if err != nil {
			return PlayUrls{}, err
		}
		return PlayUrls{}, errors.New("信令服务返回错误" + errorJsonMessage)
	}
	resJson := &PlayUrlJson{}
	json.Unmarshal(respBody, &resJson)
	return PlayUrls{
		RtmpUrl: resJson.PlayUrl.RtmpUrl,
		HlsUrl:  resJson.PlayUrl.HlsUrl,
		FlvUrl:  resJson.PlayUrl.FlvUrl,
		RtspUrl: resJson.PlayUrl.RtspUrl,
	}, nil
}

func SendStopDeviceRequest(deviceId uint64, method, url string, body interface{}, timeout time.Duration, requestId string, ctx ...context.Context) error {
	log.Printf("开始向%v发送%v请求：%v[%v]", url, method, body, requestId)
	var (
		bodyReader *bytes.Reader
		httpReq    *http.Request
		err        error
	)
	if body != nil {
		bodyB, err := json.Marshal(body)
		if err != nil {
			log.Printf("序列化body失败%v[%v]", err, requestId)
			return errors.New("序列化参数失败" + requestId)
		}
		bodyReader = bytes.NewReader(bodyB)
	}
	cCtx := context.Background()
	if len(ctx) > 0 {
		cCtx = ctx[0]
	}
	nCtx, cancel := context.WithTimeout(cCtx, timeout)
	defer cancel()
	if body != nil {
		httpReq, err = http.NewRequestWithContext(nCtx, method, url, bodyReader)
	} else {
		httpReq, err = http.NewRequestWithContext(nCtx, method, url, nil)
	}
	if err != nil {
		log.Printf("创建http请求%v失败%v[%v]", url, err, requestId)
		return errors.New("创建请求失败" + requestId)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Printf("发送http请求%v失败%v[%v]", url, err, requestId)
		return errors.New("发送请求失败" + requestId)
	}
	log.Printf("发送http请求%v返回：%v, %v[%v]", url, *resp, err, requestId)
	if resp.StatusCode != http.StatusOK {
		return errors.New("请求服务失败" + requestId)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("信令服务器返回...................", string(respBody))

	resErr := &PlayUrlErr{}
	json.Unmarshal(respBody, &resErr)
	resErrMessage := resErr.ErrorMessage
	if resErrMessage != "success" {
		if err != nil {
			return err
		}
		err = dealError(uint64(deviceId), resErr.ErrorCode)
		if err != nil {
			return err
		}
		resErrMessage = "信令服务返回错误" + resErrMessage
		return errors.New(resErrMessage + requestId)
	}
	errorJson := &PlayUrlErr{}
	json.Unmarshal(respBody, &errorJson)
	errorJsonMessage := errorJson.ErrorMessage
	if errorJsonMessage != "" {
		err = dealError(deviceId, errorJson.ErrorCode)
		if err != nil {
			return err
		}
		errorJsonMessage = "信令服务返回错误" + errorJsonMessage
		return errors.New(errorJsonMessage + requestId)
	}
	return nil
}

func SendReqWithTimeout(method, url string, body interface{}, timeout time.Duration, requestId string, ctx ...context.Context) ([]byte, error) {
	log.Printf("开始向%v发送%v请求：%v[%v]", url, method, body, requestId)

	cCtx := context.Background()
	if len(ctx) > 0 {
		cCtx = ctx[0]
	}
	nCtx, cancel := context.WithTimeout(cCtx, timeout)
	defer cancel()

	var bodyReader io.Reader
	if body != nil {
		bodyB, err := json.Marshal(body)
		if err != nil {
			log.Printf("序列化body失败%v[%v]", err, requestId)
			return nil, errors.New("序列化参数失败")
		}
		bodyReader = bytes.NewReader(bodyB)
	}

	httpReq, err := http.NewRequestWithContext(nCtx, method, url, bodyReader)
	if err != nil {
		log.Printf("创建http请求%v失败%v[%v]", url, err, requestId)
		return nil, errors.New("创建请求失败")
	}
	if method != "GET" {
		httpReq.Header.Add("Content-Type", "application/json")
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Printf("发送http请求%v失败%v[%v]", url, err, requestId)
		return nil, errors.New("发送请求失败")
	}
	defer httpRes.Body.Close()

	httpB, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		log.Printf("读取服务的返回失败%v[%v]", err, requestId)
		return nil, errors.New("请求服务失败")
	}

	log.Printf("向%v发送%v请求成功：%v[%v]", url, method, body, requestId)
	return httpB, nil
}

func IsIp(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		// 没有匹配上
		fmt.Println("\n不是ip是域名：" + ip)
		return false
	} else {
		// 匹配上
		fmt.Println("\n是ip：" + address.String())
		return true
	}
}

func GetSipServerPort(boxId uint64) (serverIp string, serverPort int, err error) {
	// 单机版单板写死本机的信令服务器IP和端口 云边版本需要维护盒子地址
	const localServerIP = "127.0.0.1"
	const localServerPort = 8088
	return localServerIP, localServerPort, nil
}

func GetRandId() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63() // 这将生成一个随机的 int64 类型的数字
}

type CommonResp struct {
	RequestId    string `json:"RequestId"`
	ErrorCode    int64  `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

type StartUpPlatformParam struct {
	ID             string `json:"ID"`
	Name           string `json:"Name"`
	CascadeRegion  string `json:"CascadeRegion"`
	LocalId        string `json:"LocalId"`
	ServerId       string `json:"ServerId"`
	ServerRealm    string `json:"ServerRealm"`
	ServerIp       string `json:"ServerIp"`
	ServerPort     string `json:"ServerPort"`
	CascadeIp      string `json:"CascadeIp"`
	CascadePort    string `json:"CascadePort"`
	UserName       string `json:"UserName"`
	UserPwd        string `json:"UserPwd"`
	Expires        string `json:"Expires"`
	KeepLive       string `json:"KeepLive"`
	Protocol       string `json:"Protocol"`
	CharacterType  string `json:"CharacterType"`
	PermissionSet  string `json:"PermissionSet"`
	CascadeNetWork string `json:"CascadeNetWork"`
	CascadeType    int32  `json:"CascadeType"`
}

type StopUpPlatformParam struct {
	ID       string `json:"ID"`
	ServerId string `json:"ServerId"`
}

// 向信令服务器发送启动上级平台级联
func SendStartUpPlatformRequest(gbId string, url string, params *StartUpPlatformParam, requestId string) (*CommonResp, error) {
	body, err := SendReqWithTimeout("PUT", url, params, 10*time.Second, requestId)
	log.Printf("信令返回：%v, %v", string(body), requestId)
	if err != nil {
		return nil, errors.New("访问信令异常" + requestId)
	}
	log.Println("信令服务器返回: ", string(body))
	resp := &CommonResp{}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, errors.New("信令返回异常" + requestId)
	}
	if resp.ErrorCode == 10003 {
		log.Printf("信令提示请求参数不正确,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, errors.New("请求信令参数异常" + requestId)
	}
	if resp.ErrorCode == 10010 {
		log.Printf("up platform has exist,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, nil
	}
	if resp.ErrorCode == 10001 {
		log.Printf("信令提示处理超时,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, errors.New("信令处理超时" + requestId)
	}
	if resp.ErrorCode == 0 {
		log.Printf("start up platform success,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, nil
	}
	log.Printf("start up platform fail,platform_id:[%v]--requestId:[%v]", gbId, requestId)
	return resp, errors.New(resp.ErrorMessage + requestId)
}

// 向信令服务器发送停止上级平台级联
func SendStopUpPlatformRequest(gbId string, url string, params *StopUpPlatformParam, requestId string) (*CommonResp, error) {
	body, err := SendReqWithTimeout("PUT", url, params, 10*time.Second, requestId)
	log.Printf("信令返回：%v, %v", string(body), requestId)
	if err != nil {
		return nil, errors.New("访问信令异常" + requestId)
	}
	log.Println("信令服务器返回: ", string(body))
	resp := &CommonResp{}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, errors.New("信令返回异常" + requestId)
	}
	if resp.ErrorCode == 10011 {
		log.Printf("up platform has not exist,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		//dao.UpdatePlatformStatusById(deviceId, 0, requestId, "off")
		//return StopUpPlatformJson{}, errs.NotFoundErr("up platform has not exist", requestId)
		return resp, nil
	}
	if resp.ErrorCode == 0 {
		log.Printf("stop up platform success,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, nil
	}
	if resp.ErrorCode == 10003 {
		log.Printf("信令提示请求参数不正确,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, errors.New("请求信令参数异常" + requestId)
	}
	if resp.ErrorCode == 10001 {
		log.Printf("信令提示处理超时,platform_id:[%v]--requestId:[%v]", gbId, requestId)
		return resp, errors.New("信令处理超时" + requestId)
	}
	log.Printf("stop up platform fail,platform_id:[%v]--requestId:[%v]", gbId, requestId)
	return resp, errors.New(resp.ErrorMessage + requestId)
}

func SendPushStreamRequest(deviceId string, url string, params *StartPlatformPushStreamParam, requestId string) (*StartPlatformPushStreamResp, error) {
	body, err := SendReqWithTimeout("PUT", url, params, 10*time.Second, requestId)
	log.Printf("信令返回：%v, %v", string(body), requestId)
	if err != nil {
		return nil, errors.New("访问信令异常" + requestId)
	}
	log.Println("信令服务器返回: ", string(body))
	resp := &StartPlatformPushStreamResp{}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, errors.New("信令返回异常" + requestId)
	}
	if resp.ErrorCode == 0 {
		log.Printf("信令提示推流成功,device_id:[%v]--requestId:[%v]", deviceId, requestId)
		return resp, nil
	}
	if resp.ErrorCode == 10003 {
		log.Printf("信令提示请求参数不正确, %v", requestId)
		return resp, errors.New("请求信令参数异常" + requestId)
	}
	if resp.ErrorCode == 10001 {
		log.Printf("信令处理出现异常，err:%v, %v", resp.ErrorMessage, requestId)
		return resp, errors.New("信令返回异常" + requestId)
	}
	if resp.ErrorCode == 10004 {
		log.Printf("信令提示平台%v状态为离线状态", deviceId)
		return resp, errors.New("平台已离线" + requestId)
	}
	log.Printf("推流失败, device_id:[%v]--requestId:[%v]", deviceId, requestId)
	return resp, errors.New("推流失败" + requestId)
}

func SendStopPushStreamRequest(deviceId string, url string, params *StopPlatformPushStreamParam, requestId string) (*CommonResp, error) {
	body, err := SendReqWithTimeout("PUT", url, params, 10*time.Second, requestId)
	log.Printf("信令返回：%v, %v", string(body), requestId)
	if err != nil {
		return nil, errors.New("访问信令异常" + requestId)
	}
	log.Println("信令服务器返回: ", string(body))
	resp := &CommonResp{}
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, errors.New("信令返回异常" + requestId)
	}
	if resp.ErrorCode == 0 {
		log.Printf("信令提示停止推流成功,device_id:[%v]--requestId:[%v]", deviceId, requestId)
		return resp, nil
	}
	if resp.ErrorCode == 10003 {
		log.Printf("信令提示请求参数不正确, %v", requestId)
		return resp, errors.New("请求信令参数异常" + requestId)
	}
	if resp.ErrorCode == 10004 {
		log.Printf("信令提示平台%v状态为离线状态", deviceId)
		return resp, errors.New("平台已离线" + requestId)
	}
	log.Printf("停止推流失败, device_id:[%v]--requestId:[%v]", deviceId, requestId)
	return resp, errors.New("推流失败" + requestId)
}