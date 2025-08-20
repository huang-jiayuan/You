<template>
  <div class="user-profile">
    <!-- 顶部导航栏 -->
    <div class="profile-header">
      <div class="header-content">
        <button class="back-btn" @click="goBack">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
            <path d="M15 18L9 12L15 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </button>
        <h1 class="header-title">{{ isOwnProfile ? '我的主页' : userInfo.name + '的主页' }}</h1>
        <div class="header-actions">
          <button v-if="isOwnProfile && !isEditing" class="action-btn" @click="startEdit">
            编辑资料
          </button>
          <template v-if="isOwnProfile && isEditing">
            <button class="action-btn cancel-btn" @click="cancelEdit">
              取消
            </button>
            <button class="action-btn save-btn" @click="saveEdit">
              保存
            </button>
          </template>
          <button v-if="!isOwnProfile" class="action-btn" @click="toggleFollow">
            {{ isFollowing ? '已关注' : '关注' }}
          </button>
          <button v-if="!isOwnProfile" class="action-btn" @click="sendMessage">
            发消息
          </button>
        </div>
      </div>
    </div>

    <!-- 用户信息卡片 -->
    <div class="user-info-card">
      <div class="user-avatar-section">
        <div class="avatar-container">
          <img 
            :src="userInfo.avatar" 
            :alt="userInfo.name" 
            class="user-avatar"
            @error="handleAvatarError"
          />
          <div v-if="userInfo.isOnline" class="online-indicator"></div>
        </div>
        <div class="user-basic-info">
          <div v-if="!isEditing">
            <h2 class="user-name">{{ userInfo.name }}</h2>
            <p class="user-id">ID: {{ userInfo.id }}</p>
            <p class="user-signature">{{ userInfo.signature || '这个人很懒，什么都没有留下~' }}</p>
          </div>
          <div v-else class="edit-form">
            <div class="form-group">
              <label>头像：</label>
              <div class="avatar-upload-section">
                <div class="current-avatar">
                  <img :src="editForm.avatar || '/default-avatar.png'" alt="当前头像" class="preview-avatar" />
                </div>
                <div class="upload-controls">
                  <input 
                    ref="avatarInput" 
                    type="file" 
                    accept="image/*" 
                    @change="handleAvatarUpload" 
                    style="display: none;"
                  />
                  <button type="button" @click="$refs.avatarInput.click()" class="upload-btn">
                    选择头像
                  </button>
                  <button type="button" @click="resetAvatar" class="reset-btn">
                    重置默认
                  </button>
                </div>
              </div>
            </div>
            <div class="form-group">
              <label>用户昵称：</label>
              <input v-model="editForm.name" type="text" class="form-input" placeholder="请输入用户昵称">
            </div>
            <div class="form-group">
              <label>个性签名：</label>
              <textarea v-model="editForm.signature" class="form-textarea" placeholder="请输入个性签名"></textarea>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 统计信息 -->
      <div class="stats-section">
        <div class="stat-item" @click="showFollowingList">
          <span class="stat-number">{{ userInfo.followingCount || 0 }}</span>
          <span class="stat-label">关注</span>
        </div>
        <div class="stat-item" @click="showFollowersList">
          <span class="stat-number">{{ userInfo.followersCount || 0 }}</span>
          <span class="stat-label">粉丝</span>
        </div>
        <div class="stat-item">
          <span class="stat-number">{{ userInfo.postsCount || 0 }}</span>
          <span class="stat-label">动态</span>
        </div>
      </div>
    </div>

    <!-- 标签页导航 -->
    <div class="tabs-container">
      <div class="tabs-nav">
        <button 
          v-for="tab in tabs" 
          :key="tab.key"
          :class="['tab-btn', { active: activeTab === tab.key }]"
          @click="activeTab = tab.key"
        >
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- 内容区域 -->
    <div class="content-area">
      <!-- 动态列表 -->
      <div v-if="activeTab === 'posts'" class="posts-list">
        <div v-if="posts.length === 0" class="empty-state">
          <div class="empty-icon">📝</div>
          <p>{{ isOwnProfile ? '还没有发布动态' : 'TA还没有发布动态' }}</p>
        </div>
        <div v-else>
          <div v-for="post in posts" :key="post.id" class="post-item">
            <div class="post-header">
              <img :src="userInfo.avatar" :alt="userInfo.name" class="post-avatar" />
              <div class="post-info">
                <span class="post-author">{{ userInfo.name }}</span>
                <span class="post-time">{{ formatTime(post.createdAt) }}</span>
              </div>
            </div>
            <div class="post-content">
              <p>{{ post.content }}</p>
              <div v-if="post.images && post.images.length > 0" class="post-images">
                <img v-for="image in post.images" :key="image" :src="image" class="post-image" />
              </div>
            </div>
            <div class="post-actions">
              <button class="action-btn-small" @click="likePost(post)">
                <span class="action-icon">❤️</span>
                <span>{{ post.likesCount || 0 }}</span>
              </button>
              <button class="action-btn-small" @click="commentPost(post)">
                <span class="action-icon">💬</span>
                <span>{{ post.commentsCount || 0 }}</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 相册 -->
      <div v-if="activeTab === 'photos'" class="photos-grid">
        <div v-if="photos.length === 0" class="empty-state">
          <div class="empty-icon">📷</div>
          <p>{{ isOwnProfile ? '还没有上传照片' : 'TA还没有上传照片' }}</p>
        </div>
        <div v-else class="photos-container">
          <div v-for="photo in photos" :key="photo.id" class="photo-item" @click="viewPhoto(photo)">
            <img :src="photo.thumbnail" :alt="photo.title" class="photo-thumbnail" />
          </div>
        </div>
      </div>

      <!-- 关于 -->
      <div v-if="activeTab === 'about'" class="about-section">
        <div class="info-group">
          <h3>基本信息</h3>
          <div v-if="!isEditing">
            <div class="info-item">
              <span class="info-label">性别：</span>
              <span class="info-value">{{ userInfo.gender || '未设置' }}</span>
            </div>
            <div class="info-item">
              <span class="info-label">年龄：</span>
              <span class="info-value">{{ userInfo.age || '未设置' }}</span>
            </div>
            <div class="info-item">
              <span class="info-label">地区：</span>
              <span class="info-value">{{ userInfo.location || '未设置' }}</span>
            </div>
            <div class="info-item">
              <span class="info-label">注册时间：</span>
              <span class="info-value">{{ formatDate(userInfo.createdAt) }}</span>
            </div>
          </div>
          <div v-else class="edit-form">
            <div class="form-group">
              <label class="info-label">性别：</label>
              <select v-model="editForm.gender" class="form-select">
                <option value="男">男</option>
                <option value="女">女</option>
                <option value="保密">保密</option>
              </select>
            </div>
            <div class="form-group">
              <label class="info-label">年龄：</label>
              <input v-model="editForm.age" type="number" class="form-input" placeholder="请输入年龄" min="1" max="120">
            </div>
            <div class="form-group">
              <label class="info-label">地区：</label>
              <input v-model="editForm.location" type="text" class="form-input" placeholder="请输入地区">
            </div>
            <div class="info-item">
              <span class="info-label">注册时间：</span>
              <span class="info-value">{{ formatDate(userInfo.createdAt) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { userAPI } from '@/api/user.js'

const route = useRoute()
const router = useRouter()

// 响应式数据
const activeTab = ref('posts')
const isFollowing = ref(false)
const isEditing = ref(false)
const userInfo = ref({
  id: '',
  name: '',
  avatar: '',
  signature: '',
  isOnline: false,
  followingCount: 0,
  followersCount: 0,
  postsCount: 0,
  gender: '',
  age: '',
  location: '',
  createdAt: ''
})

// 编辑表单数据
const editForm = ref({
  name: '',
  signature: '',
  gender: '',
  age: '',
  location: '',
  avatar: '' // 添加头像字段
})

const posts = ref([])
const photos = ref([])

// 标签页配置
const tabs = [
  { key: 'posts', label: '动态' },
  { key: 'photos', label: '相册' },
  { key: 'about', label: '关于' }
]

// 计算属性
const isOwnProfile = computed(() => {
  const userId = route.params.userId
  return !userId || userId === 'me' || userId === getCurrentUserId()
})

// 方法
const goBack = () => {
  router.go(-1)
}

const getCurrentUserId = () => {
  // 这里应该从用户状态管理中获取当前用户ID
  return localStorage.getItem('userId') || '1'
}

const loadUserInfo = async () => {
  try {
    const userId = route.params.userId || getCurrentUserId()
    
    console.log('Loading user info for userId:', userId)
    
    // 调用API获取用户信息
    const response = await userAPI.getUserInfo(userId)
    
    console.log('API Response:', response)
    console.log('response.data:', response.data)
    console.log('response.data.list:', response.data.list)
    console.log('response.data.list.length:', response.data.list ? response.data.list.length : 'undefined')
    console.log('条件判断结果:', !!(response && response.data && response.data.list && response.data.list.length > 0))
    
    if (response && response.data && response.data.list && response.data.list.length > 0) {
      console.log('✅ 进入了if分支 - 使用API数据')
      // 获取数组中的第一个用户数据
      const userData = response.data.list[0]
      
      console.log('User Data:', userData)
      
      console.log('Nickname fields:', {
        Nickname: userData.Nickname,
        nickname: userData.nickname, 
        name: userData.name
      })
      userInfo.value = {
        id: userData.id || userId,
        name: userData.Nickname || userData.nickname || userData.name || '用户' + (userData.id || userId),
        avatar: '/default-avatar.png',
        signature: userData.signature || '这个人很懒，什么都没有留下~',
        isOnline: userData.is_online || false,
        followingCount: userData.following_count || 0,
        followersCount: userData.followers_count || 0,
        postsCount: userData.posts_count || 0,
        // 从API获取性别信息并转换显示
        gender: userData.Gender === '0' ? '女' : userData.Gender === '1' ? '男' : '未设置',
        age: userData.age || '未设置',
        location: userData.location || '未设置',
        createdAt: userData.created_at || new Date().toISOString(),
        // 从API获取余额、钻石等信息
        balance: userData.Balance || 0,
        diamond: userData.Diamond || 0,
        level: userData.Level || 1,
        vipStatus: userData.VipStatus || '0'
      }
      
      console.log('Final userInfo name:', userInfo.value.name)
      console.log('userData.Nickname value:', userData.Nickname)
      console.log('Is entering else branch?')
    } else {
      console.log('❌ 进入了else分支 - 使用默认数据，这就是问题所在！')
      console.log('No valid data received from API')
      // 如果API没有返回有效数据，使用默认数据
      userInfo.value = {
        id: userId,
        name: isOwnProfile.value ? '我' : '用户' + userId,
        avatar: '/default-avatar.png',
        signature: '这是一个个性签名',
        isOnline: false,
        followingCount: 0,
        followersCount: 0,
        postsCount: 0,
        gender: '未设置',
        age: '未设置',
        location: '未设置',
        createdAt: new Date().toISOString()
      }
    }
  } catch (error) {
    console.error('Error loading user info:', error)
    // 错误处理
    if (error.response?.status === 401) {
      router.push('/login')
    } else {
      userInfo.value = {
        id: getCurrentUserId(),
        name: '桑梨', // 临时硬编码测试
        avatar: '/default-avatar.png',
        signature: '这个人很懒，什么都没有留下~',
        isOnline: false,
        followingCount: 0,
        followersCount: 0,
        postsCount: 0,
        gender: '未知',
        age: '未知',
        location: '未知',
        createdAt: new Date().toISOString()
      }
    }
  }
}

const toggleFollow = async () => {
  try {
    isFollowing.value = !isFollowing.value
    // 这里应该调用关注/取消关注的API
    console.log(isFollowing.value ? '已关注' : '已取消关注')
  } catch (error) {
    console.error('操作失败:', error)
  }
}

const sendMessage = () => {
  router.push(`/chat/${userInfo.value.id}`)
}

const showFollowingList = () => {
  console.log('显示关注列表')
}

const showFollowersList = () => {
  console.log('显示粉丝列表')
}

const likePost = (post) => {
  post.likesCount = (post.likesCount || 0) + 1
}

const commentPost = (post) => {
  console.log('评论动态:', post.id)
}

const viewPhoto = (photo) => {
  console.log('查看照片:', photo.id)
}

const formatTime = (dateString) => {
  const date = new Date(dateString)
  const now = new Date()
  const diff = now - date
  
  if (diff < 60000) return '刚刚'
  if (diff < 3600000) return Math.floor(diff / 60000) + '分钟前'
  if (diff < 86400000) return Math.floor(diff / 3600000) + '小时前'
  return Math.floor(diff / 86400000) + '天前'
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('zh-CN')
}

// 编辑相关方法
const startEdit = () => {
  // 复制当前用户信息到编辑表单
  editForm.value = {
    name: userInfo.value.name,
    signature: userInfo.value.signature,
    gender: userInfo.value.gender,
    age: userInfo.value.age,
    location: userInfo.value.location,
    avatar: userInfo.value.avatar // 添加当前头像
  }
  isEditing.value = true
}

const cancelEdit = () => {
  isEditing.value = false
  // 重置编辑表单
  editForm.value = {
    name: '',
    signature: '',
    gender: '',
    age: '',
    location: ''
  }
}

const saveEdit = async () => {
  try {
    // 表单验证
    if (!editForm.value.name.trim()) {
      alert('用户昵称不能为空')
      return
    }
    
    if (editForm.value.age && (isNaN(editForm.value.age) || editForm.value.age < 1 || editForm.value.age > 120)) {
      alert('请输入有效的年龄（1-120）')
      return
    }
    
    // 准备要更新的数据
    const updateData = {
      nickname: editForm.value.name,
      signature: editForm.value.signature,
      gender: editForm.value.gender,
      age: editForm.value.age,
      location: editForm.value.location,
      avatar: editForm.value.avatar // 添加头像字段
    }
    
    // ... existing code ...
    
    // 更新本地用户信息
    userInfo.value.name = editForm.value.name
    userInfo.value.signature = editForm.value.signature
    userInfo.value.gender = editForm.value.gender
    userInfo.value.age = editForm.value.age
    userInfo.value.location = editForm.value.location
    userInfo.value.avatar = editForm.value.avatar // 更新头像
    
    isEditing.value = false
    console.log('用户信息已保存')
  } catch (error) {
    console.error('保存失败:', error)
    alert('保存失败，请重试')
  }
}

// 生命周期
onMounted(() => {
  loadUserInfo()
})

// 添加头像上传处理方法
const handleAvatarUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    // 检查文件类型
    if (!file.type.startsWith('image/')) {
      alert('请选择图片文件')
      return
    }
    
    // 检查文件大小（限制为2MB）
    if (file.size > 2 * 1024 * 1024) {
      alert('图片大小不能超过2MB')
      return
    }
    
    // 创建FileReader来预览图片
    const reader = new FileReader()
    reader.onload = (e) => {
      editForm.value.avatar = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

// 重置头像为默认头像
const resetAvatar = () => {
  editForm.value.avatar = '/default-avatar.png'
}
</script>

<style scoped>
.avatar-upload-section {
  display: flex;
  align-items: center;
  gap: 15px;
}

.current-avatar {
  flex-shrink: 0;
}

.preview-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e0e0e0;
}

.upload-controls {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.upload-btn, .reset-btn {
  padding: 8px 16px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s ease;
}

.upload-btn:hover {
  background: #f0f0f0;
  border-color: #999;
}

.reset-btn {
  background: #f5f5f5;
  color: #666;
}

.reset-btn:hover {
  background: #e0e0e0;
}

.user-profile {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.profile-header {
  position: sticky;
  top: 0;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  z-index: 100;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem;
  max-width: 800px;
  margin: 0 auto;
}

.back-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 50%;
  transition: background-color 0.3s;
}

.back-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.header-title {
  font-size: 1.2rem;
  font-weight: 600;
  margin: 0;
}

.header-actions {
  display: flex;
  gap: 0.5rem;
}

.action-btn {
  background: rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

.cancel-btn {
  background: rgba(255, 107, 107, 0.8) !important;
}

.cancel-btn:hover {
  background: rgba(255, 107, 107, 1) !important;
}

.save-btn {
  background: rgba(103, 194, 58, 0.8) !important;
}

.save-btn:hover {
  background: rgba(103, 194, 58, 1) !important;
}

.edit-form {
    width: 100%;
    animation: fadeIn 0.3s ease-in-out;
  }
  
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(-10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  
  .form-group {
    margin-bottom: 1rem;
    position: relative;
  }

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500;
}

.form-input,
  .form-textarea,
  .form-select {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid rgba(255, 255, 255, 0.3);
    border-radius: 8px;
    background: rgba(255, 255, 255, 0.1);
    color: white;
    font-size: 0.9rem;
    transition: all 0.3s ease;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  
  .form-input:focus,
  .form-textarea:focus,
  .form-select:focus {
    outline: none;
    border-color: rgba(255, 255, 255, 0.6);
    background: rgba(255, 255, 255, 0.15);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2), 0 0 0 2px rgba(255, 255, 255, 0.1);
    transform: translateY(-1px);
  }

.form-input::placeholder,
.form-textarea::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

.form-textarea {
  min-height: 80px;
  resize: vertical;
}

.form-select {
  cursor: pointer;
}

.form-select option {
  background: #333;
  color: white;
}

.user-info-card {
  padding: 2rem 1rem;
  max-width: 800px;
  margin: 0 auto;
}

.user-avatar-section {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2rem;
}

.avatar-container {
  position: relative;
}

.user-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  border: 3px solid rgba(255, 255, 255, 0.3);
}

.online-indicator {
  position: absolute;
  bottom: 5px;
  right: 5px;
  width: 16px;
  height: 16px;
  background: #4ade80;
  border: 2px solid white;
  border-radius: 50%;
}

.user-basic-info {
  flex: 1;
}

.user-name {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0 0 0.25rem 0;
}

.user-id {
  color: rgba(255, 255, 255, 0.7);
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
}

.user-signature {
  color: rgba(255, 255, 255, 0.8);
  margin: 0;
  font-style: italic;
}

.stats-section {
  display: flex;
  justify-content: space-around;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 15px;
  padding: 1rem;
}

.stat-item {
  text-align: center;
  cursor: pointer;
  transition: transform 0.2s;
}

.stat-item:hover {
  transform: translateY(-2px);
}

.stat-number {
  display: block;
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.stat-label {
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.9rem;
}

.tabs-container {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  position: sticky;
  top: 73px;
  z-index: 90;
}

.tabs-nav {
  display: flex;
  max-width: 800px;
  margin: 0 auto;
}

.tab-btn {
  flex: 1;
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.7);
  padding: 1rem;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s;
  border-bottom: 2px solid transparent;
}

.tab-btn.active {
  color: white;
  border-bottom-color: white;
}

.tab-btn:hover {
  color: white;
  background: rgba(255, 255, 255, 0.1);
}

.content-area {
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem;
}

.empty-state {
  text-align: center;
  padding: 3rem 1rem;
  color: rgba(255, 255, 255, 0.7);
}

.empty-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.posts-list {
  space-y: 1rem;
}

.post-item {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 15px;
  padding: 1rem;
  margin-bottom: 1rem;
}

.post-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.75rem;
}

.post-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.post-info {
  flex: 1;
}

.post-author {
  display: block;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.post-time {
  color: rgba(255, 255, 255, 0.6);
  font-size: 0.8rem;
}

.post-content p {
  margin: 0 0 0.75rem 0;
  line-height: 1.5;
}

.post-images {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}

.post-image {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 8px;
}

.post-actions {
  display: flex;
  gap: 1rem;
}

.action-btn-small {
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.5rem;
  border-radius: 15px;
  transition: all 0.3s;
}

.action-btn-small:hover {
  color: white;
  background: rgba(255, 255, 255, 0.1);
}

.photos-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 0.5rem;
}

.photo-item {
  aspect-ratio: 1;
  cursor: pointer;
  border-radius: 8px;
  overflow: hidden;
  transition: transform 0.2s;
}

.photo-item:hover {
  transform: scale(1.05);
}

.photo-thumbnail {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.about-section {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 15px;
  padding: 1.5rem;
}

.info-group h3 {
  margin: 0 0 1rem 0;
  font-size: 1.1rem;
  font-weight: 600;
}

.info-item {
  display: flex;
  margin-bottom: 0.75rem;
}

.info-label {
  width: 80px;
  color: rgba(255, 255, 255, 0.7);
}

.info-value {
  flex: 1;
}

@media (max-width: 768px) {
  .user-avatar-section {
    flex-direction: column;
    text-align: center;
  }
  
  .stats-section {
    flex-direction: column;
    gap: 1rem;
  }
  
  .photos-container {
    grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  }
}
</style>
const handleAvatarError = (event) => {
  console.log('Avatar load error, using default avatar')
  event.target.src = '/default-avatar.png'
}