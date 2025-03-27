package mqtt

import (
	"fmt"
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/redis"
)

func Subscribe(topic string, callback mqtt.MessageHandler) error {
	token := MQTTClient.Subscribe(topic, 0, callback)
	token.Wait()
	if token.Error() != nil {
		log.Printf("[LOG INFO] 订阅topic失败 %s: %v", topic, token.Error())
		return token.Error()
	}
	log.Printf("[LOG INFO] 成功订阅至topic %s", topic)
	return nil
}

func InitSub() error {
	return Subscribe("/sys/+/+/thing/property/post", UploadDeviceProperties)
}

func UploadDeviceProperties(client mqtt.Client, msg mqtt.Message) {
	topic, payload := msg.Topic(), msg.Payload()
	parts := strings.Split(topic, "/")
	if len(parts) >= 5 {
		productKey, deviceName := parts[2], parts[3]
		key := fmt.Sprintf("%s:%s", productKey, deviceName) // 移除多余的 `:`

		// 确保 payload 以字符串存入 Redis
		err := redis.RedisClient.Set(redis.Ctx, key, string(payload), time.Hour*72).Err()
		if err != nil {
			log.Println("[LOG ERROR] Redis Set Error:", err)
		}
	}
}
