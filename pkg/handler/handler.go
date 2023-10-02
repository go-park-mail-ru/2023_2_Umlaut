package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"net/http"

	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("/auth/login", h.signIn)
	mux.HandleFunc("/auth/logout", h.logout)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	mux.HandleFunc("/api/feed", h.feed)
	mux.HandleFunc("/api/user", h.user)

	return mux
}
