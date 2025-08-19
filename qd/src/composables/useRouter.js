import { computed, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { getGlobalAppStore } from '../stores/app.js'

/**
 * 路由导航组合式函数
 * 提供增强的路由导航功能
 */
export function useRouterNavigation() {
  const router = useRouter()
  const route = useRoute()
  const store = getGlobalAppStore()
  
  // 导航历史记录
  const navigationHistory = ref([])
  const maxHistoryLength = 10
  
  // 计算属性
  const currentRoute = computed(() => route)
  const currentPath = computed(() => route.path)
  const currentName = computed(() => route.name)
  const routeMeta = computed(() => route.meta)
  
  // 面包屑导航
  const breadcrumbs = computed(() => {
    const matched = route.matched.filter(item => item.meta && item.meta.title)
    return matched.map(item => ({
      name: item.name,
      path: item.path,
      title: item.meta.title,
      icon: item.meta.icon
    }))
  })
  
  // 导航菜单项
  const navigationItems = computed(() => {
    return router.getRoutes()
      .filter(route => route.meta && route.meta.title && !route.meta.hidden)
      .map(route => ({
        name: route.name,
        path: route.path,
        title: route.meta.title,
        icon: route.meta.icon,
        requiresAuth: route.meta.requiresAuth,
        active: route.path === currentPath.value
      }))
      .sort((a, b) => (a.order || 0) - (b.order || 0))
  })
  
  // 导航方法
  const navigate = {
    // 基础导航
    to: (location, options = {}) => {
      const { replace = false, force = false } = options
      
      if (force || location !== currentPath.value) {
        addToHistory(currentRoute.value)
        
        if (replace) {
          return router.replace(location)
        } else {
          return router.push(location)
        }
      }
    },
    
    // 返回上一页
    back: () => {
      if (navigationHistory.value.length > 0) {
        const previousRoute = navigationHistory.value.pop()
        return router.push(previousRoute.path)
      } else if (window.history.length > 1) {
        return router.go(-1)
      } else {
        return router.push('/')
      }
    },
    
    // 前进
    forward: () => {
      return router.go(1)
    },
    
    // 刷新当前页面
    refresh: () => {
      return router.go(0)
    },
    
    // 导航到首页
    home: () => {
      return navigate.to('/')
    },
    
    // 外部链接
    external: (url, target = '_blank') => {
      window.open(url, target)
    }
  }
  
  // 路由守卫
  const guards = {
    // 添加路由守卫
    beforeEach: (guard) => {
      return router.beforeEach(guard)
    },
    
    // 添加解析守卫
    beforeResolve: (guard) => {
      return router.beforeResolve(guard)
    },
    
    // 添加后置钩子
    afterEach: (hook) => {
      return router.afterEach(hook)
    }
  }
  
  // 路由状态检查
  const is = {
    // 是否为当前路由
    current: (routeName) => {
      return currentName.value === routeName
    },
    
    // 是否为子路由
    child: (parentRoute) => {
      return currentPath.value.startsWith(parentRoute)
    },
    
    // 是否需要认证
    protected: () => {
      return routeMeta.value.requiresAuth === true
    },
    
    // 是否为首页
    home: () => {
      return currentPath.value === '/'
    }
  }
  
  // 路由参数管理
  const params = {
    // 获取路由参数
    get: (key) => {
      return route.params[key]
    },
    
    // 获取查询参数
    query: (key) => {
      return route.query[key]
    },
    
    // 更新查询参数
    updateQuery: (newQuery, options = {}) => {
      const { replace = false, merge = true } = options
      const query = merge ? { ...route.query, ...newQuery } : newQuery
      
      const location = {
        path: route.path,
        query
      }
      
      if (replace) {
        return router.replace(location)
      } else {
        return router.push(location)
      }
    },
    
    // 清除查询参数
    clearQuery: (keys = null) => {
      let query = { ...route.query }
      
      if (keys) {
        if (Array.isArray(keys)) {
          keys.forEach(key => delete query[key])
        } else {
          delete query[keys]
        }
      } else {
        query = {}
      }
      
      return router.replace({
        path: route.path,
        query
      })
    }
  }
  
  // 路由元数据管理
  const meta = {
    // 获取元数据
    get: (key) => {
      return routeMeta.value[key]
    },
    
    // 设置页面标题
    setTitle: (title) => {
      document.title = title ? `${title} - Vue3 前端应用` : 'Vue3 前端应用'
    },
    
    // 设置页面描述
    setDescription: (description) => {
      let metaDesc = document.querySelector('meta[name="description"]')
      if (!metaDesc) {
        metaDesc = document.createElement('meta')
        metaDesc.name = 'description'
        document.head.appendChild(metaDesc)
      }
      metaDesc.content = description
    }
  }
  
  // 历史记录管理
  const addToHistory = (route) => {
    // 避免重复添加相同路由
    if (navigationHistory.value.length === 0 || 
        navigationHistory.value[navigationHistory.value.length - 1].path !== route.path) {
      navigationHistory.value.push({
        name: route.name,
        path: route.path,
        title: route.meta?.title,
        timestamp: Date.now()
      })
      
      // 限制历史记录长度
      if (navigationHistory.value.length > maxHistoryLength) {
        navigationHistory.value.shift()
      }
      
      // 保存到localStorage
      saveHistoryToStorage()
    }
  }
  
  const saveHistoryToStorage = () => {
    try {
      localStorage.setItem('navigation-history', JSON.stringify(navigationHistory.value))
    } catch (error) {
      console.warn('无法保存导航历史:', error)
    }
  }
  
  const loadHistoryFromStorage = () => {
    try {
      const saved = localStorage.getItem('navigation-history')
      if (saved) {
        navigationHistory.value = JSON.parse(saved)
      }
    } catch (error) {
      console.warn('无法加载导航历史:', error)
    }
  }
  
  // 路由预加载
  const preload = {
    // 预加载路由组件
    component: async (routeName) => {
      const route = router.resolve({ name: routeName })
      if (route.matched.length > 0) {
        const component = route.matched[0].components?.default
        if (typeof component === 'function') {
          try {
            await component()
            console.log(`预加载组件成功: ${routeName}`)
          } catch (error) {
            console.warn(`预加载组件失败: ${routeName}`, error)
          }
        }
      }
    },
    
    // 预加载多个路由
    routes: async (routeNames) => {
      const promises = routeNames.map(name => preload.component(name))
      await Promise.allSettled(promises)
    }
  }
  
  // 路由分析
  const analytics = {
    // 记录页面访问
    trackPageView: () => {
      store.incrementPageViews()
      
      // 可以发送到分析服务
      if (window.gtag) {
        window.gtag('config', 'GA_MEASUREMENT_ID', {
          page_title: document.title,
          page_location: window.location.href
        })
      }
    },
    
    // 记录导航事件
    trackNavigation: (from, to) => {
      console.log('导航事件:', { from: from.path, to: to.path })
      
      // 可以发送到分析服务
      if (window.gtag) {
        window.gtag('event', 'page_view', {
          page_title: to.meta?.title || to.name,
          page_location: window.location.href,
          page_path: to.path
        })
      }
    }
  }
  
  // 监听路由变化
  watch(route, (to, from) => {
    if (from.name) {
      addToHistory(from)
      analytics.trackNavigation(from, to)
    }
    analytics.trackPageView()
  }, { immediate: true })
  
  // 初始化
  loadHistoryFromStorage()
  
  return {
    // 状态
    currentRoute,
    currentPath,
    currentName,
    routeMeta,
    breadcrumbs,
    navigationItems,
    navigationHistory: computed(() => navigationHistory.value),
    
    // 导航方法
    navigate,
    
    // 路由守卫
    guards,
    
    // 状态检查
    is,
    
    // 参数管理
    params,
    
    // 元数据管理
    meta,
    
    // 预加载
    preload,
    
    // 分析
    analytics
  }
}