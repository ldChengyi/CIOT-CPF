<template>
	<div class="device-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="关键词">
					<el-input v-model="searchForm.filters.keywords" placeholder="设备名称/设备密钥" clearable @keyup.enter="handleSearch">
						<template #prefix>
							<el-icon><Search /></el-icon>
						</template>
					</el-input>
				</el-form-item>

				<el-form-item label="产品">
					<el-select v-model="searchForm.filters.product_id" placeholder="请选择产品" clearable>
						<el-option v-for="item in productOptions" :key="item.id" :label="item.product_name" :value="item.id" />
					</el-select>
				</el-form-item>

				<el-form-item label="模块">
					<el-select v-model="searchForm.filters.module_id" placeholder="请选择模块" clearable>
						<el-option v-for="item in moduleOptions" :key="item.id" :label="item.module_name" :value="item.id" />
					</el-select>
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
				<el-table-column prop="device_name" label="设备名称" min-width="auto" />
				<el-table-column prop="device_secret" label="设备密码" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="product_id" label="所属产品" min-width="auto">
					<template #default="{ row }">
						{{ getProductName(row.product_id) }}
					</template>
				</el-table-column>
				<el-table-column prop="module_id" label="关联模块" min-width="auto">
					<template #default="{ row }">
						{{ getModuleName(row.module_id) }}
					</template>
				</el-table-column>
				<el-table-column label="状态" width="auto">
					<template #default="{ row }">
						<el-tag :type="row.status === 1 ? 'success' : 'danger'">
							{{ row.status === 1 ? '在线' : '离线' }}
						</el-tag>
					</template>
				</el-table-column>

				<el-table-column prop="last_online" label="最后上线时间" width="auto">
					<template #default="{ row }">
						<span class="time">{{ formatDateTime(row.last_online) }}</span>
					</template>
				</el-table-column>
				<el-table-column prop="created_at" label="创建时间" width="auto">
					<template #default="{ row }">
						<span class="time">{{ formatDateTime(row.created_at) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="操作" width="auto" fixed="right">
					<template #default="{ row }">
						<el-button link type="primary" @click="handleEdit(row)">编辑</el-button>
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
			<el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
				<el-form-item label="设备名称" prop="device_name">
					<el-input v-model="formData.device_name" />
				</el-form-item>
				<el-form-item label="所属产品" prop="product_id">
					<el-select v-model="formData.product_id" style="width: 100%">
						<el-option v-for="item in productOptions" :key="item.id" :label="item.product_name" :value="item.id" />
					</el-select>
				</el-form-item>
				<el-form-item label="关联模块" prop="module_id">
					<el-select v-model="formData.module_id" style="width: 100%">
						<el-option v-for="item in moduleOptions" :key="item.id" :label="item.module_name" :value="item.id" />
					</el-select>
				</el-form-item>
				<el-form-item label="描述" prop="desc">
					<el-input v-model="formData.desc" type="textarea" rows="3" />
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
import { getDevices, saveOrUpdateDevice, deleteDevice } from '@/api/device';
import { getProducts } from '@/api/product';
import { getModules } from '@/api/module';
import { formatDateTime } from '@/utils/format.js';
import { useRouter } from 'vue-router';

//路由
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
		keywords: '',
		module_id: null,
		product_id: null
	}
});

// 状态选项
const statusOptions = [
	{ label: '在线', value: 1 },
	{ label: '离线', value: 0 }
];

// 产品选项
const productOptions = ref([]);

// 模块选项
const moduleOptions = ref([]);

// 对话框控制
const dialogVisible = ref(false);
const dialogTitle = ref('');
const formRef = ref(null);

// 表单数据
const formData = reactive({
	device_name: '',
	device_secret: '',
	product_id: null,
	module_id: null,
	desc: ''
});

// 表单验证规则
const formRules = {
	device_name: [
		{ required: true, message: '请输入设备名称', trigger: 'blur' },
		{ min: 3, max: 50, message: '长度在 3 到 50 个字符', trigger: 'blur' }
	],
	product_id: [{ required: true, message: '请选择所属产品', trigger: 'change' }],
	module_id: [{ required: true, message: '请选择关联模块', trigger: 'change' }]
};

// 获取产品列表
const fetchProducts = async () => {
	try {
		const res = await getProducts({ size: 999 });
		productOptions.value = res.data.list;
	} catch (error) {
		console.error('获取产品列表失败:', error);
	}
};

// 获取模块列表
const fetchModules = async () => {
	try {
		const res = await getModules({ size: 999 });
		moduleOptions.value = res.data.list;
	} catch (error) {
		console.error('获取模块列表失败:', error);
	}
};

// 获取设备列表
const fetchData = async () => {
	loading.value = true;
	try {
		const params = {
			...searchForm
		};
		const res = await getDevices(params);
		tableData.value = res.data.list;
		total.value = res.data.total;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};

// 重置表单
const resetForm = () => {
	formData.device_name = '';
	formData.device_secret = '';
	formData.product_id = null;
	formData.module_id = null;
	formData.desc = '';
};

// 获取产品名称
const getProductName = (productId) => {
	const product = productOptions.value.find((item) => item.id === productId);
	return product ? product.product_name : '-';
};

// 获取模块名称
const getModuleName = (moduleId) => {
	const module = moduleOptions.value.find((item) => item.id === moduleId);
	return module ? module.module_name : '-';
};

// 处理搜索
const handleSearch = () => {
	searchForm.page = 1;
	fetchData();
};

// 处理重置
const handleReset = () => {
	searchForm.keywords = '';
	searchForm.filters.product_id = undefined;
	searchForm.filters.module_id = undefined;
	searchForm.page = 1;
	handleSearch();
};

// 处理新增
const handleAdd = () => {
	resetForm();
	dialogTitle.value = '新增设备';
	dialogVisible.value = true;
};

// 处理编辑
const handleEdit = (row) => {
	Object.assign(formData, row);
	dialogTitle.value = '编辑设备';
	dialogVisible.value = true;
};

// 处理删除
const handleDelete = (row) => {
	ElMessageBox.confirm('确认删除该设备?', '提示', {
		type: 'warning'
	}).then(async () => {
		try {
			await deleteDevice(row.id);
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
		await saveOrUpdateDevice(formData);
		ElMessage.success(formData.id ? '更新成功' : '添加成功');
		dialogVisible.value = false;
		fetchData();
	} catch (error) {
		console.error('提交失败:', error);
	}
};

// 分页处理
const handleSizeChange = (val) => {
	searchForm.page = val;
	fetchData();
};

const handleCurrentChange = (val) => {
	searchForm.page_size = val;
	fetchData();
};

// 初始化
onMounted(() => {
	fetchProducts();
	fetchModules();
	fetchData();
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
