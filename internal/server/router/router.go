package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/server/router/routes"
	corsMiddleware "github.com/ldChengyi/CIOT-CPF/pkg/middleware/cors"
	jwtMiddleware "github.com/ldChengyi/CIOT-CPF/pkg/middleware/jwt"

	_ "github.com/ldChengyi/CIOT-CPF/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	//使用跨域请求中间件
	r.Use(corsMiddleware.InitCORS())
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerRouters(r)
	return r
}

func registerRouters(r *gin.Engine) {

	auth := r.Group("/auth")
	routes.RegisterAuthRoutes(auth)

	admin := r.Group("/admin")
	admin.Use(jwtMiddleware.JWT())
	routes.RegisterProductRoutes(admin)
	routes.RegisterDeviceRoutes(admin)
	routes.RegisterModuleRoutes(admin)
	routes.RegisterPropertyRoutes(admin)
	routes.RegisterDeviceVideoRoutes(admin)

	user := r.Group("/user")
	routes.RegisterMQTTPubRouters(user)
}
