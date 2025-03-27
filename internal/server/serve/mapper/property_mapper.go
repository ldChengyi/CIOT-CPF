package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	gormutil "github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper/gorm_util"
)

func ListProperties(page int, pageSize int, filters map[string]interface{}) ([]*entity.Property, int64, error) {
	var properties []*entity.Property
	var total int64

	query := db.DB.Model(&entity.Property{})
	query = gormutil.BuildQuery(query, filters, []string{"property_name", "identifier"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&properties).Error

	return properties, total, err
}

func GetPropertyByID(id int64) (*entity.Property, error) {
	var property entity.Property
	err := db.DB.Where("id = ?", id).First(&property).Error
	return &property, err
}

func GetPropertyByPropertyName(propertyName string) (*entity.Property, error) {
	var property entity.Property
	err := db.DB.Where("property_name = ?", propertyName).First(&property).Error
	return &property, err
}

func GetAllPropertyIDsByModuleID(moduleId int64) ([]int64, error) {
	var propertyIDs []int64

	err := db.DB.Model(&entity.ModuleProperty{}).
		Where("module_id = ?", moduleId).
		Pluck("property_id", &propertyIDs).Error

	return propertyIDs, err
}

func GetAllPropertiesByModuleID(moduleID int64) ([]*entity.Property, error) {
	var properties []*entity.Property

	err := db.DB.
		Where("id IN (?)",
			db.DB.Model(&entity.ModuleProperty{}).
				Select("property_id").
				Where("module_id = ?", moduleID),
		).
		Find(&properties).Error

	return properties, err
}

func SaveOrUpdateProperty(property *entity.Property) error {
	return db.DB.Save(property).Error
}

func DeleteProperty(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&entity.Property{}).Error
}
