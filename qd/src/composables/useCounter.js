import { ref, readonly, computed } from 'vue'

/**
 * 计数器组合式函数
 * @param {number} initialValue - 初始值，默认为0
 * @returns {object} 计数器相关的响应式数据和方法
 */
export function useCounter(initialValue = 0) {
  // 响应式数据
  const count = ref(initialValue)
  const initialCount = ref(initialValue)

  // 计算属性
  const isZero = computed(() => count.value === 0)
  const isPositive = computed(() => count.value > 0)
  const isNegative = computed(() => count.value < 0)
  const doubleCount = computed(() => count.value * 2)

  // 方法
  const increment = (step = 1) => {
    count.value += step
  }

  const decrement = (step = 1) => {
    count.value -= step
  }

  const reset = () => {
    count.value = initialCount.value
  }

  const set = (value) => {
    if (typeof value === 'number') {
      count.value = value
    }
  }

  const incrementBy = (amount) => {
    increment(amount)
  }

  const decrementBy = (amount) => {
    decrement(amount)
  }

  // 返回只读的count和操作方法
  return {
    // 响应式数据（只读）
    count: readonly(count),
    
    // 计算属性
    isZero,
    isPositive,
    isNegative,
    doubleCount,
    
    // 操作方法
    increment,
    decrement,
    reset,
    set,
    incrementBy,
    decrementBy
  }
}

/**
 * 带限制的计数器组合式函数
 * @param {number} initialValue - 初始值
 * @param {number} min - 最小值
 * @param {number} max - 最大值
 * @returns {object} 带限制的计数器相关数据和方法
 */
export function useBoundedCounter(initialValue = 0, min = -Infinity, max = Infinity) {
  const count = ref(Math.max(min, Math.min(max, initialValue)))
  const initialCount = ref(initialValue)

  // 计算属性
  const isAtMin = computed(() => count.value <= min)
  const isAtMax = computed(() => count.value >= max)
  const canIncrement = computed(() => count.value < max)
  const canDecrement = computed(() => count.value > min)

  // 方法
  const increment = (step = 1) => {
    const newValue = count.value + step
    count.value = Math.min(max, newValue)
  }

  const decrement = (step = 1) => {
    const newValue = count.value - step
    count.value = Math.max(min, newValue)
  }

  const reset = () => {
    count.value = Math.max(min, Math.min(max, initialCount.value))
  }

  const set = (value) => {
    if (typeof value === 'number') {
      count.value = Math.max(min, Math.min(max, value))
    }
  }

  return {
    count: readonly(count),
    isAtMin,
    isAtMax,
    canIncrement,
    canDecrement,
    increment,
    decrement,
    reset,
    set,
    min,
    max
  }
}