# 房间管理功能使用说明

## 功能概述

本模块实现了房间管理的核心功能：
- **踢人功能**：将指定用户移出房间，被踢用户10分钟内不可重新进入
- **禁言功能**：禁止指定用户发言，支持1小时/24小时/永久三种时长
- **解除禁言**：手动解除用户的禁言状态
- **状态检查**：查询用户的踢出和禁言状态

## API接口

### 1. 踢人接口
**POST** `/api/room/kick`

请求参数：
```json
{
  "room_id": 123,
  "user_id": 456,
  "operator_id": 789,
  "reason": "违反房间规则"
}
```

响应：
```json
{
  "code": 200,
  "msg": "踢出用户成功",
  "data": {
    "success": true,
    "message": "用户已被踢出房间，10分钟内不可重新进入"
  }
}
```

### 2. 禁言接口
**POST** `/api/room/mute`

请求参数：
```json
{
  "room_id": 123,
  "user_id": 456,
  "operator_id": 789,
  "duration_type": 1,
  "reason": "发送不当言论"
}
```

duration_type说明：
- 1: 1小时
- 2: 24小时  
- 3: 永久

响应：
```json
{
  "code": 200,
  "msg": "禁言用户成功",
  "data": {
    "success": true,
    "message": "用户已被禁言1小时"
  }
}
```

### 3. 解除禁言接口
**POST** `/api/room/unmute`

请求参数：
```json
{
  "room_id": 123,
  "user_id": 456,
  "operator_id": 789
}
```

响应：
```json
{
  "code": 200,
  "msg": "解除禁言成功",
  "data": {
    "success": true,
    "message": "用户禁言已解除"
  }
}
```

### 4. 检查用户状态接口
**POST** `/api/room/status`

请求参数：
```json
{
  "room_id": 123,
  "user_id": 456
}
```

响应：
```json
{
  "code": 200,
  "msg": "获取用户状态成功",
  "data": {
    "is_kicked": false,
    "is_muted": true,
    "kick_expire_time": 0,
    "mute_expire_time": 1703123456,
    "message": "用户已被禁言"
  }
}
```

## 数据库表结构

### 踢人记录表 (room_kick_records)
```sql
CREATE TABLE room_kick_records (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  room_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  operator_id BIGINT NOT NULL,
  reason VARCHAR(500),
  kick_time DATETIME NOT NULL,
  expire_time DATETIME NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_room_user (room_id, user_id),
  INDEX idx_expire_time (expire_time)
);
```

### 禁言记录表 (room_mute_records)
```sql
CREATE TABLE room_mute_records (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  room_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  operator_id BIGINT NOT NULL,
  duration_type INT NOT NULL,
  reason VARCHAR(500),
  mute_time DATETIME NOT NULL,
  expire_time DATETIME NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_room_user (room_id, user_id),
  INDEX idx_expire_time (expire_time),
  INDEX idx_active (is_active)
);
```

## 辅助函数

### 检查用户是否可以进入房间
```go
canEnter, message := models.CanUserEnterRoom(roomID, userID)
if !canEnter {
    // 拒绝用户进入，返回message中的原因
}
```

### 检查用户是否可以发言
```go
canSpeak, message := models.CanUserSpeak(roomID, userID)
if !canSpeak {
    // 拒绝用户发言，返回message中的原因
}
```

## 使用注意事项

1. **权限验证**：在实际使用中，需要在API层添加权限验证，确保只有房间管理员才能执行踢人和禁言操作。

2. **实时通知**：踢人和禁言操作后，建议通过WebSocket等方式实时通知房间内的其他用户。

3. **日志记录**：所有管理操作都会记录在数据库中，便于后续审计和追踪。

4. **自动清理**：建议定期清理过期的踢人和禁言记录，避免数据库膨胀。

5. **业务集成**：在用户进入房间和发送消息时，需要调用相应的检查函数验证用户状态。

## 错误处理

所有接口都包含完整的错误处理，常见错误包括：
- 参数验证失败
- 数据库操作失败
- 用户状态冲突（如重复踢人、禁言等）
- 权限不足（操作自己等）

错误信息会在响应的message字段中详细说明。