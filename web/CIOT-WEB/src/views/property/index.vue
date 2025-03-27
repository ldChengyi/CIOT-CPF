<template>
	<div class="property-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="关键词">
					<el-input v-model="searchForm.filters.keywords" placeholder="属性名称/标识符" clearable @keyup.enter="handleSearch">
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
				<el-table-column prop="property_name" label="功能名称" min-width="auto" />
				<el-table-column prop="identifier" label="标识符" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="data_type" label="功能类型" width="auto">
					<template #default="{ row }">
						<el-tag>{{ row.data_type }}</el-tag>
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
				<el-form-item label="属性名称" prop="property_name">
					<el-input v-model="formData.property_name" placeholder="请输入功能名称" />
				</el-form-item>
				<el-form-item label="属性标识符" prop="identifier">
					<el-input v-model="formData.identifier" placeholder="请输入标识符" />
				</el-form-item>
				<el-form-item label="数据类型" prop="data_type">
					<el-select v-model="formData.data_type" placeholder="请选择数据类型">
						<el-option label="整数" value="int" />
						<el-option label="浮点数" value="float" />
						<el-option label="字符串" value="string" />
						<el-option label="布尔值" value="bool" />
						<el-option label="枚举" value="enum" />
						<el-option label="数组" value="array" />
					</el-select>
				</el-form-item>
				<el-form-item label="描述" prop="desc">
					<el-input v-model="formData.desc" type="textarea" rows="3" placeholder="请输入描述信息" />
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
import { getProperties, saveOrUpdateProperty, deleteProperty } from '@/api/property';
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

// 对话框控制
const dialogVisible = ref(false);
const dialogTitle = ref('');
const formRef = ref(null);

// 表单数据
const formData = reactive({
	property_name: '',
	identifier: '',
	data_type: '',
	desc: ''
});

// 表单验证规则
const formRules = {
	property_name: [
		{ required: true, message: '请输入功能名称', trigger: 'blur' },
		{ min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
	],
	identifier: [
		{ required: true, message: '请输入标识符', trigger: 'blur' },
		{ pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/, message: '标识符只能包含字母、数字和下划线，且必须以字母开头', trigger: 'blur' }
	],
	data_type: [{ required: true, message: '请选择数据类型', trigger: 'change' }]
};

// 获取数据
const fetchData = async () => {
	loading.value = true;
	try {
		const params = {
			...searchForm
		};
		const res = await getProperties(params);
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

// 新增
const handleAdd = () => {
	Object.assign(formData, {
		property_name: '',
		identifier: '',
		data_type: '',
		desc: ''
	});
	dialogTitle.value = '新增功能';
	dialogVisible.value = true;
};

// 编辑
const handleEdit = (row) => {
	Object.assign(formData, row);
	dialogTitle.value = '编辑功能';
	dialogVisible.value = true;
};

// 重置
const handleReset = () => {
	searchForm.page = 1;
	searchForm.filters.keywords = '';
	handleSearch();
};

// 删除
const handleDelete = (row) => {
	ElMessageBox.confirm('确认删除该功能?', '提示', {
		type: 'warning'
	}).then(async () => {
		try {
			await deleteProperty(row.id);
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
		await saveOrUpdateProperty(formData);
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
.property-container {
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
