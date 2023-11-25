package handler

import (
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model/ws"
	"net/http"

	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handler struct {
	services *service.Service
	hub      *ws.Hub
	logger   *zap.Logger
}

func NewHandler(services *service.Service, hub *ws.Hub, logger *zap.Logger) *Handler {
	return &Handler{services: services, hub: hub, logger: logger}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s:8000/swagger/doc.json", static.Host)),
	))

	authRouter := r.PathPrefix("/api/v1/auth").Subrouter()
	authRouter.HandleFunc("/login", h.signIn).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", h.logout)

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(
		//h.csrfMiddleware,
		h.authMiddleware,
	)
	apiRouter.HandleFunc("/feed", h.feed).Methods("GET")
	apiRouter.HandleFunc("/feed/users", h.getNextUsers).Methods("GET")
	apiRouter.HandleFunc("/user", h.user).Methods("GET")
	apiRouter.HandleFunc("/user", h.updateUser).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.updateUserPhoto).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.deleteUserPhoto).Methods("DELETE")
	apiRouter.HandleFunc("/like", h.createLike).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/dialogs", h.getDialogs).Methods("GET")
	apiRouter.HandleFunc("/dialogs/{id}/message", h.getDialogMessage).Methods("GET")
	apiRouter.HandleFunc("/tag", h.getAllTags).Methods("GET")
	apiRouter.HandleFunc("/ws/messenger", h.registerUserToHub).Methods("GET")
	apiRouter.HandleFunc("/complaint_types", h.getAllComplaintTypes).Methods("GET")

	r.Use(
		h.loggingMiddleware,
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
