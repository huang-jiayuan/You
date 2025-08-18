/**
 * 音频播放器组合式函数
 * 用于管理移动端音频播放状态和控制
 */

import { ref, reactive, computed, onUnmounted } from 'vue'

interface AudioTrack {
  id: string
  title: string
  artist: string
  url: string
  cover?: string
  duration?: number
}

interface AudioPlayerState {
  currentTrack: AudioTrack | null
  isPlaying: boolean
  currentTime: number
  duration: number
  volume: number
  isLoading: boolean
  error: string | null
}

export function useAudioPlayer() {
  // 音频播放器状态
  const state = reactive<AudioPlayerState>({
    currentTrack: null,
    isPlaying: false,
    currentTime: 0,
    duration: 0,
    volume: 1,
    isLoading: false,
    error: null
  })

  // 音频元素引用
  const audioElement = ref<HTMLAudioElement | null>(null)
  
  // 播放列表
  const playlist = ref<AudioTrack[]>([])
  const currentIndex = ref(-1)

  // 计算属性
  const progress = computed(() => {
    return state.duration > 0 ? (state.currentTime / state.duration) * 100 : 0
  })

  const formattedCurrentTime = computed(() => {
    return formatTime(state.currentTime)
  })

  const formattedDuration = computed(() => {
    return formatTime(state.duration)
  })

  const canPlayPrevious = computed(() => {
    return currentIndex.value > 0
  })

  const canPlayNext = computed(() => {
    return currentIndex.value < playlist.value.length - 1
  })

  // 初始化音频元素
  const initAudioElement = () => {
    if (!audioElement.value) {
      audioElement.value = new Audio()
      setupAudioEventListeners()
    }
    return audioElement.value
  }

  // 设置音频事件监听器
  const setupAudioEventListeners = () => {
    const audio = audioElement.value
    if (!audio) return

    audio.addEventListener('loadstart', () => {
      state.isLoading = true
      state.error = null
    })

    audio.addEventListener('loadedmetadata', () => {
      state.duration = audio.duration
      state.isLoading = false
    })

    audio.addEventListener('timeupdate', () => {
      state.currentTime = audio.currentTime
    })

    audio.addEventListener('play', () => {
      state.isPlaying = true
    })

    audio.addEventListener('pause', () => {
      state.isPlaying = false
    })

    audio.addEventListener('ended', () => {
      state.isPlaying = false
      playNext()
    })

    audio.addEventListener('error', (e) => {
      state.isLoading = false
      state.isPlaying = false
      state.error = '音频加载失败'
      console.error('Audio error:', e)
    })

    audio.addEventListener('volumechange', () => {
      state.volume = audio.volume
    })
  }

  // 加载音频轨道
  const loadTrack = async (track: AudioTrack) => {
    const audio = initAudioElement()
    
    try {
      state.currentTrack = track
      state.error = null
      state.isLoading = true
      
      audio.src = track.url
      audio.load()
      
      // 等待音频元数据加载
      await new Promise((resolve, reject) => {
        const onLoadedMetadata = () => {
          audio.removeEventListener('loadedmetadata', onLoadedMetadata)
          audio.removeEventListener('error', onError)
          resolve(void 0)
        }
        
        const onError = () => {
          audio.removeEventListener('loadedmetadata', onLoadedMetadata)
          audio.removeEventListener('error', onError)
          reject(new Error('Failed to load audio'))
        }
        
        audio.addEventListener('loadedmetadata', onLoadedMetadata)
        audio.addEventListener('error', onError)
      })
      
    } catch (error) {
      state.error = '音频加载失败'
      state.isLoading = false
      throw error
    }
  }

  // 播放音频
  const play = async () => {
    const audio = audioElement.value
    if (!audio || !state.currentTrack) return

    try {
      await audio.play()
    } catch (error) {
      state.error = '播放失败'
      console.error('Play error:', error)
      throw error
    }
  }

  // 暂停音频
  const pause = () => {
    const audio = audioElement.value
    if (audio) {
      audio.pause()
    }
  }

  // 切换播放/暂停
  const togglePlay = async () => {
    if (state.isPlaying) {
      pause()
    } else {
      await play()
    }
  }

  // 设置播放时间
  const seek = (time: number) => {
    const audio = audioElement.value
    if (audio && state.duration > 0) {
      audio.currentTime = Math.max(0, Math.min(time, state.duration))
    }
  }

  // 设置音量
  const setVolume = (volume: number) => {
    const audio = audioElement.value
    if (audio) {
      audio.volume = Math.max(0, Math.min(1, volume))
    }
  }

  // 播放指定轨道
  const playTrack = async (track: AudioTrack) => {
    try {
      await loadTrack(track)
      await play()
    } catch (error) {
      console.error('Failed to play track:', error)
    }
  }

  // 设置播放列表
  const setPlaylist = (tracks: AudioTrack[], startIndex = 0) => {
    playlist.value = tracks
    currentIndex.value = startIndex
    
    if (tracks.length > 0 && startIndex >= 0 && startIndex < tracks.length) {
      loadTrack(tracks[startIndex])
    }
  }

  // 播放上一首
  const playPrevious = async () => {
    if (canPlayPrevious.value) {
      currentIndex.value--
      await playTrack(playlist.value[currentIndex.value])
    }
  }

  // 播放下一首
  const playNext = async () => {
    if (canPlayNext.value) {
      currentIndex.value++
      await playTrack(playlist.value[currentIndex.value])
    }
  }

  // 停止播放
  const stop = () => {
    const audio = audioElement.value
    if (audio) {
      audio.pause()
      audio.currentTime = 0
    }
  }

  // 格式化时间
  const formatTime = (seconds: number): string => {
    if (!isFinite(seconds)) return '0:00'
    
    const mins = Math.floor(seconds / 60)
    const secs = Math.floor(seconds % 60)
    return `${mins}:${secs.toString().padStart(2, '0')}`
  }

  // 清理资源
  const cleanup = () => {
    const audio = audioElement.value
    if (audio) {
      audio.pause()
      audio.src = ''
      audio.load()
    }
  }

  // 组件卸载时清理
  onUnmounted(() => {
    cleanup()
  })

  return {
    // 状态
    state,
    playlist,
    currentIndex,
    
    // 计算属性
    progress,
    formattedCurrentTime,
    formattedDuration,
    canPlayPrevious,
    canPlayNext,
    
    // 方法
    loadTrack,
    play,
    pause,
    togglePlay,
    seek,
    setVolume,
    playTrack,
    setPlaylist,
    playPrevious,
    playNext,
    stop,
    cleanup
  }
}