import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  esbuild: {
    target: 'es2020'
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    },
    extensions: ['.js', '.ts', '.jsx', '.tsx', '.json', '.vue']
  },
  test: {
    globals: true,
    environment: 'jsdom'
  },
  server: {
    port: 3000,
    host: '127.0.0.1',
    proxy: {
      // 房间相关API代理到8080端口
      '/api/room': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false
      },
      // 其他API（用户、认证等）代理到8081端口
      '/api': {
        target: 'http://localhost:8081',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
