package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"strings"

	sq "github.com/Masterminds/squirrel"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/jackc/pgx/v5"
)

type LikePostgres struct {
	db mock_repository.PgxPoolInterface
}

func NewLikePostgres(db mock_repository.PgxPoolInterface) *LikePostgres {
	return &LikePostgres{db: db}
}

func (r *LikePostgres) CreateLike(ctx context.Context, like model.Like) (model.Like, error) {
	query, args, err := psql.Insert(likeTable).
		Columns("liked_by_user_id", "liked_to_user_id", "is_like").
		Values(like.LikedByUserId, like.LikedToUserId, like.IsLike).
		ToSql()

	if err != nil {
		return model.Like{}, fmt.Errorf("failed to create like. err: %w", err)
	}

	query += fmt.Sprintf(" RETURNING %s", static.LikeDbField)
	var newLike model.Like
	row := r.db.QueryRow(ctx, query, args...)
	err = scanLike(row, &newLike)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return newLike, static.ErrAlreadyExists
		}
	}

	return newLike, err
}

func (r *LikePostgres) IsMutualLike(ctx context.Context, like model.Like) (bool, error) {
	tmp := like.LikedToUserId
	like.LikedToUserId = like.LikedByUserId
	like.LikedByUserId = tmp

	query, args, err := psql.Select(static.LikeDbField).
		From(likeTable).
		Where(sq.Eq{"liked_by_user_id": like.LikedByUserId, "liked_to_user_id": like.LikedToUserId, "is_like": like.IsLike}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to check is like exists. err: %w", err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanLike(row, &like)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return true, err
}

func scanLike(row pgx.Row, like *model.Like) error {
	err := row.Scan(
		&like.LikedByUserId,
		&like.LikedToUserId,
	)
	return err
}
