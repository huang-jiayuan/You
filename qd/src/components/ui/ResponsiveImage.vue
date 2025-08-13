<template>
  <picture class="responsive-image">
    <source 
      v-if="webpSrc" 
      :srcset="webpSrcSet" 
      type="image/webp"
    />
    <img 
      :src="src"
      :srcset="srcSet"
      :alt="alt"
      :loading="loading"
      :class="imageClasses"
      @load="handleLoad"
      @error="handleError"
    />
  </picture>
</template>

<script>
import { ref, computed } from 'vue'

export default {
  name: 'ResponsiveImage',
  props: {
    src: {
      type: String,
      required: true
    },
    srcSet: {
      type: String,
      default: ''
    },
    webpSrc: {
      type: String,
      default: ''
    },
    webpSrcSet: {
      type: String,
      default: ''
    },
    alt: {
      type: String,
      required: true
    },
    loading: {
      type: String,
      default: 'lazy',
      validator: (value) => ['lazy', 'eager'].includes(value)
    },
    aspectRatio: {
      type: String,
      default: 'auto'
    },
    objectFit: {
      type: String,
      default: 'cover',
      validator: (value) => ['cover', 'contain', 'fill', 'none', 'scale-down'].includes(value)
    },
    rounded: {
      type: Boolean,
      default: false
    }
  },
  emits: ['load', 'error'],
  setup(props, { emit }) {
    const isLoaded = ref(false)
    const hasError = ref(false)

    const imageClasses = computed(() => ({
      'responsive-image__img': true,
      'responsive-image__img--loaded': isLoaded.value,
      'responsive-image__img--error': hasError.value,
      'responsive-image__img--rounded': props.rounded
    }))

    const handleLoad = (event) => {
      isLoaded.value = true
      emit('load', event)
    }

    const handleError = (event) => {
      hasError.value = true
      emit('error', event)
    }

    return {
      imageClasses,
      handleLoad,
      handleError
    }
  }
}
</script>

<style scoped>
.responsive-image {
  display: block;
  width: 100%;
  height: auto;
}

.responsive-image__img {
  width: 100%;
  height: auto;
  object-fit: v-bind(objectFit);
  aspect-ratio: v-bind(aspectRatio);
  transition: opacity var(--transition-normal);
  opacity: 0;
}

.responsive-image__img--loaded {
  opacity: 1;
}

.responsive-image__img--error {
  opacity: 0.5;
  background-color: var(--color-bg-tertiary);
}

.responsive-image__img--rounded {
  border-radius: var(--border-radius-lg);
}

/* 加载状态 */
.responsive-image__img:not(.responsive-image__img--loaded):not(.responsive-image__img--error) {
  background: linear-gradient(
    90deg,
    var(--color-bg-secondary) 25%,
    var(--color-bg-tertiary) 50%,
    var(--color-bg-secondary) 75%
  );
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

@keyframes loading {
  0% {
    background-position: 200% 0;
  }
  100% {
    background-position: -200% 0;
  }
}
</style>