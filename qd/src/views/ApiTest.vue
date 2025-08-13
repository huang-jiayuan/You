<template>
  <div class="api-test">
    <h1>前后端联调测试</h1>
    
    <div class="test-section">
      <h2>1. 发送验证码测试</h2>
      <input v-model="phone" placeholder="输入手机号" />
      <button @click="testSendSMS">发送验证码</button>
      <p v-if="smsResult">{{ smsResult }}</p>
    </div>

    <div class="test-section">
      <h2>2. 验证码登录测试</h2>
      <input v-model="loginPhone" placeholder="手机号" />
      <input v-model="code" placeholder="验证码" />
      <button @click="testLogin">验证码登录</button>
      <p v-if="loginResult">{{ loginResult }}</p>
    </div>

    <div class="test-section">
      <h2>3. 密码登录测试</h2>
      <input v-model="passwordPhone" placeholder="手机号" />
      <input v-model="password" type="password" placeholder="密码" />
      <button @click="testPasswordLogin">密码登录</button>
      <p v-if="passwordResult">{{ passwordResult }}</p>
    </div>

    <div class="test-section" v-if="token">
      <h2>4. 获取用户信息测试</h2>
      <button @click="testGetProfile">获取用户信息</button>
      <pre v-if="profileResult">{{ profileResult }}</pre>
    </div>
  </div>
</template>

<script>
import { authAPI } from '@/api'

export default {
  name: 'ApiTest',
  data() {
    return {
      phone: '13800138000',
      loginPhone: '13800138000',
      passwordPhone: '13800138000',
      code: '123456',
      password: '123456',
      token: '',
      smsResult: '',
      loginResult: '',
      passwordResult: '',
      profileResult: ''
    }
  },
  methods: {
    async testSendSMS() {
      try {
        const result = await authAPI.sendSMS(this.phone, 'login')
        this.smsResult = `成功: ${JSON.stringify(result)}`
      } catch (error) {
        this.smsResult = `失败: ${error.message}`
      }
    },

    async testLogin() {
      try {
        const result = await authAPI.login(this.loginPhone, this.code)
        this.token = result.token
        this.loginResult = `成功: ${JSON.stringify(result)}`
      } catch (error) {
        this.loginResult = `失败: ${error.message}`
      }
    },

    async testPasswordLogin() {
      try {
        const result = await authAPI.passwordLogin(this.passwordPhone, this.password, true)
        this.token = result.token
        this.passwordResult = `成功: ${JSON.stringify(result)}`
      } catch (error) {
        this.passwordResult = `失败: ${error.message}`
      }
    },

    async testGetProfile() {
      try {
        const result = await authAPI.getProfile()
        this.profileResult = JSON.stringify(result, null, 2)
      } catch (error) {
        this.profileResult = `失败: ${error.message}`
      }
    }
  }
}
</script>

<style scoped>
.api-test {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.test-section {
  margin: 20px 0;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

input {
  margin: 5px;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 3px;
}

button {
  margin: 5px;
  padding: 8px 15px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}

button:hover {
  background: #0056b3;
}

pre {
  background: #f8f9fa;
  padding: 10px;
  border-radius: 3px;
  overflow-x: auto;
}
</style>