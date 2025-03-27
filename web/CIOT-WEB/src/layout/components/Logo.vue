<template>
  <div class="logo">
    <img src="@/assets/logo.jpg" alt="logo">
    <h1 v-show="!isCollapse">CIOT 物联网平台</h1>
  </div>
</template>

<script setup>
defineProps({
  isCollapse: {
    type: Boolean,
    default: false
  }
})
</script>

<style lang="scss" scoped>
.logo {
  height: 50px;
  display: flex;
  align-items: center;
  padding: 0 16px;
  color: #fff;
  
  img {
    width: 32px;
    height: 32px;
    margin-right: 12px;
  }
  
  h1 {
    font-size: 16px;
    font-weight: 600;
    margin: 0;
    white-space: nowrap;
  }

    &.is-collapsed {
    padding: 16px;
    
    h1 {
      display: none;
    }
  }
}
</style>