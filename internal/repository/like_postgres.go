package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	static2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/dto"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5"
)

type LikePostgres struct {
	db PgxPoolInterface
}

func NewLikePostgres(db PgxPoolInterface) *LikePostgres {
	return &LikePostgres{db: db}
}

func (r *LikePostgres) CreateLike(ctx context.Context, like core.Like) (core.Like, error) {
	query, args, err := psql.Insert(likeTable).
		Columns("liked_by_user_id", "liked_to_user_id", "is_like").
		Values(like.LikedByUserId, like.LikedToUserId, like.IsLike).
		ToSql()

	if err != nil {
		return core.Like{}, fmt.Errorf("failed to create like. err: %w", err)
	}

	query += fmt.Sprintf(" RETURNING %s", static2.LikeDbField)
	var newLike core.Like
	row := r.db.QueryRow(ctx, query, args...)
	err = scanLike(row, &newLike)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return newLike, static2.ErrAlreadyExists
		}
	}

	return newLike, err
}

func (r *LikePostgres) IsMutualLike(ctx context.Context, like core.Like) (bool, error) {
	tmp := like.LikedToUserId
	like.LikedToUserId = like.LikedByUserId
	like.LikedByUserId = tmp

	query, args, err := psql.Select(static2.LikeDbField).
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

func (r *LikePostgres) GetUserLikedToLikes(ctx context.Context, userId int) ([]dto.PremiumLike, error) {
	query, args, err := psql.Select("liked_by_user_id", "u.image_paths").
		From(likeTable).
		Join(userTable + " u ON liked_by_user_id = u.id").
		Where(sq.And{
			sq.Eq{"liked_to_user_id": userId},
			sq.Eq{"is_like": true},
		}).
		Where(fmt.Sprintf("liked_by_user_id NOT IN (SELECT l.liked_to_user_id FROM \"like\" l WHERE l.liked_by_user_id = %d)", userId)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get premium likes. err: %w", err)
	}
	
	var likes []dto.PremiumLike
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get premium likes. err: %w", err)
	}
	defer rows.Close()

	likes, err = scanPremiumLikes(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to get premium likes. err: %w", err)
	}

	return likes, nil
}

func (r *LikePostgres) ResetDislike(ctx context.Context) error {
	_, err := r.db.Exec(ctx, fmt.Sprintf("delete from %s where is_like = false;", likeTable))
	return err

}

func scanPremiumLikes(rows pgx.Rows) ([]dto.PremiumLike, error) {
	var likes []dto.PremiumLike
	var err error
	for rows.Next() {
		var like dto.PremiumLike
		err = rows.Scan(
			&like.LikedByUserId,
			&like.ImagePaths,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		likes = append(likes, like)
	}
	if err != nil {
		return nil, fmt.Errorf("scan like error: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows like error: %v", rows.Err())
	}

	return likes, nil
}

func scanLike(row pgx.Row, like *core.Like) error {
	err := row.Scan(
		&like.LikedByUserId,
		&like.LikedToUserId,
	)
	return err
}
