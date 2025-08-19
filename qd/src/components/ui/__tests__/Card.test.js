import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Card from '../Card.vue'

describe('Card.vue', () => {
  it('渲染基本卡片', () => {
    const wrapper = mount(Card, {
      slots: {
        default: '卡片内容'
      }
    })
    
    expect(wrapper.text()).toBe('卡片内容')
    expect(wrapper.classes()).toContain('card')
    expect(wrapper.classes()).toContain('card--shadow')
  })

  it('渲染带标题的卡片', () => {
    const wrapper = mount(Card, {
      props: {
        title: '卡片标题'
      },
      slots: {
        default: '卡片内容'
      }
    })
    
    expect(wrapper.find('.card__title').text()).toBe('卡片标题')
    expect(wrapper.find('.card__header').exists()).toBe(true)
  })

  it('渲染无阴影的卡片', () => {
    const wrapper = mount(Card, {
      props: {
        shadow: false
      },
      slots: {
        default: '卡片内容'
      }
    })
    
    expect(wrapper.classes()).toContain('card--no-shadow')
    expect(wrapper.classes()).not.toContain('card--shadow')
  })

  it('渲染自定义头部插槽', () => {
    const wrapper = mount(Card, {
      slots: {
        header: '<div class="custom-header">自定义头部</div>',
        default: '卡片内容'
      }
    })
    
    expect(wrapper.find('.custom-header').text()).toBe('自定义头部')
    expect(wrapper.find('.card__header').exists()).toBe(true)
  })

  it('渲染操作区域插槽', () => {
    const wrapper = mount(Card, {
      slots: {
        default: '卡片内容',
        actions: '<button class="action-btn">操作</button>'
      }
    })
    
    expect(wrapper.find('.action-btn').text()).toBe('操作')
    expect(wrapper.find('.card__actions').exists()).toBe(true)
  })

  it('不显示空的头部和操作区域', () => {
    const wrapper = mount(Card, {
      slots: {
        default: '卡片内容'
      }
    })
    
    expect(wrapper.find('.card__header').exists()).toBe(false)
    expect(wrapper.find('.card__actions').exists()).toBe(false)
  })

  it('同时渲染标题和操作区域', () => {
    const wrapper = mount(Card, {
      props: {
        title: '卡片标题'
      },
      slots: {
        default: '卡片内容',
        actions: '<button>确定</button><button>取消</button>'
      }
    })
    
    expect(wrapper.find('.card__title').text()).toBe('卡片标题')
    expect(wrapper.find('.card__header').exists()).toBe(true)
    expect(wrapper.find('.card__actions').exists()).toBe(true)
    expect(wrapper.find('.card__content').text()).toBe('卡片内容')
  })
})