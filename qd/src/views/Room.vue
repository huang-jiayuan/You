<template>
  <div class="voice-room">
    <!-- 关闭房间确认对话框 -->
    <div v-if="showCloseRoomConfirm" class="modal-overlay">
      <div class="modal-content">
        <h3>确认关闭房间</h3>
        <p>关闭房间后，所有用户将被移出房间，且房间将不再可用。</p>
        <p>此操作不可逆，确定要关闭房间吗？</p>
        <div class="modal-buttons">
          <button class="cancel-btn" @click="showCloseRoomConfirm = false">取消</button>
          <button class="confirm-btn" @click="closeRoom">确认关闭</button>
        </div>
      </div>
    </div>

    <!-- 公告模态框 -->
    <div v-if="announcementModal.visible" class="modal-overlay announcement-modal">
      <div class="modal-content announcement-content">
        <h3>房间公告</h3>
        
        <!-- 查看模式 -->
        <div v-if="!announcementModal.isEditing" class="announcement-view">
          <div class="announcement-text">
            {{ roomInfo.announcement || '暂无公告' }}
          </div>
          <div class="modal-buttons">
            <button class="cancel-btn" @click="closeAnnouncementModal">关闭</button>
            <button class="confirm-btn" @click="startEditAnnouncement">编辑</button>
          </div>
        </div>
        
        <!-- 编辑模式 -->
        <div v-if="announcementModal.isEditing" class="announcement-edit">
          <textarea 
            v-model="announcementModal.content"
            class="announcement-textarea"
            placeholder="请输入房间公告内容..."
            maxlength="500"
            rows="6"
          ></textarea>
          <div class="char-count">{{ announcementModal.content.length }}/500</div>
          <div class="modal-buttons">
            <button class="cancel-btn" @click="cancelEditAnnouncement" :disabled="announcementModal.loading">取消</button>
            <button class="confirm-btn" @click="saveAnnouncement" :disabled="announcementModal.loading">
              {{ announcementModal.loading ? '保存中...' : '保存' }}
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- 顶部房间信息 -->
    <div class="room-header">
      <button class="back-btn" @click="goBack">←</button>
      <div class="room-info">
        <div class="room-title">{{ roomInfo.name }}</div>
        <div class="room-id">🏠 ID: {{ roomInfo.id }}</div>
      </div>
      <div class="room-actions">
        <div class="room-avatar">
          <img :src="roomInfo.ownerAvatar" alt="房主" />
          <span class="user-count">{{ roomInfo.userCount }}</span>
        </div>
        <div class="menu-container">
          <button class="more-btn" @click="toggleMenu">⋮</button>
          <div v-if="showMenu" class="menu-dropdown">
            <div v-if="isRoomOwner" class="menu-item close-room" @click="showCloseRoomConfirm = true">
              <span class="menu-icon">🚫</span>
              <span class="menu-text">关闭房间</span>
            </div>
            <div class="menu-item" @click="goBack">
              <span class="menu-icon">🚪</span>
              <span class="menu-text">离开房间</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 房间状态栏 -->
    <div class="room-status">
      <div class="status-item">
        <span class="trophy">🏆</span>
        <span>荣誉</span>
        <span class="rank">排名 99+</span>
      </div>
      <div class="action-buttons">
        <button class="action-btn" @click="openAnnouncementModal">公告</button>
        <button class="action-btn">🎵</button>
      </div>
    </div>

    <!-- 麦位区域 -->
    <div class="mic-area">
      <div class="mic-grid">
        <div 
          v-for="(mic, index) in micPositions" 
          :key="index"
          :class="['mic-position', { 
            'occupied': mic.user, 
            'owner': mic.isOwner,
            'admin': mic.isAdmin,
            'current-user': mic.user?.id === currentUser.id
          }]"
          @click="handleMicClick(index)"
        >
          <div class="mic-avatar">
            <img v-if="mic.user" :src="mic.user.avatar" :alt="mic.user.nickname" />
            <div v-else class="empty-mic">👤</div>
          </div>
          <div class="mic-label">{{ mic.label }}</div>
          <div v-if="mic.isOwner" class="owner-badge">房主</div>
          <div v-if="mic.isAdmin" class="admin-badge">老板麦位</div>
        </div>
      </div>
    </div>

    <!-- 申请上麦按钮 -->
    <div class="mic-controls">
      <!-- 未申请状态 -->
      <button 
        v-if="!isRoomOwner && !micStatus.isOnMic && !micStatus.isApplying"
        class="apply-mic-btn"
        @click="applyForMic"
      >
        🎤
      </button>
      
      <!-- 申请中状态 -->
      <div v-if="!isRoomOwner && micStatus.isApplying" class="applying-status">
        <div class="applying-info">
          <div class="applying-text">
            <span v-if="micStatus.countdown > 0">
              申请中，{{ micStatus.countdown }}秒后自动上麦
            </span>
            <span v-else>正在处理申请...</span>
          </div>
          <div class="applying-progress" v-if="micStatus.countdown > 0">
            <div class="progress-bar" :style="{ width: `${((3 - micStatus.countdown) / 3) * 100}%` }"></div>
          </div>
        </div>
        <button class="cancel-btn" @click="cancelMicApplication">取消</button>
      </div>
      
      <!-- 已上麦状态 -->
      <button 
        v-if="!isRoomOwner && micStatus.isOnMic"
        class="leave-mic-btn"
        @click="leaveMic"
      >
        下麦
      </button>
      
      <!-- 房主状态 -->
      <div v-if="isRoomOwner" class="owner-status">
        <span>👑 房主</span>
      </div>
    </div>

    <!-- 聊天区域 -->
    <div class="chat-area">
      <div class="chat-header">
        <div class="system-avatar">
          <span class="me-logo">ME</span>
        </div>
        <span class="system-name">ME团队</span>
      </div>
      <div class="chat-content">
        <p>欢迎您进入ME语音房，请注意文明用语，遵守平台规则，房内严禁传播违法信息，低俗色情等违法违规内容，严禁未成年人进房间等行为，平台严打击任何形式的线下交易行为，任何用户不得以任何方式诱导打击，如遇他人要求您进行线下操作的，请注意财产安全，谨防诈骗，如遇疑似违法违规行为，请及时举报。</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import roomAPI from '../api/room.js'
import { getCurrentUserId, getCurrentUserInfo } from '../utils/userUtils.js'

export default {
  name: 'VoiceRoom',
  setup() {
    const route = useRoute()
    const router = useRouter()
    
    // 当前用户信息（从用户工具函数获取）
    const currentUserId = getCurrentUserId()
    const currentUserInfo = getCurrentUserInfo()
    
    // 如果没有用户ID，使用默认值或引导登录
    const finalUserId = currentUserId || 999 // 使用默认用户ID 999
    const finalUserInfo = currentUserInfo || { nickname: '游客用户' }
    
    const currentUser = reactive({
      id: finalUserId,
      nickname: finalUserInfo?.nickname || '游客用户',
      avatar: finalUserInfo?.avatar || generateAvatar(finalUserInfo?.nickname || '游客用户', '4CAF50')
    })
    
    // 如果用户未登录，在控制台输出提示
    if (!currentUserId) {
      console.warn('用户未登录，使用默认用户ID:', finalUserId)
    }
    
    // 房间信息
    const roomInfo = reactive({
      id: route.params.id || '201690922',
      name: `${route.params.id || '201690922'}的房间`,
      ownerId: null, // 房主ID，从后端API获取
      userCount: 0, // 初始人数为0，进入房间后会更新
      ownerAvatar: generateAvatar('房', 'ff6b9d'),
      announcement: '' // 房间公告
    })
    
    // 判断当前用户是否是房主
    const isRoomOwner = ref(false)
    
    // 公告模态框状态
    const announcementModal = reactive({
      visible: false,
      isEditing: false,
      content: '',
      loading: false
    })
    
    // 麦位申请状态
    const micStatus = reactive({
      isOnMic: false,
      isApplying: false,
      applicationId: null,
      autoApprovalTimer: null,
      countdown: 0
    })
    
    // 生成头像
    function generateAvatar(text, color = '4CAF50') {
      const svg = `
        <svg width="40" height="40" xmlns="http://www.w3.org/2000/svg">
          <rect width="40" height="40" fill="#${color}" rx="20"/>
          <text x="20" y="28" font-family="Arial, sans-serif" 
                font-size="16" font-weight="bold" text-anchor="middle" 
                fill="white">
            ${text}
          </text>
        </svg>
      `
      return `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`
    }
    
    // 初始化麦位
    const initializeMicPositions = () => {
      const positions = []
      
      for (let i = 0; i < 8; i++) {
        const position = {
          index: i,
          user: null,
          label: '点击上麦',
          isOwner: false,
          isAdmin: false
        }
        
        // 第1个麦位：房主麦位
        if (i === 0) {
          if (isRoomOwner.value) {
            // 如果当前用户是房主，显示在房主麦位
            position.user = currentUser
            position.label = '房主'
            position.isOwner = true
            micStatus.isOnMic = true
          } else {
            // 如果当前用户不是房主，显示房主信息
            position.user = { nickname: '房主', avatar: generateAvatar('房', 'ff6b9d') }
            position.label = '房主'
            position.isOwner = true
          }
        }
        
        // 第8个麦位：管理员麦位
        if (i === 7) {
          position.user = { nickname: '老板', avatar: generateAvatar('老', 'ffa500') }
          position.label = '老板麦位'
          position.isAdmin = true
        }
        
        positions.push(position)
      }
      
      return positions
    }
    
    const micPositions = ref(initializeMicPositions())
    
    // 申请上麦
    const applyForMic = async () => {
      if (micStatus.isApplying) {
        console.log('正在申请中...')
        return
      }
      
      if (micStatus.isOnMic) {
        console.log('已经在麦上了')
        return
      }
      
      if (isRoomOwner.value) {
        console.log('房主无需申请')
        return
      }
      
      // 查找空闲麦位
      const emptyMicIndex = micPositions.value.findIndex(pos => !pos.user && !pos.isOwner && !pos.isAdmin)
      
      if (emptyMicIndex === -1) {
        console.log('麦位已满')
        return
      }
      
      console.log('申请上麦...')
      micStatus.isApplying = true
      
      try {
        // 调用真实的申请上麦API
        const response = await roomAPI.applyMic(parseInt(roomInfo.id))
        
        console.log('申请上麦API响应:', response)
        
        if (response.code === 200 && response.data?.application_id) {
          // 申请成功，获取申请ID
          micStatus.applicationId = response.data.application_id
          console.log('申请上麦成功，申请ID:', micStatus.applicationId)
          
          // 开始3秒倒计时自动上麦
          startAutoApprovalCountdown(emptyMicIndex)
          
        } else {
          // 申请失败
          console.error('申请上麦失败:', response.msg || '未知错误')
          micStatus.isApplying = false
          alert(response.msg || '申请上麦失败')
        }
        
      } catch (error) {
        console.error('申请上麦网络错误:', error)
        micStatus.isApplying = false
        alert('网络错误，申请上麦失败')
      }
    }
    
    // 开始自动审批倒计时
    const startAutoApprovalCountdown = (micIndex) => {
      micStatus.countdown = 3
      
      const countdownTimer = setInterval(() => {
        micStatus.countdown--
        console.log(`自动上麦倒计时: ${micStatus.countdown}秒`)
        
        if (micStatus.countdown <= 0) {
          clearInterval(countdownTimer)
          // 3秒后自动上麦
          autoApproveMic(micIndex)
        }
      }, 1000)
      
      micStatus.autoApprovalTimer = countdownTimer
    }
    
    // 自动审批上麦
    const autoApproveMic = async (micIndex) => {
      try {
        console.log('开始自动审批上麦...')
        
        if (!micStatus.applicationId) {
          console.error('没有申请ID，无法自动审批')
          micStatus.isApplying = false
          micStatus.autoApprovalTimer = null
          micStatus.countdown = 0
          return
        }
        
        // 调用真实的审批API，自动通过申请
        const response = await roomAPI.handleMicApplication(
          micStatus.applicationId,
          currentUser.id, // 处理者ID（这里用当前用户ID，实际可能需要房主ID）
          1, // 审批动作：1=批准，2=拒绝
          '3秒自动通过' // 审批原因
        )
        
        console.log('自动审批API响应:', response)
        
        if (response.code === 200) {
          // 审批成功，用户上麦
          micPositions.value[micIndex].user = currentUser
          micPositions.value[micIndex].label = currentUser.nickname
          micStatus.isOnMic = true
          micStatus.isApplying = false
          micStatus.applicationId = null
          micStatus.autoApprovalTimer = null
          micStatus.countdown = 0
          
          console.log('自动上麦成功')
        } else {
          console.error('自动审批失败:', response.msg || '未知错误')
          micStatus.isApplying = false
          micStatus.applicationId = null
          micStatus.autoApprovalTimer = null
          micStatus.countdown = 0
        }
        
      } catch (error) {
        console.error('自动上麦失败:', error)
        micStatus.isApplying = false
        micStatus.applicationId = null
        micStatus.autoApprovalTimer = null
        micStatus.countdown = 0
      }
    }
    
    // 取消申请上麦
    const cancelMicApplication = () => {
      if (micStatus.autoApprovalTimer) {
        clearInterval(micStatus.autoApprovalTimer)
        micStatus.autoApprovalTimer = null
      }
      
      micStatus.isApplying = false
      micStatus.applicationId = null
      micStatus.countdown = 0
      
      console.log('已取消申请上麦')
    }
    
    // 下麦
    const leaveMic = () => {
      if (!micStatus.isOnMic) {
        console.log('不在麦上')
        return
      }
      
      if (isRoomOwner.value) {
        console.log('房主无法下麦')
        return
      }
      
      console.log('下麦...')
      
      // 直接在前端清空麦位，不调用后端API
      const userMicIndex = micPositions.value.findIndex(pos => pos.user?.id === currentUser.id)
      
      if (userMicIndex !== -1) {
        micPositions.value[userMicIndex].user = null
        micPositions.value[userMicIndex].label = '点击上麦'
        micStatus.isOnMic = false
        console.log('下麦成功')
      }
    }
    
    // 切换到指定麦位
    const switchToMic = (targetIndex) => {
      if (!micStatus.isOnMic) {
        console.log('用户未上麦，无法切换')
        return
      }
      
      if (isRoomOwner.value) {
        console.log('房主无法切换麦位')
        return
      }
      
      const targetMic = micPositions.value[targetIndex]
      
      // 检查目标麦位是否为空
      if (targetMic.user) {
        console.log('目标麦位已被占用')
        return
      }
      
      // 检查目标麦位是否为房主或管理员麦位
      if (targetMic.isOwner || targetMic.isAdmin) {
        console.log('不能切换到房主或管理员麦位')
        return
      }
      
      console.log(`切换到麦位 ${targetIndex + 1}...`)
      
      // 找到当前用户所在的麦位
      const currentMicIndex = micPositions.value.findIndex(pos => pos.user?.id === currentUser.id)
      
      if (currentMicIndex !== -1) {
        // 清空当前麦位
        micPositions.value[currentMicIndex].user = null
        micPositions.value[currentMicIndex].label = '点击上麦'
        
        // 移动到目标麦位
        micPositions.value[targetIndex].user = currentUser
        micPositions.value[targetIndex].label = currentUser.nickname
        
        console.log(`成功从麦位 ${currentMicIndex + 1} 切换到麦位 ${targetIndex + 1}`)
      } else {
        console.error('找不到当前用户的麦位')
      }
    }
    
    // 处理麦位点击
    const handleMicClick = (index) => {
      const mic = micPositions.value[index]
      
      // 如果麦位有人且不是当前用户，查看用户信息
      if (mic.user && mic.user.id !== currentUser.id) {
        console.log('查看用户信息:', mic.user.nickname)
        return
      }
      
      // 如果点击的是当前用户所在的麦位，不做任何操作
      if (mic.user && mic.user.id === currentUser.id) {
        console.log('点击了自己的麦位')
        return
      }
      
      // 如果是房主或管理员麦位，不允许切换
      if (mic.isOwner || mic.isAdmin) {
        console.log('不能切换到房主或管理员麦位')
        return
      }
      
      // 如果是空麦位
      if (!mic.user) {
        // 如果用户未上麦，申请上麦
        if (!isRoomOwner.value && !micStatus.isOnMic) {
          applyForMic()
        }
        // 如果用户已上麦，切换到这个空麦位
        else if (!isRoomOwner.value && micStatus.isOnMic) {
          switchToMic(index)
        }
      }
    }
    
    // 进入房间
    const joinRoom = async () => {
      try {
        console.log('正在进入房间...', roomInfo.id)
        
        // 清理之前的申请状态，确保重新进入房间时没有申请记录
        micStatus.isApplying = false
        micStatus.applicationId = null
        micStatus.countdown = 0
        if (micStatus.autoApprovalTimer) {
          clearTimeout(micStatus.autoApprovalTimer)
          micStatus.autoApprovalTimer = null
        }
        console.log('已清理之前的申请记录')
        
        const response = await roomAPI.joinRoom(parseInt(roomInfo.id), currentUser.id)
        
        console.log('进入房间API响应:', response)
        
        if (response.code === 200) {
          console.log('成功进入房间')
          
          // 根据后端返回的数据更新房间信息
          if (response.data) {
            // 更新房间信息
            if (response.data.room_info) {
              const roomData = response.data.room_info
              
              // 更新房间基本信息
              if (roomData.room_name) {
                roomInfo.name = roomData.room_name
              }
              
              // 更新房主ID
              if (roomData.user_id) {
                roomInfo.ownerId = roomData.user_id
                console.log('更新房主ID:', roomInfo.ownerId)
                
                // 重新判断当前用户是否是房主
                isRoomOwner.value = currentUser.id === roomInfo.ownerId
                console.log('当前用户是否是房主:', isRoomOwner.value, '当前用户ID:', currentUser.id, '房主ID:', roomInfo.ownerId)
              }
              
              // 更新房主头像
              if (roomData.owner_avatar) {
                roomInfo.ownerAvatar = roomData.owner_avatar
              }
              
              // 更新在线人数
              if (typeof roomData.member_count === 'number') {
                roomInfo.userCount = roomData.member_count
                console.log('更新房间人数:', roomInfo.userCount)
              }
              
              // 更新房间公告
              if (roomData.announcement !== undefined) {
                roomInfo.announcement = roomData.announcement || ''
                console.log('更新房间公告:', roomInfo.announcement)
              }
            }
            
            // 兼容旧的数据格式
            if (typeof response.data.user_count === 'number') {
              roomInfo.userCount = response.data.user_count
              console.log('更新房间人数(兼容格式):', roomInfo.userCount)
            } else if (typeof response.data.current_member_count === 'number') {
              roomInfo.userCount = response.data.current_member_count
              console.log('更新房间人数(当前人数):', roomInfo.userCount)
            } else if (!response.data.room_info) {
              // 如果后端没有返回人数，本地+1
              roomInfo.userCount += 1
              console.log('本地增加人数:', roomInfo.userCount)
            }
            
            // 如果后端返回了房间名称，更新房间名称（兼容旧格式）
            if (response.data.room_name && !response.data.room_info) {
              roomInfo.name = response.data.room_name
            }
            
            // 如果后端返回了麦位信息，可以更新麦位状态
            if (response.data.mic_positions) {
              console.log('更新麦位信息:', response.data.mic_positions)
              // 这里可以根据后端返回的麦位信息更新本地麦位状态
            }
          } else {
            // 如果没有返回数据，本地人数+1
            roomInfo.userCount += 1
            console.log('进入房间成功，本地人数+1:', roomInfo.userCount)
          }
        } else {
          console.error('进入房间失败:', response.msg || '未知错误')
          alert(response.msg || '进入房间失败')
        }
        
      } catch (error) {
        console.error('进入房间网络错误:', error)
        // 网络错误时也可以本地+1，模拟进入成功
        roomInfo.userCount += 1
        console.log('网络错误，本地模拟进入房间，人数+1:', roomInfo.userCount)
      }
    }
    
    // 离开房间
    const leaveRoom = async () => {
      try {
        console.log('正在离开房间...', roomInfo.id)
        
        // 如果有进行中的申请，先取消
        if (micStatus.isApplying) {
          cancelMicApplication()
        }
        
        // 清理所有申请相关的状态
        micStatus.isApplying = false
        micStatus.applicationId = null
        micStatus.countdown = 0
        if (micStatus.autoApprovalTimer) {
          clearTimeout(micStatus.autoApprovalTimer)
          micStatus.autoApprovalTimer = null
        }
        console.log('已清理申请记录和状态')
        
        // 调用离开房间的API
        const response = await roomAPI.leaveRoom(parseInt(roomInfo.id), currentUser.id)
        
        console.log('离开房间API响应:', response)
        
        if (response.code === 200) {
          console.log('成功离开房间')
          
          // 根据后端返回的数据更新房间信息
          if (response.data && typeof response.data.user_count === 'number') {
            roomInfo.userCount = response.data.user_count
            console.log('更新房间人数:', roomInfo.userCount)
          } else {
            // 如果后端没有返回人数，本地-1
            roomInfo.userCount = Math.max(0, roomInfo.userCount - 1)
            console.log('本地减少人数:', roomInfo.userCount)
          }
        } else {
          console.error('离开房间失败:', response.msg || '未知错误')
          // 即使API失败，也本地-1（因为用户确实要离开）
          roomInfo.userCount = Math.max(0, roomInfo.userCount - 1)
          console.log('API失败，本地减少人数:', roomInfo.userCount)
        }
        
      } catch (error) {
        console.error('离开房间网络错误:', error)
        // 网络错误时也本地-1
        roomInfo.userCount = Math.max(0, roomInfo.userCount - 1)
        console.log('网络错误，本地模拟离开房间，人数-1:', roomInfo.userCount)
      }
    }
    
    // 组件挂载时进入房间
    onMounted(() => {
      console.log('Room组件已挂载，房间ID:', roomInfo.id)
      joinRoom()
    })
    
    // 返回上一页
    const goBack = async () => {
      console.log('用户点击返回按钮')
      // 先离开房间，再返回
      await leaveRoom()
      // 返回上一页
      router.go(-1)
    }
    
    // 公告相关方法
    const openAnnouncementModal = () => {
      announcementModal.visible = true
      announcementModal.isEditing = false
      announcementModal.content = roomInfo.announcement || ''
    }
    
    const closeAnnouncementModal = () => {
      announcementModal.visible = false
      announcementModal.isEditing = false
      announcementModal.loading = false
    }
    
    const startEditAnnouncement = () => {
      // 临时允许所有用户编辑公告，方便测试
      // if (!isRoomOwner.value) {
      //   alert('只有房主可以编辑公告')
      //   return
      // }
      announcementModal.isEditing = true
      announcementModal.content = roomInfo.announcement || ''
    }
    
    const cancelEditAnnouncement = () => {
      announcementModal.isEditing = false
      announcementModal.content = roomInfo.announcement || ''
    }
    
    const saveAnnouncement = async () => {
      // 临时允许所有用户编辑公告，方便测试
      // if (!isRoomOwner.value) {
      //   alert('只有房主可以编辑公告')
      //   return
      // }
      
      if (announcementModal.content.length > 500) {
        alert('公告内容不能超过500个字符')
        return
      }
      
      announcementModal.loading = true
      
      try {
        console.log('保存公告:', announcementModal.content)
        const response = await roomAPI.updateRoomInfo(parseInt(roomInfo.id), announcementModal.content)
        
        console.log('更新公告API响应:', response)
        
        if (response.code === 200) {
          // 更新成功
          roomInfo.announcement = announcementModal.content
          announcementModal.isEditing = false
          announcementModal.loading = false
          console.log('公告更新成功')
        } else {
          // 更新失败
          console.error('更新公告失败:', response.msg || '未知错误')
          alert('更新公告失败: ' + (response.msg || '未知错误'))
          announcementModal.loading = false
        }
        
      } catch (error) {
        console.error('更新公告网络错误:', error)
        alert('更新公告失败: ' + (error.message || '网络错误'))
        announcementModal.loading = false
      }
    }
    
    // 菜单相关状态
    const showMenu = ref(false)
    const showCloseRoomConfirm = ref(false)
    
    // 切换菜单显示状态
    const toggleMenu = () => {
      showMenu.value = !showMenu.value
    }
    
    // 关闭房间
    const closeRoom = async () => {
      try {
        console.log('正在关闭房间...', roomInfo.id)
        
        // 调用关闭房间的API
        const response = await roomAPI.closeRoom(parseInt(roomInfo.id), currentUser.id)
        
        console.log('关闭房间API响应:', response)
        
        if (response.code === 200) {
          console.log('成功关闭房间')
          alert('房间已关闭')
          // 关闭确认对话框
          showCloseRoomConfirm.value = false
          // 返回上一页
          router.go(-1)
        } else {
          console.error('关闭房间失败:', response.msg || '未知错误')
          alert(response.msg || '关闭房间失败')
          // 关闭确认对话框
          showCloseRoomConfirm.value = false
        }
      } catch (error) {
        console.error('关闭房间网络错误:', error)
        alert('网络错误，关闭房间失败')
        // 关闭确认对话框
        showCloseRoomConfirm.value = false
      }
    }
    
    // 组件卸载时离开房间
    onUnmounted(() => {
      console.log('Room组件即将卸载')
      leaveRoom()
    })
    
    return {
      roomInfo,
      micPositions,
      currentUser,
      isRoomOwner,
      micStatus,
      applyForMic,
      leaveMic,
      handleMicClick,
      cancelMicApplication,
      joinRoom,
      leaveRoom,
      goBack,
      showMenu,
      showCloseRoomConfirm,
      toggleMenu,
      closeRoom,
      // 公告相关
      announcementModal,
      openAnnouncementModal,
      closeAnnouncementModal,
      startEditAnnouncement,
      cancelEditAnnouncement,
      saveAnnouncement
    }
  }
}
</script>

<style scoped>
.voice-room {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  display: flex;
  flex-direction: column;
}

/* 顶部房间信息 */
.room-header {
  display: flex;
  align-items: flex-start;
  padding: 20px 20px 10px;
  gap: 15px;
}

.back-btn {
  background: none;
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
  padding: 5px;
  border-radius: 50%;
  transition: background 0.3s ease;
  flex-shrink: 0;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.room-info {
  flex: 1;
  min-width: 0;
}

.room-title {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 5px;
}

.room-id {
  font-size: 14px;
  opacity: 0.8;
  display: flex;
  align-items: center;
  gap: 5px;
}

.room-actions {
  display: flex;
  align-items: center;
  gap: 15px;
  flex-shrink: 0;
}

.menu-container {
  position: relative;
}

.more-btn {
  background: none;
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
  padding: 5px;
  border-radius: 50%;
  transition: background 0.3s ease;
}

.more-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.menu-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  background: rgba(0, 0, 0, 0.8);
  border-radius: 10px;
  padding: 10px 0;
  min-width: 150px;
  z-index: 100;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(10px);
}

.menu-item {
  display: flex;
  align-items: center;
  padding: 10px 15px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.menu-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

.menu-icon {
  margin-right: 10px;
  font-size: 16px;
}

.menu-text {
  font-size: 14px;
}

.close-room {
  color: #ff6b6b;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(5px);
}

.modal-content {
  background: rgba(30, 30, 40, 0.9);
  border-radius: 15px;
  padding: 25px;
  width: 80%;
  max-width: 400px;
  text-align: center;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.5);
}

.modal-content h3 {
  margin-top: 0;
  color: #ff6b6b;
  font-size: 20px;
}

.modal-content p {
  margin: 15px 0;
  font-size: 14px;
  line-height: 1.5;
}

.modal-buttons {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-top: 20px;
}

.modal-buttons button {
  padding: 10px 20px;
  border-radius: 20px;
  border: none;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.cancel-btn {
  background: rgba(255, 255, 255, 0.2);
  color: white;
}

.cancel-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

.confirm-btn {
  background: #ff6b6b;
  color: white;
}

.confirm-btn:hover {
  background: #ff4f4f;
  transform: translateY(-2px);
}

.room-avatar {
  position: relative;
}

.room-avatar img {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.user-count {
  position: absolute;
  bottom: -5px;
  right: -5px;
  background: #ff4757;
  color: white;
  font-size: 12px;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 20px;
  text-align: center;
}

.more-btn {
  background: none;
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
  padding: 5px;
}

/* 房间状态栏 */
.room-status {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px 20px;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.1);
  padding: 8px 12px;
  border-radius: 20px;
  font-size: 14px;
}

.trophy {
  font-size: 16px;
}

.rank {
  opacity: 0.8;
}

.action-buttons {
  display: flex;
  gap: 10px;
}

.action-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 14px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

/* 麦位区域 */
.mic-area {
  flex: 1;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.mic-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: repeat(2, 1fr);
  gap: 20px;
  max-width: 400px;
  width: 100%;
}

.mic-position {
  position: relative;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 20px;
  padding: 15px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 2px solid transparent;
  backdrop-filter: blur(10px);
}

.mic-position:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
}

.mic-position.occupied {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
}

.mic-position.owner {
  border-color: #ffd700;
  background: rgba(255, 215, 0, 0.2);
}

.mic-position.admin {
  border-color: #ffa500;
  background: rgba(255, 165, 0, 0.2);
}

.mic-position.current-user {
  border-color: #4CAF50;
  background: rgba(76, 175, 80, 0.2);
  box-shadow: 0 0 20px rgba(76, 175, 80, 0.3);
}

.mic-avatar {
  width: 60px;
  height: 60px;
  margin: 0 auto 10px;
  border-radius: 50%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.1);
}

.mic-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.empty-mic {
  font-size: 30px;
  opacity: 0.5;
}

.mic-label {
  font-size: 12px;
  opacity: 0.9;
  margin-top: 5px;
}

.owner-badge {
  position: absolute;
  top: -8px;
  left: 50%;
  transform: translateX(-50%);
  background: #ffd700;
  color: #333;
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 10px;
  font-weight: bold;
}

.admin-badge {
  position: absolute;
  bottom: -8px;
  left: 50%;
  transform: translateX(-50%);
  background: #ffa500;
  color: white;
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 10px;
  font-weight: bold;
}

/* 申请上麦按钮 */
.mic-controls {
  padding: 20px;
  text-align: center;
}

.apply-mic-btn {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.apply-mic-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.1);
}

.apply-mic-btn.applying {
  background: rgba(255, 193, 7, 0.3);
  animation: pulse 1.5s infinite;
}

.apply-mic-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.leave-mic-btn {
  background: rgba(244, 67, 54, 0.3);
  border: none;
  color: white;
  padding: 12px 24px;
  border-radius: 25px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.leave-mic-btn:hover {
  background: rgba(244, 67, 54, 0.5);
  transform: translateY(-2px);
}

.owner-status {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px 24px;
  background: rgba(255, 215, 0, 0.3);
  border-radius: 25px;
  font-size: 14px;
  font-weight: bold;
  backdrop-filter: blur(10px);
}

@keyframes pulse {
  0% { transform: scale(1); }
  50% { transform: scale(1.05); }
  100% { transform: scale(1); }
}

/* 申请中状态样式 */
.applying-status {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
  padding: 15px;
  background: rgba(255, 193, 7, 0.2);
  border-radius: 15px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 193, 7, 0.3);
}

.applying-info {
  text-align: center;
  width: 100%;
}

.applying-text {
  font-size: 14px;
  margin-bottom: 10px;
  color: #fff;
  font-weight: 500;
}

.applying-progress {
  width: 200px;
  height: 4px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 2px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #ffc107, #ff9800);
  border-radius: 2px;
  transition: width 1s linear;
}

.cancel-btn {
  background: rgba(244, 67, 54, 0.3);
  border: none;
  color: white;
  padding: 8px 16px;
  border-radius: 20px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.cancel-btn:hover {
  background: rgba(244, 67, 54, 0.5);
  transform: translateY(-1px);
}

/* 聊天区域 */
.chat-area {
  background: rgba(0, 0, 0, 0.3);
  margin: 0 20px 20px;
  border-radius: 15px;
  padding: 15px;
  backdrop-filter: blur(10px);
}

.chat-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.system-avatar {
  width: 30px;
  height: 30px;
  background: #4CAF50;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.me-logo {
  color: white;
  font-size: 12px;
  font-weight: bold;
}

.system-name {
  font-size: 14px;
  font-weight: bold;
  color: #4CAF50;
}

.chat-content {
  font-size: 14px;
  line-height: 1.5;
  opacity: 0.9;
}

.chat-content p {
  margin: 0;
}

/* 公告模态框样式 */
.announcement-content {
  max-width: 500px;
  width: 90%;
}

.announcement-view {
  text-align: left;
}

.announcement-text {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  padding: 15px;
  margin: 15px 0;
  min-height: 80px;
  line-height: 1.5;
  font-size: 14px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.announcement-edit {
  text-align: left;
}

.announcement-textarea {
  width: 100%;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 10px;
  padding: 15px;
  color: white;
  font-size: 14px;
  line-height: 1.5;
  resize: vertical;
  margin: 15px 0 10px 0;
  box-sizing: border-box;
}

.announcement-textarea::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

.announcement-textarea:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.5);
  background: rgba(255, 255, 255, 0.15);
}

.char-count {
  text-align: right;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 15px;
}

.modal-buttons button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .mic-grid {
    gap: 15px;
    max-width: 320px;
  }
  
  .mic-position {
    padding: 10px;
  }
  
  .mic-avatar {
    width: 50px;
    height: 50px;
  }
  
  .room-header {
    padding: 15px 15px 10px;
  }
  
  .room-title {
    font-size: 18px;
  }
  
  .announcement-content {
    width: 95%;
    max-width: none;
  }
  
  .announcement-textarea {
    font-size: 16px; /* 防止iOS缩放 */
  }
}
</style>