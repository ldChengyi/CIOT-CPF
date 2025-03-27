import request from '@/utils/request'

// 获取模块列表
export function getModules(params) {
	return request({
		url: '/admin/modules',
		method: 'get',
		params
	})
}

// 新增或更新模块
export function saveOrUpdateModule(data) {
	return request({
		url: '/admin/modules',
		method: 'post',
		data
	})
}

// 删除模块
export function deleteModule(id) {
	return request({
		url: `/admin/modules`,
		method: 'delete',
		data: {
			"id": id
		}
	})
}