package handler

import (
	"context"
	"errors"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"runtime/debug"
	"time"
)

func corsMiddleware(next http.Handler) http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.origins"),
		AllowedMethods:   viper.GetStringSlice("cors.methods"),
		AllowCredentials: true,
	})

	return corsMiddleware.Handler(next)
}

func authMiddleware(h *Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := r.Cookie("session_id")
			if errors.Is(err, http.ErrNoCookie) {
				newErrorClientResponseDto(w, http.StatusUnauthorized, "Необходимо авторизироваться")
				return
			}

			id, err := h.services.GetSessionValue(r.Context(), session.Value)
			if err != nil {
				logrus.Printf("[RequestID] %s [error] %v", r.Context().Value("RequestID"), err.Error())
				newErrorClientResponseDto(w, http.StatusUnauthorized, "Необходимо авторизироваться")
				return
			}

			ctx := context.WithValue(r.Context(), "userID", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.Printf("[RequestID] %s [Method] %s [RequestURI] %s [time] %s",
			r.Context().Value("RequestID"), r.Method, r.RequestURI, time.Since(start))
	})
}

func panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				newErrorClientResponseDto(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				logrus.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
