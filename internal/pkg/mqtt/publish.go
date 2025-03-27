package mqtt

import "log"

func Publish(topic string, payload string) error {
	token := MQTTClient.Publish(topic, 0, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("[LOG ERROR] 发送消息至topic失败 %s: %v", topic, token.Error())
		return token.Error()
	}
	log.Printf("[LOG INFO] 已经发布消息值topic %s: %s", topic, payload)
	return nil
}

func InitPub() {
}
