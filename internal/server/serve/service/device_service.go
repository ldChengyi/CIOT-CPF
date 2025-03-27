package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
	util "github.com/ldChengyi/CIOT-CPF/pkg/util"
)

func ListDevices(req *dto.QueryDTO, c *gin.Context) ([]*vo.DeviceAdminVo, int64, error) {
	devices, total, err := mapper.ListDevices(req.Page, req.PageSize, req.Filters)
	if err != nil {
		return nil, 0, err
	}

	deviceVos := []*vo.DeviceAdminVo{}
	for _, device := range devices {
		deviceVo, err := util.MapEntityToVO(device, &vo.DeviceAdminVo{})
		if err != nil {
			return nil, 0, err
		}

		deviceVos = append(deviceVos, deviceVo.(*vo.DeviceAdminVo))
	}
	return deviceVos, total, err
}

func SaveOrUpdateDevice(req *dto.DeviceDTO, c *gin.Context) error {
	var device *entity.Device

	if req.ID > 0 {
		if d, err := mapper.GetDeviceByID(req.ID); err == nil {
			device = d
		}
	}

	if device == nil {
		device = &entity.Device{}
	}

	device.ProductID = req.ProductID
	device.DeviceName = req.DeviceName
	device.DeviceSecret = req.DeviceSecret
	device.ModuleID = req.ModuleID
	device.Desc = req.Desc
	return mapper.SaveOrUpdateDevice(device)
}

func DeleteDevice(req *dto.IdDeleteDTO, c *gin.Context) error {
	return mapper.DeleteDevice(req.ID)
}
