package ws

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
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
		c.Conn.Close()
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
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.Info("WS",
					zap.String("Place", "client.go: 59"),
					zap.String("Message", "error"),
					zap.Error(err),
				)
			}
			break
		}
		if len(m) < 1 {
			continue
		}

		var receivedMessage Message
		err = json.Unmarshal(m, &receivedMessage)
		isEdit := false
		if receivedMessage.Id > 0 {
			isEdit = true
		}
		if err != nil {
			c.Logger.Info("WS",
				zap.String("Place", "client.go: 71"),
				zap.String("Message", "error"),
				zap.Error(err),
			)
			break
		}
		newMessage, err := services.Message.SaveOrUpdateMessage(ctx, model.Message{
			Id:          &receivedMessage.Id,
			SenderId:    &receivedMessage.SenderId,
			DialogId:    &receivedMessage.DialogId,
			RecipientId: &receivedMessage.RecipientId,
			Text:        &receivedMessage.Text,
			IsRead:      &receivedMessage.IsRead,
		})
		if err != nil {
			c.Logger.Info("WS",
				zap.String("Place", "client.go: 84"),
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
				Type:    static.Message,
				Payload: message,
			}
		}
	}
}
