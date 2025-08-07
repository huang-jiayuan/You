<template>
  <el-dialog
    :model-value="visible"
    title="轮播图预览"
    width="800px"
    :before-close="handleClose"
    @update:model-value="$emit('update:visible', $event)"
  >
    <div class="preview-container">
      <div class="carousel-wrapper" ref="carouselWrapper">
        <div 
          class="carousel-track" 
          :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
        >
          <div 
            v-for="(item, index) in carouselList" 
            :key="index" 
            class="carousel-slide"
            @click="handleSlideClick(item)"
          >
            <el-image
              :src="item.image"
              :alt="item.title"
              fit="cover"
              class="slide-image"
            />
            <div class="carousel-overlay">
              <h3 class="slide-title">{{ item.title }}</h3>
              <p class="slide-url">{{ item.url }}</p>
            </div>
          </div>
        </div>
        
        <!-- 左右切换按钮 -->
        <button 
          v-if="carouselList.length > 1"
          class="carousel-btn prev" 
          @click="prevSlide"
          :disabled="currentIndex === 0"
        >
          <el-icon><ArrowLeft /></el-icon>
        </button>
        <button 
          v-if="carouselList.length > 1"
          class="carousel-btn next" 
          @click="nextSlide"
          :disabled="currentIndex === carouselList.length - 1"
        >
          <el-icon><ArrowRight /></el-icon>
        </button>
        
        <!-- 指示器 -->
        <div v-if="carouselList.length > 1" class="carousel-indicators">
          <span 
            v-for="(item, index) in carouselList" 
            :key="index"
            :class="['indicator', { active: index === currentIndex }]"
            @click="goToSlide(index)"
          ></span>
        </div>
      </div>

      <!-- 轮播图信息 -->
      <div class="carousel-info">
        <div class="info-item">
          <span class="info-label">当前显示:</span>
          <span class="info-value">{{ currentIndex + 1 }} / {{ carouselList.length }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">标题:</span>
          <span class="info-value">{{ currentCarousel?.title || '-' }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">链接:</span>
          <el-link 
            :href="currentCarousel?.url" 
            target="_blank" 
            type="primary"
            class="info-value"
          >
            {{ currentCarousel?.url || '-' }}
          </el-link>
        </div>
        <div class="info-item">
          <span class="info-label">状态:</span>
          <el-tag 
            :type="currentCarousel?.status === 'active' ? 'success' : 'danger'"
            size="small"
            class="info-value"
          >
            {{ currentCarousel?.status === 'active' ? '激活' : '禁用' }}
          </el-tag>
        </div>
      </div>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">关闭</el-button>
        <el-button 
          v-if="carouselList.length > 1"
          type="primary" 
          @click="startAutoPlay"
          :disabled="autoPlaying"
        >
          {{ autoPlaying ? '自动播放中...' : '开始自动播放' }}
        </el-button>
        <el-button 
          v-if="autoPlaying"
          @click="stopAutoPlay"
        >
          停止自动播放
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, computed, watch, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'

defineOptions({
  name: 'CarouselPreview'
})

// Props
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  carouselList: {
    type: Array,
    default: () => []
  }
})

// Emits
const emit = defineEmits(['update:visible'])

// 响应式数据
const currentIndex = ref(0)
const carouselWrapper = ref(null)
const autoPlaying = ref(false)
let autoPlayTimer = null

// 计算属性
const currentCarousel = computed(() => {
  return props.carouselList[currentIndex.value] || null
})

// 监听弹窗显示状态
watch(() => props.visible, (newVal) => {
  if (newVal) {
    currentIndex.value = 0
    stopAutoPlay()
  } else {
    stopAutoPlay()
  }
})

// 监听轮播图列表变化
watch(() => props.carouselList, () => {
  currentIndex.value = 0
  stopAutoPlay()
})

// 下一张
const nextSlide = () => {
  if (currentIndex.value < props.carouselList.length - 1) {
    currentIndex.value++
  } else if (autoPlaying.value) {
    // 自动播放模式下循环
    currentIndex.value = 0
  }
}

// 上一张
const prevSlide = () => {
  if (currentIndex.value > 0) {
    currentIndex.value--
  } else if (autoPlaying.value) {
    // 自动播放模式下循环
    currentIndex.value = props.carouselList.length - 1
  }
}

// 跳转到指定幻灯片
const goToSlide = (index) => {
  currentIndex.value = index
}

// 开始自动播放
const startAutoPlay = () => {
  if (props.carouselList.length <= 1) {
    ElMessage.warning('只有一张轮播图，无法自动播放')
    return
  }

  autoPlaying.value = true
  autoPlayTimer = setInterval(() => {
    if (currentIndex.value < props.carouselList.length - 1) {
      currentIndex.value++
    } else {
      currentIndex.value = 0
    }
  }, 3000)
  
  ElMessage.success('开始自动播放')
}

// 停止自动播放
const stopAutoPlay = () => {
  if (autoPlayTimer) {
    clearInterval(autoPlayTimer)
    autoPlayTimer = null
  }
  autoPlaying.value = false
}

// 点击轮播图
const handleSlideClick = (item) => {
  ElMessage.info(`点击了轮播图: ${item.title}`)
}

// 关闭处理
const handleClose = () => {
  stopAutoPlay()
  emit('update:visible', false)
}

// 组件卸载时清理定时器
onUnmounted(() => {
  stopAutoPlay()
})

// 键盘事件处理
const handleKeydown = (event) => {
  if (!props.visible) return
  
  switch (event.key) {
    case 'ArrowLeft':
      event.preventDefault()
      prevSlide()
      break
    case 'ArrowRight':
      event.preventDefault()
      nextSlide()
      break
    case 'Escape':
      event.preventDefault()
      handleClose()
      break
  }
}

// 添加键盘事件监听
watch(() => props.visible, (newVal) => {
  if (newVal) {
    document.addEventListener('keydown', handleKeydown)
  } else {
    document.removeEventListener('keydown', handleKeydown)
  }
})
</script>

<style scoped lang="scss">
.preview-container {
  .carousel-wrapper {
    position: relative;
    width: 100%;
    height: 300px;
    overflow: hidden;
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    margin-bottom: 20px;
  }

  .carousel-track {
    display: flex;
    width: 100%;
    height: 100%;
    transition: transform 0.3s ease-in-out;
  }

  .carousel-slide {
    flex: 0 0 100%;
    position: relative;
    cursor: pointer;
    overflow: hidden;
  }

  .slide-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .carousel-overlay {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    background: linear-gradient(transparent, rgba(0, 0, 0, 0.7));
    color: white;
    padding: 20px;
    
    .slide-title {
      margin: 0 0 8px 0;
      font-size: 18px;
      font-weight: 600;
      text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
    }

    .slide-url {
      margin: 0;
      font-size: 14px;
      opacity: 0.9;
      text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
    }
  }

  .carousel-btn {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background: rgba(255, 255, 255, 0.9);
    border: none;
    border-radius: 50%;
    width: 40px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.3s ease;
    z-index: 10;

    &:hover:not(:disabled) {
      background: white;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    }

    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    &.prev {
      left: 15px;
    }

    &.next {
      right: 15px;
    }
  }

  .carousel-indicators {
    position: absolute;
    bottom: 15px;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    gap: 8px;
    z-index: 10;

    .indicator {
      width: 8px;
      height: 8px;
      border-radius: 50%;
      background: rgba(255, 255, 255, 0.5);
      cursor: pointer;
      transition: all 0.3s ease;

      &.active {
        background: white;
        transform: scale(1.2);
      }

      &:hover {
        background: rgba(255, 255, 255, 0.8);
      }
    }
  }

  .carousel-info {
    background: #f8f9fa;
    border-radius: 6px;
    padding: 16px;

    .info-item {
      display: flex;
      align-items: center;
      margin-bottom: 12px;

      &:last-child {
        margin-bottom: 0;
      }

      .info-label {
        font-weight: 500;
        color: #606266;
        min-width: 80px;
        margin-right: 12px;
      }

      .info-value {
        color: #303133;
        flex: 1;
      }
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

// 响应式设计
@media (max-width: 768px) {
  :deep(.el-dialog) {
    width: 95% !important;
    margin: 5vh auto;
  }

  .preview-container {
    .carousel-wrapper {
      height: 200px;
    }

    .carousel-overlay {
      padding: 12px;

      .slide-title {
        font-size: 16px;
        margin-bottom: 4px;
      }

      .slide-url {
        font-size: 12px;
      }
    }

    .carousel-btn {
      width: 32px;
      height: 32px;

      &.prev {
        left: 8px;
      }

      &.next {
        right: 8px;
      }
    }

    .carousel-indicators {
      bottom: 8px;

      .indicator {
        width: 6px;
        height: 6px;
      }
    }

    .carousel-info {
      padding: 12px;

      .info-item {
        flex-direction: column;
        align-items: flex-start;
        margin-bottom: 8px;

        .info-label {
          min-width: auto;
          margin-right: 0;
          margin-bottom: 4px;
        }
      }
    }
  }

  .dialog-footer {
    flex-direction: column;
    gap: 8px;

    .el-button {
      width: 100%;
    }
  }
}

// 动画效果
.carousel-slide {
  transition: opacity 0.3s ease;
}

// 加载状态
:deep(.el-image) {
  .el-image__placeholder {
    background: #f5f7fa;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #c0c4cc;
    font-size: 14px;
  }
}
</style>