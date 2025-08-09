package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源的连接，生产环境中应限制具体来源
		return true
	},
}

// User 表示一个用户
type User struct {
	ID    int64
	Conn  *websocket.Conn // 仅在线用户有值
	Mutex sync.Mutex
}

// Message 表示WebSocket传输的消息结构
type Message struct {
	Type       string `json:"type"` // "text", "image", "error", "message_sent", "read_ack", "system"
	SenderID   int64  `json:"sender_id"`
	ReceiverID int64  `json:"receiver_id"`
	Content    string `json:"content"` // 文本内容或图片URL/base64
	Timestamp  int64  `json:"timestamp"`
	MessageID  string `json:"message_id"` // 消息唯一标识
	IsRead     bool   `json:"is_read"`    // 消息是否已读
}

// 在线用户连接存储
var onlineUsers = struct {
	sync.RWMutex
	m map[int64]*User
}{m: make(map[int64]*User)}

// 离线消息存储，key为接收者ID，value为消息列表
var offlineMessages = struct {
	sync.RWMutex
	m map[int64][]Message
}{m: make(map[int64][]Message)}

// 未读消息存储
var unreadMessages = struct {
	sync.RWMutex
	m map[int64]int // key: 用户ID, value: 未读消息数
}{m: make(map[int64]int)}

func HandleWebSocket(c *gin.Context) {
	// 获取用户ID，实际应用中应从认证系统获取
	userID := int64(c.GetUint("userId"))
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId参数是必需的"})
		return
	}

	// 将HTTP连接升级为WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升级连接失败:", err)
		return
	}
	defer conn.Close()

	// 创建用户并添加到在线用户列表
	user := &User{
		ID:   userID,
		Conn: conn,
	}

	onlineUsers.Lock()
	onlineUsers.m[userID] = user
	onlineUsers.Unlock()

	log.Printf("用户 %d 已连接", userID)

	// 用户上线后，发送所有离线消息
	sendOfflineMessages(user)

	// 连接关闭时从在线列表移除用户
	defer func() {
		onlineUsers.Lock()
		delete(onlineUsers.m, userID)
		onlineUsers.Unlock()
		log.Printf("用户 %d 已断开连接", userID)
	}()

	// 循环读取并处理消息
	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("用户 %d 读取消息错误: %v", userID, err)
			break
		}

		var msg Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			log.Printf("解析消息错误: %v", err)
			continue
		}

		// 确保发送者ID正确
		msg.SenderID = userID
		msg.Timestamp = time.Now().Unix()

		// 根据消息类型处理
		switch msg.Type {
		case "text", "image":
			// 生成消息ID
			msg.MessageID = uuid.New().String()
			msg.IsRead = false
			handleMessage(user, &msg)
		case "read_ack":
			handleReadAck(&msg)
		default:
			log.Printf("未知消息类型: %s", msg.Type)
		}
	}
}

// handleMessage 处理消息发送逻辑
func handleMessage(sender *User, msg *Message) {
	if msg.ReceiverID == 0 {
		log.Println("接收者ID不能为空")
		return
	}

	// 检查接收者是否在线
	onlineUsers.RLock()
	receiver, exists := onlineUsers.m[msg.ReceiverID]
	onlineUsers.RUnlock()

	if exists {
		// 接收者在线，直接发送消息
		if err := sendMessageToUser(receiver.Conn, msg); err != nil {
			log.Printf("发送消息失败: %v", err)
			sendError(sender.Conn, "发送消息失败")
			return
		}

		// 增加接收者的未读消息计数
		unreadMessages.Lock()
		unreadMessages.m[msg.ReceiverID]++
		unreadMessages.Unlock()
	} else {
		// 接收者离线，存储消息
		offlineMessages.Lock()
		offlineMessages.m[msg.ReceiverID] = append(offlineMessages.m[msg.ReceiverID], *msg)
		offlineMessages.Unlock()

		// 增加接收者的未读消息计数
		unreadMessages.Lock()
		unreadMessages.m[msg.ReceiverID]++
		unreadMessages.Unlock()

		log.Printf("用户 %d 离线，消息已存储", msg.ReceiverID)
	}

	// 发送确认消息给发送者
	confirmation := *msg
	confirmation.Type = "message_sent"
	confirmation.Content = fmt.Sprintf("%s (接收者%s在线)", msg.Content,
		map[bool]string{true: "已", false: "未"}[exists])
	sendMessageToUser(sender.Conn, &confirmation)
}

// sendOfflineMessages 发送用户的所有离线消息
func sendOfflineMessages(user *User) {
	offlineMessages.RLock()
	messages, exists := offlineMessages.m[user.ID]
	offlineMessages.RUnlock()

	if exists && len(messages) > 0 {
		// 发送离线消息通知
		sendSystemMessage(user.Conn, fmt.Sprintf("你有 %d 条离线消息", len(messages)))

		// 逐条发送离线消息
		for _, msg := range messages {
			sendMessageToUser(user.Conn, &msg)
			time.Sleep(100 * time.Millisecond) // 避免消息拥挤
		}

		// 清除已发送的离线消息
		offlineMessages.Lock()
		delete(offlineMessages.m, user.ID)
		offlineMessages.Unlock()

		log.Printf("已向用户 %d 发送 %d 条离线消息", user.ID, len(messages))
	}
}

// handleReadAck 处理消息已读确认
func handleReadAck(msg *Message) {
	if msg.MessageID == "" {
		log.Println("消息ID不能为空")
		return
	}

	// 查找消息的发送者并转发已读确认
	// 这里简化处理，实际应用中应该从消息存储中查找消息的发送者
	// 然后将已读确认发送给原消息的发送者

	// 创建已读确认响应
	ackMsg := Message{
		Type:      "read_ack",
		MessageID: msg.MessageID,
		Timestamp: time.Now().Unix(),
	}
	// 减少未读消息计数
	// 在实际应用中，应该根据消息ID找到对应的接收者
	if msg.ReceiverID != 0 {
		unreadMessages.Lock()
		if count, exists := unreadMessages.m[msg.ReceiverID]; exists && count > 0 {
			unreadMessages.m[msg.ReceiverID] = count - 1
		}
		unreadMessages.Unlock()
	}

	// 如果发送者在线，发送已读确认
	onlineUsers.RLock()
	if sender, exists := onlineUsers.m[msg.ReceiverID]; exists {
		sendMessageToUser(sender.Conn, &ackMsg)
	}
	onlineUsers.RUnlock()
}

// sendMessageToUser 发送消息到指定用户连接
func sendMessageToUser(conn *websocket.Conn, msg *Message) error {
	if conn == nil {
		return fmt.Errorf("连接不存在")
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, msgBytes)
}

// sendError 发送错误消息
func sendError(conn *websocket.Conn, message string) {
	if conn == nil {
		return
	}

	errMsg := Message{
		Type:      "error",
		Content:   message,
		Timestamp: time.Now().Unix(),
	}
	sendMessageToUser(conn, &errMsg)
}

// sendSystemMessage 发送系统消息
func sendSystemMessage(conn *websocket.Conn, message string) {
	if conn == nil {
		return
	}

	sysMsg := Message{
		Type:      "system",
		Content:   message,
		Timestamp: time.Now().Unix(),
	}
	sendMessageToUser(conn, &sysMsg)
}
