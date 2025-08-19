<template>
  <Teleport to="body">
    <Transition
      name="toast"
      enter-active-class="animate-slide-up"
      leave-active-class="animate-fade-out"
    >
      <div
        v-if="visible"
        class="toast-container"
        :class="[`toast-${type}`, { 'toast-mobile': isMobile }]"
        @click="close"
      >
        <div class="toast-content">
          <div class="toast-icon">
            <span v-if="type === 'success'">✅</span>
            <span v-else-if="type === 'error'">❌</span>
            <span v-else-if="type === 'warning'">⚠️</span>
            <span v-else-if="type === 'info'">ℹ️</span>
            <span v-else>📢</span>
          </div>
          <div class="toast-message">
            <div v-if="title" class="toast-title">{{ title }}</div>
            <div class="toast-text">{{ message }}</div>
          </div>
          <button v-if="closable" class="toast-close" @click.stop="close">
            ×
          </button>
        </div>
        <div v-if="showProgress" class="toast-progress">
          <div 
            class="toast-progress-bar" 
            :style="{ width: `${progress}%` }"
          ></div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'

export default {
  name: 'Toast',
  props: {
    type: {
      type: String,
      default: 'info',
      validator: (value) => ['success', 'error', 'warning', 'info'].includes(value)
    },
    title: {
      type: String,
      default: ''
    },
    message: {
      type: String,
      required: true
    },
    duration: {
      type: Number,
      default: 3000
    },
    closable: {
      type: Boolean,
      default: true
    },
    showProgress: {
      type: Boolean,
      default: true
    }
  },
  emits: ['close'],
  setup(props, { emit }) {
    const visible = ref(false)
    const progress = ref(100)
    
    let timer = null
    let progressTimer = null
    
    const isMobile = computed(() => {
      return window.innerWidth <= 768
    })
    
    const show = () => {
      visible.value = true
      
      if (props.duration > 0) {
        // 进度条动画
        if (props.showProgress) {
          const interval = 50
          const step = (interval / props.duration) * 100
          
          progressTimer = setInterval(() => {
            progress.value -= step
            if (progress.value <= 0) {
              clearInterval(progressTimer)
            }
          }, interval)
        }
        
        // 自动关闭
        timer = setTimeout(() => {
          close()
        }, props.duration)
      }
    }
    
    const close = () => {
      visible.value = false
      
      if (timer) {
        clearTimeout(timer)
        timer = null
      }
      
      if (progressTimer) {
        clearInterval(progressTimer)
        progressTimer = null
      }
      
      emit('close')
    }
    
    onMounted(() => {
      show()
    })
    
    onUnmounted(() => {
      if (timer) {
        clearTimeout(timer)
      }
      if (progressTimer) {
        clearInterval(progressTimer)
      }
    })
    
    return {
      visible,
      progress,
      isMobile,
      close
    }
  }
}
</script>

<style scoped>
.toast-container {
  position: fixed;
  top: 20px;
  right: 20px;
  min-width: 300px;
  max-width: 500px;
  background: var(--surface);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  overflow: hidden;
  z-index: var(--z-tooltip);
  cursor: pointer;
  transition: all var(--duration-normal) var(--ease-out);
}

.toast-container:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-xl);
}

.toast-mobile {
  left: 16px;
  right: 16px;
  top: 16px;
  min-width: auto;
  max-width: none;
}

.toast-content {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px;
}

.toast-icon {
  font-size: 20px;
  flex-shrink: 0;
  margin-top: 2px;
}

.toast-message {
  flex: 1;
  min-width: 0;
}

.toast-title {
  font-size: var(--font-size-md);
  font-weight: var(--font-weight-semibold);
  color: var(--text-primary);
  margin-bottom: 4px;
}

.toast-text {
  font-size: var(--font-size-base);
  color: var(--text-secondary);
  line-height: var(--line-height-normal);
  word-wrap: break-word;
}

.toast-close {
  background: none;
  border: none;
  font-size: 20px;
  color: var(--text-tertiary);
  cursor: pointer;
  padding: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all var(--duration-fast) var(--ease-out);
  flex-shrink: 0;
}

.toast-close:hover {
  background: rgba(0, 0, 0, 0.1);
  color: var(--text-primary);
}

.toast-progress {
  height: 3px;
  background: rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.toast-progress-bar {
  height: 100%;
  background: currentColor;
  transition: width linear;
}

/* 类型样式 */
.toast-success {
  border-left: 4px solid var(--success);
}

.toast-success .toast-progress-bar {
  background: var(--success);
}

.toast-error {
  border-left: 4px solid var(--error);
}

.toast-error .toast-progress-bar {
  background: var(--error);
}

.toast-warning {
  border-left: 4px solid var(--warning);
}

.toast-warning .toast-progress-bar {
  background: var(--warning);
}

.toast-info {
  border-left: 4px solid var(--info);
}

.toast-info .toast-progress-bar {
  background: var(--info);
}

/* 动画 */
.animate-slide-up {
  animation: slideUp var(--duration-normal) var(--ease-out);
}

.animate-fade-out {
  animation: fadeOut var(--duration-normal) var(--ease-out);
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes fadeOut {
  from {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
  to {
    opacity: 0;
    transform: translateY(-10px) scale(0.95);
  }
}

/* 移动端优化 */
@media (max-width: 480px) {
  .toast-container {
    font-size: var(--font-size-sm);
  }
  
  .toast-content {
    padding: 12px;
    gap: 10px;
  }
  
  .toast-icon {
    font-size: 18px;
  }
  
  .toast-title {
    font-size: var(--font-size-base);
  }
  
  .toast-text {
    font-size: var(--font-size-sm);
  }
}
</style>