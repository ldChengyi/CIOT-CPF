package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

// ListProducts godoc
// @Summary 获取产品列表
// @Description 根据查询条件获取产品列表，并返回分页数据
// @Accept  json
// @Produce  json
// @Param   page       query     int    true  "页码"
// @Param   page_size query     int    true  "每页大小"
// @Param   filter     query     string true  "过滤条件，格式为 filter[key]=value, 多个过滤条件用 & 分隔"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/products [get]
func ListProducts(c *gin.Context) {
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

	products, total, err := service.ListProducts(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list":  products,
		"total": total,
	}))
}

// SaveOrUpdateProduct godoc
// @Summary 保存或更新产品
// @Description 根据提供的产品信息进行产品的保存或更新
// @Accept  json
// @Produce  json
// @Param   product body dto.ProductDTO true "产品信息"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/products [post]
func SaveOrUpdateProduct(c *gin.Context) {
	req := new(dto.ProductDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.SaveOrUpdateProduct(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}

// DeleteProduct godoc
// @Summary 删除产品
// @Description 根据产品ID删除指定的产品
// @Accept  json
// @Produce  json
// @Param   id      query     int64  true "产品ID"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/products [delete]
func DeleteProduct(c *gin.Context) {
	req := new(dto.IdDeleteDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.DeleteProduct(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}
