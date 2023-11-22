package utils

import (
	"context"
	"log"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
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

func InitLogger() (*zap.Logger, error) {
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

func InitPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return repository.NewPostgresDB(ctx, repository.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		Password: viper.GetString("postgres.password"), //os.Getenv("DB_PASSWORD"),
	})
}

func InitRedis() (*redis.Client, error) {
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

func InitMinioClient() (*minio.Client, error) {
	return repository.NewMinioClient(repository.MinioConfig{
		User:     viper.GetString("minio.user"),
		Password: viper.GetString("minio.password"),
		SSLMode:  viper.GetBool("minio.sslmode"),
		Endpoint: viper.GetString("minio.endpoint"),
	})
}
