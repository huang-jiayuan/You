/**
 * API响应处理工具类
 * 用于统一处理后端响应格式和错误处理
 */

import { BUSINESS_CODE } from './config.js'

export class ResponseHandler {
  /**
   * 处理API响应数据
   * @param {Object} response - 后端响应数据
   * @returns {*} 处理后的数据
   */
  static handleApiResponse(response) {
    console.log('ResponseHandler 处理响应:', response)
    
    // 检查响应格式是否符合后端统一格式
    if (typeof response === 'object' && response !== null) {
      // 处理标准格式 { code, msg, data }
      if ('code' in response) {
        const { code, msg, data } = response
        
        // 处理成功响应 (code === 0 表示成功)
        if (code === BUSINESS_CODE.SUCCESS) {
          return data || response
        }
        
        // 处理业务错误
        console.warn('业务错误响应:', { code, msg })
        throw new Error(msg || '请求失败')
      }
      
      // 处理其他可能的成功格式
      if (response.success === true || response.status === 'success') {
        return response.data || response
      }
      
      // 处理错误格式
      if (response.error || response.status === 'error') {
        throw new Error(response.message || response.error || '请求失败')
      }
    }
    
    // 如果不是统一格式，直接返回原始数据
    return response
  }

  /**
   * 处理特定的业务错误码
   * @param {number} code - 错误码
   * @param {string} message - 错误消息
   * @returns {Error} 错误对象
   */
  static createBusinessError(code, message) {
    const errorMap = {
      [BUSINESS_CODE.TOKEN_EXPIRED]: '登录已过期，请重新登录',
      [BUSINESS_CODE.UNAUTHORIZED]: '未授权访问，请重新登录',
      [BUSINESS_CODE.FORBIDDEN]: '禁止访问',
      [BUSINESS_CODE.NOT_FOUND]: '请求的资源不存在',
      [BUSINESS_CODE.SMS_CODE_ERROR]: '验证码错误或已过期',
      [BUSINESS_CODE.PASSWORD_ERROR]: '密码错误',
      [BUSINESS_CODE.ACCOUNT_BANNED]: '账号已被封禁，请联系客服',
      [BUSINESS_CODE.INVALID_PARAMS]: '请求参数错误'
    }
    
    const errorMessage = errorMap[code] || message || '请求失败'
    const error = new Error(errorMessage)
    error.code = code
    return error
  }

  /**
   * 检查是否需要重新登录
   * @param {number} code - 错误码
   * @returns {boolean} 是否需要重新登录
   */
  static shouldRelogin(code) {
    return code === BUSINESS_CODE.TOKEN_EXPIRED || code === BUSINESS_CODE.UNAUTHORIZED
  }

  /**
   * 格式化错误信息用于用户显示
   * @param {Error} error - 错误对象
   * @returns {string} 格式化后的错误信息
   */
  static formatErrorMessage(error) {
    if (error.code) {
      // 业务错误，直接返回错误信息
      return error.message
    }
    
    // 网络或其他错误，返回用户友好的信息
    if (error.message.includes('Failed to fetch')) {
      return '网络连接失败，请检查网络连接'
    }
    
    if (error.message.includes('timeout') || error.name === 'AbortError') {
      return '请求超时，请重试'
    }
    
    return error.message || '请求失败，请稍后重试'
  }
}

export default ResponseHandler