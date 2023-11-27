package handler

import (
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/monitoring"
	"net/http"

	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model/ws"
	adminProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	authProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	feedProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handler struct {
	authMicroservice  authProto.AuthorizationClient
	feedMicroservice  feedProto.FeedClient
	adminMicroservice adminProto.AdminClient
	services          *service.Service
	hub               *ws.Hub
	logger            *zap.Logger
	metrics           *monitoring.PrometheusMetrics
}

func NewHandler(
	services *service.Service,
	hub *ws.Hub,
	logger *zap.Logger,
	authMicroservice authProto.AuthorizationClient,
	feedMicroservice feedProto.FeedClient,
	adminMicroservice adminProto.AdminClient) *Handler {
	return &Handler{
		services:          services,
		hub:               hub,
		logger:            logger,
		authMicroservice:  authMicroservice,
		feedMicroservice:  feedMicroservice,
		adminMicroservice: adminMicroservice,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()
	h.metrics = monitoring.RegisterMonitoring(r)
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", static.Host)),
	))

	authRouter := r.PathPrefix("/api/v1/auth").Subrouter()
	authRouter.HandleFunc("/login", h.signIn).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", h.logout)
	authRouter.HandleFunc("/admin", h.logInAdmin)

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(
		//h.csrfMiddleware,
		h.authMiddleware,
	)
	apiRouter.HandleFunc("/feed", h.feed).Methods("GET")
	apiRouter.HandleFunc("/user", h.user).Methods("GET")
	apiRouter.HandleFunc("/user", h.updateUser).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.updateUserPhoto).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.deleteUserPhoto).Methods("DELETE")
	apiRouter.HandleFunc("/like", h.createLike).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/dialogs", h.getDialogs).Methods("GET")
	apiRouter.HandleFunc("/dialogs/{id}/message", h.getDialogMessage).Methods("GET")
	apiRouter.HandleFunc("/tag", h.getAllTags).Methods("GET")
	apiRouter.HandleFunc("/complaint_types", h.getAllComplaintTypes).Methods("GET")
	apiRouter.HandleFunc("/complaint", h.createComplaint).Methods("POST", "OPTIONS")

	apiRouter.HandleFunc("/ws/messenger", h.registerUserToHub).Methods("GET")

	apiRouter.HandleFunc("/feedback", h.createFeedback).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/feed-feedback", h.createFeedFeedback).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/recommendation", h.createRecommendation).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/show-csat", h.showCSAT).Methods("GET")

	adminRouter := r.PathPrefix("/api/v1/admin").Subrouter()
	adminRouter.Use(
		h.authAdminMiddleware,
	)
	adminRouter.HandleFunc("/feedback", h.getFeedbackStatistic).Methods("GET")
	//adminRouter.HandleFunc("/feed-feedback", h.get).Methods("GET")
	adminRouter.HandleFunc("/recommendation", h.getRecommendationStatistic).Methods("GET")

	r.Use(
		h.loggingMiddleware,
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
