package data

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

type mqttRepo struct {
	data *Data
	log  *log.Helper
}

func NewMqttRepo(data *Data, logger log.Logger) biz.MqttRepo {
	return &mqttRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m *mqttRepo) Publish(topic string, qos byte, payload any) mqtt.Token {
	return m.data.mqtt.Publish(topic, qos, false, payload)
}

func (m *mqttRepo) PublishRetained(topic string, qos byte, payload any) mqtt.Token {
	return m.data.mqtt.Publish(topic, qos, true, payload)
}

func (m *mqttRepo) Subscribe(topic string, qos byte, messageHandler mqtt.MessageHandler) mqtt.Token {
	return m.data.mqtt.Subscribe(topic, qos, messageHandler)
}

func (m *mqttRepo) Unsubscribe(topic ...string) mqtt.Token {
	return m.data.mqtt.Unsubscribe(topic...)
}
