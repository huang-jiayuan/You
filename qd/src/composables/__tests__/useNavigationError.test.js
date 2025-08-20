/**
 * useNavigationError 组合式函数单元测试
 * 测试导航错误处理功能
 */

import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { useNavigationError } from '../useNavigationError'

// Mock useToast
const mockShowError = vi.fn()
const mockShowWarning = vi.fn()
const mockShowInfo = vi.fn()
const mockUseToast = vi.fn(() => ({
  error: mockShowError,
  warning: mockShowWarning,
  info: mockShowInfo
}))

vi.mock('@/composables/useToast', () => ({
  useToast: mockUseToast
}))

describe('useNavigationError', () => {
  let navigationError

  beforeEach(() => {
    vi.clearAllMocks()
    navigationError = useNavigationError()
  })

  afterEach(() => {
    navigationError.clearErrorHistory()
  })

  describe('错误记录', () => {
    it('应该记录错误到历史记录', () => {
      const error = {
        type: 'navigation',
        message: '测试错误',
        timestamp: new Date()
      }

      navigationError.logError(error)

      expect(navigationError.lastError.value).toEqual(error)
      expect(navigationError.errorHistory.value).toContain(error)
    })

    it('应该限制错误历史记录数量', () => {
      // 添加超过10个错误
      for (let i = 0; i < 15; i++) {
        navigationError.logError({
          type: 'navigation',
          message: `错误 ${i}`,
          timestamp: new Date()
        })
      }

      expect(navigationError.errorHistory.value.length).toBe(10)
    })
  })

  describe('导航错误处理', () => {
    it('应该处理重复导航错误', () => {
      const error = { name: 'NavigationDuplicated', message: '重复导航' }
      
      navigationError.handleNavigationError(error)

      // 重复导航不应该显示错误提示
      expect(mockShowError).not.toHaveBeenCalled()
      expect(navigationError.lastError.value.type).toBe('navigation')
      expect(navigationError.lastError.value.code).toBe('NAVIGATION_DUPLICATED')
    })

    it('应该处理超时错误', () => {
      const error = { message: 'timeout occurred' }
      
      navigationError.handleNavigationError(error)

      expect(mockShowError).toHaveBeenCalledWith(
        '导航超时，请检查网络连接',
        expect.any(Object)
      )
      expect(navigationError.lastError.value.type).toBe('timeout')
      expect(navigationError.lastError.value.code).toBe('NAVIGATION_TIMEOUT')
    })

    it('应该处理网络错误', () => {
      const error = { message: 'network connection failed' }
      
      navigationError.handleNavigationError(error)

      expect(mockShowError).toHaveBeenCalledWith(
        '网络连接失败，请稍后重试',
        expect.any(Object)
      )
      expect(navigationError.lastError.value.type).toBe('network')
      expect(navigationError.lastError.value.code).toBe('NETWORK_ERROR')
    })

    it('应该处理权限错误', () => {
      const error = { message: 'permission denied' }
      
      navigationError.handleNavigationError(error)

      expect(mockShowError).toHaveBeenCalledWith(
        '没有访问权限',
        expect.any(Object)
      )
      expect(navigationError.lastError.value.type).toBe('permission')
      expect(navigationError.lastError.value.code).toBe('PERMISSION_DENIED')
    })

    it('应该处理未知错误', () => {
      const error = { message: '未知错误' }
      
      navigationError.handleNavigationError(error)

      expect(mockShowError).toHaveBeenCalledWith(
        '未知错误',
        expect.any(Object)
      )
      expect(navigationError.lastError.value.code).toBe('UNKNOWN_ERROR')
    })

    it('应该包含用户上下文信息', () => {
      const error = { message: '导航失败' }
      const context = { userName: '测试用户' }
      
      navigationError.handleNavigationError(error, context)

      expect(mockShowError).toHaveBeenCalledWith(
        '无法打开与测试用户的聊天，导航失败',
        expect.any(Object)
      )
    })
  })

  describe('参数错误处理', () => {
    it('应该处理用户ID参数错误', () => {
      navigationError.handleParameterError('userId', null)

      expect(mockShowError).toHaveBeenCalledWith('无效的聊天链接，请重新尝试')
      expect(navigationError.lastError.value.type).toBe('parameter')
      expect(navigationError.lastError.value.code).toBe('INVALID_PARAMETER')
    })

    it('应该处理其他参数错误', () => {
      navigationError.handleParameterError('testParam', 'invalidValue')

      expect(mockShowError).toHaveBeenCalledWith('参数错误: testParam')
      expect(navigationError.lastError.value.details.paramName).toBe('testParam')
      expect(navigationError.lastError.value.details.paramValue).toBe('invalidValue')
    })
  })

  describe('网络错误处理', () => {
    it('应该处理网络错误', () => {
      const error = new Error('网络连接失败')
      
      navigationError.handleNetworkError(error, '加载数据')

      expect(mockShowError).toHaveBeenCalledWith('网络连接失败，加载数据无法完成，请检查网络后重试')
      expect(navigationError.lastError.value.type).toBe('network')
      expect(navigationError.lastError.value.details.operation).toBe('加载数据')
    })

    it('应该使用默认操作名称', () => {
      const error = new Error('网络错误')
      
      navigationError.handleNetworkError(error)

      expect(mockShowError).toHaveBeenCalledWith('网络连接失败，操作无法完成，请检查网络后重试')
    })
  })

  describe('超时错误处理', () => {
    it('应该处理超时错误', () => {
      navigationError.handleTimeoutError('加载页面', 5000)

      expect(mockShowError).toHaveBeenCalledWith('加载页面超时，请稍后重试')
      expect(navigationError.lastError.value.type).toBe('timeout')
      expect(navigationError.lastError.value.details.timeout).toBe(5000)
    })

    it('应该使用默认参数', () => {
      navigationError.handleTimeoutError()

      expect(mockShowError).toHaveBeenCalledWith('操作超时，请稍后重试')
      expect(navigationError.lastError.value.details.operation).toBe('操作')
      expect(navigationError.lastError.value.details.timeout).toBe(5000)
    })
  })

  describe('用户反馈消息', () => {
    it('应该显示成功消息', () => {
      navigationError.showSuccessMessage('操作成功')

      expect(mockShowInfo).toHaveBeenCalledWith('操作成功', {
        duration: 2000
      })
    })

    it('应该显示警告消息', () => {
      navigationError.showWarningMessage('警告信息')

      expect(mockShowWarning).toHaveBeenCalledWith('警告信息', {
        duration: 3000
      })
    })

    it('应该支持自定义选项', () => {
      const options = { duration: 5000, closable: true }
      
      navigationError.showSuccessMessage('成功', options)

      expect(mockShowInfo).toHaveBeenCalledWith('成功', {
        duration: 5000,
        closable: true
      })
    })
  })

  describe('工具方法', () => {
    it('应该清除错误历史', () => {
      // 添加一些错误
      navigationError.logError({
        type: 'navigation',
        message: '错误1',
        timestamp: new Date()
      })
      navigationError.logError({
        type: 'parameter',
        message: '错误2',
        timestamp: new Date()
      })

      expect(navigationError.errorHistory.value.length).toBe(2)

      navigationError.clearErrorHistory()

      expect(navigationError.errorHistory.value.length).toBe(0)
      expect(navigationError.lastError.value).toBe(null)
    })

    it('应该获取错误统计', () => {
      // 添加不同类型的错误
      navigationError.logError({
        type: 'navigation',
        message: '导航错误1',
        timestamp: new Date()
      })
      navigationError.logError({
        type: 'navigation',
        message: '导航错误2',
        timestamp: new Date()
      })
      navigationError.logError({
        type: 'parameter',
        message: '参数错误',
        timestamp: new Date()
      })

      const stats = navigationError.getErrorStats()

      expect(stats.total).toBe(3)
      expect(stats.byType.navigation).toBe(2)
      expect(stats.byType.parameter).toBe(1)
    })
  })

  describe('开发环境行为', () => {
    it('应该在开发环境输出详细错误', () => {
      const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {})
      
      // 模拟开发环境
      vi.stubEnv('DEV', true)
      
      const error = {
        type: 'navigation',
        message: '测试错误',
        timestamp: new Date()
      }

      navigationError.logError(error)

      expect(consoleSpy).toHaveBeenCalledWith('导航错误:', error)
      
      consoleSpy.mockRestore()
    })
  })
})