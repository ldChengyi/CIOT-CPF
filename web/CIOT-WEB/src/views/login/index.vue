<template>
  <div class="login-container">
    <div class="login-box">
      <!-- Logo和标题 -->
      <div class="title-container">
        <h3 class="title">CIOT Platform</h3>
      </div>

      <!-- 登录/注册表单 -->
      <el-form 
        ref="formRef"
        :model="formData"
        :rules="formRules"
        size="large"
      >
        <el-tabs 
          v-model="activeTab" 
          class="login-tabs"
          @tab-change="handleTabChange"
        >
          <!-- 登录标签页 -->
          <el-tab-pane name="login" label="登录">
            <el-form-item prop="username">
              <el-input
                v-model="formData.username"
                placeholder="用户名"
                prefix-icon="User"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                v-model="formData.password"
                type="password"
                placeholder="密码"
                prefix-icon="Lock"
                show-password
                @keyup.enter="handleSubmit"
              />
            </el-form-item>
            <el-form-item>
              <el-button 
                type="primary" 
                class="submit-btn"
                :loading="loading"
                @click="handleSubmit"
              >
                登录
              </el-button>
            </el-form-item>
          </el-tab-pane>

          <!-- 注册标签页 -->
          <el-tab-pane name="register" label="注册">
            <el-form-item prop="username">
              <el-input
                v-model="formData.username"
                placeholder="用户名"
                prefix-icon="User"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                v-model="formData.password"
                type="password"
                placeholder="密码"
                prefix-icon="Lock"
                show-password
              />
            </el-form-item>
            <el-form-item 
              v-if="activeTab === 'register'" 
              prop="confirmPassword"
            >
              <el-input
                v-model="formData.confirmPassword"
                type="password"
                placeholder="确认密码"
                prefix-icon="Lock"
                show-password
                @keyup.enter="handleSubmit"
              />
            </el-form-item>
            <el-form-item>
              <el-button 
                type="primary" 
                class="submit-btn"
                :loading="loading"
                @click="handleSubmit"
              >
                注册
              </el-button>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const formRef = ref(null)
const loading = ref(false)
const activeTab = ref('login')

// 表单数据
const formData = reactive({
  username: '',
  password: '',
  confirmPassword: ''
})

const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    // 使用 validate() 而不是 validate(validateFields)
    await formRef.value.validate()
    loading.value = true

    if (activeTab.value === 'login') {
      const loginData = {
        username: formData.username,
        password: formData.password
      }
      await userStore.login(loginData)
      const redirect = route.query.redirect || '/'
      router.push(redirect)
    } else {
      const registerData = {
        username: formData.username,
        password: formData.password
      }
      await userStore.register(registerData)
      activeTab.value = 'login'
      formData.password = ''
      formData.confirmPassword = ''
    }
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

// 表单验证规则
const formRules = reactive({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    {
      required: true,
      message: '请确认密码',
      trigger: 'blur',
      validator: (rule, value, callback) => {
        if (activeTab.value === 'register') {
          if (!value) {
            callback(new Error('请确认密码'))
          } else if (value !== formData.password) {
            callback(new Error('两次输入密码不一致'))
          } else {
            callback()
          }
        } else {
          callback()
        }
      }
    }
  ]
})

// 切换标签时重置表单
const handleTabChange = () => {
  formRef.value?.resetFields()
  formData.confirmPassword = ''
}
</script>

<style lang="scss" scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f0f2f5;
  //background-image: url('@/assets/login-bg.svg');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100%;
  
  .login-box {
    width: 420px;
    padding: 40px;
    background: #fff;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    
    .title-container {
      text-align: center;
      margin-bottom: 30px;
      
      .title {
        margin: 0;
        font-size: 24px;
        color: #1f2f3d;
      }
    }
    
    .login-tabs {
      :deep(.el-tabs__nav) {
        width: 100%;
        
        .el-tabs__item {
          width: 50%;
          text-align: center;
        }
      }
    }
    
    .submit-btn {
      width: 100%;
      margin-top: 10px;
    }
  }
}
</style>