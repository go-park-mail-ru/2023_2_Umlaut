package main

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		Password: os.Getenv("PG_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	redisDb, err := strconv.Atoi(viper.GetString("redis.db"))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	sessionStore, err := repository.NewRedisClient(repository.RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDb,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer sessionStore.Close()

	repos := repository.NewRepository(db, sessionStore)
	handlers := handler.NewHandler(repos)

	srv := new(umlaut.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
