package ws

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *model.Message
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
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

func (c *Client) ReadMessage(hub *Hub) {
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
		text := string(m)
		msg := &model.Message{
			Text:     &text,
			DialogId: &c.DialogId,
			SenderId: &c.Id,
		}

		hub.Broadcast <- msg
	}
}
