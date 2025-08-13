<template>
  <div class="theme-selector">
    <div class="theme-selector__header">
      <h3 class="theme-selector__title">选择主题</h3>
      <p class="theme-selector__description">个性化您的界面体验</p>
    </div>
    
    <div class="theme-selector__grid">
      <div 
        v-for="themeOption in themeOptions" 
        :key="themeOption.value"
        class="theme-option"
        :class="{ 'theme-option--active': currentTheme === themeOption.value }"
        @click="selectTheme(themeOption.value)"
        :title="themeOption.description"
      >
        <div class="theme-option__preview" :data-theme="themeOption.value">
          <div class="theme-preview__header"></div>
          <div class="theme-preview__content">
            <div class="theme-preview__text"></div>
            <div class="theme-preview__text theme-preview__text--secondary"></div>
            <div class="theme-preview__button"></div>
          </div>
        </div>
        <div class="theme-option__info">
          <h4 class="theme-option__name">{{ themeOption.name }}</h4>
          <p class="theme-option__desc">{{ themeOption.description }}</p>
        </div>
        <div v-if="currentTheme === themeOption.value" class="theme-option__check">
          ✓
        </div>
      </div>
    </div>
    
    <div class="theme-selector__actions">
      <Button 
        type="secondary" 
        size="small"
        @click="resetToSystem"
      >
        跟随系统
      </Button>
      <Button 
        type="primary" 
        size="small"
        @click="savePreferences"
      >
        保存设置
      </Button>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useTheme } from '../../composables/useTheme.js'
import Button from './Button.vue'

export default {
  name: 'ThemeSelector',
  components: {
    Button
  },
  setup() {
    const { theme, setTheme, themes } = useTheme()

    // 主题选项配置
    const themeOptions = [
      {
        value: 'light',
        name: '明亮主题',
        description: '经典的明亮界面，适合日间使用',
        category: 'basic'
      },
      {
        value: 'dark',
        name: '暗色主题',
        description: '护眼的暗色界面，适合夜间使用',
        category: 'basic'
      },
      {
        value: 'blue',
        name: '蓝色主题',
        description: '专业的蓝色调，商务风格',
        category: 'color'
      },
      {
        value: 'green',
        name: '绿色主题',
        description: '清新的绿色调，自然风格',
        category: 'color'
      },
      {
        value: 'purple',
        name: '紫色主题',
        description: '优雅的紫色调，创意风格',
        category: 'color'
      },
      {
        value: 'red',
        name: '红色主题',
        description: '活力的红色调，热情风格',
        category: 'color'
      },
      {
        value: 'orange',
        name: '橙色主题',
        description: '温暖的橙色调，活泼风格',
        category: 'color'
      },
      {
        value: 'pink',
        name: '粉色主题',
        description: '温柔的粉色调，浪漫风格',
        category: 'color'
      },
      {
        value: 'cyan',
        name: '青色主题',
        description: '清爽的青色调，科技风格',
        category: 'color'
      },
      {
        value: 'high-contrast',
        name: '高对比度',
        description: '高对比度界面，提升可读性',
        category: 'accessibility'
      },
      {
        value: 'high-contrast-dark',
        name: '高对比度暗色',
        description: '暗色高对比度界面',
        category: 'accessibility'
      },
      {
        value: 'eye-care',
        name: '护眼主题',
        description: '温和的护眼色调，减少视觉疲劳',
        category: 'accessibility'
      }
    ]

    const currentTheme = computed(() => theme.value)

    // 方法
    const selectTheme = (themeValue) => {
      setTheme(themeValue)
    }

    const resetToSystem = () => {
      // 清除本地存储的主题设置，让系统跟随浏览器偏好
      localStorage.removeItem('theme')
      const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
      setTheme(systemTheme)
    }

    const savePreferences = () => {
      // 主题已经自动保存到localStorage，这里可以添加其他保存逻辑
      alert('主题设置已保存！')
    }

    return {
      themeOptions,
      currentTheme,
      selectTheme,
      resetToSystem,
      savePreferences
    }
  }
}
</script>

<style scoped>
.theme-selector {
  padding: var(--spacing-lg);
  background-color: var(--color-bg-primary);
  border-radius: var(--border-radius-lg);
  border: 1px solid var(--color-border);
}

.theme-selector__header {
  text-align: center;
  margin-bottom: var(--spacing-xl);
}

.theme-selector__title {
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-sm);
}

.theme-selector__description {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  margin: 0;
}

.theme-selector__grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
}

.theme-option {
  position: relative;
  padding: var(--spacing-md);
  border: 2px solid var(--color-border);
  border-radius: var(--border-radius-md);
  cursor: pointer;
  transition: all var(--transition-fast);
  background-color: var(--color-bg-secondary);
}

.theme-option:hover {
  border-color: var(--color-primary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.theme-option--active {
  border-color: var(--color-primary);
  background-color: var(--color-primary-bg);
}

.theme-option__preview {
  width: 100%;
  height: 80px;
  border-radius: var(--border-radius-sm);
  overflow: hidden;
  margin-bottom: var(--spacing-sm);
  border: 1px solid var(--color-border-light);
}

.theme-preview__header {
  height: 20px;
  background-color: var(--color-bg-secondary);
  border-bottom: 1px solid var(--color-border);
}

.theme-preview__content {
  padding: var(--spacing-sm);
  background-color: var(--color-bg-primary);
  height: calc(100% - 20px);
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.theme-preview__text {
  height: 8px;
  background-color: var(--color-text-primary);
  border-radius: var(--border-radius-sm);
  opacity: 0.8;
}

.theme-preview__text--secondary {
  background-color: var(--color-text-secondary);
  width: 70%;
  opacity: 0.6;
}

.theme-preview__button {
  height: 12px;
  width: 40px;
  background-color: var(--color-primary);
  border-radius: var(--border-radius-sm);
  margin-top: auto;
}

.theme-option__info {
  text-align: center;
}

.theme-option__name {
  font-size: var(--font-size-sm);
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.theme-option__desc {
  font-size: var(--font-size-xs);
  color: var(--color-text-secondary);
  margin: 0;
  line-height: 1.4;
}

.theme-option__check {
  position: absolute;
  top: var(--spacing-sm);
  right: var(--spacing-sm);
  width: 20px;
  height: 20px;
  background-color: var(--color-primary);
  color: var(--color-white);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--font-size-xs);
  font-weight: 700;
}

.theme-selector__actions {
  display: flex;
  justify-content: center;
  gap: var(--spacing-md);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--color-border);
}

/* 主题预览的特定样式 */
.theme-option__preview[data-theme="light"] {
  --color-bg-primary: #ffffff;
  --color-bg-secondary: #f9fafb;
  --color-text-primary: #1f2937;
  --color-text-secondary: #6b7280;
  --color-primary: #3b82f6;
  --color-border: #e5e7eb;
}

.theme-option__preview[data-theme="dark"] {
  --color-bg-primary: #111827;
  --color-bg-secondary: #1f2937;
  --color-text-primary: #f9fafb;
  --color-text-secondary: #d1d5db;
  --color-primary: #60a5fa;
  --color-border: #374151;
}

.theme-option__preview[data-theme="blue"] {
  --color-primary: #2563eb;
}

.theme-option__preview[data-theme="green"] {
  --color-primary: #059669;
}

.theme-option__preview[data-theme="purple"] {
  --color-primary: #7c3aed;
}

.theme-option__preview[data-theme="red"] {
  --color-primary: #dc2626;
}

.theme-option__preview[data-theme="orange"] {
  --color-primary: #ea580c;
}

.theme-option__preview[data-theme="pink"] {
  --color-primary: #db2777;
}

.theme-option__preview[data-theme="cyan"] {
  --color-primary: #0891b2;
}

.theme-option__preview[data-theme="high-contrast"] {
  --color-bg-primary: #ffffff;
  --color-bg-secondary: #ffffff;
  --color-text-primary: #000000;
  --color-text-secondary: #000000;
  --color-primary: #0000ff;
  --color-border: #000000;
}

.theme-option__preview[data-theme="high-contrast-dark"] {
  --color-bg-primary: #000000;
  --color-bg-secondary: #000000;
  --color-text-primary: #ffffff;
  --color-text-secondary: #ffffff;
  --color-primary: #00ffff;
  --color-border: #ffffff;
}

.theme-option__preview[data-theme="eye-care"] {
  --color-bg-primary: #f3e5ab;
  --color-bg-secondary: #f0e68c;
  --color-text-primary: #3e2723;
  --color-text-secondary: #5d4037;
  --color-primary: #8b5a3c;
  --color-border: #d4c5a9;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .theme-selector__grid {
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    gap: var(--spacing-sm);
  }
  
  .theme-option {
    padding: var(--spacing-sm);
  }
  
  .theme-option__preview {
    height: 60px;
  }
  
  .theme-selector__actions {
    flex-direction: column;
    align-items: center;
  }
}
</style>