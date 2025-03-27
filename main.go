package main

import (
	"fmt"
	"net/http"

	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/mqtt"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/redis"
	"github.com/ldChengyi/CIOT-CPF/internal/server/router"
	"github.com/ldChengyi/CIOT-CPF/setting"
)

func main() {
	//初始化配置
	setting.Init()

	//初始化各种配置
	db.InitDB()
	defer db.CloseDB()

	redis.InitRedis()
	defer redis.CloseRedis()

	mqtt.InitMQTT()
	defer mqtt.CloseMQTT()

	router := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
