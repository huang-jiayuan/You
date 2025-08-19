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
    host: '127.0.0.1', // 强制使用IPv4
    open: true,
    // 代理配置，用于开发环境跨域
    proxy: {
      // room 相关接口代理到 8080 端口
      '/api/room': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '/api'), // 保持/api前缀
        configure: (proxy, options) => {
          proxy.on('error', (err, req, res) => {
            console.log('Room service proxy error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('Sending Request to Room Service:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('Received Response from Room Service:', proxyRes.statusCode, req.url);
          });
        }
      },
      // user 和 auth 相关接口代理到 8081 端口
      '/api/user': {
        target: 'http://127.0.0.1:8081',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '/api'), // 保持/api前缀
        configure: (proxy, options) => {
          proxy.on('error', (err, req, res) => {
            console.log('User service proxy error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('Sending Request to User Service:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('Received Response from User Service:', proxyRes.statusCode, req.url);
          });
        }
      },
      // auth 相关接口也代理到 8081 端口
      '/api/auth': {
        target: 'http://127.0.0.1:8081',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '/api'), // 保持/api前缀
        configure: (proxy, options) => {
          proxy.on('error', (err, req, res) => {
            console.log('Auth service proxy error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('Sending Request to Auth Service:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('Received Response from Auth Service:', proxyRes.statusCode, req.url);
          });
        }
      },
      // 其他 API 请求默认代理到 user 服务（保持向后兼容）
      '/api': {
        target: 'http://127.0.0.1:8081',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '/api'), // 保持/api前缀
        configure: (proxy, options) => {
          proxy.on('error', (err, req, res) => {
            console.log('Default API proxy error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, res) => {
            console.log('Sending Request to Default API:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, res) => {
            console.log('Received Response from Default API:', proxyRes.statusCode, req.url);
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