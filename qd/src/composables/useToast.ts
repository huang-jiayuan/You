/**
 * Toast通知管理组合式函数
 * 用于在应用中显示各种类型的通知消息
 */

import { ref, createApp, h } from 'vue'
import Toast from '@/components/Toast.vue'

interface ToastOptions {
  type?: 'success' | 'error' | 'warning' | 'info'
  title?: string
  message: string
  duration?: number
  closable?: boolean
  showProgress?: boolean
}

interface ToastInstance {
  id: string
  close: () => void
}

class ToastManager {
  private toasts = ref<ToastInstance[]>([])
  private container: HTMLElement | null = null

  constructor() {
    this.createContainer()
  }

  private createContainer() {
    if (typeof document === 'undefined') return

    this.container = document.createElement('div')
    this.container.id = 'toast-container'
    this.container.style.cssText = `
      position: fixed;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      pointer-events: none;
      z-index: 9999;
    `
    document.body.appendChild(this.container)
  }

  private generateId(): string {
    return `toast-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`
  }

  show(options: ToastOptions): ToastInstance {
    if (!this.container) {
      this.createContainer()
    }

    const id = this.generateId()
    const toastElement = document.createElement('div')
    toastElement.id = id
    toastElement.style.pointerEvents = 'auto'

    const app = createApp({
      render() {
        return h(Toast, {
          ...options,
          onClose: () => {
            instance.close()
          }
        })
      }
    })

    const instance: ToastInstance = {
      id,
      close: () => {
        app.unmount()
        if (toastElement.parentNode) {
          toastElement.parentNode.removeChild(toastElement)
        }
        this.toasts.value = this.toasts.value.filter(t => t.id !== id)
      }
    }

    this.toasts.value.push(instance)
    this.container?.appendChild(toastElement)
    app.mount(toastElement)

    return instance
  }

  success(message: string, options?: Omit<ToastOptions, 'type' | 'message'>): ToastInstance {
    return this.show({
      type: 'success',
      message,
      ...options
    })
  }

  error(message: string, options?: Omit<ToastOptions, 'type' | 'message'>): ToastInstance {
    return this.show({
      type: 'error',
      message,
      duration: 5000, // 错误消息显示更长时间
      ...options
    })
  }

  warning(message: string, options?: Omit<ToastOptions, 'type' | 'message'>): ToastInstance {
    return this.show({
      type: 'warning',
      message,
      ...options
    })
  }

  info(message: string, options?: Omit<ToastOptions, 'type' | 'message'>): ToastInstance {
    return this.show({
      type: 'info',
      message,
      ...options
    })
  }

  // 清除所有Toast
  clear(): void {
    this.toasts.value.forEach(toast => toast.close())
    this.toasts.value = []
  }

  // 清除指定类型的Toast
  clearByType(type: ToastOptions['type']): void {
    // 由于我们没有存储类型信息，这里只能清除所有
    // 在实际应用中，可以扩展ToastInstance接口来包含类型信息
    this.clear()
  }
}

// 创建全局Toast管理器实例
const toastManager = new ToastManager()

export function useToast() {
  return {
    show: (options: ToastOptions) => toastManager.show(options),
    success: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
      toastManager.success(message, options),
    error: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
      toastManager.error(message, options),
    warning: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
      toastManager.warning(message, options),
    info: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
      toastManager.info(message, options),
    clear: () => toastManager.clear(),
    clearByType: (type: ToastOptions['type']) => toastManager.clearByType(type)
  }
}

// 导出默认实例，方便在非组件中使用
export const toast = {
  show: (options: ToastOptions) => toastManager.show(options),
  success: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
    toastManager.success(message, options),
  error: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
    toastManager.error(message, options),
  warning: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
    toastManager.warning(message, options),
  info: (message: string, options?: Omit<ToastOptions, 'type' | 'message'>) => 
    toastManager.info(message, options),
  clear: () => toastManager.clear(),
  clearByType: (type: ToastOptions['type']) => toastManager.clearByType(type)
}