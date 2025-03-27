package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	gormutil "github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper/gorm_util"
)

func ListModules(page int, pageSize int, filters map[string]interface{}) ([]*entity.Module, int64, error) {
	var modules []*entity.Module
	var total int64

	query := db.DB.Model(&entity.Module{})
	query = gormutil.BuildQuery(query, filters, []string{"module_name"})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Scopes(
		gormutil.PaginateScope(page, pageSize),
	).Find(&modules).Error

	return modules, total, err
}

func GetModuleByID(id int64) (*entity.Module, error) {
	var module entity.Module
	err := db.DB.Where("id = ?", id).First(&module).Error
	return &module, err
}

func GetModuleByModuleName(moduleName string) (*entity.Module, error) {
	var module entity.Module
	err := db.DB.Where("module_name = ?", moduleName).First(&module).Error
	return &module, err
}

func SaveOrUpdateModule(module *entity.Module) error {
	return db.DB.Save(module).Error
}

func DeleteModule(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&entity.Module{}).Error
}
