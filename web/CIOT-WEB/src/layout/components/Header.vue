<template>
  <div class="header">
    <div class="left">
      <el-button type="text" @click="toggleSidebar">
        <el-icon>
          <component :is="appStore.sidebar.isCollapse ? 'Expand' : 'Fold'" />
        </el-icon>
      </el-button>
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>设备管理</el-breadcrumb-item>
        <el-breadcrumb-item>设备分组</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    
    <div class="right">
      <el-tooltip content="搜索" placement="bottom">
        <el-button type="text">
          <el-icon><Search /></el-icon>
        </el-button>
      </el-tooltip>
      
      <el-tooltip content="全屏" placement="bottom">
        <el-button type="text">
          <el-icon><FullScreen /></el-icon>
        </el-button>
      </el-tooltip>
      
      <el-dropdown @command="handleCommand">
        <span class="user-info">
          <el-avatar size="small" />
          <span>管理员</span>
          <el-icon class="el-icon--right">
            <CaretBottom />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon><User /></el-icon>个人信息
            </el-dropdown-item>
            <el-dropdown-item divided command="logout">
              <el-icon><SwitchButton /></el-icon>退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { 
  Fold, 
  Expand, 
  Search, 
  FullScreen, 
  CaretBottom,
  User,
  SwitchButton
} from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { useAppStore } from '@/stores/app'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const appStore = useAppStore()
const userStore = useUserStore()

const toggleSidebar = () => {
  appStore.toggleSidebar()
}

// 处理下拉菜单命令
const handleCommand = (command) => {
  switch (command) {
    case 'profile':
      // 跳转到个人信息页面
      router.push('/profile')
      break
    case 'logout':
      handleLogout()
      break
  }
}

// 处理登出
const handleLogout = () => {
  ElMessageBox.confirm(
    '确定要退出登录吗？',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(() => {
    userStore.logout()
    router.push('/login')
  })
}
</script>

<style lang="scss" scoped>
.header {
  height: 50px;
  padding: 0 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
  
  .left {
    display: flex;
    align-items: center;
    gap: 20px;
  }
  
  .right {
    display: flex;
    align-items: center;
    gap: 16px;
  }
  
  .user-info {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    
    .el-icon {
      font-size: 12px;
    }
  }
}

:deep(.el-dropdown-menu__item) {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 20px;
  
  .el-icon {
    margin-right: 4px;
  }
}
</style>