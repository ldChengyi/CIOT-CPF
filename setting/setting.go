package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Config *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string

	MQTTBroker   string
	MQTTPort     int
	MQTTClientID string
	MQTTUsername string
	MQTTPassword string

	DBType        string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBName        string
	DBTablePrefix string

	RedisAddr     string
	RedisPassword string
	RedisDB       int
)

func Init() {
	var err error
	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取配置文件失败，请检查 'conf/app.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadApp()
	loadDB()
	loadRedis()
	loadMQTT()
}

func loadBase() {
	RunMode = Config.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Config.GetSection("server")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取文件配置模块[server]失败: %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := Config.GetSection("app")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取文件配置模块[app]失败: %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

func loadMQTT() {
	sec, err := Config.GetSection("mqtt")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取文件配置模块[mqtt]失败: %v", err)
	}

	MQTTBroker = sec.Key("BROKER").MustString("localhost")
	MQTTPort = sec.Key("PORT").MustInt(1883)
	MQTTClientID = sec.Key("CLIENT_ID").MustString("ciot_plateform")
	MQTTUsername = sec.Key("USERNAME").MustString("")
	MQTTPassword = sec.Key("PASSWORD").MustString("")
}

func loadDB() {
	sec, err := Config.GetSection("database")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取文件配置模块[database]失败: %v", err)
	}
	DBType = sec.Key("TYPE").String()
	DBName = sec.Key("NAME").String()
	DBUser = sec.Key("USER").String()
	DBPassword = sec.Key("PASSWORD").String()
	DBHost = sec.Key("HOST").String()
	DBTablePrefix = sec.Key("TABLE_PREFIX").String()
}

func loadRedis() {
	sec, err := Config.GetSection("redis")
	if err != nil {
		log.Fatalf("[LOG WARNING] 尝试获取文件配置模块[redis]失败: %v", err)
	}
	RedisAddr = fmt.Sprintf("%s:%d", sec.Key("HOST").String(), sec.Key("PORT").MustInt(6379))
	RedisPassword = sec.Key("PASSWORD").String()
	RedisDB, err = sec.Key("DB").Int()
	if err != nil {
		log.Fatalf("[LOG WARNING] 请在配置文件中选择启用的RedisDB: %v", err)
	}

}
