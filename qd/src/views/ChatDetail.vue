<template>
  <div class="chat-detail-page">
    <!-- 顶部导航栏 -->
    <div class="chat-header">
      <div class="header-left">
        <button class="back-btn" @click="goBack">
          <span class="back-icon">←</span>
        </button>
        <div class="user-info">
          <h2 class="user-name">{{ chatUser.name }}</h2>
          <div class="user-status" :class="onlineStatusClass">
            {{ onlineStatusText }}
          </div>
        </div>
      </div>
      <div class="header-right">
        <button class="action-btn">关注</button>
        <button class="menu-btn">⋮</button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="isLoading" class="loading-container">
      <div class="loading-spinner"></div>
      <p>正在加载用户信息...</p>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="loadError" class="error-container">
      <div class="error-icon">❌</div>
      <p>{{ loadError }}</p>
      <button class="retry-btn" @click="loadChatUser">重试</button>
    </div>

    <!-- 用户资料卡片 -->
    <div v-else class="user-profile-card">
      <div class="user-avatar">
        <img :src="chatUser.avatar" :alt="chatUser.name" />
        <div v-if="chatUser.isOnline" class="online-indicator"></div>
      </div>
      <div class="user-details">
        <div class="user-tags">
          <span class="tag vip-tag">VV</span>
          <span class="tag heart-tag">💖</span>
          <span class="tag heart-tag">🤍</span>
          <span class="tag heart-tag">🤍</span>
          <span class="tag heart-tag">🤍</span>
        </div>
        <div class="user-description">
          <p>{{ chatUser.description || '这个人很懒，什么都没有留下' }}</p>
        </div>
        <div class="user-interests">
          <p>他的标签：<span class="interest-tags">{{ chatUser.interests || '暂无标签' }}</span></p>
        </div>
      </div>
      <div class="expand-btn">
        <span>⌄</span>
      </div>
    </div>

    <!-- 聊天消息区域 -->
    <div class="chat-messages" ref="messagesContainer">
      <div class="date-divider">
        <span>{{ formatDate(new Date()) }}</span>
      </div>
      
      <div 
        v-for="message in messages" 
        :key="message.id"
        class="message-item"
        :class="{ 'own-message': message.isOwn }"
      >
        <div v-if="!message.isOwn" class="message-avatar">
          <img :src="chatUser.avatar" :alt="chatUser.name" />
        </div>
        <div class="message-content">
          <div class="message-bubble">
            <p>{{ message.content }}</p>
          </div>
          <div class="message-time">
            {{ formatTime(message.timestamp) }}
          </div>
        </div>
      </div>
    </div>

    <!-- 底部输入区域 -->
    <div class="chat-input-area">
      <div class="input-container">
        <button class="voice-btn">🎤</button>
        <div class="text-input-wrapper">
          <input 
            v-model="newMessage"
            type="text" 
            placeholder="输入想说的内容"
            class="message-input"
            @keyup.enter="sendMessage"
          />
        </div>
        <button class="emoji-btn">😊</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useToast } from '@/composables/useToast'
import { useNavigationError } from '@/composables/useNavigationError'

export default {
  name: 'ChatDetail',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const { error: showError, warning: showWarning } = useToast()
    const { 
      handleParameterError, 
      handleNetworkError,
      showWarningMessage 
    } = useNavigationError()
    
    const newMessage = ref('')
    const messagesContainer = ref(null)
    const isLoading = ref(true)
    const loadError = ref(null)
    
    // 聊天用户信息
    const chatUser = ref({
      id: null,
      name: '用户',
      avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%234facfe"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E用%3C/text%3E%3C/svg%3E',
      description: '他刚来ME不久，还是一个小萌新',
      interests: '视频聊、健身达人',
      isOnline: false
    })

    // 计算属性：在线状态显示
    const onlineStatusText = computed(() => {
      return chatUser.value.isOnline ? '[在线]' : '[离线]'
    })

    // 计算属性：在线状态样式类
    const onlineStatusClass = computed(() => {
      return chatUser.value.isOnline ? 'online' : 'offline'
    })
    
    // 聊天消息
    const messages = ref([
      {
        id: 1,
        content: '晚上好在干嘛呢',
        isOwn: false,
        timestamp: new Date()
      }
    ])

    // 生成用户头像函数
    const generateUserAvatar = (char, userId) => {
      // 根据用户ID生成颜色
      const colors = [
        '#4facfe', '#667eea', '#fa709a', '#764ba2', 
        '#ff6b9d', '#ff9a56', '#4CAF50', '#2196F3'
      ]
      const colorIndex = parseInt(userId) % colors.length
      const color = colors[colorIndex]
      
      return `data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="${encodeURIComponent(color)}"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E${encodeURIComponent(char)}%3C/text%3E%3C/svg%3E`
    }

    // 参数验证函数
    const validateParams = () => {
      const userId = route.params.userId
      
      if (!userId) {
        loadError.value = '缺少用户ID参数'
        handleParameterError('userId', userId)
        return false
      }

      // 验证用户ID格式（应该是数字或字符串）
      if (typeof userId !== 'string' && typeof userId !== 'number') {
        loadError.value = '用户ID格式无效'
        handleParameterError('userId格式', userId)
        return false
      }

      return true
    }

    // 从路由参数获取用户信息
    const loadChatUser = async () => {
      try {
        isLoading.value = true
        loadError.value = null

        // 参数验证
        if (!validateParams()) {
          isLoading.value = false
          return
        }

        const userId = route.params.userId
        const userName = route.query.name || '用户'
        const isOnline = route.query.isOnline === 'true'

        // 生成用户头像（基于用户名的第一个字符）
        const firstChar = userName.charAt(0) || '用'
        const userAvatar = generateUserAvatar(firstChar, userId)

        // 参数完整性检查
        if (!userName || userName === '用户') {
          showWarningMessage('用户名信息不完整，使用默认显示')
        }

        // 更新用户信息
        chatUser.value = {
          id: userId,
          name: userName,
          avatar: userAvatar,
          description: '他刚来ME不久，还是一个小萌新',
          interests: '视频聊、健身达人',
          isOnline: isOnline
        }

        console.log('用户信息加载成功:', chatUser.value)

      } catch (error) {
        console.error('加载用户信息失败:', error)
        loadError.value = error.message || '加载用户信息失败'
        
        // 根据错误类型使用不同的处理方式
        if (error.message?.includes('network') || error.message?.includes('fetch')) {
          handleNetworkError(error, '加载用户信息')
        } else {
          showError('加载用户信息失败，请稍后重试')
        }
      } finally {
        isLoading.value = false
      }
    }

    // 发送消息
    const sendMessage = () => {
      if (!newMessage.value.trim()) return
      
      const message = {
        id: Date.now(),
        content: newMessage.value,
        isOwn: true,
        timestamp: new Date()
      }
      
      messages.value.push(message)
      newMessage.value = ''
    }

    // 返回上一页
    const goBack = () => {
      router.go(-1)
    }

    // 格式化日期
    const formatDate = (date) => {
      return '今天'
    }

    // 格式化时间
    const formatTime = (timestamp) => {
      const date = new Date(timestamp)
      const hours = date.getHours().toString().padStart(2, '0')
      const minutes = date.getMinutes().toString().padStart(2, '0')
      return `${hours}:${minutes}`
    }

    onMounted(() => {
      loadChatUser()
    })

    return {
      chatUser,
      messages,
      newMessage,
      messagesContainer,
      isLoading,
      loadError,
      onlineStatusText,
      onlineStatusClass,
      sendMessage,
      goBack,
      loadChatUser,
      formatDate,
      formatTime
    }
  }
}
</script>

<style scoped>
.chat-detail-page {
  min-height: 100vh;
  background: #f5f5f5;
  display: flex;
  flex-direction: column;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  animation: slideInFromRight 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

@keyframes slideInFromRight {
  0% {
    opacity: 0;
    transform: translateX(100%);
  }
  100% {
    opacity: 1;
    transform: translateX(0);
  }
}

/* 顶部导航栏 */
.chat-header {
  background: white;
  padding: 12px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #eee;
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.back-btn {
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
  padding: 4px;
}

.user-name {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.user-status {
  font-size: 12px;
  margin-top: 2px;
  font-weight: 500;
}

.user-status.online {
  color: #4CAF50;
}

.user-status.offline {
  color: #999;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.action-btn {
  background: #667eea;
  color: white;
  border: none;
  border-radius: 16px;
  padding: 6px 16px;
  font-size: 14px;
  cursor: pointer;
}

.menu-btn {
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
  padding: 4px;
}

/* 用户资料卡片 */
.user-profile-card {
  background: white;
  padding: 16px;
  margin: 8px 16px;
  border-radius: 12px;
  display: flex;
  gap: 12px;
  align-items: flex-start;
  animation: fadeInUp 0.4s ease-out 0.1s both;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.3s ease;
}

.user-profile-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.user-avatar {
  width: 60px;
  height: 60px;
  flex-shrink: 0;
  position: relative;
}

.user-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.online-indicator {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 16px;
  height: 16px;
  background: #4CAF50;
  border: 2px solid white;
  border-radius: 50%;
  animation: onlinePulse 2s ease-in-out infinite;
}

@keyframes onlinePulse {
  0%, 100% {
    transform: scale(1);
    box-shadow: 0 0 0 0 rgba(76, 175, 80, 0.7);
  }
  50% {
    transform: scale(1.1);
    box-shadow: 0 0 0 4px rgba(76, 175, 80, 0);
  }
}

.user-details {
  flex: 1;
}

.user-tags {
  display: flex;
  gap: 4px;
  margin-bottom: 8px;
}

.tag {
  padding: 2px 6px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: bold;
}

.vip-tag {
  background: #ff6b9d;
  color: white;
}

.heart-tag {
  font-size: 14px;
}

.user-description {
  margin-bottom: 8px;
}

.user-description p {
  font-size: 14px;
  color: #666;
  line-height: 1.4;
  margin: 0;
}

.user-interests {
  font-size: 14px;
  color: #666;
}

.interest-tags {
  color: #333;
}

.expand-btn {
  color: #999;
  cursor: pointer;
  font-size: 16px;
}

/* 聊天消息区域 */
.chat-messages {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.date-divider {
  text-align: center;
  margin: 16px 0;
}

.date-divider span {
  background: rgba(0, 0, 0, 0.1);
  color: #666;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
}

.message-item {
  display: flex;
  gap: 8px;
  align-items: flex-start;
  animation: messageSlideIn 0.3s ease-out;
}

.message-item.own-message {
  flex-direction: row-reverse;
}

@keyframes messageSlideIn {
  0% {
    opacity: 0;
    transform: translateY(10px) scale(0.95);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.message-avatar {
  width: 32px;
  height: 32px;
  flex-shrink: 0;
}

.message-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.message-content {
  max-width: 70%;
}

.message-bubble {
  background: white;
  padding: 12px 16px;
  border-radius: 18px;
  margin-bottom: 4px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
  cursor: pointer;
}

.message-bubble:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  transform: translateY(-1px);
}

.own-message .message-bubble {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.own-message .message-bubble:hover {
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.message-bubble p {
  margin: 0;
  font-size: 16px;
  line-height: 1.4;
}

.message-time {
  font-size: 12px;
  color: #999;
  text-align: center;
}

.own-message .message-time {
  text-align: right;
}

/* 底部输入区域 */
.chat-input-area {
  background: white;
  padding: 12px 16px;
  border-top: 1px solid #eee;
}

.input-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.voice-btn, .emoji-btn {
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  transition: all 0.2s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  -webkit-tap-highlight-color: transparent;
}

.voice-btn:hover, .emoji-btn:hover {
  background: #f0f0f0;
  transform: scale(1.1);
}

.voice-btn:active, .emoji-btn:active {
  background: #e0e0e0;
  transform: scale(0.95);
}

.text-input-wrapper {
  flex: 1;
}

.message-input {
  width: 100%;
  border: 1px solid #ddd;
  border-radius: 20px;
  padding: 10px 16px;
  font-size: 16px;
  outline: none;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  background: #fafafa;
}

.message-input:focus {
  border-color: #667eea;
  background: white;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
  transform: scale(1.02);
}

.message-input::placeholder {
  color: #999;
  transition: color 0.3s ease;
}

.message-input:focus::placeholder {
  color: #ccc;
}

/* 加载状态 */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  background: white;
  margin: 8px 16px;
  border-radius: 12px;
}

.loading-spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #e0e0e0;
  border-top: 3px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-container p {
  color: #666;
  font-size: 14px;
  margin: 0;
}

/* 错误状态 */
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  background: white;
  margin: 8px 16px;
  border-radius: 12px;
}

.error-icon {
  font-size: 32px;
  margin-bottom: 16px;
}

.error-container p {
  color: #666;
  font-size: 14px;
  margin: 0 0 16px 0;
  text-align: center;
}

.retry-btn {
  background: #667eea;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 20px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.retry-btn:hover {
  background: #5a6fd8;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .user-profile-card {
    margin: 8px 12px;
    padding: 12px;
  }
  
  .loading-container,
  .error-container {
    margin: 8px 12px;
    padding: 30px 16px;
  }
  
  .chat-messages {
    padding: 12px;
  }
  
  .chat-input-area {
    padding: 8px 12px;
  }
  
  .message-content {
    max-width: 80%;
  }
  
  .user-name {
    font-size: 16px;
  }
  
  .user-status {
    font-size: 11px;
  }
}
</style>