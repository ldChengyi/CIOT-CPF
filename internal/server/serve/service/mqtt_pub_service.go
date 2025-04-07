package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/mqtt"
)

func PublishTireControl(req *dto.CommandTransferDTO, c *gin.Context) error {
	// 1：前进, 2：后退, 3：左转, 4：右转, 5：停止
	// 消息格式: {"code":200,"data":<command>,"dataTime":"<timestamp>"}
	if err := mqtt.PublishTireControl("/sys/zzuli-flc/FLC1/thing/property/command", int(req.Command)); err != nil {
		return err
	}

	return nil
}
