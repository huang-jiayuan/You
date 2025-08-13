<template>
  <div class="login-page">
    <!-- 背景装饰 -->
    <div class="login-background">
      <div class="background-gradient"></div>
      <div class="background-stars">
        <div class="star star-1"></div>
        <div class="star star-2"></div>
        <div class="star star-3"></div>
        <div class="star star-4"></div>
        <div class="star star-5"></div>
      </div>
      <div class="magic-wand">
        <div class="wand-trail"></div>
        <div class="wand-star">⭐</div>
      </div>
    </div>

    <!-- 主要内容 -->
    <div class="login-content">
      <!-- Logo 和标题区域 -->
      <div class="login-header">
        <div class="app-logo">
          <div class="logo-icon">You</div>
        </div>
        <h1 class="app-title">声音交友</h1>
        <p class="app-subtitle">听见温暖的声音</p>
      </div>

      <!-- 登录按钮区域 -->
      <div class="login-actions">
        <button class="login-btn login-btn--phone" @click="handlePhoneLogin">
          <span class="login-btn__icon">📱</span>
          <span class="login-btn__text">手机登录</span>
        </button>

        <button class="login-link" @click="handlePasswordLogin">
          密码登录
        </button>
      </div>

      <!-- 底部协议 -->
      <div class="login-footer">
        <label class="agreement-checkbox">
          <input type="checkbox" v-model="agreedToTerms" />
          <span class="checkbox-mark"></span>
          <span class="agreement-text">
            同意You平台《用户协议》《隐私政策》《登录政策》等协议反馈
          </span>
        </label>
      </div>

      <!-- 底部指示器 -->
      <div class="bottom-indicator"></div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const agreedToTerms = ref(false)

    // 登录方法
    const handlePhoneLogin = () => {
      if (!agreedToTerms.value) {
        alert('请先同意用户协议和隐私政策')
        return
      }
      console.log('跳转到手机验证码登录')
      // 跳转到手机验证码登录页面
      router.push('/phone-login')
    }

    const handlePasswordLogin = () => {
      if (!agreedToTerms.value) {
        alert('请先同意用户协议和隐私政策')
        return
      }
      console.log('跳转到密码登录')
      // 跳转到密码登录页面
      router.push('/password-login')
    }

    return {
      agreedToTerms,
      handlePhoneLogin,
      handlePasswordLogin
    }
  }
}
</script>

<style scoped>
.login-page {
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

.background-stars {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.star {
  position: absolute;
  width: 4px;
  height: 4px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 50%;
  animation: twinkle 2s infinite;
}

.star-1 { top: 20%; left: 10%; animation-delay: 0s; }
.star-2 { top: 30%; right: 20%; animation-delay: 0.5s; }
.star-3 { top: 60%; left: 15%; animation-delay: 1s; }
.star-4 { top: 70%; right: 10%; animation-delay: 1.5s; }
.star-5 { top: 40%; left: 80%; animation-delay: 0.8s; }

@keyframes twinkle {
  0%, 100% { opacity: 0.3; transform: scale(1); }
  50% { opacity: 1; transform: scale(1.2); }
}

/* 魔法棒动画 */
.magic-wand {
  position: absolute;
  top: 25%;
  left: 20%;
  animation: float 3s ease-in-out infinite;
}

.wand-star {
  font-size: 24px;
  filter: drop-shadow(0 0 10px rgba(255, 255, 255, 0.8));
  animation: rotate 4s linear infinite;
}

.wand-trail {
  position: absolute;
  top: 10px;
  left: -50px;
  width: 80px;
  height: 2px;
  background: linear-gradient(90deg, 
    transparent 0%, 
    rgba(255, 255, 255, 0.8) 50%, 
    transparent 100%);
  border-radius: 1px;
  animation: trail 2s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px) rotate(-10deg); }
  50% { transform: translateY(-10px) rotate(-5deg); }
}

@keyframes rotate {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@keyframes trail {
  0%, 100% { opacity: 0.3; transform: scaleX(0.8); }
  50% { opacity: 1; transform: scaleX(1.2); }
}

/* 主要内容 */
.login-content {
  position: relative;
  z-index: 2;
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 60px 40px 40px;
  color: white;
}

/* Logo 和标题 */
.login-header {
  text-align: center;
  margin-top: 80px;
}

.app-logo {
  margin-bottom: 20px;
}

.logo-icon {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  font-weight: bold;
  color: white;
  margin: 0 auto;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.app-title {
  font-size: 24px;
  font-weight: 600;
  margin: 20px 0 8px;
  color: white;
}

.app-subtitle {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.8);
  margin: 0;
}

/* 登录按钮区域 */
.login-actions {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 16px;
  margin: 40px 0;
}

.login-btn {
  width: 100%;
  height: 52px;
  border: none;
  border-radius: 26px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.login-btn--phone {
  background: rgba(255, 255, 255, 0.95);
  color: #333;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 25px rgba(0, 0, 0, 0.15);
}

.login-btn:active {
  transform: translateY(0);
}

.login-btn__icon {
  font-size: 20px;
}

.login-link {
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.8);
  font-size: 16px;
  cursor: pointer;
  padding: 16px;
  text-align: center;
  transition: color 0.3s ease;
}

.login-link:hover {
  color: white;
}

/* 底部协议 */
.login-footer {
  margin-bottom: 20px;
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

.checkbox-mark {
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

/* 底部指示器 */
.bottom-indicator {
  width: 134px;
  height: 5px;
  background: rgba(255, 255, 255, 0.3);
  border-radius: 3px;
  margin: 0 auto;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .login-content {
    padding: 40px 24px 24px;
  }
  
  .login-header {
    margin-top: 40px;
  }
  
  .logo-icon {
    width: 70px;
    height: 70px;
    font-size: 24px;
  }
  
  .app-title {
    font-size: 20px;
  }
  
  .app-subtitle {
    font-size: 14px;
  }
  
  .login-btn {
    height: 48px;
    font-size: 15px;
  }
  
  .agreement-text {
    font-size: 11px;
  }
}

@media (max-width: 375px) {
  .login-content {
    padding: 30px 20px 20px;
  }
  
  .magic-wand {
    display: none;
  }
}

/* 深色主题适配 */
@media (prefers-color-scheme: dark) {
  .login-btn--phone {
    background: rgba(255, 255, 255, 0.9);
    color: #333;
  }
}

/* 高对比度模式 */
@media (prefers-contrast: high) {
  .login-btn {
    border: 2px solid rgba(255, 255, 255, 0.5);
  }
  
  .checkbox-mark {
    border-width: 3px;
  }
}

/* 减少动画模式 */
@media (prefers-reduced-motion: reduce) {
  .star,
  .magic-wand,
  .wand-star,
  .wand-trail {
    animation: none;
  }
  
  .login-btn {
    transition: none;
  }
}
</style>