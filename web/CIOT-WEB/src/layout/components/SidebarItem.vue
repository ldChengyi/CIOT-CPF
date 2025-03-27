<template>
  <!-- 不显示的路由直接跳过 -->
  <template v-if="!route.meta?.isHidden">
    <!-- 没有子路由或只有一个可见的子路由 -->
    <el-menu-item 
      v-if="!hasVisibleChildren(route)"
      :index="resolvePath(route)"
    >
      <el-icon v-if="getIcon(route)">
        <component :is="getIcon(route)"></component>
      </el-icon>
      <template #title>{{ getTitle(route) }}</template>
    </el-menu-item>

    <!-- 有多个可见的子路由 -->
    <el-sub-menu 
      v-else 
      :index="route.path"
    >
      <template #title>
        <el-icon v-if="getIcon(route)">
          <component :is="getIcon(route)"></component>
        </el-icon>
        <span>{{ getTitle(route) }}</span>
      </template>

      <!-- 递归调用自身，处理子路由 -->
      <sidebar-item
        v-for="child in getVisibleChildren(route)"
        :key="child.path"
        :route="child"
        :base-path="resolvePath(route)"
      />
    </el-sub-menu>
  </template>
</template>

<script setup>
import { computed } from 'vue'
import path from 'path-browserify'

const props = defineProps({
  route: {
    type: Object,
    required: true
  },
  basePath: {
    type: String,
    default: ''
  }
})

// 获取可见的子路由
const getVisibleChildren = (route) => {
  if (!route.children) return []
  return route.children.filter(child => !child.meta?.isHidden)
}

// 判断是否有可见的子路由
const hasVisibleChildren = (route) => {
  const visibleChildren = getVisibleChildren(route)
  return visibleChildren.length > 1
}

// 获取图标
const getIcon = (route) => {
  const visibleChildren = getVisibleChildren(route)
  if (visibleChildren.length === 1) {
    return visibleChildren[0].meta?.icon || route.meta?.icon
  }
  return route.meta?.icon
}

// 获取标题
const getTitle = (route) => {
  const visibleChildren = getVisibleChildren(route)
  if (visibleChildren.length === 1) {
    return visibleChildren[0].meta?.label || visibleChildren[0].name || route.meta?.label || route.name
  }
  return route.meta?.label || route.name
}

// 解析路由路径
const resolvePath = (route) => {
  const visibleChildren = getVisibleChildren(route)
  if (visibleChildren.length === 1) {
    return path.resolve(props.basePath || '', route.path, visibleChildren[0].path)
  }
  return path.resolve(props.basePath || '', route.path)
}
</script>

<style lang="scss" scoped>
.el-menu-item, .el-sub-menu__title {
  display: flex;
  align-items: center;
  
  .el-icon {
    margin-right: 12px;
  }
}
</style>