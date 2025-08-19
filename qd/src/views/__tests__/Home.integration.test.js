/**
 * Home 组件集成测试
 * 验证房间数据集成功能
 */

import { describe, it, expect, beforeEach, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import Home from '../Home.vue'
import { roomAPI } from '@/api'

// Mock API
vi.mock('@/api', () => ({
  authAPI: {
    getProfile: vi.fn()
  },
  roomAPI: {
    getRecommendRooms: vi.fn(),
    getRoomsByCategory: vi.fn(),
    searchRooms: vi.fn(),
    getRoomTags: vi.fn(),
    joinRoom: vi.fn()
  }
}))

// Mock composables
vi.mock('@/composables/mobile-chat/useAudioPlayer', () => ({
  useAudioPlayer: () => ({})
}))

vi.mock('@/composables/mobile-chat/usePerformanceOptimization', () => ({
  usePerformanceOptimization: () => ({})
}))

vi.mock('@/composables/useToast', () => ({
  useToast: () => ({
    success: vi.fn(),
    error: vi.fn()
  })
}))

describe('Home Component Integration', () => {
  let wrapper

  const mockRoomsResponse = {
    code: 200,
    msg: '房间展示成功',
    data: {
      rooms: [
        {
          id: 1,
          name: '测试房间1',
          cover: 'https://example.com/cover1.jpg',
          tag: '热门',
          userCount: 'x10'
        },
        {
          id: 2,
          name: '测试房间2',
          cover: 'https://example.com/cover2.jpg',
          tag: '娱乐',
          userCount: 'x5'
        }
      ],
      total: 2,
      page: 1,
      pageSize: 10
    }
  }

  const mockTagsResponse = {
    code: 200,
    msg: '获取成功',
    data: [
      { id: 0, name: '全部', color: '#667eea' },
      { id: 1, name: '热门', color: '#ff6b9d' },
      { id: 2, name: '娱乐', color: '#4facfe' }
    ]
  }

  beforeEach(() => {
    vi.clearAllMocks()
    
    // Mock localStorage
    Object.defineProperty(window, 'localStorage', {
      value: {
        getItem: vi.fn(() => 'mock-token'),
        setItem: vi.fn(),
        removeItem: vi.fn()
      }
    })

    // Setup default API responses
    roomAPI.getRecommendRooms.mockResolvedValue(mockRoomsResponse)
    roomAPI.getRoomTags.mockResolvedValue(mockTagsResponse)
  })

  afterEach(() => {
    if (wrapper) {
      wrapper.unmount()
    }
  })

  it('should load recommend rooms on mount', async () => {
    wrapper = mount(Home)

    // Wait for component to mount and API calls to complete
    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    expect(roomAPI.getRecommendRooms).toHaveBeenCalledWith(1, 10)
    expect(roomAPI.getRoomTags).toHaveBeenCalled()
  })

  it('should display room cards when data is loaded', async () => {
    wrapper = mount(Home)

    // Wait for data to load
    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    const roomCards = wrapper.findAll('.room-card')
    expect(roomCards.length).toBe(2)

    // Check first room card content
    const firstCard = roomCards[0]
    expect(firstCard.find('h4').text()).toBe('测试房间1')
    expect(firstCard.find('.room-tag').text()).toBe('热门')
    expect(firstCard.find('.room-count').text()).toBe('x10')
  })

  it('should show loading state while fetching data', async () => {
    // Mock a delayed response
    roomAPI.getRecommendRooms.mockImplementation(() => 
      new Promise(resolve => setTimeout(() => resolve(mockRoomsResponse), 1000))
    )

    wrapper = mount(Home)

    await nextTick()

    // Should show loading state
    expect(wrapper.find('.loading-container').exists()).toBe(true)
    expect(wrapper.find('.loading-spinner').exists()).toBe(true)
  })

  it('should handle API errors gracefully', async () => {
    const mockError = new Error('API Error')
    roomAPI.getRecommendRooms.mockRejectedValue(mockError)

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Should show error state
    expect(wrapper.find('.error-container').exists()).toBe(true)
    expect(wrapper.find('.error-message').text()).toContain('API Error')
    expect(wrapper.find('.retry-btn').exists()).toBe(true)
  })

  it('should filter rooms by tag when tag is selected', async () => {
    const mockCategoryResponse = {
      ...mockRoomsResponse,
      data: {
        ...mockRoomsResponse.data,
        rooms: [mockRoomsResponse.data.rooms[1]] // Only entertainment room
      }
    }

    roomAPI.getRoomsByCategory.mockResolvedValue(mockCategoryResponse)

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Click filter button to show tags
    await wrapper.find('.filter-btn').trigger('click')
    await nextTick()

    // Should show tag filter
    expect(wrapper.find('.tag-filter').exists()).toBe(true)

    // Click on entertainment tag (assuming it's the third tag)
    const tagButtons = wrapper.findAll('.tag-item')
    await tagButtons[2].trigger('click')

    await nextTick()

    expect(roomAPI.getRoomsByCategory).toHaveBeenCalledWith(2, 1, 10)
  })

  it('should search rooms when search input is used', async () => {
    const mockSearchResponse = {
      ...mockRoomsResponse,
      data: {
        ...mockRoomsResponse.data,
        rooms: [mockRoomsResponse.data.rooms[0]] // Only first room
      }
    }

    roomAPI.searchRooms.mockResolvedValue(mockSearchResponse)

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Click search button to show search input
    await wrapper.find('.search-btn').trigger('click')
    await nextTick()

    // Should show search input
    expect(wrapper.find('.search-input').exists()).toBe(true)

    // Type in search input
    const searchInput = wrapper.find('.search-input')
    await searchInput.setValue('测试')

    // Wait for debounce
    await new Promise(resolve => setTimeout(resolve, 600))

    expect(roomAPI.searchRooms).toHaveBeenCalledWith('测试', 1, 10)
  })

  it('should show empty state when no rooms are found', async () => {
    const emptyResponse = {
      ...mockRoomsResponse,
      data: {
        ...mockRoomsResponse.data,
        rooms: []
      }
    }

    roomAPI.getRecommendRooms.mockResolvedValue(emptyResponse)

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Should show empty state
    expect(wrapper.find('.empty-container').exists()).toBe(true)
    expect(wrapper.find('.empty-message').text()).toContain('暂无房间数据')
  })

  it('should call joinRoom API when room card is clicked', async () => {
    roomAPI.joinRoom.mockResolvedValue({ code: 200, msg: '进入房间成功' })

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Click on first room card
    const firstRoomCard = wrapper.find('.room-card')
    await firstRoomCard.trigger('click')

    expect(roomAPI.joinRoom).toHaveBeenCalledWith(1)
  })

  it('should retry loading rooms when retry button is clicked', async () => {
    // First call fails
    roomAPI.getRecommendRooms.mockRejectedValueOnce(new Error('Network Error'))
    // Second call succeeds
    roomAPI.getRecommendRooms.mockResolvedValueOnce(mockRoomsResponse)

    wrapper = mount(Home)

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Should show error state
    expect(wrapper.find('.error-container').exists()).toBe(true)

    // Click retry button
    await wrapper.find('.retry-btn').trigger('click')

    await nextTick()
    await new Promise(resolve => setTimeout(resolve, 100))

    // Should call API again
    expect(roomAPI.getRecommendRooms).toHaveBeenCalledTimes(2)
  })
})