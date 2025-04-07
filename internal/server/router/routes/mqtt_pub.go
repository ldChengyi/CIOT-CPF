package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterMQTTPubRouters(rg *gin.RouterGroup) {
	rg.POST("/mqtt/publish", controller.PublishTireControl)
}
