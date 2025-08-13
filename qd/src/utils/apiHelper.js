/**
 * API调用辅助工具
 * 提供统一的API调用和错误处理方法
 */

import { ResponseHandler } from '@/api'

/**
 * 安全的API调用包装器
 * @param {Function} apiCall - API调用函数
 * @param {Object} options - 选项
 * @param {Function} options.onSuccess - 成功回调
 * @param {Function} options.onError - 错误回调
 * @param {boolean} options.showLoading - 是否显示加载状态
 * @returns {Promise} API调用结果
 */
export async function safeApiCall(apiCall, options = {}) {
  const {
    onSuccess,
    onError,
    showLoading = false
  } = options

  try {
    if (showLoading) {
      // 这里可以集成全局loading状态管理
      console.log('Loading started...')
    }

    const result = await apiCall()
    
    if (onSuccess) {
      onSuccess(result)
    }
    
    return result
  } catch (error) {
    const formattedError = ResponseHandler.formatErrorMessage(error)
    
    if (onError) {
      onError(formattedError, error)
    } else {
      // 默认错误处理 - 可以集成全局消息提示
      console.error('API Error:', formattedError)
    }
    
    throw error
  } finally {
    if (showLoading) {
      console.log('Loading finished...')
    }
  }
}

/**
 * 创建带有默认错误处理的API调用函数
 * @param {Function} apiCall - API调用函数
 * @param {string} errorMessage - 默认错误消息
 * @returns {Function} 包装后的API调用函数
 */
export function createApiCall(apiCall, errorMessage = '操作失败') {
  return async (...args) => {
    try {
      return await apiCall(...args)
    } catch (error) {
      const message = ResponseHandler.formatErrorMessage(error)
      console.error(`${errorMessage}: ${message}`)
      throw error
    }
  }
}

/**
 * 批量API调用
 * @param {Array} apiCalls - API调用函数数组
 * @param {Object} options - 选项
 * @returns {Promise<Array>} 所有API调用的结果
 */
export async function batchApiCall(apiCalls, options = {}) {
  const { failFast = false } = options
  
  if (failFast) {
    // 快速失败模式：任何一个失败就停止
    return Promise.all(apiCalls.map(call => call()))
  } else {
    // 容错模式：收集所有结果，包括错误
    const results = await Promise.allSettled(apiCalls.map(call => call()))
    return results.map(result => {
      if (result.status === 'fulfilled') {
        return result.value
      } else {
        console.error('Batch API call failed:', result.reason)
        return null
      }
    })
  }
}