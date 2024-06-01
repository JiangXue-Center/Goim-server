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
	upgrader   = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients    = make(map[string]*websocket.Conn)
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

func WebSocketHandler(c *gin.Context) {
	id := c.Param("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	clientsMux.Lock()
	clients[id] = conn
	clientsMux.Unlock()

	// 启动一个新的 goroutine 来处理这个连接
	go handleWebSocketConnection(id, conn)
}

func handleWebSocketConnection(id string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		clientsMux.Lock()
		delete(clients, id)
		clientsMux.Unlock()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}
		log.Printf("Received %s type message from %s to %s: %s", msg.SenderID, msg.RecipientID, msg.Content, msg.MessageType)

		sendToClient(msg)
	}
}

func sendToClient(msg Message) {
	clientsMux.Lock()
	defer clientsMux.Unlock()

	if conn, ok := clients[msg.RecipientID]; ok {
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Marshal error: %v", err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			log.Printf("Error sending message to %s: %v", msg.RecipientID, err)
		} else {
			log.Printf("Message sent to %s: %s", msg.RecipientID, msg.Content)
		}
	} else {
		log.Printf("Client %s not found", msg.RecipientID)
		// 如果目标用户不在线，可以选择缓存消息或进行其他处理
		// 例如，将消息存储在数据库中，等目标用户上线后再发送
	}
}
