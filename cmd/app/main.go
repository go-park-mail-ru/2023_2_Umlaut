package main

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host umlaut-bmstu.me:8000
// @BasePath /
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close(context.Background())

	redisDb, err := strconv.Atoi(viper.GetString("redis.db"))
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	sessionStore, err := repository.NewRedisClient(repository.RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: "", //os.Getenv("DB_PASSWORD"),
		DB:       redisDb,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer sessionStore.Close()

	fileClient, err := repository.NewMinioClient(repository.MinioConfig{
		User:     viper.GetString("minio.user"),
		Password: viper.GetString("minio.password"),
		SSLMode:  viper.GetBool("minio.sslmode"),
		Endpoint: viper.GetString("minio.endpoint"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db, sessionStore, fileClient)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(umlaut.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
