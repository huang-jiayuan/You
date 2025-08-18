<template>
  <div class="not-found-page">
    <!-- 页面头部 -->
    <Header 
      title="页面未找到" 
      subtitle="抱歉，您访问的页面不存在"
    />

    <div class="not-found-content">
      <Card title="404 错误" :shadow="true" class="not-found-card">
        <template #default>
          <div class="error-display">
            <div class="error-icon">🔍</div>
            <div class="error-code">404</div>
            <div class="error-message">
              <h3>页面未找到</h3>
              <p>您访问的页面可能已被删除、重命名或暂时不可用。</p>
            </div>
          </div>
          
          <div class="suggestions">
            <h4>您可以尝试：</h4>
            <ul class="suggestion-list">
              <li>检查URL地址是否正确</li>
              <li>返回首页重新导航</li>
              <li>使用搜索功能查找内容</li>
              <li>联系我们获取帮助</li>
            </ul>
          </div>
          
          <div class="quick-links">
            <h4>快速链接：</h4>
            <div class="link-grid">
              <router-link 
                v-for="link in quickLinks" 
                :key="link.path"
                :to="link.path" 
                class="quick-link"
              >
                <span class="link-icon">{{ link.icon }}</span>
                <span class="link-text">{{ link.name }}</span>
              </router-link>
            </div>
          </div>
        </template>
        
        <template #actions>
          <div class="action-buttons">
            <Button 
              type="primary" 
              size="medium"
              @click="goHome"
            >
              返回首页
            </Button>
            <Button 
              type="secondary" 
              size="medium"
              @click="goBack"
            >
              返回上页
            </Button>
            <Button 
              type="secondary" 
              size="medium"
              @click="reportIssue"
            >
              报告问题
            </Button>
          </div>
        </template>
      </Card>
      
      <!-- 最近访问的页面 -->
      <Card title="最近访问" :shadow="true" v-if="recentPages.length > 0">
        <template #default>
          <div class="recent-pages">
            <router-link 
              v-for="page in recentPages" 
              :key="page.path"
              :to="page.path" 
              class="recent-page"
            >
              <span class="page-icon">{{ page.icon }}</span>
              <div class="page-info">
                <h5 class="page-title">{{ page.title }}</h5>
                <p class="page-description">{{ page.description }}</p>
              </div>
            </router-link>
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useNotifications } from '../composables/useNotifications.js'
import Header from '../components/common/Header.vue'
import Card from '../components/ui/Card.vue'
import Button from '../components/ui/Button.vue'

export default {
  name: 'NotFound',
  components: {
    Header,
    Card,
    Button
  },
  setup() {
    const router = useRouter()
    const { notify } = useNotifications()
    
    // 快速链接
    const quickLinks = ref([
      {
        path: '/',
        name: '首页',
        icon: '🏠'
      },
      {
        path: '/about',
        name: '关于我们',
        icon: 'ℹ️'
      },
      {
        path: '/contact',
        name: '联系我们',
        icon: '📞'
      }
    ])
    
    // 最近访问的页面（从localStorage获取）
    const recentPages = ref([])
    
    // 方法
    const goHome = () => {
      router.push('/')
      notify.info('已返回首页')
    }
    
    const goBack = () => {
      if (window.history.length > 1) {
        router.go(-1)
      } else {
        router.push('/')
      }
      notify.info('已返回上一页')
    }
    
    const reportIssue = () => {
      const currentUrl = window.location.href
      const subject = encodeURIComponent('页面访问问题报告')
      const body = encodeURIComponent(`我在访问以下页面时遇到了404错误：\n\n页面URL: ${currentUrl}\n时间: ${new Date().toLocaleString()}\n\n请帮助解决此问题。`)
      
      // 可以打开邮件客户端或者跳转到联系页面
      router.push({
        path: '/contact',
        query: {
          subject: '页面访问问题',
          message: `页面URL: ${currentUrl}`
        }
      })
      
      notify.info('已跳转到联系页面，您可以在那里报告问题')
    }
    
    // 获取最近访问的页面
    const loadRecentPages = () => {
      try {
        const recent = localStorage.getItem('recent-pages')
        if (recent) {
          const pages = JSON.parse(recent)
          recentPages.value = pages.slice(0, 3) // 只显示最近3个
        }
      } catch (error) {
        console.warn('无法加载最近访问的页面:', error)
      }
    }
    
    // 生命周期
    onMounted(() => {
      loadRecentPages()
      
      // 记录404错误
      console.warn('404页面访问:', window.location.href)
      
      // 可以发送404统计到分析服务
      if (window.gtag) {
        window.gtag('event', 'page_not_found', {
          page_location: window.location.href,
          page_title: document.title
        })
      }
    })
    
    return {
      quickLinks,
      recentPages,
      goHome,
      goBack,
      reportIssue
    }
  }
}
</script>

<style scoped>
.not-found-page {
  min-height: 100vh;
}

.not-found-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-2xl);
  padding: var(--spacing-2xl) 0;
  max-width: 800px;
  margin: 0 auto;
}

.not-found-card {
  text-align: center;
}

.error-display {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.error-icon {
  font-size: 4rem;
  opacity: 0.6;
}

.error-code {
  font-size: 6rem;
  font-weight: 900;
  color: var(--color-primary);
  line-height: 1;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
}

.error-message h3 {
  font-size: var(--font-size-2xl);
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-md);
  font-weight: 600;
}

.error-message p {
  font-size: var(--font-size-base);
  color: var(--color-text-secondary);
  line-height: 1.6;
  max-width: 500px;
  margin: 0 auto;
}

.suggestions {
  text-align: left;
  margin-bottom: var(--spacing-xl);
  padding: var(--spacing-lg);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-md);
}

.suggestions h4 {
  font-size: var(--font-size-lg);
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-md);
  font-weight: 600;
}

.suggestion-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.suggestion-list li {
  padding: var(--spacing-sm) 0;
  color: var(--color-text-secondary);
  position: relative;
  padding-left: var(--spacing-lg);
}

.suggestion-list li::before {
  content: '•';
  color: var(--color-primary);
  font-weight: bold;
  position: absolute;
  left: 0;
}

.quick-links {
  text-align: left;
  margin-bottom: var(--spacing-xl);
}

.quick-links h4 {
  font-size: var(--font-size-lg);
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-md);
  font-weight: 600;
}

.link-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: var(--spacing-md);
}

.quick-link {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-md);
  text-decoration: none;
  color: var(--color-text-primary);
  transition: all var(--transition-fast);
}

.quick-link:hover {
  background-color: var(--color-primary-bg);
  color: var(--color-primary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}

.link-icon {
  font-size: var(--font-size-lg);
}

.link-text {
  font-weight: 500;
}

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  justify-content: center;
  flex-wrap: wrap;
}

.recent-pages {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.recent-page {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-md);
  text-decoration: none;
  color: var(--color-text-primary);
  transition: all var(--transition-fast);
}

.recent-page:hover {
  background-color: var(--color-primary-bg);
  transform: translateX(4px);
}

.page-icon {
  font-size: var(--font-size-xl);
  flex-shrink: 0;
}

.page-info {
  flex: 1;
}

.page-title {
  font-size: var(--font-size-base);
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.page-description {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  margin: 0;
  line-height: 1.4;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .not-found-content {
    padding: var(--spacing-xl) 0;
    gap: var(--spacing-xl);
  }
  
  .error-code {
    font-size: 4rem;
  }
  
  .error-icon {
    font-size: 3rem;
  }
  
  .link-grid {
    grid-template-columns: 1fr;
  }
  
  .action-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .suggestions,
  .quick-links {
    text-align: center;
  }
}

@media (max-width: 480px) {
  .error-code {
    font-size: 3rem;
  }
  
  .error-message h3 {
    font-size: var(--font-size-xl);
  }
  
  .suggestions {
    padding: var(--spacing-md);
  }
}
</style>