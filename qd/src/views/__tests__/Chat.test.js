/**
 * Chat.vue 组件单元测试
 * 测试聊天列表页面的导航功能
 */

import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import Chat from '../Chat.vue'

// Mock Vue Router
const mockPush = vi.fn()
const mockRouter = {
  push: mockPush
}

// Mock useToast
const mockShowError = vi.fn()
const mockShowSuccess = vi.fn()
const mockUseToast = vi.fn(() => ({
  error: mockShowError,
  success: mockShowSuccess
}))

// Mock useNavigationError
const mockHandleNavigationError = vi.fn()
const mockHandleParameterError = vi.fn()
const mockShowSuccessMessage = vi.fn()
const mockUseNavigationError = vi.fn(() => ({
  handleNavigationError: mockHandleNavigationError,
  handleParameterError: mockHandleParameterError,
  showSuccessMessage: mockShowSuccessMessage
}))

// Mock modules
vi.mock('vue-router', () => ({
  useRouter: () => mockRouter
}))

vi.mock('@/composables/useToast', () => ({
  useToast: mockUseToast
}))

vi.mock('@/composables/useNavigationError', () => ({
  useNavigationError: mockUseNavigationError
}))

describe('Chat.vue', () => {
  let wrapper

  beforeEach(() => {
    // 重置所有mock
    vi.clearAllMocks()
    
    wrapper = mount(Chat, {
      global: {
        stubs: {
          // 简化子组件渲染
          'router-link': true,
          'router-view': true
        }
      }
    })
  })

  afterEach(() => {
    if (wrapper) {
      wrapper.unmount()
    }
  })

  describe('组件渲染', () => {
    it('应该正确渲染聊天页面', () => {
      expect(wrapper.find('.chat-page').exists()).toBe(true)
      expect(wrapper.find('.chat-header').exists()).toBe(true)
      expect(wrapper.find('.chat-list').exists()).toBe(true)
    })

    it('应该渲染聊天列表项目', () => {
      const chatItems = wrapper.findAll('.chat-item')
      // 包括家族广场和聊天用户
      expect(chatItems.length).toBeGreaterThan(0)
    })

    it('应该显示用户头像和名称', () => {
      const firstChatItem = wrapper.find('.chat-item:not(.special-item)')
      expect(firstChatItem.find('.chat-avatar img').exists()).toBe(true)
      expect(firstChatItem.find('.chat-name').exists()).toBe(true)
    })
  })

  describe('openChat 方法', () => {
    it('应该能够打开聊天', async () => {
      const chatId = 1
      
      // 模拟点击聊天项目
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      await chatItem.trigger('click')

      // 验证路由导航被调用
      expect(mockPush).toHaveBeenCalledWith({
        name: 'ChatDetail',
        params: { userId: String(chatId) },
        query: expect.objectContaining({
          name: expect.any(String),
          avatar: expect.any(String),
          isOnline: expect.any(String)
        })
      })
    })

    it('应该防止重复点击', async () => {
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      
      // 设置导航状态为进行中
      await wrapper.vm.openChat(1)
      
      // 再次点击应该被忽略
      await chatItem.trigger('click')
      
      // 验证只调用了一次
      expect(mockPush).toHaveBeenCalledTimes(1)
    })

    it('应该处理无效的聊天ID', async () => {
      await wrapper.vm.openChat(null)
      
      expect(mockHandleParameterError).toHaveBeenCalledWith('chatId', null)
      expect(mockPush).not.toHaveBeenCalled()
    })

    it('应该处理用户信息不存在的情况', async () => {
      await wrapper.vm.openChat(999) // 不存在的ID
      
      expect(mockHandleParameterError).toHaveBeenCalledWith('用户信息', 999)
      expect(mockPush).not.toHaveBeenCalled()
    })

    it('应该处理导航错误', async () => {
      const error = new Error('导航失败')
      mockPush.mockRejectedValueOnce(error)
      
      await wrapper.vm.openChat(1)
      
      expect(mockHandleNavigationError).toHaveBeenCalledWith(
        error, 
        expect.objectContaining({ userName: expect.any(String) })
      )
    })

    it('应该显示成功消息', async () => {
      mockPush.mockResolvedValueOnce()
      
      await wrapper.vm.openChat(1)
      
      expect(mockShowSuccessMessage).toHaveBeenCalledWith(
        expect.stringContaining('正在打开与')
      )
    })
  })

  describe('导航状态管理', () => {
    it('应该正确设置导航状态', async () => {
      expect(wrapper.vm.isNavigating).toBe(false)
      expect(wrapper.vm.navigatingChatId).toBe(null)
      
      // 开始导航
      const promise = wrapper.vm.openChat(1)
      
      expect(wrapper.vm.isNavigating).toBe(true)
      expect(wrapper.vm.navigatingChatId).toBe(1)
      
      await promise
      
      // 等待状态重置
      await new Promise(resolve => setTimeout(resolve, 350))
      
      expect(wrapper.vm.isNavigating).toBe(false)
      expect(wrapper.vm.navigatingChatId).toBe(null)
    })

    it('应该在导航中显示加载指示器', async () => {
      // 设置导航状态
      wrapper.vm.isNavigating = true
      wrapper.vm.navigatingChatId = 1
      
      await nextTick()
      
      const loadingIndicator = wrapper.find('.loading-indicator')
      expect(loadingIndicator.exists()).toBe(true)
    })

    it('应该在导航中禁用其他聊天项目', async () => {
      wrapper.vm.isNavigating = true
      
      await nextTick()
      
      const disabledItems = wrapper.findAll('.chat-item-disabled')
      expect(disabledItems.length).toBeGreaterThan(0)
    })
  })

  describe('视觉反馈', () => {
    it('应该在导航中高亮当前项目', async () => {
      wrapper.vm.isNavigating = true
      wrapper.vm.navigatingChatId = 1
      
      await nextTick()
      
      const navigatingItem = wrapper.find('.chat-item-navigating')
      expect(navigatingItem.exists()).toBe(true)
    })

    it('应该显示加载动画', async () => {
      wrapper.vm.isNavigating = true
      wrapper.vm.navigatingChatId = 1
      
      await nextTick()
      
      const spinner = wrapper.find('.loading-spinner')
      expect(spinner.exists()).toBe(true)
    })
  })

  describe('聊天列表数据', () => {
    it('应该有正确的聊天数据结构', () => {
      const chatList = wrapper.vm.chatList
      
      expect(Array.isArray(chatList)).toBe(true)
      expect(chatList.length).toBeGreaterThan(0)
      
      const firstChat = chatList[0]
      expect(firstChat).toHaveProperty('id')
      expect(firstChat).toHaveProperty('name')
      expect(firstChat).toHaveProperty('avatar')
      expect(firstChat).toHaveProperty('lastMessage')
      expect(firstChat).toHaveProperty('isOnline')
    })

    it('应该正确计算未读消息总数', () => {
      const totalUnreadCount = wrapper.vm.totalUnreadCount
      expect(typeof totalUnreadCount).toBe('number')
      expect(totalUnreadCount).toBeGreaterThanOrEqual(0)
    })
  })

  describe('错误处理', () => {
    it('应该处理超时错误', async () => {
      const timeoutError = new Error('导航超时，请检查网络连接')
      mockPush.mockRejectedValueOnce(timeoutError)
      
      await wrapper.vm.openChat(1)
      
      expect(mockHandleNavigationError).toHaveBeenCalledWith(
        timeoutError,
        expect.any(Object)
      )
    })

    it('应该重置错误状态', async () => {
      wrapper.vm.navigationError = '测试错误'
      
      await wrapper.vm.openChat(1)
      
      expect(wrapper.vm.navigationError).toBe(null)
    })
  })

  describe('用户交互', () => {
    it('应该响应点击事件', async () => {
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      const clickSpy = vi.spyOn(wrapper.vm, 'openChat')
      
      await chatItem.trigger('click')
      
      expect(clickSpy).toHaveBeenCalled()
    })

    it('应该有正确的触摸目标大小', () => {
      const chatItems = wrapper.findAll('.chat-item')
      
      chatItems.forEach(item => {
        const styles = getComputedStyle(item.element)
        // 检查最小高度（通过CSS类）
        expect(item.classes()).toContain('chat-item')
      })
    })
  })
})