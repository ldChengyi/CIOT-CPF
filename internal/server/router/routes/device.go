package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterDeviceRoutes(rg *gin.RouterGroup) {
	rg.GET("/devices", controller.ListDevices)
	rg.GET("/devices/properties/val", controller.GetDevicePropertiesValue)
	rg.POST("/devices", controller.SaveOrUpdateDevice)
	rg.DELETE("/devices", controller.DeleteDevice)
}
