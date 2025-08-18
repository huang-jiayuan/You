import { ref, readonly, computed, watch, onMounted } from 'vue'

// 全局主题状态
const theme = ref('light')
const isInitialized = ref(false)

/**
 * 主题管理组合式函数
 * @returns {object} 主题相关的响应式数据和方法
 */
export function useTheme() {
  // 计算属性
  const isDark = computed(() => theme.value === 'dark')
  const isLight = computed(() => theme.value === 'light')
  
  // 主题选项
  const themes = [
    'light', 
    'dark', 
    'blue', 
    'green', 
    'purple', 
    'red', 
    'orange', 
    'pink', 
    'cyan',
    'high-contrast',
    'high-contrast-dark',
    'eye-care'
  ]
  
  // 方法
  const setTheme = (newTheme) => {
    if (themes.includes(newTheme)) {
      theme.value = newTheme
      applyTheme(newTheme)
      saveTheme(newTheme)
    }
  }

  const toggleTheme = () => {
    const newTheme = theme.value === 'light' ? 'dark' : 'light'
    setTheme(newTheme)
  }

  const initTheme = () => {
    if (isInitialized.value) return

    // 从localStorage获取保存的主题
    const savedTheme = getSavedTheme()
    
    // 如果没有保存的主题，检查系统偏好
    if (!savedTheme) {
      const systemTheme = getSystemTheme()
      setTheme(systemTheme)
    } else {
      setTheme(savedTheme)
    }

    // 监听系统主题变化
    watchSystemTheme()
    
    isInitialized.value = true
  }

  // 应用主题到DOM
  const applyTheme = (themeName) => {
    document.documentElement.setAttribute('data-theme', themeName)
    
    // 更新meta标签的theme-color
    updateThemeColor(themeName)
  }

  // 保存主题到localStorage
  const saveTheme = (themeName) => {
    try {
      localStorage.setItem('theme', themeName)
    } catch (error) {
      console.warn('无法保存主题设置:', error)
    }
  }

  // 从localStorage获取保存的主题
  const getSavedTheme = () => {
    try {
      return localStorage.getItem('theme')
    } catch (error) {
      console.warn('无法读取主题设置:', error)
      return null
    }
  }

  // 获取系统主题偏好
  const getSystemTheme = () => {
    if (typeof window !== 'undefined' && window.matchMedia) {
      return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
    }
    return 'light'
  }

  // 监听系统主题变化
  const watchSystemTheme = () => {
    if (typeof window !== 'undefined' && window.matchMedia) {
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
      
      const handleChange = (e) => {
        // 只有在没有用户手动设置主题时才跟随系统
        const savedTheme = getSavedTheme()
        if (!savedTheme) {
          const systemTheme = e.matches ? 'dark' : 'light'
          setTheme(systemTheme)
        }
      }

      // 现代浏览器使用addEventListener
      if (mediaQuery.addEventListener) {
        mediaQuery.addEventListener('change', handleChange)
      } else {
        // 旧版浏览器使用addListener
        mediaQuery.addListener(handleChange)
      }
    }
  }

  // 更新主题色meta标签
  const updateThemeColor = (themeName) => {
    const themeColors = {
      light: '#ffffff',
      dark: '#111827'
    }

    let metaThemeColor = document.querySelector('meta[name="theme-color"]')
    if (!metaThemeColor) {
      metaThemeColor = document.createElement('meta')
      metaThemeColor.name = 'theme-color'
      document.head.appendChild(metaThemeColor)
    }
    
    metaThemeColor.content = themeColors[themeName] || themeColors.light
  }

  // 获取主题相关的CSS变量值
  const getThemeVariable = (variableName) => {
    if (typeof window !== 'undefined') {
      return getComputedStyle(document.documentElement)
        .getPropertyValue(`--${variableName}`)
        .trim()
    }
    return ''
  }

  // 设置主题相关的CSS变量
  const setThemeVariable = (variableName, value) => {
    if (typeof window !== 'undefined') {
      document.documentElement.style.setProperty(`--${variableName}`, value)
    }
  }

  return {
    // 响应式数据（只读）
    theme: readonly(theme),
    isDark,
    isLight,
    themes,
    
    // 方法
    setTheme,
    toggleTheme,
    initTheme,
    getThemeVariable,
    setThemeVariable
  }
}

/**
 * 主题持久化组合式函数
 * 提供更高级的主题管理功能
 */
export function useThemeWithPreferences() {
  const { theme, isDark, isLight, setTheme, toggleTheme, initTheme } = useTheme()
  
  // 主题偏好设置
  const preferences = ref({
    followSystem: false,
    autoSwitch: false,
    switchTime: {
      dark: '18:00',
      light: '06:00'
    }
  })

  // 加载偏好设置
  const loadPreferences = () => {
    try {
      const saved = localStorage.getItem('theme-preferences')
      if (saved) {
        Object.assign(preferences.value, JSON.parse(saved))
      }
    } catch (error) {
      console.warn('无法加载主题偏好设置:', error)
    }
  }

  // 保存偏好设置
  const savePreferences = () => {
    try {
      localStorage.setItem('theme-preferences', JSON.stringify(preferences.value))
    } catch (error) {
      console.warn('无法保存主题偏好设置:', error)
    }
  }

  // 根据时间自动切换主题
  const checkAutoSwitch = () => {
    if (!preferences.value.autoSwitch) return

    const now = new Date()
    const currentTime = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`
    
    const { dark: darkTime, light: lightTime } = preferences.value.switchTime
    
    if (currentTime >= darkTime || currentTime < lightTime) {
      if (theme.value !== 'dark') {
        setTheme('dark')
      }
    } else {
      if (theme.value !== 'light') {
        setTheme('light')
      }
    }
  }

  // 监听偏好设置变化
  watch(preferences, savePreferences, { deep: true })

  // 初始化
  onMounted(() => {
    loadPreferences()
    
    if (preferences.value.autoSwitch) {
      checkAutoSwitch()
      // 每分钟检查一次
      setInterval(checkAutoSwitch, 60000)
    }
  })

  return {
    theme,
    isDark,
    isLight,
    preferences: readonly(preferences),
    setTheme,
    toggleTheme,
    initTheme,
    loadPreferences,
    savePreferences
  }
}