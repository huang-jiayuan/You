import { computed, onMounted, onUnmounted } from 'vue'
import { getGlobalAppStore } from '../stores/app.js'

/**
 * 应用状态管理组合式函数
 * 提供应用级别的状态管理功能
 */
export function useApp() {
  const store = getGlobalAppStore()
  
  // 计算属性
  const appInfo = computed(() => ({
    title: store.state.app.title,
    version: store.state.app.version,
    theme: store.state.app.theme
  }))
  
  const isLoading = computed(() => store.state.app.loading)
  const isOnline = computed(() => store.isOnline.value)
  const isSidebarCollapsed = computed(() => store.state.app.sidebarCollapsed)
  
  // 模态框状态
  const modals = computed(() => store.state.modals)
  
  // 统计信息
  const stats = computed(() => store.state.stats)
  const sessionInfo = computed(() => ({
    duration: store.sessionDuration.value,
    startTime: store.state.stats.sessionStart,
    lastActivity: store.state.stats.lastActivity,
    pageViews: store.state.stats.pageViews,
    userActions: store.state.stats.userActions
  }))
  
  // 应用控制
  const setLoading = (loading, message = '') => {
    store.setLoading(loading)
    if (loading && message) {
      // 可以显示加载消息
      console.log('加载中:', message)
    }
  }
  
  const toggleSidebar = () => {
    store.toggleSidebar()
  }
  
  // 模态框管理
  const modal = {
    open: (modalName) => store.openModal(modalName),
    close: (modalName) => store.closeModal(modalName),
    closeAll: () => store.closeAllModals(),
    isOpen: (modalName) => computed(() => store.state.modals[modalName] || false)
  }
  
  // 页面管理
  const page = {
    // 页面访问统计
    visit: () => {
      store.incrementPageViews()
    },
    
    // 设置页面标题
    setTitle: (title) => {
      document.title = title ? `${title} - ${appInfo.value.title}` : appInfo.value.title
    },
    
    // 设置页面元数据
    setMeta: (meta) => {
      if (meta.description) {
        let metaDesc = document.querySelector('meta[name="description"]')
        if (!metaDesc) {
          metaDesc = document.createElement('meta')
          metaDesc.name = 'description'
          document.head.appendChild(metaDesc)
        }
        metaDesc.content = meta.description
      }
      
      if (meta.keywords) {
        let metaKeywords = document.querySelector('meta[name="keywords"]')
        if (!metaKeywords) {
          metaKeywords = document.createElement('meta')
          metaKeywords.name = 'keywords'
          document.head.appendChild(metaKeywords)
        }
        metaKeywords.content = meta.keywords
      }
    }
  }
  
  // 网络状态管理
  const network = {
    isOnline: isOnline,
    
    // 检查网络连接
    checkConnection: async () => {
      try {
        const response = await fetch('/api/ping', {
          method: 'HEAD',
          cache: 'no-cache'
        })
        const online = response.ok
        store.setOnlineStatus(online)
        return online
      } catch (error) {
        store.setOnlineStatus(false)
        return false
      }
    },
    
    // 网络状态变化回调
    onStatusChange: (callback) => {
      const handleOnline = () => {
        store.setOnlineStatus(true)
        callback(true)
      }
      
      const handleOffline = () => {
        store.setOnlineStatus(false)
        callback(false)
      }
      
      window.addEventListener('online', handleOnline)
      window.addEventListener('offline', handleOffline)
      
      // 返回清理函数
      return () => {
        window.removeEventListener('online', handleOnline)
        window.removeEventListener('offline', handleOffline)
      }
    }
  }
  
  // 性能监控
  const performance = {
    // 记录性能指标
    mark: (name) => {
      if (window.performance && window.performance.mark) {
        window.performance.mark(name)
      }
    },
    
    // 测量性能
    measure: (name, startMark, endMark) => {
      if (window.performance && window.performance.measure) {
        try {
          window.performance.measure(name, startMark, endMark)
          const measure = window.performance.getEntriesByName(name)[0]
          return measure.duration
        } catch (error) {
          console.warn('性能测量失败:', error)
          return null
        }
      }
      return null
    },
    
    // 获取页面加载性能
    getPageLoadMetrics: () => {
      if (!window.performance || !window.performance.timing) {
        return null
      }
      
      const timing = window.performance.timing
      return {
        domContentLoaded: timing.domContentLoadedEventEnd - timing.navigationStart,
        loadComplete: timing.loadEventEnd - timing.navigationStart,
        domReady: timing.domComplete - timing.navigationStart,
        firstPaint: timing.responseEnd - timing.navigationStart
      }
    }
  }
  
  // 错误处理
  const error = {
    // 全局错误处理
    handle: (error, context = '') => {
      console.error('应用错误:', error, context)
      
      // 可以发送错误到监控服务
      // errorReportingService.report(error, context)
      
      // 显示用户友好的错误消息
      store.addNotification({
        type: 'error',
        title: '发生错误',
        message: '应用遇到了一个问题，请刷新页面重试',
        duration: 5000
      })
    },
    
    // 异步错误处理
    handleAsync: async (asyncFn, fallback = null) => {
      try {
        return await asyncFn()
      } catch (error) {
        error.handle(error, 'async operation')
        return fallback
      }
    }
  }
  
  // 数据持久化
  const storage = {
    save: () => store.saveToLocalStorage(),
    load: () => store.loadFromLocalStorage(),
    clear: () => store.clearLocalStorage(),
    
    // 自动保存
    enableAutoSave: (interval = 30000) => {
      const autoSaveInterval = setInterval(() => {
        store.saveToLocalStorage()
      }, interval)
      
      return () => clearInterval(autoSaveInterval)
    }
  }
  
  // 调试工具
  const debug = {
    getState: () => store.getDebugInfo(),
    logState: () => console.log('应用状态:', store.getDebugInfo()),
    
    // 性能分析
    profileComponent: (componentName, fn) => {
      const startTime = performance.now()
      const result = fn()
      const endTime = performance.now()
      console.log(`组件 ${componentName} 执行时间:`, endTime - startTime, 'ms')
      return result
    }
  }
  
  // 初始化应用
  const initialize = () => {
    let cleanup = null
    
    onMounted(() => {
      cleanup = store.initialize()
      
      // 记录应用启动
      performance.mark('app-start')
      
      // 页面访问统计
      page.visit()
    })
    
    onUnmounted(() => {
      if (cleanup) {
        cleanup()
      }
    })
  }
  
  return {
    // 状态
    appInfo,
    isLoading,
    isOnline,
    isSidebarCollapsed,
    modals,
    stats,
    sessionInfo,
    
    // 应用控制
    setLoading,
    toggleSidebar,
    
    // 模态框
    modal,
    
    // 页面管理
    page,
    
    // 网络状态
    network,
    
    // 性能监控
    performance,
    
    // 错误处理
    error,
    
    // 数据持久化
    storage,
    
    // 调试工具
    debug,
    
    // 初始化
    initialize
  }
}