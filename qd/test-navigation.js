#!/usr/bin/env node

/**
 * 导航功能测试运行脚本
 * 运行所有与聊天导航相关的测试
 */

import { spawn } from 'child_process'
import { fileURLToPath } from 'url'
import { dirname, join } from 'path'

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

// 测试文件模式
const testPatterns = [
  'src/views/__tests__/Chat.test.js',
  'src/views/__tests__/ChatDetail.test.js',
  'src/composables/__tests__/useNavigationError.test.js',
  'src/__tests__/navigation-integration.test.js'
]

console.log('🚀 开始运行聊天导航功能测试...\n')

// 运行测试
const testProcess = spawn('npx', ['vitest', 'run', ...testPatterns], {
  cwd: __dirname,
  stdio: 'inherit',
  shell: true
})

testProcess.on('close', (code) => {
  if (code === 0) {
    console.log('\n✅ 所有导航功能测试通过！')
    console.log('\n测试覆盖范围：')
    console.log('  - Chat.vue 组件测试')
    console.log('  - ChatDetail.vue 组件测试')
    console.log('  - useNavigationError 组合式函数测试')
    console.log('  - 导航功能集成测试')
  } else {
    console.log('\n❌ 测试失败，请检查错误信息')
    process.exit(code)
  }
})

testProcess.on('error', (error) => {
  console.error('❌ 测试运行失败:', error.message)
  process.exit(1)
})