package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/dto"
	"github.com/ldChengyi/CIOT-CPF/internal/domain/vo"
	"github.com/ldChengyi/CIOT-CPF/internal/server/serve/service"
	"github.com/ldChengyi/CIOT-CPF/pkg/util"
)

// ListDeviceVideos godoc
// @Summary 获取设备视频视频列表
// @Description 根据查询条件获取设备视频视频列表，并返回分页数据
// @Accept  json
// @Produce  json
// @Param   page       query     int    true  "页码"
// @Param   page_size query     int    true  "每页大小"
// @Param   filter     query     string true  "过滤条件，格式为 filter[key]=value, 多个过滤条件用 & 分隔"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/deviceVideos [get]
func ListDeviceVideos(c *gin.Context) {
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

	deviceVideos, total, err := service.ListDeviceVideos(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", map[string]interface{}{
		"list":  deviceVideos,
		"total": total,
	}))
}

// GetDeviceVideoAdminByID godoc
// @Summary 获取设备视频具体数据
// @Description 获取设备视频具体数据
// @Accept  json
// @Produce  json
// @Param   id       query     int    true  "ID"
// @Success 200 {object} vo.Response{data=map[string]interface{}} "操作成功"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/deviceVideos/:id
func GetDeviceVideoAdminByID(c *gin.Context) {
	req := new(dto.IdTransferDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	deviceVideoAdminVo, err := service.GetDeviceVideoAdminVoByID(req, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功", deviceVideoAdminVo))
}

// SaveOrUpdateDeviceVideo godoc
// @Summary 保存或更新设备视频视频
// @Description 根据提供的设备视频视频信息进行设备视频视频的保存或更新
// @Accept  json
// @Produce  json
// @Param   deviceVideo body dto.DeviceVideoDTO true "设备视频视频信息"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/deviceVideos [post]
func SaveOrUpdateDeviceVideo(c *gin.Context) {
	req := new(dto.DeviceVideoDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.SaveOrUpdateDeviceVideo(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}

// DeleteDeviceVideo godoc
// @Summary 删除设备视频视频
// @Description 根据设备视频视频ID删除指定的设备视频视频
// @Accept  json
// @Produce  json
// @Param   id      query     int64  true "设备视频视频ID"
// @Success 200 {object} vo.Response "操作成功!"
// @Failure 400 {object} vo.Response "请求参数错误"
// @Router /admin/deviceVideos [delete]
func DeleteDeviceVideo(c *gin.Context) {
	req := new(dto.IdDeleteDTO)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	if err := service.DeleteDeviceVideo(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("操作成功!", nil))
}
