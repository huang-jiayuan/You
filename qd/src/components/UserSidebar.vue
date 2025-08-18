<template>
  <div class="user-sidebar-overlay" :class="{ 'show': isVisible }" @click="closeSidebar">
    <div class="user-sidebar" :class="{ 'slide-in': isVisible }" @click.stop>
      <!-- 用户信息头部 -->
      <div class="user-header">
        <div class="user-avatar-large" @click="goToUserProfile">
          <img :src="userInfo.avatar" :alt="userInfo.nickname" />
        </div>
        <div class="user-basic-info">
          <h2 class="user-nickname">{{ userInfo.nickname }}</h2>
          <p class="user-id">ME号：{{ userInfo.meId }}</p>
          <div class="user-badges">
            <span class="badge badge-new">初来乍到</span>
            <span class="badge badge-level">{{ userInfo.level }}</span>
          </div>
        </div>
        <button class="close-btn" @click="closeSidebar">×</button>
      </div>

      <!-- 关注和粉丝 -->
      <div class="follow-stats">
        <div class="stat-item" @click="showFollowing">
          <span class="stat-label">关注</span>
          <span class="stat-count">{{ userInfo.following }}</span>
        </div>
        <div class="stat-item" @click="showFollowers">
          <span class="stat-label">粉丝</span>
          <span class="stat-count">{{ userInfo.followers }}</span>
        </div>
      </div>

      <!-- 功能菜单 -->
      <div class="menu-section">
        <div class="menu-item" @click="goToWallet">
          <div class="menu-icon">💰</div>
          <span class="menu-text">我的钱包</span>
          <div class="menu-value">{{ userInfo.coins }}</div>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToPersonalCenter">
          <div class="menu-icon">👤</div>
          <span class="menu-text">个人中心</span>
          <div class="menu-value">¥ {{ userInfo.balance }}</div>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToMemberCenter">
          <div class="menu-icon">💎</div>
          <span class="menu-text">会员中心</span>
          <div class="menu-value">普通用户</div>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToFamilyCenter">
          <div class="menu-icon">👥</div>
          <span class="menu-text">贵族中心</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToCompetitionCenter">
          <div class="menu-icon">🏆</div>
          <span class="menu-text">竞猜中心</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToListenCenter">
          <div class="menu-icon">🎧</div>
          <span class="menu-text">兑换中心</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToTeacher">
          <div class="menu-icon">👨‍🏫</div>
          <span class="menu-text">师徒</span>
          <div class="menu-stats">
            <span class="stat-blue">👥 {{ userInfo.teacherStats.disciples }}</span>
            <span class="stat-red">❤️ {{ userInfo.teacherStats.hearts }}</span>
          </div>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToBackpack">
          <div class="menu-icon">🎒</div>
          <span class="menu-text">背包管理</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToMall">
          <div class="menu-icon">🏪</div>
          <span class="menu-text">商城</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToSettings">
          <div class="menu-icon">⚙️</div>
          <span class="menu-text">设置</span>
          <div class="menu-arrow">></div>
        </div>

        <div class="menu-item" @click="goToCustomerService">
          <div class="menu-icon">🎧</div>
          <span class="menu-text">在线客服</span>
          <div class="menu-arrow">></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '@/composables/useToast'

export default {
  name: 'UserSidebar',
  props: {
    isVisible: {
      type: Boolean,
      default: false
    },
    userInfo: {
      type: Object,
      default: () => ({
        nickname: '途场',
        meId: '201691465',
        avatar: 'https://via.placeholder.com/80x80/4CAF50/ffffff?text=途',
        level: 0,
        following: 1,
        followers: 1,
        coins: 0,
        balance: '0.00',
        teacherStats: {
          disciples: 1,
          hearts: 1
        }
      })
    }
  },
  emits: ['close'],
  setup(props, { emit }) {
    const router = useRouter()
    const { info: showInfo } = useToast()

    const closeSidebar = () => {
      emit('close')
    }

    const showFollowing = () => {
      showInfo('关注列表功能开发中')
    }

    const showFollowers = () => {
      showInfo('粉丝列表功能开发中')
    }

    const goToWallet = () => {
      showInfo('我的钱包功能开发中')
      closeSidebar()
    }

    const goToPersonalCenter = () => {
      showInfo('个人中心功能开发中')
      closeSidebar()
    }

    const goToMemberCenter = () => {
      showInfo('会员中心功能开发中')
      closeSidebar()
    }

    const goToFamilyCenter = () => {
      showInfo('贵族中心功能开发中')
      closeSidebar()
    }

    const goToCompetitionCenter = () => {
      showInfo('竞猜中心功能开发中')
      closeSidebar()
    }

    const goToListenCenter = () => {
      showInfo('兑换中心功能开发中')
      closeSidebar()
    }

    const goToTeacher = () => {
      showInfo('师徒功能开发中')
      closeSidebar()
    }

    const goToBackpack = () => {
      showInfo('背包管理功能开发中')
      closeSidebar()
    }

    const goToMall = () => {
      showInfo('商城功能开发中')
      closeSidebar()
    }

    const goToSettings = () => {
      showInfo('设置功能开发中')
      closeSidebar()
    }

    const goToCustomerService = () => {
      showInfo('在线客服功能开发中')
      closeSidebar()
    }

    const goToUserProfile = () => {
      router.push('/profile/me')
      closeSidebar()
    }

    return {
      closeSidebar,
      showFollowing,
      showFollowers,
      goToWallet,
      goToPersonalCenter,
      goToMemberCenter,
      goToFamilyCenter,
      goToCompetitionCenter,
      goToListenCenter,
      goToTeacher,
      goToBackpack,
      goToMall,
      goToSettings,
      goToCustomerService,
      goToUserProfile
    }
  }
}
</script>

<style scoped>
.user-sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0);
  z-index: 1000;
  visibility: hidden;
  transition: all 0.3s ease;
}

.user-sidebar-overlay.show {
  background: rgba(0, 0, 0, 0.5);
  visibility: visible;
}

.user-sidebar {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 320px;
  max-width: 85vw;
  background: white;
  transform: translateX(-100%);
  transition: transform 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  overflow-y: auto;
  box-shadow: 2px 0 10px rgba(0, 0, 0, 0.1);
}

.user-sidebar.slide-in {
  transform: translateX(0);
}

/* 用户信息头部 */
.user-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 40px 20px 20px;
  position: relative;
  display: flex;
  align-items: center;
  gap: 15px;
}

.user-avatar-large {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  overflow: hidden;
  border: 3px solid rgba(255, 255, 255, 0.3);
  flex-shrink: 0;
  cursor: pointer;
  transition: all 0.2s ease;
  position: relative;
}

.user-avatar-large:hover {
  transform: scale(1.05);
  border-color: rgba(255, 255, 255, 0.6);
}

.user-avatar-large:active {
  transform: scale(0.95);
}

.user-avatar-large::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0);
  border-radius: 50%;
  transition: background-color 0.2s ease;
}

.user-avatar-large:hover::after {
  background: rgba(255, 255, 255, 0.1);
}

.user-avatar-large img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-basic-info {
  flex: 1;
}

.user-nickname {
  font-size: 20px;
  font-weight: 600;
  margin: 0 0 5px 0;
}

.user-id {
  font-size: 14px;
  opacity: 0.9;
  margin: 0 0 10px 0;
}

.user-badges {
  display: flex;
  gap: 8px;
}

.badge {
  padding: 3px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.badge-new {
  background: rgba(255, 107, 157, 0.2);
  color: #ff6b9d;
  border: 1px solid rgba(255, 107, 157, 0.3);
}

.badge-level {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.close-btn {
  position: absolute;
  top: 15px;
  right: 15px;
  background: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
  font-size: 24px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s ease;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

/* 关注和粉丝统计 */
.follow-stats {
  display: flex;
  background: white;
  border-bottom: 1px solid #f0f0f0;
}

.stat-item {
  flex: 1;
  padding: 20px;
  text-align: center;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.stat-item:hover {
  background: #f8f9ff;
}

.stat-item:first-child {
  border-right: 1px solid #f0f0f0;
}

.stat-label {
  display: block;
  font-size: 14px;
  color: #666;
  margin-bottom: 5px;
}

.stat-count {
  display: block;
  font-size: 18px;
  font-weight: 600;
  color: #333;
}

/* 功能菜单 */
.menu-section {
  background: white;
}

.menu-item {
  display: flex;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #f8f8f8;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.menu-item:hover {
  background: #f8f9ff;
}

.menu-item:last-child {
  border-bottom: none;
}

.menu-icon {
  font-size: 20px;
  width: 24px;
  text-align: center;
  margin-right: 15px;
}

.menu-text {
  flex: 1;
  font-size: 16px;
  color: #333;
}

.menu-value {
  font-size: 14px;
  color: #666;
  margin-right: 10px;
}

.menu-stats {
  display: flex;
  gap: 10px;
  margin-right: 10px;
}

.stat-blue {
  color: #4facfe;
  font-size: 12px;
}

.stat-red {
  color: #ff6b9d;
  font-size: 12px;
}

.menu-arrow {
  font-size: 14px;
  color: #ccc;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .user-sidebar {
    width: 280px;
  }
  
  .user-header {
    padding: 30px 15px 15px;
  }
  
  .user-avatar-large {
    width: 60px;
    height: 60px;
  }
  
  .user-nickname {
    font-size: 18px;
  }
  
  .user-id {
    font-size: 13px;
  }
  
  .stat-item {
    padding: 15px;
  }
  
  .menu-item {
    padding: 14px 15px;
  }
  
  .menu-text {
    font-size: 15px;
  }
}

/* 动画优化 */
@media (prefers-reduced-motion: reduce) {
  .user-sidebar-overlay,
  .user-sidebar,
  .menu-item,
  .stat-item,
  .close-btn {
    transition: none;
  }
}

/* 触摸设备优化 */
@media (hover: none) and (pointer: coarse) {
  .menu-item:hover,
  .stat-item:hover,
  .close-btn:hover {
    background: initial;
  }
  
  .menu-item:active {
    background: #f0f0f0;
  }
  
  .stat-item:active {
    background: #f0f0f0;
  }
}
</style>