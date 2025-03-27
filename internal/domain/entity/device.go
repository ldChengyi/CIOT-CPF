package entity

import "time"

// Device 表示设备信息，继承了 Model 基础结构，包含了与设备相关的字段
type Device struct {
	Model
	ProductID    int64      `gorm:"not null" json:"product_id"`
	DeviceName   string     `gorm:"not null;size:255" json:"device_name"`
	DeviceSecret string     `gorm:"size:255" json:"device_secret"`
	ModuleID     int64      `json:"module_id"`
	Desc         string     `gorm:"type:text" json:"desc"`
	Status       int64      `gorm:"default:0" json:"status"`
	LastOnline   *time.Time `json:"last_online"`
}
