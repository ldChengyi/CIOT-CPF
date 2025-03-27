package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ldChengyi/CIOT-CPF/setting"
)

var MQTTClient mqtt.Client

func InitMQTT() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", setting.MQTTBroker, setting.MQTTPort))
	opts.SetClientID(setting.MQTTClientID)
	opts.SetUsername(setting.MQTTUsername)
	opts.SetPassword(setting.MQTTPassword)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	MQTTClient = mqtt.NewClient(opts)
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("[LOG ERROR]: 连接至MQTTBroker失败!: %v", token.Error())
	}

	if err := InitSub(); err != nil {
		return err
	}

	return nil
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("[LOG INFO] 成功连接至MQTTBroker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("[LOG WARNING] 连接已丢失: %v", err)
}

func CloseMQTT() {
	if MQTTClient.IsConnected() {
		MQTTClient.Disconnect(250) // 250ms 超时时间
		log.Println("[LOG INFO] MQTT已关闭")
	}
}
