/**
 * 导航功能集成测试
 * 测试从聊天列表到聊天详情的完整导航流程
 */

import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router'
import { nextTick } from 'vue'
import Chat from '../views/Chat.vue'
import ChatDetail from '../views/ChatDetail.vue'

// 创建测试路由
const routes = [
  {
    path: '/chat',
    name: 'Chat',
    component: Chat
  },
  {
    path: '/chat/:userId',
    name: 'ChatDetail',
    component: ChatDetail
  }
]

// Mock useToast
const mockShowError = vi.fn()
const mockShowSuccess = vi.fn()
const mockShowWarning = vi.fn()
const mockShowInfo = vi.fn()
const mockUseToast = vi.fn(() => ({
  error: mockShowError,
  success: mockShowSuccess,
  warning: mockShowWarning,
  info: mockShowInfo
}))

// Mock useNavigationError
const mockHandleNavigationError = vi.fn()
const mockHandleParameterError = vi.fn()
const mockHandleNetworkError = vi.fn()
const mockShowSuccessMessage = vi.fn()
const mockShowWarningMessage = vi.fn()
const mockUseNavigationError = vi.fn(() => ({
  handleNavigationError: mockHandleNavigationError,
  handleParameterError: mockHandleParameterError,
  handleNetworkError: mockHandleNetworkError,
  showSuccessMessage: mockShowSuccessMessage,
  showWarningMessage: mockShowWarningMessage
}))

vi.mock('@/composables/useToast', () => ({
  useToast: mockUseToast
}))

vi.mock('@/composables/useNavigationError', () => ({
  useNavigationError: mockUseNavigationError
}))

describe('导航功能集成测试', () => {
  let router
  let wrapper

  beforeEach(async () => {
    vi.clearAllMocks()
    
    router = createRouter({
      history: createWebHistory(),
      routes
    })

    // 导航到聊天页面
    await router.push('/chat')
    await router.isReady()

    wrapper = mount(Chat, {
      global: {
        plugins: [router],
        stubs: {
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

  describe('完整导航流程', () => {
    it('应该能够从聊天列表导航到聊天详情', async () => {
      // 查找第一个聊天项目
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      expect(chatItem.exists()).toBe(true)

      // 点击聊天项目
      await chatItem.trigger('click')

      // 等待导航完成
      await nextTick()
      await router.isReady()

      // 验证路由已经改变
      expect(router.currentRoute.value.name).toBe('ChatDetail')
      expect(router.currentRoute.value.params.userId).toBeDefined()
    })

    it('应该传递正确的查询参数', async () => {
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      
      await chatItem.trigger('click')
      await nextTick()
      await router.isReady()

      const query = router.currentRoute.value.query
      expect(query.name).toBeDefined()
      expect(query.avatar).toBeDefined()
      expect(query.isOnline).toBeDefined()
    })

    it('应该在导航过程中显示加载状态', async () => {
      const chatItem = wrapper.find('.chat-item:not(.special-item)')
      
      // 开始导航
      const navigationPromise = chatItem.trigger('click')
      
      // 检查加载状态
      await nextTick()
      expect(wrapper.vm.isNavigating).toBe(true)
      expect(wrapper.find('.loading-indicator').exists()).toBe(true)
      
      // 等待导航完成
      await navigationPromise
      await new Promise(resolve => setTimeout(resolve, 350))
      
      // 检查状态重置
      expect(wrapper.vm.isNavigating).toBe(false)
    })
  })

  describe('错误场景测试', () => {
    it('应该处理无效的聊天ID', async () => {
      // 直接调用openChat方法，传入无效ID
      await wrapper.vm.openChat(null)

      expect(mockHandleParameterError).toHaveBeenCalledWith('chatId', null)
      expect(router.currentRoute.value.name).toBe('Chat') // 应该保持在聊天页面
    })

    it('应该处理不存在的用户', async () => {
      await wrapper.vm.openChat(999) // 不存在的用户ID

      expect(mockHandleParameterError).toHaveBeenCalledWith('用户信息', 999)
      expect(router.currentRoute.value.name).toBe('Chat')
    })

    it('应该处理导航失败', async () => {
      // 模拟路由导航失败
      const originalPush = router.push
      router.push = vi.fn().mockRejectedValue(new Error('导航失败'))

      await wrapper.vm.openChat(1)

      expect(mockHandleNavigationError).toHaveBeenCalled()
      
      // 恢复原始方法
      router.push = originalPush
    })
  })

  describe('状态管理测试', () => {
    it('应该正确管理导航状态', async () => {
      expect(wrapper.vm.isNavigating).toBe(false)
      expect(wrapper.vm.navigatingChatId).toBe(null)
      expect(wrapper.vm.navigationError).toBe(null)

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

    it('应该防止重复导航', async () => {
      // 设置导航状态
      wrapper.vm.isNavigating = true

      const consoleSpy = vi.spyOn(console, 'log').mockImplementation(() => {})

      await wrapper.vm.openChat(1)

      expect(consoleSpy).toHaveBeenCalledWith('导航正在进行中，忽略重复点击')
      
      consoleSpy.mockRestore()
    })
  })

  describe('用户体验测试', () => {
    it('应该显示成功消息', async () => {
      await wrapper.vm.openChat(1)

      expect(mockShowSuccessMessage).toHaveBeenCalledWith(
        expect.stringContaining('正在打开与')
      )
    })

    it('应该在导航中禁用其他聊天项目', async () => {
      wrapper.vm.isNavigating = true
      
      await nextTick()

      const disabledItems = wrapper.findAll('.chat-item-disabled')
      expect(disabledItems.length).toBeGreaterThan(0)
    })

    it('应该高亮正在导航的聊天项目', async () => {
      wrapper.vm.isNavigating = true
      wrapper.vm.navigatingChatId = 1

      await nextTick()

      const navigatingItem = wrapper.find('.chat-item-navigating')
      expect(navigatingItem.exists()).toBe(true)
    })
  })

  describe('性能测试', () => {
    it('应该在合理时间内完成导航', async () => {
      const startTime = Date.now()
      
      await wrapper.vm.openChat(1)
      
      const endTime = Date.now()
      const duration = endTime - startTime

      // 导航应该在1秒内完成
      expect(duration).toBeLessThan(1000)
    })

    it('应该正确处理超时', async () => {
      // 模拟超时
      const timeoutError = new Error('导航超时，请检查网络连接')
      router.push = vi.fn().mockRejectedValue(timeoutError)

      await wrapper.vm.openChat(1)

      expect(mockHandleNavigationError).toHaveBeenCalledWith(
        timeoutError,
        expect.any(Object)
      )
    })
  })

  describe('数据完整性测试', () => {
    it('应该传递完整的用户数据', async () => {
      const chatId = 1
      const chatUser = wrapper.vm.chatList.find(chat => chat.id === chatId)
      
      await wrapper.vm.openChat(chatId)
      await router.isReady()

      const route = router.currentRoute.value
      
      expect(route.params.userId).toBe(String(chatId))
      expect(route.query.name).toBe(chatUser.name)
      expect(route.query.avatar).toBe(chatUser.avatar)
      expect(route.query.isOnline).toBe(String(chatUser.isOnline))
    })

    it('应该处理特殊字符', async () => {
      // 修改聊天用户名包含特殊字符
      wrapper.vm.chatList[0].name = '测试用户@#$%'
      
      await wrapper.vm.openChat(1)
      await router.isReady()

      const route = router.currentRoute.value
      expect(route.query.name).toBe('测试用户@#$%')
    })
  })
})