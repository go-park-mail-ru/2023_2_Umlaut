package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ctxKey string

const (
	keyUserID    ctxKey = "user_id"
	keyAdminID   ctxKey = "admin_id"
	keyLogger    ctxKey = "logger"
	keyRequestId ctxKey = "request_id"
	secret              = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.origins"),
		AllowedMethods:   viper.GetStringSlice("cors.methods"),
		AllowedHeaders:   viper.GetStringSlice("cors.headers"),
		AllowCredentials: true,
	})

	return corsMiddleware.Handler(next)
}

func (h *Handler) authAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("admin_session_id")
		if errors.Is(err, http.ErrNoCookie) {
			newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}

		id, err := h.services.Authorization.GetSessionValue(r.Context(), session.Value)
		if err != nil {
			newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}

		ctx := context.WithValue(r.Context(), keyAdminID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}

		id, err := h.services.Authorization.GetSessionValue(r.Context(), session.Value)
		if err != nil {
			newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) csrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}
		session, _ := r.Cookie("session_id")
		jwtToken := NewJwtToken(secret)

		CSRFToken := r.Header.Get("X-Csrf-Token")

		valid, err := jwtToken.Check(session.Value, CSRFToken)
		if err != nil || !valid {
			newErrorClientResponseDto(r.Context(), w, http.StatusForbidden, "Need csrf token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		childLogger := h.logger.With(zap.String("RequestID", uuid.NewString()))
		ctx := context.WithValue(r.Context(), keyLogger, childLogger)
		ctx = context.WithValue(ctx, keyRequestId, uuid.NewString())

		next.ServeHTTP(w, r.WithContext(ctx))

		method := r.Method
		path := r.RequestURI
		timing := time.Since(start)

		childLogger.Info("Request handled",
			zap.String("Method", method),
			zap.String("RequestURI", path),
			zap.Duration("Time", timing),
		)

		status := 200
		h.metrics.Hits.WithLabelValues(strconv.Itoa(status), path, method).Inc()
		h.metrics.Duration.WithLabelValues(strconv.Itoa(status), path, method).Observe(timing.Seconds())
	})
}

func (h *Handler) panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger, ok := r.Context().Value(keyLogger).(*zap.Logger)
				if !ok {
					log.Println("Logger not found in context")
				}

				logger.Error("Panic",
					zap.String("Method", r.Method),
					zap.String("RequestURI", r.RequestURI),
					zap.String("Error", err.(string)),
					zap.String("Message", string(debug.Stack())),
				)
				newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
