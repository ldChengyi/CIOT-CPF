package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterPropertyRoutes(rg *gin.RouterGroup) {
	rg.GET("/properties", controller.ListProperties)
	rg.POST("/properties", controller.SaveOrUpdateProperty)
	rg.DELETE("/properties", controller.DeleteProperty)
}
