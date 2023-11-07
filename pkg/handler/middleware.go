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

const (
	keyUserID  ctxKey = "user_id"
	keyStatus  ctxKey = "status"
	keyMessage ctxKey = "message"
	secret            = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

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
			newErrorClientResponseDto(h.ctx, w, http.StatusUnauthorized, "Need auth")
			return
		}

		id, err := h.services.GetSessionValue(r.Context(), session.Value)
		if err != nil {
			newErrorClientResponseDto(h.ctx, w, http.StatusUnauthorized, "Need auth")
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (h *Handler) csrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r.WithContext(*h.ctx))
			return
		}
		session, _ := r.Cookie("session_id")
		jwtToken := NewJwtToken(h.ctx, secret)

		CSRFToken := r.Header.Get("X-Csrf-Token")

		valid, err := jwtToken.Check(session.Value, CSRFToken)
		if err != nil || !valid {
			newErrorClientResponseDto(h.ctx, w, http.StatusForbidden, "Need csrf token")
			return
		}

		next.ServeHTTP(w, r.WithContext(*h.ctx))
	})
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger, ok := (*h.ctx).Value("logger").(*zap.Logger)
		if !ok {
			log.Println("Logger not found in context")
		}

		logger.Info("Request handled",
			zap.String("Method", r.Method),
			zap.String("RequestURI", r.RequestURI),
			zap.Any("Status", (*h.ctx).Value(keyStatus)),
			zap.Any("Message", (*h.ctx).Value(keyMessage)),
			zap.Duration("Time", time.Since(start)),
		)
	})
}

func (h *Handler) panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger, ok := (*h.ctx).Value("logger").(*zap.Logger)
				if !ok {
					log.Println("Logger not found in context")
				}

				logger.Error("Panic",
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
