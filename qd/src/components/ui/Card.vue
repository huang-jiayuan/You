<template>
  <div :class="cardClasses">
    <!-- 卡片头部 -->
    <div v-if="title || $slots.header" class="card__header">
      <slot name="header">
        <h3 class="card__title">{{ title }}</h3>
      </slot>
    </div>

    <!-- 卡片内容 -->
    <div class="card__content">
      <slot />
    </div>

    <!-- 卡片操作区域 -->
    <div v-if="$slots.actions" class="card__actions">
      <slot name="actions" />
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'

export default {
  name: 'Card',
  props: {
    title: {
      type: String,
      default: ''
    },
    shadow: {
      type: Boolean,
      default: true
    }
  },
  setup(props) {
    // 计算卡片的CSS类
    const cardClasses = computed(() => {
      return [
        'card',
        {
          'card--shadow': props.shadow,
          'card--no-shadow': !props.shadow
        }
      ]
    })

    return {
      cardClasses
    }
  }
}
</script>

<style scoped>
.card {
  background-color: var(--color-bg-primary);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-lg);
  overflow: hidden;
  transition: all var(--transition-fast);
}

.card--shadow {
  box-shadow: var(--shadow-sm);
}

.card--shadow:hover {
  box-shadow: var(--shadow-md);
  transform: translateY(-2px);
}

.card--no-shadow {
  box-shadow: none;
}

.card__header {
  padding: var(--spacing-lg);
  border-bottom: 1px solid var(--color-border);
  background-color: var(--color-bg-secondary);
}

.card__title {
  margin: 0;
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-text-primary);
}

.card__content {
  padding: var(--spacing-lg);
}

.card__actions {
  padding: var(--spacing-md) var(--spacing-lg);
  border-top: 1px solid var(--color-border);
  background-color: var(--color-bg-tertiary);
  display: flex;
  gap: var(--spacing-sm);
  justify-content: flex-end;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .card__header,
  .card__content {
    padding: var(--spacing-md);
  }
  
  .card__actions {
    padding: var(--spacing-sm) var(--spacing-md);
    flex-direction: column;
  }
  
  .card__title {
    font-size: var(--font-size-base);
  }
}
</style>