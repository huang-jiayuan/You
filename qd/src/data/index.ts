/**
 * Mock Data Index
 * 模拟数据入口文件
 */

// 导出所有用户相关数据和方法
export * from './mockUsers'

// 导出所有消息相关数据和方法
export * from './mockMessages'

// 导出所有对话相关数据和方法
export * from './mockConversations'

// 导出所有快捷操作相关数据和方法
export * from './mockQuickActions'

// 导出所有导航标签相关数据和方法
export * from './mockNavigationTabs'

// 导出数据服务
export * from './mockDataService'

// 默认导出数据服务实例
export { mockDataService as default } from './mockDataService'