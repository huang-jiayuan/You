<template>
  <div class="search-filter-debug">
    <h3>搜索和筛选调试工具</h3>
    
    <!-- 搜索测试 -->
    <div class="debug-section">
      <h4>搜索测试</h4>
      <input 
        v-model="testKeyword" 
        placeholder="输入搜索关键词" 
        @keyup.enter="testSearch"
      />
      <button @click="testSearch">搜索</button>
      <div v-if="searchResult" class="result">
        <pre>{{ JSON.stringify(searchResult, null, 2) }}</pre>
      </div>
    </div>
    
    <!-- 筛选测试 -->
    <div class="debug-section">
      <h4>标签筛选测试</h4>
      <select v-model="testTagId">
        <option value="0">全部</option>
        <option value="1">热门</option>
        <option value="2">娱乐</option>
        <option value="3">交友</option>
        <option value="4">才艺</option>
        <option value="5">电台音乐</option>
      </select>
      <button @click="testFilter">筛选</button>
      <div v-if="filterResult" class="result">
        <pre>{{ JSON.stringify(filterResult, null, 2) }}</pre>
      </div>
    </div>
    
    <!-- 推荐房间测试 -->
    <div class="debug-section">
      <h4>推荐房间测试</h4>
      <button @click="testRecommend">获取推荐房间</button>
      <div v-if="recommendResult" class="result">
        <pre>{{ JSON.stringify(recommendResult, null, 2) }}</pre>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import roomAPI from '../api/room.js'

export default {
  name: 'SearchFilterDebug',
  setup() {
    const testKeyword = ref('房间')
    const testTagId = ref('1')
    const searchResult = ref(null)
    const filterResult = ref(null)
    const recommendResult = ref(null)
    
    const testSearch = async () => {
      try {
        console.log('测试搜索，关键词:', testKeyword.value)
        const result = await roomAPI.searchRooms(testKeyword.value, 1, 10)
        searchResult.value = result
        console.log('搜索结果:', result)
      } catch (error) {
        console.error('搜索失败:', error)
        searchResult.value = { error: error.message }
      }
    }
    
    const testFilter = async () => {
      try {
        console.log('测试筛选，标签ID:', testTagId.value)
        const result = await roomAPI.getRoomsByCategory(parseInt(testTagId.value), 1, 10)
        filterResult.value = result
        console.log('筛选结果:', result)
      } catch (error) {
        console.error('筛选失败:', error)
        filterResult.value = { error: error.message }
      }
    }
    
    const testRecommend = async () => {
      try {
        console.log('测试推荐房间')
        const result = await roomAPI.getRecommendRooms(1, 10)
        recommendResult.value = result
        console.log('推荐房间结果:', result)
      } catch (error) {
        console.error('获取推荐房间失败:', error)
        recommendResult.value = { error: error.message }
      }
    }
    
    return {
      testKeyword,
      testTagId,
      searchResult,
      filterResult,
      recommendResult,
      testSearch,
      testFilter,
      testRecommend
    }
  }
}
</script>

<style scoped>
.search-filter-debug {
  position: fixed;
  top: 10px;
  right: 10px;
  width: 400px;
  background: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 9999;
  max-height: 80vh;
  overflow-y: auto;
}

.debug-section {
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid #eee;
}

.debug-section:last-child {
  border-bottom: none;
}

.debug-section h4 {
  margin: 0 0 10px 0;
  color: #333;
}

.debug-section input,
.debug-section select,
.debug-section button {
  margin: 4px;
  padding: 6px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.debug-section button {
  background: #007bff;
  color: white;
  cursor: pointer;
}

.debug-section button:hover {
  background: #0056b3;
}

.result {
  margin-top: 10px;
  padding: 10px;
  background: #f8f9fa;
  border-radius: 4px;
  font-size: 12px;
  max-height: 200px;
  overflow-y: auto;
}
</style>