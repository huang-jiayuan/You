<template>
  <el-dialog
    :model-value="visible"
    :title="title"
    width="500px"
    :before-close="handleClose"
    @update:model-value="$emit('update:visible', $event)"
  >
    <div class="delete-confirm-content">
      <div class="warning-icon">
        <el-icon size="48" color="#E6A23C"><WarningFilled /></el-icon>
      </div>
      
      <div class="confirm-message">
        <p class="main-message">{{ message }}</p>
        
        <div v-if="carouselData" class="carousel-info">
          <div class="info-card">
            <div class="carousel-preview">
              <el-image
                :src="carouselData.image"
                :alt="carouselData.title"
                style="width: 60px; height: 40px; border-radius: 4px;"
                fit="cover"
              />
            </div>
            <div class="carousel-details">
              <p><strong>标题：</strong>{{ carouselData.title }}</p>
              <p><strong>链接：</strong>{{ carouselData.url }}</p>
              <p><strong>排序：</strong>{{ carouselData.orderId }}</p>
              <p><strong>状态：</strong>
                <el-tag 
                  :type="carouselData.status === 'active' ? 'success' : 'danger'"
                  size="small"
                >
                  {{ carouselData.status === 'active' ? '激活' : '禁用' }}
                </el-tag>
              </p>
            </div>
          </div>
        </div>

        <div v-if="batchData && batchData.length > 0" class="batch-info">
          <div class="batch-summary">
            <p><strong>将要删除 {{ batchData.length }} 张轮播图：</strong></p>
          </div>
          <div class="batch-list">
            <div 
              v-for="(item, index) in batchData.slice(0, 5)" 
              :key="index"
              class="batch-item"
            >
              <el-image
                :src="item.image"
                :alt="item.title"
                style="width: 40px; height: 25px; border-radius: 2px;"
                fit="cover"
              />
              <span class="batch-title">{{ item.title }}</span>
            </div>
            <div v-if="batchData.length > 5" class="more-items">
              还有 {{ batchData.length - 5 }} 张轮播图...
            </div>
          </div>
        </div>
        
        <div class="warning-text">
          <el-icon><WarningFilled /></el-icon>
          <span>此操作不可恢复，请谨慎操作！</span>
        </div>
      </div>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button 
          type="danger" 
          :loading="loading"
          @click="handleConfirm"
        >
          确定删除
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { WarningFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'DeleteConfirm'
})

// Props
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: '删除确认'
  },
  message: {
    type: String,
    default: '确定要删除吗？'
  },
  carouselData: {
    type: Object,
    default: null
  },
  batchData: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
})

// Emits
const emit = defineEmits(['update:visible', 'confirm', 'cancel'])

// 确认处理
const handleConfirm = () => {
  emit('confirm')
}

// 取消处理
const handleCancel = () => {
  emit('cancel')
}

// 关闭处理
const handleClose = (done) => {
  if (props.loading) {
    return
  }
  done()
  emit('cancel')
}
</script>

<style scoped lang="scss">
.delete-confirm-content {
  text-align: center;
  padding: 20px 0;

  .warning-icon {
    margin-bottom: 20px;
  }

  .confirm-message {
    .main-message {
      font-size: 16px;
      color: #303133;
      margin-bottom: 20px;
      font-weight: 500;
    }

    .carousel-info {
      margin-bottom: 20px;

      .info-card {
        background: #f5f7fa;
        border-radius: 8px;
        padding: 16px;
        display: flex;
        align-items: flex-start;
        gap: 12px;
        text-align: left;

        .carousel-preview {
          flex-shrink: 0;
        }

        .carousel-details {
          flex: 1;

          p {
            margin: 4px 0;
            font-size: 14px;
            color: #606266;

            strong {
              color: #303133;
            }
          }
        }
      }
    }

    .batch-info {
      margin-bottom: 20px;
      text-align: left;

      .batch-summary {
        margin-bottom: 12px;

        p {
          font-size: 14px;
          color: #303133;
          margin: 0;
        }
      }

      .batch-list {
        background: #f5f7fa;
        border-radius: 8px;
        padding: 12px;
        max-height: 200px;
        overflow-y: auto;

        .batch-item {
          display: flex;
          align-items: center;
          gap: 8px;
          margin-bottom: 8px;
          font-size: 14px;
          color: #606266;

          &:last-child {
            margin-bottom: 0;
          }

          .batch-title {
            flex: 1;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }

        .more-items {
          text-align: center;
          color: #909399;
          font-size: 12px;
          margin-top: 8px;
          padding-top: 8px;
          border-top: 1px solid #e4e7ed;
        }
      }
    }

    .warning-text {
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 8px;
      color: #E6A23C;
      font-size: 14px;
      font-weight: 500;
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
    width: 90% !important;
    margin: 5vh auto;
  }

  .delete-confirm-content {
    padding: 10px 0;

    .confirm-message {
      .carousel-info .info-card {
        flex-direction: column;
        text-align: center;

        .carousel-details {
          text-align: left;
        }
      }

      .batch-info .batch-list {
        .batch-item {
          flex-direction: column;
          align-items: flex-start;
          gap: 4px;
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
</style>