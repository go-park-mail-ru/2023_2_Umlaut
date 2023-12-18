package main

import (
	"context"
	"fmt"
	umlaut "github.com/go-park-mail-ru/2023_2_Umlaut/internal/api"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/api/handler"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core/chat"
	"log"
	"os"
	"os/signal"
	"syscall"

	initial "github.com/go-park-mail-ru/2023_2_Umlaut/cmd"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host umlaut-bmstu.me
// @BasePath /
func main() {
	initial.InitConfig()
	logger, err := initial.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting server...")
	defer logger.Sync()

	ctx := context.Background()

	db, err := initial.InitPostgres(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres: %s", err.Error())))
	}

	dbAdmin, err := initial.InitPostgresAdmin(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres admin: %s", err.Error())))
	}

	sessionStore, err := initial.InitRedis()
	if err != nil {
		logger.Error("initialize redisDb",
			zap.String("Error", fmt.Sprintf("failed to initialize redisDb: %s", err.Error())))
	}
	defer sessionStore.Close()

	fileClient, err := initial.InitMinioClient()
	if err != nil {
		logger.Error("initialize Minio",
			zap.String("Error", fmt.Sprintf("failed to initialize Minio: %s", err.Error())))
	}

	authClient, feedConn, adminConn, err := initial.InitMicroservices()
	if err != nil {
		logger.Error("initialize Microservices",
			zap.String("Error", fmt.Sprintf("failed to initialize microservices: %s", err.Error())))
	}

	hub := chat.NewHub(logger)
	repos := repository.NewRepository(db, dbAdmin, sessionStore, fileClient)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, hub, logger, authClient, feedConn, adminConn)

	srv := new(umlaut.Server)
	go hub.Run()
	if err = services.Background.ResetDislike(ctx); err != nil {
		logger.Error("initialize Background Service",
			zap.String("Error", fmt.Sprintf("ResetDislike: %s", err.Error())))
	}
	if err = services.Background.ResetLikeCounter(ctx); err != nil {
		logger.Error("initialize Background Service",
			zap.String("Error", fmt.Sprintf("ResetLikeCounter: %s", err.Error())))
	}
	go func() {
		if err = srv.Serve(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Error("running http server",
				zap.String("Error", fmt.Sprintf("error occured while running http server: %s", err.Error())))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("TodoApp Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on server shutting down: %s",
			zap.Error(err))
	}
}
