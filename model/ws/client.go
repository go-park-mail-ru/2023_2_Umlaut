package ws

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

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
	Conn    *websocket.Conn
	Message chan *Message
	Id      int          `json:"id" db:"id"`
	Dialogs map[int]bool `json:"clients"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
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
				log.Printf("error: %v", err)
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
			log.Printf("error: %v", err)
			break
		}
		newMessage, err := services.Message.SaveOrUpdateMessage(ctx, model.Message{
			Id:       &receivedMessage.Id,
			SenderId: &receivedMessage.SenderId,
			DialogId: &receivedMessage.DialogId,
			Text:     &receivedMessage.Text,
			IsRead:   receivedMessage.IsRead,
		})
		if err != nil {
			//TODO: do something
		} else if !isEdit {
			hub.Broadcast <- &Message{
				Id:        *newMessage.Id,
				SenderId:  *newMessage.SenderId,
				DialogId:  *newMessage.DialogId,
				Text:      *newMessage.Text,
				IsRead:    newMessage.IsRead,
				CreatedAt: *newMessage.CreatedAt,
			}
		}
	}
}
