package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
)

func ListDeviceVideos(req *dto.QueryDTO, c *gin.Context) ([]*vo.DeviceVideoAdminVo, int64, error) {
	deviceVideoAdminVos, total, err := mapper.ListDeviceVideoAdminVos(req.Page, req.PageSize, req.Filters)
	if err != nil {
		return nil, 0, err
	}

	return deviceVideoAdminVos, total, err
}

func GetDeviceVideoAdminVoByID(req *dto.IdTransferDTO, c *gin.Context) (*vo.DeviceVideoAdminVo, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return mapper.GetDeviceVideoAdminVoByID(id)
}

func SaveOrUpdateDeviceVideo(req *dto.DeviceVideoDTO, c *gin.Context) error {
	var deviceVideo *entity.DeviceVideo

	if req.ID > 0 {
		if dv, err := mapper.GetDeviceVideoByID(req.ID); err == nil {
			deviceVideo = dv
		}
	}

	if deviceVideo == nil {
		deviceVideo = &entity.DeviceVideo{}
	}

	deviceVideo.DeviceID = req.DeviceID
	deviceVideo.EnableVideo = req.EnableVideo
	deviceVideo.VideoURL = req.VideoURL
	deviceVideo.StreamType = req.StreamType
	deviceVideo.Status = req.Status

	return mapper.SaveOrUpdateDeviceVideo(deviceVideo)
}

func DeleteDeviceVideo(req *dto.IdDeleteDTO, c *gin.Context) error {
	return mapper.DeleteDeviceVideo(req.ID)
}
