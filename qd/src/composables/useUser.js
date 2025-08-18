import { computed } from 'vue'
import { getGlobalAppStore } from '../stores/app.js'

/**
 * 用户状态管理组合式函数
 * 提供用户相关的状态管理功能
 */
export function useUser() {
  const store = getGlobalAppStore()
  
  // 计算属性
  const user = computed(() => store.state.user)
  const isLoggedIn = computed(() => store.isLoggedIn.value)
  const currentUser = computed(() => store.currentUser.value)
  const preferences = computed(() => store.state.user.preferences)
  
  // 用户信息管理
  const updateProfile = (profileData) => {
    const { preferences: prefs, ...userInfo } = profileData
    
    // 更新用户基本信息
    if (Object.keys(userInfo).length > 0) {
      store.setUser(userInfo)
    }
    
    // 更新用户偏好设置
    if (prefs) {
      store.updateUserPreferences(prefs)
    }
    
    store.updateActivity()
  }
  
  const updatePreferences = (newPreferences) => {
    store.updateUserPreferences(newPreferences)
  }
  
  const login = async (credentials) => {
    try {
      // 这里应该调用实际的登录API
      // const response = await authAPI.login(credentials)
      
      // 模拟登录过程
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      const userData = {
        name: credentials.username || '用户',
        email: credentials.email || 'user@example.com',
        avatar: credentials.avatar || '',
        preferences: {
          language: 'zh-CN',
          timezone: 'Asia/Shanghai',
          notifications: true
        }
      }
      
      store.setUser(userData)
      
      return {
        success: true,
        user: userData
      }
    } catch (error) {
      return {
        success: false,
        error: error.message
      }
    }
  }
  
  const logout = async () => {
    try {
      // 这里应该调用实际的登出API
      // await authAPI.logout()
      
      // 清除用户数据
      store.clearUser()
      
      // 清除本地存储中的敏感数据
      const currentState = JSON.parse(localStorage.getItem('app-state') || '{}')
      if (currentState.user) {
        delete currentState.user
        localStorage.setItem('app-state', JSON.stringify(currentState))
      }
      
      return { success: true }
    } catch (error) {
      return {
        success: false,
        error: error.message
      }
    }
  }
  
  const register = async (userData) => {
    try {
      // 这里应该调用实际的注册API
      // const response = await authAPI.register(userData)
      
      // 模拟注册过程
      await new Promise(resolve => setTimeout(resolve, 1500))
      
      const newUser = {
        name: userData.name,
        email: userData.email,
        avatar: userData.avatar || '',
        preferences: {
          language: userData.language || 'zh-CN',
          timezone: userData.timezone || 'Asia/Shanghai',
          notifications: userData.notifications !== false
        }
      }
      
      store.setUser(newUser)
      
      return {
        success: true,
        user: newUser
      }
    } catch (error) {
      return {
        success: false,
        error: error.message
      }
    }
  }
  
  // 偏好设置管理
  const setLanguage = (language) => {
    updatePreferences({ language })
  }
  
  const setTimezone = (timezone) => {
    updatePreferences({ timezone })
  }
  
  const toggleNotifications = () => {
    updatePreferences({ 
      notifications: !preferences.value.notifications 
    })
  }
  
  // 用户头像管理
  const updateAvatar = async (avatarFile) => {
    try {
      // 这里应该上传头像到服务器
      // const response = await uploadAPI.uploadAvatar(avatarFile)
      
      // 模拟上传过程
      await new Promise(resolve => setTimeout(resolve, 2000))
      
      const avatarUrl = URL.createObjectURL(avatarFile)
      store.setUser({ avatar: avatarUrl })
      
      return {
        success: true,
        avatarUrl
      }
    } catch (error) {
      return {
        success: false,
        error: error.message
      }
    }
  }
  
  // 用户统计信息
  const getUserStats = () => {
    return computed(() => ({
      sessionDuration: store.sessionDuration.value,
      pageViews: store.state.stats.pageViews,
      userActions: store.state.stats.userActions,
      lastActivity: store.state.stats.lastActivity,
      joinDate: user.value.joinDate || Date.now()
    }))
  }
  
  // 用户验证
  const validateUser = () => {
    const errors = {}
    
    if (!user.value.name || user.value.name.trim().length < 2) {
      errors.name = '姓名至少需要2个字符'
    }
    
    if (!user.value.email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(user.value.email)) {
      errors.email = '请输入有效的邮箱地址'
    }
    
    return {
      isValid: Object.keys(errors).length === 0,
      errors
    }
  }
  
  // 用户权限检查（示例）
  const hasPermission = (permission) => {
    // 这里可以实现复杂的权限检查逻辑
    const userPermissions = user.value.permissions || []
    return userPermissions.includes(permission) || userPermissions.includes('admin')
  }
  
  const hasRole = (role) => {
    const userRoles = user.value.roles || []
    return userRoles.includes(role)
  }
  
  // 用户活动跟踪
  const trackActivity = (action, data = {}) => {
    store.incrementUserActions()
    
    // 这里可以发送活动数据到分析服务
    console.log('用户活动:', {
      user: user.value.email,
      action,
      data,
      timestamp: Date.now()
    })
  }
  
  return {
    // 状态
    user,
    isLoggedIn,
    currentUser,
    preferences,
    
    // 用户管理
    updateProfile,
    updatePreferences,
    login,
    logout,
    register,
    
    // 偏好设置
    setLanguage,
    setTimezone,
    toggleNotifications,
    
    // 头像管理
    updateAvatar,
    
    // 统计信息
    getUserStats,
    
    // 验证
    validateUser,
    
    // 权限
    hasPermission,
    hasRole,
    
    // 活动跟踪
    trackActivity
  }
}