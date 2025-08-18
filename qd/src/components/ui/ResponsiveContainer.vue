<template>
  <div :class="containerClasses">
    <slot />
  </div>
</template>

<script>
import { computed } from 'vue'

export default {
  name: 'ResponsiveContainer',
  props: {
    size: {
      type: String,
      default: 'default',
      validator: (value) => ['narrow', 'default', 'wide', 'fluid'].includes(value)
    },
    padding: {
      type: String,
      default: 'default',
      validator: (value) => ['none', 'sm', 'default', 'lg', 'xl'].includes(value)
    },
    center: {
      type: Boolean,
      default: true
    }
  },
  setup(props) {
    const containerClasses = computed(() => ({
      'responsive-container': true,
      [`responsive-container--${props.size}`]: true,
      [`responsive-container--padding-${props.padding}`]: props.padding !== 'default',
      'responsive-container--center': props.center
    }))

    return {
      containerClasses
    }
  }
}
</script>

<style scoped>
.responsive-container {
  width: 100%;
}

.responsive-container--center {
  margin: 0 auto;
}

/* 容器尺寸 */
.responsive-container--narrow {
  max-width: 800px;
}

.responsive-container--default {
  max-width: 1200px;
}

.responsive-container--wide {
  max-width: 1440px;
}

.responsive-container--fluid {
  max-width: none;
}

/* 内边距 */
.responsive-container {
  padding: 0 var(--spacing-md);
}

.responsive-container--padding-none {
  padding: 0;
}

.responsive-container--padding-sm {
  padding: 0 var(--spacing-sm);
}

.responsive-container--padding-lg {
  padding: 0 var(--spacing-lg);
}

.responsive-container--padding-xl {
  padding: 0 var(--spacing-xl);
}

/* 响应式断点 */
@media (max-width: 479px) {
  .responsive-container {
    padding: 0 var(--spacing-sm);
  }
  
  .responsive-container--padding-lg,
  .responsive-container--padding-xl {
    padding: 0 var(--spacing-md);
  }
}

@media (min-width: 768px) {
  .responsive-container {
    padding: 0 var(--spacing-lg);
  }
}

@media (min-width: 1024px) {
  .responsive-container {
    padding: 0 var(--spacing-xl);
  }
}
</style>