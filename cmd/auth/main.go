package main

import (
	"context"
	"log"
	"net"

	utils "github.com/go-park-mail-ru/2023_2_Umlaut/cmd"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	db, err := utils.InitPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres: %s", err.Error())
	}

	sessionStore, err := utils.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize redisDb: %s", err.Error())
	}
	defer sessionStore.Close()

	authService := service.NewAuthService(
		repository.NewUserPostgres(db),
		repository.NewRedisStore(sessionStore),
	)

	server := server.NewAuthServer(authService)

	listen, err := net.Listen("tcp", ":"+viper.GetString("authorization.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("authorization.port"), err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAuthorizationServer(grpcServer, server)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("authorization.port"), err.Error())
	}
}
