package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
	util "github.com/ldChengyi/CIOT-CPF/pkg/util"
)

func ListProperties(req *dto.QueryDTO, c *gin.Context) ([]*vo.PropertyAdminVo, int64, error) {
	properties, total, err := mapper.ListProperties(req.Page, req.PageSize, req.Filters)
	if err != nil {
		return nil, 0, err
	}

	propertyVos := []*vo.PropertyAdminVo{}
	for _, property := range properties {
		propertyVo, err := util.MapEntityToVO(property, &vo.PropertyAdminVo{})
		if err != nil {
			return nil, 0, err
		}

		propertyVos = append(propertyVos, propertyVo.(*vo.PropertyAdminVo))
	}
	return propertyVos, total, err
}

func SaveOrUpdateProperty(req *dto.PropertyDTO, c *gin.Context) error {
	var property *entity.Property

	if req.ID > 0 {
		if p, err := mapper.GetPropertyByID(req.ID); err == nil {
			property = p
		}
	}

	if property == nil {
		property = &entity.Property{}
	}

	property.PropertyName = req.PropertyName
	property.Identifier = req.Identifier
	property.DataType = req.DataType
	property.Desc = req.Desc

	return mapper.SaveOrUpdateProperty(property)
}

func DeleteProperty(req *dto.IdDeleteDTO, c *gin.Context) error {
	return mapper.DeleteProperty(req.ID)
}
