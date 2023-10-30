package handler

import (
	"context"
	"errors"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type ctxKey string

const keyUserID ctxKey = "user_id"

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.origins"),
		AllowedMethods:   viper.GetStringSlice("cors.methods"),
		AllowCredentials: true,
	})

	return corsMiddleware.Handler(next)
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			newErrorClientResponseDto(h.ctx, w, http.StatusUnauthorized, "Необходимо авторизироваться")
			return
		}

		id, err := h.services.GetSessionValue(r.Context(), session.Value)
		if err != nil {
			logger, ok := h.ctx.Value("Logger").(*zap.Logger)
			if !ok {
				log.Fatal("Logger not found in context")
			}

			logger.Error("Request handled",
				zap.String("Method", r.Method),
				zap.String("RequestURI", r.RequestURI),
				//zap.Any("Status", h.ctx.Value("Status")),
				//zap.Any("Message", h.ctx.Value("Message")),
				zap.String("Error", err.Error()),
			)
			newErrorClientResponseDto(h.ctx, w, http.StatusUnauthorized, "Необходимо авторизироваться")
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(h.ctx))

		logger, ok := h.ctx.Value("Logger").(*zap.Logger)
		if !ok {
			log.Fatal("Logger not found in context")
		}

		logger.Info("Request handled",
			zap.String("Method", r.Method),
			zap.String("RequestURI", r.RequestURI),
			zap.Any("Status", h.ctx.Value("Status")),
			zap.Any("Message", h.ctx.Value("Message")),
			zap.Duration("Time", time.Since(start)),
		)
	})
}

func (h *Handler) panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger, ok := h.ctx.Value("Logger").(*zap.SugaredLogger)
				if !ok {
					log.Fatal("Logger not found in context")
				}

				logger.Error("Request handled",
					zap.String("Method", r.Method),
					zap.String("RequestURI", r.RequestURI),
					zap.String("Message", string(debug.Stack())),
				)
				newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
