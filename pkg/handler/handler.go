package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
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
	mux.HandleFunc("/auth/login", h.loginHandler)
	mux.HandleFunc("/auth/logout", h.logoutHandler)
	mux.HandleFunc("/auth/sign-up", h.signUpHandler)

	mux.HandleFunc("/api/feed/", h.feedHandler)

	return mux
}
