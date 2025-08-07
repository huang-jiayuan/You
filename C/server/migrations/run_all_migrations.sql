-- 运行所有迁移脚本的主文件
-- 按照依赖关系顺序执行

-- 1. 首先创建基础表（无外键依赖）
SOURCE create_user_table.sql;

-- 2. 创建房间表
SOURCE create_room_table.sql;

-- 3. 创建依赖于用户和房间表的表
SOURCE create_room_member_table.sql;
SOURCE create_room_mic_table.sql;
SOURCE create_mic_application_table.sql;
SOURCE create_user_room_table.sql;

-- 显示创建的表
SHOW TABLES LIKE '%room%';
SHOW TABLES LIKE '%mic%';
SHOW TABLES LIKE '%user%';