package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/keepalive"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/interceptors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/server"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
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

	db_admin, err := utils.InitPostgresAdmin(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres admin: %s", err.Error())
	}

	db_umlaut, err := utils.InitPostgres(ctx)
	if err != nil {
		log.Fatalf("failed to initialize Postgres umlaut: %s", err.Error())
	}

	sessionStore, err := utils.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize redisDb: %s", err.Error())
	}
	defer sessionStore.Close()

	adminService := service.NewAdminService(repository.NewAdminPostgres(db_admin))
	complaintService := service.NewComplaintService(repository.NewComplaintPostgres(db_umlaut))

	adminServer := server.NewAdminServer(adminService, complaintService)
	viper.GetString("admin.port")
	listen, err := net.Listen("tcp", ":"+viper.GetString("admin.port"))
	if err != nil {
		log.Fatalf("Cannot listen port: %s. Err: %s", viper.GetString("admin.port"), err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc.UnaryServerInterceptor(interceptors.PanicRecoveryInterceptor),
			grpc.UnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)

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
