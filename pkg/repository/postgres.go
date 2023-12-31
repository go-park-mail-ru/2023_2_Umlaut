package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	userTable          = "\"user\""
	likeTable          = "\"like\""
	dialogTable        = "\"dialog\""
	tagTable           = "\"tag\""
	messageTable       = "\"message\""
	complaintTable     = "\"complaint\""
	complaintTypeTable = "\"complaint_type\""

	adminTable          = "\"admin\""
	feedbackTable       = "\"feedback\""
	recommendationTable = "\"recommendation\""
	feedFeedbackTable   = "\"feed_feedback\""
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(ctx context.Context, cfg PostgresConfig) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
