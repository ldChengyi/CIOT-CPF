<template>
	<div class="device-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="设备">
					<el-select v-model="searchForm.id" placeholder="请选择设备" clearable>
						<el-option v-for="item in deviceOptions" :key="item.id" :label="item.device_name" :value="item.id" />
					</el-select>
				</el-form-item>

				<el-form-item class="search-buttons">
					<el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button>
				</el-form-item>

				<el-form-item class="refresh-switch">
					<el-switch v-model="autoRefreshEnabled" active-text="开启实时刷新" inactive-text="关闭实时刷新"></el-switch>
				</el-form-item>
			</el-form>
		</div>

		<!-- 表格卡片 -->
		<div class="table-card">
			<el-row :gutter="20">
				<el-col :span="8" v-for="property in propertiesData" :key="property.id">
					<el-card class="property-card" shadow="hover">
						<div class="header">
							<h3>{{ property.name }}</h3>
						</div>
						<div class="content">
							<p>
								<strong>Identifier:</strong>
								<el-tag effect="dark" size="large" style="font-size: 14px; font-weight: 700">{{ property.identifier }}</el-tag>
							</p>
							<p>
								<strong>Value:</strong>
								<el-text class="mx-1" type="primary" size="large" style="font-weight: 700" v-if="property.value !== ''">{{ property.value }}</el-text>
								<span v-else class="empty-value">暂无数据</span>
							</p>
						</div>
					</el-card>
				</el-col>
			</el-row>

			<el-empty description="请选择设备" />
		</div>
	</div>
</template>

<script setup>
import { ref, reactive, onMounted, watch, onBeforeUnmount } from 'vue';
import { Search, Refresh, Plus } from '@element-plus/icons-vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { getDevices, getDevicePropertiesVal } from '@/api/device';
import { formatDateTime } from '@/utils/format.js';
import { useRouter } from 'vue-router';

//路由
const router = useRouter();

//加载
const loading = ref(false);

//自动刷新定时器
const autoRefreshEnabled = ref(false);
const refreshInterval = 1000;
let intervalId = null;

//获得数据
const deviceOptions = ref([]);
const propertiesData = ref([]);

//请求参数
const searchForm = reactive({
	id: null
});

// 获取设备列表
const fetchDevices = async () => {
	loading.value = true;
	try {
		const res = await getDevices({ page_size: 999 });
		deviceOptions.value = res.data.list;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};

const fetchData = async () => {
	loading.value = true;
	try {
		const res = await getDevicePropertiesVal({ ...searchForm });
		propertiesData.value = res.data.list;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};

const handleSearch = () => {
	fetchData();
};

// 监听开关状态变化
watch(autoRefreshEnabled, (newVal) => {
	if (newVal) {
		// 开关打开时启动定时刷新
		intervalId = setInterval(fetchData, refreshInterval);
		// 可选：初始加载一次数据
		fetchData();
	} else {
		// 关闭时清除定时器
		if (intervalId) {
			clearInterval(intervalId);
			intervalId = null;
		}
	}
});

// 组件卸载时清除定时器

// 初始化
onMounted(() => {
	fetchDevices();
});

onBeforeUnmount(() => {
	if (intervalId) {
		clearInterval(intervalId);
	}
});
</script>

<style lang="scss" scoped>
.device-container {
	padding: 20px;

	.search-card {
		background: #fff;
		border-radius: 4px;
		box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
		margin-bottom: 20px;
		padding: 24px;

		.search-form {
			display: flex;
			flex-wrap: wrap;
			gap: 16px;

			:deep(.el-form-item) {
				margin-bottom: 0;
				margin-right: 0;

				.el-input {
					width: 240px;

					.el-input__wrapper {
						box-shadow: 0 0 0 1px #dcdfe6 inset;

						&:hover {
							box-shadow: 0 0 0 1px #c0c4cc inset;
						}

						&.is-focus {
							box-shadow: 0 0 0 1px var(--el-color-primary) inset;
						}
					}
				}

				.el-select {
					width: 180px;
				}
			}

			.search-buttons {
				display: flex;
				gap: 8px;

				.el-button {
					padding: 8px 16px;

					.el-icon {
						margin-right: 4px;
					}
				}
			}

			.refresh-switch {
				display: flex;
				margin-left: auto;
				margin-right: 100px;
			}
		}
	}
}
.table-card {
	padding: 20px;
	background-color: #f9f9f9;

	.el-row {
		margin-bottom: 20px;
	}

	.property-card {
		border-radius: 8px;
		overflow: hidden;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
		transition: transform 0.3s ease, box-shadow 0.3s ease;
		margin-bottom: 20px;

		&:hover {
			transform: translateY(-5px);
			box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
		}

		.header {
			background-color: #409eff;
			padding: 10px;

			h3 {
				color: #fff;
				margin: 0;
				font-size: 18px;
			}
		}

		.content {
			padding: 15px;

			p {
				margin: 8px 0;
				font-size: 14px;
				color: #333;

				strong {
					display: inline-block;
					width: 90px;
					color: #555;
				}

				.empty-value {
					color: #999;
					font-style: italic;
				}
			}
		}
	}
}

// 响应式布局
@media screen and (max-width: 768px) {
	.device-container {
		.search-card {
			.search-form {
				:deep(.el-form-item) {
					margin-bottom: 16px;

					.el-input,
					.el-select {
						width: 100%;
					}
				}

				.search-buttons {
					width: 100%;
					justify-content: flex-end;
				}
			}
		}
	}
}
</style>
