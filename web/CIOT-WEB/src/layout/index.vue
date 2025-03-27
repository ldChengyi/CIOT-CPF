<template>
  <div class="app-wrapper">
    <!-- 使用store中的状态控制侧边栏 -->
    <div class="sidebar-container" :class="{ 'is-collapsed': appStore.sidebar.isCollapse }">
      <Logo :is-collapse="appStore.sidebar.isCollapse" />
      <Sidebar :is-collapse="appStore.sidebar.isCollapse" />
    </div>
    
    <div class="main-container">
      <Header />
      <div class="app-main">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAppStore } from '@/stores/app'
import Sidebar from './components/Siderbar.vue'
import Header from './components/Header.vue'
import Logo from './components/Logo.vue'

const appStore = useAppStore()
</script>

<style lang="scss" scoped>
.app-wrapper {
  display: flex;
  width: 100%;
  height: 100vh;
}

.sidebar-container {
  width: 220px;
  height: 100%;
  background: #001529;
  transition: width 0.3s;
  
  &.is-collapsed {
    width: 64px;
  }
}

.main-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background: #f0f2f5;
}

.app-main {
  padding: 20px;
  flex: 1;
  overflow: auto;
}
</style>