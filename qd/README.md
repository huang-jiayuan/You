# Vue3 前端应用

一个基于Vue3的现代前端应用模板，展示了Vue3的核心特性和最佳实践。

## 技术栈

- **Vue 3.x** - 渐进式JavaScript框架
- **Vite 5.x** - 下一代前端构建工具
- **Vue Router 4.x** - Vue.js官方路由管理器
- **Vitest** - 基于Vite的单元测试框架

## 项目结构

```
vue3-frontend-app/
├── public/                 # 静态资源
├── src/
│   ├── assets/            # 资源文件
│   │   ├── images/        # 图片资源
│   │   └── styles/        # 样式文件
│   ├── components/        # Vue组件
│   │   ├── common/        # 通用组件
│   │   └── ui/           # UI组件
│   ├── views/            # 页面组件
│   ├── composables/      # 组合式函数
│   ├── router/           # 路由配置
│   ├── stores/           # 状态管理
│   ├── App.vue           # 根组件
│   └── main.js           # 应用入口
├── package.json          # 项目配置
├── vite.config.js        # Vite配置
└── README.md            # 项目说明
```

## 开发指南

### 安装依赖

```bash
npm install
```

### 启动开发服务器

```bash
npm run dev
```

### 构建生产版本

```bash
npm run build
```

### 运行测试

```bash
npm run test
```

### 预览生产构建

```bash
npm run preview
```

## 特性

- ✅ Vue3 Composition API
- ✅ 响应式设计
- ✅ 组件化架构
- ✅ 现代CSS设计系统
- ✅ 主题切换支持
- ✅ 单元测试配置
- ✅ 开发热重载

## 浏览器支持

- Chrome >= 87
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 许可证

MIT License