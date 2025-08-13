<template>
  <div class="contact-page">
    <!-- 页面头部 -->
    <Header 
      title="联系我们" 
      subtitle="有任何问题或建议？欢迎与我们取得联系"
    />

    <div class="contact-content">
      <!-- 联系表单 -->
      <Card title="发送消息" :shadow="true" class="contact-form-card">
        <template #default>
          <form @submit.prevent="handleSubmit" class="contact-form" v-if="!showSuccess">
            <div class="form-row">
              <div class="form-group">
                <label for="name" class="form-label">姓名 *</label>
                <input
                  id="name"
                  v-model="form.name"
                  type="text"
                  class="form-input"
                  :class="{ 'form-input--error': errors.name }"
                  placeholder="请输入您的姓名"
                />
                <span v-if="errors.name" class="error-message">{{ errors.name }}</span>
              </div>

              <div class="form-group">
                <label for="email" class="form-label">邮箱 *</label>
                <input
                  id="email"
                  v-model="form.email"
                  type="email"
                  class="form-input"
                  :class="{ 'form-input--error': errors.email }"
                  placeholder="请输入您的邮箱地址"
                />
                <span v-if="errors.email" class="error-message">{{ errors.email }}</span>
              </div>
            </div>

            <div class="form-group">
              <label for="subject" class="form-label">主题</label>
              <input
                id="subject"
                v-model="form.subject"
                type="text"
                class="form-input"
                placeholder="请输入消息主题"
              />
            </div>

            <div class="form-group">
              <label for="message" class="form-label">消息内容 *</label>
              <textarea
                id="message"
                v-model="form.message"
                class="form-textarea"
                :class="{ 'form-input--error': errors.message }"
                rows="5"
                placeholder="请输入您的消息内容"
              ></textarea>
              <span v-if="errors.message" class="error-message">{{ errors.message }}</span>
              <div class="character-count">
                {{ form.message.length }}/500 字符
              </div>
            </div>

            <div class="form-actions">
              <Button 
                type="primary" 
                size="medium"
                :disabled="isSubmitting"
                @click="handleSubmit"
              >
                {{ isSubmitting ? '发送中...' : '发送消息' }}
              </Button>
              <Button 
                type="secondary" 
                size="medium"
                @click="resetForm"
                :disabled="isSubmitting"
              >
                重置表单
              </Button>
            </div>
          </form>

          <!-- 成功提示 -->
          <div v-if="showSuccess" class="success-message">
            <div class="success-icon">✅</div>
            <h3>消息发送成功！</h3>
            <p>感谢您的反馈，我们会尽快回复您。</p>
            <Button 
              type="primary" 
              size="small"
              @click="resetSuccess"
              class="success-button"
            >
              发送新消息
            </Button>
          </div>
        </template>
      </Card>

      <!-- 联系信息 -->
      <div class="contact-info-grid">
        <Card title="联系方式" :shadow="true">
          <template #default>
            <div class="contact-methods">
              <div class="contact-method" v-for="method in contactMethods" :key="method.type">
                <div class="method-icon">{{ method.icon }}</div>
                <div class="method-content">
                  <h4 class="method-title">{{ method.title }}</h4>
                  <p class="method-value">{{ method.value }}</p>
                  <span class="method-description">{{ method.description }}</span>
                </div>
              </div>
            </div>
          </template>
        </Card>

        <Card title="工作时间" :shadow="true">
          <template #default>
            <div class="working-hours">
              <div class="hours-item" v-for="schedule in workingHours" :key="schedule.day">
                <span class="hours-day">{{ schedule.day }}</span>
                <span class="hours-time" :class="{ 'hours-time--closed': schedule.closed }">
                  {{ schedule.closed ? '休息' : schedule.time }}
                </span>
              </div>
            </div>
            <div class="timezone-info">
              <p>⏰ 时区：北京时间 (UTC+8)</p>
              <p>📞 紧急联系：support@vue3app.com</p>
            </div>
          </template>
        </Card>

        <Card title="常见问题" :shadow="true">
          <template #default>
            <div class="faq-section">
              <div 
                class="faq-item" 
                v-for="(faq, index) in faqs" 
                :key="index"
                @click="toggleFaq(index)"
                :class="{ 'faq-item--open': openFaqIndex === index }"
              >
                <div class="faq-question">
                  <span>{{ faq.question }}</span>
                  <span class="faq-toggle">{{ openFaqIndex === index ? '−' : '+' }}</span>
                </div>
                <div class="faq-answer" v-if="openFaqIndex === index">
                  {{ faq.answer }}
                </div>
              </div>
            </div>
          </template>
          <template #actions>
            <Button type="secondary" size="small" @click="viewAllFaqs">
              查看更多FAQ
            </Button>
          </template>
        </Card>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed } from 'vue'
import Header from '../components/common/Header.vue'
import Card from '../components/ui/Card.vue'
import Button from '../components/ui/Button.vue'

export default {
  name: 'Contact',
  components: {
    Header,
    Card,
    Button
  },
  setup() {
    // 表单数据
    const form = reactive({
      name: '',
      email: '',
      subject: '',
      message: ''
    })

    // 错误信息
    const errors = reactive({
      name: '',
      email: '',
      message: ''
    })

    // 状态管理
    const isSubmitting = ref(false)
    const showSuccess = ref(false)
    const openFaqIndex = ref(null)

    // 联系方式数据
    const contactMethods = ref([
      {
        type: 'email',
        icon: '📧',
        title: '邮箱地址',
        value: 'contact@vue3app.com',
        description: '我们会在24小时内回复'
      },
      {
        type: 'website',
        icon: '🌐',
        title: '官方网站',
        value: 'www.vue3app.com',
        description: '查看更多产品信息'
      },
      {
        type: 'github',
        icon: '💻',
        title: 'GitHub',
        value: 'github.com/vue3app',
        description: '查看开源代码'
      },
      {
        type: 'support',
        icon: '💬',
        title: '在线支持',
        value: 'support@vue3app.com',
        description: '技术支持和帮助'
      }
    ])

    // 工作时间数据
    const workingHours = ref([
      { day: '周一', time: '09:00 - 18:00', closed: false },
      { day: '周二', time: '09:00 - 18:00', closed: false },
      { day: '周三', time: '09:00 - 18:00', closed: false },
      { day: '周四', time: '09:00 - 18:00', closed: false },
      { day: '周五', time: '09:00 - 18:00', closed: false },
      { day: '周六', time: '', closed: true },
      { day: '周日', time: '', closed: true }
    ])

    // 常见问题数据
    const faqs = ref([
      {
        question: '如何开始使用这个Vue3应用？',
        answer: '您可以克隆我们的GitHub仓库，然后运行npm install安装依赖，最后执行npm run dev启动开发服务器。'
      },
      {
        question: '这个项目支持哪些浏览器？',
        answer: '我们支持所有现代浏览器，包括Chrome、Firefox、Safari和Edge的最新版本。'
      },
      {
        question: '可以用于商业项目吗？',
        answer: '是的，这个项目采用MIT许可证，您可以自由地用于个人和商业项目。'
      },
      {
        question: '如何贡献代码？',
        answer: '欢迎提交Pull Request！请先阅读我们的贡献指南，确保代码符合项目规范。'
      }
    ])

    // 表单验证
    const validateForm = () => {
      // 清空之前的错误
      errors.name = ''
      errors.email = ''
      errors.message = ''

      let isValid = true

      // 验证姓名
      if (!form.name.trim()) {
        errors.name = '请输入您的姓名'
        isValid = false
      } else if (form.name.trim().length < 2) {
        errors.name = '姓名至少需要2个字符'
        isValid = false
      }

      // 验证邮箱
      if (!form.email.trim()) {
        errors.email = '请输入邮箱地址'
        isValid = false
      } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
        errors.email = '请输入有效的邮箱地址'
        isValid = false
      }

      // 验证消息内容
      if (!form.message.trim()) {
        errors.message = '请输入消息内容'
        isValid = false
      } else if (form.message.trim().length < 10) {
        errors.message = '消息内容至少需要10个字符'
        isValid = false
      } else if (form.message.length > 500) {
        errors.message = '消息内容不能超过500个字符'
        isValid = false
      }

      return isValid
    }

    // 处理表单提交
    const handleSubmit = async () => {
      if (!validateForm()) {
        return
      }

      isSubmitting.value = true

      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 2000))
        
        // 显示成功消息
        showSuccess.value = true
        
        // 重置表单
        resetForm()
        
      } catch (error) {
        console.error('表单提交失败:', error)
        alert('发送失败，请稍后重试')
      } finally {
        isSubmitting.value = false
      }
    }

    // 重置表单
    const resetForm = () => {
      form.name = ''
      form.email = ''
      form.subject = ''
      form.message = ''
      
      // 清空错误信息
      errors.name = ''
      errors.email = ''
      errors.message = ''
    }

    // 重置成功状态
    const resetSuccess = () => {
      showSuccess.value = false
    }

    // 切换FAQ显示
    const toggleFaq = (index) => {
      openFaqIndex.value = openFaqIndex.value === index ? null : index
    }

    // 查看所有FAQ
    const viewAllFaqs = () => {
      alert('更多FAQ请访问我们的帮助中心！')
    }

    return {
      // 表单相关
      form,
      errors,
      isSubmitting,
      showSuccess,
      
      // 数据
      contactMethods,
      workingHours,
      faqs,
      openFaqIndex,
      
      // 方法
      handleSubmit,
      resetForm,
      resetSuccess,
      toggleFaq,
      viewAllFaqs
    }
  }
}
</script>

<style scoped>
.contact-page {
  min-height: 100vh;
}

.contact-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: var(--spacing-2xl);
  align-items: start;
  padding: var(--spacing-2xl) 0;
}

/* 表单样式 */
.contact-form-card {
  height: fit-content;
}

.contact-form {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-lg);
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  font-weight: 500;
  color: var(--color-text-primary);
  margin-bottom: var(--spacing-sm);
  font-size: var(--font-size-sm);
}

.form-input,
.form-textarea {
  padding: var(--spacing-md);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-md);
  font-size: var(--font-size-base);
  background-color: var(--color-bg-secondary);
  color: var(--color-text-primary);
  transition: all var(--transition-fast);
  font-family: inherit;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-input--error {
  border-color: var(--color-danger);
}

.form-textarea {
  resize: vertical;
  min-height: 120px;
}

.error-message {
  color: var(--color-danger);
  font-size: var(--font-size-sm);
  margin-top: var(--spacing-xs);
}

.character-count {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
  text-align: right;
  margin-top: var(--spacing-xs);
}

.form-actions {
  display: flex;
  gap: var(--spacing-md);
  margin-top: var(--spacing-md);
}

/* 成功消息样式 */
.success-message {
  text-align: center;
  padding: var(--spacing-2xl);
}

.success-icon {
  font-size: 4rem;
  margin-bottom: var(--spacing-lg);
}

.success-message h3 {
  color: var(--color-success);
  margin-bottom: var(--spacing-md);
  font-weight: 600;
  font-size: var(--font-size-xl);
}

.success-message p {
  color: var(--color-text-secondary);
  margin-bottom: var(--spacing-lg);
  line-height: 1.6;
}

.success-button {
  margin-top: var(--spacing-md);
}

/* 联系信息网格 */
.contact-info-grid {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xl);
}

/* 联系方式样式 */
.contact-methods {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.contact-method {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border-radius: var(--border-radius-md);
  background-color: var(--color-bg-tertiary);
  transition: all var(--transition-fast);
}

.contact-method:hover {
  background-color: var(--color-bg-secondary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-sm);
}

.method-icon {
  font-size: var(--font-size-xl);
  flex-shrink: 0;
  margin-top: var(--spacing-xs);
}

.method-content {
  flex: 1;
}

.method-title {
  font-size: var(--font-size-base);
  font-weight: 600;
  color: var(--color-text-primary);
  margin: 0 0 var(--spacing-xs) 0;
}

.method-value {
  font-size: var(--font-size-sm);
  color: var(--color-primary);
  font-weight: 500;
  margin: 0 0 var(--spacing-xs) 0;
}

.method-description {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
}

/* 工作时间样式 */
.working-hours {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-lg);
}

.hours-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  border-radius: var(--border-radius-sm);
}

.hours-day {
  font-size: var(--font-size-sm);
  font-weight: 500;
  color: var(--color-text-primary);
}

.hours-time {
  font-size: var(--font-size-sm);
  color: var(--color-success);
  font-weight: 500;
}

.hours-time--closed {
  color: var(--color-text-muted);
}

.timezone-info {
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--color-border);
}

.timezone-info p {
  font-size: var(--font-size-xs);
  color: var(--color-text-secondary);
  margin: var(--spacing-xs) 0;
  line-height: 1.5;
}

/* FAQ样式 */
.faq-section {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.faq-item {
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-md);
  overflow: hidden;
  transition: all var(--transition-fast);
  cursor: pointer;
}

.faq-item:hover {
  border-color: var(--color-primary);
}

.faq-item--open {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-sm);
}

.faq-question {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  background-color: var(--color-bg-tertiary);
  font-size: var(--font-size-sm);
  font-weight: 500;
  color: var(--color-text-primary);
}

.faq-toggle {
  font-size: var(--font-size-lg);
  font-weight: 700;
  color: var(--color-primary);
  transition: transform var(--transition-fast);
}

.faq-item--open .faq-toggle {
  transform: rotate(180deg);
}

.faq-answer {
  padding: var(--spacing-md);
  background-color: var(--color-bg-primary);
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  line-height: 1.6;
  border-top: 1px solid var(--color-border);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .contact-content {
    grid-template-columns: 1fr;
    gap: var(--spacing-xl);
    padding: var(--spacing-xl) 0;
  }
  
  .form-row {
    grid-template-columns: 1fr;
    gap: var(--spacing-md);
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .contact-method {
    padding: var(--spacing-sm);
  }
  
  .method-icon {
    font-size: var(--font-size-lg);
  }
  
  .hours-item {
    padding: var(--spacing-xs) var(--spacing-sm);
  }
  
  .faq-question {
    padding: var(--spacing-sm);
  }
  
  .faq-answer {
    padding: var(--spacing-sm);
  }
}

@media (max-width: 480px) {
  .success-message {
    padding: var(--spacing-lg);
  }
  
  .success-icon {
    font-size: 3rem;
  }
  
  .contact-method {
    flex-direction: column;
    text-align: center;
    gap: var(--spacing-sm);
  }
  
  .method-icon {
    align-self: center;
  }
}
</style>