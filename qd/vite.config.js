import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  test: {
    globals: true,
    environment: 'jsdom'
  },
  server: {
    port: 3000,
    host: '127.0.0.1', // 强制使用IPv4
    open: true,
    // 代理配置，用于开发环境跨域
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8081', // 使用IPv4地址避免IPv6问题
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '/api'), // 保持/api前缀
        configure: (proxy, options) => {
          proxy.on('error', (err, req, res) => {
            console.log('proxy error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('Sending Request to the Target:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('Received Response from the Target:', proxyRes.statusCode, req.url);
          });
        }
      }
    }
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
    // 构建优化
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['vue', 'vue-router'],
          api: ['./src/api/request.js', './src/api/auth.js', './src/api/user.js']
        }
      }
    }
  }
})