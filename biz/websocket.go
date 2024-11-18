package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/blues120/ias-core/data/iam"
	"strings"
)

type WebsocketClientCtx struct {
	TenantID      string // 租户ID
	AccessOrgList string // 能访问用户资源的所有层级，从当前所在层级回溯到父层级获得，使用 # 分隔，如 "#1#2#"
}

// 租户信息详情
type TenantWebsocketMsg struct {
	TenantID      string
	AccessOrgList string // 组织架构
	Message       string // 消息内容，json字符串
}

func CreateWebsocketConn(ctx context.Context, upgrader websocket.Upgrader) (*WebsocketClientCtx, *websocket.Conn, error) {
	httpCtx, ok := GetHttpCtx(ctx)
	if !ok {
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
		return nil, nil, err
	}
	// 将新连接添加到客户端集合
	clientCtx := new(WebsocketClientCtx)
	clientCtx.TenantID = tenantID
	clientCtx.AccessOrgList = accessOrgList
	return clientCtx, conn, nil
}

func DecodeMsgWithTenant(msgByte []byte) (*TenantWebsocketMsg, error) {
	tenantWebsocketMsg := new(TenantWebsocketMsg)
	err := json.Unmarshal(msgByte, &tenantWebsocketMsg)
	if err != nil {
		return nil, err
	}
	return tenantWebsocketMsg, nil
}

// 检查是否可以发送
func CheckSend(clientCtx *WebsocketClientCtx, tenantWebsocketMsg *TenantWebsocketMsg) bool {
	// 判断是否是同一租户
	if clientCtx.TenantID != tenantWebsocketMsg.TenantID {
		return false
	}
	// 检查消息发送权限
	if !strings.Contains(tenantWebsocketMsg.AccessOrgList, clientCtx.AccessOrgList) {
		return false
	}
	return true
}
