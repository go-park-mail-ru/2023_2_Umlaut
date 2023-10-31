package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const (
	userTable = "\"user\""
	likeTable = "\"like\""
	dialogTable = "\"dialog\""
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
