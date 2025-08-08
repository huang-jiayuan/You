package models

import (
	"server/room/basic/global"
	"time"
)

// 房间基础信息表
type Room struct {
	Id             int64     `gorm:"column:id;type:bigint;comment:房间唯一ID;primaryKey;not null;" json:"id"`                                      // 房间唯一ID
	RoomName       string    `gorm:"column:room_name;type:varchar(50);comment:房间名称;not null;" json:"room_name"`                                // 房间名称
	UserId         uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:房主用户ID，关联 user.id;not null;" json:"user_id"`                 // 房主用户ID，关联 user.id
	RoomType       string    `gorm:"column:room_type;type:varchar(1);comment:房间类型(0-公开房,1-私密房);not null;default:0;" json:"room_type"`    // 房间类型(0-公开房,1-私密房)
	Status         string    `gorm:"column:status;type:varchar(1);comment:房间状态(0-正常,1-已解散,2-全局禁言);not null;default:0;" json:"status"` // 房间状态(0-正常,1-已解散,2-全局禁言)
	Announcement   string    `gorm:"column:announcement;type:varchar(200);comment:房间公告;default:NULL;" json:"announcement"`                     // 房间公告
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	ClosedAt       time.Time `gorm:"column:closed_at;type:datetime;comment:解散时间;default:NULL;" json:"closed_at"`                                         // 解散时间
	LastActiveTime time.Time `gorm:"column:last_active_time;type:datetime;comment:最后活跃时间;not null;default:CURRENT_TIMESTAMP;" json:"last_active_time"` // 最后活跃时间
	FkMemberRoom   int32     `gorm:"column:fk_member_room;type:int;comment:房间人数;not null;default:0;" json:"fk_member_room"`                              // 房间人数
	Image          string    `gorm:"column:image;type:varchar(255);comment:房间封面图;default:NULL;" json:"image"`                                           // 房间封面图
	TagId          uint64    `gorm:"column:tag_id;type:bigint UNSIGNED;comment:关联的标签ID（一个房间只能选一个）;not null;" json:"tag_id"`                    // 关联的标签ID（一个房间只能选一个）
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy      int64     `gorm:"column:created_by;type:bigint;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy      int64     `gorm:"column:updated_by;type:bigint;comment:修改者;default:NULL;" json:"updated_by"` // 修改者
	DeletedBy      int64     `gorm:"column:deleted_by;type:bigint;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (r *Room) TableName() string {
	return "room"
}

// 房间标签表
type RoomTagDict struct {
	Id        uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:标签唯一ID;primaryKey;not null;" json:"id"`            // 标签唯一ID
	TagName   string    `gorm:"column:tag_name;type:varchar(30);comment:标签名称;not null;" json:"tag_name"`                 // 标签名称
	TagDesc   string    `gorm:"column:tag_desc;type:varchar(100);comment:标签描述;default:NULL;" json:"tag_desc"`            // 标签描述
	Status    string    `gorm:"column:status;type:varchar(1);comment:状态(0-启用,1-停用);not null;default:0;" json:"status"` // 状态(0-启用,1-停用)
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;default:NULL;" json:"created_at"`         // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;default:NULL;" json:"updated_at"`         // 修改时间
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;comment:删除时间;default:NULL;" json:"deleted_at"`            // 删除时间
	CreatedBy int64     `gorm:"column:created_by;type:bigint;comment:创建者;default:NULL;" json:"created_by"`                // 创建者
	UpdatedBy int64     `gorm:"column:updated_by;type:bigint;comment:更新者;default:NULL;" json:"updated_by"`                // 更新者
	DeletedBy int64     `gorm:"column:deleted_by;type:bigint;comment:删除者;default:NULL;" json:"deleted_by"`                // 删除者
}

func (r *RoomTagDict) TableName() string {
	return "room_tag_dict"
}

// 房间成员表
type RoomMember struct {
	Id          int64     `gorm:"column:id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"id"`                              // 记录唯一ID
	RoomId      int64     `gorm:"column:room_id;type:bigint;comment:关联房间ID;not null;" json:"room_id"`                               // 关联房间ID
	UserId      uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:成员用户ID，关联 user.id;not null;" json:"user_id"`         // 成员用户ID，关联 user.id
	Role        int8      `gorm:"column:role;type:tinyint;comment:成员角色(0-普通,1-管理员,2-房主);not null;default:0;" json:"role"`    // 成员角色(0-普通,1-管理员,2-房主)
	JoinTime    time.Time `gorm:"column:join_time;type:datetime;comment:加入时间;not null;default:CURRENT_TIMESTAMP;" json:"join_time"` // 加入时间
	LeaveTime   time.Time `gorm:"column:leave_time;type:datetime;comment:离开时间;default:NULL;" json:"leave_time"`                     // 离开时间
	IsMuted     int8      `gorm:"column:is_muted;type:tinyint;comment:是否禁言(0-正常,1-禁言);not null;default:0;" json:"is_muted"`     // 是否禁言(0-正常,1-禁言)
	MuteEndTime time.Time `gorm:"column:mute_end_time;type:datetime;comment:禁言结束时间;default:NULL;" json:"mute_end_time"`           // 禁言结束时间
}

func (r *RoomMember) TableName() string {
	return "room_member"
}

// 房间消息表
type RoomMessage struct {
	Id          int64     `gorm:"column:id;type:bigint;comment:消息唯一ID;primaryKey;not null;" json:"id"`                                                 // 消息唯一ID
	RoomId      int64     `gorm:"column:room_id;type:bigint;comment:关联房间ID;not null;" json:"room_id"`                                                  // 关联房间ID
	SenderId    uint64    `gorm:"column:sender_id;type:bigint UNSIGNED;comment:发送者用户ID，关联 user.id;not null;" json:"sender_id"`                      // 发送者用户ID，关联 user.id
	Content     string    `gorm:"column:content;type:varchar(500);comment:消息内容;not null;" json:"content"`                                              // 消息内容
	MessageType int8      `gorm:"column:message_type;type:tinyint;comment:消息类型(0-文字,1-礼物通知,2-系统通知);not null;default:0;" json:"message_type"` // 消息类型(0-文字,1-礼物通知,2-系统通知)
	SendTime    time.Time `gorm:"column:send_time;type:datetime;comment:发送时间;not null;default:CURRENT_TIMESTAMP;" json:"send_time"`                    // 发送时间
	IsDeleted   int8      `gorm:"column:is_deleted;type:tinyint;comment:是否删除(0-正常,1-删除);not null;default:0;" json:"is_deleted"`                    // 是否删除(0-正常,1-删除)
}

func (r *RoomMessage) TableName() string {
	return "room_message"
}

// 麦位管理表
type RoomMic struct {
	Id          int64     `gorm:"column:id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"id"`                                // 记录唯一ID
	RoomId      int64     `gorm:"column:room_id;type:bigint;comment:关联房间ID;not null;" json:"room_id"`                                 // 关联房间ID
	MicPosition int8      `gorm:"column:mic_position;type:tinyint;comment:麦位序号(1-8);not null;" json:"mic_position"`                   // 麦位序号(1-8)
	Status      int8      `gorm:"column:status;type:tinyint;comment:麦位状态(0-空闲,1-占用,2-禁用);not null;default:0;" json:"status"`    // 麦位状态(0-空闲,1-占用,2-禁用)
	UserId      uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:占用麦位的用户ID，关联 user.id;default:NULL;" json:"user_id"` // 占用麦位的用户ID，关联 user.id
	OccupyTime  time.Time `gorm:"column:occupy_time;type:datetime;comment:占用时间;default:NULL;" json:"occupy_time"`                     // 占用时间
	IsLocked    int8      `gorm:"column:is_locked;type:tinyint;comment:是否锁麦(0-未锁,1-已锁);not null;default:0;" json:"is_locked"`     // 是否锁麦(0-未锁,1-已锁)
}

func (r *RoomMic) TableName() string {
	return "room_mic"
}

// 麦位申请表
type MicApplication struct {
	Id         int64     `gorm:"column:id;type:bigint;comment:申请唯一ID;primaryKey;not null;" json:"id"`                                            // 申请唯一ID
	RoomId     int64     `gorm:"column:room_id;type:bigint;comment:关联房间ID;not null;" json:"room_id"`                                             // 关联房间ID
	UserId     uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:申请用户ID;not null;" json:"user_id"`                                    // 申请用户ID
	Status     int8      `gorm:"column:status;type:tinyint;comment:申请状态(0-待审批,1-已批准,2-已拒绝,3-已取消);not null;default:0;" json:"status"` // 申请状态(0-待审批,1-已批准,2-已拒绝,3-已取消)
	ApplyTime  time.Time `gorm:"column:apply_time;type:datetime;comment:申请时间;not null;default:CURRENT_TIMESTAMP;" json:"apply_time"`             // 申请时间
	HandleTime time.Time `gorm:"column:handle_time;type:datetime;comment:处理时间;default:NULL;" json:"handle_time"`                                 // 处理时间
	HandlerId  uint64    `gorm:"column:handler_id;type:bigint UNSIGNED;comment:处理者ID;default:NULL;" json:"handler_id"`                            // 处理者ID
	Reason     string    `gorm:"column:reason;type:varchar(200);comment:拒绝原因;default:NULL;" json:"reason"`                                       // 拒绝原因
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
}

func (m *MicApplication) TableName() string {
	return "mic_application"
}

// 用户礼物背包表
type UserGiftBackpack struct {
	Id           int64     `gorm:"column:id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"id"`                                  // 记录唯一ID
	UserId       uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:所属用户ID，关联 user.id;not null;" json:"user_id"`             // 所属用户ID，关联 user.id
	GiftId       int64     `gorm:"column:gift_id;type:bigint;comment:礼物ID;not null;" json:"gift_id"`                                       // 礼物ID
	Count        int32     `gorm:"column:count;type:int;comment:持有数量;not null;default:0;" json:"count"`                                  // 持有数量
	ObtainSource string    `gorm:"column:obtain_source;type:varchar(50);comment:获得来源;not null;" json:"obtain_source"`                    // 获得来源
	ObtainTime   time.Time `gorm:"column:obtain_time;type:datetime;comment:获得时间;not null;default:CURRENT_TIMESTAMP;" json:"obtain_time"` // 获得时间
	ExpireTime   time.Time `gorm:"column:expire_time;type:datetime;comment:过期时间(NULL-永久);default:NULL;" json:"expire_time"`            // 过期时间(NULL-永久)
	IsValid      int8      `gorm:"column:is_valid;type:tinyint(1);comment:是否有效(1-是,0-否);not null;default:1;" json:"is_valid"`          // 是否有效(1-是,0-否)
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"update_time"` // 更新时间
}

func (r *UserGiftBackpack) TableName() string {
	return "user_gift_backpack"
}

// 礼物基础信息表
type GiftInfo struct {
	GiftId         int64     `gorm:"column:gift_id;type:bigint;comment:礼物唯一ID;primaryKey;not null;" json:"gift_id"`                        // 礼物唯一ID
	GiftName       string    `gorm:"column:gift_name;type:varchar(100);comment:礼物名称;not null;" json:"gift_name"`                           // 礼物名称
	GiftType       int8      `gorm:"column:gift_type;type:tinyint;comment:礼物类型(1-普通,2-特效,3-稀有,4-自定义);not null;" json:"gift_type"` // 礼物类型(1-普通,2-特效,3-稀有,4-自定义)
	Price          float64   `gorm:"column:price;type:decimal(10, 2);comment:虚拟币价格;not null;" json:"price"`                               // 虚拟币价格
	DiamondPrice   int32     `gorm:"column:diamond_price;type:int;comment:钻石价格(可为空);default:NULL;" json:"diamond_price"`                // 钻石价格(可为空)
	IconUrl        string    `gorm:"column:icon_url;type:varchar(255);comment:礼物图标URL;not null;" json:"icon_url"`                          // 礼物图标URL
	AnimationUrl   string    `gorm:"column:animation_url;type:varchar(255);comment:动画URL(特效礼物);default:NULL;" json:"animation_url"`      // 动画URL(特效礼物)
	Weight         int32     `gorm:"column:weight;type:int;comment:排序权重;not null;default:0;" json:"weight"`                                // 排序权重
	IsHot          int8      `gorm:"column:is_hot;type:tinyint(1);comment:是否热门(1-是,0-否);not null;default:0;" json:"is_hot"`              // 是否热门(1-是,0-否)
	IsLimit        int8      `gorm:"column:is_limit;type:tinyint(1);comment:是否限时(1-是,0-否);not null;default:0;" json:"is_limit"`          // 是否限时(1-是,0-否)
	LimitStartTime time.Time `gorm:"column:limit_start_time;type:datetime;comment:限时开始时间;default:NULL;" json:"limit_start_time"`         // 限时开始时间
	LimitEndTime   time.Time `gorm:"column:limit_end_time;type:datetime;comment:限时结束时间;default:NULL;" json:"limit_end_time"`             // 限时结束时间
	Status         int8      `gorm:"column:status;type:tinyint;comment:状态(0-下架,1-上架,2-待审核);not null;default:1;" json:"status"`        // 状态(0-下架,1-上架,2-待审核)
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"create_time"` // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"update_time"` // 更新时间
}

func (r *GiftInfo) TableName() string {
	return "gift_info"
}
func (r *GiftInfo) GetFindGiftById(id int64) error {
	return global.DB.Where("id=?", id).Find(&r).Error
}

// 礼物赠送记录表
type GiftSendRecord struct {
	RecordId      int64     `gorm:"column:record_id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"record_id"`                          // 记录唯一ID
	SendUserId    uint64    `gorm:"column:send_user_id;type:bigint UNSIGNED;comment:赠送者用户ID，关联 user.id;not null;" json:"send_user_id"`       // 赠送者用户ID，关联 user.id
	ReceiveUserId uint64    `gorm:"column:receive_user_id;type:bigint UNSIGNED;comment:接收者用户ID，关联 user.id;not null;" json:"receive_user_id"` // 接收者用户ID，关联 user.id
	RoomId        int64     `gorm:"column:room_id;type:bigint;comment:房间ID(NULL-私聊);default:NULL;" json:"room_id"`                              // 房间ID(NULL-私聊)
	GiftId        int64     `gorm:"column:gift_id;type:bigint;comment:礼物ID;not null;" json:"gift_id"`                                             // 礼物ID
	SendCount     int32     `gorm:"column:send_count;type:int;comment:赠送数量;not null;default:1;" json:"send_count"`                              // 赠送数量
	TotalPrice    float64   `gorm:"column:total_price;type:decimal(10, 2);comment:总消耗(虚拟币);not null;" json:"total_price"`                     // 总消耗(虚拟币)
	TotalDiamond  int32     `gorm:"column:total_diamond;type:int;comment:总消耗(钻石);default:0;" json:"total_diamond"`                             // 总消耗(钻石)
	SendType      int8      `gorm:"column:send_type;type:tinyint;comment:赠送方式(1-背包,2-直接购买);not null;" json:"send_type"`                   // 赠送方式(1-背包,2-直接购买)
	Message       string    `gorm:"column:message;type:varchar(255);comment:赠送附言;default:NULL;" json:"message"`                                 // 赠送附言
	SendTime      time.Time `gorm:"column:send_time;type:datetime;comment:赠送时间;not null;default:CURRENT_TIMESTAMP;" json:"send_time"`           // 赠送时间
	Status        int8      `gorm:"column:status;type:tinyint;comment:状态(0-失败,1-成功,2-已撤回);not null;default:1;" json:"status"`              // 状态(0-失败,1-成功,2-已撤回)
	ClientIp      string    `gorm:"column:client_ip;type:varchar(50);comment:赠送者IP;not null;" json:"client_ip"`                                  // 赠送者IP
}

func (r *GiftSendRecord) TableName() string {
	return "gift_send_record"
}

// 用户房间关系表 - 记录用户在线状态
type UserRoom struct {
	Id        int64     `gorm:"column:id;type:bigint;comment:记录唯一ID;primaryKey;not null;" json:"id"`                              // 记录唯一ID
	UserId    uint64    `gorm:"column:user_id;type:bigint UNSIGNED;comment:用户ID，关联 user.id;not null;" json:"user_id"`             // 用户ID，关联 user.id
	RoomId    int64     `gorm:"column:room_id;type:bigint;comment:房间ID，关联 room.id;not null;" json:"room_id"`                      // 房间ID，关联 room.id
	IsOnline  int8      `gorm:"column:is_online;type:tinyint;comment:是否在线(0-离线,1-在线);not null;default:0;" json:"is_online"`   // 是否在线(0-离线,1-在线)
	JoinTime  time.Time `gorm:"column:join_time;type:datetime;comment:进入时间;not null;default:CURRENT_TIMESTAMP;" json:"join_time"` // 进入时间
	LeaveTime time.Time `gorm:"column:leave_time;type:datetime;comment:离开时间;default:NULL;" json:"leave_time"`                     // 离开时间
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
}

func (r *UserRoom) TableName() string {
	return "user_room"
}
