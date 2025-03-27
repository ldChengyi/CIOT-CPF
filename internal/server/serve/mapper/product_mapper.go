package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	gormutil "github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper/gorm_util"
)

func ListProducts(page int, pageSize int, filters map[string]interface{}) ([]*entity.Product, int64, error) {
	var products []*entity.Product
	var total int64

	query := db.DB.Model(&entity.Product{})
	query = gormutil.BuildQuery(query, filters, []string{"product_name", "product_key"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&products).Error

	return products, total, err
}

func GetProductByID(id int64) (*entity.Product, error) {
	var product entity.Product
	err := db.DB.Where("id = ?", id).First(&product).Error
	return &product, err
}

func GetProductByProductName(productName string) (*entity.Product, error) {
	var product entity.Product
	err := db.DB.Where("product_name = ?", productName).First(&product).Error
	return &product, err
}

func SaveOrUpdateProduct(product *entity.Product) error {
	return db.DB.Save(product).Error
}

func DeleteProduct(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&entity.Product{}).Error
}
