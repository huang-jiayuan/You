/**
 * Mock Quick Actions Data
 * 模拟快捷操作数据
 */

import type { QuickAction } from '@/types/mobile-chat'

export const mockQuickActions: QuickAction[] = [
  {
    id: 'recent-contacts',
    label: '最近来往',
    icon: '👥',
    color: '#FB923C', // 橙色
    route: '/recent-contacts'
  },
  {
    id: 'gift-records',
    label: '收礼记录',
    icon: '🎁',
    color: '#F472B6', // 粉色
    route: '/gift-records'
  },
  {
    id: 'profile-views',
    label: '谁看过我',
    icon: '👁️',
    color: '#8B7CF6', // 紫色
    route: '/profile-views'
  }
]

/**
 * 根据ID获取快捷操作
 */
export function getQuickActionById(id: string): QuickAction | undefined {
  return mockQuickActions.find(action => action.id === id)
}

/**
 * 获取所有快捷操作
 */
export function getAllQuickActions(): QuickAction[] {
  return [...mockQuickActions]
}

/**
 * 添加新的快捷操作
 */
export function addQuickAction(action: Omit<QuickAction, 'id'>): QuickAction {
  const newAction: QuickAction = {
    ...action,
    id: `action-${Date.now()}`
  }
  mockQuickActions.push(newAction)
  return newAction
}

/**
 * 更新快捷操作
 */
export function updateQuickAction(id: string, updates: Partial<Omit<QuickAction, 'id'>>): boolean {
  const actionIndex = mockQuickActions.findIndex(action => action.id === id)
  if (actionIndex !== -1) {
    mockQuickActions[actionIndex] = { ...mockQuickActions[actionIndex], ...updates }
    return true
  }
  return false
}

/**
 * 删除快捷操作
 */
export function removeQuickAction(id: string): boolean {
  const actionIndex = mockQuickActions.findIndex(action => action.id === id)
  if (actionIndex !== -1) {
    mockQuickActions.splice(actionIndex, 1)
    return true
  }
  return false
}