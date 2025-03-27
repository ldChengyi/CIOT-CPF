package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterModuleRoutes(rg *gin.RouterGroup) {
	rg.GET("/modules", controller.ListModules)
	rg.POST("/modules", controller.SaveOrUpdateModule)
	rg.DELETE("/modules", controller.DeleteModule)
}
