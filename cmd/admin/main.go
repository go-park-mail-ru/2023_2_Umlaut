package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/keepalive"

	initial "github.com/go-park-mail-ru/2023_2_Umlaut/cmd"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/microservices/interceptors"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/microservices/admin/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	initial.InitConfig()
	grpc_prometheus.EnableHandlingTimeHistogram()
	ctx := context.Background()

	dbAdmin, err := initial.InitPostgresAdmin(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres admin: %s", err.Error())
	}

	dbUmlaut, err := initial.InitPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres umlaut: %s", err.Error())
	}

	sessionStore, err := initial.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize redisDb: %s", err.Error())
	}
	defer sessionStore.Close()

	adminService := service.NewAdminService(
		repository.NewAdminPostgres(dbAdmin),
		repository.NewUserPostgres(dbUmlaut),
	)
	complaintService := service.NewComplaintService(repository.NewComplaintPostgres(dbUmlaut))

	adminServer := server.NewAdminServer(adminService, complaintService)
	viper.GetString("admin.port")
	listen, err := net.Listen("tcp", ":"+viper.GetString("admin.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("admin.port"), err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.PanicRecoveryInterceptor,
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)

	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{
		Addr:    ":9093",
		Handler: nil,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Printf("Failed to start Prometheus metrics server: %s\n", err)
		}
	}()

	proto.RegisterAdminServer(grpcServer, adminServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("admin.port"), err.Error())
	}
}
