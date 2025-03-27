package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
)

// Login godoc
// @Summary 用户登录
// @Description 用户通过用户名和密码登录，并返回 token
// @Accept  json
// @Produce  json
// @Param   login body dto.UserLoginDTO true "用户登录信息"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "登录成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	req := new(dto.UserLoginDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	token, err := service.Login(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("登录成功", map[string]interface{}{
		"token": token,
	}))

}

// Register godoc
// @Summary 用户注册
// @Description 用户注册，提供用户名和密码进行注册
// @Accept  json
// @Produce  json
// @Param   register body dto.UserRegisterDTO true "用户注册信息"
// @Success 200 {object} vo.Response "注册成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	req := new(dto.UserRegisterDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.Register(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("注册成功", nil))
}
