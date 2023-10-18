package handler

import (
	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://37.139.32.76:8000/swagger/doc.json"),
	))

	r.HandleFunc("/auth/login", h.signIn).Methods("POST")
	r.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	r.HandleFunc("/auth/logout", h.logout)

	r.HandleFunc("/api/feed", h.feed).Methods("GET")

	r.HandleFunc("/api/user", h.user).Methods("GET")
	r.HandleFunc("/api/user", h.updateUser).Methods("POST")

	r.Use(
		corsMiddleware,
		loggingMiddleware,
		panicRecoveryMiddleware,
	)

	return r
}
