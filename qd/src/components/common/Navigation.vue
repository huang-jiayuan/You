<template>
  <nav class="navigation">
    <div class="container">
      <div class="nav__content">
        <!-- 品牌标识 -->
        <div class="nav__brand">
          <router-link to="/" class="nav__brand-link">
            <h1 class="nav__brand-title">Vue3 App</h1>
          </router-link>
        </div>

        <!-- 桌面端导航菜单 -->
        <div class="nav__menu" :class="{ 'nav__menu--open': isMenuOpen }">
          <router-link 
            v-for="item in navItems" 
            :key="item.path"
            :to="item.path" 
            class="nav__link" 
            active-class="nav__link--active"
            @click="closeMenu"
          >
            {{ item.name }}
          </router-link>
          
          <!-- 主题切换按钮 -->
          <button @click="toggleTheme" class="nav__theme-toggle" :title="themeToggleTitle">
            {{ currentTheme === 'light' ? '🌙' : '☀️' }}
          </button>
        </div>

        <!-- 移动端汉堡菜单按钮 -->
        <button 
          class="nav__mobile-toggle" 
          @click="toggleMenu"
          :aria-expanded="isMenuOpen"
          aria-label="切换导航菜单"
        >
          <span class="nav__hamburger"></span>
          <span class="nav__hamburger"></span>
          <span class="nav__hamburger"></span>
        </button>
      </div>
    </div>
  </nav>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'

export default {
  name: 'Navigation',
  setup() {
    const isMenuOpen = ref(false)
    const currentTheme = ref('light')
    
    // 导航项配置
    const navItems = [
      { path: '/', name: '首页' },
      { path: '/about', name: '关于' },
      { path: '/contact', name: '联系' }
    ]

    // 计算属性
    const themeToggleTitle = computed(() => {
      return currentTheme.value === 'light' ? '切换到暗色主题' : '切换到亮色主题'
    })

    // 方法
    const toggleMenu = () => {
      isMenuOpen.value = !isMenuOpen.value
    }

    const closeMenu = () => {
      isMenuOpen.value = false
    }

    const toggleTheme = () => {
      currentTheme.value = currentTheme.value === 'light' ? 'dark' : 'light'
      localStorage.setItem('theme', currentTheme.value)
      document.documentElement.setAttribute('data-theme', currentTheme.value)
    }

    // 处理窗口大小变化
    const handleResize = () => {
      if (window.innerWidth > 768) {
        isMenuOpen.value = false
      }
    }

    // 处理点击外部关闭菜单
    const handleClickOutside = (event) => {
      const nav = event.target.closest('.navigation')
      if (!nav && isMenuOpen.value) {
        closeMenu()
      }
    }

    // 生命周期
    onMounted(() => {
      // 初始化主题
      const savedTheme = localStorage.getItem('theme')
      if (savedTheme) {
        currentTheme.value = savedTheme
        document.documentElement.setAttribute('data-theme', savedTheme)
      }

      // 添加事件监听器
      window.addEventListener('resize', handleResize)
      document.addEventListener('click', handleClickOutside)
    })

    onUnmounted(() => {
      // 清理事件监听器
      window.removeEventListener('resize', handleResize)
      document.removeEventListener('click', handleClickOutside)
    })

    return {
      isMenuOpen,
      currentTheme,
      navItems,
      themeToggleTitle,
      toggleMenu,
      closeMenu,
      toggleTheme
    }
  }
}
</script>

<style scoped>
.navigation {
  background-color: var(--color-bg-primary);
  border-bottom: 1px solid var(--color-border);
  position: sticky;
  top: 0;
  z-index: 100;
  backdrop-filter: blur(10px);
}

.nav__content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) 0;
}

.nav__brand-link {
  text-decoration: none;
  color: inherit;
}

.nav__brand-title {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--color-primary);
  margin: 0;
  transition: color var(--transition-fast);
}

.nav__brand-link:hover .nav__brand-title {
  color: var(--color-primary-dark);
}

.nav__menu {
  display: flex;
  align-items: center;
  gap: var(--spacing-lg);
}

.nav__link {
  color: var(--color-text-secondary);
  font-weight: 500;
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--border-radius-md);
  transition: all var(--transition-fast);
  text-decoration: none;
}

.nav__link:hover,
.nav__link--active {
  color: var(--color-primary);
  background-color: var(--color-bg-tertiary);
}

.nav__theme-toggle {
  background: none;
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-md);
  padding: var(--spacing-sm);
  font-size: var(--font-size-lg);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.nav__theme-toggle:hover {
  background-color: var(--color-bg-tertiary);
  border-color: var(--color-primary);
}

.nav__mobile-toggle {
  display: none;
  flex-direction: column;
  background: none;
  border: none;
  cursor: pointer;
  padding: var(--spacing-sm);
  gap: 4px;
}

.nav__hamburger {
  width: 24px;
  height: 2px;
  background-color: var(--color-text-primary);
  transition: all var(--transition-fast);
}

/* 移动端样式 */
@media (max-width: 768px) {
  .nav__mobile-toggle {
    display: flex;
  }

  .nav__menu {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background-color: var(--color-bg-primary);
    border-top: 1px solid var(--color-border);
    flex-direction: column;
    padding: var(--spacing-lg);
    gap: var(--spacing-md);
    transform: translateY(-100%);
    opacity: 0;
    visibility: hidden;
    transition: all var(--transition-normal);
  }

  .nav__menu--open {
    transform: translateY(0);
    opacity: 1;
    visibility: visible;
  }

  .nav__link {
    width: 100%;
    text-align: center;
    padding: var(--spacing-md);
  }

  .nav__theme-toggle {
    align-self: center;
  }

  /* 汉堡菜单动画 */
  .nav__mobile-toggle[aria-expanded="true"] .nav__hamburger:nth-child(1) {
    transform: rotate(45deg) translate(6px, 6px);
  }

  .nav__mobile-toggle[aria-expanded="true"] .nav__hamburger:nth-child(2) {
    opacity: 0;
  }

  .nav__mobile-toggle[aria-expanded="true"] .nav__hamburger:nth-child(3) {
    transform: rotate(-45deg) translate(6px, -6px);
  }
}
</style>