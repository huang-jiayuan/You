/**
 * 房间 API 测试文件
 * 用于验证房间相关接口的功能
 */

import { describe, it, expect, beforeEach, vi } from 'vitest'
import roomAPI from '../room.js'
import http from '../request.js'

// Mock http 请求
vi.mock('../request.js', () => ({
  default: {
    post: vi.fn()
  }
}))

describe('Room API', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  describe('getRecommendRooms', () => {
    it('should fetch recommend rooms successfully', async () => {
      const mockResponse = {
        code: 200,
        msg: '房间展示成功',
        data: {
          rooms: [
            {
              id: 1,
              room_name: '测试房间',
              cover_url: 'https://example.com/cover.jpg',
              tag_name: '热门',
              user_count: 10
            }
          ],
          total: 1,
          page: 1,
          pageSize: 10
        }
      }

      http.post.mockResolvedValue(mockResponse)

      const result = await roomAPI.getRecommendRooms(1, 10)

      expect(http.post).toHaveBeenCalledWith('/room/getRecommendRooms', {
        page: 1,
        pageSize: 10
      })

      expect(result.data.rooms).toHaveLength(1)
      expect(result.data.rooms[0]).toMatchObject({
        id: 1,
        name: '测试房间',
        tag: '热门',
        userCount: 'x10'
      })
    })

    it('should handle API error', async () => {
      const mockError = new Error('网络错误')
      http.post.mockRejectedValue(mockError)

      await expect(roomAPI.getRecommendRooms()).rejects.toThrow('网络错误')
    })
  })

  describe('getRoomsByCategory', () => {
    it('should fetch rooms by category successfully', async () => {
      const mockResponse = {
        code: 200,
        msg: '房间展示成功',
        data: {
          rooms: [
            {
              id: 2,
              room_name: '娱乐房间',
              tag_name: '娱乐',
              user_count: 5
            }
          ]
        }
      }

      http.post.mockResolvedValue(mockResponse)

      const result = await roomAPI.getRoomsByCategory(2, 1, 10)

      expect(http.post).toHaveBeenCalledWith('/room/getRoomsByCategory', {
        tagId: 2,
        page: 1,
        pageSize: 10
      })

      expect(result.data.rooms[0].tag).toBe('娱乐')
    })
  })

  describe('searchRooms', () => {
    it('should search rooms successfully', async () => {
      const mockResponse = {
        code: 200,
        msg: '房间搜索成功',
        data: {
          rooms: [
            {
              id: 3,
              room_name: '搜索结果房间',
              user_count: 8
            }
          ]
        }
      }

      http.post.mockResolvedValue(mockResponse)

      const result = await roomAPI.searchRooms('测试', 1, 10)

      expect(http.post).toHaveBeenCalledWith('/room/searchRooms', {
        keyword: '测试',
        page: 1,
        pageSize: 10
      })

      expect(result.data.rooms[0].name).toBe('搜索结果房间')
    })

    it('should validate search keyword', async () => {
      await expect(roomAPI.searchRooms('')).rejects.toThrow('搜索关键词不能为空')
      await expect(roomAPI.searchRooms('a'.repeat(51))).rejects.toThrow('搜索关键词长度不能超过50个字符')
    })
  })

  describe('transformRoomData', () => {
    it('should transform backend room data correctly', () => {
      const backendRoom = {
        id: 1,
        room_name: '测试房间',
        cover_url: 'https://example.com/cover.jpg',
        tag_name: '热门',
        user_count: 100,
        owner_info: {
          nickname: '房主',
          avatar: 'https://example.com/avatar.jpg'
        }
      }

      const transformed = roomAPI.transformRoomData(backendRoom)

      expect(transformed).toMatchObject({
        id: 1,
        name: '测试房间',
        cover: 'https://example.com/cover.jpg',
        tag: '热门',
        userCount: 'x100',
        owner: {
          nickname: '房主',
          avatar: 'https://example.com/avatar.jpg'
        }
      })
    })

    it('should handle missing data with defaults', () => {
      const backendRoom = {
        id: 1
      }

      const transformed = roomAPI.transformRoomData(backendRoom)

      expect(transformed.name).toBe('未命名房间')
      expect(transformed.tag).toBe('默认')
      expect(transformed.userCount).toBe('x0')
      expect(transformed.cover).toContain('placeholder')
    })
  })

  describe('generateDefaultCover', () => {
    it('should generate consistent default cover for same room name', () => {
      const cover1 = roomAPI.generateDefaultCover('测试房间')
      const cover2 = roomAPI.generateDefaultCover('测试房间')

      expect(cover1).toBe(cover2)
      expect(cover1).toContain('placeholder')
      expect(cover1).toContain('测')
    })
  })

  describe('formatUserCount', () => {
    it('should format user count correctly', () => {
      expect(roomAPI.formatUserCount(0)).toBe('x0')
      expect(roomAPI.formatUserCount(50)).toBe('x50')
      expect(roomAPI.formatUserCount(1500)).toBe('x1.5k')
      expect(roomAPI.formatUserCount(15000)).toBe('x15.0k')
    })
  })
})