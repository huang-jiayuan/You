// 简单的 Node.js 测试脚本，用于验证 room.js 的语法
const fs = require('fs');
const path = require('path');

try {
  // 读取 room.js 文件内容
  const roomJsPath = path.join(__dirname, 'src/api/room.js');
  const content = fs.readFileSync(roomJsPath, 'utf8');
  
  console.log('✅ room.js 文件读取成功');
  console.log('文件大小:', content.length, '字符');
  
  // 检查是否包含必要的函数
  const requiredFunctions = [
    'transformRoomData',
    'generateDefaultCover', 
    'formatUserCount',
    'hashCode',
    'getRecommendRooms',
    'getRoomsByCategory',
    'searchRooms'
  ];
  
  const missingFunctions = requiredFunctions.filter(func => !content.includes(func));
  
  if (missingFunctions.length === 0) {
    console.log('✅ 所有必要的函数都存在');
  } else {
    console.log('❌ 缺少函数:', missingFunctions);
  }
  
  // 检查语法错误（简单检查）
  const syntaxIssues = [];
  
  if (content.includes('this.transformRoomData')) {
    syntaxIssues.push('仍然存在 this.transformRoomData 调用');
  }
  
  if (content.includes('this.generateDefaultCover')) {
    syntaxIssues.push('仍然存在 this.generateDefaultCover 调用');
  }
  
  if (content.includes('this.formatUserCount')) {
    syntaxIssues.push('仍然存在 this.formatUserCount 调用');
  }
  
  if (content.includes('this.hashCode')) {
    syntaxIssues.push('仍然存在 this.hashCode 调用');
  }
  
  if (syntaxIssues.length === 0) {
    console.log('✅ 没有发现明显的语法问题');
  } else {
    console.log('❌ 发现潜在问题:');
    syntaxIssues.forEach(issue => console.log('  -', issue));
  }
  
} catch (error) {
  console.error('❌ 测试失败:', error.message);
}