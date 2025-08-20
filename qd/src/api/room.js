/**
 * 房间相关API接口
 * 对接后端 Go 语言 room 接口
 */

import http from './request.js'
import { API_ENDPOINTS } from './config.js'

/**
 * 工具函数：将后端房间数据转换为前端组件所需格式
 */
function transformRoomData(backendRoom) {
  if (!backendRoom) {
    console.warn('房间数据为空:', backendRoom)
    return null
  }

  console.log('转换房间数据:', backendRoom)

  const transformed = {
    id: backendRoom.id || backendRoom.room_id || 0,
    name: backendRoom.room_name || backendRoom.name || '未命名房间',
    description: backendRoom.description || backendRoom.content || backendRoom.room_name || '',
    cover: backendRoom.cover_url || backendRoom.cover || generateDefaultCover(backendRoom.room_name || backendRoom.name || '房'),
    tag: backendRoom.tag_name || backendRoom.tag || backendRoom.room_type || '热门',
    userCount: formatUserCount(backendRoom.user_count || backendRoom.online_count || backendRoom.fk_member_room || 0),
    owner: backendRoom.owner_info ? {
      nickname: backendRoom.owner_info.nickname || '房主',
      avatar: backendRoom.owner_info.avatar || ''
    } : {
      nickname: backendRoom.owner_nickname || '房主',
      avatar: ''
    },
    // 保留原始字段以便前端使用
    room_type: backendRoom.room_type,
    tag_name: backendRoom.tag_name,
    user_count: backendRoom.user_count,
    fk_member_room: backendRoom.fk_member_room,
    owner_nickname: backendRoom.owner_nickname,
    // 保留原始数据以备后用
    _raw: backendRoom
  }

  console.log('转换后的房间数据:', transformed)
  return transformed
}

/**
 * 工具函数：生成默认封面图片
 */
function generateDefaultCover(roomName = '房') {
  const colors = ['667eea', '764ba2', 'f093fb', '4facfe', 'ff6b9d', 'fa709a', 'fee140']
  const colorIndex = Math.abs(hashCode(roomName)) % colors.length
  const color = colors[colorIndex]
  const firstChar = roomName.charAt(0) || '房'
  
  // 生成本地 SVG 图片，避免外部依赖
  const svg = `
    <svg width="80" height="80" xmlns="http://www.w3.org/2000/svg">
      <rect width="80" height="80" fill="#${color}"/>
      <text x="40" y="50" font-family="Arial, sans-serif" font-size="24" font-weight="bold" 
            text-anchor="middle" dominant-baseline="middle" fill="white">
        ${firstChar}
      </text>
    </svg>
  `
  
  return `data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`
}

/**
 * 工具函数：格式化用户数量显示
 */
function formatUserCount(count) {
  if (count === 0) {
    return 'x0'
  }
  
  if (count >= 1000) {
    return `x${(count / 1000).toFixed(1)}k`
  }
  
  return `x${count}`
}

/**
 * 工具函数：字符串哈希函数，用于生成一致的颜色
 */
function hashCode(str) {
  let hash = 0
  if (str.length === 0) return hash
  
  for (let i = 0; i < str.length; i++) {
    const char = str.charCodeAt(i)
    hash = ((hash << 5) - hash) + char
    hash = hash & hash // 转换为32位整数
  }
  
  return hash
}

/**
 * 房间API服务类
 */
class RoomAPI {
  /**
   * 获取推荐房间列表
   * @param {number} page - 页码，默认1
   * @param {number} pageSize - 每页数量，默认10
   * @returns {Promise} 推荐房间数据
   */
  async getRecommendRooms(page = 1, pageSize = 10) {
    try {
      console.log('发送推荐房间请求:', { page, pageSize })
      const response = await http.post(API_ENDPOINTS.ROOM.GET_RECOMMEND_ROOMS, {
        page,
        page_size: pageSize  // 后端期望的字段名是 page_size
      })
      
      console.log('推荐房间原始响应:', response)
      
      // 转换数据格式为前端组件所需格式
      if (response && response.data) {
        const transformedRooms = response.data.rooms ? response.data.rooms.map(room => transformRoomData(room)) : []
        console.log('转换后的房间数据:', transformedRooms)
        
        return {
          ...response,
          data: {
            ...response.data,
            rooms: transformedRooms
          }
        }
      }
      
      return response
    } catch (error) {
      console.error('获取推荐房间失败:', error)
      throw new Error(error.message || '获取推荐房间失败')
    }
  }

  /**
   * 根据分类获取房间列表
   * @param {number} tagId - 标签ID
   * @param {number} page - 页码，默认1
   * @param {number} pageSize - 每页数量，默认10
   * @returns {Promise} 分类房间数据
   */
  async getRoomsByCategory(tagId, page = 1, pageSize = 10) {
    try {
      console.log('发送标签筛选请求:', { tagId, page, pageSize })
      
      const response = await http.post(API_ENDPOINTS.ROOM.GET_ROOMS_BY_CATEGORY, {
        tag_id: tagId,  // 后端期望的字段名是 tag_id
        page,
        page_size: pageSize  // 后端期望的字段名是 page_size
      })
      
      console.log('标签筛选原始响应:', response)
      
      // 转换数据格式
      if (response && response.data) {
        const transformedRooms = response.data.rooms ? response.data.rooms.map(room => transformRoomData(room)) : []
        console.log('标签筛选转换后的房间数据:', transformedRooms)
        
        return {
          ...response,
          data: {
            ...response.data,
            rooms: transformedRooms
          }
        }
      }
      
      return response
    } catch (error) {
      console.error('获取分类房间失败:', error)
      throw new Error(error.message || '获取分类房间失败')
    }
  }

  /**
   * 搜索房间
   * @param {string} keyword - 搜索关键词
   * @param {number} page - 页码，默认1
   * @param {number} pageSize - 每页数量，默认10
   * @returns {Promise} 搜索结果数据
   */
  async searchRooms(keyword, page = 1, pageSize = 10) {
    try {
      // 验证搜索关键词
      if (!keyword || keyword.trim().length === 0) {
        throw new Error('搜索关键词不能为空')
      }
      
      if (keyword.length > 50) {
        throw new Error('搜索关键词长度不能超过50个字符')
      }

      const response = await http.post(API_ENDPOINTS.ROOM.SEARCH_ROOMS, {
        keyword: keyword.trim(),
        page,
        page_size: pageSize  // 后端期望的字段名是 page_size
      })
      
      // 转换数据格式
      if (response && response.data) {
        return {
          ...response,
          data: {
            ...response.data,
            rooms: response.data.rooms ? response.data.rooms.map(room => transformRoomData(room)) : []
          }
        }
      }
      
      return response
    } catch (error) {
      console.error('搜索房间失败:', error)
      throw new Error(error.message || '搜索房间失败')
    }
  }

  /**
   * 进入房间
   * @param {number} roomId - 房间ID
   * @param {number} userId - 用户ID（可选，如果没有登录token时使用）
   * @returns {Promise} 进入房间结果
   */
  async joinRoom(roomId, userId = 1) {
    try {
      const response = await http.post(API_ENDPOINTS.ROOM.JOIN_ROOM, {
        room_id: roomId,
        user_id: userId  // 暂时使用默认用户ID 1
      })
      
      return response
    } catch (error) {
      console.error('进入房间失败:', error)
      throw new Error(error.message || '进入房间失败')
    }
  }

  /**
   * 创建房间
   * @param {Object} roomData - 房间数据
   * @param {string} roomData.room_name - 房间名称
   * @param {number} roomData.tag_id - 标签ID
   * @param {string} roomData.content - 房间描述
   * @param {string} roomData.id_card - 身份证号（后端必填）
   * @param {string} roomData.real_name - 真实姓名（后端必填）
   * @returns {Promise} 创建房间结果
   */
  async createRoom(roomData) {
    try {
      console.log('发送创建房间请求:', roomData)
      
      const response = await http.post(API_ENDPOINTS.ROOM.CREATE_ROOM, roomData)
      
      console.log('创建房间响应:', response)
      return response
    } catch (error) {
      console.error('创建房间失败:', error)
      throw new Error(error.message || '创建房间失败')
    }
  }



  /**
   * 更新房间信息
   * @param {number} roomId - 房间ID
   * @param {string} announcement - 公告内容
   * @returns {Promise} 更新结果
   */
  async updateRoomInfo(roomId, announcement) {
    try {
      console.log('发送更新房间信息请求:', { roomId, announcement })
      
      // 验证参数
      if (!roomId || roomId <= 0) {
        throw new Error('房间ID不能为空')
      }
      
      if (typeof announcement !== 'string') {
        throw new Error('公告内容必须是字符串')
      }
      
      if (announcement.length > 500) {
        throw new Error('公告内容不能超过500个字符')
      }

      const response = await http.post(API_ENDPOINTS.ROOM.UPDATE_ROOM, {
        id: roomId,
        announcement: announcement.trim()
      })
      
      console.log('更新房间信息原始响应:', response)
      
      // 确保返回完整的响应对象，不要在这里抛出错误
      // 让上层组件根据 response.code 来判断成功或失败
      return response
    } catch (error) {
      console.error('更新房间信息网络错误:', error)
      
      // 如果是网络错误或请求错误，重新抛出
      if (error.name === 'NetworkError' || error.message.includes('fetch') || !error.response) {
        throw error
      }
      
      // 如果是 HTTP 错误但有响应体，返回响应体
      if (error.response && error.response.data) {
        console.log('HTTP错误响应:', error.response.data)
        return error.response.data
      }
      
      // 其他情况抛出错误
      throw new Error(error.message || '更新房间信息失败')
    }
  }

  /**
   * 获取房间标签列表
   * @returns {Promise} 标签列表
   */
  async getRoomTags() {
    try {
      // 根据第二张图的六个标签配置
      return {
        code: 200,
        msg: '获取成功',
        data: [
          { id: 1, name: '娱乐', color: '#4facfe' },
          { id: 2, name: '才艺', color: '#f093fb' },
          { id: 3, name: '交友速配', color: '#fa709a' },
          { id: 4, name: '音乐', color: '#764ba2' },
          { id: 5, name: '聊天', color: '#ff6b9d' },
          { id: 6, name: '陪伴', color: '#667eea' }
        ]
      }
    } catch (error) {
      console.error('获取房间标签失败:', error)
      throw new Error(error.message || '获取房间标签失败')
    }
  }

  /**
 * 申请上麦
 * @param {number} roomId - 房间ID
 * @param {number} userId - 用户ID（可选，如果没有登录token时使用）
 * @returns {Promise} 申请结果
 */
async applyMic(roomId, userId = 1) {
  try {
    console.log('发送申请上麦请求:', { roomId, userId })
    
    const response = await http.post(API_ENDPOINTS.ROOM.APPLY_MIC, {
      room_id: roomId,
      user_id: userId  // 显式传递用户ID，方便测试时不依赖JWT
    })
    
    console.log('申请上麦响应:', response)
    return response
  } catch (error) {
    console.error('申请上麦失败:', error)
    throw new Error(error.message || '申请上麦失败')
  }
}

  /**
   * 下麦
   * @param {number} roomId - 房间ID
   * @param {number} userId - 用户ID（可选，如果没有登录token时使用）
   * @returns {Promise} 下麦结果
   */
  async leaveMic(roomId, userId = 1) {
    try {
      console.log('发送下麦请求:', { roomId, userId })
      
      const response = await http.post(API_ENDPOINTS.ROOM.LEAVE_MIC, {
        room_id: roomId,
        user_id: userId  // 显式传递用户ID，方便测试时不依赖JWT
      })
      
      console.log('下麦响应:', response)
      return response
    } catch (error) {
      console.error('下麦失败:', error)
      throw new Error(error.message || '下麦失败')
    }
  }

  /**
   * 离开房间
   * @param {number} roomId - 房间ID
   * @param {number} userId - 用户ID（可选，如果没有登录token时使用）
   * @returns {Promise} 离开房间结果
   */
  async leaveRoom(roomId, userId = 1) {
    try {
      console.log('发送离开房间请求:', { roomId, userId })
      
      // 使用C文件夹中的leaveRoom接口
      const response = await http.post('/room/leaveRoom', {
        room_id: roomId,
        user_id: userId
      })
      
      console.log('离开房间响应:', response)
      return response
    } catch (error) {
      console.error('离开房间失败:', error)
      throw new Error(error.message || '离开房间失败')
    }
  }

  /**
   * 关闭房间（房主功能）
   * @param {number} roomId - 房间ID
   * @param {number} userId - 用户ID（可选，如果没有登录token时使用）
   * @returns {Promise} 关闭房间结果
   */
  async closeRoom(roomId, userId = 1) {
    try {
      console.log('发送关闭房间请求:', { roomId, userId })
      
      // 使用C文件夹中的closeRoom接口
      const response = await http.post('/room/closeRoom', {
        room_id: roomId,
        user_id: userId
      })
      
      console.log('关闭房间响应:', response)
      return response
    } catch (error) {
      console.error('关闭房间失败:', error)
      throw new Error(error.message || '关闭房间失败')
    }
  }

  /**
   * 处理麦位申请（管理员功能）
   * @param {number} applicationId - 申请ID
   * @param {number} handlerId - 处理者ID
   * @param {string} action - 操作类型 ('approve' 或 'reject')
   * @param {string} reason - 处理原因
   * @returns {Promise} 处理结果
   */
  async handleMicApplication(applicationId, handlerId, action, reason = '') {
    try {
      console.log('发送处理麦位申请请求:', { applicationId, handlerId, action, reason })
      
      const response = await http.post(API_ENDPOINTS.ROOM.HANDLE_MIC_APPLICATION, {
        application_id: applicationId,
        handler_id: handlerId,
        action: action,
        reason: reason
      })
      
      console.log('处理麦位申请响应:', response)
      return response
    } catch (error) {
      console.error('处理麦位申请失败:', error)
      throw new Error(error.message || '处理麦位申请失败')
    }
  }
}

// 创建并导出房间API实例
const roomAPI = new RoomAPI()

export default roomAPI