import { describe, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import Button from '../Button.vue'

describe('Button.vue', () => {
  it('渲染默认按钮', () => {
    const wrapper = mount(Button, {
      slots: {
        default: '点击我'
      }
    })
    
    expect(wrapper.text()).toBe('点击我')
    expect(wrapper.classes()).toContain('btn')
    expect(wrapper.classes()).toContain('btn--primary')
    expect(wrapper.classes()).toContain('btn--medium')
  })

  it('渲染不同类型的按钮', () => {
    const wrapper = mount(Button, {
      props: {
        type: 'secondary'
      },
      slots: {
        default: '次要按钮'
      }
    })
    
    expect(wrapper.classes()).toContain('btn--secondary')
  })

  it('渲染不同尺寸的按钮', () => {
    const wrapper = mount(Button, {
      props: {
        size: 'large'
      },
      slots: {
        default: '大按钮'
      }
    })
    
    expect(wrapper.classes()).toContain('btn--large')
  })

  it('处理禁用状态', () => {
    const wrapper = mount(Button, {
      props: {
        disabled: true
      },
      slots: {
        default: '禁用按钮'
      }
    })
    
    expect(wrapper.classes()).toContain('btn--disabled')
    expect(wrapper.attributes('disabled')).toBeDefined()
  })

  it('触发点击事件', async () => {
    const wrapper = mount(Button, {
      slots: {
        default: '点击我'
      }
    })
    
    await wrapper.trigger('click')
    expect(wrapper.emitted('click')).toBeTruthy()
    expect(wrapper.emitted('click')).toHaveLength(1)
  })

  it('禁用状态下不触发点击事件', async () => {
    const wrapper = mount(Button, {
      props: {
        disabled: true
      },
      slots: {
        default: '禁用按钮'
      }
    })
    
    await wrapper.trigger('click')
    expect(wrapper.emitted('click')).toBeFalsy()
  })

  it('验证props类型', () => {
    const wrapper = mount(Button, {
      props: {
        type: 'danger',
        size: 'small',
        disabled: false
      },
      slots: {
        default: '危险按钮'
      }
    })
    
    expect(wrapper.classes()).toContain('btn--danger')
    expect(wrapper.classes()).toContain('btn--small')
    expect(wrapper.classes()).not.toContain('btn--disabled')
  })
})