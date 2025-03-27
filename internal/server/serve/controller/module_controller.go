package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

// ListModules godoc
// @Summary 获取模块列表
// @Description 根据查询条件获取模块列表，并返回分页数据
// @Accept  json
// @Produce  json
// @Param   page       query     int    true  "页码"
// @Param   page_size query     int    true  "每页大小"
// @Param   filter     query     string true  "过滤条件，格式为 filter[key]=value, 多个过滤条件用 & 分隔"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/modules [get]
func ListModules(c *gin.Context) {
	req := new(dto.QueryDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	filters, err := util.ParseFilters(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}
	req.Filters = filters

	modules, total, err := service.ListModules(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list":  modules,
		"total": total,
	}))
}

// SaveOrUpdateModule godoc
// @Summary 保存或更新模块
// @Description 根据提供的模块信息进行模块的保存或更新
// @Accept  json
// @Produce  json
// @Param   module body dto.ModuleDTO true "模块信息"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/modules [post]
func SaveOrUpdateModule(c *gin.Context) {
	req := new(dto.ModuleDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.SaveOrUpdateModule(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}

// DeleteModule godoc
// @Summary 删除模块
// @Description 根据模块ID删除指定的模块
// @Accept  json
// @Produce  json
// @Param   id      query     int64  true "模块ID"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/modules [delete]
func DeleteModule(c *gin.Context) {
	req := new(dto.IdDeleteDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.DeleteModule(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}
