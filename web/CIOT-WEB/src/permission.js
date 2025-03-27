import router from './router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

// 白名单路由
const whiteList = ['/login','/register']

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  
  if (userStore.token) {
    if (to.path === '/login') {
      // 已登录且要跳转的页面是登录页
      next({ path: '/' })
    } else {
      // 已登录访问其他页面
      next()
    }
  } else {
    // 未登录
    if (whiteList.indexOf(to.path) !== -1) {
      // 在免登录白名单，直接进入
      next()
    } else {
      // 其他没有访问权限的页面将被重定向到登录页面
      next(`/login?redirect=${to.path}`)
      ElMessage.warning('请先登录')
    }
  }
})