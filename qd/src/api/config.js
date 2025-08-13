/**
 * API配置文件
 * 用于配置后端接口地址和相关设置
 */

// 环境配置
const ENV_CONFIG = {
  development: {
    baseURL: '/api', // 开发环境使用代理，避免跨域问题
    timeout: 10000
  },
  production: {
    baseURL: '/api', // 生产环境后端地址
    timeout: 10000
  }
}

// 获取当前环境配置
const currentEnv = import.meta.env.MODE || 'development'
export const API_CONFIG = ENV_CONFIG[currentEnv]

// API接口路径（根据后端实际路由配置）
export const API_ENDPOINTS = {
  // 用户认证相关
  USER: {
    SEND_SMS: '/user/sendsms',               // 发送短信验证码
    LOGIN: '/user/login',                    // 验证码登录
    PASSWORD_LOGIN: '/user/login/password',  // 密码登录
    UPDATE_PASSWORD: '/user/update/password', // 修改密码
    IMPROVE_INFO: '/user/improve/message',   // 完善用户信息
    FOLLOW: '/user/follow',                  // 关注用户
    UNFOLLOW: '/user/unfollow',              // 取消关注
    FOLLOW_LIST: '/user/follow/list',        // 关注列表
    WEBSOCKET: '/user/ws'                    // WebSocket连接
  },
  
  // 需要认证的接口
  AUTH: {
    PROFILE: '/auth/profile',                // 获取用户信息
    LOGOUT: '/auth/logout'                   // 退出登录
  },
  
  // 其他业务接口
  COMMON: {
    UPLOAD: '/common/upload',                // 文件上传
    CONFIG: '/common/config'                 // 获取配置信息
  }
}

// 请求状态码
export const HTTP_STATUS = {
  SUCCESS: 200,
  CREATED: 201,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  INTERNAL_SERVER_ERROR: 500
}

// 业务状态码（根据后端实际响应码）
export const BUSINESS_CODE = {
  SUCCESS: 0,        // 后端成功状态码是0
  FAIL: 1,
  REQUEST_ERROR: 1000,     // 请求失败！无法找到获取的资源
  SERVER_ERROR: 1001,      // 请求失败！服务器内部错误
  TOKEN_EXPIRED: 1001,
  INVALID_PARAMS: 1000,
  UNAUTHORIZED: 401,
  FORBIDDEN: 403,
  NOT_FOUND: 404,
  SMS_CODE_ERROR: 1002,
  PASSWORD_ERROR: 1004,
  ACCOUNT_BANNED: 1005
}