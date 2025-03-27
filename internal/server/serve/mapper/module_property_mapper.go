package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
)

func CreateModuleProperty(mp *entity.ModuleProperty) error {
	return db.DB.Create(mp).Error
}

func DeleteAllModulePropertiesByModuleID(moduleId int64) error {
	return db.DB.Where("module_id = ?", moduleId).Delete(&entity.ModuleProperty{}).Error
}
