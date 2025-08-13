import { createRouter, createWebHistory } from 'vue-router'

// 导入页面组件
import Login from '../views/Login.vue'
import PhoneLogin from '../views/PhoneLogin.vue'
import PasswordLogin from '../views/PasswordLogin.vue'
import ResetPassword from '../views/ResetPassword.vue'
import Home from '../views/Home.vue'
import Chat from '../views/Chat.vue'
import About from '../views/About.vue'
import Contact from '../views/Contact.vue'

// 定义路由配置
const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login,
    meta: {
      title: '登录',
      description: '声音交友登录页面',
      keywords: '登录, 声音交友, You',
      requiresAuth: false,
      keepAlive: false,
      transition: 'fade',
      icon: '🔐'
    }
  },
  {
    path: '/phone-login',
    name: 'PhoneLogin',
    component: PhoneLogin,
    meta: {
      title: '手机验证码登录',
      description: '手机验证码登录页面',
      keywords: '手机登录, 验证码, You',
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
      description: '密码登录页面',
      keywords: '密码登录, You',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-right',
      icon: '🔑'
    }
  },
  {
    path: '/reset-password',
    name: 'ResetPassword',
    component: ResetPassword,
    meta: {
      title: '修改密码',
      description: '修改密码页面',
      keywords: '修改密码, 找回密码, You',
      requiresAuth: false,
      keepAlive: false,
      transition: 'slide-up',
      icon: '🔄'
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
    path: '/chat',
    name: 'Chat',
    component: Chat,
    meta: {
      title: '聊天',
      description: '声音交友聊天页面，与好友畅聊',
      keywords: '聊天, 消息, 声音交友, You',
      requiresAuth: true,
      keepAlive: true,
      transition: 'slide-left',
      icon: '💬'
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
  routes
})

export default router