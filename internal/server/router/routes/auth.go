package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", controller.Login)
	rg.POST("/register", controller.Register)
}
