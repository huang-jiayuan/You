<template>
  <button 
    :class="buttonClasses" 
    :disabled="disabled"
    @click="handleClick"
  >
    <slot />
  </button>
</template>

<script>
import { computed } from 'vue'

export default {
  name: 'Button',
  props: {
    type: {
      type: String,
      default: 'primary',
      validator: (value) => ['primary', 'secondary', 'danger'].includes(value)
    },
    size: {
      type: String,
      default: 'medium',
      validator: (value) => ['small', 'medium', 'large'].includes(value)
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  emits: ['click'],
  setup(props, { emit }) {
    // 计算按钮的CSS类
    const buttonClasses = computed(() => {
      return [
        'btn',
        `btn--${props.type}`,
        `btn--${props.size}`,
        {
          'btn--disabled': props.disabled
        }
      ]
    })

    // 处理点击事件
    const handleClick = (event) => {
      if (!props.disabled) {
        emit('click', event)
      }
    }

    return {
      buttonClasses,
      handleClick
    }
  }
}
</script>

<style scoped>
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
  border-radius: var(--border-radius-md);
  font-weight: 500;
  text-decoration: none;
  cursor: pointer;
  transition: all var(--transition-fast);
  font-family: inherit;
  line-height: 1;
  user-select: none;
}

.btn:focus {
  outline: 2px solid var(--color-primary);
  outline-offset: 2px;
}

/* 按钮类型样式 */
.btn--primary {
  background-color: var(--color-primary);
  color: var(--color-white);
  border: 1px solid var(--color-primary);
}

.btn--primary:hover:not(.btn--disabled) {
  background-color: var(--color-primary-dark);
  border-color: var(--color-primary-dark);
}

.btn--secondary {
  background-color: transparent;
  color: var(--color-primary);
  border: 1px solid var(--color-primary);
}

.btn--secondary:hover:not(.btn--disabled) {
  background-color: var(--color-primary);
  color: var(--color-white);
}

.btn--danger {
  background-color: var(--color-danger);
  color: var(--color-white);
  border: 1px solid var(--color-danger);
}

.btn--danger:hover:not(.btn--disabled) {
  background-color: var(--color-danger-dark);
  border-color: var(--color-danger-dark);
}

/* 按钮尺寸样式 */
.btn--small {
  padding: var(--spacing-xs) var(--spacing-sm);
  font-size: var(--font-size-sm);
  min-height: 32px;
}

.btn--medium {
  padding: var(--spacing-sm) var(--spacing-md);
  font-size: var(--font-size-base);
  min-height: 40px;
}

.btn--large {
  padding: var(--spacing-md) var(--spacing-lg);
  font-size: var(--font-size-lg);
  min-height: 48px;
}

/* 禁用状态 */
.btn--disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn--disabled:hover {
  transform: none;
}
</style>