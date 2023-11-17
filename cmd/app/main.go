package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	umlaut "github.com/go-park-mail-ru/2023_2_Umlaut"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config_local")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host umlaut-bmstu.me:8000
// @BasePath /
func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting server...")
	defer logger.Sync()

	ctx := context.WithValue(context.Background(), "logger", logger)

	db, err := initPostgres(ctx)
	if err != nil {
		logger.Error("initialize Postgres",
			zap.String("Error", fmt.Sprintf("failed to initialize Postgres: %s", err.Error())))
	}

	sessionStore, err := initRedis()
	if err != nil {
		logger.Error("initialize redisDb",
			zap.String("Error", fmt.Sprintf("failed to initialize redisDb: %s", err.Error())))
	}
	defer sessionStore.Close()

	fileClient, err := initMinioClient()
	if err != nil {
		logger.Error("initialize Minio",
			zap.String("Error", fmt.Sprintf("failed to initialize Minio: %s", err.Error())))
	}

	repos := repository.NewRepository(db, sessionStore, fileClient)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, ctx)
	srv := new(umlaut.Server)

	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("running http server",
			zap.String("Error", fmt.Sprintf("error occured while running http server: %s", err.Error())))
	}
}

func initLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config.Build()
}

func initPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return repository.NewPostgresDB(ctx, repository.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		Password: viper.GetString("postgres.password"), //os.Getenv("DB_PASSWORD"),
	})
}

func initRedis() (*redis.Client, error) {
	redisDb, err := strconv.Atoi(viper.GetString("redis.db"))
	if err != nil {
		return nil, err
	}
	return repository.NewRedisClient(repository.RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: "", //os.Getenv("DB_PASSWORD"),
		DB:       redisDb,
	})
}

func initMinioClient() (*minio.Client, error) {
	return repository.NewMinioClient(repository.MinioConfig{
		User:     viper.GetString("minio.user"),
		Password: viper.GetString("minio.password"),
		SSLMode:  viper.GetBool("minio.sslmode"),
		Endpoint: viper.GetString("minio.endpoint"),
	})
}
