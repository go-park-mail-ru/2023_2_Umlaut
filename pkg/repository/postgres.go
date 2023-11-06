package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	userTable   = "\"user\""
	likeTable   = "\"like\""
	dialogTable = "\"dialog\""
	tagTable    = "\"tag\""
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	//Пробросить контекст
	if err != nil {
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
