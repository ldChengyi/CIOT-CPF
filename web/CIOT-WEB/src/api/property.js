import request from '@/utils/request'

// 获取产品列表
export function getProperties(params) {
	return request({
		url: '/admin/properties',
		method: 'get',
		params
	})
}

// 新增或更新产品
export function saveOrUpdateProperty(data) {
	return request({
		url: '/admin/properties',
		method: 'post',
		data
	})
}

// 删除产品
export function deleteProperty(id) {
	return request({
		url: `/admin/properties`,
		method: 'delete',
		data: {
			"id": id
		}
	})
}