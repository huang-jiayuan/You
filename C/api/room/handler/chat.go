package handler

import (
	"encoding/json"
	"hash/crc32"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 消息结构体
type Message struct {
	Type    string    `json:"type"`    // 消息类型：chat, join, leave等
	Sender  int64     `json:"sender"`  // 发送者ID
	Content string    `json:"content"` // 消息内容
	Time    time.Time `json:"time"`    // 发送时间
}

// 客户端结构体
type Client struct {
	ID     int64
	Conn   *websocket.Conn
	Server *Server
	send   chan []byte
}

// 服务器结构体
type Server struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex
}

// 创建新服务器
func NewServer() *Server {
	return &Server{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// 运行服务器
func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
			log.Printf("新客户端连接: %d, 当前在线人数: %d", client.ID, len(s.clients))

		case client := <-s.unregister:
			s.mutex.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
				log.Printf("客户端断开连接: %d, 当前在线人数: %d", client.ID, len(s.clients))
			}
			s.mutex.Unlock()

		case message := <-s.broadcast:
			s.mutex.Lock()
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
			s.mutex.Unlock()
		}
	}
}

// 客户端读取消息
func (c *Client) Read() {
	defer func() {
		// 发送离开消息
		leaveMsg := Message{
			Type:    "leave",
			Sender:  c.ID,
			Content: "用户离开了聊天室",
			Time:    time.Now(),
		}
		if msgBytes, err := json.Marshal(leaveMsg); err == nil {
			c.Server.broadcast <- msgBytes
		}

		c.Server.unregister <- c
		c.Conn.Close()
	}()

	// 发送加入消息
	joinMsg := Message{
		Type:    "join",
		Sender:  c.ID,
		Content: "新用户加入了聊天室",
		Time:    time.Now(),
	}
	if msgBytes, err := json.Marshal(joinMsg); err == nil {
		c.Server.broadcast <- msgBytes
	}

	for {
		// 只接受文本消息，拒绝二进制消息（图片等）
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("错误: %v", err)
			}
			break
		}

		// 尝试解析消息
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			// 如果不是JSON格式，创建一个默认的聊天消息
			msg = Message{
				Type:    "chat",
				Sender:  c.ID,
				Content: string(message),
				Time:    time.Now(),
			}
			message, _ = json.Marshal(msg)
		} else {
			// 确保发送者ID是正确的
			msg.Sender = c.ID
			msg.Time = time.Now()
			message, _ = json.Marshal(msg)
		}

		// 将消息广播给所有客户端
		c.Server.broadcast <- message
	}
}

// 客户端写入消息
func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

// WebSocket升级配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源的连接
		return true
	},
}

// 处理WebSocket连接
func (s *Server) HandleWebSocket(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket升级失败:", err)
		return
	}

	// 生成客户端ID（使用远程地址和时间戳组合）
	clientIP := c.ClientIP()
	timestamp := time.Now().UnixNano()
	// 使用IP地址的哈希值和时间戳生成唯一ID
	hashValue := int64(crc32.ChecksumIEEE([]byte(clientIP)))
	clientID := hashValue ^ timestamp

	// 创建客户端
	client := &Client{
		ID:     clientID,
		Conn:   conn,
		Server: s,
		send:   make(chan []byte, 256),
	}

	log.Printf("新WebSocket连接: %s, ID: %d", clientIP, clientID)

	// 注册客户端
	s.register <- client

	// 启动读写协程
	go client.Read()
	go client.Write()
}

// 全局服务器实例
var GlobalServer *Server

// 初始化全局服务器
func InitServer() {
	GlobalServer = NewServer()
	go GlobalServer.Run()
}

// 全局WebSocket处理函数
func HandleWebSocket(context *gin.Context) {
	if GlobalServer == nil {
		InitServer()
	}
	GlobalServer.HandleWebSocket(context)
}
