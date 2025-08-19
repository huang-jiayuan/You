/**
 * Data Validation Utilities
 * 数据验证工具
 */

import type { User, Conversation, QuickAction, NavigationTab, Message } from '@/types/mobile-chat'

/**
 * 验证用户数据
 */
export function validateUser(user: any): user is User {
  if (!user || typeof user !== 'object') {
    return false
  }

  return (
    typeof user.id === 'string' &&
    user.id.length > 0 &&
    typeof user.name === 'string' &&
    user.name.length > 0 &&
    typeof user.avatar === 'string' &&
    typeof user.isOnline === 'boolean'
  )
}

/**
 * 验证消息数据
 */
export function validateMessage(message: any): message is Message {
  if (!message || typeof message !== 'object') {
    return false
  }

  return (
    typeof message.id === 'string' &&
    message.id.length > 0 &&
    typeof message.content === 'string' &&
    message.content.length > 0 &&
    message.timestamp instanceof Date &&
    typeof message.sender === 'string' &&
    message.sender.length > 0 &&
    ['text', 'image', 'file'].includes(message.type)
  )
}

/**
 * 验证对话数据
 */
export function validateConversation(conversation: any): conversation is Conversation {
  if (!conversation || typeof conversation !== 'object') {
    return false
  }

  return (
    typeof conversation.id === 'string' &&
    conversation.id.length > 0 &&
    validateUser(conversation.participant) &&
    validateMessage(conversation.lastMessage) &&
    typeof conversation.unreadCount === 'number' &&
    conversation.unreadCount >= 0 &&
    typeof conversation.isPinned === 'boolean' &&
    conversation.updatedAt instanceof Date
  )
}

/**
 * 验证快捷操作数据
 */
export function validateQuickAction(action: any): action is QuickAction {
  if (!action || typeof action !== 'object') {
    return false
  }

  return (
    typeof action.id === 'string' &&
    action.id.length > 0 &&
    typeof action.label === 'string' &&
    action.label.length > 0 &&
    typeof action.icon === 'string' &&
    action.icon.length > 0 &&
    typeof action.color === 'string' &&
    action.color.length > 0 &&
    typeof action.route === 'string' &&
    action.route.length > 0
  )
}

/**
 * 验证导航标签数据
 */
export function validateNavigationTab(tab: any): tab is NavigationTab {
  if (!tab || typeof tab !== 'object') {
    return false
  }

  return (
    typeof tab.id === 'string' &&
    tab.id.length > 0 &&
    typeof tab.label === 'string' &&
    tab.label.length > 0 &&
    typeof tab.icon === 'string' &&
    tab.icon.length > 0 &&
    typeof tab.route === 'string' &&
    tab.route.length > 0 &&
    (tab.badge === undefined || (typeof tab.badge === 'number' && tab.badge >= 0)) &&
    (tab.isActive === undefined || typeof tab.isActive === 'boolean')
  )
}

/**
 * 验证用户数组
 */
export function validateUsers(users: any[]): users is User[] {
  if (!Array.isArray(users)) {
    return false
  }
  return users.every(validateUser)
}

/**
 * 验证对话数组
 */
export function validateConversations(conversations: any[]): conversations is Conversation[] {
  if (!Array.isArray(conversations)) {
    return false
  }
  return conversations.every(validateConversation)
}

/**
 * 验证快捷操作数组
 */
export function validateQuickActions(actions: any[]): actions is QuickAction[] {
  if (!Array.isArray(actions)) {
    return false
  }
  return actions.every(validateQuickAction)
}

/**
 * 验证导航标签数组
 */
export function validateNavigationTabs(tabs: any[]): tabs is NavigationTab[] {
  if (!Array.isArray(tabs)) {
    return false
  }
  return tabs.every(validateNavigationTab)
}

/**
 * 数据验证错误类
 */
export class ValidationError extends Error {
  constructor(message: string, public field?: string) {
    super(message)
    this.name = 'ValidationError'
  }
}

/**
 * 验证并抛出错误
 */
export function validateOrThrow<T>(
  data: any,
  validator: (data: any) => data is T,
  errorMessage: string,
  field?: string
): T {
  if (!validator(data)) {
    throw new ValidationError(errorMessage, field)
  }
  return data
}

/**
 * 安全验证 - 返回验证结果和错误信息
 */
export function safeValidate<T>(
  data: any,
  validator: (data: any) => data is T
): { isValid: boolean; data?: T; error?: string } {
  try {
    if (validator(data)) {
      return { isValid: true, data }
    } else {
      return { isValid: false, error: '数据格式不正确' }
    }
  } catch (error) {
    return { 
      isValid: false, 
      error: error instanceof Error ? error.message : '验证过程中发生错误' 
    }
  }
}