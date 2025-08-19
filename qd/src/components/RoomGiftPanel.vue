<template>
  <div v-if="visible" class="gift-overlay" @click="$emit('close')">
    <div class="gift-panel" @click.stop>
      <!-- 礼物列表 -->
      <div class="gift-grid">
        <div 
          v-for="gift in gifts" 
          :key="gift.id"
          class="gift-item"
          @click="sendGift(gift)"
        >
          <div class="gift-icon">{{ gift.icon }}</div>
          <div class="gift-name">{{ gift.name }}</div>
          <div class="gift-price">{{ gift.price }}币</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { roomGiftAPI } from '@/api/roomGift.js'

export default {
  name: 'RoomGiftPanel',
  props: {
    visible: Boolean,
    roomId: String,
    receiverId: String
  },
  emits: ['close', 'gift-sent'],
  setup(props, { emit }) {
    const gifts = ref([
      { id: 1, name: '玫瑰', icon: '🌹', price: 10 },
      { id: 2, name: '啤酒', icon: '🍺', price: 20 },
      { id: 3, name: '跑车', icon: '🏎️', price: 500 },
      { id: 4, name: '火箭', icon: '🚀', price: 1000 }
    ])

    const sendGift = async (gift) => {
      try {
        const response = await roomGiftAPI.sendRoomGift({
          roomId: props.roomId,
          receiverId: props.receiverId,
          giftId: gift.id,
          quantity: 1
        })

        if (response.success) {
          emit('gift-sent', { gift, receiver: props.receiverId })
          emit('close')
        }
      } catch (error) {
        console.error('发送礼物失败:', error)
      }
    }

    return {
      gifts,
      sendGift
    }
  }
}
</script>

<style scoped>
.gift-overlay {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  top: 0;
  background: rgba(0,0,0,0.5);
  z-index: 1000;
  display: flex;
  align-items: flex-end;
}

.gift-panel {
  background: white;
  border-radius: 16px 16px 0 0;
  width: 100%;
  padding: 20px;
}

.gift-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.gift-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px;
  border-radius: 12px;
  cursor: pointer;
  transition: background 0.2s;
}

.gift-item:hover {
  background: #f5f5f5;
}

.gift-icon {
  font-size: 32px;
  margin-bottom: 8px;
}

.gift-name {
  font-size: 12px;
  margin-bottom: 4px;
}

.gift-price {
  font-size: 12px;
  color: #ff6b35;
  font-weight: bold;
}
</style>