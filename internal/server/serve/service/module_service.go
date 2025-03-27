package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/mapper"
	util "github.com/ldChengyi/CIOT-CPF/pkg/util"
)

func ListModules(req *dto.QueryDTO, c *gin.Context) ([]*vo.ModuleAdminVo, int64, error) {
	modules, total, err := mapper.ListModules(req.Page, req.PageSize, req.Filters)
	if err != nil {
		return nil, 0, err
	}

	moduleVos := []*vo.ModuleAdminVo{}
	for _, module := range modules {
		moduleVo, err := util.MapEntityToVO(module, &vo.ModuleAdminVo{})
		if err != nil {
			return nil, 0, err
		}

		propertyIds, err := mapper.GetAllPropertyIDsByModuleID(module.ID)
		if err != nil {
			return nil, 0, err
		}
		fmt.Println(propertyIds)
		moduleVo.(*vo.ModuleAdminVo).PropertyIDs = propertyIds

		moduleVos = append(moduleVos, moduleVo.(*vo.ModuleAdminVo))
	}
	return moduleVos, total, err
}

func SaveOrUpdateModule(req *dto.ModuleDTO, c *gin.Context) error {
	tx := db.DB.Begin()
	var module *entity.Module

	if req.ID > 0 {
		if m, err := mapper.GetModuleByID(req.ID); err == nil {
			module = m
		}
	}

	if module == nil {
		module = &entity.Module{}
	}

	module.ModuleName = req.ModuleName
	module.Desc = req.Desc
	if err := mapper.SaveOrUpdateModule(module); err != nil {
		tx.Rollback()
		return err
	}

	if err := mapper.DeleteAllModulePropertiesByModuleID(req.ID); err != nil {
		tx.Rollback()
		return err
	}
	for _, id := range req.PropertyIDs {
		moduleProperty := &entity.ModuleProperty{
			ModuleID:   req.ID,
			PropertyID: id,
		}
		if err := mapper.CreateModuleProperty(moduleProperty); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func DeleteModule(req *dto.IdDeleteDTO, c *gin.Context) error {
	tx := db.DB.Begin()
	if err := mapper.DeleteModule(req.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := mapper.DeleteAllModulePropertiesByModuleID(req.ID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
