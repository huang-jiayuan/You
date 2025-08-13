/**
 * Mock Message Data
 * 模拟消息数据
 */

import type { Message } from '@/types/mobile-chat'

export const mockMessages: Message[] = [
  {
    id: 'msg-1',
    content: '你好，最近怎么样？',
    timestamp: new Date('2024-01-15T10:30:00'),
    sender: 'user-1',
    type: 'text'
  },
  {
    id: 'msg-2',
    content: '还不错，你呢？',
    timestamp: new Date('2024-01-15T10:32:00'),
    sender: 'current-user',
    type: 'text'
  },
  {
    id: 'msg-3',
    content: '今天天气真好',
    timestamp: new Date('2024-01-15T09:15:00'),
    sender: 'user-2',
    type: 'text'
  },
  {
    id: 'msg-4',
    content: '是的，适合出去走走',
    timestamp: new Date('2024-01-15T09:16:00'),
    sender: 'current-user',
    type: 'text'
  },
  {
    id: 'msg-5',
    content: '晚上一起吃饭吗？',
    timestamp: new Date('2024-01-14T18:20:00'),
    sender: 'user-3',
    type: 'text'
  },
  {
    id: 'msg-6',
    content: '好的，几点？',
    timestamp: new Date('2024-01-14T18:22:00'),
    sender: 'current-user',
    type: 'text'
  },
  {
    id: 'msg-7',
    content: '工作顺利吗？',
    timestamp: new Date('2024-01-14T14:30:00'),
    sender: 'user-4',
    type: 'text'
  },
  {
    id: 'msg-8',
    content: '还可以，谢谢关心',
    timestamp: new Date('2024-01-14T14:35:00'),
    sender: 'current-user',
    type: 'text'
  },
  {
    id: 'msg-9',
    content: '周末有什么计划？',
    timestamp: new Date('2024-01-13T16:45:00'),
    sender: 'user-5',
    type: 'text'
  },
  {
    id: 'msg-10',
    content: '可能去看电影',
    timestamp: new Date('2024-01-13T16:50:00'),
    sender: 'current-user',
    type: 'text'
  },
  {
    id: 'msg-11',
    content: '新年快乐！',
    timestamp: new Date('2024-01-01T00:01:00'),
    sender: 'user-6',
    type: 'text'
  },
  {
    id: 'msg-12',
    content: '新年快乐！祝你身体健康！',
    timestamp: new Date('2024-01-01T00:02:00'),
    sender: 'current-user',
    type: 'text'
  }
]

/**
 * 根据发送者获取消息
 */
export function getMessagesBySender(senderId: string): Message[] {
  return mockMessages.filter(message => message.sender === senderId)
}

/**
 * 获取两个用户之间的对话消息
 */
export function getConversationMessages(userId1: string, userId2: string): Message[] {
  return mockMessages
    .filter(message => 
      (message.sender === userId1) || (message.sender === userId2)
    )
    .sort((a, b) => a.timestamp.getTime() - b.timestamp.getTime())
}

/**
 * 获取最新消息
 */
export function getLatestMessage(userId: string): Message | undefined {
  const userMessages = mockMessages
    .filter(message => message.sender === userId || message.sender === 'current-user')
    .sort((a, b) => b.timestamp.getTime() - a.timestamp.getTime())
  
  return userMessages[0]
}

/**
 * 创建新消息
 */
export function createMessage(content: string, sender: string, type: 'text' | 'image' | 'file' = 'text'): Message {
  return {
    id: `msg-${Date.now()}`,
    content,
    timestamp: new Date(),
    sender,
    type
  }
}