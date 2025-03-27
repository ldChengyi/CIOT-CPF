<template>
  <el-menu
    :default-active="activeMenu"
    background-color="#001529"
    text-color="#fff"
    active-text-color="#409EFF"
    :router="true"
    :collapse="isCollapse"
    :collapse-transition="false"
  >
    <sidebar-item 
      v-for="route in routes" 
      :key="route.path" 
      :route="route" 
    />
  </el-menu>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { constantRoutes} from '@/router'
import SidebarItem from './SidebarItem.vue'

defineProps({
  isCollapse: {
    type: Boolean,
    default: false
  }
})

const route = useRoute()

// 合并路由
const routes = computed(() => {
  return [...constantRoutes]
})

// 当前激活的菜单
const activeMenu = computed(() => route.path)
</script>

<style lang="scss" scoped>
.el-menu {
  border-right: none;
  user-select: none;
}
</style>