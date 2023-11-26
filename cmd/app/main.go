package main

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model/ws"
	"log"

	umlaut "github.com/go-park-mail-ru/2023_2_Umlaut"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	adminProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	authProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	feedProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initMicroservices() (authProto.AuthorizationClient, feedProto.FeedClient, adminProto.AdminClient, error) {
	authConn, err := grpc.Dial(
		viper.GetString("authorization.host")+":"+viper.GetString("authorization.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, nil, err
	}
	feedConn, err := grpc.Dial(
		viper.GetString("feed.host")+":"+viper.GetString("feed.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, nil, err
	}
	adminConn, err := grpc.Dial(
		viper.GetString("admin.host")+":"+viper.GetString("admin.port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return authProto.NewAuthorizationClient(authConn), feedProto.NewFeedClient(feedConn), adminProto.NewAdminClient(adminConn), nil
}

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host umlaut-bmstu.me:8000
// @BasePath /
func main() {
	logger, err := utils.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting server...")
	defer logger.Sync()

	ctx := context.WithValue(context.Background(), "logger", logger)

	db, err := utils.InitPostgres(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres: %s", err.Error())))
	}

	dbAdmin, err := utils.InitPostgresAdmin(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres admin: %s", err.Error())))
	}

	sessionStore, err := utils.InitRedis()
	if err != nil {
		logger.Error("initialize redisDb",
			zap.String("Error", fmt.Sprintf("failed to initialize redisDb: %s", err.Error())))
	}
	defer sessionStore.Close()

	fileClient, err := utils.InitMinioClient()
	if err != nil {
		logger.Error("initialize Minio",
			zap.String("Error", fmt.Sprintf("failed to initialize Minio: %s", err.Error())))
	}

	authClient, feedConn, adminConn, err := initMicroservices()
	if err != nil {
		logger.Error("initialize Microservices",
			zap.String("Error", fmt.Sprintf("failed to initialize microservices: %s", err.Error())))
	}

	hub := ws.NewHub()
	repos := repository.NewRepository(db, dbAdmin, sessionStore, fileClient)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, hub, logger, authClient, feedConn, adminConn)

	srv := new(umlaut.Server)
	go hub.Run()

	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("running http server",
			zap.String("Error", fmt.Sprintf("error occured while running http server: %s", err.Error())))
	}
}
