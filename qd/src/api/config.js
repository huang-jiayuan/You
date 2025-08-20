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
  
  // 房间相关接口
  ROOM: {
    GIFTS: '/room/:roomId/gifts',            // 获取房间礼物列表
    SEND_GIFT: '/room/gift/send',            // 发送礼物
    GIFT_RECORDS: '/room/:roomId/gift/records' // 礼物记录
  },

  // 用户信息管理接口
  PROFILE: {
    GET_INFO: '/user/profile',               // 获取当前用户信息
    GET_USER_INFO: '/user/center',           // 获取用户详情信息
    UPDATE_INFO: '/user/improve-info'        // 更新用户信息
  },

  // 房间相关接口
  ROOM: {
    GET_RECOMMEND_ROOMS: '/room/getRecommendRooms',     // 获取推荐房间
    GET_ROOMS_BY_CATEGORY: '/room/getRoomsByCategory',  // 根据分类获取房间
    SEARCH_ROOMS: '/room/searchRooms',                  // 搜索房间
    JOIN_ROOM: '/room/joinRoom',                        // 进入房间
    LEAVE_ROOM: '/room/leaveRoom',                      // 离开房间
    CREATE_ROOM: '/room/createroom',                    // 创建房间
    UPDATE_ROOM: '/room/updateroom',                    // 更新房间
    CLOSE_ROOM: '/room/closeRoom',                      // 关闭房间
    // 麦位管理接口
    APPLY_MIC: '/room/applyMic',                        // 申请上麦
    LEAVE_MIC: '/room/leaveMic',                        // 下麦
    HANDLE_MIC_APPLICATION: '/room/handleMicApplication', // 处理麦位申请
    KICK_FROM_MIC: '/room/kickFromMic',                 // 踢人下麦
    MUTE_MIC_USER: '/room/muteMicUser'                  // 禁言麦位用户
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