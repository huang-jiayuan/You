<template>
  <div class="chat-page">
    <!-- 顶部导航 -->
    <div class="chat-header">
      <div class="header-tabs">
        <div class="tab-item active">聊天</div>
        <div class="tab-item">好友</div>
      </div>
    </div>

    <!-- 消息通知卡片 -->
    <div class="notification-card">
      <div class="notification-content">
        <div class="notification-text">
          <div class="notification-title">打开消息通知，不再错过TA的问好</div>
          <div class="notification-subtitle">支持应用提升300%</div>
        </div>
        <button class="notification-btn">去开启</button>
      </div>
    </div>

    <!-- 功能按钮区域 -->
    <div class="function-buttons">
      <div class="function-item">
        <div class="function-icon orange">
          <span>👤</span>
        </div>
        <span class="function-text">最近来往</span>
      </div>
      <div class="function-item">
        <div class="function-icon pink">
          <span>💬</span>
        </div>
        <span class="function-text">收礼记录</span>
      </div>
      <div class="function-item">
        <div class="function-icon purple">
          <span>👁️</span>
        </div>
        <span class="function-text">谁看过我</span>
      </div>
    </div>

    <!-- 聊天列表 -->
    <div class="chat-list">
      <!-- 家族广场 -->
      <div class="chat-item special-item">
        <div class="chat-avatar family-avatar">
          <span>🏠</span>
        </div>
        <div class="chat-info">
          <div class="chat-name">家族广场</div>
          <div class="chat-preview">参与家族动态，可领珍稀道具奖励</div>
        </div>
      </div>

      <!-- 聊天记录列表 -->
      <div 
        v-for="chat in chatList" 
        :key="chat.id"
        class="chat-item"
        :class="{ 
          'chat-item-navigating': isNavigating && navigatingChatId === chat.id,
          'chat-item-disabled': isNavigating
        }"
        @click="openChat(chat.id)"
      >
        <div class="chat-avatar">
          <img :src="chat.avatar" :alt="chat.name" />
          <div v-if="chat.isOnline" class="online-status">
            <span class="status-text">[在线]</span>
          </div>
        </div>
        <div class="chat-info">
          <div class="chat-name">{{ chat.name }}</div>
          <div class="chat-preview" :class="{ 'online-preview': chat.isOnline }">
            <span v-if="chat.isOnline" class="online-tag">[在线]</span>
            {{ chat.lastMessage }}
          </div>
        </div>
        <div class="chat-meta">
          <div class="chat-time">{{ chat.time }}</div>
          <div v-if="chat.unreadCount > 0" class="unread-badge">
            {{ chat.unreadCount }}
          </div>
          <!-- 导航加载指示器 -->
          <div v-if="isNavigating && navigatingChatId === chat.id" class="loading-indicator">
            <div class="loading-spinner"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部导航栏 -->
    <div class="bottom-navigation">
      <div class="nav-item" @click="navigateTo('home')">
        <span class="nav-icon">🏠</span>
        <span class="nav-text">首页</span>
      </div>
      <div class="nav-item" @click="navigateTo('room')">
        <span class="nav-icon">🎙️</span>
        <span class="nav-text">房间</span>
      </div>
      <div class="nav-item" @click="navigateTo('dynamic')">
        <span class="nav-icon">⭐</span>
        <span class="nav-text">动态</span>
      </div>
      <div class="nav-item active" @click="navigateTo('chat')">
        <span class="nav-icon">💬</span>
        <span class="nav-text">聊天</span>
        <div v-if="totalUnreadCount > 0" class="notification-badge">{{ totalUnreadCount }}</div>
      </div>
    </div>

    <!-- 浮动消息按钮 -->
    <div class="floating-message-btn" @click="showMessageOptions">
      <span class="message-icon">💬</span>
      <span class="message-text">寻友</span>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '@/composables/useToast'
import { useNavigationError } from '@/composables/useNavigationError'

export default {
  name: 'ChatPage',
  setup() {
    const router = useRouter()
    const { error: showError, success: showSuccess } = useToast()
    const { 
      handleNavigationError, 
      handleParameterError, 
      showSuccessMessage 
    } = useNavigationError()
    
    // 导航状态管理
    const isNavigating = ref(false)
    const navigatingChatId = ref(null)
    const navigationError = ref(null)
    
    // 聊天列表数据
    const chatList = ref([
      {
        id: 1,
        name: '高Sir',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23333333"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E高%3C/text%3E%3C/svg%3E',
        lastMessage: '帅子声的小仙女！',
        time: '刚刚',
        unreadCount: 0,
        isOnline: true
      },
      {
        id: 2,
        name: '进',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%234facfe"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E进%3C/text%3E%3C/svg%3E',
        lastMessage: 'hi小可爱，晚安喽呀',
        time: '刚刚',
        unreadCount: 0,
        isOnline: true
      },
      {
        id: 3,
        name: '再再',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23fa709a"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E再%3C/text%3E%3C/svg%3E',
        lastMessage: '接触觉你挺特别',
        time: '2分钟前',
        unreadCount: 0,
        isOnline: false
      },
      {
        id: 4,
        name: '想要了',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23667eea"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E想%3C/text%3E%3C/svg%3E',
        lastMessage: '可以聊聊吗美女',
        time: '3分钟前',
        unreadCount: 0,
        isOnline: true
      },
      {
        id: 5,
        name: '小m',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23ff6b9d"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E小%3C/text%3E%3C/svg%3E',
        lastMessage: '我觉觉你挺特别',
        time: '2分钟前',
        unreadCount: 0,
        isOnline: true
      },
      {
        id: 6,
        name: '峰俊男',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23764ba2"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E峰%3C/text%3E%3C/svg%3E',
        lastMessage: '想要吗',
        time: '4分钟前',
        unreadCount: 0,
        isOnline: false
      },
      {
        id: 7,
        name: '榊涂蛋',
        avatar: 'data:image/svg+xml,%3Csvg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="24" cy="24" r="24" fill="%23f093fb"/%3E%3Ctext x="24" y="28" font-family="Arial" font-size="16" font-weight="bold" fill="white" text-anchor="middle"%3E榊%3C/text%3E%3C/svg%3E',
        lastMessage: '前新小仙女你好呢一',
        time: '',
        unreadCount: 1,
        isOnline: true
      }
    ])

    // 计算总未读消息数
    const totalUnreadCount = computed(() => {
      return chatList.value.reduce((total, chat) => total + chat.unreadCount, 0)
    })

    // 本地错误处理函数（包装全局错误处理）
    const handleChatNavigationError = (error, chatUser) => {
      // 使用全局错误处理函数
      handleNavigationError(error, { userName: chatUser?.name })
      
      // 重置导航状态
      isNavigating.value = false
      navigatingChatId.value = null
      navigationError.value = error.message || '导航失败'
    }

    // 方法
    const openChat = async (chatId) => {
      // 防止重复点击
      if (isNavigating.value) {
        console.log('导航正在进行中，忽略重复点击')
        return
      }

      // 参数验证
      if (!chatId) {
        handleParameterError('chatId', chatId)
        return
      }

      console.log('打开聊天:', chatId)
      
      // 查找对应的聊天用户信息
      const chatUser = chatList.value.find(chat => chat.id === chatId)
      if (!chatUser) {
        console.error('未找到聊天用户信息')
        handleParameterError('用户信息', chatId)
        return
      }
      
      // 设置导航状态
      isNavigating.value = true
      navigatingChatId.value = chatId
      navigationError.value = null
      
      try {
        // 构建导航参数
        const params = { userId: String(chatId) }
        const query = {
          name: chatUser.name,
          isOnline: String(chatUser.isOnline)
          // 不传递avatar，避免URL过长问题
        }
        
        // 调试信息
        console.log('准备导航到:', {
          name: 'ChatDetail',
          params,
          query
        })

        // 执行导航
        await router.push({
          name: 'ChatDetail',
          params,
          query
        })
        
        console.log('导航成功:', { params, query })
        showSuccessMessage(`正在打开与${chatUser.name}的聊天`)
        
      } catch (error) {
        handleChatNavigationError(error, chatUser)
      } finally {
        // 延迟重置状态，确保页面转换完成
        setTimeout(() => {
          isNavigating.value = false
          navigatingChatId.value = null
        }, 300)
      }
    }

    const showMessageOptions = () => {
      console.log('显示消息选项')
      // 这里可以显示消息选项弹窗
    }

    const navigateTo = (page) => {
      console.log('导航到:', page)
      switch(page) {
        case 'home':
          router.push('/home')
          break
        case 'room':
          // router.push('/room')
          break
        case 'dynamic':
          // router.push('/dynamic')
          break
        case 'chat':
          // 已经在聊天页面
          break
        default:
          console.log('导航到:', page)
      }
    }

    onMounted(() => {
      // 页面加载时的初始化操作
      console.log('聊天页面已加载')
    })

    return {
      chatList,
      totalUnreadCount,
      isNavigating,
      navigatingChatId,
      navigationError,
      openChat,
      showMessageOptions,
      navigateTo
    }
  }
}
</script>

<style scoped>
.chat-page {
  min-height: 100vh;
  background: #f8f9ff;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  padding-bottom: 80px;
}

/* 顶部导航 */
.chat-header {
  background: white;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
}

.header-tabs {
  display: flex;
  justify-content: center;
  gap: 40px;
}

.tab-item {
  font-size: 18px;
  font-weight: 500;
  color: #999;
  cursor: pointer;
  position: relative;
  padding-bottom: 8px;
}

.tab-item.active {
  color: #333;
  font-weight: 600;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 20px;
  height: 3px;
  background: #4CAF50;
  border-radius: 2px;
}

/* 消息通知卡片 */
.notification-card {
  margin: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 16px;
  color: white;
}

.notification-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.notification-text {
  flex: 1;
}

.notification-title {
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 4px;
}

.notification-subtitle {
  font-size: 12px;
  opacity: 0.8;
}

.notification-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  color: white;
  font-size: 14px;
  cursor: pointer;
  backdrop-filter: blur(10px);
}

/* 功能按钮区域 */
.function-buttons {
  display: flex;
  justify-content: space-around;
  padding: 20px 16px;
  background: white;
  margin: 0 16px;
  border-radius: 12px;
  margin-bottom: 16px;
}

.function-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.function-icon {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.function-icon.orange {
  background: linear-gradient(135deg, #ff9a56 0%, #ff6b35 100%);
}

.function-icon.pink {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8a9b 100%);
}

.function-icon.purple {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.function-text {
  font-size: 12px;
  color: #666;
}

/* 聊天列表 */
.chat-list {
  background: white;
  margin: 0 16px;
  border-radius: 12px;
  overflow: hidden;
}

.chat-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #f8f9ff;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  position: relative;
  min-height: 44px; /* 确保触摸友好的最小高度 */
  overflow: hidden;
  -webkit-tap-highlight-color: transparent; /* 移除移动端点击高亮 */
}

.chat-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(76, 175, 80, 0.1), transparent);
  transition: left 0.5s ease;
}

.chat-item:hover {
  background: #f8f9ff;
  transform: translateX(4px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.chat-item:hover::before {
  left: 100%;
}

.chat-item:active {
  background: #e8f5e8;
  transform: scale(0.98) translateX(2px);
  transition: all 0.1s ease;
}

/* 移动端触摸优化 */
@media (hover: none) and (pointer: coarse) {
  .chat-item:hover {
    transform: none;
    box-shadow: none;
  }
  
  .chat-item:active {
    background: #e8f5e8;
    transform: scale(0.96);
  }
}

/* 导航状态样式 */
.chat-item-navigating {
  background: linear-gradient(135deg, #e8f5e8 0%, #f0f8f0 100%);
  position: relative;
  transform: scale(0.98);
  box-shadow: inset 0 0 0 2px rgba(76, 175, 80, 0.3);
}

.chat-item-navigating::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    90deg, 
    transparent 0%, 
    rgba(76, 175, 80, 0.2) 50%, 
    transparent 100%
  );
  animation: loading-shimmer 1.2s ease-in-out infinite;
}

.chat-item-disabled {
  pointer-events: none;
  opacity: 0.5;
  filter: grayscale(0.3);
  transition: all 0.3s ease;
}

@keyframes loading-shimmer {
  0% {
    transform: translateX(-100%) skewX(-15deg);
  }
  50% {
    transform: translateX(0%) skewX(-15deg);
  }
  100% {
    transform: translateX(100%) skewX(-15deg);
  }
}

/* 脉冲动画用于加载指示器 */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.7;
    transform: scale(1.05);
  }
}

.chat-item:last-child {
  border-bottom: none;
}

.special-item {
  background: linear-gradient(135deg, rgba(76, 175, 80, 0.1) 0%, rgba(139, 195, 74, 0.1) 100%);
}

.chat-avatar {
  position: relative;
  width: 48px;
  height: 48px;
  margin-right: 12px;
  flex-shrink: 0;
}

.chat-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.family-avatar {
  background: linear-gradient(135deg, #4CAF50 0%, #8BC34A 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.online-status {
  position: absolute;
  bottom: -2px;
  left: -2px;
  background: #4CAF50;
  color: white;
  font-size: 8px;
  padding: 1px 3px;
  border-radius: 6px;
}

.status-text {
  font-size: 8px;
}

.chat-info {
  flex: 1;
  min-width: 0;
}

.chat-name {
  font-size: 16px;
  font-weight: 500;
  color: #333;
  margin-bottom: 4px;
}

.chat-preview {
  font-size: 14px;
  color: #666;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.online-preview {
  color: #4CAF50;
}

.online-tag {
  color: #4CAF50;
  font-size: 12px;
}

.chat-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 4px;
}

.chat-time {
  font-size: 12px;
  color: #999;
}

.unread-badge {
  background: #ff4444;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 16px;
  text-align: center;
}

/* 加载指示器 */
.loading-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-top: 4px;
  animation: pulse 1.5s ease-in-out infinite;
}

.loading-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(76, 175, 80, 0.2);
  border-top: 2px solid #4CAF50;
  border-right: 2px solid #4CAF50;
  border-radius: 50%;
  animation: spin 0.8s cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite;
}

@keyframes spin {
  0% { 
    transform: rotate(0deg) scale(1);
  }
  50% {
    transform: rotate(180deg) scale(1.1);
  }
  100% { 
    transform: rotate(360deg) scale(1);
  }
}

/* 底部导航栏 */
.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  background: white;
  border-top: 1px solid #eee;
  padding: 8px 0;
  z-index: 100;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  padding: 8px;
  cursor: pointer;
  position: relative;
}

.nav-item.active {
  color: #4CAF50;
}

.nav-icon {
  font-size: 20px;
}

.nav-text {
  font-size: 10px;
}

.notification-badge {
  position: absolute;
  top: 4px;
  right: 20px;
  background: #ff4444;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 16px;
  text-align: center;
}

/* 浮动消息按钮 */
.floating-message-btn {
  position: fixed;
  bottom: 90px;
  right: 16px;
  background: linear-gradient(135deg, #ff6b35 0%, #f7931e 100%);
  color: white;
  border-radius: 24px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  box-shadow: 0 4px 12px rgba(255, 107, 53, 0.3);
  z-index: 99;
}

.message-icon {
  font-size: 16px;
}

.message-text {
  font-size: 14px;
  font-weight: 600;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .chat-header {
    padding: 12px 16px;
  }
  
  .notification-card {
    margin: 12px;
    padding: 12px;
  }
  
  .function-buttons {
    padding: 16px;
    margin: 0 12px 12px;
  }
  
  .function-icon {
    width: 40px;
    height: 40px;
    font-size: 18px;
  }
  
  .chat-list {
    margin: 0 12px;
  }
  
  .chat-item {
    padding: 12px;
    min-height: 48px; /* 移动端更大的触摸目标 */
  }
  
  .floating-message-btn {
    bottom: 90px;
    right: 20px;
    width: 56px;
    height: 56px;
    border-radius: 50%;
    padding: 0;
    justify-content: center;
  }
  
  .message-text {
    display: none;
  }
}
</style>