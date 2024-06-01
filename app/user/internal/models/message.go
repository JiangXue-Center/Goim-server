package models

type Message struct {

	//消息发送者
	SenderId string `json:"sender_id"`

	//消息接收者
	RecipientId string `json:"recipient_id"`

	//消息内容
	Content string `json:"content"`

	//消息类型
	MessageType string `json:"message_type"`
}
