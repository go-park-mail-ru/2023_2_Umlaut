package handler

import (
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		logrus.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func panicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				newErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				logrus.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}
