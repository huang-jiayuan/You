<template>
  <div class="home-page">
    <!-- 顶部轮播图 -->
    <div class="carousel-section">
      <div class="carousel-container" ref="carouselContainer">
        <div 
          class="carousel-wrapper"
          :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
          @touchstart="handleTouchStart"
          @touchmove="handleTouchMove"
          @touchend="handleTouchEnd"
        >
          <div 
            v-for="(item, index) in carouselImages" 
            :key="index"
            class="carousel-slide"
            @click="handleCarouselClick(item)"
          >
            <img :src="item.image" :alt="item.title" />
            <div class="carousel-overlay">
              <h3>{{ item.title }}</h3>
            </div>
          </div>
        </div>
        
        <!-- 左右切换按钮 -->
        <button class="carousel-btn prev" @click="prevSlide" v-if="carouselImages.length > 1">
          <el-icon><ArrowLeft /></el-icon>
        </button>
        <button class="carousel-btn next" @click="nextSlide" v-if="carouselImages.length > 1">
          <el-icon><ArrowRight /></el-icon>
        </button>
        
        <!-- 指示器 -->
        <div class="carousel-indicators" v-if="carouselImages.length > 1">
          <span 
            v-for="(item, index) in carouselImages" 
            :key="index"
            :class="['indicator', { active: index === currentIndex }]"
            @click="goToSlide(index)"
          ></span>
        </div>
      </div>
    </div>

    <!-- 主标题 -->
    <div class="main-header">
      <h1 class="main-title">防诈骗指南</h1>
      <div class="leaf-decoration">🍃</div>
    </div>

    <!-- 功能按钮区域 -->
    <div class="feature-buttons">
      <div class="feature-btn purple" @click="handleFeatureClick('match')">
        <div class="btn-icon">💕</div>
        <span>牵手速配</span>
      </div>
      <div class="feature-btn orange" @click="handleFeatureClick('auction')">
        <div class="btn-icon">🔥</div>
        <span>激情拍卖</span>
      </div>
      <div class="feature-btn pink" @click="handleFeatureClick('romance')">
        <div class="btn-icon">🏠</div>
        <span>浪漫满屋</span>
      </div>
    </div>

    <!-- 分类标签 -->
    <div class="category-tabs">
      <div 
        v-for="(category, index) in categories" 
        :key="index"
        :class="['category-tab', { active: activeCategory === category.key }]"
        @click="switchCategory(category.key)"
      >
        <span class="tab-icon">{{ category.icon }}</span>
        <span class="tab-text">{{ category.name }}</span>
      </div>
    </div>

    <!-- 内容列表 -->
    <div class="content-list">
      <div 
        v-for="(item, index) in getCurrentContentList()" 
        :key="index"
        class="content-item"
        @click="handleContentClick(item)"
      >
        <div class="content-image">
          <img :src="item.image" :alt="item.title" />
          <div v-if="item.isHot" class="hot-badge">🔥</div>
          <div v-if="item.isNew" class="new-badge">NEW</div>
        </div>
        <div class="content-info">
          <h3 class="content-title">{{ item.title }}</h3>
          <div class="content-description">{{ item.description }}</div>
          <div class="content-meta">
            <span class="tag" :style="{ background: item.tagColor }">{{ item.tag }}</span>
            <div class="stats">
              <span class="views">👁 {{ item.views }}</span>
              <span class="likes">❤ {{ item.likes }}</span>
              <span v-if="item.rating" class="rating">⭐ {{ item.rating }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部导航 -->
    <div class="bottom-nav">
      <div class="nav-item active" @click="handleNavClick('home')">
        <el-icon><House /></el-icon>
        <span>首页</span>
      </div>
      <div class="nav-item" @click="handleNavClick('category')">
        <el-icon><Grid /></el-icon>
        <span>分类</span>
      </div>
      <div class="nav-item" @click="handleNavClick('profile')">
        <el-icon><User /></el-icon>
        <span>我的</span>
      </div>
      <div class="nav-item" @click="handleNavClick('more')">
        <el-icon><More /></el-icon>
        <span>更多</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { ElMessage } from 'element-plus'
import { House, Grid, User, More, ArrowLeft, ArrowRight } from '@element-plus/icons-vue'
import { getCarouselImagePublic } from '@/api/carouse/carouselImage'

defineOptions({
  name: 'HomePage'
})

// 轮播图数据
const carouselImages = ref([
  {
    id: 1,
    image: 'https://picsum.photos/400/200?random=1',
    title: '防诈骗知识普及',
    url: '/guide/1'
  },
  {
    id: 2,
    image: 'https://picsum.photos/400/200?random=2',
    title: '网络安全防护',
    url: '/guide/2'
  },
  {
    id: 3,
    image: 'https://picsum.photos/400/200?random=3',
    title: '金融诈骗识别',
    url: '/guide/3'
  },
  {
    id: 4,
    image: 'https://picsum.photos/400/200?random=4',
    title: '电信诈骗预防',
    url: '/guide/4'
  },
  {
    id: 5,
    image: 'https://picsum.photos/400/200?random=5',
    title: '网购安全指南',
    url: '/guide/5'
  }
])

// 分类数据
const categories = ref([
  { key: 'selected', name: '精选', icon: '⭐' },
  { key: 'hot', name: '热门', icon: '🔥' },
  { key: 'entertainment', name: '娱乐', icon: '🎮' },
  { key: 'lifestyle', name: '生活', icon: '🏠' }
])

// 当前激活的分类
const activeCategory = ref('selected')

// 内容数据
const selectedContentList = ref([
  {
    id: 1,
    image: 'https://picsum.photos/80/80?random=4',
    title: '[Home] 许愿新头像馆',
    description: '精美头像，个性定制',
    tag: '女生头像',
    tagColor: '#e91e63',
    views: '1.0万',
    likes: '520',
    url: '/content/1'
  },
  {
    id: 2,
    image: 'https://picsum.photos/80/80?random=5',
    title: '[无忧] 超高颜值恋爱头像',
    description: '甜蜜恋爱风格头像',
    tag: '恋爱头像',
    tagColor: '#f368e0',
    views: '8.5千',
    likes: '365',
    url: '/content/2'
  },
  {
    id: 3,
    image: 'https://picsum.photos/80/80?random=6',
    title: '[IPO] 恋爱头像',
    description: '浪漫情侣专属',
    tag: '情侣头像',
    tagColor: '#ff6348',
    views: '6.2千',
    likes: '298',
    url: '/content/3'
  },
  {
    id: 4,
    image: 'https://picsum.photos/80/80?random=7',
    title: '[江南] 男生春梦高端娱乐厅',
    description: '高端娱乐体验',
    tag: '娱乐休闲',
    tagColor: '#2ed573',
    views: '4.8千',
    likes: '156',
    url: '/content/4'
  }
])

const hotContentList = ref([
  {
    id: 1,
    image: 'https://picsum.photos/80/80?random=10',
    title: '[热门] 今日最火话题',
    description: '全网都在讨论的热点内容',
    tag: '热门话题',
    tagColor: '#ff4757',
    views: '15.2万',
    likes: '8.9千',
    rating: '4.8',
    isHot: true,
    url: '/content/hot1'
  },
  {
    id: 2,
    image: 'https://picsum.photos/80/80?random=11',
    title: '[爆款] 网红推荐好物',
    description: '明星同款，限时优惠',
    tag: '好物推荐',
    tagColor: '#ff6b6b',
    views: '12.8万',
    likes: '6.7千',
    rating: '4.9',
    isHot: true,
    url: '/content/hot2'
  },
  {
    id: 3,
    image: 'https://picsum.photos/80/80?random=12',
    title: '[火爆] 热门游戏攻略',
    description: '最新游戏通关秘籍',
    tag: '游戏攻略',
    tagColor: '#ff9ff3',
    views: '9.5万',
    likes: '4.2千',
    rating: '4.7',
    url: '/content/hot3'
  },
  {
    id: 4,
    image: 'https://picsum.photos/80/80?random=13',
    title: '[热议] 社会热点分析',
    description: '深度解读当下热点',
    tag: '社会话题',
    tagColor: '#ffa502',
    views: '7.8万',
    likes: '3.1千',
    url: '/content/hot4'
  }
])

const entertainmentContentList = ref([
  {
    id: 1,
    image: 'https://picsum.photos/80/80?random=14',
    title: '[娱乐] 最新电影推荐',
    description: '院线热映，口碑佳作',
    tag: '电影推荐',
    tagColor: '#3742fa',
    views: '5.6万',
    likes: '2.8千',
    rating: '4.6',
    isNew: true,
    url: '/content/ent1'
  },
  {
    id: 2,
    image: 'https://picsum.photos/80/80?random=15',
    title: '[综艺] 爆笑综艺合集',
    description: '欢声笑语，解压必备',
    tag: '综艺节目',
    tagColor: '#ff3838',
    views: '4.2万',
    likes: '1.9千',
    rating: '4.5',
    url: '/content/ent2'
  },
  {
    id: 3,
    image: 'https://picsum.photos/80/80?random=16',
    title: '[音乐] 热门歌曲榜单',
    description: '流行金曲，循环播放',
    tag: '流行音乐',
    tagColor: '#ff6b35',
    views: '3.8万',
    likes: '1.5千',
    url: '/content/ent3'
  },
  {
    id: 4,
    image: 'https://picsum.photos/80/80?random=17',
    title: '[直播] 网红主播推荐',
    description: '颜值与才华并存',
    tag: '直播娱乐',
    tagColor: '#7bed9f',
    views: '2.9万',
    likes: '1.2千',
    isHot: true,
    url: '/content/ent4'
  }
])

const lifestyleContentList = ref([
  {
    id: 1,
    image: 'https://picsum.photos/80/80?random=18',
    title: '[美食] 家常菜制作教程',
    description: '简单易学，营养美味',
    tag: '美食制作',
    tagColor: '#ff7675',
    views: '3.2万',
    likes: '1.8千',
    rating: '4.7',
    url: '/content/life1'
  },
  {
    id: 2,
    image: 'https://picsum.photos/80/80?random=19',
    title: '[健康] 养生保健知识',
    description: '科学养生，健康生活',
    tag: '健康养生',
    tagColor: '#00b894',
    views: '2.8万',
    likes: '1.4千',
    url: '/content/life2'
  },
  {
    id: 3,
    image: 'https://picsum.photos/80/80?random=20',
    title: '[家居] 装修设计灵感',
    description: '温馨家居，品质生活',
    tag: '家居装修',
    tagColor: '#6c5ce7',
    views: '2.1万',
    likes: '980',
    isNew: true,
    url: '/content/life3'
  },
  {
    id: 4,
    image: 'https://picsum.photos/80/80?random=21',
    title: '[旅行] 周末出游攻略',
    description: '放松心情，拥抱自然',
    tag: '旅行攻略',
    tagColor: '#fd79a8',
    views: '1.9万',
    likes: '750',
    url: '/content/life4'
  }
])

// 轮播图控制
const currentIndex = ref(0)
const carouselContainer = ref(null)
let autoPlayTimer = null

// 触摸滑动相关
let startX = 0
let startY = 0
let isDragging = false
let startTime = 0

// 轮播图方法
const nextSlide = () => {
  if (carouselImages.value.length > 0) {
    currentIndex.value = (currentIndex.value + 1) % carouselImages.value.length
  }
}

const prevSlide = () => {
  if (carouselImages.value.length > 0) {
    currentIndex.value = currentIndex.value === 0 
      ? carouselImages.value.length - 1 
      : currentIndex.value - 1
  }
}

const goToSlide = (index) => {
  currentIndex.value = index
  resetAutoPlay()
}

const startAutoPlay = () => {
  if (carouselImages.value.length > 1) {
    autoPlayTimer = setInterval(() => {
      nextSlide()
    }, 4000)
  }
}

const stopAutoPlay = () => {
  if (autoPlayTimer) {
    clearInterval(autoPlayTimer)
    autoPlayTimer = null
  }
}

const resetAutoPlay = () => {
  stopAutoPlay()
  startAutoPlay()
}

// 触摸事件处理
const handleTouchStart = (e) => {
  startX = e.touches[0].clientX
  startY = e.touches[0].clientY
  startTime = Date.now()
  isDragging = true
  stopAutoPlay()
}

const handleTouchMove = (e) => {
  if (!isDragging) return
  
  const currentX = e.touches[0].clientX
  const currentY = e.touches[0].clientY
  const diffX = startX - currentX
  const diffY = startY - currentY
  
  // 如果是水平滑动，阻止默认行为
  if (Math.abs(diffX) > Math.abs(diffY)) {
    e.preventDefault()
  }
}

const handleTouchEnd = (e) => {
  if (!isDragging) return
  
  const endX = e.changedTouches[0].clientX
  const endY = e.changedTouches[0].clientY
  const diffX = startX - endX
  const diffY = startY - endY
  const diffTime = Date.now() - startTime
  
  // 判断是否为有效的水平滑动
  if (Math.abs(diffX) > Math.abs(diffY) && Math.abs(diffX) > 50 && diffTime < 300) {
    if (diffX > 0) {
      // 向左滑动，显示下一张
      nextSlide()
    } else {
      // 向右滑动，显示上一张
      prevSlide()
    }
  }
  
  isDragging = false
  resetAutoPlay()
}

// 事件处理函数
const handleCarouselClick = (item) => {
  ElMessage.success(`点击了轮播图: ${item.title}`)
}

const handleFeatureClick = (type) => {
  const messages = {
    match: '进入牵手速配',
    auction: '打开激情拍卖',
    romance: '进入浪漫满屋'
  }
  ElMessage.success(messages[type])
}

const handleContentClick = (item) => {
  ElMessage.success(`查看内容: ${item.title}`)
}

const handleNavClick = (nav) => {
  const messages = {
    home: '首页',
    category: '分类',
    profile: '我的',
    more: '更多'
  }
  ElMessage.success(`切换到: ${messages[nav]}`)
}

const switchCategory = (categoryKey) => {
  activeCategory.value = categoryKey
}

const getCurrentContentList = () => {
  switch (activeCategory.value) {
    case 'hot':
      return hotContentList.value
    case 'selected':
      return selectedContentList.value
    case 'entertainment':
      return entertainmentContentList.value
    case 'lifestyle':
      return lifestyleContentList.value
    default:
      return selectedContentList.value
  }
}

// 加载轮播图数据
const loadCarouselImages = async () => {
  try {
    const response = await getCarouselImagePublic()
    if (response.code === 0 && response.data && response.data.length > 0) {
      carouselImages.value = response.data
    }
  } catch (error) {
    console.error('加载轮播图失败:', error)
    // 使用默认数据
  }
}

// 生命周期
onMounted(() => {
  loadCarouselImages()
  startAutoPlay()
})

onUnmounted(() => {
  stopAutoPlay()
})
</script>

<style scoped>
.home-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  padding-bottom: 70px;
}

/* 轮播图样式 */
.carousel-section {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
  border-radius: 0 0 20px 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.carousel-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.carousel-wrapper {
  display: flex;
  width: 100%;
  height: 100%;
  transition: transform 0.3s ease-in-out;
  touch-action: pan-y;
}

.carousel-slide {
  flex: 0 0 100%;
  position: relative;
  cursor: pointer;
}

.carousel-slide img {
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
}

.carousel-overlay h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
}

.carousel-indicators {
  position: absolute;
  bottom: 15px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 8px;
  z-index: 10;
}

.indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.5);
  cursor: pointer;
  transition: all 0.3s ease;
}

.indicator.active {
  background: white;
  transform: scale(1.2);
}

/* 轮播图按钮 */
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
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.carousel-btn:hover {
  background: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  transform: translateY(-50%) scale(1.1);
}

.carousel-btn.prev {
  left: 15px;
}

.carousel-btn.next {
  right: 15px;
}

.carousel-btn .el-icon {
  font-size: 18px;
  color: #333;
}

/* 主标题 */
.main-header {
  text-align: center;
  padding: 20px;
  position: relative;
}

.main-title {
  font-size: 28px;
  font-weight: bold;
  color: white;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.leaf-decoration {
  position: absolute;
  top: 10px;
  right: 20px;
  font-size: 24px;
  transform: rotate(15deg);
}

/* 功能按钮区域 */
.feature-buttons {
  display: flex;
  justify-content: space-between;
  gap: 15px;
  margin: 0 20px 30px;
}

.feature-btn {
  flex: 1;
  background: white;
  border-radius: 15px;
  padding: 20px 10px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  position: relative;
  overflow: hidden;
}

.feature-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  transition: all 0.3s ease;
}

.feature-btn.purple::before {
  background: linear-gradient(90deg, #667eea, #764ba2);
}

.feature-btn.orange::before {
  background: linear-gradient(90deg, #f093fb, #f5576c);
}

.feature-btn.pink::before {
  background: linear-gradient(90deg, #4facfe, #00f2fe);
}

.feature-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.btn-icon {
  font-size: 24px;
  margin-bottom: 8px;
}

.feature-btn span {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #333;
}

/* 分类标签 */
.category-tabs {
  display: flex;
  justify-content: space-around;
  margin: 0 20px 25px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 25px;
  padding: 5px;
}

.category-tab {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px 8px;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: rgba(255, 255, 255, 0.7);
}

.category-tab.active {
  background: white;
  color: #333;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.tab-icon {
  font-size: 20px;
  margin-bottom: 4px;
}

.tab-text {
  font-size: 12px;
  font-weight: 500;
}

/* 内容列表 */
.content-list {
  background: white;
  border-radius: 20px 20px 0 0;
  padding: 20px;
  margin: 0 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.content-item {
  display: flex;
  align-items: center;
  padding: 15px;
  margin-bottom: 15px;
  background: #f8f9fa;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid #e9ecef;
}

.content-item:last-child {
  margin-bottom: 0;
}

.content-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  border-color: #dee2e6;
}

.content-image {
  position: relative;
  margin-right: 15px;
  flex-shrink: 0;
}

.content-image img {
  width: 60px;
  height: 60px;
  border-radius: 8px;
  object-fit: cover;
}

.hot-badge, .new-badge {
  position: absolute;
  top: -5px;
  right: -5px;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  font-weight: bold;
}

.hot-badge {
  background: #ff4757;
  color: white;
}

.new-badge {
  background: #2ed573;
  color: white;
}

.content-info {
  flex: 1;
  min-width: 0;
}

.content-title {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 8px 0;
  line-height: 1.3;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.content-description {
  font-size: 13px;
  color: #666;
  margin-bottom: 10px;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.content-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tag {
  font-size: 11px;
  color: white;
  padding: 3px 8px;
  border-radius: 12px;
  font-weight: 500;
}

.stats {
  display: flex;
  gap: 8px;
  font-size: 11px;
  color: #999;
}

/* 底部导航 */
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: white;
  display: flex;
  justify-content: space-around;
  padding: 10px 0;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 5px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #999;
}

.nav-item.active {
  color: #667eea;
}

.nav-item span {
  font-size: 12px;
  margin-top: 2px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .feature-buttons {
    flex-direction: column;
    gap: 10px;
  }
  
  .category-tabs {
    flex-wrap: wrap;
    gap: 5px;
  }
  
  .category-tab {
    min-width: 70px;
  }
  
  .content-item {
    flex-direction: column;
    text-align: center;
  }
  
  .content-image {
    margin-right: 0;
    margin-bottom: 10px;
  }
  
  .content-title,
  .content-description {
    white-space: normal;
  }
}

/* 动画效果 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.content-item {
  animation: fadeInUp 0.6s ease forwards;
}

.content-item:nth-child(1) { animation-delay: 0.1s; }
.content-item:nth-child(2) { animation-delay: 0.2s; }
.content-item:nth-child(3) { animation-delay: 0.3s; }
.content-item:nth-child(4) { animation-delay: 0.4s; }
</style>