package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/constant"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/redis"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
	util "github.com/ldChengyi/CIOT-CPF/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(req *dto.UserLoginDTO, c *gin.Context) (string, error) {

	if err := CheckAuth(req.Username, req.Password); err != nil {
		return "", err
	}

	token, err := util.GenerateToken(req.Username, req.Password)
	if err != nil {
		return "", err
	}

	err = redis.RedisClient.Set(redis.Ctx, fmt.Sprintf("%s%s", constant.USER_LOGIN, req.Username), token, time.Hour*24).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func Register(req *dto.UserRegisterDTO, c *gin.Context) error {
	_, err := mapper.GetUserByUsername(req.Username)
	if err == nil {
		return fmt.Errorf("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newuser := entity.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	return mapper.CreateUser(&newuser)
}

func CheckAuth(username string, password string) error {
	user, err := mapper.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("未查询到用户，请检查用户名是否正确")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("密码校验失败，请检查密码是否正确")
	}
	return nil
}
