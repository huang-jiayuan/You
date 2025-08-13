<template>
  <div class="counter">
    <div class="counter__display">
      <span class="counter__label" v-if="label">{{ label }}</span>
      <div class="counter__value" :class="valueClasses">
        {{ count }}
      </div>
      <div class="counter__info" v-if="showInfo">
        <span class="counter__status">{{ statusText }}</span>
        <span class="counter__double" v-if="showDouble">双倍值: {{ doubleCount }}</span>
      </div>
    </div>

    <div class="counter__controls">
      <button 
        class="counter__btn counter__btn--decrement"
        @click="handleDecrement"
        :disabled="!canDecrement"
        :title="decrementTitle"
      >
        <span class="counter__btn-icon">−</span>
        <span class="counter__btn-text">减少</span>
      </button>

      <div class="counter__step-controls" v-if="showStepControls">
        <label class="counter__step-label">
          步长:
          <select v-model="currentStep" class="counter__step-select">
            <option v-for="step in stepOptions" :key="step" :value="step">
              {{ step }}
            </option>
          </select>
        </label>
      </div>

      <button 
        class="counter__btn counter__btn--increment"
        @click="handleIncrement"
        :disabled="!canIncrement"
        :title="incrementTitle"
      >
        <span class="counter__btn-icon">+</span>
        <span class="counter__btn-text">增加</span>
      </button>
    </div>

    <div class="counter__actions" v-if="showActions">
      <button 
        class="counter__action-btn counter__action-btn--reset"
        @click="handleReset"
        :disabled="isZero"
      >
        重置
      </button>
      
      <div class="counter__quick-actions" v-if="showQuickActions">
        <button 
          v-for="action in quickActions"
          :key="action.value"
          class="counter__quick-btn"
          @click="handleSet(action.value)"
          :title="action.title"
        >
          {{ action.label }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useCounter, useBoundedCounter } from '../../composables/useCounter.js'

export default {
  name: 'Counter',
  props: {
    initialValue: {
      type: Number,
      default: 0
    },
    label: {
      type: String,
      default: ''
    },
    min: {
      type: Number,
      default: -Infinity
    },
    max: {
      type: Number,
      default: Infinity
    },
    step: {
      type: Number,
      default: 1
    },
    showInfo: {
      type: Boolean,
      default: true
    },
    showDouble: {
      type: Boolean,
      default: false
    },
    showActions: {
      type: Boolean,
      default: true
    },
    showQuickActions: {
      type: Boolean,
      default: false
    },
    showStepControls: {
      type: Boolean,
      default: false
    },
    bounded: {
      type: Boolean,
      default: false
    }
  },
  emits: ['update', 'increment', 'decrement', 'reset'],
  setup(props, { emit }) {
    // 根据是否有边界选择不同的计数器
    const counter = props.bounded 
      ? useBoundedCounter(props.initialValue, props.min, props.max)
      : useCounter(props.initialValue)

    // 当前步长
    const currentStep = ref(props.step)
    
    // 步长选项
    const stepOptions = [1, 2, 5, 10]

    // 快速操作选项
    const quickActions = [
      { value: 0, label: '0', title: '设置为0' },
      { value: 10, label: '10', title: '设置为10' },
      { value: 100, label: '100', title: '设置为100' }
    ]

    // 计算属性
    const valueClasses = computed(() => ({
      'counter__value--positive': counter.isPositive?.value || counter.count.value > 0,
      'counter__value--negative': counter.isNegative?.value || counter.count.value < 0,
      'counter__value--zero': counter.isZero?.value || counter.count.value === 0
    }))

    const statusText = computed(() => {
      const count = counter.count.value
      if (count > 0) return '正数'
      if (count < 0) return '负数'
      return '零'
    })

    const canIncrement = computed(() => {
      return props.bounded ? counter.canIncrement?.value : true
    })

    const canDecrement = computed(() => {
      return props.bounded ? counter.canDecrement?.value : true
    })

    const incrementTitle = computed(() => {
      return canIncrement.value ? `增加 ${currentStep.value}` : '已达到最大值'
    })

    const decrementTitle = computed(() => {
      return canDecrement.value ? `减少 ${currentStep.value}` : '已达到最小值'
    })

    // 方法
    const handleIncrement = () => {
      if (canIncrement.value) {
        counter.increment(currentStep.value)
        emit('increment', counter.count.value)
        emit('update', counter.count.value)
      }
    }

    const handleDecrement = () => {
      if (canDecrement.value) {
        counter.decrement(currentStep.value)
        emit('decrement', counter.count.value)
        emit('update', counter.count.value)
      }
    }

    const handleReset = () => {
      counter.reset()
      emit('reset', counter.count.value)
      emit('update', counter.count.value)
    }

    const handleSet = (value) => {
      counter.set(value)
      emit('update', counter.count.value)
    }

    return {
      // 计数器数据
      count: counter.count,
      isZero: counter.isZero,
      doubleCount: counter.doubleCount,
      
      // 本地状态
      currentStep,
      stepOptions,
      quickActions,
      
      // 计算属性
      valueClasses,
      statusText,
      canIncrement,
      canDecrement,
      incrementTitle,
      decrementTitle,
      
      // 方法
      handleIncrement,
      handleDecrement,
      handleReset,
      handleSet
    }
  }
}
</script>

<style scoped>
.counter {
  background-color: var(--color-bg-primary);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-lg);
  padding: var(--spacing-lg);
  text-align: center;
  max-width: 400px;
  margin: 0 auto;
}

.counter__display {
  margin-bottom: var(--spacing-lg);
}

.counter__label {
  display: block;
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  margin-bottom: var(--spacing-sm);
  font-weight: 500;
}

.counter__value {
  font-size: 3rem;
  font-weight: 700;
  line-height: 1;
  margin-bottom: var(--spacing-sm);
  transition: color var(--transition-fast);
}

.counter__value--positive {
  color: var(--color-success);
}

.counter__value--negative {
  color: var(--color-danger);
}

.counter__value--zero {
  color: var(--color-text-secondary);
}

.counter__info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.counter__status {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
}

.counter__double {
  font-size: var(--font-size-sm);
  color: var(--color-primary);
  font-weight: 500;
}

.counter__controls {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
}

.counter__btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-xs);
  background-color: var(--color-primary);
  color: var(--color-white);
  border: none;
  border-radius: var(--border-radius-md);
  padding: var(--spacing-md);
  cursor: pointer;
  transition: all var(--transition-fast);
  min-width: 80px;
}

.counter__btn:hover:not(:disabled) {
  background-color: var(--color-primary-dark);
  transform: translateY(-2px);
}

.counter__btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.counter__btn--decrement {
  background-color: var(--color-danger);
}

.counter__btn--decrement:hover:not(:disabled) {
  background-color: var(--color-danger-dark);
}

.counter__btn-icon {
  font-size: var(--font-size-xl);
  font-weight: 700;
}

.counter__btn-text {
  font-size: var(--font-size-xs);
}

.counter__step-controls {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.counter__step-label {
  font-size: var(--font-size-sm);
  color: var(--color-text-secondary);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-xs);
}

.counter__step-select {
  padding: var(--spacing-xs) var(--spacing-sm);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-sm);
  background-color: var(--color-bg-primary);
  color: var(--color-text-primary);
}

.counter__actions {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.counter__action-btn {
  background-color: var(--color-secondary);
  color: var(--color-white);
  border: none;
  border-radius: var(--border-radius-md);
  padding: var(--spacing-sm) var(--spacing-md);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.counter__action-btn:hover:not(:disabled) {
  opacity: 0.8;
}

.counter__action-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.counter__quick-actions {
  display: flex;
  justify-content: center;
  gap: var(--spacing-sm);
}

.counter__quick-btn {
  background-color: var(--color-bg-tertiary);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-sm);
  padding: var(--spacing-xs) var(--spacing-sm);
  cursor: pointer;
  transition: all var(--transition-fast);
  font-size: var(--font-size-sm);
}

.counter__quick-btn:hover {
  background-color: var(--color-primary);
  color: var(--color-white);
  border-color: var(--color-primary);
}

/* 响应式设计 */
@media (max-width: 480px) {
  .counter {
    padding: var(--spacing-md);
  }
  
  .counter__value {
    font-size: 2.5rem;
  }
  
  .counter__controls {
    flex-direction: column;
    gap: var(--spacing-sm);
  }
  
  .counter__btn {
    min-width: 60px;
    padding: var(--spacing-sm);
  }
}
</style>