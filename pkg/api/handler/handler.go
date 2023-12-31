package handler

import (
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core/chat"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/monitoring"

	_ "github.com/go-park-mail-ru/2023_2_Umlaut/docs"
	adminProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	authProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	feedProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Handler struct {
	authMicroservice  authProto.AuthorizationClient
	feedMicroservice  feedProto.FeedClient
	adminMicroservice adminProto.AdminClient
	services          *service.Service
	hub               *chat.Hub
	logger            *zap.Logger
	metrics           *monitoring.PrometheusMetrics
}

func NewHandler(
	services *service.Service,
	hub *chat.Hub,
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
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	api := r.PathPrefix("/api").Subrouter()
	authRouter := api.PathPrefix("/v1/auth").Subrouter()
	authRouter.HandleFunc("/login", h.signIn).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/sign-up", h.signUp).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/vk-login", h.vkLogin)
	authRouter.HandleFunc("/vk-sign-up", h.vkSignUp)
	authRouter.HandleFunc("/logout", h.logout)
	authRouter.HandleFunc("/admin", h.logInAdmin)

	adminRouter := api.PathPrefix("/v1/admin").Subrouter()
	adminRouter.Use(
		h.authAdminMiddleware,
	)
	adminRouter.HandleFunc("/complaint", h.getNextComplaint).Methods("GET")
	adminRouter.HandleFunc("/complaint/{id}", h.deleteComplaint).Methods("DELETE", "OPTIONS")
	adminRouter.HandleFunc("/complaint/{id}", h.acceptComplaint).Methods("GET")
	adminRouter.HandleFunc("/feedback", h.getFeedbackStatistic).Methods("GET")
	//adminRouter.HandleFunc("/feed-feedback", h.get).Methods("GET")
	adminRouter.HandleFunc("/recommendation", h.getRecommendationStatistic).Methods("GET")

	apiRouter := api.PathPrefix("/v1").Subrouter()
	apiRouter.Use(
		h.csrfMiddleware,
		h.authMiddleware,
	)
	apiRouter.HandleFunc("/feed", h.feed).Methods("GET")
	apiRouter.HandleFunc("/user", h.user).Methods("GET")
	apiRouter.HandleFunc("/user", h.updateUser).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/share", h.getUserShareCridentials).Methods("GET")
	apiRouter.HandleFunc("/user/{id}", h.userById).Methods("GET")
	apiRouter.HandleFunc("/user/photo", h.updateUserPhoto).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/user/photo", h.deleteUserPhoto).Methods("DELETE", "OPTIONS")

	apiRouter.HandleFunc("/like", h.createLike).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/dialogs", h.getDialogs).Methods("GET")
	apiRouter.HandleFunc("/dialogs/{id}", h.getDialog).Methods("GET")
	apiRouter.HandleFunc("/dialogs/{id}/message", h.getDialogMessage).Methods("GET")
	apiRouter.HandleFunc("/tag", h.getAllTags).Methods("GET")
	apiRouter.HandleFunc("/complaint_types", h.getAllComplaintTypes).Methods("GET")
	apiRouter.HandleFunc("/complaint", h.createComplaint).Methods("POST", "OPTIONS")

	apiRouter.HandleFunc("/premium/likes", h.getPremiumLikes).Methods("GET")

	apiRouter.HandleFunc("/ws/messenger", h.registerUserToHub).Methods("GET")

	apiRouter.HandleFunc("/feedback", h.createFeedback).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/recommendation", h.createRecommendation).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/feed-feedback", h.createFeedFeedback).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/show-csat", h.showCSAT).Methods("GET")

	api.Use(
		h.loggingMiddleware,
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
