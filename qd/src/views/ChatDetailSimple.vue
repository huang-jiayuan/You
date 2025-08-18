<template>
  <div class="chat-detail-simple">
    <div class="header">
      <button @click="goBack">← 返回</button>
      <h2>{{ userName }}</h2>
    </div>
    
    <div class="content">
      <p>聊天详情页面</p>
      <p>用户ID: {{ userId }}</p>
      <p>用户名: {{ userName }}</p>
    </div>
  </div>
</template>

<script>
import { useRouter, useRoute } from 'vue-router'

export default {
  name: 'ChatDetailSimple',
  setup() {
    const router = useRouter()
    const route = useRoute()
    
    const userId = route.params.userId
    const userName = route.query.name || '用户'
    
    const goBack = () => {
      router.go(-1)
    }
    
    return {
      userId,
      userName,
      goBack
    }
  }
}
</script>

<style scoped>
.chat-detail-simple {
  padding: 20px;
}

.header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
}

.header button {
  padding: 8px 16px;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.content {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
}
</style>