<template>
  <div id="app" :data-theme="currentTheme">


    <!-- 主要内容区域 -->
    <main class="app-main">
      <div class="container">
        <!-- 路由出口带过渡动画 -->
        <router-view v-slot="{ Component, route }">
          <transition 
            :name="route.meta.transition || 'fade'"
            mode="out-in"
            appear
          >
            <keep-alive :include="getKeepAliveComponents()">
              <component :is="Component" :key="route.path" />
            </keep-alive>
          </transition>
        </router-view>
      </div>
    </main>


  </div>
</template>

<script>
import { onMounted } from 'vue'
import { useTheme } from './composables/useTheme.js'
import { initErrorHandler } from './utils/errorHandler'

export default {
  name: 'App',
  components: {
  },
  setup() {
    // 使用主题管理
    const { theme, initTheme } = useTheme()

    // 获取需要缓存的组件列表
    const getKeepAliveComponents = () => {
      // 返回需要缓存的组件名称数组
      return ['Home', 'About'] // Contact页面包含表单，不缓存
    }

    // 初始化主题和错误处理
    onMounted(() => {
      initTheme()
      initErrorHandler()
    })

    return {
      currentTheme: theme,
      getKeepAliveComponents
    }
  }
}
</script>

<style scoped>
/* 应用根容器 */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* 主要内容区域 */
.app-main {
  flex: 1;
  padding: var(--spacing-2xl) 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .app-main {
    padding: var(--spacing-xl) 0;
  }
}

@media (max-width: 480px) {
  .app-main {
    padding: var(--spacing-lg) 0;
  }
}

/* 页面过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity var(--transition-normal);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-left-enter-active,
.slide-left-leave-active {
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.slide-left-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.slide-left-leave-to {
  opacity: 0;
  transform: translateX(-30%);
}

/* 聊天页面专用的滑动动画 */
.chat-slide-enter-active,
.chat-slide-leave-active {
  transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.chat-slide-enter-from {
  opacity: 0;
  transform: translateX(100%) scale(0.95);
}

.chat-slide-leave-to {
  opacity: 0;
  transform: translateX(-20%) scale(1.05);
}

.slide-right-enter-active,
.slide-right-leave-active {
  transition: all var(--transition-normal);
}

.slide-right-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all var(--transition-normal);
}

.slide-up-enter-from {
  opacity: 0;
  transform: translateY(30px);
}

.slide-up-leave-to {
  opacity: 0;
  transform: translateY(-30px);
}

.scale-enter-active,
.scale-leave-active {
  transition: all var(--transition-normal);
}

.scale-enter-from {
  opacity: 0;
  transform: scale(0.9);
}

.scale-leave-to {
  opacity: 0;
  transform: scale(1.1);
}
</style>