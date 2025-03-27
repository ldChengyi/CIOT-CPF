import { defineStore } from 'pinia'
import { login, register } from '@/api/auth'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: {}
  }),

  actions: {
    // 登录
    async login(userInfo) {
      try {
        const res = await login(userInfo)
        const { token } = res.data
        this.token = token
        localStorage.setItem('token', token)
        ElMessage.success('登录成功')
        return Promise.resolve(res)
      } catch (error) {
        ElMessage.error(error.message || '登录失败')
        return Promise.reject(error)
      }
    },

    // 注册
    async register(userInfo) {
      try {
        await register(userInfo)
        ElMessage.success('注册成功，请登录')
        return Promise.resolve()
      } catch (error) {
        ElMessage.error(error.message || '注册失败')
        return Promise.reject(error)
      }
    },

    // 登出
    logout() {
      this.token = ''
      this.userInfo = {}
      localStorage.removeItem('token')
    },

    // 重置token
    resetToken() {
      this.token = ''
      localStorage.removeItem('token')
    }
  }
})