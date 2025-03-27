package vo

import "time"

type DeviceAdminVo struct {
	ID           int64      `json:"id"`
	ProductID    int64      `json:"product_id"`
	DeviceName   string     `json:"device_name"`
	DeviceSecret string     `json:"device_secret"`
	ModuleID     int64      `json:"module_id"`
	Desc         string     `json:"desc"`
	Status       int64      `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	LastOnline   *time.Time `json:"last_online"`
}
