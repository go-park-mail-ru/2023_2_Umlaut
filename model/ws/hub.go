package ws

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"log"
)

type Hub struct {
	Users      map[int]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Notification
}

func NewHub() *Hub {
	return &Hub{
		Users:      make(map[int]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Notification, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			h.Users[cl.Id] = cl
		case cl := <-h.Unregister:
			if _, ok := h.Users[cl.Id]; ok {
				delete(h.Users, cl.Id)
				close(cl.Notifications)
			}
		case m := <-h.Broadcast:
			switch m.Type {
			case static.Message:
				message := m.Payload.(Message)
				if _, ok := h.Users[message.RecipientId]; ok {
					h.Users[message.RecipientId].Notifications <- m
				}
			case static.Match:
				log.Println("[HUB] match like")
				match := m.Payload.(model.Dialog)
				if _, ok := h.Users[match.User1Id]; ok {
					h.Users[match.User1Id].Notifications <- m
				}
				if _, ok := h.Users[match.User2Id]; ok {
					h.Users[match.User2Id].Notifications <- m
				}
				log.Println("[HUB] match like send")
			}

		}
	}
}
