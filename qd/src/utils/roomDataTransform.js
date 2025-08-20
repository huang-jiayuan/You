/**
 * 房间数据转换工具函数
 * 处理后端数据格式转换为前端组件所需格式
 */

/**
 * 房间数据转换器类
 */
export class RoomDataTransformer {
  /**
   * 转换单个房间数据
   * @param {Object} backendRoom - 后端房间数据
   * @returns {Object} 前端格式的房间数据
   */
  static transformSingleRoom(backendRoom) {
    if (!backendRoom || typeof backendRoom !== 'object') {
      return null
    }

    return {
      id: this.extractRoomId(backendRoom),
      name: this.extractRoomName(backendRoom),
      description: this.extractDescription(backendRoom),
      cover: this.extractCoverUrl(backendRoom),
      tag: this.extractTag(backendRoom),
      userCount: this.formatUserCount(this.extractUserCount(backendRoom)),
      owner: this.extractOwnerInfo(backendRoom),
      status: this.extractRoomStatus(backendRoom),
      createdAt: this.extractCreatedTime(backendRoom),
      // 保留原始数据
      _raw: backendRoom
    }
  }

  /**
   * 转换房间列表数据
   * @param {Array} backendRooms - 后端房间列表
   * @returns {Array} 前端格式的房间列表
   */
  static transformRoomList(backendRooms) {
    if (!Array.isArray(backendRooms)) {
      return []
    }

    return backendRooms
      .map(room => this.transformSingleRoom(room))
      .filter(room => room !== null)
  }

  /**
   * 转换完整的房间响应数据
   * @param {Object} backendResponse - 后端响应数据
   * @returns {Object} 前端格式的响应数据
   */
  static transformRoomResponse(backendResponse) {
    if (!backendResponse || typeof backendResponse !== 'object') {
      return {
        rooms: [],
        total: 0,
        page: 1,
        pageSize: 10,
        hasMore: false
      }
    }

    const rooms = this.transformRoomList(backendResponse.rooms || backendResponse.data?.rooms || [])
    const total = backendResponse.total || backendResponse.data?.total || 0
    const page = backendResponse.page || backendResponse.data?.page || 1
    const pageSize = backendResponse.pageSize || backendResponse.data?.pageSize || 10

    return {
      rooms,
      total,
      page,
      pageSize,
      hasMore: (page * pageSize) < total
    }
  }

  /**
   * 提取房间ID
   * @param {Object} room - 房间数据
   * @returns {number} 房间ID
   */
  static extractRoomId(room) {
    return room.id || room.room_id || room.roomId || 0
  }

  /**
   * 提取房间名称
   * @param {Object} room - 房间数据
   * @returns {string} 房间名称
   */
  static extractRoomName(room) {
    const name = room.room_name || room.name || room.roomName || room.title
    return name || '未命名房间'
  }

  /**
   * 提取房间描述
   * @param {Object} room - 房间数据
   * @returns {string} 房间描述
   */
  static extractDescription(room) {
    return room.description || room.content || room.announcement || ''
  }

  /**
   * 提取封面图片URL
   * @param {Object} room - 房间数据
   * @returns {string} 封面图片URL
   */
  static extractCoverUrl(room) {
    const coverUrl = room.cover_url || room.cover || room.coverUrl || room.image
    
    if (coverUrl && this.isValidImageUrl(coverUrl)) {
      return coverUrl
    }
    
    // 生成默认封面
    return this.generateDefaultCover(this.extractRoomName(room))
  }

  /**
   * 提取房间标签
   * @param {Object} room - 房间数据
   * @returns {string} 房间标签
   */
  static extractTag(room) {
    return room.tag_name || room.tag || room.category || room.type || '默认'
  }

  /**
   * 提取用户数量
   * @param {Object} room - 房间数据
   * @returns {number} 用户数量
   */
  static extractUserCount(room) {
    return room.user_count || room.online_count || room.userCount || room.memberCount || 0
  }

  /**
   * 提取房主信息
   * @param {Object} room - 房间数据
   * @returns {Object|null} 房主信息
   */
  static extractOwnerInfo(room) {
    const ownerInfo = room.owner_info || room.owner || room.creator
    
    if (!ownerInfo) {
      return null
    }

    return {
      id: ownerInfo.id || ownerInfo.user_id || ownerInfo.userId,
      nickname: ownerInfo.nickname || ownerInfo.name || ownerInfo.username || '房主',
      avatar: ownerInfo.avatar || ownerInfo.avatar_url || ownerInfo.headImg || '',
      level: ownerInfo.level || 1
    }
  }

  /**
   * 提取房间状态
   * @param {Object} room - 房间数据
   * @returns {string} 房间状态
   */
  static extractRoomStatus(room) {
    return room.status || room.room_status || 'active'
  }

  /**
   * 提取创建时间
   * @param {Object} room - 房间数据
   * @returns {string} 创建时间
   */
  static extractCreatedTime(room) {
    return room.created_at || room.createTime || room.createdAt || ''
  }

  /**
   * 格式化用户数量显示
   * @param {number} count - 用户数量
   * @returns {string} 格式化后的用户数量
   */
  static formatUserCount(count) {
    const num = parseInt(count) || 0
    
    if (num === 0) {
      return 'x0'
    }
    
    if (num >= 10000) {
      return `x${(num / 10000).toFixed(1)}w`
    }
    
    if (num >= 1000) {
      return `x${(num / 1000).toFixed(1)}k`
    }
    
    return `x${num}`
  }

  /**
   * 生成默认封面图片
   * @param {string} roomName - 房间名称
   * @returns {string} 默认封面图片URL
   */
  static generateDefaultCover(roomName = '房') {
    const colors = [
      '667eea', '764ba2', 'f093fb', '4facfe', 'ff6b9d', 
      'fa709a', 'fee140', '4CAF50', 'FF9800', '9C27B0'
    ]
    
    const colorIndex = Math.abs(this.hashCode(roomName)) % colors.length
    const color = colors[colorIndex]
    const firstChar = this.getFirstDisplayChar(roomName)
    
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
   * 获取用于显示的第一个字符
   * @param {string} text - 文本
   * @returns {string} 第一个显示字符
   */
  static getFirstDisplayChar(text) {
    if (!text || text.length === 0) {
      return '房'
    }
    
    // 移除特殊字符，获取第一个有意义的字符
    const cleanText = text.replace(/[^\u4e00-\u9fa5a-zA-Z0-9]/g, '')
    return cleanText.charAt(0) || text.charAt(0) || '房'
  }

  /**
   * 字符串哈希函数
   * @param {string} str - 输入字符串
   * @returns {number} 哈希值
   */
  static hashCode(str) {
    let hash = 0
    if (!str || str.length === 0) return hash
    
    for (let i = 0; i < str.length; i++) {
      const char = str.charCodeAt(i)
      hash = ((hash << 5) - hash) + char
      hash = hash & hash // 转换为32位整数
    }
    
    return Math.abs(hash)
  }

  /**
   * 验证图片URL是否有效
   * @param {string} url - 图片URL
   * @returns {boolean} 是否有效
   */
  static isValidImageUrl(url) {
    if (!url || typeof url !== 'string') {
      return false
    }
    
    // 基本URL格式验证
    try {
      new URL(url)
      return true
    } catch {
      // 相对路径也认为是有效的
      return url.startsWith('/') || url.startsWith('./') || url.startsWith('../')
    }
  }

  /**
   * 获取房间标签的颜色
   * @param {string} tagName - 标签名称
   * @returns {string} 标签颜色
   */
  static getTagColor(tagName) {
    const tagColorMap = {
      '热门': '#ff6b9d',
      '娱乐': '#4facfe',
      '交友': '#fa709a',
      '才艺': '#f093fb',
      '电台音乐': '#764ba2',
      '游戏': '#667eea',
      '学习': '#4CAF50',
      '音乐': '#9C27B0',
      '默认': '#999999'
    }
    
    return tagColorMap[tagName] || tagColorMap['默认']
  }

  /**
   * 验证房间数据的完整性
   * @param {Object} room - 房间数据
   * @returns {Object} 验证结果
   */
  static validateRoomData(room) {
    const errors = []
    const warnings = []
    
    if (!room) {
      errors.push('房间数据为空')
      return { isValid: false, errors, warnings }
    }
    
    if (!this.extractRoomId(room)) {
      errors.push('缺少房间ID')
    }
    
    if (!this.extractRoomName(room)) {
      warnings.push('缺少房间名称，将使用默认名称')
    }
    
    if (!this.extractCoverUrl(room)) {
      warnings.push('缺少封面图片，将使用默认封面')
    }
    
    return {
      isValid: errors.length === 0,
      errors,
      warnings
    }
  }
}

// 导出便捷方法
export const transformRoomData = RoomDataTransformer.transformSingleRoom.bind(RoomDataTransformer)
export const transformRoomList = RoomDataTransformer.transformRoomList.bind(RoomDataTransformer)
export const transformRoomResponse = RoomDataTransformer.transformRoomResponse.bind(RoomDataTransformer)
export const generateDefaultCover = RoomDataTransformer.generateDefaultCover.bind(RoomDataTransformer)
export const formatUserCount = RoomDataTransformer.formatUserCount.bind(RoomDataTransformer)

export default RoomDataTransformer