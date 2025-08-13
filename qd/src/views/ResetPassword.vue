<template>
  <div class="reset-password-page">
    <!-- 主要内容 -->
    <div class="reset-content">
      <!-- 返回按钮和标题 -->
      <div class="header">
        <button class="back-button" @click="goBack">
          <span class="back-icon">←</span>
        </button>
        <h1 class="page-title">修改密码</h1>
      </div>

      <!-- 表单区域 -->
      <div class="form-container">
        <!-- 步骤指示器 -->
        <div class="step-indicator">
          <div class="step" :class="{ 'step--active': currentStep >= 1, 'step--completed': currentStep > 1 }">
            <div class="step-number">1</div>
            <div class="step-label">验证手机</div>
          </div>
          <div class="step-line" :class="{ 'step-line--active': currentStep > 1 }"></div>
          <div class="step" :class="{ 'step--active': currentStep >= 2, 'step--completed': currentStep > 2 }">
            <div class="step-number">2</div>
            <div class="step-label">设置新密码</div>
          </div>
        </div>

        <!-- 步骤1：验证手机号 -->
        <div v-if="currentStep === 1" class="step-content">
          <div class="step-title">请输入需要修改密码的手机号</div>
          
          <!-- 手机号输入 -->
          <div class="input-group">
            <div class="input-wrapper">
              <span class="country-code">+86</span>
              <input
                type="tel"
                v-model="phoneNumber"
                placeholder="请输入手机号码"
                class="phone-input"
                maxlength="11"
                @input="validatePhone"
              />
            </div>
            <div v-if="phoneError" class="error-message">{{ phoneError }}</div>
          </div>

          <!-- 验证码输入 -->
          <div class="input-group" v-if="showCodeInput">
            <div class="input-wrapper">
              <input
                type="text"
                v-model="verificationCode"
                placeholder="请输入验证码"
                class="code-input"
                maxlength="4"
                @input="validateCode"
              />
              <button
                class="resend-btn"
                @click="sendCode"
                :disabled="countdown > 0"
              >
                {{ countdown > 0 ? `${countdown}s` : '重新发送' }}
              </button>
            </div>
            <div v-if="codeError" class="error-message">{{ codeError }}</div>
          </div>

          <!-- 操作按钮 -->
          <div class="action-buttons">
            <button
              v-if="!showCodeInput"
              class="primary-btn"
              @click="sendCode"
              :disabled="!isPhoneValid || isLoading"
            >
              <span v-if="isLoading" class="loading-spinner"></span>
              {{ isLoading ? '发送中...' : '获取验证码' }}
            </button>

            <button
              v-if="showCodeInput"
              class="primary-btn"
              @click="verifyCode"
              :disabled="!isCodeValid || isLoading"
            >
              <span v-if="isLoading" class="loading-spinner"></span>
              {{ isLoading ? '验证中...' : '下一步' }}
            </button>
          </div>
        </div>

        <!-- 步骤2：设置新密码 -->
        <div v-if="currentStep === 2" class="step-content">
          <div class="step-title">设置新密码</div>
          <div class="step-subtitle">{{ maskedPhone }}</div>
          
          <!-- 新密码输入 -->
          <div class="input-group">
            <div class="input-wrapper">
              <input
                :type="showNewPassword ? 'text' : 'password'"
                v-model="newPassword"
                placeholder="请输入新密码（6-20位）"
                class="password-input"
                @input="validateNewPassword"
              />
              <button
                type="button"
                class="toggle-password"
                @click="toggleNewPassword"
              >
                {{ showNewPassword ? '👁️' : '👁️‍🗨️' }}
              </button>
            </div>
            <div v-if="newPasswordError" class="error-message">{{ newPasswordError }}</div>
          </div>

          <!-- 确认密码输入 -->
          <div class="input-group">
            <div class="input-wrapper">
              <input
                :type="showConfirmPassword ? 'text' : 'password'"
                v-model="confirmPassword"
                placeholder="请再次输入新密码"
                class="password-input"
                @input="validateConfirmPassword"
              />
              <button
                type="button"
                class="toggle-password"
                @click="toggleConfirmPassword"
              >
                {{ showConfirmPassword ? '👁️' : '👁️‍🗨️' }}
              </button>
            </div>
            <div v-if="confirmPasswordError" class="error-message">{{ confirmPasswordError }}</div>
          </div>

          <!-- 操作按钮 -->
          <div class="action-buttons">
            <button
              class="primary-btn"
              @click="resetPassword"
              :disabled="!isPasswordValid || isLoading"
            >
              <span v-if="isLoading" class="loading-spinner"></span>
              {{ isLoading ? '修改中...' : '完成修改' }}
            </button>
          </div>
        </div>

        <!-- 成功提示 -->
        <div v-if="currentStep === 3" class="step-content success-content">
          <div class="success-icon">✓</div>
          <div class="success-title">密码修改成功</div>
          <div class="success-subtitle">请使用新密码登录</div>
          
          <div class="action-buttons">
            <button class="primary-btn" @click="goToLogin">
              返回登录
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '@/api/auth.js'

const router = useRouter()

// 响应式数据
const currentStep = ref(1)
const phoneNumber = ref('')
const verificationCode = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const phoneError = ref('')
const codeError = ref('')
const newPasswordError = ref('')
const confirmPasswordError = ref('')
const showCodeInput = ref(false)
const showNewPassword = ref(false)
const showConfirmPassword = ref(false)
const isLoading = ref(false)
const countdown = ref(0)
let countdownTimer: number | null = null

// 计算属性
const isPhoneValid = computed(() => {
  return /^1[3-9]\d{9}$/.test(phoneNumber.value)
})

const isCodeValid = computed(() => {
  return /^\d{4}$/.test(verificationCode.value)
})

const isPasswordValid = computed(() => {
  return newPassword.value.length >= 6 && 
         newPassword.value.length <= 20 && 
         confirmPassword.value === newPassword.value
})

const maskedPhone = computed(() => {
  if (phoneNumber.value.length === 11) {
    return phoneNumber.value.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
  }
  return phoneNumber.value
})

// 方法
const goBack = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  } else {
    router.go(-1)
  }
}

const validatePhone = () => {
  phoneError.value = ''
  if (phoneNumber.value && !isPhoneValid.value) {
    phoneError.value = '请输入正确的手机号码'
  }
}

const validateCode = () => {
  codeError.value = ''
  if (verificationCode.value && !isCodeValid.value) {
    codeError.value = '请输入4位数字验证码'
  }
}

const validateNewPassword = () => {
  newPasswordError.value = ''
  if (newPassword.value && (newPassword.value.length < 6 || newPassword.value.length > 20)) {
    newPasswordError.value = '密码长度应为6-20位'
  }
  // 重新验证确认密码
  if (confirmPassword.value) {
    validateConfirmPassword()
  }
}

const validateConfirmPassword = () => {
  confirmPasswordError.value = ''
  if (confirmPassword.value && confirmPassword.value !== newPassword.value) {
    confirmPasswordError.value = '两次输入的密码不一致'
  }
}

const sendCode = async () => {
  if (!isPhoneValid.value) {
    phoneError.value = '请输入正确的手机号码'
    return
  }

  isLoading.value = true
  phoneError.value = ''

  try {
    // 调用后端API发送重置密码验证码
    await authAPI.sendCode(phoneNumber.value, 'reset')
    
    showCodeInput.value = true
    startCountdown()
    
    console.log('验证码已发送到', phoneNumber.value)
    
  } catch (error) {
    phoneError.value = error.message || '发送验证码失败，请重试'
  } finally {
    isLoading.value = false
  }
}

const verifyCode = async () => {
  if (!isCodeValid.value) {
    codeError.value = '请输入4位数字验证码'
    return
  }

  isLoading.value = true
  codeError.value = ''

  try {
    // 调用后端API验证验证码
    await authAPI.verifyCode(phoneNumber.value, verificationCode.value, 'reset')
    
    currentStep.value = 2
    
  } catch (error) {
    codeError.value = error.message || '验证码错误，请重新输入'
  } finally {
    isLoading.value = false
  }
}

const resetPassword = async () => {
  if (!isPasswordValid.value) {
    if (newPassword.value.length < 6 || newPassword.value.length > 20) {
      newPasswordError.value = '密码长度应为6-20位'
    } else if (confirmPassword.value !== newPassword.value) {
      confirmPasswordError.value = '两次输入的密码不一致'
    }
    return
  }

  isLoading.value = true
  newPasswordError.value = ''
  confirmPasswordError.value = ''

  try {
    // 调用后端API重置密码
    await authAPI.resetPassword(phoneNumber.value, verificationCode.value, newPassword.value)
    
    console.log('密码修改成功')
    currentStep.value = 3
    
  } catch (error) {
    newPasswordError.value = error.message || '密码修改失败，请重试'
  } finally {
    isLoading.value = false
  }
}

const toggleNewPassword = () => {
  showNewPassword.value = !showNewPassword.value
}

const toggleConfirmPassword = () => {
  showConfirmPassword.value = !showConfirmPassword.value
}

const startCountdown = () => {
  countdown.value = 60
  countdownTimer = window.setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(countdownTimer!)
      countdownTimer = null
    }
  }, 1000)
}

const goToLogin = () => {
  router.push('/')
}

// 清理定时器
onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
})
</script>

<style scoped>
.reset-password-page {
  min-height: 100vh;
  background-color: #f8f9fa;
  display: flex;
  flex-direction: column;
}

/* 主要内容 */
.reset-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

/* 头部区域 */
.header {
  display: flex;
  align-items: center;
  padding: 20px 20px 30px;
  background-color: white;
  position: relative;
}

.back-button {
  background: none;
  border: none;
  padding: 8px;
  cursor: pointer;
  position: absolute;
  left: 20px;
}

.back-icon {
  font-size: 24px;
  color: #333;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0;
  text-align: center;
  flex: 1;
}

/* 表单区域 */
.form-container {
  flex: 1;
  padding: 30px 20px 20px;
  background-color: white;
  margin-top: 10px;
}

/* 步骤指示器 */
.step-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 40px;
}

.step {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.step-number {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: #e0e0e0;
  color: #999;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.step--active .step-number {
  background-color: #667eea;
  color: white;
}

.step--completed .step-number {
  background-color: #10b981;
  color: white;
}

.step-label {
  font-size: 12px;
  color: #999;
  transition: color 0.3s ease;
}

.step--active .step-label {
  color: #333;
}

.step-line {
  width: 60px;
  height: 2px;
  background-color: #e0e0e0;
  margin: 0 20px;
  transition: background-color 0.3s ease;
}

.step-line--active {
  background-color: #10b981;
}

/* 步骤内容 */
.step-content {
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

.step-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  text-align: center;
  margin-bottom: 8px;
}

.step-subtitle {
  font-size: 14px;
  color: #666;
  text-align: center;
  margin-bottom: 30px;
}

/* 输入组 */
.input-group {
  margin-bottom: 24px;
}

.input-wrapper {
  display: flex;
  align-items: center;
  background: #f8f9fa;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 0 16px;
  height: 50px;
  transition: border-color 0.3s ease;
}

.input-wrapper:focus-within {
  border-color: #667eea;
}

.country-code {
  color: #666;
  font-size: 16px;
  margin-right: 12px;
  padding-right: 12px;
  border-right: 1px solid #e0e0e0;
}

.phone-input,
.code-input,
.password-input {
  flex: 1;
  border: none;
  outline: none;
  font-size: 16px;
  color: #333;
  background: transparent;
}

.phone-input::placeholder,
.code-input::placeholder,
.password-input::placeholder {
  color: #999;
}

.resend-btn,
.toggle-password {
  background: none;
  border: none;
  color: #667eea;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 6px;
  transition: background-color 0.3s ease;
}

.resend-btn:hover:not(:disabled),
.toggle-password:hover {
  background-color: rgba(102, 126, 234, 0.1);
}

.resend-btn:disabled {
  color: #999;
  cursor: not-allowed;
}

.error-message {
  color: #ff4757;
  font-size: 14px;
  margin-top: 8px;
  margin-left: 16px;
}

/* 操作按钮 */
.action-buttons {
  margin-top: 30px;
}

.primary-btn {
  width: 100%;
  height: 50px;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.primary-btn:disabled {
  background: #e0e0e0;
  color: #999;
  cursor: not-allowed;
}

.primary-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.loading-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 成功页面 */
.success-content {
  text-align: center;
  padding: 40px 20px;
}

.success-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background-color: #10b981;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 40px;
  font-weight: bold;
  margin: 0 auto 24px;
}

.success-title {
  font-size: 20px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.success-subtitle {
  font-size: 14px;
  color: #666;
  margin-bottom: 40px;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .form-container {
    padding: 20px 16px 16px;
  }
  
  .step-indicator {
    margin-bottom: 30px;
  }
  
  .step-number {
    width: 28px;
    height: 28px;
    font-size: 12px;
  }
  
  .step-line {
    width: 40px;
    margin: 0 15px;
  }
  
  .step-title {
    font-size: 16px;
  }
  
  .input-wrapper {
    height: 48px;
    padding: 0 12px;
  }
  
  .primary-btn {
    height: 48px;
    font-size: 15px;
  }
}

@media (max-width: 375px) {
  .header {
    padding: 16px 16px 20px;
  }
  
  .form-container {
    padding: 16px 12px 12px;
  }
  
  .success-icon {
    width: 60px;
    height: 60px;
    font-size: 30px;
  }
}

/* 减少动画模式 */
@media (prefers-reduced-motion: reduce) {
  .step-number,
  .step-label,
  .step-line,
  .primary-btn {
    transition: none;
  }
  
  .loading-spinner {
    animation: none;
  }
}
</style>