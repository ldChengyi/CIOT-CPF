package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	gormutil "github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper/gorm_util"
)

func ListDevices(page int, pageSize int, filters map[string]interface{}) ([]*entity.Device, int64, error) {
	var devices []*entity.Device
	var total int64

	query := db.DB.Model(&entity.Device{})
	query = gormutil.BuildQuery(query, filters, []string{"device_name", "device_key"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&devices).Error

	return devices, total, err
}

func GetDeviceByID(id int64) (*entity.Device, error) {
	var device entity.Device
	err := db.DB.Where("id = ?", id).First(&device).Error
	return &device, err
}

func GetDeviceByDeviceName(deviceName string) (*entity.Device, error) {
	var device entity.Device
	err := db.DB.Where("device_name = ?", deviceName).First(&device).Error
	return &device, err
}

func GetDeviceProductKeyByID(id int64) (string, error) {
	var productKey string

	err := db.DB.Table("ciot_device").Select("ciot_product.product_key").Joins("JOIN ciot_product ON ciot_product.id = ciot_device.product_id").Where("ciot_device.id = ?", id).
		Scan(&productKey).Error

	return productKey, err
}

func SaveOrUpdateDevice(device *entity.Device) error {
	return db.DB.Save(device).Error
}

func DeleteDevice(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&entity.Device{}).Error
}
