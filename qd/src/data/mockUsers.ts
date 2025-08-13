/**
 * Mock User Data
 * 模拟用户数据
 */

import type { User } from '@/types/mobile-chat'

export const mockUsers: User[] = [
  {
    id: 'user-1',
    name: '小明',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaoming',
    isOnline: true
  },
  {
    id: 'user-2', 
    name: '小红',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaohong',
    isOnline: true
  },
  {
    id: 'user-3',
    name: '小李',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaoli',
    isOnline: false
  },
  {
    id: 'user-4',
    name: '小王',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaowang',
    isOnline: true
  },
  {
    id: 'user-5',
    name: '小张',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaozhang',
    isOnline: false
  },
  {
    id: 'user-6',
    name: '小刘',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaoliu',
    isOnline: true
  },
  {
    id: 'user-7',
    name: '小陈',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaochen',
    isOnline: false
  },
  {
    id: 'user-8',
    name: '小杨',
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=xiaoyang',
    isOnline: true
  }
]

// 当前用户
export const currentUser: User = {
  id: 'current-user',
  name: '我',
  avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=currentuser',
  isOnline: true
}

/**
 * 根据ID获取用户
 */
export function getUserById(id: string): User | undefined {
  if (id === currentUser.id) {
    return currentUser
  }
  return mockUsers.find(user => user.id === id)
}

/**
 * 获取在线用户列表
 */
export function getOnlineUsers(): User[] {
  return mockUsers.filter(user => user.isOnline)
}

/**
 * 更新用户在线状态
 */
export function updateUserOnlineStatus(userId: string, isOnline: boolean): boolean {
  const user = mockUsers.find(u => u.id === userId)
  if (user) {
    user.isOnline = isOnline
    return true
  }
  return false
}