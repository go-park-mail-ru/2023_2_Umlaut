package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/jackc/pgx/v5"
)

type LikePostgres struct {
	db *pgx.Conn
}

func NewLikePostgres(db *pgx.Conn) *LikePostgres {
	return &LikePostgres{db: db}
}

func (r *LikePostgres) CreateLike(ctx context.Context, like model.Like) (model.Like, error) {
	query, args, err := psql.Insert(likeTable).
		Columns("liked_by_user_id", "liked_to_user_id").
		Values(like.LikedByUserId, like.LikedToUserId).
		ToSql()

	if err != nil {
		return model.Like{}, err
	}

	query += " RETURNING *"
	var newLike model.Like
	row := r.db.QueryRow(ctx, query, args...)
	err = scanLike(row, &newLike)

	return newLike, err
}

func (r *LikePostgres) Exists(ctx context.Context, like model.Like) (bool, error) {
	query, args, err := psql.Select("*").
		From(likeTable).
		Where(sq.Eq{"liked_by_user_id": like.LikedByUserId, "liked_to_user_id": like.LikedToUserId}).
		ToSql()
	if err != nil {
		return false, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanLike(row, &like)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func scanLike(row pgx.Row, like *model.Like) error {
	err := row.Scan(
		&like.LikedByUserId,
		&like.LikedToUserId,
		&like.CommittedAt,
	)
	return err
}
