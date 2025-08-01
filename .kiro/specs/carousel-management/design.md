# 轮播图管理功能设计文档

## 概述

本设计文档基于现有的gin-vue-admin项目架构，为轮播图组件添加完整的管理功能。项目采用Vue 3 + Element Plus + Vite的前端技术栈，后端使用Gin框架。设计将充分利用现有的API接口和组件结构，提供用户友好的轮播图管理界面。

## 架构

### 技术栈
- **前端**: Vue 3.5.7 + Element Plus 2.8.5 + Vite 6.2.3
- **状态管理**: Pinia 2.2.2
- **路由**: Vue Router 4.4.3
- **HTTP客户端**: Axios 1.8.2
- **样式**: SCSS + TailwindCSS 3.4.10
- **图标**: Element Plus Icons Vue 2.3.1

### 架构模式
采用MVVM架构模式，遵循现有项目的组件化和模块化设计原则：

```
View Layer (Vue Components)
    ↓
ViewModel Layer (Composition API + Pinia)
    ↓
Model Layer (API Services + Data Models)
    ↓
Backend API (Gin + GORM)
```

## 组件和接口

### 1. 核心组件结构

#### 1.1 CarouselManagement.vue (主管理组件)
- **路径**: `gin-vue-admin/web/src/view/carouse/carouselImage/CarouselManagement.vue`
- **功能**: 轮播图列表展示、操作入口
- **组件特性**:
  - 响应式表格布局
  - 搜索和筛选功能
  - 批量操作支持
  - 分页控制

#### 1.2 CarouselDialog.vue (弹窗组件)
- **路径**: `gin-vue-admin/web/src/view/carouse/carouselImage/components/CarouselDialog.vue`
- **功能**: 添加/编辑轮播图的弹窗表单
- **组件特性**:
  - 动态表单验证
  - 图片上传预览
  - 拖拽排序支持

#### 1.3 CarouselPreview.vue (预览组件)
- **路径**: `gin-vue-admin/web/src/view/carouse/carouselImage/components/CarouselPreview.vue`
- **功能**: 轮播图效果预览
- **组件特性**:
  - 实时预览效果
  - 响应式适配

### 2. API接口设计

基于现有的API接口进行扩展和优化：

#### 2.1 现有接口利用
```javascript
// 已存在的接口，直接使用
- createCarouselImage(data)     // 创建轮播图
- updateCarouselImage(data)     // 更新轮播图
- deleteCarouselImage(params)   // 删除单个轮播图
- deleteCarouselImageByIds(params) // 批量删除
- findCarouselImage(params)     // 查询单个轮播图
- getCarouselImageList(params)  // 获取轮播图列表
- getCarouselImagePublic()      // 公开接口获取轮播图
```

#### 2.2 接口数据模型
```javascript
// CarouselImage 数据模型
interface CarouselImageModel {
  ID?: number
  image: string      // 图片URL
  title: string      // 标题
  url: string        // 跳转链接
  orderId: number    // 排序序号
  status: string     // 状态 (active/inactive)
  createdAt?: string
  updatedAt?: string
}
```

### 3. 状态管理设计

#### 3.1 Pinia Store结构
```javascript
// stores/carousel.js
export const useCarouselStore = defineStore('carousel', {
  state: () => ({
    carouselList: [],
    currentCarousel: null,
    loading: false,
    pagination: {
      page: 1,
      pageSize: 10,
      total: 0
    },
    searchParams: {
      title: '',
      status: ''
    }
  }),
  
  actions: {
    async fetchCarouselList(),
    async createCarousel(data),
    async updateCarousel(data),
    async deleteCarousel(id),
    async batchDeleteCarousels(ids)
  }
})
```

## 数据模型

### 1. 轮播图数据模型
```typescript
interface CarouselImage {
  ID: number
  image: string           // 图片URL，支持相对路径和绝对路径
  title: string          // 轮播图标题，最大长度100字符
  url: string            // 点击跳转链接，支持内部路由和外部链接
  orderId: number        // 排序序号，数字越小越靠前
  status: 'active' | 'inactive'  // 状态：激活/禁用
  createdAt: string      // 创建时间
  updatedAt: string      // 更新时间
}
```

### 2. 表单验证规则
```javascript
const validationRules = {
  image: [
    { required: true, message: '请上传轮播图片', trigger: 'blur' },
    { 
      validator: (rule, value, callback) => {
        const validFormats = ['.jpg', '.jpeg', '.png', '.gif', '.webp']
        const isValidFormat = validFormats.some(format => 
          value.toLowerCase().includes(format)
        )
        if (!isValidFormat) {
          callback(new Error('图片格式不支持，请上传jpg、png、gif或webp格式'))
        } else {
          callback()
        }
      }, 
      trigger: 'blur' 
    }
  ],
  title: [
    { required: true, message: '请输入轮播图标题', trigger: 'blur' },
    { min: 1, max: 100, message: '标题长度在1到100个字符', trigger: 'blur' }
  ],
  url: [
    { required: true, message: '请输入跳转链接', trigger: 'blur' },
    { 
      pattern: /^(https?:\/\/|\/)/,
      message: '请输入有效的URL地址或路由路径',
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
```

### 3. 响应数据格式
```javascript
// API响应统一格式
interface ApiResponse<T> {
  code: number
  data: T
  msg: string
}

// 列表响应格式
interface ListResponse {
  list: CarouselImage[]
  total: number
  page: number
  pageSize: number
}
```

## 错误处理

### 1. 前端错误处理策略

#### 1.1 网络错误处理
```javascript
// 在axios拦截器中统一处理
const handleNetworkError = (error) => {
  if (error.code === 'NETWORK_ERROR') {
    ElMessage.error('网络连接失败，请检查网络设置')
    return Promise.reject(new Error('网络错误'))
  }
  
  if (error.response?.status === 500) {
    ElMessage.error('服务器内部错误，请稍后重试')
    return Promise.reject(new Error('服务器错误'))
  }
  
  return Promise.reject(error)
}
```

#### 1.2 表单验证错误
```javascript
// 表单验证失败处理
const handleValidationError = (errors) => {
  const firstError = Object.values(errors)[0]
  ElMessage.error(firstError[0].message)
  
  // 滚动到第一个错误字段
  const firstErrorField = Object.keys(errors)[0]
  const errorElement = document.querySelector(`[prop="${firstErrorField}"]`)
  errorElement?.scrollIntoView({ behavior: 'smooth' })
}
```

#### 1.3 文件上传错误
```javascript
// 图片上传错误处理
const handleUploadError = (error, file) => {
  const errorMessages = {
    'FILE_TOO_LARGE': '图片大小不能超过2MB',
    'INVALID_FORMAT': '不支持的图片格式',
    'UPLOAD_FAILED': '图片上传失败，请重试'
  }
  
  const message = errorMessages[error.code] || '上传失败'
  ElMessage.error(message)
}
```

### 2. 用户体验优化

#### 2.1 加载状态管理
```javascript
// 全局加载状态
const useLoading = () => {
  const loading = ref(false)
  
  const withLoading = async (asyncFn) => {
    loading.value = true
    try {
      return await asyncFn()
    } finally {
      loading.value = false
    }
  }
  
  return { loading, withLoading }
}
```

#### 2.2 操作确认机制
```javascript
// 删除确认
const confirmDelete = (item) => {
  return ElMessageBox.confirm(
    `确定要删除轮播图"${item.title}"吗？此操作不可恢复。`,
    '删除确认',
    {
      confirmButtonText: '确定删除',
      cancelButtonText: '取消',
      type: 'warning',
      dangerouslyUseHTMLString: true
    }
  )
}
```

## 测试策略

### 1. 单元测试

#### 1.1 组件测试
```javascript
// CarouselDialog.vue 测试用例
describe('CarouselDialog', () => {
  test('should validate required fields', async () => {
    const wrapper = mount(CarouselDialog)
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.find('.el-form-item__error').exists()).toBe(true)
  })
  
  test('should emit save event with correct data', async () => {
    const wrapper = mount(CarouselDialog)
    const formData = {
      title: '测试标题',
      image: 'test.jpg',
      url: '/test',
      orderId: 1,
      status: 'active'
    }
    
    await wrapper.setData({ formData })
    await wrapper.find('form').trigger('submit')
    
    expect(wrapper.emitted('save')).toBeTruthy()
    expect(wrapper.emitted('save')[0][0]).toEqual(formData)
  })
})
```

#### 1.2 API服务测试
```javascript
// carousel API 测试
describe('Carousel API', () => {
  test('should create carousel successfully', async () => {
    const mockData = {
      title: '测试轮播图',
      image: 'test.jpg',
      url: '/test',
      orderId: 1,
      status: 'active'
    }
    
    const response = await createCarouselImage(mockData)
    expect(response.code).toBe(0)
    expect(response.data).toHaveProperty('ID')
  })
})
```

### 2. 集成测试

#### 2.1 端到端测试场景
```javascript
// E2E测试用例
describe('Carousel Management E2E', () => {
  test('complete carousel management workflow', async () => {
    // 1. 访问管理页面
    await page.goto('/carousel/management')
    
    // 2. 添加新轮播图
    await page.click('[data-test="add-carousel"]')
    await page.fill('[data-test="title-input"]', '测试轮播图')
    await page.fill('[data-test="url-input"]', '/test')
    await page.click('[data-test="save-button"]')
    
    // 3. 验证添加成功
    await expect(page.locator('.el-message--success')).toBeVisible()
    
    // 4. 编辑轮播图
    await page.click('[data-test="edit-button"]:first-child')
    await page.fill('[data-test="title-input"]', '修改后的标题')
    await page.click('[data-test="save-button"]')
    
    // 5. 验证修改成功
    await expect(page.locator('text=修改后的标题')).toBeVisible()
    
    // 6. 删除轮播图
    await page.click('[data-test="delete-button"]:first-child')
    await page.click('[data-test="confirm-delete"]')
    
    // 7. 验证删除成功
    await expect(page.locator('.el-message--success')).toBeVisible()
  })
})
```

### 3. 性能测试

#### 3.1 组件渲染性能
```javascript
// 性能测试
describe('Performance Tests', () => {
  test('should render large carousel list efficiently', async () => {
    const startTime = performance.now()
    
    const wrapper = mount(CarouselManagement, {
      props: {
        carouselList: generateMockCarousels(1000)
      }
    })
    
    await wrapper.vm.$nextTick()
    const endTime = performance.now()
    
    expect(endTime - startTime).toBeLessThan(100) // 渲染时间小于100ms
  })
})
```

### 4. 可访问性测试

#### 4.1 键盘导航测试
```javascript
// 可访问性测试
describe('Accessibility Tests', () => {
  test('should support keyboard navigation', async () => {
    const wrapper = mount(CarouselManagement)
    
    // 测试Tab键导航
    await wrapper.trigger('keydown', { key: 'Tab' })
    expect(document.activeElement).toBe(wrapper.find('[data-test="add-button"]').element)
    
    // 测试Enter键激活
    await wrapper.trigger('keydown', { key: 'Enter' })
    expect(wrapper.emitted('add-carousel')).toBeTruthy()
  })
  
  test('should have proper ARIA labels', () => {
    const wrapper = mount(CarouselDialog)
    
    expect(wrapper.find('[role="dialog"]').exists()).toBe(true)
    expect(wrapper.find('[aria-label="轮播图标题"]').exists()).toBe(true)
  })
})
```