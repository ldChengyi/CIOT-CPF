package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	gormutil "github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper/gorm_util"
)

func ListDeviceVideos(page int, pageSize int, filters map[string]interface{}) ([]*entity.DeviceVideo, int64, error) {
	var deviceVideos []*entity.DeviceVideo
	var total int64

	query := db.DB.Model(&entity.DeviceVideo{})
	query = gormutil.BuildQuery(query, filters, []string{"video_url"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&deviceVideos).Error

	return deviceVideos, total, err
}

func ListDeviceVideoAdminVos(page int, pageSize int, filters map[string]interface{}) ([]*vo.DeviceVideoAdminVo, int64, error) {
	var deviceVideoAdminVos []*vo.DeviceVideoAdminVo
	var total int64

	query := db.DB.Table("ciot_device_video").
		Select("ciot_device_video.id, ciot_device.device_name, ciot_device_video.device_id, ciot_device_video.enable_video, ciot_device_video.video_url, ciot_device_video.stream_type, ciot_device_video.status, ciot_device_video.created_at").
		Joins("left join ciot_device on ciot_device_video.device_id = ciot_device.id")
	query = gormutil.BuildQuery(query, filters, []string{"video_url"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&deviceVideoAdminVos).Error

	return deviceVideoAdminVos, total, err
}

func GetDeviceVideoAdminVoByID(id int64) (*vo.DeviceVideoAdminVo, error) {
	var deviceVideoAdminVo vo.DeviceVideoAdminVo
	err := db.DB.Table("ciot_device_video").
		Select("ciot_device_video.id, ciot_device.device_name, ciot_device_video.device_id, ciot_device_video.enable_video, ciot_device_video.video_url, ciot_device_video.stream_type, ciot_device_video.status, ciot_device_video.created_at").
		Joins("left join ciot_device on ciot_device_video.device_id = ciot_device.id").Where("ciot_device_video.id = ?", id).Find(&deviceVideoAdminVo).Error
	return &deviceVideoAdminVo, err
}

func GetDeviceVideoByID(id int64) (*entity.DeviceVideo, error) {
	var deviceVideo entity.DeviceVideo
	err := db.DB.Where("id = ?", id).First(&deviceVideo).Error
	return &deviceVideo, err
}

func SaveOrUpdateDeviceVideo(dv *entity.DeviceVideo) error {
	return db.DB.Save(dv).Error
}

func DeleteDeviceVideo(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&entity.DeviceVideo{}).Error
}
