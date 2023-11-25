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
		Host:     viper.GetString("db_umlaut.host"),
		Port:     viper.GetString("db_umlaut.port"),
		Username: viper.GetString("db_umlaut.username"),
		DBName:   viper.GetString("db_umlaut.dbname"),
		SSLMode:  viper.GetString("db_umlaut.sslmode"),
		Password: viper.GetString("db_umlaut.password"), //os.Getenv("DB_PASSWORD"),
	})
}

func InitPostgresAdmin(ctx context.Context) (*pgxpool.Pool, error) {
	return repository.NewPostgresDB(ctx, repository.PostgresConfig{
		Host:     viper.GetString("db_admin.host"),
		Port:     viper.GetString("db_admin.port"),
		Username: viper.GetString("db_admin.username"),
		DBName:   viper.GetString("db_admin.dbname"),
		SSLMode:  viper.GetString("db_admin.sslmode"),
		Password: viper.GetString("db_admin.password"), //os.Getenv("DB_PASSWORD"),
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
