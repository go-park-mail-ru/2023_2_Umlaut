package handler

import (
	"context"
	"net/http"

	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
	ctx      *context.Context
}

func NewHandler(services *service.Service, ctx *context.Context) *Handler {
	return &Handler{services: services, ctx: ctx}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
	))

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", h.signIn).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", h.logout)

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(
		//h.csrfMiddleware,
		h.authMiddleware,
	)
	apiRouter.HandleFunc("/feed", h.feed).Methods("GET")
	apiRouter.HandleFunc("/user", h.user).Methods("GET")
	apiRouter.HandleFunc("/user", h.updateUser).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.updateUserPhoto).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/{id}/photo", h.getUserPhoto).Methods("GET")
	apiRouter.HandleFunc("/user/photo", h.deleteUserPhoto).Methods("DELETE")
	apiRouter.HandleFunc("/like", h.createLike).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/dialogs", h.getDialogs).Methods("GET")
	apiRouter.HandleFunc("/tags", h.getAllTags).Methods("GET")

	r.Use(
		h.loggingMiddleware,
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
