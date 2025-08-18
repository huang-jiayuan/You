<template>
  <footer class="footer">
    <div class="container">
      <div class="footer__content">
        <div class="footer__info">
          <p class="footer__copyright">
            &copy; {{ currentYear }} Vue3 前端应用. 使用Vue3 + Vite构建
          </p>
          <p class="footer__description">
            现代化的Vue3前端应用模板，展示最佳实践和组件化开发
          </p>
        </div>
        
        <div class="footer__links">
          <div class="footer__section">
            <h4 class="footer__section-title">快速链接</h4>
            <ul class="footer__link-list">
              <li><router-link to="/" class="footer__link">首页</router-link></li>
              <li><router-link to="/about" class="footer__link">关于</router-link></li>
              <li><router-link to="/contact" class="footer__link">联系</router-link></li>
            </ul>
          </div>
          
          <div class="footer__section">
            <h4 class="footer__section-title">技术栈</h4>
            <ul class="footer__link-list">
              <li><span class="footer__tech">Vue 3</span></li>
              <li><span class="footer__tech">Vite</span></li>
              <li><span class="footer__tech">Vue Router</span></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import { computed } from 'vue'

export default {
  name: 'Footer',
  setup() {
    const currentYear = computed(() => new Date().getFullYear())
    
    return {
      currentYear
    }
  }
}
</script>

<style scoped>
.footer {
  background-color: var(--color-bg-secondary);
  border-top: 1px solid var(--color-border);
  padding: var(--spacing-2xl) 0 var(--spacing-lg) 0;
  margin-top: auto;
}

.footer__content {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: var(--spacing-2xl);
  align-items: start;
}

.footer__info {
  max-width: 500px;
}

.footer__copyright {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  margin: 0 0 var(--spacing-sm) 0;
  font-weight: 500;
}

.footer__description {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
  margin: 0;
  line-height: 1.5;
}

.footer__links {
  display: flex;
  gap: var(--spacing-2xl);
}

.footer__section-title {
  font-size: var(--font-size-sm);
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 var(--spacing-md) 0;
}

.footer__link-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer__link-list li {
  margin-bottom: var(--spacing-xs);
}

.footer__link {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.footer__link:hover {
  color: var(--color-primary);
}

.footer__tech {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .footer__content {
    grid-template-columns: 1fr;
    gap: var(--spacing-xl);
    text-align: center;
  }
  
  .footer__links {
    justify-content: center;
    gap: var(--spacing-xl);
  }
  
  .footer__info {
    max-width: none;
  }
}

@media (max-width: 480px) {
  .footer__links {
    flex-direction: column;
    gap: var(--spacing-lg);
  }
}
</style>