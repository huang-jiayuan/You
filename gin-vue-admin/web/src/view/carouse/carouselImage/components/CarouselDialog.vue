<template>
  <el-dialog
    :model-value="visible"
    :title="dialogTitle"
    width="600px"
    :before-close="handleClose"
    destroy-on-close
    @update:model-value="$emit('update:visible', $event)"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="100px"
      label-position="right"
    >
      <el-form-item label="轮播图片" prop="image" required>
        <div class="upload-section">
          <el-upload
            ref="uploadRef"
            :action="uploadAction"
            :headers="uploadHeaders"
            :show-file-list="false"
            :before-upload="beforeUpload"
            :on-success="handleUploadSuccess"
            :on-error="handleUploadError"
            :on-progress="handleUploadProgress"
            accept="image/*"
            drag
          >
            <div v-if="!formData.image" class="upload-placeholder">
              <div v-if="uploading" class="upload-progress">
                <el-progress 
                  :percentage="uploadProgress" 
                  :stroke-width="6"
                  status="success"
                />
                <div class="progress-text">上传中... {{ uploadProgress }}%</div>
              </div>
              <div v-else class="upload-content">
                <el-icon class="upload-icon"><Plus /></el-icon>
                <div class="upload-text">点击或拖拽上传图片</div>
                <div class="upload-tip">支持 jpg、png、gif、webp 格式，大小不超过 2MB</div>
              </div>
            </div>
            <div v-else class="image-preview">
              <el-image
                :src="formData.image"
                fit="cover"
                style="width: 100%; height: 200px; border-radius: 6px;"
              />
              <div class="image-overlay">
                <el-button type="primary" size="small" @click.stop="replaceImage">
                  更换图片
                </el-button>
              </div>
            </div>
          </el-upload>
        </div>
      </el-form-item>

      <el-form-item label="标题" prop="title" required>
        <el-input
          v-model="formData.title"
          placeholder="请输入轮播图标题"
          maxlength="100"
          show-word-limit
          clearable
        />
      </el-form-item>

      <el-form-item label="跳转链接" prop="url" required>
        <el-input
          v-model="formData.url"
          placeholder="请输入跳转链接，如：/home 或 https://example.com"
          clearable
        >
          <template #prepend>
            <el-select v-model="urlType" style="width: 80px">
              <el-option label="内部" value="internal" />
              <el-option label="外部" value="external" />
            </el-select>
          </template>
        </el-input>
      </el-form-item>

      <el-form-item label="排序序号" prop="orderId" required>
        <el-input-number
          v-model="formData.orderId"
          :min="0"
          :max="9999"
          :step="1"
          placeholder="数字越小越靠前"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="状态" prop="status" required>
        <el-radio-group v-model="formData.status">
          <el-radio value="active">激活</el-radio>
          <el-radio value="inactive">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button 
          type="primary" 
          :loading="saveLoading" 
          @click="handleSave"
        >
          {{ mode === 'add' ? '添加' : '保存' }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed, watch, nextTick, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { 
  createCarouselImage, 
  updateCarouselImage 
} from '@/api/carouse/carouselImage'
// import { getDictFunc } from '@/utils/format' // 暂时注释掉，如果需要可以取消注释

defineOptions({
  name: 'CarouselDialog'
})

// Props
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  mode: {
    type: String,
    default: 'add', // 'add' | 'edit'
    validator: (value) => ['add', 'edit'].includes(value)
  },
  carouselData: {
    type: Object,
    default: null
  },
  maxOrderId: {
    type: Number,
    default: 0
  }
})

// Emits
const emit = defineEmits(['update:visible', 'save', 'cancel'])

// 响应式数据
const formRef = ref()
const uploadRef = ref()
const saveLoading = ref(false)
const urlType = ref('internal')
const uploadProgress = ref(0)
const uploading = ref(false)

// 表单数据
const formData = reactive({
  image: '',
  title: '',
  url: '',
  orderId: 0,
  status: 'active'
})

// 计算属性
const dialogTitle = computed(() => {
  return props.mode === 'add' ? '添加轮播图' : '编辑轮播图'
})

const uploadAction = computed(() => {
  // 这里应该是实际的上传接口地址
  return '/api/fileUploadAndDownload/upload'
})

const uploadHeaders = computed(() => {
  // 添加认证头部
  const token = localStorage.getItem('token')
  return token ? { 'x-token': token } : {}
})

// 表单验证规则
const formRules = {
  image: [
    { required: true, message: '请上传轮播图片', trigger: 'blur' }
  ],
  title: [
    { required: true, message: '请输入轮播图标题', trigger: 'blur' },
    { min: 1, max: 100, message: '标题长度在1到100个字符', trigger: 'blur' },
    { 
      validator: (rule, value, callback) => {
        if (value && value.trim().length === 0) {
          callback(new Error('标题不能为空白字符'))
          return
        }
        callback()
      }, 
      trigger: 'blur' 
    }
  ],
  url: [
    { required: true, message: '请输入跳转链接', trigger: 'blur' },
    { 
      validator: (rule, value, callback) => {
        if (!value || value.trim().length === 0) {
          callback(new Error('请输入跳转链接'))
          return
        }
        
        const trimmedValue = value.trim()
        
        if (urlType.value === 'external') {
          const urlPattern = /^https?:\/\/.+/
          if (!urlPattern.test(trimmedValue)) {
            callback(new Error('外部链接必须以 http:// 或 https:// 开头'))
            return
          }
        } else {
          if (!trimmedValue.startsWith('/')) {
            callback(new Error('内部链接必须以 / 开头'))
            return
          }
        }
        
        callback()
      }, 
      trigger: 'blur' 
    }
  ],
  orderId: [
    { required: true, message: '请输入排序序号', trigger: 'blur' },
    { type: 'number', min: 0, max: 9999, message: '排序序号范围0-9999', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'change' }
  ]
}

// 监听弹窗显示状态
watch(() => props.visible, (newVal) => {
  if (newVal) {
    initFormData()
  }
})

// 监听轮播图数据变化
watch(() => props.carouselData, () => {
  if (props.visible) {
    initFormData()
  }
})

// 监听URL类型变化
watch(urlType, (newType) => {
  if (formData.url) {
    if (newType === 'external' && !formData.url.startsWith('http')) {
      formData.url = 'https://'
    } else if (newType === 'internal' && formData.url.startsWith('http')) {
      formData.url = '/'
    }
  }
})

// 初始化表单数据
const initFormData = () => {
  if (props.mode === 'edit' && props.carouselData) {
    Object.assign(formData, {
      ...props.carouselData
    })
    
    // 判断URL类型
    if (formData.url && formData.url.startsWith('http')) {
      urlType.value = 'external'
    } else {
      urlType.value = 'internal'
    }
  } else {
    // 重置表单数据
    Object.assign(formData, {
      image: '',
      title: '',
      url: '/',
      orderId: getNextOrderId(),
      status: 'active'
    })
    urlType.value = 'internal'
  }
  
  // 清除表单验证
  nextTick(() => {
    formRef.value?.clearValidate()
  })
}

// 获取下一个排序ID
const getNextOrderId = () => {
  // 从父组件传入的最大orderId + 1
  return (props.maxOrderId || 0) + 1
}

// 文件上传前验证
const beforeUpload = (file) => {
  const isValidType = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'].includes(file.type)
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isValidType) {
    ElMessage.error('只能上传 jpg、png、gif、webp 格式的图片!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('上传图片大小不能超过 2MB!')
    return false
  }
  
  uploading.value = true
  uploadProgress.value = 0
  return true
}

// 上传成功处理
const handleUploadSuccess = (response) => {
  uploading.value = false
  uploadProgress.value = 100
  
  if (response.code === 0) {
    formData.image = response.data.file?.url || response.data.url || response.data
    ElMessage.success('图片上传成功')
    
    // 触发表单验证
    formRef.value?.validateField('image')
  } else {
    ElMessage.error(response.msg || '图片上传失败')
  }
}

// 上传失败处理
const handleUploadError = (error) => {
  uploading.value = false
  uploadProgress.value = 0
  console.error('图片上传失败:', error)
  ElMessage.error('图片上传失败，请稍后重试')
}

// 上传进度处理
const handleUploadProgress = (event) => {
  uploadProgress.value = Math.round((event.loaded / event.total) * 100)
}

// 更换图片
const replaceImage = () => {
  uploadRef.value?.$el.querySelector('input').click()
}

// 保存处理
const handleSave = async () => {
  try {
    const valid = await formRef.value?.validate()
    if (!valid) {
      ElMessage.warning('请完善表单信息')
      return
    }

    // 验证图片是否已上传
    if (!formData.image || formData.image.trim().length === 0) {
      ElMessage.error('请上传轮播图片')
      return
    }

    // 清理表单数据
    const submitData = {
      ...formData,
      title: formData.title.trim(),
      url: formData.url.trim(),
      orderId: Number(formData.orderId)
    }

    // 如果是编辑模式，需要包含ID
    if (props.mode === 'edit' && props.carouselData?.ID) {
      submitData.ID = props.carouselData.ID
    }

    saveLoading.value = true
    
    const apiCall = props.mode === 'add' ? createCarouselImage : updateCarouselImage
    const response = await apiCall(submitData)
    
    if (response.code === 0) {
      ElMessage.success(props.mode === 'add' ? '轮播图添加成功' : '轮播图保存成功')
      emit('save', { ...submitData, ...response.data })
    } else {
      ElMessage.error(response.msg || '操作失败，请检查网络连接')
    }
  } catch (error) {
    console.error('保存失败:', error)
    
    // 根据错误类型提供不同的提示
    if (error.response) {
      const status = error.response.status
      if (status === 401) {
        ElMessage.error('登录已过期，请重新登录')
      } else if (status === 403) {
        ElMessage.error('没有权限执行此操作')
      } else if (status === 500) {
        ElMessage.error('服务器内部错误，请稍后重试')
      } else {
        ElMessage.error('网络错误，请检查网络连接')
      }
    } else if (error.code === 'NETWORK_ERROR') {
      ElMessage.error('网络连接失败，请检查网络设置')
    } else {
      ElMessage.error('操作失败，请稍后重试')
    }
  } finally {
    saveLoading.value = false
  }
}

// 取消处理
const handleCancel = () => {
  emit('cancel')
}

// 关闭处理
const handleClose = (done) => {
  if (saveLoading.value) {
    ElMessage.warning('正在保存中，请稍候...')
    return
  }
  done()
  emit('cancel')
}
</script>

<style scoped lang="scss">
.upload-section {
  width: 100%;

  :deep(.el-upload) {
    width: 100%;
    
    .el-upload-dragger {
      width: 100%;
      height: 200px;
      border-radius: 6px;
      position: relative;
      overflow: hidden;
    }
  }

  .upload-placeholder {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    color: #8c939d;

    .upload-progress {
      width: 80%;
      text-align: center;

      .progress-text {
        margin-top: 12px;
        font-size: 14px;
        color: #67c23a;
      }
    }

    .upload-content {
      display: flex;
      flex-direction: column;
      align-items: center;

      .upload-icon {
        font-size: 48px;
        margin-bottom: 16px;
      }

      .upload-text {
        font-size: 16px;
        margin-bottom: 8px;
      }

      .upload-tip {
        font-size: 12px;
        color: #a8abb2;
      }
    }
  }

  .image-preview {
    position: relative;
    width: 100%;
    height: 200px;

    .image-overlay {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: rgba(0, 0, 0, 0.5);
      display: flex;
      align-items: center;
      justify-content: center;
      opacity: 0;
      transition: opacity 0.3s;
      border-radius: 6px;

      &:hover {
        opacity: 1;
      }
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

// 表单样式优化
:deep(.el-form) {
  .el-form-item__label {
    font-weight: 500;
    color: #606266;
  }

  .el-input-number {
    .el-input__inner {
      text-align: left;
    }
  }
}

// 响应式设计
@media (max-width: 768px) {
  :deep(.el-dialog) {
    width: 90% !important;
    margin: 5vh auto;
  }

  .upload-section {
    :deep(.el-upload-dragger) {
      height: 150px;
    }

    .upload-placeholder {
      .upload-icon {
        font-size: 36px;
        margin-bottom: 12px;
      }

      .upload-text {
        font-size: 14px;
      }
    }

    .image-preview {
      height: 150px;
    }
  }
}
</style>