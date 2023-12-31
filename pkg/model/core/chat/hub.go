package chat

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"go.uber.org/zap"
	"runtime/debug"
)

type Hub struct {
	Users      map[int]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Notification
	Services   *service.Service
	Logger     *zap.Logger
}

func NewHub(logger *zap.Logger, services *service.Service) *Hub {
	return &Hub{
		Users:      make(map[int]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Notification, 5),
		Services:   services,
		Logger:     logger,
	}
}

func (h *Hub) Run() {
	defer func() {
		if r := recover(); r != nil {
			h.Logger.Error("Panic [WS]",
				zap.String("Message", "Panic in Serve"),
				zap.String("Error", string(debug.Stack())),
			)
			h.Run()
		}
	}()
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
			case constants.Message:
				if message, ok := m.Payload.(*Message); ok {
					if user, userExists := h.Users[message.RecipientId]; userExists {
						user.Notifications <- m
					}
				} else {
					h.Logger.Info("[WS] (*Message)",
						zap.String("Message", "Ошибка преобразования типа"),
					)
				}
			case constants.Match:
				if match, ok := m.Payload.(core.Dialog); ok {
					ctx := context.Background()
					if user1, user1Exists := h.Users[match.User1Id]; user1Exists {
						dialog, err := h.Services.Dialog.GetDialog(ctx, match.Id, match.User1Id)
						if err == nil {
							m.Payload = dialog
						}
						user1.Notifications <- m
					}
					if user2, user2Exists := h.Users[match.User2Id]; user2Exists {
						dialog, err := h.Services.Dialog.GetDialog(ctx, match.Id, match.User2Id)
						if err == nil {
							m.Payload = dialog
						}
						user2.Notifications <- m
					}
				} else {
					h.Logger.Info("[WS] (Dialog)",
						zap.String("Message", "Ошибка преобразования типа"),
					)
				}
			}
		}
	}
}
