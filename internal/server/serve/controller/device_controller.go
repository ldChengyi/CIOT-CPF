package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

// ListDevices godoc
// @Summary 获取设备列表
// @Description 根据查询条件获取设备列表，并返回分页数据
// @Accept  json
// @Produce  json
// @Param   page       query     int    true  "页码"
// @Param   page_size query     int    true  "每页大小"
// @Param   filter     query     string true  "过滤条件，格式为 filter[key]=value, 多个过滤条件用 & 分隔"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/devices [get]
func ListDevices(c *gin.Context) {
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

	devices, total, err := service.ListDevices(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list":  devices,
		"total": total,
	}))
}

// GetDevicePropertiesValue godoc
// @Summary 获取设备所有属性值
// @Description 获取设备关联模块的所有属性的值, 值取于redis中
// @Accept  json
// @Produce  json
// @Param   id       query     int    true  "设备id"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/devices/properties/val [get]
func GetDevicePropertiesValue(c *gin.Context) {
	req := new(dto.IdTransferDTO)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	propertiesValue, err := service.GetDevicePropertiesValue(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list": propertiesValue,
	}))

}

// SaveOrUpdateDevice godoc
// @Summary 保存或更新设备
// @Description 根据提供的设备信息进行设备的保存或更新
// @Accept  json
// @Produce  json
// @Param   device body dto.DeviceDTO true "设备信息"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/devices [post]
func SaveOrUpdateDevice(c *gin.Context) {
	req := new(dto.DeviceDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.SaveOrUpdateDevice(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}

// DeleteDevice godoc
// @Summary 删除设备
// @Description 根据设备ID删除指定的设备
// @Accept  json
// @Produce  json
// @Param   id      query     int64  true "设备ID"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/devices [delete]
func DeleteDevice(c *gin.Context) {
	req := new(dto.IdDeleteDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.DeleteDevice(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}
