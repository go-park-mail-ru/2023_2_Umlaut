package main

import (
	"context"
	"fmt"
	"log"

	umlaut "github.com/go-park-mail-ru/2023_2_Umlaut"
	utils "github.com/go-park-mail-ru/2023_2_Umlaut/cmd"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	authProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initMicroservices() (authProto.AuthorizationClient, error) {
	authConn, err := grpc.Dial(
		static.Adress+":"+viper.GetString("auth_port"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return authProto.NewAuthorizationClient(authConn), nil
}

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host localhost:8000
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

	authClient, err := initMicroservices()
	if err != nil {
		logger.Error("initialize Microservices",
			zap.String("Error", fmt.Sprintf("failed to initialize microservices: %s", err.Error())))
	}

	repos := repository.NewRepository(db, sessionStore, fileClient)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logger, authClient)
	srv := new(umlaut.Server)

	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("running http server",
			zap.String("Error", fmt.Sprintf("error occured while running http server: %s", err.Error())))
	}
}
