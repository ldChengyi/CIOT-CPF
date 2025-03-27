import request from '@/utils/request'

// 获取设备列表
export function getDevices(params) {
	return request({
		url: '/admin/devices',
		method: 'get',
		params
	})
}

// 新增或更新设备
export function saveOrUpdateDevice(data) {
	return request({
		url: '/admin/devices',
		method: 'post',
		data
	})
}

// 删除设备
export function deleteDevice(id) {
	return request({
		url: `/admin/devices`,
		method: 'delete',
		data: {
			"id": id
		}
	})
}

export function getDevicePropertiesVal(params) {
	return request({
		url: '/admin/devices/properties/val',
		method: 'get',
		params
	})
}