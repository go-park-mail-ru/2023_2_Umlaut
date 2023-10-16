package main

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

// @title Umlaut API
// @version 1.0
// @description API Server for Umlaut Application

// @host 37.139.32.76:8000
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
	defer db.Close()

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

	repos := repository.NewRepository(db, sessionStore)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(umlaut.Server)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("cors.origins"),
		AllowedMethods:   viper.GetStringSlice("cors.methods"),
		AllowCredentials: true,
	})

	if err := srv.Run(viper.GetString("port"), corsMiddleware.Handler(handlers.InitRoutes())); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
