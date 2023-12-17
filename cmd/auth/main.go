package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/keepalive"

	initial "github.com/go-park-mail-ru/2023_2_Umlaut/cmd"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/interceptors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	initial.InitConfig()
	grpc_prometheus.EnableHandlingTimeHistogram()
	ctx := context.Background()

	db, err := initial.InitPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres: %s", err.Error())
	}

	dbAdmin, err := initial.InitPostgresAdmin(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres admin: %s", err.Error())
	}

	sessionStore, err := initial.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize redisDb: %s", err.Error())
	}
	defer sessionStore.Close()

	authService := service.NewAuthService(
		repository.NewUserPostgres(db),
		repository.NewRedisStore(sessionStore),
		repository.NewAdminPostgres(dbAdmin),
	)

	authServer := server.NewAuthServer(authService)

	listen, err := net.Listen("tcp", ":"+viper.GetString("authorization.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("authorization.port"), err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc.UnaryServerInterceptor(interceptors.PanicRecoveryInterceptor),
			grpc.UnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)

	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{
		Addr:    ":9091",
		Handler: nil,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Printf("Failed to start Prometheus metrics server: %s\n", err)
		}
	}()

	proto.RegisterAuthorizationServer(grpcServer, authServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("authorization.port"), err.Error())
	}
}
