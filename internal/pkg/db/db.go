package db

import (
	"fmt"
	"log"

	"github.com/ldChengyi/CIOT-CPF/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() error {
	// 构建 DSN
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DBUser,
		setting.DBPassword,
		setting.DBHost,
		setting.DBName,
	)

	// 打开数据库连接，配置日志级别和命名策略（表名带前缀、单数形式）
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DBTablePrefix, // 表名前缀
			SingularTable: true,                  // 使用单数表名
		},
	})
	if err != nil {
		return fmt.Errorf("[LOG ERROR] 连接至数据库失败!: %v", err)
	}

	// 获取底层 sql.DB 对象，以便设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("[LOG ERROR] 获取数据库连接池失败!: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("[LOG INFO] 成功连接至数据库")
	return nil
}

func CloseDB() {
	if DB != nil {
		// 获取底层 sql.DB 对象关闭连接
		if sqlDB, err := DB.DB(); err == nil {
			sqlDB.Close()
		}
	}
	log.Println("[LOG INFO] 数据库已关闭")
}
