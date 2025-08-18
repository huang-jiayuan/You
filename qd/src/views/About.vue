<template>
  <div class="about-page">
    <!-- 页面头部 -->
    <Header 
      title="关于我们" 
      subtitle="了解更多关于这个Vue3前端应用项目的信息"
    />

    <div class="content-grid">
      <!-- 项目介绍 -->
      <Card title="项目介绍" :shadow="true">
        <template #default>
          <div class="project-intro">
            <p class="intro-text">
              这是一个基于Vue3技术栈构建的现代前端应用模板。项目采用了最新的Vue3 Composition API，
              结合Vite构建工具，为开发者提供了一个高效、现代化的开发起点。
            </p>
            <p class="intro-text">
              该项目展示了Vue3的核心特性，包括响应式系统、组件化架构、路由管理和状态管理等功能，
              是学习和使用Vue3开发的理想参考项目。
            </p>
            <div class="project-stats">
              <div class="stat-item">
                <span class="stat-number">{{ projectStats.components }}</span>
                <span class="stat-label">组件数量</span>
              </div>
              <div class="stat-item">
                <span class="stat-number">{{ projectStats.pages }}</span>
                <span class="stat-label">页面数量</span>
              </div>
              <div class="stat-item">
                <span class="stat-number">{{ projectStats.features }}</span>
                <span class="stat-label">核心特性</span>
              </div>
            </div>
          </div>
        </template>
        <template #actions>
          <Button type="primary" size="small" @click="viewSource">
            查看源码
          </Button>
        </template>
      </Card>

      <!-- 技术栈 -->
      <Card title="技术栈" :shadow="true">
        <template #default>
          <div class="tech-stack">
            <div class="tech-category" v-for="category in techStack" :key="category.name">
              <h3 class="tech-category-title">{{ category.name }}</h3>
              <div class="tech-items">
                <div 
                  class="tech-item" 
                  v-for="tech in category.items" 
                  :key="tech.name"
                >
                  <div class="tech-icon">{{ tech.icon }}</div>
                  <div class="tech-info">
                    <h4 class="tech-name">{{ tech.name }}</h4>
                    <p class="tech-description">{{ tech.description }}</p>
                    <span class="tech-version">{{ tech.version }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </template>
      </Card>

      <!-- 核心特性 -->
      <Card title="核心特性" :shadow="true">
        <template #default>
          <div class="features-grid">
            <div 
              class="feature-item" 
              v-for="feature in coreFeatures" 
              :key="feature.title"
              @click="selectFeature(feature)"
              :class="{ 'feature-item--selected': selectedFeature?.title === feature.title }"
            >
              <div class="feature-icon">{{ feature.icon }}</div>
              <h3 class="feature-title">{{ feature.title }}</h3>
              <p class="feature-description">{{ feature.description }}</p>
              <div class="feature-tags">
                <span 
                  class="feature-tag" 
                  v-for="tag in feature.tags" 
                  :key="tag"
                >
                  {{ tag }}
                </span>
              </div>
            </div>
          </div>
          
          <!-- 选中特性的详细信息 -->
          <div v-if="selectedFeature" class="feature-details">
            <h4>{{ selectedFeature.title }} - 详细信息</h4>
            <p>{{ selectedFeature.details }}</p>
            <div class="feature-examples" v-if="selectedFeature.examples">
              <h5>示例用法：</h5>
              <ul>
                <li v-for="example in selectedFeature.examples" :key="example">
                  {{ example }}
                </li>
              </ul>
            </div>
          </div>
        </template>
        <template #actions>
          <Button 
            type="secondary" 
            size="small" 
            @click="clearSelection"
            v-if="selectedFeature"
          >
            清除选择
          </Button>
        </template>
      </Card>

      <!-- 开发团队 -->
      <Card title="开发信息" :shadow="true">
        <template #default>
          <div class="team-info">
            <div class="team-section">
              <h4>项目状态</h4>
              <div class="status-grid">
                <div class="status-item">
                  <span class="status-label">版本</span>
                  <span class="status-value">v1.0.0</span>
                </div>
                <div class="status-item">
                  <span class="status-label">构建状态</span>
                  <span class="status-value status-success">✅ 正常</span>
                </div>
                <div class="status-item">
                  <span class="status-label">最后更新</span>
                  <span class="status-value">{{ lastUpdated }}</span>
                </div>
              </div>
            </div>
            
            <div class="team-section">
              <h4>开发环境</h4>
              <div class="env-info">
                <p><strong>Node.js:</strong> {{ envInfo.node }}</p>
                <p><strong>包管理器:</strong> {{ envInfo.packageManager }}</p>
                <p><strong>构建工具:</strong> {{ envInfo.buildTool }}</p>
              </div>
            </div>
          </div>
        </template>
        <template #actions>
          <Button type="secondary" size="small" @click="checkUpdates">
            检查更新
          </Button>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import Header from '../components/common/Header.vue'
import Card from '../components/ui/Card.vue'
import Button from '../components/ui/Button.vue'

export default {
  name: 'About',
  components: {
    Header,
    Card,
    Button
  },
  setup() {
    // 响应式数据
    const selectedFeature = ref(null)
    
    // 项目统计
    const projectStats = ref({
      components: 8,
      pages: 3,
      features: 6
    })

    // 技术栈数据
    const techStack = ref([
      {
        name: '前端框架',
        items: [
          {
            name: 'Vue 3',
            icon: '🖖',
            description: '渐进式JavaScript框架',
            version: 'v3.4.0'
          },
          {
            name: 'Composition API',
            icon: '🔧',
            description: '新的组合式API',
            version: 'Built-in'
          },
          {
            name: 'Vue Router',
            icon: '🛣️',
            description: '官方路由管理器',
            version: 'v4.2.5'
          }
        ]
      },
      {
        name: '构建工具',
        items: [
          {
            name: 'Vite',
            icon: '⚡',
            description: '下一代前端构建工具',
            version: 'v5.0.0'
          },
          {
            name: 'ES6+',
            icon: '📜',
            description: '现代JavaScript语法',
            version: 'Latest'
          },
          {
            name: 'CSS3',
            icon: '🎨',
            description: '现代样式技术',
            version: 'Latest'
          }
        ]
      },
      {
        name: '开发工具',
        items: [
          {
            name: 'Vitest',
            icon: '🧪',
            description: '单元测试框架',
            version: 'v1.0.0'
          },
          {
            name: 'Vue Test Utils',
            icon: '🔍',
            description: 'Vue组件测试工具',
            version: 'v2.4.0'
          }
        ]
      }
    ])

    // 核心特性数据
    const coreFeatures = ref([
      {
        title: '现代化架构',
        icon: '🚀',
        description: '采用Vue3 Composition API，提供更好的逻辑复用和类型推导',
        tags: ['Vue3', 'Composition API', '类型安全'],
        details: 'Vue3的Composition API提供了更灵活的代码组织方式，支持更好的TypeScript集成和逻辑复用。',
        examples: [
          '使用setup()函数组织组件逻辑',
          '通过composables实现逻辑复用',
          '更好的TypeScript类型推导支持'
        ]
      },
      {
        title: '快速开发',
        icon: '⚡',
        description: '基于Vite构建，享受极速的热重载和构建体验',
        tags: ['Vite', '热重载', '快速构建'],
        details: 'Vite利用ES模块的特性，提供了极快的开发服务器启动速度和热更新体验。',
        examples: [
          '毫秒级的热更新',
          '按需编译，启动速度快',
          '原生ES模块支持'
        ]
      },
      {
        title: '响应式设计',
        icon: '📱',
        description: '完全响应式布局，适配各种设备和屏幕尺寸',
        tags: ['响应式', '移动端', 'CSS Grid'],
        details: '采用移动优先的设计理念，使用CSS Grid和Flexbox实现完全响应式的布局。',
        examples: [
          'CSS Grid自适应布局',
          '移动端汉堡菜单',
          '灵活的断点设计'
        ]
      },
      {
        title: '主题系统',
        icon: '🎨',
        description: '内置明暗主题切换，支持自定义主题配置',
        tags: ['主题', '暗色模式', 'CSS变量'],
        details: '基于CSS自定义属性实现的主题系统，支持明暗主题切换和主题持久化。',
        examples: [
          'CSS自定义属性定义主题',
          '主题状态持久化',
          '系统主题跟随'
        ]
      },
      {
        title: '组件化架构',
        icon: '🧩',
        description: '可重用的UI组件设计，清晰的职责分离',
        tags: ['组件化', '可重用', '模块化'],
        details: '采用组件化的开发方式，每个组件都有明确的职责和清晰的API设计。',
        examples: [
          '可重用的UI组件库',
          'Props和事件的标准化',
          '组件间的松耦合设计'
        ]
      },
      {
        title: '状态管理',
        icon: '🔄',
        description: '基于Composition API的状态管理解决方案',
        tags: ['状态管理', 'Composables', '响应式'],
        details: '使用Vue3的响应式系统和Composition API实现轻量级的状态管理。',
        examples: [
          'useCounter状态管理',
          'useTheme主题管理',
          '响应式数据共享'
        ]
      }
    ])

    // 环境信息
    const envInfo = ref({
      node: 'v18.0.0+',
      packageManager: 'npm',
      buildTool: 'Vite 5.0'
    })

    // 计算属性
    const lastUpdated = computed(() => {
      return new Date().toLocaleDateString('zh-CN')
    })

    // 方法
    const selectFeature = (feature) => {
      selectedFeature.value = selectedFeature.value?.title === feature.title ? null : feature
    }

    const clearSelection = () => {
      selectedFeature.value = null
    }

    const viewSource = () => {
      window.open('https://github.com/your-repo/vue3-frontend-app', '_blank')
    }

    const checkUpdates = () => {
      alert('当前已是最新版本！')
    }

    return {
      // 响应式数据
      selectedFeature,
      projectStats,
      techStack,
      coreFeatures,
      envInfo,
      
      // 计算属性
      lastUpdated,
      
      // 方法
      selectFeature,
      clearSelection,
      viewSource,
      checkUpdates
    }
  }
}
</script>

<style scoped>
.about-page {
  min-height: 100vh;
}

.content-grid {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-2xl);
  padding: var(--spacing-2xl) 0;
}

/* 项目介绍样式 */
.project-intro {
  text-align: center;
}

.intro-text {
  color: var(--color-text-secondary);
  line-height: 1.7;
  margin-bottom: var(--spacing-lg);
  font-size: var(--font-size-base);
}

.project-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: var(--spacing-lg);
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--color-border);
}

.stat-item {
  text-align: center;
  padding: var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-md);
  transition: all var(--transition-fast);
}

.stat-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}

.stat-number {
  display: block;
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--color-primary);
  margin-bottom: var(--spacing-xs);
}

.stat-label {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  font-weight: 500;
}

/* 技术栈样式 */
.tech-stack {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xl);
}

.tech-category-title {
  font-size: var(--font-size-lg);
  color: var(--color-primary);
  margin-bottom: var(--spacing-md);
  font-weight: 600;
  border-bottom: 2px solid var(--color-primary);
  padding-bottom: var(--spacing-xs);
}

.tech-items {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-md);
}

.tech-item {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-md);
  transition: all var(--transition-fast);
}

.tech-item:hover {
  background-color: var(--color-bg-secondary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}

.tech-icon {
  font-size: var(--font-size-xl);
  flex-shrink: 0;
  margin-top: var(--spacing-xs);
}

.tech-info {
  flex: 1;
}

.tech-name {
  font-size: var(--font-size-base);
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.tech-description {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  margin: 0 0 var(--spacing-xs) 0;
  line-height: 1.5;
}

.tech-version {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
  background-color: var(--color-bg-primary);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--border-radius-sm);
  font-weight: 500;
}

/* 特性网格样式 */
.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
}

.feature-item {
  padding: var(--spacing-lg);
  border-radius: var(--border-radius-md);
  background-color: var(--color-bg-tertiary);
  transition: all var(--transition-normal);
  cursor: pointer;
  border: 2px solid transparent;
}

.feature-item:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
  border-color: var(--color-primary);
}

.feature-item--selected {
  border-color: var(--color-primary);
  background-color: var(--color-bg-secondary);
}

.feature-icon {
  font-size: var(--font-size-3xl);
  margin-bottom: var(--spacing-md);
  text-align: center;
}

.feature-title {
  font-size: var(--font-size-lg);
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-sm);
  font-weight: 600;
  text-align: center;
}

.feature-description {
  color: var(--color-text-secondary);
  font-size: var(--font-size-sm);
  line-height: 1.6;
  margin-bottom: var(--spacing-md);
  text-align: center;
}

.feature-tags {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-xs);
  justify-content: center;
}

.feature-tag {
  background-color: var(--color-primary);
  color: var(--color-white);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--border-radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 500;
}

/* 特性详情样式 */
.feature-details {
  margin-top: var(--spacing-xl);
  padding: var(--spacing-lg);
  background-color: var(--color-bg-secondary);
  border-radius: var(--border-radius-md);
  border-left: 4px solid var(--color-primary);
}

.feature-details h4 {
  color: var(--color-primary);
  margin-bottom: var(--spacing-md);
  font-size: var(--font-size-lg);
}

.feature-details p {
  color: var(--color-text-secondary);
  line-height: 1.6;
  margin-bottom: var(--spacing-md);
}

.feature-examples h5 {
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-sm);
  font-size: var(--font-size-base);
}

.feature-examples ul {
  margin: 0;
  padding-left: var(--spacing-lg);
}

.feature-examples li {
  color: var(--color-text-secondary);
  margin-bottom: var(--spacing-xs);
  line-height: 1.5;
}

/* 团队信息样式 */
.team-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-xl);
}

.team-section h4 {
  color: var(--color-primary);
  margin-bottom: var(--spacing-md);
  font-size: var(--font-size-lg);
  font-weight: 600;
}

.status-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--spacing-sm);
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-sm);
}

.status-label {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  font-weight: 500;
}

.status-value {
  font-size: var(--font-size-sm);
  color: var(--color-text-primary);
  font-weight: 600;
}

.status-success {
  color: var(--color-success) !important;
}

.env-info p {
  margin: var(--spacing-sm) 0;
  color: var(--color-text-secondary);
  font-size: var(--font-size-sm);
  line-height: 1.5;
}

.env-info strong {
  color: var(--color-text-primary);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .content-grid {
    padding: var(--spacing-xl) 0;
    gap: var(--spacing-xl);
  }
  
  .project-stats {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .tech-items {
    grid-template-columns: 1fr;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .team-info {
    grid-template-columns: 1fr;
    gap: var(--spacing-lg);
  }
  
  .feature-item {
    padding: var(--spacing-md);
  }
  
  .feature-details {
    padding: var(--spacing-md);
  }
}

@media (max-width: 480px) {
  .stat-item {
    padding: var(--spacing-sm);
  }
  
  .stat-number {
    font-size: var(--font-size-xl);
  }
  
  .tech-item {
    flex-direction: column;
    text-align: center;
    gap: var(--spacing-sm);
  }
  
  .feature-tags {
    justify-content: flex-start;
  }
}
</style>