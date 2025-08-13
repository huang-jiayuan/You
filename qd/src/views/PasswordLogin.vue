<template>
  <div class="password-login-page">
    <!-- 背景装饰 -->
    <div class="login-background">
      <div class="background-gradient"></div>
    </div>

    <!-- 主要内容 -->
    <div class="login-content">
      <!-- 返回按钮 -->
      <div class="back-button" @click="goBack">
        <span class="back-icon">←</span>
        <span class="back-text">返回</span>
      </div>

      <!-- 标题区域 -->
      <div class="login-header">
        <h1 class="page-title">密码登录</h1>
        <p class="page-subtitle">使用手机号和密码登录</p>
      </div>

      <!-- 登录表单 -->
      <div class="login-form">
        <!-- 手机号输入 -->
        <div class="form-group">
          <label class="form-label">手机号</label>
          <input 
            v-model="phone" 
            type="tel" 
            placeholder="请输入手机号" 
            class="form-input"
            maxlength="11"
          />
        </div>

        <!-- 密码输入 -->
        <div class="form-group">
          <label class="form-label">密码</label>
          <div class="password-input-group">
            <input 
              v-model="password" 
              :type="showPassword ? 'text' : 'password'" 
              placeholder="请输入密码" 
              class="form-input password-input"
            />
            <button 
              class="toggle-password-btn" 
              @click="togglePassword"
              type="button"
            >
              {{ showPassword ? '👁️' : '👁️‍🗨️' }}
            </button>
          </div>
        </div>

        <!-- 记住登录状态 -->
        <div class="form-options">
          <label class="remember-checkbox">
            <input type="checkbox" v-model="rememberMe" />
            <span class="checkbox-mark"></span>
            <span class="checkbox-text">记住登录状态</span>
          </label>
        </div>

        <!-- 登录按钮 -->
        <button 
          class="login-btn" 
          @click="handleLogin"
          :disabled="!phone || !password"
        >
          登录
        </button>

        <!-- 其他登录方式 -->
        <div class="other-login">
          <button class="link-btn" @click="goToPhoneLogin">
            验证码登录
          </button>
          <span class="separator">|</span>
          <button class="link-btn" @click="goToResetPassword">
            忘记密码
          </button>
        </div>
      </div>

      <!-- 底部协议 -->
      <div class="login-footer">
        <label class="agreement-checkbox">
          <input type="checkbox" v-model="agreedToTerms" />
          <span class="checkbox-mark"></span>
          <span class="agreement-text">
            同意You平台《用户协议》《隐私政策》等协议
          </span>
        </label>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '@/api'

export default {
  name: 'PasswordLogin',
  setup() {
    const router = useRouter()
    const phone = ref('')
    const password = ref('')
    const showPassword = ref(false)
    const rememberMe = ref(false)
    const agreedToTerms = ref(true)

    // 切换密码显示
    const togglePassword = () => {
      showPassword.value = !showPassword.value
    }

    // 登录
    const handleLogin = async () => {
      if (!agreedToTerms.value) {
        alert('请先同意用户协议和隐私政策')
        return
      }

      if (!phone.value || !password.value) {
        alert('请输入手机号和密码')
        return
      }

      if (!/^1[3-9]\d{9}$/.test(phone.value)) {
        alert('请输入正确的手机号')
        return
      }

      try {
        const response = await authAPI.passwordLogin(phone.value, password.value, rememberMe.value)
        console.log('登录成功:', response)
        
        // 存储token
        if (response.token) {
          localStorage.setItem('access_token', response.token)
        }
        
        alert('登录成功！')
        router.push('/home')
        
      } catch (error) {
        console.error('登录失败:', error)
        alert('登录失败: ' + error.message)
      }
    }

    // 返回上一页
    const goBack = () => {
      router.go(-1)
    }

    // 跳转到验证码登录
    const goToPhoneLogin = () => {
      router.push('/phone-login')
    }

    // 跳转到重置密码
    const goToResetPassword = () => {
      router.push('/reset-password')
    }

    return {
      phone,
      password,
      showPassword,
      rememberMe,
      agreedToTerms,
      togglePassword,
      handleLogin,
      goBack,
      goToPhoneLogin,
      goToResetPassword
    }
  }
}
</script>

<style scoped>
.password-login-page {
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

/* 背景装饰 */
.login-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

.background-gradient {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, 
    rgba(102, 126, 234, 0.9) 0%, 
    rgba(118, 75, 162, 0.9) 50%,
    rgba(102, 126, 234, 0.8) 100%);
}

/* 主要内容 */
.login-content {
  position: relative;
  z-index: 2;
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 40px;
  color: white;
}

/* 返回按钮 */
.back-button {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  margin-bottom: 40px;
  color: rgba(255, 255, 255, 0.8);
  transition: color 0.3s ease;
}

.back-button:hover {
  color: white;
}

.back-icon {
  font-size: 20px;
}

.back-text {
  font-size: 16px;
}

/* 标题区域 */
.login-header {
  text-align: center;
  margin-bottom: 40px;
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  margin: 0 0 12px 0;
  color: white;
}

.page-subtitle {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.8);
  margin: 0;
}

/* 表单样式 */
.login-form {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 24px;
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 14px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.9);
}

.form-input {
  height: 50px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 12px;
  padding: 0 16px;
  font-size: 16px;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.form-input::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

.form-input:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.6);
  background: rgba(255, 255, 255, 0.15);
}

.password-input-group {
  position: relative;
  display: flex;
  align-items: center;
}

.password-input {
  flex: 1;
  padding-right: 50px;
}

.toggle-password-btn {
  position: absolute;
  right: 12px;
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
  font-size: 16px;
  padding: 4px;
  transition: color 0.3s ease;
}

.toggle-password-btn:hover {
  color: white;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.remember-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
}

.remember-checkbox input[type="checkbox"] {
  display: none;
}

.remember-checkbox .checkbox-mark {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 3px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.remember-checkbox input[type="checkbox"]:checked + .checkbox-mark {
  background: rgba(255, 255, 255, 0.9);
  border-color: rgba(255, 255, 255, 0.9);
}

.remember-checkbox input[type="checkbox"]:checked + .checkbox-mark::after {
  content: '✓';
  color: #667eea;
  font-size: 10px;
  font-weight: bold;
}

.login-btn {
  height: 52px;
  background: rgba(255, 255, 255, 0.95);
  color: #333;
  border: none;
  border-radius: 26px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 25px rgba(0, 0, 0, 0.15);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.other-login {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 12px;
  margin-top: 20px;
}

.link-btn {
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
  cursor: pointer;
  text-decoration: underline;
  padding: 8px;
  transition: color 0.3s ease;
}

.link-btn:hover {
  color: white;
}

.separator {
  color: rgba(255, 255, 255, 0.5);
  font-size: 14px;
}

/* 底部协议 */
.login-footer {
  margin-top: 40px;
}

.agreement-checkbox {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  cursor: pointer;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.7);
  line-height: 1.4;
}

.agreement-checkbox input[type="checkbox"] {
  display: none;
}

.agreement-checkbox .checkbox-mark {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255, 255, 255, 0.5);
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: all 0.3s ease;
}

.agreement-checkbox input[type="checkbox"]:checked + .checkbox-mark {
  background: rgba(255, 255, 255, 0.9);
  border-color: rgba(255, 255, 255, 0.9);
}

.agreement-checkbox input[type="checkbox"]:checked + .checkbox-mark::after {
  content: '✓';
  color: #667eea;
  font-size: 12px;
  font-weight: bold;
}

.agreement-text {
  flex: 1;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .login-content {
    padding: 20px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .form-input,
  .login-btn {
    height: 45px;
  }
  
  .other-login {
    flex-direction: column;
    gap: 8px;
  }
  
  .separator {
    display: none;
  }
}
</style>