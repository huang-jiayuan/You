/**
 * Validation Utilities Tests
 * 数据验证工具测试
 */

import { describe, it, expect } from 'vitest'
import {
  validateUser,
  validateMessage,
  validateConversation,
  validateQuickAction,
  validateNavigationTab,
  validateUsers,
  validateConversations,
  validateQuickActions,
  validateNavigationTabs,
  ValidationError,
  validateOrThrow,
  safeValidate
} from '../validation'

describe('Validation Utilities', () => {
  describe('validateUser', () => {
    it('should validate correct user data', () => {
      const validUser = {
        id: 'user-1',
        name: '测试用户',
        avatar: 'https://example.com/avatar.jpg',
        isOnline: true
      }
      
      expect(validateUser(validUser)).toBe(true)
    })

    it('should reject invalid user data', () => {
      const invalidUsers = [
        null,
        undefined,
        {},
        { id: '', name: '测试用户', avatar: 'avatar.jpg', isOnline: true },
        { id: 'user-1', name: '', avatar: 'avatar.jpg', isOnline: true },
        { id: 'user-1', name: '测试用户', avatar: '', isOnline: true },
        { id: 'user-1', name: '测试用户', avatar: 'avatar.jpg', isOnline: 'true' },
        { name: '测试用户', avatar: 'avatar.jpg', isOnline: true }
      ]
      
      invalidUsers.forEach(user => {
        expect(validateUser(user)).toBe(false)
      })
    })
  })

  describe('validateMessage', () => {
    it('should validate correct message data', () => {
      const validMessage = {
        id: 'msg-1',
        content: '测试消息',
        timestamp: new Date(),
        sender: 'user-1',
        type: 'text' as const
      }
      
      expect(validateMessage(validMessage)).toBe(true)
    })

    it('should reject invalid message data', () => {
      const invalidMessages = [
        null,
        undefined,
        {},
        { id: '', content: '测试消息', timestamp: new Date(), sender: 'user-1', type: 'text' },
        { id: 'msg-1', content: '', timestamp: new Date(), sender: 'user-1', type: 'text' },
        { id: 'msg-1', content: '测试消息', timestamp: 'invalid', sender: 'user-1', type: 'text' },
        { id: 'msg-1', content: '测试消息', timestamp: new Date(), sender: '', type: 'text' },
        { id: 'msg-1', content: '测试消息', timestamp: new Date(), sender: 'user-1', type: 'invalid' }
      ]
      
      invalidMessages.forEach(message => {
        expect(validateMessage(message)).toBe(false)
      })
    })
  })

  describe('validateQuickAction', () => {
    it('should validate correct quick action data', () => {
      const validAction = {
        id: 'action-1',
        label: '快捷操作',
        icon: '🔥',
        color: '#FF0000',
        route: '/action'
      }
      
      expect(validateQuickAction(validAction)).toBe(true)
    })

    it('should reject invalid quick action data', () => {
      const invalidActions = [
        null,
        undefined,
        {},
        { id: '', label: '快捷操作', icon: '🔥', color: '#FF0000', route: '/action' },
        { id: 'action-1', label: '', icon: '🔥', color: '#FF0000', route: '/action' },
        { id: 'action-1', label: '快捷操作', icon: '', color: '#FF0000', route: '/action' },
        { id: 'action-1', label: '快捷操作', icon: '🔥', color: '', route: '/action' },
        { id: 'action-1', label: '快捷操作', icon: '🔥', color: '#FF0000', route: '' }
      ]
      
      invalidActions.forEach(action => {
        expect(validateQuickAction(action)).toBe(false)
      })
    })
  })

  describe('validateNavigationTab', () => {
    it('should validate correct navigation tab data', () => {
      const validTab = {
        id: 'tab-1',
        label: '导航标签',
        icon: '🏠',
        route: '/home',
        badge: 5,
        isActive: true
      }
      
      expect(validateNavigationTab(validTab)).toBe(true)
    })

    it('should validate tab without optional fields', () => {
      const validTab = {
        id: 'tab-1',
        label: '导航标签',
        icon: '🏠',
        route: '/home'
      }
      
      expect(validateNavigationTab(validTab)).toBe(true)
    })

    it('should reject invalid navigation tab data', () => {
      const invalidTabs = [
        null,
        undefined,
        {},
        { id: '', label: '导航标签', icon: '🏠', route: '/home' },
        { id: 'tab-1', label: '', icon: '🏠', route: '/home' },
        { id: 'tab-1', label: '导航标签', icon: '', route: '/home' },
        { id: 'tab-1', label: '导航标签', icon: '🏠', route: '' },
        { id: 'tab-1', label: '导航标签', icon: '🏠', route: '/home', badge: -1 },
        { id: 'tab-1', label: '导航标签', icon: '🏠', route: '/home', isActive: 'true' }
      ]
      
      invalidTabs.forEach(tab => {
        expect(validateNavigationTab(tab)).toBe(false)
      })
    })
  })

  describe('array validators', () => {
    it('should validate arrays of valid data', () => {
      const validUsers = [
        { id: 'user-1', name: '用户1', avatar: 'avatar1.jpg', isOnline: true },
        { id: 'user-2', name: '用户2', avatar: 'avatar2.jpg', isOnline: false }
      ]
      
      expect(validateUsers(validUsers)).toBe(true)
    })

    it('should reject arrays with invalid data', () => {
      const invalidUsers = [
        { id: '', name: '用户1', avatar: 'avatar1.jpg', isOnline: true },
        { id: 'user-2', name: '用户2', avatar: 'avatar2.jpg', isOnline: false }
      ]
      
      expect(validateUsers(invalidUsers)).toBe(false)
    })

    it('should reject non-arrays', () => {
      expect(validateUsers(null as any)).toBe(false)
      expect(validateUsers('not an array' as any)).toBe(false)
      expect(validateUsers({} as any)).toBe(false)
    })
  })

  describe('ValidationError', () => {
    it('should create validation error with message', () => {
      const error = new ValidationError('测试错误')
      
      expect(error.name).toBe('ValidationError')
      expect(error.message).toBe('测试错误')
      expect(error.field).toBeUndefined()
    })

    it('should create validation error with field', () => {
      const error = new ValidationError('测试错误', 'testField')
      
      expect(error.name).toBe('ValidationError')
      expect(error.message).toBe('测试错误')
      expect(error.field).toBe('testField')
    })
  })

  describe('validateOrThrow', () => {
    it('should return data if validation passes', () => {
      const validUser = {
        id: 'user-1',
        name: '测试用户',
        avatar: 'avatar.jpg',
        isOnline: true
      }
      
      const result = validateOrThrow(validUser, validateUser, '用户数据无效')
      expect(result).toBe(validUser)
    })

    it('should throw ValidationError if validation fails', () => {
      const invalidUser = { id: '', name: '测试用户' }
      
      expect(() => {
        validateOrThrow(invalidUser, validateUser, '用户数据无效', 'user')
      }).toThrow(ValidationError)
    })
  })

  describe('safeValidate', () => {
    it('should return success result for valid data', () => {
      const validUser = {
        id: 'user-1',
        name: '测试用户',
        avatar: 'avatar.jpg',
        isOnline: true
      }
      
      const result = safeValidate(validUser, validateUser)
      
      expect(result.isValid).toBe(true)
      expect(result.data).toBe(validUser)
      expect(result.error).toBeUndefined()
    })

    it('should return error result for invalid data', () => {
      const invalidUser = { id: '', name: '测试用户' }
      
      const result = safeValidate(invalidUser, validateUser)
      
      expect(result.isValid).toBe(false)
      expect(result.data).toBeUndefined()
      expect(result.error).toBe('数据格式不正确')
    })
  })
})