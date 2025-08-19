<template>
  <header class="header">
    <div class="container">
      <div class="header__content">
        <h1 v-if="title" class="header__title">{{ title }}</h1>
        <p v-if="subtitle" class="header__subtitle">{{ subtitle }}</p>
        <slot />
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header',
  props: {
    title: {
      type: String,
      default: ''
    },
    subtitle: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
.header {
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-dark) 100%);
  color: var(--color-white);
  padding: var(--spacing-2xl) 0;
  text-align: center;
}

.header__content {
  max-width: 800px;
  margin: 0 auto;
}

.header__title {
  font-size: var(--font-size-3xl);
  font-weight: 700;
  margin: 0 0 var(--spacing-md) 0;
  line-height: 1.2;
}

.header__subtitle {
  font-size: var(--font-size-lg);
  font-weight: 400;
  margin: 0;
  opacity: 0.9;
  line-height: 1.5;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .header {
    padding: var(--spacing-xl) 0;
  }
  
  .header__title {
    font-size: var(--font-size-2xl);
  }
  
  .header__subtitle {
    font-size: var(--font-size-base);
  }
}
</style>