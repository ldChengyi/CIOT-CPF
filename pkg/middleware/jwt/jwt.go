package jwtMiddleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/constant"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/redis"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

var (
	TOKEN_ERROR = 401
	SUCCESS     = 200
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = TOKEN_ERROR
		} else {
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}

			claims, err := util.ParseToken(token)
			if err != nil {
				code = TOKEN_ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = TOKEN_ERROR
			} else {
				redisToken := redis.RedisClient.Get(c, fmt.Sprintf("%s%s", constant.USER_LOGIN, claims.Username)).Val()
				if redisToken == "" || redisToken != token {
					code = TOKEN_ERROR
				} else {
					// 刷新 token 的过期时间
					err = redis.RedisClient.Expire(c, fmt.Sprintf("%s%s", constant.USER_LOGIN, claims.Username), time.Hour*24).Err()
					if err != nil {
						code = TOKEN_ERROR
					}
				}
			}
		}

		if code != SUCCESS {
			if !c.Writer.Written() {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  "Token异常，请重新登录",
					"data": data,
				})
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
