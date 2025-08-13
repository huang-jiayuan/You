/**
 * 认证相关API服务
 */

import http from './request.js'
import { API_ENDPOINTS } from './config.js'

export const authAPI = {
  /**
   * 发送短信验证码
   * @param {string} username - 手机号
   * @param {string} source - 验证码用途 (login|register)
   */
  sendSMS(username, source = 'login') {
    return http.post(API_ENDPOINTS.USER.SEND_SMS, {
      username,
      source
    })
  },

  /**
   * 验证码登录
   * @param {string} username - 手机号
   * @param {string} code - 验证码
   */
  login(username, code) {
    return http.post(API_ENDPOINTS.USER.LOGIN, {
      username,
      code
    })
  },

  /**
   * 密码登录
   * @param {string} username - 手机号
   * @param {string} password - 密码
   * @param {boolean} remember_me - 是否记住登录状态
   */
  passwordLogin(username, password, remember_me = false) {
    return http.post(API_ENDPOINTS.USER.PASSWORD_LOGIN, {
      username,
      password,
      remember_me
    })
  },

  /**
   * 获取用户信息（需要认证）
   */
  getProfile() {
    return http.get(API_ENDPOINTS.AUTH.PROFILE)
  },

  /**
   * 退出登录（需要认证）
   */
  logout() {
    return http.post(API_ENDPOINTS.AUTH.LOGOUT)
  }
}

export default authAPI