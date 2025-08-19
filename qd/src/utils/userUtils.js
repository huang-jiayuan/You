/**
 * 用户信息工具函数
 * 用于获取和管理用户身份信息
 */

/**
 * 从JWT token中解析用户信息
 * @param {string} token - JWT token
 * @returns {object|null} 解析出的用户信息
 */
export function parseTokenUserInfo(token) {
  if (!token) {
    return null
  }

  try {
    // JWT token格式: header.payload.signature
    const parts = token.split('.')
    if (parts.length !== 3) {
      console.warn('Invalid JWT token format')
      return null
    }

    // 解析payload部分
    const payload = parts[1]
    // Base64解码
    const decodedPayload = atob(payload)
    const userInfo = JSON.parse(decodedPayload)

    return {
      id: userInfo.id || userInfo.ID || userInfo.userId || userInfo.user_id,
      username: userInfo.username || userInfo.Username,
      nickname: userInfo.nickname || userInfo.NickName,
      exp: userInfo.exp, // 过期时间
      iat: userInfo.iat  // 签发时间
    }
  } catch (error) {
    console.error('Failed to parse JWT token:', error)
    return null
  }
}

/**
 * 获取当前用户ID
 * @returns {number|null} 用户ID
 */
export function getCurrentUserId() {
  // 首先尝试从localStorage获取token
  const token = localStorage.getItem('access_token')
  if (!token) {
    console.warn('No access token found in localStorage')
    return null
  }

  const userInfo = parseTokenUserInfo(token)
  if (!userInfo || !userInfo.id) {
    console.warn('Failed to get user ID from token')
    return null
  }

  return parseInt(userInfo.id)
}

/**
 * 获取当前用户完整信息
 * @returns {object|null} 用户信息对象
 */
export function getCurrentUserInfo() {
  const token = localStorage.getItem('access_token')
  if (!token) {
    return null
  }

  return parseTokenUserInfo(token)
}

/**
 * 检查token是否有效（未过期）
 * @param {string} token - JWT token
 * @returns {boolean} token是否有效
 */
export function isTokenValid(token) {
  if (!token) {
    return false
  }

  const userInfo = parseTokenUserInfo(token)
  if (!userInfo || !userInfo.exp) {
    return false
  }

  // 检查是否过期（exp是秒级时间戳）
  const currentTime = Math.floor(Date.now() / 1000)
  return userInfo.exp > currentTime
}

/**
 * 检查当前用户是否已登录
 * @returns {boolean} 是否已登录
 */
export function isUserLoggedIn() {
  const token = localStorage.getItem('access_token')
  return token && isTokenValid(token)
}

/**
 * 清除用户认证信息
 */
export function clearUserAuth() {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  console.log('User authentication cleared')
}

/**
 * 设置用户认证token
 * @param {string} accessToken - 访问token
 * @param {string} refreshToken - 刷新token（可选）
 */
export function setUserAuth(accessToken, refreshToken = null) {
  if (accessToken) {
    localStorage.setItem('access_token', accessToken)
  }
  if (refreshToken) {
    localStorage.setItem('refresh_token', refreshToken)
  }
  console.log('User authentication set')
}