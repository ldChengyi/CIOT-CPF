package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/ldChengyi/CIOT-CPF/setting"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     setting.RedisAddr,
		Password: setting.RedisPassword,
		DB:       setting.RedisDB,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		return fmt.Errorf("[LOG ERROR] 连接至Redis失败!: %v", err)
	}

	log.Println("[LOG INFO] 成功连接至Redis")
	return nil
}

func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Println("[LOG INFO] 关闭Redis失败:", err)
	} else {
		log.Println("[LOG INFO] Redis已关闭")
	}
}
