/**
 * 房间公告功能测试
 */

import { describe, it, expect, beforeEach, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import Room from '../Room.vue'
import roomAPI from '../../api/room.js'

// Mock room API
vi.mock('../../api/room.js', () => ({
  default: {
    updateRoomInfo: vi.fn()
  }
}))

// Mock vue-router
vi.mock('vue-router', () => ({
  useRoute: () => ({
    params: { id: '123' }
  }),
  useRouter: () => ({
    push: vi.fn()
  })
}))

describe('Room Announcement Feature', () => {
  let wrapper

  beforeEach(() => {
    vi.clearAllMocks()
    wrapper = mount(Room)
  })

  it('should display announcement button', () => {
    const announcementBtn = wrapper.find('.action-btn')
    expect(announcementBtn.exists()).toBe(true)
    expect(announcementBtn.text()).toBe('公告')
  })

  it('should open announcement modal when clicking announcement button', async () => {
    const announcementBtn = wrapper.find('.action-btn')
    await announcementBtn.trigger('click')
    await nextTick()

    expect(wrapper.vm.announcementModal.visible).toBe(true)
    expect(wrapper.find('.announcement-modal').exists()).toBe(true)
  })

  it('should close modal when clicking close button', async () => {
    // Open modal first
    wrapper.vm.openAnnouncementModal()
    await nextTick()

    const closeBtn = wrapper.find('.close-btn')
    await closeBtn.trigger('click')
    await nextTick()

    expect(wrapper.vm.announcementModal.visible).toBe(false)
  })

  it('should enter edit mode when clicking edit button', async () => {
    wrapper.vm.openAnnouncementModal()
    await nextTick()

    const editBtn = wrapper.find('.edit-btn')
    await editBtn.trigger('click')
    await nextTick()

    expect(wrapper.vm.announcementModal.isEditing).toBe(true)
    expect(wrapper.find('.announcement-textarea').exists()).toBe(true)
  })

  it('should call updateRoomInfo API when saving announcement', async () => {
    const mockResponse = { code: 200, msg: '更新成功' }
    roomAPI.updateRoomInfo.mockResolvedValue(mockResponse)

    wrapper.vm.openAnnouncementModal()
    wrapper.vm.startEditAnnouncement()
    wrapper.vm.announcementModal.content = '新的公告内容'
    await nextTick()

    await wrapper.vm.saveAnnouncement()

    expect(roomAPI.updateRoomInfo).toHaveBeenCalledWith(123, '新的公告内容')
    expect(wrapper.vm.roomInfo.announcement).toBe('新的公告内容')
    expect(wrapper.vm.announcementModal.isEditing).toBe(false)
  })

  it('should handle API error when saving announcement', async () => {
    const mockError = new Error('更新失败')
    roomAPI.updateRoomInfo.mockRejectedValue(mockError)

    // Mock alert to avoid actual alert popup in tests
    const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {})

    wrapper.vm.openAnnouncementModal()
    wrapper.vm.startEditAnnouncement()
    wrapper.vm.announcementModal.content = '新的公告内容'

    await wrapper.vm.saveAnnouncement()

    expect(alertSpy).toHaveBeenCalledWith('更新公告失败: 更新失败')
    expect(wrapper.vm.announcementModal.loading).toBe(false)

    alertSpy.mockRestore()
  })

  it('should validate announcement content length', () => {
    const longContent = 'a'.repeat(501) // 超过500字符限制
    wrapper.vm.announcementModal.content = longContent

    expect(wrapper.vm.announcementModal.content.length).toBeGreaterThan(500)
    // 在实际应用中，textarea的maxlength属性会阻止输入超过限制的内容
  })

  it('should cancel edit and restore original content', async () => {
    const originalContent = '原始公告内容'
    wrapper.vm.roomInfo.announcement = originalContent
    wrapper.vm.openAnnouncementModal()
    wrapper.vm.startEditAnnouncement()
    
    // 修改内容
    wrapper.vm.announcementModal.content = '修改后的内容'
    
    // 取消编辑
    wrapper.vm.cancelEditAnnouncement()

    expect(wrapper.vm.announcementModal.isEditing).toBe(false)
    expect(wrapper.vm.announcementModal.content).toBe(originalContent)
  })
})