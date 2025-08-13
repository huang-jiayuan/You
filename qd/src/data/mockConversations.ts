/**
 * Mock Conversation Data
 * 模拟对话数据
 */

import type { Conversation } from '@/types/mobile-chat'
import { mockUsers } from './mockUsers'
import { getLatestMessage } from './mockMessages'

/**
 * 生成模拟对话数据
 */
export function generateMockConversations(): Conversation[] {
  const conversations: Conversation[] = []
  
  // 为每个用户创建对话
  mockUsers.forEach((user, index) => {
    const lastMessage = getLatestMessage(user.id)
    
    if (lastMessage) {
      conversations.push({
        id: `conv-${user.id}`,
        participant: user,
        lastMessage,
        unreadCount: Math.floor(Math.random() * 5), // 随机未读数量
        isPinned: index < 2, // 前两个对话置顶
        updatedAt: lastMessage.timestamp
      })
    }
  })
  
  // 按更新时间排序，置顶的在前面
  return conversations.sort((a, b) => {
    if (a.isPinned && !b.isPinned) return -1
    if (!a.isPinned && b.isPinned) return 1
    return b.updatedAt.getTime() - a.updatedAt.getTime()
  })
}

export const mockConversations = generateMockConversations()

/**
 * 根据ID获取对话
 */
export function getConversationById(id: string): Conversation | undefined {
  return mockConversations.find(conv => conv.id === id)
}

/**
 * 获取置顶对话
 */
export function getPinnedConversations(): Conversation[] {
  return mockConversations.filter(conv => conv.isPinned)
}

/**
 * 获取有未读消息的对话
 */
export function getUnreadConversations(): Conversation[] {
  return mockConversations.filter(conv => conv.unreadCount > 0)
}

/**
 * 更新对话未读数量
 */
export function updateUnreadCount(conversationId: string, count: number): boolean {
  const conversation = mockConversations.find(conv => conv.id === conversationId)
  if (conversation) {
    conversation.unreadCount = Math.max(0, count)
    return true
  }
  return false
}

/**
 * 标记对话为已读
 */
export function markConversationAsRead(conversationId: string): boolean {
  return updateUnreadCount(conversationId, 0)
}

/**
 * 切换对话置顶状态
 */
export function toggleConversationPin(conversationId: string): boolean {
  const conversation = mockConversations.find(conv => conv.id === conversationId)
  if (conversation) {
    conversation.isPinned = !conversation.isPinned
    return true
  }
  return false
}