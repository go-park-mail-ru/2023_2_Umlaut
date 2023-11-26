package ws

type Hub struct {
	Users      map[int]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Users:      make(map[int]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
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
				close(cl.Message)
			}
		case m := <-h.Broadcast:
			if _, ok := h.Users[m.RecipientId]; ok {
				h.Users[m.RecipientId].Message <- m
			}
		}
	}
}
