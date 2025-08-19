/**
 * 房间礼物功能组合式函数
 */
import { ref } from 'vue'
import { roomGiftAPI } from '@/api/roomGift.js'

export function useRoomGift() {
  const showGiftPanel = ref(false)
  const selectedReceiver = ref(null)
  const currentRoom = ref(null)

  // 打开礼物面板
  const openGiftPanel = (roomId, receiverId) => {
    currentRoom.value = roomId
    selectedReceiver.value = receiverId
    showGiftPanel.value = true
  }

  // 关闭礼物面板
  const closeGiftPanel = () => {
    showGiftPanel.value = false
    selectedReceiver.value = null
  }

  // 发送礼物
  const sendGift = async (giftData) => {
    try {
      const response = await roomGiftAPI.sendRoomGift({
        roomId: currentRoom.value,
        receiverId: selectedReceiver.value,
        giftId: giftData.giftId,
        quantity: giftData.quantity || 1
      })

      if (response.success) {
        // 触发礼物动画
        showGiftAnimation(giftData)
        return { success: true, data: response.data }
      }
      
      return { success: false, error: response.message }
    } catch (error) {
      console.error('发送礼物失败:', error)
      return { success: false, error: error.message }
    }
  }

  // 显示礼物动画
  const showGiftAnimation = (giftData) => {
    // 创建礼物动画元素
    const giftElement = document.createElement('div')
    giftElement.className = 'gift-animation'
    giftElement.innerHTML = `
      <div class="gift-fly">
        <span class="gift-icon">${giftData.icon}</span>
        <span class="gift-text">${giftData.name} x${giftData.quantity}</span>
      </div>
    `
    
    // 添加样式
    giftElement.style.cssText = `
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      z-index: 9999;
      pointer-events: none;
      animation: giftFly 3s ease-out forwards;
    `

    document.body.appendChild(giftElement)

    // 3秒后移除元素
    setTimeout(() => {
      document.body.removeChild(giftElement)
    }, 3000)
  }

  // 获取房间礼物记录
  const getRoomGiftRecords = async (roomId, params = {}) => {
    try {
      const response = await roomGiftAPI.getRoomGiftRecords(roomId, params)
      return response.success ? response.data : []
    } catch (error) {
      console.error('获取礼物记录失败:', error)
      return []
    }
  }

  return {
    showGiftPanel,
    selectedReceiver,
    currentRoom,
    openGiftPanel,
    closeGiftPanel,
    sendGift,
    getRoomGiftRecords
  }
}