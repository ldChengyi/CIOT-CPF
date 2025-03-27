package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterDeviceVideoRoutes(rg *gin.RouterGroup) {
	rg.GET("/deviceVideos", controller.ListDeviceVideos)
	rg.GET("/deviceVideos/:id", controller.GetDeviceVideoAdminByID)
	rg.POST("/deviceVideos", controller.SaveOrUpdateDeviceVideo)
	rg.DELETE("/deviceVideos", controller.DeleteDeviceVideo)
}
