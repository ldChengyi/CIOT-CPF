package dto

type DeviceDTO struct {
	ID           int64  `json:"id" xml:"id" form:"id" query:"id"`
	ProductID    int64  `json:"product_id" xml:"product_id" form:"product_id" query:"product_id"`
	DeviceName   string `json:"device_name" xml:"device_name" form:"device_name" query:"device_name"`
	DeviceSecret string `json:"device_secret" xml:"device_secret" form:"device_secret" query:"device_secret"`
	ModuleID     int64  `json:"module_id" xml:"module_id" form:"module_id" query:"module_id"`
	Desc         string `json:"desc" xml:"desc" form:"desc" query:"desc"`
}
