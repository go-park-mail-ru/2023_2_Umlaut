package ws

import "github.com/go-park-mail-ru/2023_2_Umlaut/model"

type Dialog struct {
	Id                  int              `json:"id" db:"id"`
	User1Id             int              `json:"user1_id" db:"user1_id"`
	User2Id             int              `json:"user2_id" db:"user2_id"`
	Сompanion           string           `json:"companion"`
	СompanionImagePaths *[]string        `json:"сompanion_image_paths"`
	LastMessage         *model.Message   `json:"last_message"`
	Clients             *map[int]*Client `json:"clients" swaggerignore:"true"`
}

type Hub struct {
	Dialogs    map[int]*Dialog
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *model.Message
}

func NewHub() *Hub {
	return &Hub{
		Dialogs:    make(map[int]*Dialog),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *model.Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Dialogs[cl.DialogId]; ok {
				r := h.Dialogs[cl.DialogId]

				if _, ok = (*r.Clients)[cl.Id]; !ok {
					(*r.Clients)[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Dialogs[cl.DialogId]; ok {
				if _, ok = (*h.Dialogs[cl.DialogId].Clients)[cl.Id]; ok {
					delete(*h.Dialogs[cl.DialogId].Clients, cl.Id)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Dialogs[*m.DialogId]; ok {

				for _, cl := range *h.Dialogs[*m.DialogId].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
