<template>
  <div class="mobile-voice-home">
    <!-- 顶部用户信息区域 -->
    <div class="top-section">
      <div class="user-avatar" @click="showUserSidebar">
        <img :src="userInfo.avatar || generateAvatar('我', '4CAF50', 40)" :alt="userInfo.nickname" />
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
              <img :src="generateAvatar('西', 'ff6b9d', 40)" alt="西法" />
              <span>西法</span>
            </div>
            <div class="avatar-item">
              <img :src="generateAvatar('通', 'ff6b9d', 40)" alt="通阿里" />
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
            <img :src="generateAvatar('哥', '4facfe', 60)" alt="小哥哥" />
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
        <div class="header-actions">
          <button class="search-btn" @click="showSearch = !showSearch">
            <span class="icon">🔍</span>
          </button>
          <button class="filter-btn" @click="showTagFilter = !showTagFilter">筛选</button>
        </div>
      </div>

      <!-- 搜索框 -->
      <div v-if="showSearch" class="search-container">
        <input
          v-model="searchKeyword"
          type="text"
          placeholder="搜索房间名称或房主名字..."
          class="search-input"
          @input="searchRooms(searchKeyword)"
          @keyup.enter="searchRooms(searchKeyword)"
        />
        <button v-if="searchKeyword" @click="searchKeyword = ''; loadRecommendRooms()" class="clear-search">
          ✕
        </button>
      </div>

      <!-- 标签筛选 -->
      <div class="tag-filter">
        <div class="tag-list">
          <button
            v-for="tag in roomTags"
            :key="tag.id"
            :class="['tag-item', { active: selectedTag === tag.id }]"
            @click="loadRoomsByTag(tag.id)"
          >
            {{ tag.name }}
          </button>
        </div>
      </div>

      <!-- 加载状态 -->
      <div v-if="roomsLoading" class="loading-container">
        <div class="loading-spinner"></div>
        <span>加载中...</span>
      </div>

      <!-- 错误状态 -->
      <div v-else-if="roomsError" class="error-container">
        <div class="error-message">{{ roomsError }}</div>
        <button class="retry-btn" @click="retryLoadRooms">重试</button>
      </div>

      <!-- 房间列表 -->
      <div v-else-if="popularRooms.length > 0" class="rooms-list">
        <div 
          v-for="room in popularRooms" 
          :key="room.id || room.room_id"
          class="room-item"
          @click="console.log('点击了房间卡片!'); enterRoom(room.id || room.room_id)"
          style="cursor: pointer;"
        >
          <!-- 房间封面 -->
          <div class="room-cover">
            <img
              :src="room.cover || generateDefaultCover(room.room_name || room.name || '房间')"
              :alt="room.room_name || room.name || '房间'"
              @error="handleImageError"
            />
            <div class="room-count">{{ formatUserCount(room.user_count || room.fk_member_room || 0) }}</div>
          </div>

          <!-- 房间信息 -->
          <div class="room-content">
            <!-- 房间标题 -->
            <div class="room-title">
              <span class="room-icon">🔥</span>
              <h4>{{ room.room_name || room.name || '未命名房间' }}</h4>
              <span class="room-emoji">✨</span>
            </div>

            <!-- 房间标签 -->
            <div class="room-tags">
              <span class="room-tag">{{ getRoomTagName(room) }}</span>
              <span class="room-status">🎵 梦幻邮轮中</span>
            </div>

            <!-- 房主信息 -->
            <div class="room-owner">
              <span class="owner-name">{{ room.owner_nickname || '房主' }}</span>
            </div>

            <!-- 在线用户头像 -->
            <div class="room-users">
              <div
                v-for="i in Math.min(4, Math.max(1, Math.floor((room.fk_member_room || 0) / 30)))"
                :key="i"
                class="user-avatar"
              >
                <img :src="generateAvatar(`用户${i}`, getRandomColor(), 24)" :alt="`用户${i}`" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 无数据状态 -->
      <div v-else class="empty-container">
        <div class="empty-icon">🏠</div>
        <div class="empty-message">
          {{ searchKeyword ? '未找到相关房间' : '暂无房间数据' }}
        </div>
        <button v-if="searchKeyword" @click="searchKeyword = ''; loadRecommendRooms()" class="reset-btn">
          查看推荐房间
        </button>
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

    <!-- 用户侧边栏 -->
    <UserSidebar
      :isVisible="sidebarVisible"
      :userInfo="sidebarUserInfo"
      @close="hideUserSidebar"
    />

  </div>
  
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '@/api/auth'
import roomAPI from '@/api/room'
import { useAudioPlayer } from '@/composables/mobile-chat/useAudioPlayer'
import { usePerformanceOptimization } from '@/composables/mobile-chat/usePerformanceOptimization'
import { useToast } from '@/composables/useToast'
import UserSidebar from '@/components/UserSidebar.vue'

const router = useRouter()
const { showToast } = useToast()


    

    
    // 响应式数据
    const currentTime = ref('')
    const userInfo = ref({
      id: null,
      nickname: '用户',
      avatar: '',
      level: 1,
      vipStatus: false
    })
    
    // 房间相关数据
    const popularRooms = ref([])
    const roomsLoading = ref(false)
    const roomsError = ref(null)
    const searchKeyword = ref('')
    const selectedTag = ref(null)
    const roomTags = ref([{
      id: 1,
      name: '娱乐',
      color: '#4facfe'
    }, {
      id: 2, 
      name: '才艺',
      color: '#f093fb'
    }, {
      id: 3,
      name: '交友速配', 
      color: '#fa709a'
    }, {
      id: 4,
      name: '音乐',
      color: '#764ba2'
    }, {
      id: 5, 
      name: '聊天',
      color: '#ff6b9d'
    }, {
      id: 6,
      name: '陪伴',
      color: '#667eea'
    }])
    const showSearch = ref(false)
    const showTagFilter = ref(false)
    const showDebug = ref(false)
    
    const isPlaying = ref(false)
    const sidebarVisible = ref(false)

    // 侧边栏用户信息
    const sidebarUserInfo = computed(() => ({
      nickname: userInfo.value.nickname || '途场',
      meId: '201691465',
      avatar: userInfo.value.avatar || 'https://via.placeholder.com/80x80/4CAF50/ffffff?text=途',
      level: userInfo.value.level || 0,
      following: 1,
      followers: 1,
      coins: 0,
      balance: '0.00',
      teacherStats: {
        disciples: 1,
        hearts: 1
      }
    }))

    // 生成本地 SVG 头像
    const generateAvatar = (text, color = '4CAF50', size = 50) => {
      const svg = `
        <svg width="${size}" height="${size}" xmlns="http://www.w3.org/2000/svg">
          <rect width="${size}" height="${size}" fill="#${color}" rx="${size/10}"/>
          <text x="${size/2}" y="${size/2 + size/8}" font-family="Arial, sans-serif"
                font-size="${size/2.5}" font-weight="bold" text-anchor="middle"
                dominant-baseline="middle" fill="white">
            ${text}
          </text>
        </svg>
      `
      return `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`
    }

const broadcastUsers = ref([
      {
        id: 1,
        nickname: '处对象，希望非',
        age: 33,
        avatar: generateAvatar('处', 'ff6b9d', 48),
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
        avatar: generateAvatar('清', '4facfe', 48),
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
        avatar: generateAvatar('茶', 'fa709a', 48),
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
        avatar: generateAvatar('聊', '667eea', 48),
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
          // 暂时注释掉API调用，避免404错误
          // const response = await authAPI.getProfile()

          // 使用模拟数据，等后端接口准备好后再启用
          userInfo.value = {
            id: 1,
            nickname: '用户',
            avatar: 'data:image/svg+xml,%3Csvg width="50" height="50" viewBox="0 0 50 50" fill="none" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="25" cy="25" r="25" fill="%234CAF50"/%3E%3Ctext x="25" y="30" font-family="Arial" font-size="18" font-weight="bold" fill="white" text-anchor="middle"%3E我%3C/text%3E%3C/svg%3E',
            level: 1,
            vipStatus: false
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

    // 加载推荐房间数据
    const loadRecommendRooms = async () => {
      if (roomsLoading.value) return

      try {
        roomsLoading.value = true
        roomsError.value = null

        console.log('开始加载推荐房间数据...')

        const response = await roomAPI.getRecommendRooms(1, 10)
        console.log('收到推荐房间响应:', response)

        // 处理响应数据
        let roomsData = []
        if (response && response.code === 200) {
          if (response.data && response.data.rooms) {
            roomsData = response.data.rooms
          } else if (response.data && Array.isArray(response.data)) {
            roomsData = response.data
          }
        }

        popularRooms.value = roomsData
        console.log('推荐房间数据:', popularRooms.value)

        // 重置筛选状态
        selectedTag.value = null

      } catch (error) {
        console.error('加载推荐房间失败:', error)
        roomsError.value = error.message || '加载房间数据失败'
        popularRooms.value = []
      } finally {
        roomsLoading.value = false
      }
    }

    // 格式化用户数量显示的辅助函数
    const formatUserCount = (count) => {
      const num = parseInt(count) || 0
      if (num === 0) return 'x0'
      if (num >= 1000) return `x${(num / 1000).toFixed(1)}k`
      return `x${num}`
    }

    // 生成默认房间封面
    const generateDefaultCover = (roomName = '房') => {
      const colors = ['667eea', '764ba2', 'f093fb', '4facfe', 'ff6b9d', 'fa709a', 'fee140']
      const colorIndex = Math.abs(hashCode(roomName)) % colors.length
      const color = colors[colorIndex]
      const firstChar = roomName.charAt(0) || '房'

      const svg = `
        <svg width="80" height="80" xmlns="http://www.w3.org/2000/svg">
          <rect width="80" height="80" fill="#${color}" rx="8"/>
          <text x="40" y="50" font-family="Arial, sans-serif" font-size="24" font-weight="bold"
                text-anchor="middle" dominant-baseline="middle" fill="white">
            ${firstChar}
          </text>
        </svg>
      `

      return `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`
    }

    // 字符串哈希函数
    const hashCode = (str) => {
      let hash = 0
      if (str.length === 0) return hash

      for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i)
        hash = ((hash << 5) - hash) + char
        hash = hash & hash
      }

      return Math.abs(hash)
    }

    // 获取房间标签名称
    const getRoomTagName = (room) => {
      // 优先使用后端返回的标签名称
      if (room.tag_name) {
        return room.tag_name
      }

      // 如果有room_type，根据类型映射
      if (room.room_type) {
        const tagMap = {
          '1': '交友速配',
          '2': '才艺',
          '3': '点唱',
          '4': '电台音乐'
        }
        return tagMap[room.room_type] || '热门'
      }

      // 默认返回
      return '热门'
    }

    // 获取随机颜色
    const getRandomColor = () => {
      const colors = ['ff6b9d', '4facfe', 'fa709a', 'f093fb', '764ba2', '667eea', '4CAF50', 'FF9800']
      return colors[Math.floor(Math.random() * colors.length)]
    }

    // 根据标签加载房间
    const loadRoomsByTag = async (tagId) => {
      if (roomsLoading.value) return

      try {
        roomsLoading.value = true
        roomsError.value = null

        console.log('开始根据标签加载房间，标签ID:', tagId)

        let response
        if (tagId === null || tagId === 0) {
          // 加载推荐房间
          console.log('加载推荐房间...')
          response = await roomAPI.getRecommendRooms(1, 10)
        } else {
          // 根据标签加载房间
          console.log('根据标签加载房间，标签ID:', tagId)
          response = await roomAPI.getRoomsByCategory(tagId, 1, 10)
        }

        console.log('标签筛选响应:', response)

        // 处理响应数据
        let roomsData = []
        if (response && response.code === 200) {
          if (response.data && response.data.rooms) {
            roomsData = response.data.rooms
          } else if (response.data && Array.isArray(response.data)) {
            roomsData = response.data
          } else if (response.rooms) {
            roomsData = response.rooms
          }
        }

        // 确保数据是数组格式
        if (!Array.isArray(roomsData)) {
          console.warn('标签筛选返回的数据不是数组格式:', roomsData)
          roomsData = []
        }

        popularRooms.value = roomsData
        selectedTag.value = tagId

        console.log('标签筛选结果:', popularRooms.value)
        console.log('标签筛选结果数量:', popularRooms.value.length)

      } catch (error) {
        console.error('加载分类房间失败:', error)
        roomsError.value = error.message || '加载房间数据失败'
        popularRooms.value = []
      } finally {
        roomsLoading.value = false
      }
    }

    // 搜索防抖定时器
    let searchTimer = null

    // 搜索房间（带防抖）
    const searchRooms = (keyword) => {
      // 清除之前的定时器
      if (searchTimer) {
        clearTimeout(searchTimer)
      }

      // 设置新的定时器
      searchTimer = setTimeout(async () => {
        await performSearch(keyword)
      }, 500) // 500ms 防抖延迟
    }

    // 执行搜索
    const performSearch = async (keyword) => {
      if (!keyword || keyword.trim().length === 0) {
        // 如果搜索关键词为空，重新加载推荐房间
        console.log('搜索关键词为空，加载推荐房间')
        await loadRecommendRooms()
        return
      }

      if (roomsLoading.value) return

      try {
        roomsLoading.value = true
        roomsError.value = null

        console.log('开始搜索房间，关键词:', keyword.trim())

        const response = await roomAPI.searchRooms(keyword.trim(), 1, 10)

        console.log('搜索房间响应:', response)

        // 处理响应数据
        let roomsData = []
        if (response && response.code === 200) {
          if (response.data && response.data.rooms) {
            roomsData = response.data.rooms
          } else if (response.data && Array.isArray(response.data)) {
            roomsData = response.data
          }
        }

        popularRooms.value = roomsData

        console.log('搜索结果:', popularRooms.value)

      } catch (error) {
        console.error('搜索房间失败:', error)
        roomsError.value = error.message || '搜索房间失败'
        popularRooms.value = []
      } finally {
        roomsLoading.value = false
      }
    }

    // 加载房间标签
    const loadRoomTags = async () => {
      try {
        console.log('开始加载房间标签...')
        const response = await roomAPI.getRoomTags()
        console.log('房间标签响应:', response)

        // 始终确保热门标签在第一位
        let backendTags = []
        if (response && response.code === 200 && response.data) {
          backendTags = response.data
        } else {
          // 使用默认标签
          backendTags = [
            { id: 1, name: '娱乐', color: '#4facfe' },
            { id: 2, name: '才艺', color: '#f093fb' },
            { id: 3, name: '交友速配', color: '#fa709a' },
            { id: 4, name: '音乐', color: '#764ba2' },
            { id: 5, name: '聊天', color: '#ff6b9d' },
            { id: 6, name: '陪伴', color: '#667eea' }
          ]
        }

        // 确保热门标签始终在第一位
        roomTags.value = [
          { id: 0, name: '热门', color: '#FF6B35' },
          ...backendTags
        ]

        console.log('房间标签数据:', roomTags.value)

        // 默认选中热门标签
        if (roomTags.value.length > 0) {
          selectedTag.value = 0
        }

      } catch (error) {
        console.error('加载房间标签失败:', error)
        // 使用默认标签，确保热门标签在第一位
        const defaultTags = [
          { id: 1, name: '娱乐', color: '#4facfe' },
          { id: 2, name: '才艺', color: '#f093fb' },
          { id: 3, name: '交友速配', color: '#fa709a' },
          { id: 4, name: '音乐', color: '#764ba2' },
          { id: 5, name: '聊天', color: '#ff6b9d' },
          { id: 6, name: '陪伴', color: '#667eea' }
        ]

        roomTags.value = [
          { id: 0, name: '热门', color: '#FF6B35' },
          ...defaultTags
        ]

        // 默认选中热门标签
        if (roomTags.value.length > 0) {
          selectedTag.value = 0
        }
      }
    }

    // 重试加载房间数据
    const retryLoadRooms = () => {
      if (searchKeyword.value) {
        searchRooms(searchKeyword.value)
      } else if (selectedTag.value !== null) {
        loadRoomsByTag(selectedTag.value)
      } else {
        loadRecommendRooms()
      }
    }

    const enterRoom = async (roomId) => {
      console.log('🔥 enterRoom函数被调用，房间ID:', roomId)
      alert(`点击了房间 ${roomId}`)

      // 直接跳转，先不调用API
      console.log('🚀 直接跳转到房间页面')
      router.push(`/room/${roomId}`)
    }

    const viewUserProfile = (userId) => {
      console.log('查看用户资料:', userId)
      // 这里可以导航到用户资料页面
      // router.push(`/profile/${userId}`)
    }

    const startVoiceChat = () => {
      console.log('点击创建房间按钮')
      // 跳转到创建房间页面
      router.push('/create-room')
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

    const showUserSidebar = () => {
      showSidebar.value = true
    }

    const hideUserSidebar = () => {
      sidebarVisible.value = false
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      const img = event.target
      const roomName = img.alt || '房间'
      img.src = generateDefaultCover(roomName)
    }

// 生命周期
onMounted(async () => {
  updateTime()
  setInterval(updateTime, 60000)
  
  await loadUserInfo()
  await loadRecommendRooms()
})

onUnmounted(() => {
  if (timeInterval) {
    clearInterval(timeInterval)
  }

  // 清除搜索定时器
  if (searchTimer) {
    clearTimeout(searchTimer)
  }
})
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
  cursor: pointer;
  transition: transform 0.2s ease;
}

.user-avatar:hover {
  transform: scale(1.05);
}

.user-avatar:active {
  transform: scale(0.95);
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

/* 头像占位符样式 */
.avatar-placeholder {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 12px;
  font-weight: bold;
}

.brother-avatar-placeholder {
  width: 50px;
  height: 50px;
  background: #4facfe;
  font-size: 16px;
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

.header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.search-btn {
  background: none;
  border: 1px solid #ddd;
  border-radius: 50%;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 14px;
  color: #666;
  transition: all 0.2s ease;
}

.search-btn:hover {
  background: #f8f9ff;
  border-color: #4CAF50;
}

.search-btn .icon {
  font-size: 14px;
}

/* 房间列表样式 */
.rooms-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.room-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: white;
  border-radius: 12px;
  padding: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.room-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.room-cover {
  position: relative;
  width: 80px;
  height: 80px;
  border-radius: 12px;
  overflow: hidden;
  flex-shrink: 0;
}

.room-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.room-count {
  position: absolute;
  bottom: 4px;
  left: 4px;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 2px 6px;
  border-radius: 8px;
  font-size: 10px;
  font-weight: 600;
}

.room-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.room-title {
  display: flex;
  align-items: center;
  gap: 4px;
}

.room-icon {
  font-size: 16px;
}

.room-title h4 {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.room-emoji {
  font-size: 14px;
}

.room-tags {
  display: flex;
  align-items: center;
  gap: 8px;
}

.room-tag {
  background: linear-gradient(135deg, #ff6b9d, #ff8a9b);
  color: white;
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.room-status {
  color: #666;
  font-size: 12px;
  display: flex;
  align-items: center;
  gap: 2px;
}

.room-owner {
  color: #666;
  font-size: 12px;
}

.owner-name {
  color: #333;
  font-weight: 500;
}

.room-users {
  display: flex;
  gap: -4px;
  margin-top: 4px;
}

.room-users .user-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  border: 2px solid white;
  overflow: hidden;
  margin-left: -4px;
}

.room-users .user-avatar:first-child {
  margin-left: 0;
}

.room-users .user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
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
  padding: 6px 12px;
  font-size: 12px;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.filter-btn:hover {
  background: #f8f9ff;
  border-color: #4CAF50;
}

/* 搜索容器 */
.search-container {
  position: relative;
  margin-bottom: 12px;
  animation: slideDown 0.3s ease-out;
}

.search-input {
  width: 100%;
  padding: 10px 40px 10px 12px;
  border: 1px solid #ddd;
  border-radius: 20px;
  font-size: 14px;
  background: white;
  outline: none;
  transition: border-color 0.2s ease;
}

.search-input:focus {
  border-color: #4CAF50;
}

.clear-search {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #999;
  cursor: pointer;
  font-size: 16px;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.clear-search:hover {
  color: #666;
}

/* 标签筛选 */
.tag-filter {
  margin-bottom: 12px;
  animation: slideDown 0.3s ease-out;
}

.tag-list {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.tag-item {
  background: white;
  border: 1px solid #ddd;
  border-radius: 16px;
  padding: 6px 12px;
  font-size: 12px;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tag-item:hover {
  background: #f8f9ff;
  border-color: #4CAF50;
}

.tag-item.active {
  background: #4CAF50;
  border-color: #4CAF50;
  color: white;
}

/* 加载状态 */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  color: #666;
}

.loading-spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #f3f3f3;
  border-top: 3px solid #4CAF50;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 12px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 错误状态 */
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  text-align: center;
}

.error-message {
  color: #ff4444;
  font-size: 14px;
  margin-bottom: 12px;
}

.retry-btn {
  background: #4CAF50;
  color: white;
  border: none;
  border-radius: 16px;
  padding: 8px 16px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.retry-btn:hover {
  background: #45a049;
}

/* 无数据状态 */
.empty-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  text-align: center;
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 12px;
  opacity: 0.5;
}

.empty-message {
  color: #666;
  font-size: 14px;
  margin-bottom: 12px;
}

.reset-btn {
  background: #4CAF50;
  color: white;
  border: none;
  border-radius: 16px;
  padding: 8px 16px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.reset-btn:hover {
  background: #45a049;
}

/* 动画效果 */
@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
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
  
  .avatar-placeholder {
    width: 28px;
    height: 28px;
  }
  
  .avatar-item span {
    font-size: 10px;
  }
  
  .brother-avatar-placeholder {
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
@media (prefers-contrast: more) {
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