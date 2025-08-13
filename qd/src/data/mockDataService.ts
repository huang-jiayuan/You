/**
 * Mock Data Service
 * 模拟数据服务 - 提供数据获取和更新逻辑
 */

import type { 
  User, 
  Conversation, 
  QuickAction, 
  NavigationTab, 
  ApiResponse,
  ChatInterfaceState 
} from '@/types/mobile-chat'

import { currentUser, mockUsers, getUserById, updateUserOnlineStatus } from './mockUsers'
import { mockConversations, getConversationById, updateUnreadCount, markConversationAsRead } from './mockConversations'
import { mockQuickActions, getQuickActionById } from './mockQuickActions'
import { mockNavigationTabs, mockHeaderTabs, setActiveNavigationTab, setActiveHeaderTab, updateTabBadge } from './mockNavigationTabs'

/**
 * 模拟API延迟
 */
const API_DELAY = 300

/**
 * 模拟网络延迟
 */
function simulateDelay(ms: number = API_DELAY): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms))
}

/**
 * 模拟API错误
 */
function simulateError(errorRate: number = 0.1): boolean {
  return Math.random() < errorRate
}

/**
 * 创建API响应
 */
function createApiResponse<T>(data: T, success: boolean = true, message?: string): ApiResponse<T> {
  return {
    success,
    data,
    message,
    error: success ? undefined : message
  }
}

/**
 * 数据服务类
 */
export class MockDataService {
  private static instance: MockDataService
  private isOnline: boolean = true

  private constructor() {}

  static getInstance(): MockDataService {
    if (!MockDataService.instance) {
      MockDataService.instance = new MockDataService()
    }
    return MockDataService.instance
  }

  /**
   * 设置网络状态
   */
  setNetworkStatus(isOnline: boolean): void {
    this.isOnline = isOnline
  }

  /**
   * 获取当前用户信息
   */
  async getCurrentUser(): Promise<ApiResponse<User>> {
    await simulateDelay()
    
    if (!this.isOnline) {
      return createApiResponse(currentUser, false, '网络连接失败')
    }

    if (simulateError(0.05)) {
      return createApiResponse(currentUser, false, '获取用户信息失败')
    }

    return createApiResponse(currentUser)
  }

  /**
   * 获取对话列表
   */
  async getConversations(): Promise<ApiResponse<Conversation[]>> {
    await simulateDelay()
    
    if (!this.isOnline) {
      return createApiResponse([], false, '网络连接失败')
    }

    if (simulateError(0.1)) {
      return createApiResponse([], false, '获取对话列表失败')
    }

    return createApiResponse([...mockConversations])
  }

  /**
   * 获取快捷操作列表
   */
  async getQuickActions(): Promise<ApiResponse<QuickAction[]>> {
    await simulateDelay(100)
    
    if (!this.isOnline) {
      return createApiResponse([], false, '网络连接失败')
    }

    return createApiResponse([...mockQuickActions])
  }

  /**
   * 获取导航标签
   */
  async getNavigationTabs(): Promise<ApiResponse<NavigationTab[]>> {
    await simulateDelay(50)
    
    if (!this.isOnline) {
      return createApiResponse([], false, '网络连接失败')
    }

    return createApiResponse([...mockNavigationTabs])
  }

  /**
   * 获取头部标签
   */
  async getHeaderTabs(): Promise<ApiResponse<NavigationTab[]>> {
    await simulateDelay(50)
    
    return createApiResponse([...mockHeaderTabs])
  }

  /**
   * 标记对话为已读
   */
  async markAsRead(conversationId: string): Promise<ApiResponse<boolean>> {
    await simulateDelay(200)
    
    if (!this.isOnline) {
      return createApiResponse(false, false, '网络连接失败')
    }

    if (simulateError(0.05)) {
      return createApiResponse(false, false, '标记已读失败')
    }

    const success = markConversationAsRead(conversationId)
    return createApiResponse(success, success, success ? '标记成功' : '对话不存在')
  }

  /**
   * 更新用户在线状态
   */
  async updateOnlineStatus(userId: string, isOnline: boolean): Promise<ApiResponse<boolean>> {
    await simulateDelay(100)
    
    if (!this.isOnline) {
      return createApiResponse(false, false, '网络连接失败')
    }

    const success = updateUserOnlineStatus(userId, isOnline)
    return createApiResponse(success, success, success ? '状态更新成功' : '用户不存在')
  }

  /**
   * 切换导航标签
   */
  async switchNavigationTab(tabId: string): Promise<ApiResponse<boolean>> {
    await simulateDelay(50)
    
    const success = setActiveNavigationTab(tabId)
    return createApiResponse(success, success, success ? '切换成功' : '标签不存在')
  }

  /**
   * 切换头部标签
   */
  async switchHeaderTab(tabId: string): Promise<ApiResponse<boolean>> {
    await simulateDelay(50)
    
    const success = setActiveHeaderTab(tabId)
    return createApiResponse(success, success, success ? '切换成功' : '标签不存在')
  }

  /**
   * 获取完整的聊天界面状态
   */
  async getChatInterfaceState(): Promise<ApiResponse<ChatInterfaceState>> {
    await simulateDelay(500)
    
    if (!this.isOnline) {
      return createApiResponse({
        activeTab: 'chat',
        conversations: [],
        quickActions: [],
        navigationTabs: [],
        currentUser: null,
        isLoading: false,
        error: '网络连接失败'
      } as ChatInterfaceState, false, '网络连接失败')
    }

    if (simulateError(0.1)) {
      return createApiResponse({
        activeTab: 'chat',
        conversations: [],
        quickActions: [],
        navigationTabs: [],
        currentUser: null,
        isLoading: false,
        error: '数据加载失败'
      } as ChatInterfaceState, false, '数据加载失败')
    }

    const state: ChatInterfaceState = {
      activeTab: 'chat',
      conversations: [...mockConversations],
      quickActions: [...mockQuickActions],
      navigationTabs: [...mockNavigationTabs],
      currentUser: currentUser,
      isLoading: false,
      error: null
    }

    return createApiResponse(state)
  }

  /**
   * 刷新数据
   */
  async refreshData(): Promise<ApiResponse<boolean>> {
    await simulateDelay(800)
    
    if (!this.isOnline) {
      return createApiResponse(false, false, '网络连接失败')
    }

    if (simulateError(0.05)) {
      return createApiResponse(false, false, '刷新失败')
    }

    // 模拟数据更新
    // 这里可以添加更多的数据刷新逻辑
    return createApiResponse(true, true, '刷新成功')
  }
}

// 导出单例实例
export const mockDataService = MockDataService.getInstance()

// 导出便捷方法
export const {
  getCurrentUser,
  getConversations,
  getQuickActions,
  getNavigationTabs,
  getHeaderTabs,
  markAsRead,
  updateOnlineStatus,
  switchNavigationTab,
  switchHeaderTab,
  getChatInterfaceState,
  refreshData
} = mockDataService