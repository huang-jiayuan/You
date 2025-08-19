<template>
  <div class="room-container">
    <!-- 房间头部 -->
    <div class="room-header">
      <div class="room-info">
        <h2>{{ roomInfo.name }}</h2>
        <span class="room-id">ID: {{ roomInfo.id }}</span>
      </div>
      <div class="room-stats">
        <span class="online-count">{{ onlineCount }}人在线</span>
      </div>
    </div>

    <!-- 用户列表 -->
    <div class="user-list">
      <div 
        v-for="user in roomUsers" 
        :key="user.id"
        class="user-item"
        @click="selectUser(user)"
      >
        <img :src="user.avatar" :alt="user.name" class="user-avatar">
        <span class="user-name">{{ user.name }}</span>
        <button 
          class="gift-btn"
          @click.stop="openGiftPanel(roomInfo.id, user.id)"
        >
          🎁
        </button>
      </div>
    </div>

    <!-- 聊天区域 -->
    <div class="chat-area">
      <div class="message-list" ref="messageList">
        <div 
          v-for="message in messages" 
          :key="message.id"
          :class="['message-item', message.type]"
        >
          <div v-if="message.type === 'gift'" class="gift-message">
            <span class="gift-sender">{{ message.senderName }}</span>
            送给
            <span class="gift-receiver">{{ message.receiverName }}</span>
            <span class="gift-info">{{ message.giftName }} x{{ message.quantity }}</span>
            <span class="gift-icon">{{ message.giftIcon }}</span>
          </div>
          <div v-else class="text-message">
            <span class="sender">{{ message.senderName }}:</span>
            <span class="content">{{ message.content }}</span>
          </div>
        </div>
      </div>
      
      <div class="input-area">
        <input 
          v-model="messageInput"
          @keyup.enter="sendMessage"
          placeholder="输入消息..."
          class="message-input"
        >
        <button @click="sendMessage" class="send-btn">发送</button>
      </div>
    </div>

    <!-- 礼物面板 -->
    <RoomGiftPanel
      :visible="showGiftPanel"
      :room-id="currentRoom"
      :receiver-id="selectedReceiver"
      @close="closeGiftPanel"
      @gift-sent="onGiftSent"
    />
  </div>
</template>

<script>
import { ref, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import RoomGiftPanel from '@/components/RoomGiftPanel.vue'
import { useRoomGift } from '@/composables/useRoomGift.js'

export default {
  name: 'Room',
  components: {
    RoomGiftPanel
  },
  setup() {
    const route = useRoute()
    const messageList = ref(null)
    const messageInput = ref('')
    
    // 房间数据
    const roomInfo = ref({
      id: route.params.id || '12345',
      name: '聊天室'
    })
    
    const roomUsers = ref([
      { id: '1', name: '小明', avatar: 'https://via.placeholder.com/40x40/333/fff?text=明' },
      { id: '2', name: '小红', avatar: 'https://via.placeholder.com/40x40/f06/fff?text=红' },
      { id: '3', name: '小李', avatar: 'https://via.placeholder.com/40x40/06f/fff?text=李' }
    ])
    
    const messages = ref([
      { id: 1, type: 'text', senderName: '小明', content: '大家好！', timestamp: Date.now() }
    ])
    
    const onlineCount = ref(roomUsers.value.length)

    // 使用礼物功能
    const {
      showGiftPanel,
      selectedReceiver,
      currentRoom,
      openGiftPanel,
      closeGiftPanel
    } = useRoomGift()

    // 选择用户
    const selectUser = (user) => {
      console.log('选择用户:', user.name)
    }

    // 发送消息
    const sendMessage = () => {
      if (!messageInput.value.trim()) return
      
      const newMessage = {
        id: Date.now(),
        type: 'text',
        senderName: '我',
        content: messageInput.value,
        timestamp: Date.now()
      }
      
      messages.value.push(newMessage)
      messageInput.value = ''
      
      // 滚动到底部
      nextTick(() => {
        if (messageList.value) {
          messageList.value.scrollTop = messageList.value.scrollHeight
        }
      })
    }

    // 礼物发送成功回调
    const onGiftSent = (giftData) => {
      // 添加礼物消息到聊天记录
      const giftMessage = {
        id: Date.now(),
        type: 'gift',
        senderName: '我',
        receiverName: getUserName(giftData.receiver),
        giftName: giftData.gift.name,
        giftIcon: giftData.gift.icon,
        quantity: giftData.quantity || 1,
        timestamp: Date.now()
      }
      
      messages.value.push(giftMessage)
      
      // 滚动到底部
      nextTick(() => {
        if (messageList.value) {
          messageList.value.scrollTop = messageList.value.scrollHeight
        }
      })
    }

    // 获取用户名
    const getUserName = (userId) => {
      const user = roomUsers.value.find(u => u.id === userId)
      return user ? user.name : '未知用户'
    }

    onMounted(() => {
      // 初始化房间数据
      console.log('房间初始化完成')
    })

    return {
      roomInfo,
      roomUsers,
      messages,
      onlineCount,
      messageInput,
      messageList,
      showGiftPanel,
      selectedReceiver,
      currentRoom,
      selectUser,
      sendMessage,
      openGiftPanel,
      closeGiftPanel,
      onGiftSent
    }
  }
}
</script>

<style scoped>
.room-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #f5f5f5;
}

.room-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  background: white;
  border-bottom: 1px solid #eee;
}

.room-info h2 {
  margin: 0;
  font-size: 18px;
}

.room-id {
  font-size: 12px;
  color: #666;
}

.online-count {
  font-size: 14px;
  color: #4CAF50;
}

.user-list {
  display: flex;
  gap: 12px;
  padding: 12px 20px;
  background: white;
  border-bottom: 1px solid #eee;
  overflow-x: auto;
}

.user-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 60px;
  cursor: pointer;
  position: relative;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-bottom: 4px;
}

.user-name {
  font-size: 12px;
  color: #333;
  text-align: center;
}

.gift-btn {
  position: absolute;
  top: -5px;
  right: -5px;
  width: 20px;
  height: 20px;
  border: none;
  border-radius: 50%;
  background: #ff6b35;
  color: white;
  font-size: 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chat-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: white;
  margin: 8px 16px 16px;
  border-radius: 12px;
  overflow: hidden;
}

.message-list {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
}

.message-item {
  margin-bottom: 12px;
  padding: 8px 12px;
  border-radius: 8px;
}

.message-item.text {
  background: #f8f9ff;
}

.message-item.gift {
  background: linear-gradient(135deg, #ff6b35 0%, #f7931e 100%);
  color: white;
}

.gift-message {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
}

.gift-sender, .gift-receiver {
  font-weight: bold;
}

.gift-info {
  background: rgba(255,255,255,0.2);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.gift-icon {
  font-size: 18px;
}

.text-message .sender {
  font-weight: bold;
  color: #333;
}

.input-area {
  display: flex;
  padding: 12px 16px;
  border-top: 1px solid #eee;
  gap: 8px;
}

.message-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 20px;
  outline: none;
}

.send-btn {
  background: #4CAF50;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  cursor: pointer;
}
</style>