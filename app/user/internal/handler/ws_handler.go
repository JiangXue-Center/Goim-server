package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	// WebSocket 升级器，将 HTTP 连接升级为 WebSocket 连接
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	// 存储所有 WebSocket 连接的 map，键为用户 ID，值为 WebSocket 连接
	clients = make(map[string]*websocket.Conn)

	// 保护 clients map 并发访问的互斥锁
	clientsMux sync.Mutex
)

// Message 是我们通过 WebSocket 发送和接收的 JSON 数据结构
type Message struct {

	//消息发送者
	SenderID string `json:"sender_id"`

	//消息接收者
	RecipientID string `json:"recipient_id"`

	//消息内容
	Content string `json:"content"`

	//消息类型
	MessageType string `json:"message_type"`
}

// WebSocketHandler 处理 WebSocket 连接请求
func WebSocketHandler(c *gin.Context) {
	// 从 URL 参数中获取用户 ID
	id := c.Param("id")

	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// 将新的 WebSocket 连接添加到 clients map 中
	clientsMux.Lock()
	clients[id] = conn
	clientsMux.Unlock()

	// 启动一个新的 goroutine 来处理这个连接，以免阻塞当前的 HTTP 处理流程
	go handleWebSocketConnection(id, conn)
}

// handleWebSocketConnection 处理具体的 WebSocket 连接逻辑
func handleWebSocketConnection(id string, conn *websocket.Conn) {
	// 在函数结束时关闭连接并从 clients map 中删除该连接
	defer func() {
		conn.Close()
		clientsMux.Lock()
		delete(clients, id)
		clientsMux.Unlock()
	}()

	// 无限循环，读取来自 WebSocket 连接的消息
	for {
		// 读取消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// 将消息反序列化为 Message 结构体
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		// 打印接收到的消息信息
		log.Printf("Received %s type message from %s to %s: %s",
			msg.MessageType, msg.SenderID, msg.RecipientID, msg.Content)

		// 将消息发送给目标客户端
		sendToClient(msg)
	}
}

// sendToClient 将消息发送给目标客户端
func sendToClient(msg Message) {
	// 保护 clients map 并发访问的互斥锁
	clientsMux.Lock()
	defer clientsMux.Unlock()

	// 查找目标用户的 WebSocket 连接
	if conn, ok := clients[msg.RecipientID]; ok {
		// 将 Message 结构体序列化为 JSON
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Marshal error: %v", err)
			return
		}

		// 发送消息给目标客户端
		if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			log.Printf("Error sending message to %s: %v", msg.RecipientID, err)
		} else {
			log.Printf("Message sent to %s: %s", msg.RecipientID, msg.Content)
		}
	} else {
		// 如果目标用户不在线，记录日志或执行其他处理逻辑
		log.Printf("Client %s not found", msg.RecipientID)
		// 可以选择将消息存储在数据库中，等目标用户上线后再发送
	}
}
