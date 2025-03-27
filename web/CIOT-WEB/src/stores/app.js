import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    sidebar: {
      isCollapse: false
    }
  }),
  
  actions: {
    toggleSidebar() {
      this.sidebar.isCollapse = !this.sidebar.isCollapse
    }
  }
})