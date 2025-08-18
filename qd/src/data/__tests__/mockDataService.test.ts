/**
 * Mock Data Service Tests
 * 模拟数据服务测试
 */

import { describe, it, expect, beforeEach, vi } from 'vitest'
import { mockDataService } from '../mockDataService'
import { validateUser, validateConversations, validateQuickActions, validateNavigationTabs } from '../validation'

describe('MockDataService', () => {
  beforeEach(() => {
    // 重置网络状态
    mockDataService.setNetworkStatus(true)
  })

  describe('getCurrentUser', () => {
    it('should return current user data', async () => {
      const response = await mockDataService.getCurrentUser()
      
      expect(response.success).toBe(true)
      expect(response.data).toBeDefined()
      expect(validateUser(response.data)).toBe(true)
      expect(response.data.id).toBe('current-user')
    })

    it('should handle network failure', async () => {
      mockDataService.setNetworkStatus(false)
      const response = await mockDataService.getCurrentUser()
      
      expect(response.success).toBe(false)
      expect(response.error).toBe('网络连接失败')
    })
  })

  describe('getConversations', () => {
    it('should return conversations list', async () => {
      const response = await mockDataService.getConversations()
      
      expect(response.success).toBe(true)
      expect(Array.isArray(response.data)).toBe(true)
      expect(validateConversations(response.data)).toBe(true)
      expect(response.data.length).toBeGreaterThan(0)
    })

    it('should handle network failure', async () => {
      mockDataService.setNetworkStatus(false)
      const response = await mockDataService.getConversations()
      
      expect(response.success).toBe(false)
      expect(response.data).toEqual([])
      expect(response.error).toBe('网络连接失败')
    })
  })

  describe('getQuickActions', () => {
    it('should return quick actions list', async () => {
      const response = await mockDataService.getQuickActions()
      
      expect(response.success).toBe(true)
      expect(Array.isArray(response.data)).toBe(true)
      expect(validateQuickActions(response.data)).toBe(true)
      expect(response.data.length).toBe(3)
    })

    it('should contain expected quick actions', async () => {
      const response = await mockDataService.getQuickActions()
      const actionIds = response.data.map(action => action.id)
      
      expect(actionIds).toContain('recent-contacts')
      expect(actionIds).toContain('gift-records')
      expect(actionIds).toContain('profile-views')
    })
  })

  describe('getNavigationTabs', () => {
    it('should return navigation tabs list', async () => {
      const response = await mockDataService.getNavigationTabs()
      
      expect(response.success).toBe(true)
      expect(Array.isArray(response.data)).toBe(true)
      expect(validateNavigationTabs(response.data)).toBe(true)
      expect(response.data.length).toBe(4)
    })

    it('should have chat tab with badge', async () => {
      const response = await mockDataService.getNavigationTabs()
      const chatTab = response.data.find(tab => tab.id === 'chat')
      
      expect(chatTab).toBeDefined()
      expect(chatTab?.badge).toBe(1)
      expect(chatTab?.isActive).toBe(true)
    })
  })

  describe('getHeaderTabs', () => {
    it('should return header tabs list', async () => {
      const response = await mockDataService.getHeaderTabs()
      
      expect(response.success).toBe(true)
      expect(Array.isArray(response.data)).toBe(true)
      expect(response.data.length).toBe(2)
    })

    it('should contain chat and friends tabs', async () => {
      const response = await mockDataService.getHeaderTabs()
      const tabIds = response.data.map(tab => tab.id)
      
      expect(tabIds).toContain('chat')
      expect(tabIds).toContain('friends')
    })
  })

  describe('markAsRead', () => {
    it('should mark conversation as read', async () => {
      const conversationsResponse = await mockDataService.getConversations()
      const firstConversation = conversationsResponse.data[0]
      
      const response = await mockDataService.markAsRead(firstConversation.id)
      
      expect(response.success).toBe(true)
      expect(response.data).toBe(true)
    })

    it('should handle invalid conversation id', async () => {
      const response = await mockDataService.markAsRead('invalid-id')
      
      expect(response.success).toBe(false)
      expect(response.data).toBe(false)
      expect(response.message).toBe('对话不存在')
    })
  })

  describe('switchNavigationTab', () => {
    it('should switch navigation tab successfully', async () => {
      const response = await mockDataService.switchNavigationTab('home')
      
      expect(response.success).toBe(true)
      expect(response.data).toBe(true)
    })

    it('should handle invalid tab id', async () => {
      const response = await mockDataService.switchNavigationTab('invalid-tab')
      
      expect(response.success).toBe(false)
      expect(response.data).toBe(false)
    })
  })

  describe('switchHeaderTab', () => {
    it('should switch header tab successfully', async () => {
      const response = await mockDataService.switchHeaderTab('friends')
      
      expect(response.success).toBe(true)
      expect(response.data).toBe(true)
    })
  })

  describe('getChatInterfaceState', () => {
    it('should return complete chat interface state', async () => {
      const response = await mockDataService.getChatInterfaceState()
      
      expect(response.success).toBe(true)
      expect(response.data).toBeDefined()
      expect(response.data.currentUser).toBeDefined()
      expect(Array.isArray(response.data.conversations)).toBe(true)
      expect(Array.isArray(response.data.quickActions)).toBe(true)
      expect(Array.isArray(response.data.navigationTabs)).toBe(true)
      expect(response.data.isLoading).toBe(false)
      expect(response.data.error).toBeNull()
    })

    it('should handle network failure for complete state', async () => {
      mockDataService.setNetworkStatus(false)
      const response = await mockDataService.getChatInterfaceState()
      
      expect(response.success).toBe(false)
      expect(response.data.error).toBe('网络连接失败')
      expect(response.data.currentUser).toBeNull()
    })
  })

  describe('refreshData', () => {
    it('should refresh data successfully', async () => {
      const response = await mockDataService.refreshData()
      
      expect(response.success).toBe(true)
      expect(response.data).toBe(true)
      expect(response.message).toBe('刷新成功')
    })

    it('should handle network failure during refresh', async () => {
      mockDataService.setNetworkStatus(false)
      const response = await mockDataService.refreshData()
      
      expect(response.success).toBe(false)
      expect(response.error).toBe('网络连接失败')
    })
  })
})