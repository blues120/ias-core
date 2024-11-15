package biz

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/log"
)

// 告警类型管理
type MqttRepo interface {
	// 发布告警信息
	Publish(topic string, qos byte, payload any) mqtt.Token

	// PublishRetained 发布保留告警信息
	PublishRetained(topic string, qos byte, payload any) mqtt.Token

	// 查询指定告警类型
	Subscribe(topic string, qos byte, messageHandler mqtt.MessageHandler) mqtt.Token

	// Unsubscribe 取消订阅
	Unsubscribe(topic ...string) mqtt.Token
}

type MqttUsecase struct {
	repo MqttRepo
	log  *log.Helper
}

func NewMqttUsecase(repo MqttRepo, logger log.Logger) *MqttUsecase {
	return &MqttUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (m *MqttUsecase) Publish(topic string, qos byte, payload any) mqtt.Token {
	return m.repo.Publish(topic, qos, payload)
}

func (m *MqttUsecase) PublishRetained(topic string, qos byte, payload any) mqtt.Token {
	return m.repo.PublishRetained(topic, qos, payload)
}

func (m *MqttUsecase) Subscribe(topic string, qos byte, messageHandler mqtt.MessageHandler) mqtt.Token {
	return m.repo.Subscribe(topic, qos, messageHandler)
}

func (m *MqttUsecase) Unsubscribe(topic ...string) mqtt.Token {
	return m.repo.Unsubscribe(topic...)
}
