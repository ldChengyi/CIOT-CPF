package corsMiddleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSConfig 定义 CORS 中间件的配置
type corsConfig struct {
	AllowedOrigins   []string // 允许的源
	AllowedMethods   []string // 允许的方法
	AllowedHeaders   []string // 允许的头部
	AllowCredentials bool     // 是否允许携带证书
}

func defaultCORSConfig() corsConfig {
	return corsConfig{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Client-Info", "X-Client-Version", "X-Client-Data", "X-Request-Id"},
		AllowCredentials: false,
	}
}

// InitCORS 返回一个 Gin 中间件函数，用于初始化 CORS 配置
func InitCORS() gin.HandlerFunc {
	return corsWithConfig(defaultCORSConfig())
}

// corsWithConfig 返回一个 CORS 中间件函数
func corsWithConfig(config corsConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 头部
		c.Header("Access-Control-Allow-Origin", strings.Join(config.AllowedOrigins, ","))
		c.Header("Access-Control-Allow-Methods", strings.Join(config.AllowedMethods, ","))
		c.Header("Access-Control-Allow-Headers", strings.Join(config.AllowedHeaders, ","))

		if config.AllowCredentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 处理预检请求，缓存预检请求结果 24 小时
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 记录 CORS 请求（可选，假设你有类似的全局日志）
		// if global.BizLog != nil {
		// 	global.BizLog.Info("CORS request",
		// 		"method", c.Request.Method,
		// 		"path", c.Request.URL.Path,
		// 		"origin", c.Request.Header.Get("Origin"),
		// 		"headers", c.Request.Header,
		// 	)
		// }

		c.Next()
	}
}
