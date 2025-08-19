<template>
  <div class="create-room">
    <!-- 顶部标题栏 -->
    <div class="header">
      <button class="close-btn" @click="goBack">×</button>
      <h1 class="title">创建房间中</h1>
    </div>

    <!-- 用户头像 -->
    <div class="user-avatar-section">
      <div class="avatar-container">
        <img :src="userAvatar" alt="用户头像" class="user-avatar" />
      </div>
    </div>

    <!-- 房间标题输入 -->
    <div class="room-title-section">
      <label class="section-label">房间标题</label>
      <div class="input-container">
        <input 
          v-model="roomForm.roomName"
          type="text" 
          placeholder="郑潇的房间"
          class="room-title-input"
          maxlength="20"
        />
        <button v-if="roomForm.roomName" @click="clearRoomName" class="clear-btn">×</button>
      </div>
    </div>

    <!-- 房间分类选择 -->
    <div class="room-category-section">
      <label class="section-label">房间分类</label>
      <div class="category-grid">
        <button 
          v-for="category in categories" 
          :key="category.id"
          :class="['category-btn', { active: roomForm.tagId === category.id }]"
          @click="selectCategory(category.id)"
        >
          {{ category.name }}
        </button>
      </div>
    </div>

    <!-- 用户协议 -->
    <div class="agreement-section">
      <div class="agreement-text">
        <p>1.请勿发送涉嫌违法的文字、图片及语音信息</p>
        <p>2.请勿使用色情、违法或其他不适的头像或资料</p>
        <p>3.请勿发布含有广告、恶意信息、诈骗信息</p>
      </div>
      <div class="agreement-link">
        <a href="#" class="link">查看《ME用户协议》</a>
      </div>
    </div>

    <!-- 创建按钮 -->
    <div class="create-button-section">
      <button 
        :class="['create-btn', { disabled: !canCreate }]"
        :disabled="!canCreate"
        @click="createRoom"
      >
        {{ isCreating ? '创建中...' : '同意并创建房间' }}
      </button>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { roomAPI } from '@/api'

export default {
  name: 'CreateRoom',
  setup() {
    const router = useRouter()
    
    // 表单数据
    const roomForm = ref({
      roomName: '郑潇的房间',
      tagId: 1, // 默认选择第一个分类
      content: '',
      idCard: '123456789012345678', // 后端要求必填字段
      realName: '郑潇' // 后端要求必填字段
    })
    
    // 房间分类
    const categories = ref([
      { id: 1, name: '娱乐' },
      { id: 2, name: '交友速配' },
      { id: 3, name: '才艺' },
      { id: 4, name: '电台音乐' },
      { id: 5, name: '萌萌闲聊' },
      { id: 6, name: '安静陪伴' }
    ])
    
    // 状态
    const isCreating = ref(false)
    const userAvatar = ref('')
    
    // 计算属性
    const canCreate = computed(() => {
      return roomForm.value.roomName.trim().length > 0 && 
             roomForm.value.tagId > 0 && 
             !isCreating.value
    })
    
    // 生成用户头像
    const generateAvatar = (text, color = 'ff6b9d', size = 120) => {
      const svg = `
        <svg width="${size}" height="${size}" xmlns="http://www.w3.org/2000/svg">
          <rect width="${size}" height="${size}" fill="#${color}" rx="${size/10}"/>
          <text x="${size/2}" y="${size/2 + size/8}" font-family="Arial, sans-serif" 
                font-size="${size/3}" font-weight="bold" text-anchor="middle" 
                dominant-baseline="middle" fill="white">
            ${text}
          </text>
        </svg>
      `
      return `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`
    }
    
    // 方法
    const goBack = () => {
      router.push('/')
    }
    
    const clearRoomName = () => {
      roomForm.value.roomName = ''
    }
    
    const selectCategory = (categoryId) => {
      roomForm.value.tagId = categoryId
    }
    
    const createRoom = async () => {
      if (!canCreate.value) return
      
      try {
        isCreating.value = true
        console.log('🏗️ 开始创建房间:', roomForm.value)
        
        // 调用创建房间API
        const response = await roomAPI.createRoom({
          room_name: roomForm.value.roomName,
          tag_id: roomForm.value.tagId,
          content: roomForm.value.content || roomForm.value.roomName,
          id_card: roomForm.value.idCard,
          real_name: roomForm.value.realName
        })
        
        console.log('🏗️ 创建房间响应:', response)
        console.log('🔍 响应数据类型:', typeof response)
        console.log('🔍 响应数据结构:', JSON.stringify(response, null, 2))
        
        // 检查多种可能的成功响应格式
        const hasId = response && (response.id || (response.data && response.data.id))
        const isCodeSuccess = response && response.code === 200
        const isSuccess = isCodeSuccess || hasId
        
        console.log('🔍 检查结果 - hasId:', hasId, 'isCodeSuccess:', isCodeSuccess, 'isSuccess:', isSuccess)
        
        if (isSuccess) {
          console.log('✅ 房间创建成功!')
          
          // 创建成功后跳转到房间页面
          const roomId = response.id || response.data?.id || response.data?.room_id || Date.now()
          console.log('🚀 跳转到房间页面，房间ID:', roomId)
          router.push(`/room/${roomId}`)
        } else {
          console.error('❌ 房间创建失败:', response)
          // 即使判断失败，如果有ID也尝试跳转
          if (response && (response.id || (response.data && response.data.id))) {
            const roomId = response.id || response.data.id
            console.log('🔄 虽然判断失败，但有ID，仍尝试跳转:', roomId)
            router.push(`/room/${roomId}`)
          } else {
            alert('房间创建失败，请重试')
          }
        }
      } catch (error) {
        console.error('❌ 创建房间出错:', error)
        alert('创建房间出错: ' + (error.message || '未知错误'))
      } finally {
        isCreating.value = false
      }
    }
    
    onMounted(() => {
      // 生成用户头像
      userAvatar.value = generateAvatar('郑', 'ff6b9d', 120)
      console.log('📱 创建房间页面已加载')
    })
    
    return {
      roomForm,
      categories,
      isCreating,
      userAvatar,
      canCreate,
      goBack,
      clearRoomName,
      selectCategory,
      createRoom
    }
  }
}
</script>

<style scoped>
.create-room {
  min-height: 100vh;
  background: linear-gradient(180deg, #2d1b69 0%, #1a0f3a 100%);
  color: white;
  padding: 0;
  position: relative;
}

/* 顶部标题栏 */
.header {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px 20px;
  position: relative;
}

.close-btn {
  position: absolute;
  right: 20px;
  background: none;
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.title {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
}

/* 用户头像区域 */
.user-avatar-section {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}

.avatar-container {
  position: relative;
}

.user-avatar {
  width: 120px;
  height: 120px;
  border-radius: 16px;
  object-fit: cover;
  border: 3px solid rgba(255, 255, 255, 0.2);
}

/* 房间标题区域 */
.room-title-section {
  padding: 0 20px 24px;
}

.section-label {
  display: block;
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 12px;
  color: white;
}

.input-container {
  position: relative;
}

.room-title-input {
  width: 100%;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 16px;
  color: white;
  font-size: 16px;
  outline: none;
  box-sizing: border-box;
}

.room-title-input::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

.room-title-input:focus {
  border-color: rgba(255, 255, 255, 0.4);
  background: rgba(255, 255, 255, 0.15);
}

.clear-btn {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(255, 255, 255, 0.3);
  border: none;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

/* 房间分类区域 */
.room-category-section {
  padding: 0 20px 32px;
}

.category-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
}

.category-btn {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  padding: 12px 8px;
  color: white;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: center;
}

.category-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
}

.category-btn.active {
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-color: #667eea;
  transform: scale(1.02);
}

/* 用户协议区域 */
.agreement-section {
  padding: 0 20px 32px;
}

.agreement-text {
  margin-bottom: 12px;
}

.agreement-text p {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.7);
  margin: 4px 0;
  line-height: 1.4;
}

.agreement-link {
  text-align: center;
}

.link {
  color: #667eea;
  text-decoration: none;
  font-size: 14px;
}

.link:hover {
  text-decoration: underline;
}

/* 创建按钮区域 */
.create-button-section {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 20px;
  background: linear-gradient(180deg, transparent 0%, rgba(45, 27, 105, 0.9) 50%, rgba(45, 27, 105, 1) 100%);
}

.create-btn {
  width: 100%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border: none;
  border-radius: 25px;
  padding: 16px;
  color: white;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.create-btn:hover:not(.disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.4);
}

.create-btn.disabled {
  background: rgba(255, 255, 255, 0.2);
  color: rgba(255, 255, 255, 0.5);
  cursor: not-allowed;
}

.create-btn:active:not(.disabled) {
  transform: translateY(0);
}

/* 响应式设计 */
@media (max-width: 480px) {
  .category-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .user-avatar {
    width: 100px;
    height: 100px;
  }
}
</style>