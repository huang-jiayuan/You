import { createApp } from 'vue'
import App from './App.vue'
import router from './router/temp.js'
import './assets/styles/main.css'
import './styles/mobile-design-system.css'
import { getGlobalAppStore } from './stores/app.js'

// 创建Vue应用实例
const app = createApp(App)

// 使用路由
app.use(router)

// 初始化全局状态管理
const globalStore = getGlobalAppStore()
const cleanupStore = globalStore.initialize()

// 全局错误处理
app.config.errorHandler = (err, vm, info) => {
  console.error('Vue应用错误:', err)
  console.error('错误信息:', info)
  console.error('组件实例:', vm)
  
  // 添加错误通知
  globalStore.addNotification({
    type: 'error',
    title: 'Vue应用错误',
    message: err.message || '应用发生未知错误',
    duration: 5000
  })
  
  // 在开发环境中显示更详细的错误信息
  if (import.meta.env.DEV) {
    console.error('错误堆栈:', err.stack)
  }
}

// 全局警告处理
app.config.warnHandler = (msg, vm, trace) => {
  console.warn('Vue警告:', msg)
  console.warn('组件追踪:', trace)
  
  // 在开发环境中添加警告通知
  if (import.meta.env.DEV) {
    globalStore.addNotification({
      type: 'warning',
      title: 'Vue警告',
      message: msg,
      duration: 3000
    })
  }
}

// 全局属性配置
app.config.globalProperties.$appName = 'Vue3 前端应用'
app.config.globalProperties.$version = '1.0.0'
app.config.globalProperties.$store = globalStore

// 性能追踪（仅在开发环境）
if (import.meta.env.DEV) {
  app.config.performance = true
}

// 路由守卫 - 页面访问统计
router.beforeEach((to, from, next) => {
  // 页面访问统计
  globalStore.incrementPageViews()
  
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - Vue3 前端应用`
  } else {
    document.title = 'Vue3 前端应用'
  }
  
  // 设置页面描述
  if (to.meta.description) {
    const metaDescription = document.querySelector('meta[name="description"]')
    if (metaDescription) {
      metaDescription.setAttribute('content', to.meta.description)
    }
  }
  
  // 开发环境下的路由日志
  if (import.meta.env.DEV) {
    console.log(`🧭 路由导航: ${from.path} -> ${to.path}`)
  }
  
  next()
})

// 挂载应用
app.mount('#app')

// 页面卸载时清理资源
window.addEventListener('beforeunload', () => {
  globalStore.saveToLocalStorage()
  if (cleanupStore) {
    cleanupStore()
  }
})

// 开发环境下的调试信息
if (import.meta.env.DEV) {
  console.log('🚀 Vue3应用已启动')
  console.log('📦 应用版本:', app.config.globalProperties.$version)
  console.log('🔧 开发模式已启用')
  console.log('🗃️ 全局状态管理已初始化')
  
  // 将全局store暴露到window对象，方便调试
  window.__APP_STORE__ = globalStore
}