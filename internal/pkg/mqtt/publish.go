package mqtt

import (
	"fmt"
	"log"
	"time"
)

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

// 1：前进, 2：后退, 3：左转, 4：右转, 5：停止
// 消息格式: {"code":200,"data":<command>,"dataTime":"<timestamp>"}
func PublishTireControl(topic string, command int) error {

	now := time.Now()
	message := fmt.Sprintf(`{"code":200,"data":%d,"dataTime":"%s"}`,
		command, now.Format(time.RFC3339))
	token := MQTTClient.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		log.Printf("[LOG ERROR] 发送轮胎控制消息至 topic 失败 %s: %v", topic, token.Error())
		return token.Error()
	}
	log.Printf("[LOG INFO] 已发布轮胎控制消息至 topic %s: %s", topic, message)
	return nil
}
