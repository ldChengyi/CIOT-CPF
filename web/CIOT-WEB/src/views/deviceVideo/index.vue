<template>
	<div class="device-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="关键词">
					<el-input v-model="searchForm.filters.keywords" placeholder="视频流地址/流类型" clearable @keyup.enter="handleSearch">
						<template #prefix>
							<el-icon><Search /></el-icon>
						</template>
					</el-input>
				</el-form-item>

				<el-form-item class="search-buttons">
					<el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button>
					<el-button :icon="Refresh" @click="handleReset">重置</el-button>
					<el-button type="primary" :icon="Plus" @click="handleAdd">新增</el-button>
				</el-form-item>
			</el-form>
		</div>

		<!-- 表格卡片 -->
		<div class="table-card">
			<el-table v-loading="loading" :data="tableData" style="width: 100%">
				<el-table-column prop="device_id" label="设备ID" min-width="auto" />
				<el-table-column prop="device_name" label="设备名称" min-width="auto" />
				<el-table-column label="启用视频" width="auto">
					<template #default="{ row }">
						<el-tag :type="row.enable_video ? 'success' : 'danger'">
							{{ row.enable_video ? '启用' : '禁用' }}
						</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="video_url" label="视频流地址" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="stream_type" label="视频流类型" min-width="auto" />
				<el-table-column label="状态" width="auto">
					<template #default="{ row }">
						<el-tag :type="row.status === 1 ? 'success' : 'danger'">
							{{ row.status === 1 ? '在线' : '离线' }}
						</el-tag>
					</template>
				</el-table-column>
				<el-table-column prop="created_at" label="创建时间" width="auto">
					<template #default="{ row }">
						<span class="time">{{ formatDateTime(row.created_at) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="操作" width="200" fixed="right">
					<template #default="{ row }">
						<el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
						<el-button link type="primary" @click="CheckVideo(row)">查看视频</el-button>
						<el-button link type="danger" @click="handleDelete(row)">删除</el-button>
					</template>
				</el-table-column>
			</el-table>

			<!-- 分页 -->
			<div class="pagination">
				<el-pagination
					v-model:current-page="searchForm.page"
					v-model:page-size="searchForm.page_size"
					:total="total"
					:page-sizes="[10, 20, 50, 100]"
					layout="total, sizes, prev, pager, next, jumper"
					@size-change="handleSizeChange"
					@current-change="handleCurrentChange"
				/>
			</div>
		</div>

		<!-- 新增/编辑对话框 -->
		<el-dialog :title="dialogTitle" v-model="dialogVisible" width="500px">
			<el-form ref="formRef" :model="formData" :rules="formRules" label-width="120px">
				<el-form-item label="关联设备" prop="device_id">
					<el-select v-model="formData.device_id" style="width: 100%">
						<el-option v-for="item in deviceOptions" :key="item.id" :label="item.device_name" :value="item.id" />
					</el-select>
				</el-form-item>
				<el-form-item label="启用视频" prop="enable_video">
					<el-radio-group v-model="formData.enable_video">
						<el-radio :label="true">启用</el-radio>
						<el-radio :label="false">禁用</el-radio>
					</el-radio-group>
				</el-form-item>
				<el-form-item label="视频流地址" prop="video_url">
					<el-input v-model="formData.video_url" />
				</el-form-item>
				<el-form-item label="视频流类型" prop="stream_type">
					<el-input v-model="formData.stream_type" />
				</el-form-item>
				<el-form-item label="状态" prop="status">
					<el-radio-group v-model="formData.status">
						<el-radio :label="1">在线</el-radio>
						<el-radio :label="0">离线</el-radio>
					</el-radio-group>
				</el-form-item>
			</el-form>
			<template #footer>
				<el-button @click="dialogVisible = false">取消</el-button>
				<el-button type="primary" @click="handleSubmit">确定</el-button>
			</template>
		</el-dialog>
	</div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Search, Refresh, Plus } from '@element-plus/icons-vue';
import { ElMessage, ElMessageBox } from 'element-plus';
// 替换为DeviceVideo相关API
import { getDeviceVideos, saveOrUpdateDeviceVideo, deleteDeviceVideo } from '@/api/deviceVideo';
import { getDevices } from '../../api/device';
import { formatDateTime } from '@/utils/format.js';
import { useRouter } from 'vue-router';

// 路由
const router = useRouter();

// 表格数据
const tableData = ref([]);
const loading = ref(false);
const total = ref(0);

// 搜索表单
const searchForm = reactive({
	page: 1,
	page_size: 10,
	filters: {
		keywords: ''
	}
});

// 状态选项
const statusOptions = [
	{ label: '在线', value: 1 },
	{ label: '离线', value: 0 }
];
const deviceOptions = ref([]);

// 对话框控制
const dialogVisible = ref(false);
const dialogTitle = ref('');
const formRef = ref(null);

// 表单数据
const formData = reactive({
	device_id: '',
	enable_video: false,
	video_url: '',
	stream_type: '',
	status: 1
});

// 表单验证规则
const formRules = {
	device_id: [{ required: true, message: '请输入设备ID', trigger: 'blur' }],
	video_url: [{ required: true, message: '请输入视频流地址', trigger: 'blur' }],
	stream_type: [{ required: true, message: '请输入视频流类型', trigger: 'blur' }]
};

// 获取设备列表
const fetchDevices = async () => {
	try {
		const res = await getDevices({ size: 999 });
		deviceOptions.value = res.data.list;
	} catch (error) {
		console.error('获取产品列表失败:', error);
	}
};

// 获取设备视频列表
const fetchData = async () => {
	loading.value = true;
	try {
		const params = {
			...searchForm
		};
		const res = await getDeviceVideos(params);
		tableData.value = res.data.list;
		total.value = res.data.total;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};

// 重置表单
const resetForm = () => {
	formData.device_id = '';
	formData.enable_video = false;
	formData.video_url = '';
	formData.stream_type = '';
	formData.status = 1;
};

// 处理搜索
const handleSearch = () => {
	searchForm.current = 1;
	fetchData();
};

// 处理重置
const handleReset = () => {
	searchForm.filter.keywords = '';
	searchForm.page_size = 10;
	searchForm.page = 1;
	handleSearch();
};

// 处理新增
const handleAdd = () => {
	resetForm();
	dialogTitle.value = '新增视频设备';
	dialogVisible.value = true;
};

// 处理编辑
const handleEdit = (row) => {
	Object.assign(formData, row);
	// 将 enable_video 转换为 boolean 类型
	formData.enable_video = Boolean(row.enable_video);
	dialogTitle.value = '编辑视频设备';
	dialogVisible.value = true;
};

// 处理详情
const CheckVideo = (row) => {
	router.push(`/iot/deviceVideo/${row.id}`);
};

// 处理删除
const handleDelete = (row) => {
	ElMessageBox.confirm('确认删除该视频设备?', '提示', {
		type: 'warning'
	}).then(async () => {
		try {
			await deleteDeviceVideo(row.id);
			ElMessage.success('删除成功');
			fetchData();
		} catch (error) {
			ElMessage.error('删除失败');
		}
	});
};

// 处理提交
const handleSubmit = async () => {
	if (!formRef.value) return;

	try {
		await formRef.value.validate();
		await saveOrUpdateDeviceVideo(formData);
		ElMessage.success(formData.id ? '更新成功' : '添加成功');
		dialogVisible.value = false;
		fetchData();
	} catch (error) {
		console.error('提交失败:', error);
	}
};

// 分页处理
const handleSizeChange = (val) => {
	searchForm.page_size = val;
	fetchData();
};

const handleCurrentChange = (val) => {
	searchForm.page = val;
	fetchData();
};

// 初始化
onMounted(() => {
	fetchData();
	fetchDevices();
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
				margin-left: auto;
				display: flex;
				gap: 8px;

				.el-button {
					padding: 8px 16px;

					.el-icon {
						margin-right: 4px;
					}
				}
			}
		}
	}

	.table-card {
		background: #fff;
		border-radius: 4px;
		padding: 20px;
		box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

		.pagination {
			margin-top: 20px;
			display: flex;
			justify-content: flex-end;
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
