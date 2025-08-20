/**
 * 全局错误处理工具
 * 用于捕获和处理应用中的未处理错误
 */

import { useNavigationError } from '@/composables/useNavigationError'

interface ErrorInfo {
  message: string
  stack?: string
  componentStack?: string
  errorInfo?: any
}

class GlobalErrorHandler {
  private navigationErrorHandler: ReturnType<typeof useNavigationError> | null = null

  constructor() {
    this.setupErrorListeners()
  }

  // 设置错误监听器
  private setupErrorListeners() {
    // 捕获未处理的Promise拒绝
    window.addEventListener('unhandledrejection', (event) => {
      console.error('未处理的Promise拒绝:', event.reason)
      this.handleError(event.reason, 'unhandledrejection')
      event.preventDefault() // 阻止默认的错误处理
    })

    // 捕获JavaScript运行时错误
    window.addEventListener('error', (event) => {
      console.error('JavaScript错误:', event.error)
      this.handleError(event.error, 'javascript')
    })

    // 捕获资源加载错误
    window.addEventListener('error', (event) => {
      if (event.target !== window) {
        console.error('资源加载错误:', event.target)
        this.handleResourceError(event)
      }
    }, true)
  }

  // 初始化导航错误处理器
  initNavigationErrorHandler() {
    if (!this.navigationErrorHandler) {
      this.navigationErrorHandler = useNavigationError()
    }
  }

  // 处理一般错误
  private handleError(error: any, type: string) {
    const errorInfo: ErrorInfo = {
      message: error?.message || '未知错误',
      stack: error?.stack,
      errorInfo: {
        type,
        timestamp: new Date().toISOString(),
        userAgent: navigator.userAgent,
        url: window.location.href
      }
    }

    // 根据错误类型进行分类处理
    if (this.isNavigationError(error)) {
      this.handleNavigationError(error)
    } else if (this.isNetworkError(error)) {
      this.handleNetworkError(error)
    } else {
      this.handleGenericError(errorInfo)
    }
  }

  // 处理资源加载错误
  private handleResourceError(event: Event) {
    const target = event.target as HTMLElement
    const tagName = target.tagName?.toLowerCase()
    const src = (target as any).src || (target as any).href

    console.error(`资源加载失败: ${tagName} - ${src}`)

    // 对于关键资源加载失败，显示用户提示
    if (tagName === 'script' || tagName === 'link') {
      this.navigationErrorHandler?.showWarningMessage('部分资源加载失败，可能影响功能使用')
    }
  }

  // 判断是否为导航相关错误
  private isNavigationError(error: any): boolean {
    const message = error?.message?.toLowerCase() || ''
    return message.includes('navigation') || 
           message.includes('router') || 
           message.includes('route')
  }

  // 判断是否为网络错误
  private isNetworkError(error: any): boolean {
    const message = error?.message?.toLowerCase() || ''
    return message.includes('network') || 
           message.includes('fetch') || 
           message.includes('timeout') ||
           message.includes('connection')
  }

  // 处理导航错误
  private handleNavigationError(error: any) {
    if (this.navigationErrorHandler) {
      this.navigationErrorHandler.handleNavigationError(error)
    } else {
      console.error('导航错误处理器未初始化:', error)
    }
  }

  // 处理网络错误
  private handleNetworkError(error: any) {
    if (this.navigationErrorHandler) {
      this.navigationErrorHandler.handleNetworkError(error, '网络请求')
    } else {
      console.error('网络错误:', error)
    }
  }

  // 处理一般错误
  private handleGenericError(errorInfo: ErrorInfo) {
    console.error('应用错误:', errorInfo)
    
    // 在开发环境显示详细错误信息
    if (import.meta.env.DEV) {
      console.table(errorInfo)
    }

    // 生产环境只显示用户友好的错误信息
    if (this.navigationErrorHandler) {
      this.navigationErrorHandler.showWarningMessage('应用遇到了一些问题，请刷新页面重试')
    }
  }

  // 手动报告错误
  reportError(error: any, context?: string) {
    console.error(`手动报告错误 [${context}]:`, error)
    this.handleError(error, 'manual')
  }
}

// 创建全局错误处理器实例
export const globalErrorHandler = new GlobalErrorHandler()

// 导出便捷方法
export const reportError = (error: any, context?: string) => {
  globalErrorHandler.reportError(error, context)
}

// 初始化错误处理器
export const initErrorHandler = () => {
  globalErrorHandler.initNavigationErrorHandler()
}