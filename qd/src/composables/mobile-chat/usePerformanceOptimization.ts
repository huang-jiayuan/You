/**
 * 移动端性能优化组合式函数
 * 用于优化移动端应用的性能表现
 */

import { ref, onMounted, onUnmounted, nextTick } from 'vue'

interface PerformanceMetrics {
  loadTime: number
  renderTime: number
  memoryUsage: number
  fps: number
}

export function usePerformanceOptimization() {
  const metrics = ref<PerformanceMetrics>({
    loadTime: 0,
    renderTime: 0,
    memoryUsage: 0,
    fps: 0
  })

  const isVisible = ref(true)
  const isOnline = ref(navigator.onLine)

  // 图片懒加载
  const createLazyImageObserver = () => {
    if (!('IntersectionObserver' in window)) {
      return null
    }

    return new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const img = entry.target as HTMLImageElement
          const src = img.dataset.src
          
          if (src) {
            img.src = src
            img.removeAttribute('data-src')
            img.classList.remove('lazy')
            img.classList.add('loaded')
          }
          
          observer?.unobserve(img)
        }
      })
    }, {
      rootMargin: '50px 0px',
      threshold: 0.1
    })
  }

  let observer: IntersectionObserver | null = null

  // 初始化懒加载
  const initLazyLoading = () => {
    observer = createLazyImageObserver()
    
    if (observer) {
      const lazyImages = document.querySelectorAll('img[data-src]')
      lazyImages.forEach((img) => observer?.observe(img))
    }
  }

  // 添加懒加载图片
  const addLazyImage = (img: HTMLImageElement) => {
    if (observer) {
      observer.observe(img)
    }
  }

  // 页面可见性检测
  const handleVisibilityChange = () => {
    isVisible.value = !document.hidden
    
    // 页面不可见时暂停不必要的操作
    if (!isVisible.value) {
      // 暂停动画
      document.body.classList.add('page-hidden')
    } else {
      // 恢复动画
      document.body.classList.remove('page-hidden')
    }
  }

  // 网络状态检测
  const handleOnlineStatus = () => {
    isOnline.value = navigator.onLine
    
    if (!isOnline.value) {
      // 离线时的处理
      console.log('应用已离线')
    } else {
      // 重新上线时的处理
      console.log('应用已重新上线')
    }
  }

  // 防抖函数
  const debounce = <T extends (...args: any[]) => any>(
    func: T,
    wait: number
  ): ((...args: Parameters<T>) => void) => {
    let timeout: NodeJS.Timeout | null = null
    
    return (...args: Parameters<T>) => {
      if (timeout) {
        clearTimeout(timeout)
      }
      
      timeout = setTimeout(() => {
        func.apply(null, args)
      }, wait)
    }
  }

  // 节流函数
  const throttle = <T extends (...args: any[]) => any>(
    func: T,
    limit: number
  ): ((...args: Parameters<T>) => void) => {
    let inThrottle: boolean = false
    
    return (...args: Parameters<T>) => {
      if (!inThrottle) {
        func.apply(null, args)
        inThrottle = true
        setTimeout(() => inThrottle = false, limit)
      }
    }
  }

  // 预加载关键资源
  const preloadResource = (url: string, type: 'image' | 'audio' | 'video' = 'image') => {
    return new Promise((resolve, reject) => {
      let element: HTMLImageElement | HTMLAudioElement | HTMLVideoElement
      
      switch (type) {
        case 'image':
          element = new Image()
          break
        case 'audio':
          element = new Audio()
          break
        case 'video':
          element = document.createElement('video')
          break
      }
      
      element.onload = () => resolve(element)
      element.onerror = reject
      
      if (element instanceof HTMLImageElement) {
        element.src = url
      } else if (element instanceof HTMLAudioElement || element instanceof HTMLVideoElement) {
        element.src = url
        element.load()
      }
    })
  }

  // 批量预加载资源
  const preloadResources = async (urls: string[], type: 'image' | 'audio' | 'video' = 'image') => {
    const promises = urls.map(url => preloadResource(url, type))
    
    try {
      await Promise.all(promises)
      console.log(`成功预加载 ${urls.length} 个${type}资源`)
    } catch (error) {
      console.warn('部分资源预加载失败:', error)
    }
  }

  // 内存使用监控
  const monitorMemoryUsage = () => {
    if ('memory' in performance) {
      const memory = (performance as any).memory
      metrics.value.memoryUsage = memory.usedJSHeapSize / 1024 / 1024 // MB
    }
  }

  // FPS监控
  let fpsCounter = 0
  let lastTime = performance.now()
  
  const measureFPS = () => {
    const now = performance.now()
    const delta = now - lastTime
    
    if (delta >= 1000) {
      metrics.value.fps = Math.round((fpsCounter * 1000) / delta)
      fpsCounter = 0
      lastTime = now
    }
    
    fpsCounter++
    requestAnimationFrame(measureFPS)
  }

  // 优化滚动性能
  const optimizeScrolling = () => {
    const scrollElements = document.querySelectorAll('[data-scroll-optimize]')
    
    scrollElements.forEach((element) => {
      element.addEventListener('scroll', throttle(() => {
        // 滚动时的优化处理
        element.classList.add('scrolling')
        
        setTimeout(() => {
          element.classList.remove('scrolling')
        }, 150)
      }, 16)) // 60fps
    })
  }

  // 清理函数
  const cleanup = () => {
    if (observer) {
      observer.disconnect()
    }
    
    document.removeEventListener('visibilitychange', handleVisibilityChange)
    window.removeEventListener('online', handleOnlineStatus)
    window.removeEventListener('offline', handleOnlineStatus)
  }

  // 初始化
  onMounted(async () => {
    await nextTick()
    
    // 记录加载时间
    metrics.value.loadTime = performance.now()
    
    // 初始化各种优化
    initLazyLoading()
    optimizeScrolling()
    
    // 开始监控
    measureFPS()
    monitorMemoryUsage()
    
    // 设置定时监控
    const monitorInterval = setInterval(() => {
      monitorMemoryUsage()
    }, 5000)
    
    // 添加事件监听器
    document.addEventListener('visibilitychange', handleVisibilityChange)
    window.addEventListener('online', handleOnlineStatus)
    window.addEventListener('offline', handleOnlineStatus)
    
    // 清理定时器
    onUnmounted(() => {
      clearInterval(monitorInterval)
      cleanup()
    })
  })

  onUnmounted(() => {
    cleanup()
  })

  return {
    // 状态
    metrics,
    isVisible,
    isOnline,
    
    // 方法
    addLazyImage,
    debounce,
    throttle,
    preloadResource,
    preloadResources,
    cleanup
  }
}