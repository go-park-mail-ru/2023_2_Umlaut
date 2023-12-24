package chat

import (
	"context"
	"encoding/json"
	"runtime/debug"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Notification struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Message struct {
	Id          int       `json:"id"`
	SenderId    int       `json:"sender_id"`
	RecipientId int       `json:"recipient_id"`
	DialogId    int       `json:"dialog_id"`
	Text        string    `json:"message_text"`
	IsRead      bool      `json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}

type Client struct {
	Conn          *websocket.Conn
	Notifications chan *Notification
	Logger        *zap.Logger
	Id            int `json:"id" db:"id"`
}

func (c *Client) WriteMessage() {
	defer func() {
		if r := recover(); r != nil {
			c.Logger.Error("Panic [WS]",
				zap.String("Message", "Panic in WriteMessage"),
				zap.String("Error", string(debug.Stack())),
			)
			c.WriteMessage()
		} else {
			c.Conn.Close()
		}
	}()
	for {
		message, ok := <-c.Notifications
		if !ok {
			c.Logger.Info("WS",
				zap.String("Place", "client.go: 43"),
				zap.String("Message", "error"),
			)
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(ctx context.Context, hub *Hub, services *service.Service) {
	defer func() {
		if r := recover(); r != nil {
			c.Logger.Error("Panic [WS]",
				zap.String("Message", "Panic in WriteMessage"),
				zap.String("Error", string(debug.Stack())),
			)
			c.ReadMessage(ctx, hub, services)
		} else {
			hub.Unregister <- c
			c.Conn.Close()
		}
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.Info("WS",
					zap.String("Message", "error"),
					zap.Error(err),
				)
			}
			break
		}

		var receivedMessage Message
		err = json.Unmarshal(m, &receivedMessage)
		isEdit := false
		if receivedMessage.Id > 0 {
			isEdit = true
		}
		if err != nil {
			c.Logger.Info("WS",
				zap.String("Message", "error"),
				zap.Error(err),
			)
			continue
		}
		newMessage, err := services.Message.SaveOrUpdateMessage(ctx, core.Message{
			Id:          &receivedMessage.Id,
			SenderId:    &receivedMessage.SenderId,
			DialogId:    &receivedMessage.DialogId,
			RecipientId: &receivedMessage.RecipientId,
			Text:        &receivedMessage.Text,
			IsRead:      &receivedMessage.IsRead,
		})
		if err != nil {
			c.Logger.Info("WS",
				zap.String("Message", "error"),
				zap.Error(err),
			)
		} else if !isEdit {
			message := &Message{
				Id:          *newMessage.Id,
				SenderId:    *newMessage.SenderId,
				RecipientId: *newMessage.RecipientId,
				DialogId:    *newMessage.DialogId,
				Text:        *newMessage.Text,
				IsRead:      *newMessage.IsRead,
				CreatedAt:   *newMessage.CreatedAt,
			}
			hub.Broadcast <- &Notification{
				Type:    constants.Message,
				Payload: message,
			}
		}
	}
}
