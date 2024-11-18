package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/data/iam"
)

const warningAlertStreamName string = "ias:WarningAlert" //缓存队列名

// 告警类型管理
type WarningAlertRepo interface {
	// 发布告警信息
	Publish(ctx context.Context, channel string, message any) error

	// 查询指定告警类型
	Subscribe(ctx context.Context, channel string) <-chan *redis.Message

	// 消息缓存队列，写入数据
	XAdd(ctx context.Context, args *redis.XAddArgs) error

	// 消息缓存队列，读取数据
	XRead(ctx context.Context, args *redis.XReadArgs) ([]redis.XStream, error)
}

type WarningAlertUsecase struct {
	repo WarningAlertRepo
	conf *conf.WarnAlert // 提供channel
	log  *log.Helper
}

func NewWarningAlertUsecase(repo WarningAlertRepo, conf *conf.WarnAlert, logger log.Logger) *WarningAlertUsecase {
	return &WarningAlertUsecase{repo: repo, conf: conf, log: log.NewHelper(logger)}
}

// PublishByTenant 发布租户告警信息
func (uc *WarningAlertUsecase) PublishMsgWithTenant(ctx context.Context, tenantId, accessOrgList string, message any) error {
	tenantMessage, err := uc.EncodeTenantWarnMsg(tenantId, accessOrgList, message)
	if err != nil {
		return err
	}
	return uc.PublishData(ctx, tenantMessage)
}

// Publish 发布告警信息, interface/any，任意类型
func (uc *WarningAlertUsecase) Publish(ctx context.Context, message any) error {
	return uc.PublishData(ctx, &message)
}

// Subscribe 订阅告警数据，返回信息需要注意
func (uc *WarningAlertUsecase) Subscribe(ctx context.Context) chan []byte {
	return uc.SubscribeByChannel(ctx, uc.conf.Channel)
}

// 订阅channel
func (uc *WarningAlertUsecase) SubscribeByChannel(ctx context.Context, channel string) chan []byte {
	rch := uc.repo.Subscribe(ctx, channel)
	ch := make(chan []byte, 1024)
	go func() {
		for {
			msg := <-rch
			if msg == nil || msg.Payload == "" {
				continue
			}
			data := []byte(msg.Payload)
			select {
			case ch <- data:
			default:
				<-ch
				ch <- data
			}
		}
	}()
	return ch
}

// 通用走统一channel
func (uc *WarningAlertUsecase) PublishData(ctx context.Context, message any) error {
	return uc.PublishDataByChannel(ctx, uc.conf.Channel, message)
}

// 根据channel发布
func (uc *WarningAlertUsecase) PublishDataByChannel(ctx context.Context, channel string, message any) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		uc.log.Errorf("序列化告警信息失败: %v", err)
		return err
	}
	return uc.repo.Publish(ctx, channel, string(jsonData))
}

// 根据channel发布
func (uc *WarningAlertUsecase) PublishWithTenantByChannel(ctx context.Context, channel, tenantId, accessOrgList string, message any) error {
	tenantMessage, err := uc.EncodeTenantWarnMsg(tenantId, accessOrgList, message)
	if err != nil {
		return err
	}
	return uc.PublishDataByChannel(ctx, channel, tenantMessage)
}

func (uc *WarningAlertUsecase) MsgQueueAdd(ctx context.Context, streamName string, maxLen int64, message []byte) error {
	args := &redis.XAddArgs{
		Stream:     streamName, // 设置流stream的 key，消息队列名
		NoMkStream: false,      //为false，key不存在会新建
		MaxLen:     maxLen,     //消息队列最大长度，队列长度超过设置最大长度后，旧消息会被删除
		Approx:     false,      //默认false，设为true时，模糊指定stram的长度
		ID:         "*",        //消息ID，* 表示由Redis自动生成
		Values: []interface{}{
			"msg", message,
		},
	}
	return uc.repo.XAdd(ctx, args)
}

func (uc *WarningAlertUsecase) MsgQueueRead(ctx context.Context, streamName string, time time.Duration) []string {
	var Msgs []string
	args := &redis.XReadArgs{
		Streams: []string{streamName, "0"}, // 0,从头开始读取
		Block:   time,                      // 阻塞等待时间（0表示一直等待）
	}
	message, err := uc.repo.XRead(ctx, args)
	if err != nil {
		return Msgs
	}
	for _, msgs := range message {
		for _, m := range msgs.Messages {
			Msgs = append(Msgs, fmt.Sprintf("%v", m.Values["msg"]))
		}
	}
	return Msgs
}

func (uc *WarningAlertUsecase) EncodeTenantWarnMsg(tenantId, accessOrgList string, message any) (*TenantWebsocketMsg, error) {
	jsonMessageData, err := json.Marshal(message)
	if err != nil {
		uc.log.Errorf("序列化告警信息失败: %v", err)
		return nil, err
	}
	tenantMessage := new(TenantWebsocketMsg)
	tenantMessage.TenantID = tenantId
	tenantMessage.AccessOrgList = accessOrgList
	tenantMessage.Message = string(jsonMessageData)
	return tenantMessage, nil
}

func (uc *WarningAlertUsecase) DecodeMsgWithTenant(msgByte []byte) (*TenantWebsocketMsg, error) {
	tenantWarnMsg := new(TenantWebsocketMsg)
	err := json.Unmarshal(msgByte, &tenantWarnMsg)
	if err != nil {
		return nil, err
	}
	return tenantWarnMsg, nil
}

// 根据当前组织，获取所有组织
func (uc *WarningAlertUsecase) GetAccessOrgList(accessOrgList string) []string {
	list := strings.Split(strings.Trim(accessOrgList, "#"), "#")
	listLen := len(list)
	var orgList []string
	for i := 0; i < listLen; i++ {
		result := "#" + strings.Join(list[:i+1], "#") + "#"
		orgList = append(orgList, result)
	}
	return orgList
}

func (uc *WarningAlertUsecase) GetTenantStreamName(tenantId, accessOrgList string) string {
	tenantStreamName := uc.GetWarningAlertStreamName() + "/tenantId:%s/accessOrgList:%s"
	return fmt.Sprintf(tenantStreamName, tenantId, accessOrgList)
}

func (uc *WarningAlertUsecase) GetWarningAlertStreamName() string {
	return warningAlertStreamName
}

type WarnAlertClientCtx struct {
	TenantID      string // 租户ID
	AccessOrgList string // 能访问用户资源的所有层级，从当前所在层级回溯到父层级获得，使用 # 分隔，如 "#1#2#"
}

type httpCtxKey struct {
}

func SetHttpCtx(ctx context.Context, httpCxt http.Context) context.Context {
	return context.WithValue(ctx, httpCtxKey{}, httpCxt)
}

func GetHttpCtx(ctx context.Context) (http.Context, bool) {
	u, ok := ctx.Value(httpCtxKey{}).(http.Context)
	return u, ok
}

func (uc *WarningAlertUsecase) CreateWebsocketConn(ctx context.Context, upgrader websocket.Upgrader) (*WarnAlertClientCtx, *websocket.Conn, error) {
	httpCtx, ok := GetHttpCtx(ctx)
	if !ok {
		uc.log.Errorf("WebSocket连接升级失败: http请求信息获取失败")
		return nil, nil, errors.New("WebSocket连接升级失败, http请求信息获取失败")
	}
	w := httpCtx.Response()
	r := httpCtx.Request()
	userInfo, ok := iam.GetUserInfo(ctx)
	var tenantID, accessOrgList string
	if ok && userInfo != nil {
		tenantID = userInfo.TenantID
	}
	roles, ok := iam.GetOrgInfo(ctx)
	if ok && roles != nil {
		accessOrgList = roles.AccessOrgList
	}
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		uc.log.Errorf("WebSocket连接升级失败: %s", err)
		return nil, nil, err
	}
	// 将新连接添加到客户端集合
	clientCtx := new(WarnAlertClientCtx)
	clientCtx.TenantID = tenantID
	clientCtx.AccessOrgList = accessOrgList
	return clientCtx, conn, nil
}

// 检查是否可以发送
func (uc *WarningAlertUsecase) CheckSend(clientCtx *WarnAlertClientCtx, tenantWarnMsg *TenantWebsocketMsg) bool {
	// 判断是否是同一租户
	if clientCtx.TenantID != tenantWarnMsg.TenantID {
		return false
	}
	// 检查消息发送权限
	if !strings.Contains(tenantWarnMsg.AccessOrgList, clientCtx.AccessOrgList) {
		return false
	}
	return true
}
