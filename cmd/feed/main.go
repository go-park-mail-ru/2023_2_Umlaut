package main

import (
	"context"
	"log"
	"net"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
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

	feedService := service.NewFeedService(
		repository.NewUserPostgres(db),
		repository.NewRedisStore(sessionStore),
		repository.NewDialogPostgres(db),
	)

	feedServer := server.NewFeedServer(feedService)
	viper.GetString("feed.port")
	listen, err := net.Listen("tcp", ":"+viper.GetString("feed.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("feed.port"), err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterFeedServer(grpcServer, feedServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("feed.port"), err.Error())
	}
}
