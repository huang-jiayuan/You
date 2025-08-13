# 前后端对接配置说明

## 📋 概述

本项目已配置完成与后端C#服务的对接，包含完整的API请求封装、错误处理、token管理等功能。

## 🔧 配置说明

### 1. 环境配置

#### 开发环境 (`.env.development`)
```env
VITE_APP_TITLE=You声音交友
VITE_API_BASE_URL=http://localhost:8081/api
VITE_APP_ENV=development
```

#### 生产环境 (`.env.production`)
```env
VITE_APP_TITLE=You声音交友
VITE_API_BASE_URL=/api
VITE_APP_ENV=production
```

### 2. 代理配置 (`vite.config.js`)

开发环境下，Vite会自动将 `/api` 开头的请求代理到后端服务：

```javascript
proxy: {
  '/api': {
    target: 'http://localhost:8081', // 后端服务地址
    changeOrigin: true,
    secure: false
  }
}
```

## 📁 API文件结构

```
src/api/
├── config.js      # API配置和常量
├── request.js     # HTTP请求封装
├── auth.js        # 认证相关API
├── user.js        # 用户相关API
└── index.js       # 统一导出
```

## 🔌 API接口说明

### 用户认证接口（无需认证）

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 发送短信验证码 | POST | `/user/sendsms` | 发送注册/登录验证码 |
| 验证码登录 | POST | `/user/login` | 手机号+验证码登录 |
| 密码登录 | POST | `/user/login/password` | 手机号+密码登录 |

### 用户管理接口（需要认证）

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 修改密码 | POST | `/user/update/password` | 修改用户登录密码 |
| 完善用户信息 | POST | `/user/improve/message` | 更新用户个人信息 |
| 关注用户 | POST | `/user/follow` | 关注指定用户 |
| 取消关注 | POST | `/user/unfollow` | 取消关注指定用户 |
| 关注列表 | POST | `/user/follow/list` | 获取用户关注列表 |
| WebSocket连接 | GET | `/user/ws` | 建立WebSocket连接 |

### 认证相关接口（需要认证）

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 获取用户信息 | GET | `/auth/profile` | 获取当前用户信息 |
| 退出登录 | POST | `/auth/logout` | 用户退出登录 |

## 📝 请求/响应格式

### 请求格式

```javascript
// 密码登录
{
  "username": "13812345678",
  "password": "123456",
  "remember_me": true
}

// 发送验证码
{
  "username": "13812345678",
  "source": "login" // login|register
}

// 验证码登录
{
  "username": "13812345678",
  "code": "1234"
}
```

### 响应格式

```javascript
// 成功响应
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "refresh_token_here",
    "user": {
      "id": 1,
      "phone": "13812345678",
      "nickname": "用户昵称"
    }
  }
}

// 错误响应
{
  "code": 1005,
  "message": "密码错误",
  "data": null
}
```

## 🔐 Token管理

### Token存储
- `access_token`: 存储在 localStorage 中
- `refresh_token`: 存储在 localStorage 中

### 自动刷新
当API返回token过期错误时，会自动清除token并跳转到登录页面。

### 请求头
所有需要认证的请求会自动添加Authorization头：
```
Authorization: Bearer {access_token}
```

## 🚀 使用示例

### 在Vue组件中使用

```javascript
import { authAPI, userAPI } from '@/api'

// 发送验证码
const sendSMS = async () => {
  try {
    const response = await authAPI.sendSMS('13800138000', 'login')
    console.log('验证码发送成功', response)
  } catch (error) {
    console.error('发送失败', error.message)
  }
}

// 验证码登录
const login = async () => {
  try {
    const response = await authAPI.login('13800138000', '123456')
    console.log('登录成功', response)
    // 存储token
    localStorage.setItem('access_token', response.token)
  } catch (error) {
    console.error('登录失败', error.message)
  }
}

// 密码登录
const passwordLogin = async () => {
  try {
    const response = await authAPI.passwordLogin('13800138000', '123456', true)
    console.log('登录成功', response)
    localStorage.setItem('access_token', response.token)
  } catch (error) {
    console.error('登录失败', error.message)
  }
}

// 获取用户信息
const getUserInfo = async () => {
  try {
    const userInfo = await authAPI.getProfile()
    console.log('用户信息', userInfo)
  } catch (error) {
    console.error('获取用户信息失败', error.message)
  }
}
```

## ⚙️ 后端要求

### 1. CORS配置
后端需要配置CORS允许前端域名访问：

```csharp
// C# ASP.NET Core 示例
services.AddCors(options =>
{
    options.AddPolicy("AllowFrontend", builder =>
    {
        builder.WithOrigins("http://localhost:3000")
               .AllowAnyMethod()
               .AllowAnyHeader()
               .AllowCredentials();
    });
});
```

### 2. 响应格式
后端API需要返回统一的JSON格式：

```csharp
public class ApiResponse<T>
{
    public int Code { get; set; }
    public string Message { get; set; }
    public T Data { get; set; }
}
```

### 3. 状态码约定
- `code: 0` - 成功
- `code: 1001` - Token过期
- `code: 1004` - 验证码错误
- `code: 1005` - 密码错误

## 🔧 开发调试

### 1. 启动前端开发服务器
```bash
cd qd
npm run dev
```

### 2. 确保后端服务运行在
```
http://localhost:8081
```

### 3. 查看网络请求
在浏览器开发者工具的Network面板中可以查看所有API请求和响应。

## 📞 联调测试

### 快速测试
访问专门的测试页面：`http://localhost:3000/api-test`

### 测试账号
- 手机号: `13800138000`
- 密码: `123456`
- 验证码: `123456`

### 测试流程
1. 启动后端服务：`cd C && start-dev.bat`
2. 启动前端服务：`cd qd && npm run dev`
3. 访问测试页面：`http://localhost:3000/api-test`
4. 按顺序测试各个接口功能
5. 查看浏览器控制台和网络请求
6. 验证所有接口返回正确的数据格式

### 健康检查
- 后端健康检查：`http://localhost:8081/health`
- 前端页面访问：`http://localhost:3000`

## 🐛 常见问题

### 1. 跨域问题
确保后端配置了正确的CORS策略，允许前端域名访问。

### 2. 代理不生效
检查 `vite.config.js` 中的代理配置是否正确。

### 3. Token过期
检查后端JWT配置和过期时间设置。

### 4. 接口404
确认后端路由配置和前端API路径是否一致。

## 📚 相关文档

- [Vite代理配置](https://vitejs.dev/config/server-options.html#server-proxy)
- [Fetch API](https://developer.mozilla.org/zh-CN/docs/Web/API/Fetch_API)
- [JWT Token](https://jwt.io/)