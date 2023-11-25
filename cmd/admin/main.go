package main

import (
	"context"
	"log"
	"net"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	db, err := utils.InitPostgresAdmin(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres: %s", err.Error())
	}

	sessionStore, err := utils.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize redisDb: %s", err.Error())
	}
	defer sessionStore.Close()

	adminService := service.NewAdminService(repository.NewAdminPostgres(db))

	adminServer := server.NewAdminServer(adminService)
	viper.GetString("admin.port")
	listen, err := net.Listen("tcp", ":"+viper.GetString("admin.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("admin.port"), err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAdminServer(grpcServer, adminServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("admin.port"), err.Error())
	}
}
