import request from '@/utils/request'

// 获取产品列表
export function getProducts(params) {
  return request({
    url: '/admin/products',
    method: 'get',
    params
  })
}

// 新增或更新产品
export function saveOrUpdateProduct(data) {
  return request({
    url: '/admin/products',
    method: 'post',
    data
  })
}

// 删除产品
export function deleteProduct(id) {
  return request({
    url: `/admin/products`,
    method: 'delete',
    data: {
      "id" : id
    }
  })
}