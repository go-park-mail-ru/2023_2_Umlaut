package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
)

type AdminPostgres struct {
	db *pgxpool.Pool
}

func NewAdminPostgres(db *pgxpool.Pool) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) GetAdmin(ctx context.Context, mail string) (model.Admin, error) {
	var admin model.Admin

	query, args, err := psql.Select(static.AdminDbField).From(adminTable).Where(sq.Eq{"mail": mail}).ToSql()

	if err != nil {
		return admin, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanAdmin(row, &admin)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Admin{}, fmt.Errorf("admin with mail: %s not found", mail)
	}

	return admin, err
}

func (r *AdminPostgres) CreateStatistic(ctx context.Context, stat model.Statistic) (int, error) {
	var id int
	query, args, err := psql.Insert(statisticTable).
		Columns("user_id", "rating", "liked", "need_fix", "comment_fix", "comment", "created_at").
		Values(stat.UserId, stat.Rating, stat.Liked, stat.NeedFix, stat.CommentFix, stat.Comment, stat.CreatedAt).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create statistic. err: %w", err)
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func scanAdmin(row pgx.Row, admin *model.Admin) error {
	err := row.Scan(
		&admin.Id,
		&admin.Mail,
		&admin.PasswordHash,
		&admin.Salt,
	)

	return err
}
