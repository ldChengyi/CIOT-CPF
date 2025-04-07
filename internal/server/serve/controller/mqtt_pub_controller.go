package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
)

func PublishTireControl(c *gin.Context) {
	req := new(dto.CommandTransferDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.PublishTireControl(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))

}
