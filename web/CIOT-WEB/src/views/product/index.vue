<template>
	<div class="product-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="关键词">
					<el-input v-model="searchForm.filters.keywords" placeholder="产品名称/ProductKey" clearable @keyup.enter="handleSearch">
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
				<el-table-column prop="id" label="ID" min-width="auto" />
				<el-table-column prop="product_key" label="ProductKey" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="product_name" label="产品名称" min-width="auto" />
				<el-table-column prop="desc" label="描述" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="created_at" label="创建时间" width="auto">
					<template #default="{ row }">
						<span class="time">{{ formatDateTime(row.created_at) }}</span>
					</template>
				</el-table-column>
				<el-table-column label="操作" width="200" fixed="right">
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
				<el-form-item label="产品密钥" prop="product_key">
					<el-input v-model="formData.product_key" />
				</el-form-item>
				<el-form-item label="产品名称" prop="product_name">
					<el-input v-model="formData.product_name" />
				</el-form-item>
				<el-form-item label="产品密码" prop="product_secret">
					<el-input v-model="formData.product_secret" />
				</el-form-item>
				<el-form-item label="产品描述" prop="description">
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
import { getProducts, saveOrUpdateProduct, deleteProduct } from '@/api/product';
import { useRouter } from 'vue-router';
import { formatDateTime } from '@/utils/format';

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
	{ label: '启用', value: 1 },
	{ label: '禁用', value: 0 }
];

// 对话框控制
const dialogVisible = ref(false);
const dialogTitle = ref('');
const formRef = ref(null);

// 表单数据
const formData = reactive({
	id: null,
	product_key: '',
	product_name: '',
	product_secret: '',
	desc: ''
});

// 表单验证规则
const formRules = {
	product_name: [
		{ required: true, message: '请输入产品名称', trigger: 'blur' },
		{ min: 3, max: 50, message: '长度在 3 到 50 个字符', trigger: 'blur' }
	],
	product_key: [
		{ required: true, message: '请输入产品密钥', trigger: 'blur' },
		{ min: 3, max: 50, message: '长度在 3 到 50 个字符', trigger: 'blur' }
	],
	product_secret: [
		{ message: '请输入产品密码', trigger: 'blur' },
		{ min: 6, max: 50, message: '长度在 3 到 50 个字符', trigger: 'blur' }
	]
};

// 获取数据
const fetchData = async () => {
	loading.value = true;
	try {
		const params = {
			...searchForm
		};
		const res = await getProducts(params);
		tableData.value = res.data.list;
		total.value = res.data.total;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};

// 搜索
const handleSearch = () => {
	searchForm.page = 1;
	fetchData();
};

// 重置
const handleReset = () => {
	searchForm.filters.keywords = '';
	searchForm.page = 1;
	handleSearch();
};

// 新增
const handleAdd = () => {
	resetForm();
	dialogTitle.value = '新增产品';
	dialogVisible.value = true;
};

// 重置表单
const resetForm = () => {
	formData.id = null;
	formData.product_name = '';
	formData.product_key = '';
	formData.product_secret = '';
	formData.desc = '';
};

// 编辑
const handleEdit = (row) => {
	Object.assign(formData, row);
	dialogTitle.value = '编辑产品';
	dialogVisible.value = true;
};

// 删除
const handleDelete = (row) => {
	ElMessageBox.confirm('确认删除该产品?', '提示', {
		type: 'warning'
	}).then(async () => {
		try {
			await deleteProduct(row.id);
			ElMessage.success('删除成功');
			fetchData();
		} catch (error) {
			ElMessage.error('删除失败');
		}
	});
};

// 提交表单
const handleSubmit = async () => {
	if (!formRef.value) return;

	try {
		await formRef.value.validate();
		await saveOrUpdateProduct(formData);
		ElMessage.success(formData.id ? '更新成功' : '新增成功');
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
});
</script>

<style lang="scss" scoped>
.product-container {
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
	.product-container {
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
