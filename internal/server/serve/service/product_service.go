package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
	util "github.com/ldChengyi/CIOT-CPF/pkg/util"
)

func ListProducts(req *dto.QueryDTO, c *gin.Context) ([]*vo.ProductAdminVo, int64, error) {
	products, total, err := mapper.ListProducts(req.Page, req.PageSize, req.Filters)
	if err != nil {
		return nil, 0, err
	}

	productVos := []*vo.ProductAdminVo{}
	for _, product := range products {
		productVo, err := util.MapEntityToVO(product, &vo.ProductAdminVo{})
		if err != nil {
			return nil, 0, err
		}

		productVos = append(productVos, productVo.(*vo.ProductAdminVo))
	}
	return productVos, total, err
}

func SaveOrUpdateProduct(req *dto.ProductDTO, c *gin.Context) error {
	var product *entity.Product

	// 如果传入的ID大于0，则尝试查询已有记录
	if req.ID > 0 {
		if p, err := mapper.GetProductByID(req.ID); err == nil {
			product = p
		}
	}

	// 如果未查到记录，则新建一个产品实例
	if product == nil {
		product = &entity.Product{}
	}

	// 更新字段
	product.ProductKey = req.ProductKey
	product.ProductName = req.ProductName
	product.ProductSecret = req.ProductSecret
	product.Desc = req.Desc

	return mapper.SaveOrUpdateProduct(product)
}

func DeleteProduct(req *dto.IdDeleteDTO, c *gin.Context) error {
	return mapper.DeleteProduct(req.ID)
}
