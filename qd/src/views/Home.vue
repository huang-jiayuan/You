<template>
  <div class="mobile-voice-home">
    <!-- 顶部用户信息区域 -->
    <div class="top-section">
      <div class="user-avatar">
        <img :src="userInfo.avatar || 'https://via.placeholder.com/40x40/4CAF50/ffffff?text=我'" :alt="userInfo.nickname" />
        <div class="online-indicator"></div>
      </div>
      <div class="top-actions">
        <button class="icon-btn">
          <span class="icon">👤</span>
          <span class="text">青少年</span>
        </button>
        <button class="icon-btn">
          <span class="icon">✓</span>
          <span class="text">签到</span>
        </button>
      </div>
    </div>

    <!-- 功能卡片区域 -->
    <div class="feature-cards">
      <!-- ME头像卡片 -->
      <div class="feature-card me-card">
        <div class="card-header">
          <h3>ME头像</h3>
          <span class="subtitle">许愿圆梦</span>
        </div>
        <div class="card-content">
          <div class="avatar-showcase">
            <div class="avatar-item">
              <img src="https://via.placeholder.com/40x40/ff6b9d/ffffff?text=西" alt="西法" />
              <span>西法</span>
            </div>
            <div class="avatar-item">
              <img src="https://via.placeholder.com/40x40/ff6b9d/ffffff?text=通" alt="通阿里" />
              <span>通阿里</span>
            </div>
          </div>
          <div class="card-badge">
            <span class="badge-text">神主生日</span>
            <span class="badge-count">x1</span>
          </div>
        </div>
      </div>

      <!-- 小哥哥卡片 -->
      <div class="feature-card brother-card">
        <div class="card-header">
          <h3>小哥哥</h3>
          <span class="subtitle">青春有你</span>
        </div>
        <div class="card-content">
          <div class="brother-avatar">
            <img src="https://via.placeholder.com/60x60/4facfe/ffffff?text=哥" alt="小哥哥" />
          </div>
          <div class="voice-controls">
            <button class="voice-btn" @click="togglePlay">
              <span v-if="!isPlaying">▶️</span>
              <span v-else>⏸️</span>
            </button>
            <button class="voice-btn">🎵</button>
            <button class="voice-btn">⏭️</button>
          </div>
        </div>
      </div>

      <!-- 家族聚会卡片 -->
      <div class="feature-card family-card">
        <div class="card-header">
          <h3>家族聚会</h3>
          <span class="subtitle">立即聊天</span>
        </div>
        <div class="card-content">
          <div class="family-icon">🍷</div>
          <button class="join-btn">去聊聊天</button>
        </div>
      </div>
    </div>

    <!-- 人气房间 -->
    <div class="popular-rooms">
      <div class="section-header">
        <h2>人气房间</h2>
      </div>
      <div class="rooms-grid">
        <div 
          v-for="room in popularRooms" 
          :key="room.id"
          class="room-card"
          @click="enterRoom(room.id)"
        >
          <div class="room-cover">
            <img :src="room.cover" :alt="room.name" />
            <div class="room-overlay">
              <div class="room-tag">{{ room.tag }}</div>
              <div class="room-count">{{ room.userCount }}</div>
            </div>
          </div>
          <div class="room-info">
            <h4>{{ room.name }}</h4>
          </div>
        </div>
      </div>
    </div>

    <!-- 广播/在线用户 -->
    <div class="broadcast-section">
      <div class="section-header">
        <h2>广播</h2>
        <span class="subtitle">在线的人</span>
        <button class="filter-btn">筛选</button>
      </div>
      <div class="broadcast-list">
        <div 
          v-for="user in onlineUsers" 
          :key="user.id"
          class="broadcast-item"
          @click="viewUserProfile(user.id)"
        >
          <div class="user-avatar">
            <img :src="user.avatar" :alt="user.nickname" />
            <div class="user-level">{{ user.level }}</div>
            <div v-if="user.isPlaying" class="playing-indicator">
              <span class="play-icon">⏸️</span>
            </div>
          </div>
          <div class="user-info">
            <div class="user-name">
              <span class="nickname">{{ user.nickname }}</span>
              <span class="age">{{ user.age }}岁</span>
              <span v-if="user.isVip" class="vip-badge">👑</span>
              <span v-if="user.statusEmoji" class="status-emoji">{{ user.statusEmoji }}</span>
            </div>
            <div class="user-status">
              <span class="status-text">{{ user.statusText }}</span>
            </div>
          </div>
          <div class="user-stats">
            <div class="stat-item">
              <span class="stat-icon">👥</span>
              <span class="stat-count">{{ user.followers }}</span>
            </div>
            <div class="stat-item">
              <span class="stat-icon">❤️</span>
              <span class="stat-count">{{ user.likes }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部导航栏 -->
    <div class="bottom-navigation">
      <div class="nav-item active" @click="navigateTo('home')">
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
      <div class="nav-item" @click="navigateTo('chat')">
        <span class="nav-icon">💬</span>
        <span class="nav-text">聊天</span>
      </div>
    </div>

    <!-- 浮动语音按钮 -->
    <div class="floating-voice-btn" @click="startVoiceChat">
      <span class="voice-icon">🎤</span>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '@/api'
import { useAudioPlayer } from '@/composables/mobile-chat/useAudioPlayer'
import { usePerformanceOptimization } from '@/composables/mobile-chat/usePerformanceOptimization'
import { useToast } from '@/composables/useToast'

export default {
  name: 'VoiceChatHome',
  setup() {
    const router = useRouter()
    
    // 使用组合式函数
    const audioPlayer = useAudioPlayer()
    const performanceOptimization = usePerformanceOptimization()
    const toast = useToast()
    
    // 响应式数据
    const currentTime = ref('')
    const userInfo = ref({
      id: null,
      nickname: '用户',
      avatar: '',
      level: 1,
      vipStatus: false
    })
    
    const popularRooms = ref([
      {
        id: 1,
        name: '夜长梦多',
        description: '友友连麦',
        cover: 'https://via.placeholder.com/80x80/667eea/ffffff?text=夜',
        tag: '热门',
        userCount: 'x10'
      },
      {
        id: 2,
        name: '友友连麦',
        description: '友友连麦',
        cover: 'https://via.placeholder.com/80x80/764ba2/ffffff?text=友',
        tag: '热门',
        userCount: 'x10'
      },
      {
        id: 3,
        name: '友友连麦',
        description: '友友连麦',
        cover: 'https://via.placeholder.com/80x80/f093fb/ffffff?text=连',
        tag: '热门',
        userCount: 'x10'
      },
      {
        id: 4,
        name: '友友连麦',
        description: '友友连麦',
        cover: 'https://via.placeholder.com/80x80/4facfe/ffffff?text=麦',
        tag: '热门',
        userCount: 'x10'
      }
    ])
    
    const isPlaying = ref(false)
    
    const onlineUsers = ref([
      {
        id: 1,
        nickname: '处对象，希望非',
        age: 33,
        avatar: 'https://via.placeholder.com/48x48/ff6b9d/ffffff?text=处',
        level: 4,
        statusText: '天友连麦-千青',
        statusEmoji: '😊',
        followers: 4,
        likes: 2,
        isVip: false,
        isPlaying: false
      },
      {
        id: 2,
        nickname: '没有节操的清欢',
        age: 69,
        avatar: 'https://via.placeholder.com/48x48/4facfe/ffffff?text=清',
        level: 1,
        statusText: '电台音乐',
        statusEmoji: '🎵',
        followers: 1,
        likes: 1,
        isVip: true,
        isPlaying: true
      },
      {
        id: 3,
        nickname: '茶',
        age: 15,
        avatar: 'https://via.placeholder.com/48x48/fa709a/ffffff?text=茶',
        level: 2,
        statusText: '电台音乐',
        statusEmoji: '🎵',
        followers: 2,
        likes: 0,
        isVip: false,
        isPlaying: false
      },
      {
        id: 4,
        nickname: '聊五块美金的清欢',
        age: 69,
        avatar: 'https://via.placeholder.com/48x48/667eea/ffffff?text=聊',
        level: 1,
        statusText: '天友连麦',
        statusEmoji: '💬',
        followers: 0,
        likes: 0,
        isVip: false,
        isPlaying: false
      }
    ])

    // 计算属性
    const isLoggedIn = computed(() => !!userInfo.value.id)

    // 方法
    const updateTime = () => {
      const now = new Date()
      currentTime.value = now.toLocaleTimeString('zh-CN', { 
        hour: '2-digit', 
        minute: '2-digit' 
      })
    }

    const loadUserInfo = async () => {
      try {
        const token = localStorage.getItem('access_token')
        if (token) {
          const response = await authAPI.getProfile()
          userInfo.value = {
            id: response.id,
            nickname: response.nickname || '用户',
            avatar: response.avatar || 'https://via.placeholder.com/50x50/4CAF50/ffffff?text=我',
            level: response.level || 1,
            vipStatus: response.vip_status === '1'
          }
        }
      } catch (error) {
        console.error('获取用户信息失败:', error)
        // 如果获取失败，可能需要重新登录
        if (error.message.includes('登录已过期')) {
          router.push('/')
        }
      }
    }

    const enterRoom = (roomId) => {
      console.log('进入房间:', roomId)
      // 这里可以导航到房间页面
      // router.push(`/room/${roomId}`)
    }

    const viewUserProfile = (userId) => {
      console.log('查看用户资料:', userId)
      // 这里可以导航到用户资料页面
      // router.push(`/profile/${userId}`)
    }

    const startVoiceChat = () => {
      console.log('开始语音聊天')
      // 这里可以打开语音聊天功能
    }

    const togglePlay = () => {
      isPlaying.value = !isPlaying.value
    }

    const navigateTo = (page) => {
      console.log('导航到:', page)
      switch(page) {
        case 'chat':
          router.push('/chat')
          break
        case 'about':
          router.push('/about')
          break
        case 'contact':
          router.push('/contact')
          break
        default:
          console.log('导航到:', page)
      }
    }

    // 生命周期
    let timeInterval = null

    onMounted(() => {
      updateTime()
      timeInterval = setInterval(updateTime, 1000)
      loadUserInfo()
    })

    onUnmounted(() => {
      if (timeInterval) {
        clearInterval(timeInterval)
      }
    })

    return {
      currentTime,
      userInfo,
      popularRooms,
      onlineUsers,
      isLoggedIn,
      isPlaying,
      enterRoom,
      viewUserProfile,
      startVoiceChat,
      togglePlay,
      navigateTo
    }
  }
}
</script>

<style scoped>
.mobile-voice-home {
  min-height: 100vh;
  background: linear-gradient(180deg, #f8f9ff 0%, #e8f0ff 100%);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  animation: fadeIn 0.5s ease-out;
  padding-bottom: 80px;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 顶部用户信息区域 */
.top-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
}

.user-avatar {
  position: relative;
  width: 50px;
  height: 50px;
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
  width: 12px;
  height: 12px;
  background: #4CAF50;
  border: 2px solid white;
  border-radius: 50%;
}

.top-actions {
  display: flex;
  gap: 12px;
}

.icon-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: #666;
  font-size: 12px;
  cursor: pointer;
}

.icon-btn .icon {
  font-size: 20px;
}

/* 功能卡片区域 */
.feature-cards {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  gap: 6px;
  padding: 0 16px 16px;
  height: 150px;
}

.feature-card {
  border-radius: 10px;
  padding: 8px;
  color: white;
  position: relative;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.2s ease;
  display: flex;
  flex-direction: column;
}

.feature-card:hover {
  transform: scale(1.02);
}

.me-card {
  background: linear-gradient(135deg, #ff6b9d 0%, #ff8a9b 100%);
  grid-column: 1;
  grid-row: 1 / 3;
}

.brother-card {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  grid-column: 2;
  grid-row: 1;
}

.family-card {
  background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
  grid-column: 2;
  grid-row: 2;
}

.card-header h3 {
  font-size: 14px;
  font-weight: 600;
  margin: 0 0 2px 0;
}

.card-header .subtitle {
  font-size: 10px;
  opacity: 0.8;
}

.card-content {
  margin-top: 6px;
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.avatar-showcase {
  display: flex;
  gap: 6px;
  margin-bottom: 6px;
}

.avatar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.avatar-item img {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  object-fit: cover;
}

.avatar-item span {
  font-size: 8px;
}

.card-badge {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  padding: 2px 6px;
  font-size: 10px;
  display: inline-block;
}

.brother-avatar {
  text-align: center;
  margin-bottom: 4px;
}

.brother-avatar img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.voice-controls {
  display: flex;
  gap: 4px;
  justify-content: center;
}

.voice-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 12px;
}

.family-icon {
  font-size: 32px;
  text-align: center;
  margin-bottom: 4px;
}

.join-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  border-radius: 12px;
  padding: 4px 8px;
  color: white;
  font-size: 12px;
  cursor: pointer;
  width: 100%;
}

/* 人气房间 */
.popular-rooms {
  padding: 0 16px 20px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.section-header h2 {
  font-size: 18px;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.rooms-grid {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  padding-bottom: 8px;
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
}

.rooms-grid::-webkit-scrollbar {
  height: 4px;
}

.rooms-grid::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
  border-radius: 2px;
}

.rooms-grid::-webkit-scrollbar-thumb {
  background: rgba(0, 0, 0, 0.3);
  border-radius: 2px;
}

.rooms-grid::-webkit-scrollbar-thumb:hover {
  background: rgba(0, 0, 0, 0.5);
}

.room-card {
  cursor: pointer;
  transition: transform 0.2s ease;
  flex-shrink: 0;
  width: 80px;
}

.room-card:hover {
  transform: scale(1.05);
}

.room-cover {
  position: relative;
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 8px;
}

.room-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.room-overlay {
  position: absolute;
  top: 8px;
  left: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.room-tag {
  background: #FFD700;
  color: #333;
  padding: 2px 6px;
  border-radius: 8px;
  font-size: 10px;
  font-weight: 600;
}

.room-count {
  background: rgba(0, 0, 0, 0.6);
  color: white;
  padding: 2px 6px;
  border-radius: 8px;
  font-size: 10px;
}

.room-info h4 {
  font-size: 12px;
  font-weight: 600;
  color: #333;
  margin: 0 0 2px 0;
}

.room-info p {
  font-size: 10px;
  color: #666;
  margin: 0;
}

/* 广播/在线用户 */
.broadcast-section {
  padding: 0 16px 100px;
}

.broadcast-section .section-header {
  align-items: center;
  gap: 8px;
}

.broadcast-section .section-header .subtitle {
  font-size: 14px;
  color: #666;
}

.filter-btn {
  background: none;
  border: 1px solid #ddd;
  border-radius: 16px;
  padding: 4px 12px;
  font-size: 12px;
  color: #666;
  cursor: pointer;
}

.broadcast-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.broadcast-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: white;
  border-radius: 12px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.broadcast-item:hover {
  background: #f8f9ff;
}

.broadcast-item .user-avatar {
  position: relative;
  width: 48px;
  height: 48px;
}

.broadcast-item .user-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.user-level {
  position: absolute;
  bottom: -2px;
  right: -2px;
  background: #4CAF50;
  color: white;
  font-size: 10px;
  padding: 2px 4px;
  border-radius: 8px;
  min-width: 16px;
  text-align: center;
}

.playing-indicator {
  position: absolute;
  bottom: -2px;
  left: -2px;
  background: #4facfe;
  color: white;
  font-size: 8px;
  padding: 2px 4px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.play-icon {
  font-size: 8px;
}

.user-info {
  flex: 1;
}

.user-name {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 4px;
}

.nickname {
  font-size: 14px;
  font-weight: 600;
  color: #333;
}

.age {
  font-size: 12px;
  color: #666;
}

.vip-badge {
  font-size: 12px;
}

.user-status {
  display: flex;
  align-items: center;
  gap: 4px;
}

.status-text {
  font-size: 12px;
  color: #666;
}

.status-emoji {
  font-size: 12px;
}

.user-stats {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
}

.stat-icon {
  font-size: 12px;
}

.stat-count {
  color: #666;
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

/* 浮动语音按钮 */
.floating-voice-btn {
  position: fixed;
  bottom: 80px;
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

.voice-icon {
  font-size: 16px;
}

.voice-text {
  font-size: 14px;
  font-weight: 600;
}

/* 响应式设计 - 默认浏览器大小，自适应屏幕 */

/* 默认桌面样式 - 不限制最大宽度 */
.mobile-voice-home {
  padding-bottom: 80px;
}

/* 小屏幕设备 (手机) */
@media (max-width: 768px) {
  .mobile-voice-home {
    padding-bottom: 80px;
  }
  
  .top-section {
    padding: 12px 16px;
  }
  
  .feature-cards {
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr 1fr;
    gap: 8px;
    padding: 0 16px 20px;
    height: 160px;
  }
  
  .me-card {
    grid-column: 1;
    grid-row: 1 / 3;
    padding: 12px;
  }
  
  .brother-card {
    grid-column: 2;
    grid-row: 1;
    padding: 10px;
  }
  
  .family-card {
    grid-column: 2;
    grid-row: 2;
    padding: 10px;
  }
  
  .card-header h3 {
    font-size: 16px;
    margin-bottom: 4px;
  }
  
  .card-header .subtitle {
    font-size: 12px;
  }
  
  .card-content {
    margin-top: 8px;
  }
  
  .avatar-showcase {
    margin-bottom: 8px;
  }
  
  .avatar-item img {
    width: 28px;
    height: 28px;
  }
  
  .avatar-item span {
    font-size: 10px;
  }
  
  .brother-avatar img {
    width: 50px;
    height: 50px;
  }
  
  .voice-controls {
    margin-top: 8px;
  }
  
  .voice-btn {
    width: 28px;
    height: 28px;
    font-size: 14px;
  }
  
  .family-icon {
    font-size: 40px;
    margin-bottom: 8px;
  }
  
  .join-btn {
    padding: 8px 12px;
    font-size: 14px;
    border-radius: 16px;
  }
  
  .rooms-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
  }
  
  .popular-rooms, .broadcast-section {
    padding: 0 16px 20px;
  }
  
  .broadcast-section {
    padding-bottom: 100px;
  }
  
  .floating-voice-btn {
    bottom: 90px;
    right: 16px;
    width: 56px;
    height: 56px;
    border-radius: 50%;
    padding: 0;
    justify-content: center;
  }
  
  .voice-text {
    display: none;
  }
}

/* 平板设备 */
@media (min-width: 769px) and (max-width: 1024px) {
  .mobile-voice-home {
    padding: 0 40px 80px;
  }
  
  .top-section {
    padding: 20px 0;
  }
  
  .user-avatar {
    width: 60px;
    height: 60px;
  }
  
  .feature-cards {
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr 1fr;
    gap: 20px;
    padding: 0 0 32px;
    height: 240px;
  }
  
  .me-card {
    grid-column: 1;
    grid-row: 1 / 3;
  }
  
  .brother-card {
    grid-column: 2;
    grid-row: 1;
  }
  
  .family-card {
    grid-column: 2;
    grid-row: 2;
  }
  
  .card-header h3 {
    font-size: 16px;
  }
  
  .card-header .subtitle {
    font-size: 12px;
  }
  
  .rooms-grid {
    grid-template-columns: repeat(4, 1fr);
    gap: 16px;
  }
  
  .popular-rooms, .broadcast-section {
    padding: 0 0 32px;
  }
  
  .broadcast-section {
    padding-bottom: 120px;
  }
  
  .section-header h2 {
    font-size: 20px;
  }
  
  .floating-voice-btn {
    bottom: 100px;
    right: 40px;
    padding: 16px 20px;
  }
  
  .voice-text {
    display: inline;
  }
}

/* 桌面设备 */
@media (min-width: 1025px) {
  .mobile-voice-home {
    padding: 0 60px 80px;
  }
  
  .top-section {
    padding: 24px 0;
  }
  
  .user-avatar {
    width: 60px;
    height: 60px;
  }
  
  .feature-cards {
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: 1fr;
    gap: 24px;
    padding: 0 0 40px;
    height: 200px;
  }
  
  .me-card, .brother-card, .family-card {
    grid-column: auto;
    grid-row: 1;
    min-height: auto;
  }
  
  .card-header h3 {
    font-size: 18px;
  }
  
  .card-header .subtitle {
    font-size: 14px;
  }
  
  .rooms-grid {
    grid-template-columns: repeat(6, 1fr);
    gap: 20px;
  }
  
  .popular-rooms, .broadcast-section {
    padding: 0 0 40px;
  }
  
  .broadcast-section {
    padding-bottom: 120px;
  }
  
  .section-header h2 {
    font-size: 22px;
  }
  
  .floating-voice-btn {
    bottom: 100px;
    right: 60px;
    padding: 16px 24px;
  }
  
  .voice-text {
    display: inline;
  }
}

/* 大屏幕桌面设备 */
@media (min-width: 1440px) {
  .mobile-voice-home {
    padding: 0 80px 80px;
  }
  
  .top-section {
    padding: 32px 0;
  }
  
  .feature-cards {
    gap: 32px;
    padding: 0 0 48px;
    height: 220px;
  }
  
  .rooms-grid {
    grid-template-columns: repeat(8, 1fr);
    gap: 24px;
  }
  
  .popular-rooms, .broadcast-section {
    padding: 0 0 48px;
  }
  
  .broadcast-section {
    padding-bottom: 140px;
  }
  
  .floating-voice-btn {
    bottom: 120px;
    right: 80px;
  }
}

/* 横屏适配 */
@media (orientation: landscape) and (max-height: 600px) {
  .mobile-voice-home {
    padding-bottom: 70px;
  }
  
  .top-section {
    padding: 8px 16px;
  }
  
  .feature-cards {
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: 1fr;
    gap: 12px;
    height: 120px;
    padding: 0 16px 16px;
  }
  
  .me-card, .brother-card, .family-card {
    grid-column: auto;
    grid-row: 1;
    min-height: auto;
  }
  
  .popular-rooms, .broadcast-section {
    padding: 0 16px 16px;
  }
  
  .broadcast-section {
    padding-bottom: 80px;
  }
  
  .floating-voice-btn {
    bottom: 75px;
    right: 16px;
    width: 48px;
    height: 48px;
    border-radius: 50%;
    padding: 0;
    justify-content: center;
  }
  
  .voice-text {
    display: none;
  }
  
  .bottom-navigation {
    padding: 6px 0;
  }
  
  .nav-item {
    padding: 6px 4px;
  }
  
  .nav-icon {
    font-size: 18px;
  }
  
  .nav-text {
    font-size: 9px;
  }
}

/* 触摸设备优化 */
@media (hover: none) and (pointer: coarse) {
  .feature-card:hover {
    transform: none;
  }
  
  .room-card:hover {
    transform: none;
  }
  
  .broadcast-item:hover {
    background: white;
  }
  
  .feature-card:active {
    transform: scale(0.98);
  }
  
  .room-card:active {
    transform: scale(0.95);
  }
}

/* 高对比度模式支持 */
@media (prefers-contrast: high) {
  .feature-card {
    border: 2px solid rgba(255, 255, 255, 0.5);
  }
  
  .broadcast-item {
    border: 1px solid #ccc;
  }
}

/* 减少动画偏好支持 */
@media (prefers-reduced-motion: reduce) {
  .feature-card,
  .room-card,
  .broadcast-item,
  .floating-voice-btn {
    transition: none;
  }
}
</style>