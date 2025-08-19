/**
 * 房间礼物相关API接口
 */
import { request } from './request.js'

// 房间礼物API
export const roomGiftAPI = {
  // 获取房间礼物列表
  getRoomGifts(roomId) {
    return request({
      url: `/room/${roomId}/gifts`,
      method: 'GET'
    })
  },

  // 在房间发送礼物
  sendRoomGift(data) {
    return request({
      url: '/room/gift/send',
      method: 'POST',
      data: {
        roomId: data.roomId,
        receiverId: data.receiverId,
        giftId: data.giftId,
        quantity: data.quantity || 1
      }
    })
  },

  // 获取房间礼物记录
  getRoomGiftRecords(roomId, params = {}) {
    return request({
      url: `/room/${roomId}/gift/records`,
      method: 'GET',
      params: {
        page: params.page || 1,
        pageSize: params.pageSize || 20
      }
    })
  }
}