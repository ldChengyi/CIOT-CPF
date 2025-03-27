package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/redis"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
)

func GetDevicePropertiesValue(req *dto.IdTransferDTO, c *gin.Context) ([]map[string]interface{}, error) {
	response := []map[string]interface{}{}

	device, err := mapper.GetDeviceByID(req.ID)
	if err != nil {
		return response, err
	}

	properties, err := mapper.GetAllPropertiesByModuleID(device.ModuleID)
	if err != nil {
		return response, err
	}

	productKey, err := mapper.GetDeviceProductKeyByID(device.ID)
	if err != nil || productKey == "" {
		return response, err
	}

	redisKey := fmt.Sprintf("%s:%s", productKey, device.DeviceName)
	redisData, _ := redis.RedisClient.Get(redis.Ctx, redisKey).Result()

	var data map[string]interface{}
	if redisData != "" {
		json.Unmarshal([]byte(redisData), &data)
	}

	for _, property := range properties {
		value := data[property.Identifier]
		if value == nil {
			value = ""
		}
		response = append(response, map[string]interface{}{
			"id":         property.ID,
			"name":       property.PropertyName,
			"identifier": property.Identifier,
			"value":      value,
		})
	}

	return response, nil
}
