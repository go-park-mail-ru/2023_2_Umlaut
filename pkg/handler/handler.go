package handler

import (
	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handler struct {
	Repositories *repository.Repository
}

func NewHandler(repositories *repository.Repository) *Handler {
	return &Handler{Repositories: repositories}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)

	mux.HandleFunc("/auth/login", h.signIn)
	mux.HandleFunc("/auth/logout", h.logout)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	mux.HandleFunc("/api/feed", h.feed)

	return mux
}
