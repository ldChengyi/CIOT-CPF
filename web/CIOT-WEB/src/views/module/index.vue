<template>
	<div class="property-container">
		<!-- 搜索卡片 -->
		<div class="search-card">
			<el-form :model="searchForm" inline class="search-form">
				<el-form-item label="关键词">
					<el-input v-model="searchForm.filters.keywords" placeholder="模块名称" clearable @keyup.enter="handleSearch">
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
				<el-table-column prop="module_name" label="模块名称" min-width="auto" />
				<el-table-column prop="desc" label="描述" min-width="auto" show-overflow-tooltip />
				<el-table-column prop="property_ids" label="模块属性" min-width="auto" show-overflow-tooltip="">
					<template #default="{ row }">
						<el-tag v-for="id in row.property_ids">{{ getPropertyName(id) }}</el-tag>
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
				<el-form-item label="模块名称" prop="module_name">
					<el-input v-model="formData.module_name" placeholder="请输入模块名称" />
				</el-form-item>
				<el-form-item label="描述" prop="desc">
					<el-input v-model="formData.desc" type="textarea" rows="3" placeholder="请输入描述信息" />
				</el-form-item>

				<!-- 属性选择区域（修正后） -->
				<el-form-item label="属性">
					<div class="property-selection">
						<el-popover :visible="showPropertyPopover" placement="bottom" :width="400" trigger="manual">
							<!-- 关键修正：将触发按钮放在reference插槽内 -->
							<template #reference>
								<el-button @click="showPropertyPopover = true">添加属性</el-button>
							</template>
							<!-- 弹窗内容 -->
							<el-checkbox-group v-model="selectedProperties">
								<el-checkbox v-for="property in properties" :key="property.id" :value="property.id">
									{{ property.property_name }}
								</el-checkbox>
							</el-checkbox-group>
							<div style="text-align: right; margin: 10px 0">
								<el-button size="small" @click="showPropertyPopover = false">取消</el-button>
								<el-button size="small" type="primary" @click="handlePropertyConfirm">确定</el-button>
							</div>
						</el-popover>

						<!-- 已选属性标签 -->
						<div class="selected-tags">
							<el-tag v-for="id in formData.property_ids" :key="id" closable @close="removeProperty(id)">
								{{ getPropertyName(id) }}
							</el-tag>
						</div>
					</div>
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
import { getModules, deleteModule, saveOrUpdateModule } from '@/api/module';
import { getProperties } from '@/api/property';
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

// 属性状态
const showPropertyPopover = ref(false);
const properties = ref([]); // 存储所有属性数据
const selectedProperties = ref([]); // 临时存储选中的属性

// 表单数据
const formData = reactive({
	module_name: '',
	desc: '',
	property_ids: []
});

// 表单验证规则
const formRules = {
	module_name: [
		{ required: true, message: '请输入模块名称', trigger: 'blur' },
		{ min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
	]
};

// 获取数据
const fetchData = async () => {
	loading.value = true;
	try {
		const params = {
			...searchForm
		};
		const res = await getModules(params);
		tableData.value = res.data.list;
		total.value = res.data.total;
	} catch (error) {
		ElMessage.error('获取数据失败');
	}
	loading.value = false;
};
// 获取属性数据
const fetchProperties = async () => {
	try {
		const res = await getProperties({ page_size: 999 });
		properties.value = res.data.list;
	} catch (error) {
		ElMessage.error('获取属性失败');
	}
};

// 搜索
const handleSearch = () => {
	searchForm.page = 1;
	fetchData();
};

// 确认属性选择
const handlePropertyConfirm = () => {
	// 合并并去重
	console.log(properties.value);
	const newIds = selectedProperties.value.filter((id) => !formData.property_ids.includes(id));
	formData.property_ids = [...formData.property_ids, ...newIds];
	showPropertyPopover.value = false;
	selectedProperties.value = [];
};

// 移除属性
const removeProperty = (id) => {
	formData.property_ids = formData.property_ids.filter((item) => item !== id);
};

// 获取属性名称
const getPropertyName = (id) => {
	const property = properties.value.find((p) => p.id === id);
	return property?.property_name || '未知属性';
};

// 新增
const handleAdd = () => {
	fetchProperties();
	Object.assign(formData, {
		module_name: '',
		desc: '',
		property_ids: []
	});
	dialogTitle.value = '新增功能';
	dialogVisible.value = true;
};

// 编辑
const handleEdit = (row) => {
	fetchProperties();
	Object.assign(formData, {
		...row,
		property_ids: row.property_ids || []
	});
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
			await deleteModule(row.id);
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
		await saveOrUpdateModule(formData);
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
	fetchProperties();
});
</script>

<style lang="scss" scoped>
.el-popover {
	z-index: 9999 !important;
}

.property-selection {
	display: flex;
	flex-direction: column;
	gap: 8px;

	.selected-tags {
		display: flex;
		flex-wrap: wrap;
		gap: 6px;

		.el-tag {
			margin-right: 6px;
		}
	}
}

.property-item {
	padding: 8px;
	border-bottom: 1px solid #eee;
	&:last-child {
		border-bottom: none;
	}
}

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
