import {
	createRouter,
	createWebHistory
} from 'vue-router'
import Layout from '@/layout/index.vue'

// 静态路由
// export const constantRoutes = [
//     {
//     path: '/',
//     component: Layout,
//     redirect: '/dashboard',
//     children: [
//       {
//         path: 'dashboard',
//         name: 'Dashboard',
//         component: () => import('@/views/dashboard/index.vue'),
//         meta: {
//           icon: 'Monitor',
//           label: '控制台',
//           isHidden: false
//         }
//       }
//     ]
//   },
//   {
//     path: '/iot',
//     component: Layout,
//     redirect: '/iot/product',
//     meta: {
//       icon: 'DataBoard',
//       label: 'IOT管理',
//       isHidden: false
//     },
//     children: [
//       {
//         path: 'product',
//         name: 'Product',
//         component: () => import('@/views/product/index.vue'),
//         meta: {
//           icon: 'Box',
//           label: '产品管理',
//           isHidden: false
//         }
//       },
//       {
//         path: 'product/:id',
//         name: 'ProductDetail',
//         component: () => import('@/views/product/detail.vue'),
//         meta: {
//           title: '产品详情',
//           activeMenu: '/iot/product',
//           isHidden: true
//         }
//       },
//       {
//         path: 'device',
//         name: 'DeviceList',
//         component: () => import('@/views/device/index.vue'),
//         meta: {
//           icon: 'Platform',
//           label: '设备管理',
//           isHidden: false
//         }
//       },
//       {
//         path: 'device/:id',
//         name: 'DeviceDetail',
//         component: () => import('@/views/device/detail.vue'),
//         meta: {
//           title: '设备详情',
//           activeMenu: '/iot/device',
//           isHidden: true
//         }
//       },
//       {
//         path: 'module',
//         name: 'ModuleList',
//         component: () => import('@/views/module/index.vue'),
//         meta: {
//           icon: 'Grid',
//           label: '模块管理',
//           isHidden: false
//         }
//       },
//       {
//         path: 'module/:id',
//         name: 'ModuleDetail',
//         component: () => import('@/views/module/detail.vue'),
//         meta: {
//           title: '模块详情',
//           activeMenu: '/iot/module',
//           isHidden: true
//         }
//       },
//       {
//         path: 'function',
//         name: 'FunctionList',
//         component: () => import('@/views/functions/index.vue'),
//         meta: {
//           icon: 'CollectionTag',
//           label: '功能管理',
//           isHidden: false
//         }
//       },
//       {
//         path: 'function/:id',
//         name: 'FunctionDetail',
//         component: () => import('@/views/functions/detail.vue'),
//         meta: {
//           title: '功能详情',
//           activeMenu: '/iot/function',
//           isHidden: true
//         }
//       },
//       {
//         path: 'deviceAttribute',
//         name: 'deviceAttribute',
//         component: () => import('@/views/deviceatt/index.vue'),
//         meta: {
//           icon: 'Monitor',
//           label: '设备属性',
//           isHidden: false
//         }
//       },
//       {
//         path: 'deviceVideo',
//         name: 'deviceVideo',
//         component: () => import("@/views/devicevideo/index.vue"),
//         meta: {
//           icon: 'Camera',
//           label: '设备视频',
//           isHidden: false
//         }
//       },
//       {
//         path: 'deviceVideo/:id',
//         name: 'video',
//         component: () => import('@/views/devicevideo/video.vue'),
//         meta: {
//           title: '视频',
//           activeMenu: '/iot/devicevideo',
//           isHidden: true
//         }
//       }]
//   },
//   {
//     path: '/system',
//     component: Layout,
//     meta: {
//       icon: 'Setting',
//       label: '系统管理',
//       isHidden: false
//     },
//     children: [
//       {
//         path: 'exception',
//         name: 'Exception',
//         component: () => import('@/views/exceptionlog/index.vue'),
//         meta: {
//           icon: 'Warning',
//           label: '异常日志',
//           isHidden: false
//         }
//       }
//     ]
//   },
//   {
//     path: '/login',
//     name: 'Login',
//     component: () => import('@/views/login/index.vue'),
//     meta: { isHidden: true }
//   }
// ]

export const constantRoutes = [{
		path: '/',
		component: Layout,
		redirect: '/dashboard',
		children: [{
			path: 'dashboard',
			name: 'Dashboard',
			component: () => import('@/views/dashboard/index.vue'),
			meta: {
				icon: 'Monitor',
				label: '控制台',
				isHidden: false
			}
		}]
	},
	{
		path: '/map',
		component: Layout,
		children: [{
			path: 'map',
			name: 'map',
			component: () => import('@/views/map/index.vue'),
			meta: {
				icon: 'Location',
				label: '大地图',
				isHidden: false
			}
		}]
	},
	{
		path: '/iot',
		component: Layout,
		redirect: '/iot/product',
		meta: {
			icon: 'DataBoard',
			label: 'IOT管理',
			isHidden: false
		},
		children: [{
				path: 'product',
				name: 'Product',
				component: () => import('@/views/product/index.vue'),
				meta: {
					icon: 'Box',
					label: '产品管理',
					isHidden: false
				}
			},
			{
				path: 'device',
				name: 'DeviceList',
				component: () => import('@/views/device/index.vue'),
				meta: {
					icon: 'Platform',
					label: '设备管理',
					isHidden: false
				}
			},
			{
				path: 'module',
				name: 'ModuleList',
				component: () => import('@/views/module/index.vue'),
				meta: {
					icon: 'Grid',
					label: '模块管理',
					isHidden: false
				}
			},
			{
				path: 'property',
				name: 'PropertyList',
				component: () => import('@/views/property/index.vue'),
				meta: {
					icon: 'CollectionTag',
					label: '属性管理',
					isHidden: false
				}
			},
			{
				path: 'deviceProperty',
				name: 'deviceProperty',
				component: () => import('@/views/deviceProperty/index.vue'),
				meta: {
					icon: 'Monitor',
					label: '设备属性',
					isHidden: false
				}
			},
			{
				path: 'deviceVideo',
				name: 'deviceVideo',
				component: () => import("@/views/deviceVideo/index.vue"),
				meta: {
					icon: 'Camera',
					label: '设备视频',
					isHidden: false
				}
			},
			{
				path: 'deviceVideo/:id',
				name: 'video',
				component: () => import('@/views/devicevideo/video.vue'),
				meta: {
					title: '视频',
					activeMenu: '/iot/devicevideo',
					isHidden: true
				}
			}
		]
	},
	{
		path: '/login',
		name: 'Login',
		component: () => import('@/views/login/index.vue'),
		meta: {
			isHidden: true
		}
	}
]
const router = createRouter({
	history: createWebHistory(),
	routes: constantRoutes
})

export default router