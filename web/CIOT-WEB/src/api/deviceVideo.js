import request from '@/utils/request'

export function getDeviceVideos(params) {
	return request({
		url: '/admin/deviceVideos',
		method: 'get',
		params
	})
}


export function getDeviceVideoDetail(id) {
	return request({
		url: `/admin/deviceVideos/${id}`,
		method: 'get',
	})
}


// 新增或更新设备
export function saveOrUpdateDeviceVideo(data) {
	return request({
		url: '/admin/deviceVideos',
		method: 'post',
		data
	})
}

// 删除设备
export function deleteDeviceVideo(id) {
	return request({
		url: `/admin/deviceVideos`,
		method: 'delete',
		data: {
			"id": id
		}
	})
}