import { createRouter, createWebHistory } from 'vue-router'

// 导入页面组件
import Login from '../views/Login.vue'
import PhoneLogin from '../views/PhoneLogin.vue'
import PasswordLogin from '../views/PasswordLogin.vue'
import Home from '../views/Home.vue'
import Room from '../views/Room.vue'
import CreateRoom from '../views/CreateRoom.vue'
import About from '../views/About.vue'
import Contact from '../views/Contact.vue'
import Chat from '../views/Chat.vue'
// import ChatDetail from '../views/ChatDetailSimple.vue'
import ApiTest from '../views/ApiTest.vue'

// 定义路由配置
const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: '首页',
      description: '声音交友首页，发现有趣的人和房间',
      keywords: '声音交友, 首页, 房间, ME',
      requiresAuth: false,
      keepAlive: true,
      transition: 'fade',
      icon: '🏠'
    }
  },
  {
    path: '/login',
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
    path: '/room/:id',
    name: 'Room',
    component: Room,
    meta: {
      title: '语音房间',
      description: '语音聊天房间，与好友畅聊',
      keywords: '语音房间, 聊天, 麦位, ME',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-up',
      icon: '🎤'
    }
  },
  {
    path: '/create-room',
    name: 'CreateRoom',
    component: CreateRoom,
    meta: {
      title: '创建房间',
      description: '创建语音聊天房间',
      keywords: '创建房间, 语音房间, ME',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-up',
      icon: '🏗️'
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
    path: '/chat/:userId',
    name: 'ChatDetail',
    component: () => import('../views/ChatDetail.vue'),
    meta: {
      title: '聊天详情',
      description: '与好友进行私聊',
      keywords: 'Vue3, 聊天, 私聊, 消息',
      requiresAuth: true,
      keepAlive: false,
      transition: 'slide-left',
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
  {
    path: '/profile/:userId?',
    name: 'UserProfile',
    component: () => import('../views/UserProfile.vue'),
    meta: {
      title: '用户主页',
      description: '查看用户个人资料和动态',
      keywords: 'Vue3, 用户, 主页, 资料',
      requiresAuth: true,
      keepAlive: false,
      transition: 'slide-left',
      icon: '👤'
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
    // 权限检查
    if (to.meta.requiresAuth) {
      const token = localStorage.getItem('access_token')
      if (!token) {
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

// 全局后置钩子
router.afterEach((to, from, failure) => {
  // 清除加载状态
  if (loadingTimer) {
    clearTimeout(loadingTimer)
    loadingTimer = null
  }
  
  // 如果导航失败
  if (failure) {
    console.error('路由导航失败:', failure)
    return
  }
  
  // 路由切换完成后的处理
  if (import.meta.env.DEV) {
    console.log(`✅ 路由切换完成: ${to.path}`)
  }
  
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
  
  // 在开发环境中显示更详细的错误信息
  if (import.meta.env.DEV) {
    console.error('错误堆栈:', error.stack)
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



export default router