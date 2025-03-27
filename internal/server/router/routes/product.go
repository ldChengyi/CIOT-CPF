package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/controller"
)

func RegisterProductRoutes(rg *gin.RouterGroup) {
	rg.GET("/products", controller.ListProducts)
	rg.POST("/products", controller.SaveOrUpdateProduct)
	rg.DELETE("/products", controller.DeleteProduct)
}
