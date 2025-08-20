/**
 * ChatDetail.vue 组件单元测试
 * 测试聊天详情页面的参数处理和用户信息显示
 */

import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import ChatDetail from '../ChatDetail.vue'

// Mock Vue Router
const mockGo = vi.fn()
const mockRouter = {
  go: mockGo
}

const mockRoute = {
  params: { userId: '1' },
  query: { 
    name: '测试用户', 
    avatar: 'test-avatar.jpg',
    isOnline: 'true'
  }
}

// Mock useToast
const mockShowError = vi.fn()
const mockShowWarning = vi.fn()
const mockUseToast = vi.fn(() => ({
  error: mockShowError,
  warning: mockShowWarning
}))

// Mock useNavigationError
const mockHandleParameterError = vi.fn()
const mockHandleNetworkError = vi.fn()
const mockShowWarningMessage = vi.fn()
const mockUseNavigationError = vi.fn(() => ({
  handleParameterError: mockHandleParameterError,
  handleNetworkError: mockHandleNetworkError,
  showWarningMessage: mockShowWarningMessage
}))

// Mock modules
vi.mock('vue-router', () => ({
  useRouter: () => mockRouter,
  useRoute: () => mockRoute
}))

vi.mock('@/composables/useToast', () => ({
  useToast: mockUseToast
}))

vi.mock('@/composables/useNavigationError', () => ({
  useNavigationError: mockUseNavigationError
}))

describe('ChatDetail.vue', () => {
  let wrapper

  beforeEach(() => {
    // 重置所有mock
    vi.clearAllMocks()
    
    wrapper = mount(ChatDetail, {
      global: {
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

  describe('组件渲染', () => {
    it('应该正确渲染聊天详情页面', () => {
      expect(wrapper.find('.chat-detail-page').exists()).toBe(true)
      expect(wrapper.find('.chat-header').exists()).toBe(true)
      expect(wrapper.find('.user-profile-card').exists()).toBe(true)
      expect(wrapper.find('.chat-messages').exists()).toBe(true)
      expect(wrapper.find('.chat-input-area').exists()).toBe(true)
    })

    it('应该显示用户信息', async () => {
      // 等待组件加载完成
      await nextTick()
      
      expect(wrapper.find('.user-name').text()).toBe('测试用户')
      expect(wrapper.find('.user-avatar img').attributes('src')).toBe('test-avatar.jpg')
    })

    it('应该显示在线状态', async () => {
      await nextTick()
      
      const onlineStatus = wrapper.find('.user-status')
      expect(onlineStatus.exists()).toBe(true)
      expect(onlineStatus.text()).toBe('[在线]')
      expect(onlineStatus.classes()).toContain('online')
    })

    it('应该显示在线指示器', async () => {
      await nextTick()
      
      const onlineIndicator = wrapper.find('.online-indicator')
      expect(onlineIndicator.exists()).toBe(true)
    })
  })

  describe('参数处理', () => {
    it('应该正确加载用户信息', async () => {
      await wrapper.vm.loadChatUser()
      
      expect(wrapper.vm.chatUser.id).toBe('1')
      expect(wrapper.vm.chatUser.name).toBe('测试用户')
      expect(wrapper.vm.chatUser.avatar).toBe('test-avatar.jpg')
      expect(wrapper.vm.chatUser.isOnline).toBe(true)
    })

    it('应该处理缺少用户ID的情况', async () => {
      // 模拟缺少userId参数
      mockRoute.params.userId = null
      
      const result = wrapper.vm.validateParams()
      
      expect(result).toBe(false)
      expect(mockHandleParameterError).toHaveBeenCalledWith('userId', null)
    })

    it('应该处理无效的用户ID格式', async () => {
      // 模拟无效的userId格式
      mockRoute.params.userId = {}
      
      const result = wrapper.vm.validateParams()
      
      expect(result).toBe(false)
      expect(mockHandleParameterError).toHaveBeenCalledWith('userId格式', {})
    })

    it('应该处理不完整的用户名', async () => {
      mockRoute.query.name = '用户'
      
      await wrapper.vm.loadChatUser()
      
      expect(mockShowWarningMessage).toHaveBeenCalledWith('用户名信息不完整，使用默认显示')
    })

    it('应该使用默认头像', async () => {
      mockRoute.query.avatar = ''
      
      await wrapper.vm.loadChatUser()
      
      // 应该使用默认头像
      expect(wrapper.vm.chatUser.avatar).toContain('data:image/svg+xml')
    })
  })

  describe('加载状态', () => {
    it('应该显示加载状态', async () => {
      wrapper.vm.isLoading = true
      
      await nextTick()
      
      const loadingContainer = wrapper.find('.loading-container')
      expect(loadingContainer.exists()).toBe(true)
      expect(loadingContainer.find('.loading-spinner').exists()).toBe(true)
    })

    it('应该显示错误状态', async () => {
      wrapper.vm.isLoading = false
      wrapper.vm.loadError = '加载失败'
      
      await nextTick()
      
      const errorContainer = wrapper.find('.error-container')
      expect(errorContainer.exists()).toBe(true)
      expect(errorContainer.text()).toContain('加载失败')
    })

    it('应该有重试按钮', async () => {
      wrapper.vm.isLoading = false
      wrapper.vm.loadError = '加载失败'
      
      await nextTick()
      
      const retryBtn = wrapper.find('.retry-btn')
      expect(retryBtn.exists()).toBe(true)
      
      const loadChatUserSpy = vi.spyOn(wrapper.vm, 'loadChatUser')
      await retryBtn.trigger('click')
      
      expect(loadChatUserSpy).toHaveBeenCalled()
    })
  })

  describe('错误处理', () => {
    it('应该处理网络错误', async () => {
      const networkError = new Error('network error')
      
      // 模拟网络错误
      vi.spyOn(wrapper.vm, 'loadChatUser').mockRejectedValueOnce(networkError)
      
      try {
        await wrapper.vm.loadChatUser()
      } catch (error) {
        expect(mockHandleNetworkError).toHaveBeenCalledWith(networkError, '加载用户信息')
      }
    })

    it('应该处理一般错误', async () => {
      const generalError = new Error('一般错误')
      
      vi.spyOn(wrapper.vm, 'loadChatUser').mockRejectedValueOnce(generalError)
      
      try {
        await wrapper.vm.loadChatUser()
      } catch (error) {
        expect(mockShowError).toHaveBeenCalledWith('加载用户信息失败，请稍后重试')
      }
    })
  })

  describe('计算属性', () => {
    it('应该正确计算在线状态文本', () => {
      wrapper.vm.chatUser.isOnline = true
      expect(wrapper.vm.onlineStatusText).toBe('[在线]')
      
      wrapper.vm.chatUser.isOnline = false
      expect(wrapper.vm.onlineStatusText).toBe('[离线]')
    })

    it('应该正确计算在线状态样式类', () => {
      wrapper.vm.chatUser.isOnline = true
      expect(wrapper.vm.onlineStatusClass).toBe('online')
      
      wrapper.vm.chatUser.isOnline = false
      expect(wrapper.vm.onlineStatusClass).toBe('offline')
    })
  })

  describe('消息功能', () => {
    it('应该能够发送消息', async () => {
      const messageInput = wrapper.find('.message-input')
      const testMessage = '测试消息'
      
      await messageInput.setValue(testMessage)
      expect(wrapper.vm.newMessage).toBe(testMessage)
      
      const initialMessageCount = wrapper.vm.messages.length
      
      await wrapper.vm.sendMessage()
      
      expect(wrapper.vm.messages.length).toBe(initialMessageCount + 1)
      expect(wrapper.vm.newMessage).toBe('')
      
      const lastMessage = wrapper.vm.messages[wrapper.vm.messages.length - 1]
      expect(lastMessage.content).toBe(testMessage)
      expect(lastMessage.isOwn).toBe(true)
    })

    it('应该忽略空消息', async () => {
      wrapper.vm.newMessage = '   '
      
      const initialMessageCount = wrapper.vm.messages.length
      
      await wrapper.vm.sendMessage()
      
      expect(wrapper.vm.messages.length).toBe(initialMessageCount)
    })

    it('应该支持回车发送', async () => {
      const messageInput = wrapper.find('.message-input')
      await messageInput.setValue('测试消息')
      
      const sendMessageSpy = vi.spyOn(wrapper.vm, 'sendMessage')
      
      await messageInput.trigger('keyup.enter')
      
      expect(sendMessageSpy).toHaveBeenCalled()
    })
  })

  describe('导航功能', () => {
    it('应该能够返回上一页', async () => {
      const backBtn = wrapper.find('.back-btn')
      
      await backBtn.trigger('click')
      
      expect(mockGo).toHaveBeenCalledWith(-1)
    })
  })

  describe('时间格式化', () => {
    it('应该正确格式化时间', () => {
      const testDate = new Date('2023-01-01 14:30:00')
      const formattedTime = wrapper.vm.formatTime(testDate)
      
      expect(formattedTime).toBe('14:30')
    })

    it('应该正确格式化日期', () => {
      const testDate = new Date()
      const formattedDate = wrapper.vm.formatDate(testDate)
      
      expect(formattedDate).toBe('今天')
    })
  })

  describe('响应式设计', () => {
    it('应该在移动端有正确的样式', () => {
      // 检查是否有移动端相关的CSS类
      expect(wrapper.find('.chat-detail-page').exists()).toBe(true)
      
      // 检查用户资料卡片
      const profileCard = wrapper.find('.user-profile-card')
      expect(profileCard.exists()).toBe(true)
    })
  })
})