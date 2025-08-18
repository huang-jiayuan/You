/**
 * 用户相关API服务
 */

import http from './request.js'
import { API_ENDPOINTS } from './config.js'

export const userAPI = {
  /**
   * 修改密码（需要认证）
   * @param {string} password - 新密码
   */
  updatePassword(password) {
    return http.post(API_ENDPOINTS.USER.UPDATE_PASSWORD, {
      password
    })
  },

  /**
   * 完善用户信息（需要认证）
   * @param {Object} userData - 用户数据
   * @param {string} userData.nickname - 用户昵称
   * @param {string} userData.gender - 性别
   * @param {string} userData.voice_tag - 声音标签
   * @param {string} userData.interest_tag - 兴趣标签
   */
  improveInfo(userData) {
    return http.post(API_ENDPOINTS.USER.IMPROVE_INFO, userData)
  },

  /**
   * 关注用户（需要认证）
   * @param {number} followed_id - 被关注用户ID
   */
  follow(followed_id) {
    return http.post(API_ENDPOINTS.USER.FOLLOW, {
      followed_id
    })
  },

  /**
   * 取消关注用户（需要认证）
   * @param {number} followed_id - 被取消关注用户ID
   */
  unfollow(followed_id) {
    return http.post(API_ENDPOINTS.USER.UNFOLLOW, {
      followed_id
    })
  },

  /**
   * 获取关注列表（需要认证）
   */
  getFollowList() {
    return http.post(API_ENDPOINTS.USER.FOLLOW_LIST)
  }
}

export default userAPI