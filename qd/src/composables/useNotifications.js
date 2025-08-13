import { computed } from 'vue'
import { getGlobalAppStore } from '../stores/app.js'

/**
 * 通知系统组合式函数
 * 提供便捷的通知管理功能
 */
export function useNotifications() {
  const store = getGlobalAppStore()
  
  // 计算属性
  const notifications = computed(() => store.state.notifications)
  const unreadCount = computed(() => store.unreadNotifications.value)
  const hasUnread = computed(() => unreadCount.value > 0)
  
  // 便捷的通知方法
  const notify = {
    // 成功通知
    success: (message, options = {}) => {
      return store.addNotification({
        type: 'success',
        title: '成功',
        message,
        duration: 3000,
        ...options
      })
    },
    
    // 错误通知
    error: (message, options = {}) => {
      return store.addNotification({
        type: 'error',
        title: '错误',
        message,
        duration: 5000,
        ...options
      })
    },
    
    // 警告通知
    warning: (message, options = {}) => {
      return store.addNotification({
        type: 'warning',
        title: '警告',
        message,
        duration: 4000,
        ...options
      })
    },
    
    // 信息通知
    info: (message, options = {}) => {
      return store.addNotification({
        type: 'info',
        title: '提示',
        message,
        duration: 3000,
        ...options
      })
    },
    
    // 自定义通知
    custom: (notification) => {
      return store.addNotification(notification)
    }
  }
  
  // 批量操作
  const batch = {
    // 标记所有为已读
    markAllAsRead: () => {
      notifications.value.forEach(notification => {
        if (!notification.read) {
          store.markNotificationAsRead(notification.id)
        }
      })
    },
    
    // 清除所有已读
    clearRead: () => {
      store.clearReadNotifications()
    },
    
    // 清除所有
    clearAll: () => {
      store.clearAllNotifications()
    },
    
    // 清除指定类型
    clearByType: (type) => {
      const toRemove = notifications.value
        .filter(n => n.type === type)
        .map(n => n.id)
      
      toRemove.forEach(id => store.removeNotification(id))
    }
  }
  
  // 过滤器
  const filter = {
    // 按类型过滤
    byType: (type) => {
      return computed(() => notifications.value.filter(n => n.type === type))
    },
    
    // 未读通知
    unread: computed(() => notifications.value.filter(n => !n.read)),
    
    // 已读通知
    read: computed(() => notifications.value.filter(n => n.read)),
    
    // 最近的通知
    recent: (count = 5) => {
      return computed(() => 
        notifications.value
          .slice()
          .sort((a, b) => b.timestamp - a.timestamp)
          .slice(0, count)
      )
    }
  }
  
  return {
    // 状态
    notifications,
    unreadCount,
    hasUnread,
    
    // 基础操作
    add: store.addNotification,
    remove: store.removeNotification,
    markAsRead: store.markNotificationAsRead,
    
    // 便捷方法
    notify,
    
    // 批量操作
    batch,
    
    // 过滤器
    filter
  }
}