package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

// ListPropertys godoc
// @Summary 获取属性列表
// @Description 根据查询条件获取属性列表，并返回分页数据
// @Accept  json
// @Produce  json
// @Param   page       query     int    true  "页码"
// @Param   page_size query     int    true  "每页大小"
// @Param   filter     query     string true  "过滤条件，格式为 filter[key]=value, 多个过滤条件用 & 分隔"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/properties [get]
func ListProperties(c *gin.Context) {
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

	properties, total, err := service.ListProperties(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list":  properties,
		"total": total,
	}))
}

// SaveOrUpdateProperty godoc
// @Summary 保存或更新属性
// @Description 根据提供的属性信息进行属性的保存或更新
// @Accept  json
// @Produce  json
// @Param   property body dto.PropertyDTO true "属性信息"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/properties [post]
func SaveOrUpdateProperty(c *gin.Context) {
	req := new(dto.PropertyDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.SaveOrUpdateProperty(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}

// DeleteProperty godoc
// @Summary 删除属性
// @Description 根据属性ID删除指定的属性
// @Accept  json
// @Produce  json
// @Param   id      query     int64  true "属性ID"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/properties [delete]
func DeleteProperty(c *gin.Context) {
	req := new(dto.IdDeleteDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.DeleteProperty(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}
