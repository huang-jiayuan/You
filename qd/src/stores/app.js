import { ref, reactive, readonly, computed } from 'vue'

// 全局应用状态
const state = reactive({
  // 用户信息
  user: {
    name: '',
    email: '',
    avatar: '',
    preferences: {
      language: 'zh-CN',
      timezone: 'Asia/Shanghai',
      notifications: true
    }
  },
  
  // 应用设置
  app: {
    title: 'Vue3 前端应用',
    version: '1.0.0',
    theme: 'light',
    loading: false,
    online: navigator.onLine,
    sidebarCollapsed: false
  },
  
  // 通知系统
  notifications: [],
  
  // 模态框状态
  modals: {
    themeSelector: false,
    userProfile: false,
    settings: false
  },
  
  // 统计数据
  stats: {
    pageViews: 0,
    userActions: 0,
    sessionStart: Date.now(),
    lastActivity: Date.now()
  }
})

// 私有状态（不对外暴露）
const _internal = {
  notificationId: 0,
  activityTimer: null
}

/**
 * 应用状态管理 Store
 * 基于 Vue 3 Composition API 的轻量级状态管理
 */
export function useAppStore() {
  
  // ========== 计算属性 ==========
  
  const isLoggedIn = computed(() => {
    return !!(state.user.name && state.user.email)
  })
  
  const sessionDuration = computed(() => {
    return Date.now() - state.stats.sessionStart
  })
  
  const unreadNotifications = computed(() => {
    return state.notifications.filter(n => !n.read).length
  })
  
  const isOnline = computed(() => state.app.online)
  
  const currentUser = computed(() => ({
    ...state.user,
    isLoggedIn: isLoggedIn.value
  }))
  
  // ========== 用户管理 ==========
  
  const setUser = (userData) => {
    Object.assign(state.user, userData)
    updateActivity()
  }
  
  const updateUserPreferences = (preferences) => {
    Object.assign(state.user.preferences, preferences)
    updateActivity()
  }
  
  const clearUser = () => {
    state.user.name = ''
    state.user.email = ''
    state.user.avatar = ''
    state.user.preferences = {
      language: 'zh-CN',
      timezone: 'Asia/Shanghai',
      notifications: true
    }
    updateActivity()
  }
  
  // ========== 应用设置 ==========
  
  const setTheme = (theme) => {
    state.app.theme = theme
    updateActivity()
  }
  
  const setLoading = (loading) => {
    state.app.loading = loading
  }
  
  const toggleSidebar = () => {
    state.app.sidebarCollapsed = !state.app.sidebarCollapsed
    updateActivity()
  }
  
  const setOnlineStatus = (online) => {
    state.app.online = online
  }
  
  // ========== 通知系统 ==========
  
  const addNotification = (notification) => {
    const id = ++_internal.notificationId
    const newNotification = {
      id,
      type: 'info',
      title: '',
      message: '',
      duration: 5000,
      read: false,
      timestamp: Date.now(),
      ...notification
    }
    
    state.notifications.unshift(newNotification)
    
    // 自动移除通知
    if (newNotification.duration > 0) {
      setTimeout(() => {
        removeNotification(id)
      }, newNotification.duration)
    }
    
    updateActivity()
    return id
  }
  
  const removeNotification = (id) => {
    const index = state.notifications.findIndex(n => n.id === id)
    if (index > -1) {
      state.notifications.splice(index, 1)
    }
  }
  
  const markNotificationAsRead = (id) => {
    const notification = state.notifications.find(n => n.id === id)
    if (notification) {
      notification.read = true
      updateActivity()
    }
  }
  
  const clearAllNotifications = () => {
    state.notifications.length = 0
    updateActivity()
  }
  
  const clearReadNotifications = () => {
    state.notifications = state.notifications.filter(n => !n.read)
    updateActivity()
  }
  
  // ========== 模态框管理 ==========
  
  const openModal = (modalName) => {
    if (modalName in state.modals) {
      state.modals[modalName] = true
      updateActivity()
    }
  }
  
  const closeModal = (modalName) => {
    if (modalName in state.modals) {
      state.modals[modalName] = false
      updateActivity()
    }
  }
  
  const closeAllModals = () => {
    Object.keys(state.modals).forEach(key => {
      state.modals[key] = false
    })
    updateActivity()
  }
  
  // ========== 统计管理 ==========
  
  const incrementPageViews = () => {
    state.stats.pageViews++
    updateActivity()
  }
  
  const incrementUserActions = () => {
    state.stats.userActions++
    updateActivity()
  }
  
  const updateActivity = () => {
    state.stats.lastActivity = Date.now()
    
    // 清除之前的定时器
    if (_internal.activityTimer) {
      clearTimeout(_internal.activityTimer)
    }
    
    // 设置新的定时器，用于检测用户不活跃状态
    _internal.activityTimer = setTimeout(() => {
      // 可以在这里添加不活跃状态的处理逻辑
      console.log('用户不活跃状态')
    }, 30 * 60 * 1000) // 30分钟
  }
  
  const resetStats = () => {
    state.stats = {
      pageViews: 0,
      userActions: 0,
      sessionStart: Date.now(),
      lastActivity: Date.now()
    }
  }
  
  // ========== 数据持久化 ==========
  
  const saveToLocalStorage = () => {
    try {
      const dataToSave = {
        user: state.user,
        app: {
          theme: state.app.theme,
          sidebarCollapsed: state.app.sidebarCollapsed
        },
        stats: state.stats
      }
      localStorage.setItem('app-state', JSON.stringify(dataToSave))
    } catch (error) {
      console.warn('无法保存应用状态到本地存储:', error)
    }
  }
  
  const loadFromLocalStorage = () => {
    try {
      const savedData = localStorage.getItem('app-state')
      if (savedData) {
        const parsedData = JSON.parse(savedData)
        
        // 合并用户数据
        if (parsedData.user) {
          Object.assign(state.user, parsedData.user)
        }
        
        // 合并应用设置
        if (parsedData.app) {
          Object.assign(state.app, parsedData.app)
        }
        
        // 合并统计数据
        if (parsedData.stats) {
          Object.assign(state.stats, parsedData.stats)
        }
      }
    } catch (error) {
      console.warn('无法从本地存储加载应用状态:', error)
    }
  }
  
  const clearLocalStorage = () => {
    try {
      localStorage.removeItem('app-state')
    } catch (error) {
      console.warn('无法清除本地存储:', error)
    }
  }
  
  // ========== 初始化 ==========
  
  const initialize = () => {
    // 加载本地存储的数据
    loadFromLocalStorage()
    
    // 监听在线状态变化
    const handleOnline = () => setOnlineStatus(true)
    const handleOffline = () => setOnlineStatus(false)
    
    window.addEventListener('online', handleOnline)
    window.addEventListener('offline', handleOffline)
    
    // 监听页面可见性变化
    const handleVisibilityChange = () => {
      if (document.visibilityState === 'visible') {
        updateActivity()
      }
    }
    
    document.addEventListener('visibilitychange', handleVisibilityChange)
    
    // 定期保存状态
    const saveInterval = setInterval(saveToLocalStorage, 30000) // 每30秒保存一次
    
    // 页面卸载时保存状态
    const handleBeforeUnload = () => {
      saveToLocalStorage()
    }
    
    window.addEventListener('beforeunload', handleBeforeUnload)
    
    // 返回清理函数
    return () => {
      window.removeEventListener('online', handleOnline)
      window.removeEventListener('offline', handleOffline)
      document.removeEventListener('visibilitychange', handleVisibilityChange)
      window.removeEventListener('beforeunload', handleBeforeUnload)
      clearInterval(saveInterval)
      if (_internal.activityTimer) {
        clearTimeout(_internal.activityTimer)
      }
    }
  }
  
  // ========== 调试工具 ==========
  
  const getDebugInfo = () => {
    return {
      state: JSON.parse(JSON.stringify(state)),
      computed: {
        isLoggedIn: isLoggedIn.value,
        sessionDuration: sessionDuration.value,
        unreadNotifications: unreadNotifications.value,
        isOnline: isOnline.value
      },
      internal: {
        notificationId: _internal.notificationId
      }
    }
  }
  
  // ========== 返回公共API ==========
  
  return {
    // 只读状态
    state: readonly(state),
    
    // 计算属性
    isLoggedIn,
    sessionDuration,
    unreadNotifications,
    isOnline,
    currentUser,
    
    // 用户管理
    setUser,
    updateUserPreferences,
    clearUser,
    
    // 应用设置
    setTheme,
    setLoading,
    toggleSidebar,
    setOnlineStatus,
    
    // 通知系统
    addNotification,
    removeNotification,
    markNotificationAsRead,
    clearAllNotifications,
    clearReadNotifications,
    
    // 模态框管理
    openModal,
    closeModal,
    closeAllModals,
    
    // 统计管理
    incrementPageViews,
    incrementUserActions,
    updateActivity,
    resetStats,
    
    // 数据持久化
    saveToLocalStorage,
    loadFromLocalStorage,
    clearLocalStorage,
    
    // 初始化
    initialize,
    
    // 调试工具
    getDebugInfo
  }
}

// 创建全局实例
let globalStore = null

/**
 * 获取全局应用状态实例
 * 确保整个应用使用同一个状态实例
 */
export function getGlobalAppStore() {
  if (!globalStore) {
    globalStore = useAppStore()
  }
  return globalStore
}