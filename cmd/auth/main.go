package main

import (
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	reg = prometheus.NewRegistry()

	grpcMetrics = grpc_prometheus.NewServerMetrics()

	info = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "grpcMetrics",
		Help: "help",
	}, []string{"name"})
)

func init() {
	reg.MustRegister(grpcMetrics, info)
	info.WithLabelValues("Test")
}

func main() {
	ctx := context.Background()

	db, err := utils.InitPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres: %s", err.Error())
	}

	dbAdmin, err := utils.InitPostgresAdmin(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres admin: %s", err.Error())
	}

	sessionStore, err := utils.InitRedis()
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
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)

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
