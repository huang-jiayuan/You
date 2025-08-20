/**
 * 导航错误处理组合式函数
 * 专门处理聊天导航相关的错误和用户反馈
 */

import { ref } from 'vue'
import { useToast } from '@/composables/useToast'

interface NavigationError {
  type: 'navigation' | 'parameter' | 'network' | 'timeout' | 'permission'
  code?: string
  message: string
  details?: any
  timestamp: Date
}

export function useNavigationError() {
  const { error: showError, warning: showWarning, info: showInfo } = useToast()
  
  const lastError = ref<NavigationError | null>(null)
  const errorHistory = ref<NavigationError[]>([])

  // 记录错误
  const logError = (error: NavigationError) => {
    lastError.value = error
    errorHistory.value.push(error)
    
    // 保持错误历史记录在合理范围内
    if (errorHistory.value.length > 10) {
      errorHistory.value.shift()
    }
    
    // 开发环境下输出详细错误信息
    if (import.meta.env.DEV) {
      console.error('导航错误:', error)
    }
  }

  // 处理导航错误
  const handleNavigationError = (error: any, context?: any) => {
    const navigationError: NavigationError = {
      type: 'navigation',
      message: '导航失败',
      timestamp: new Date()
    }

    // 根据错误类型进行分类处理
    if (error?.name === 'NavigationDuplicated') {
      navigationError.type = 'navigation'
      navigationError.message = '重复导航请求'
      navigationError.code = 'NAVIGATION_DUPLICATED'
      
      // 重复导航不需要显示错误提示
      logError(navigationError)
      return
    }

    if (error?.message?.includes('timeout')) {
      navigationError.type = 'timeout'
      navigationError.message = '导航超时，请检查网络连接'
      navigationError.code = 'NAVIGATION_TIMEOUT'
    } else if (error?.message?.includes('network')) {
      navigationError.type = 'network'
      navigationError.message = '网络连接失败，请稍后重试'
      navigationError.code = 'NETWORK_ERROR'
    } else if (error?.message?.includes('permission')) {
      navigationError.type = 'permission'
      navigationError.message = '没有访问权限'
      navigationError.code = 'PERMISSION_DENIED'
    } else {
      navigationError.message = error?.message || '未知导航错误'
      navigationError.code = 'UNKNOWN_ERROR'
    }

    navigationError.details = {
      originalError: error,
      context: context
    }

    logError(navigationError)
    
    // 显示用户友好的错误提示
    const userMessage = context?.userName 
      ? `无法打开与${context.userName}的聊天，${navigationError.message}`
      : navigationError.message

    showError(userMessage, {
      duration: 5000,
      closable: true
    })
  }

  // 处理参数错误
  const handleParameterError = (paramName: string, paramValue?: any) => {
    const paramError: NavigationError = {
      type: 'parameter',
      message: `参数错误: ${paramName}`,
      code: 'INVALID_PARAMETER',
      timestamp: new Date(),
      details: {
        paramName,
        paramValue
      }
    }

    logError(paramError)

    const userMessage = paramName === 'userId' 
      ? '无效的聊天链接，请重新尝试'
      : `参数错误: ${paramName}`

    showError(userMessage)
  }

  // 处理网络错误
  const handleNetworkError = (error: any, operation: string = '操作') => {
    const networkError: NavigationError = {
      type: 'network',
      message: `网络错误: ${operation}失败`,
      code: 'NETWORK_ERROR',
      timestamp: new Date(),
      details: {
        operation,
        originalError: error
      }
    }

    logError(networkError)
    showError(`网络连接失败，${operation}无法完成，请检查网络后重试`)
  }

  // 处理超时错误
  const handleTimeoutError = (operation: string = '操作', timeout: number = 5000) => {
    const timeoutError: NavigationError = {
      type: 'timeout',
      message: `超时错误: ${operation}超时`,
      code: 'OPERATION_TIMEOUT',
      timestamp: new Date(),
      details: {
        operation,
        timeout
      }
    }

    logError(timeoutError)
    showError(`${operation}超时，请稍后重试`)
  }

  // 显示成功消息
  const showSuccessMessage = (message: string, options?: any) => {
    showInfo(message, {
      duration: 2000,
      ...options
    })
  }

  // 显示警告消息
  const showWarningMessage = (message: string, options?: any) => {
    showWarning(message, {
      duration: 3000,
      ...options
    })
  }

  // 清除错误历史
  const clearErrorHistory = () => {
    errorHistory.value = []
    lastError.value = null
  }

  // 获取错误统计
  const getErrorStats = () => {
    const stats = {
      total: errorHistory.value.length,
      byType: {} as Record<string, number>
    }

    errorHistory.value.forEach(error => {
      stats.byType[error.type] = (stats.byType[error.type] || 0) + 1
    })

    return stats
  }

  return {
    // 状态
    lastError,
    errorHistory,
    
    // 错误处理方法
    handleNavigationError,
    handleParameterError,
    handleNetworkError,
    handleTimeoutError,
    
    // 用户反馈方法
    showSuccessMessage,
    showWarningMessage,
    
    // 工具方法
    logError,
    clearErrorHistory,
    getErrorStats
  }
}