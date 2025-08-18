/**
 * Mock Navigation Tabs Data
 * 模拟导航标签数据
 */

import type { NavigationTab } from '@/types/mobile-chat'

export const mockNavigationTabs: NavigationTab[] = [
  {
    id: 'home',
    label: '首页',
    icon: '🏠',
    route: '/home',
    badge: 0,
    isActive: false
  },
  {
    id: 'nearby',
    label: '附近',
    icon: '📍',
    route: '/nearby',
    badge: 0,
    isActive: false
  },
  {
    id: 'moments',
    label: '动态',
    icon: '📱',
    route: '/moments',
    badge: 0,
    isActive: false
  },
  {
    id: 'chat',
    label: '聊天',
    icon: '💬',
    route: '/chat',
    badge: 1, // 有1个未读消息
    isActive: true
  }
]

// 头部导航标签（聊天和好友）
export const mockHeaderTabs: NavigationTab[] = [
  {
    id: 'chat',
    label: '聊天',
    icon: '💬',
    route: '/chat',
    isActive: true
  },
  {
    id: 'friends',
    label: '好友',
    icon: '👥',
    route: '/friends',
    isActive: false
  }
]

/**
 * 根据ID获取导航标签
 */
export function getNavigationTabById(id: string): NavigationTab | undefined {
  return mockNavigationTabs.find(tab => tab.id === id)
}

/**
 * 获取头部标签
 */
export function getHeaderTabById(id: string): NavigationTab | undefined {
  return mockHeaderTabs.find(tab => tab.id === id)
}

/**
 * 获取活动的导航标签
 */
export function getActiveNavigationTab(): NavigationTab | undefined {
  return mockNavigationTabs.find(tab => tab.isActive)
}

/**
 * 获取活动的头部标签
 */
export function getActiveHeaderTab(): NavigationTab | undefined {
  return mockHeaderTabs.find(tab => tab.isActive)
}

/**
 * 设置活动的导航标签
 */
export function setActiveNavigationTab(id: string): boolean {
  const tab = mockNavigationTabs.find(t => t.id === id)
  if (tab) {
    // 清除所有活动状态
    mockNavigationTabs.forEach(t => t.isActive = false)
    // 设置新的活动标签
    tab.isActive = true
    return true
  }
  return false
}

/**
 * 设置活动的头部标签
 */
export function setActiveHeaderTab(id: string): boolean {
  const tab = mockHeaderTabs.find(t => t.id === id)
  if (tab) {
    // 清除所有活动状态
    mockHeaderTabs.forEach(t => t.isActive = false)
    // 设置新的活动标签
    tab.isActive = true
    return true
  }
  return false
}

/**
 * 更新标签徽章数量
 */
export function updateTabBadge(id: string, count: number): boolean {
  const tab = mockNavigationTabs.find(t => t.id === id)
  if (tab) {
    tab.badge = Math.max(0, count)
    return true
  }
  return false
}

/**
 * 获取所有有徽章的标签
 */
export function getTabsWithBadges(): NavigationTab[] {
  return mockNavigationTabs.filter(tab => tab.badge && tab.badge > 0)
}

/**
 * 清除标签徽章
 */
export function clearTabBadge(id: string): boolean {
  return updateTabBadge(id, 0)
}