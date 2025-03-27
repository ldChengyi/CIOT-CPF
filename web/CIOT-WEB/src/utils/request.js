import axios from 'axios'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import config from '@/config'

const service = axios.create({
  baseURL: config.baseURL,
  timeout: config.timeout
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers['Authorization'] = `Bearer ${userStore.token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
// 响应拦截器
service.interceptors.response.use(
  // 第一个函数处理 HTTP 状态码 2xx 的响应
  response => {
    const res = response.data
    // 这里处理业务状态码，比如后端返回 { code: 401, msg: '未登录' }
    if (res.code !== 200) {
      // 虽然 HTTP 状态是 200，但业务状态是 401
      if (res.code === 401) {
        ElMessage.error('登录已过期，请重新登录')
        const userStore = useUserStore()
        userStore.resetToken()
        router.push(`/login?redirect=${router.currentRoute.value.fullPath}`)
        return Promise.reject(new Error('登录已过期'))
      }
      
      ElMessage.error(res.msg || '请求失败')
      return Promise.reject(new Error(res.msg || '请求失败'))
    }
    return res
  },

  // 第二个函数处理 HTTP 错误状态
  error => {
    // 处理 HTTP 401 状态码
    // 这种情况是服务器直接返回 401 状态码
    if (error.response && error.response.status === 401) {
      ElMessage.error('登录已过期，请重新登录')
      const userStore = useUserStore()
      userStore.resetToken()
      router.push(`/login?redirect=${router.currentRoute.value.fullPath}`)
      return Promise.reject(error)
    }
    
    // 处理其他错误情况：
    // - HTTP 4xx/5xx 错误
    // - 请求超时
    // - 网络错误等
    ElMessage.error(error.response.data.error || '请求失败')
    return Promise.reject(error)
  }
)

export default service