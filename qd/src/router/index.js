import { createRouter, createWebHistory } from 'vue-router'
import { getGlobalAppStore } from '../stores/app.js'

// 导入页面组件
import Login from '../views/Login.vue'
import PhoneLogin from '../views/PhoneLogin.vue'
import PasswordLogin from '../views/PasswordLogin.vue'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import Contact from '../views/Contact.vue'
import Chat from '../views/Chat.vue'
import Room from '../views/Room.vue'
import ApiTest from '../views/ApiTest.vue'

// 定义路由配置
const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login,
    meta: {
      title: '登录',
      description: '声音交友登录页面',
      keywords: '登录, 声音交友, ME',
      requiresAuth: false,
      keepAlive: false,
      transition: 'fade',
      icon: '🔐'
    }
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      title: '首页',
      description: 'Vue3前端应用首页，展示现代Vue3开发特性',
      keywords: 'Vue3, 前端, 应用, 首页',
      requiresAuth: true,
      keepAlive: true,
      transition: 'fade',
      icon: '🏠'
    }
  },
  {
    path: '/about',
    name: 'About',
    component: About,
    meta: {
      title: '关于我们',
      description: '了解Vue3前端应用项目的技术栈和核心特性',
      keywords: 'Vue3, 关于, 技术栈, 特性',
      requiresAuth: false,
      keepAlive: true,
      transition: 'slide-left',
      icon: 'ℹ️'
    }
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact,
    meta: {
      title: '联系我们',
      description: '与我们取得联系，获取技术支持和帮助',
      keywords: 'Vue3, 联系, 支持, 帮助',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-right',
      icon: '📞'
    }
  },
  {
    path: '/chat',
    name: 'Chat',
    component: Chat,
    meta: {
      title: '聊天',
      description: '移动端聊天界面，体验现代化的即时通讯功能',
      keywords: 'Vue3, 聊天, 移动端, 即时通讯',
      requiresAuth: true,
      keepAlive: true,
      transition: 'slide-up',
      icon: '💬'
    }
  },
  {
    path: '/phone-login',
    name: 'PhoneLogin',
    component: PhoneLogin,
    meta: {
      title: '手机验证码登录',
      description: '使用手机验证码登录',
      keywords: '手机登录, 验证码, 登录',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-left',
      icon: '📱'
    }
  },
  {
    path: '/password-login',
    name: 'PasswordLogin',
    component: PasswordLogin,
    meta: {
      title: '密码登录',
      description: '使用密码登录',
      keywords: '密码登录, 登录',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-left',
      icon: '🔑'
    }
  },
  {
    path: '/api-test',
    name: 'ApiTest',
    component: ApiTest,
    meta: {
      title: 'API测试',
      description: '前后端联调测试页面',
      keywords: 'API, 测试, 联调',
      requiresAuth: false,
      keepAlive: false,
      transition: 'fade',
      icon: '🔧'
    }
  },
  // 404页面
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue'),
    meta: {
      title: '页面未找到',
      description: '您访问的页面不存在',
      requiresAuth: false,
      keepAlive: false,
      transition: 'fade'
    }
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  // 路由切换时滚动行为
  scrollBehavior(to, from, savedPosition) {
    // 如果有保存的位置（浏览器前进/后退）
    if (savedPosition) {
      return savedPosition
    }
    
    // 如果有锚点
    if (to.hash) {
      return {
        el: to.hash,
        behavior: 'smooth'
      }
    }
    
    // 默认滚动到顶部
    return { 
      top: 0,
      behavior: 'smooth'
    }
  }
})

// 路由加载状态管理
let loadingTimer = null

// 全局前置守卫
router.beforeEach(async (to, from, next) => {
  try {
    // 获取全局状态管理
    const store = getGlobalAppStore()
    
    // 开始加载状态
    if (loadingTimer) {
      clearTimeout(loadingTimer)
    }
    
    // 延迟显示加载状态，避免快速切换时的闪烁
    loadingTimer = setTimeout(() => {
      store.setLoading(true)
    }, 200)
    
    // 页面访问统计
    store.incrementPageViews()
    
    // 权限检查
    if (to.meta.requiresAuth) {
      if (!store.isLoggedIn.value) {
        // 未登录用户重定向到登录页
        next({ name: 'Login', query: { redirect: to.fullPath } })
        return
      }
    }
    
    // 设置页面标题
    if (to.meta.title) {
      document.title = `${to.meta.title} - Vue3 前端应用`
    } else {
      document.title = 'Vue3 前端应用'
    }
    
    // 设置页面元数据
    setPageMeta(to.meta)
    
    // 添加页面切换通知
    if (from.name && to.name !== from.name) {
      store.addNotification({
        type: 'info',
        title: '页面切换',
        message: `正在前往${to.meta.title || to.name}`,
        duration: 1500
      })
    }
    
    // 开发环境下的路由日志
    if (import.meta.env.DEV) {
      console.log(`🧭 路由导航: ${from.path} -> ${to.path}`)
      console.log('路由元数据:', to.meta)
    }
    
    next()
  } catch (error) {
    console.error('路由守卫错误:', error)
    next(false) // 取消导航
  }
})

// 全局解析守卫
router.beforeResolve(async (to, from, next) => {
  // 在导航被确认之前，同时在所有组件内守卫和异步路由组件被解析之后调用
  try {
    // 这里可以进行一些异步操作，比如预加载数据
    if (to.meta.preload) {
      // 预加载数据的逻辑
      await preloadRouteData(to)
    }
    
    next()
  } catch (error) {
    console.error('路由解析错误:', error)
    next(false)
  }
})

// 全局后置钩子
router.afterEach((to, from, failure) => {
  // 清除加载状态
  if (loadingTimer) {
    clearTimeout(loadingTimer)
    loadingTimer = null
  }
  
  const store = getGlobalAppStore()
  store.setLoading(false)
  
  // 如果导航失败
  if (failure) {
    console.error('路由导航失败:', failure)
    store.addNotification({
      type: 'error',
      title: '导航失败',
      message: '页面加载失败，请重试',
      duration: 3000
    })
    return
  }
  
  // 路由切换完成后的处理
  if (import.meta.env.DEV) {
    console.log(`✅ 路由切换完成: ${to.path}`)
  }
  
  // 更新用户活动
  store.updateActivity()
  
  // 页面性能监控
  if (window.performance && window.performance.mark) {
    window.performance.mark(`route-${to.name}-loaded`)
  }
})

// 全局错误处理
router.onError((error) => {
  console.error('🚨 路由错误:', error)
  
  // 清除加载状态
  if (loadingTimer) {
    clearTimeout(loadingTimer)
    loadingTimer = null
  }
  
  const store = getGlobalAppStore()
  store.setLoading(false)
  
  // 添加错误通知
  store.addNotification({
    type: 'error',
    title: '路由错误',
    message: '页面加载出现问题，请刷新重试',
    duration: 5000
  })
  
  // 在开发环境中显示更详细的错误信息
  if (import.meta.env.DEV) {
    console.error('错误堆栈:', error.stack)
  }
  
  // 错误上报（生产环境）
  if (import.meta.env.PROD) {
    // 发送错误信息到监控服务
    reportError(error, 'router')
  }
})

// 辅助函数：设置页面元数据
function setPageMeta(meta) {
  // 设置页面描述
  if (meta.description) {
    let metaDescription = document.querySelector('meta[name="description"]')
    if (!metaDescription) {
      metaDescription = document.createElement('meta')
      metaDescription.name = 'description'
      document.head.appendChild(metaDescription)
    }
    metaDescription.content = meta.description
  }
  
  // 设置关键词
  if (meta.keywords) {
    let metaKeywords = document.querySelector('meta[name="keywords"]')
    if (!metaKeywords) {
      metaKeywords = document.createElement('meta')
      metaKeywords.name = 'keywords'
      document.head.appendChild(metaKeywords)
    }
    metaKeywords.content = meta.keywords
  }
  
  // 设置Open Graph标签
  if (meta.title) {
    let ogTitle = document.querySelector('meta[property="og:title"]')
    if (!ogTitle) {
      ogTitle = document.createElement('meta')
      ogTitle.setAttribute('property', 'og:title')
      document.head.appendChild(ogTitle)
    }
    ogTitle.content = meta.title
  }
  
  if (meta.description) {
    let ogDescription = document.querySelector('meta[property="og:description"]')
    if (!ogDescription) {
      ogDescription = document.createElement('meta')
      ogDescription.setAttribute('property', 'og:description')
      document.head.appendChild(ogDescription)
    }
    ogDescription.content = meta.description
  }
}

// 辅助函数：预加载路由数据
async function preloadRouteData(to) {
  // 这里可以根据路由预加载数据
  console.log('预加载路由数据:', to.name)
  // 模拟异步数据加载
  await new Promise(resolve => setTimeout(resolve, 100))
}

// 辅助函数：错误上报
function reportError(error, context) {
  // 这里可以集成错误监控服务，如Sentry
  console.log('上报错误:', { error, context })
}

export default router