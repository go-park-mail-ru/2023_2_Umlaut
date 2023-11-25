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

func (r *AdminPostgres) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	var id int
	query, args, err := psql.Insert(feedbackTable).
		Columns("user_id", "rating", "liked", "need_fix", "comment_fix", "comment", "show").
		Values(stat.UserId, stat.Rating, stat.Liked, stat.NeedFix, stat.CommentFix, stat.Comment, stat.Show).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create statistic. err: %w", err)
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	var id int
	query, args, err := psql.Insert(recommendationTable).
		Columns("user_id", "recommend", "show").
		Values(rec.UserId, rec.Recommend, rec.Show).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *DialogPostgres) GetFeedbacks(ctx context.Context) ([]model.Feedback, error) {
	query, args, err := psql.
		Select("d.id", "d.user1_id", "d.user2_id", "u.name", "u.image_paths", "m.id", "m.sender_id", "m.dialog_id", "m.message_text", "m.created_at").
		From(dialogTable + " d").
		LeftJoin(fmt.Sprintf("%s u on d.user1_id = u.id or d.user2_id = u.id", userTable)).
		LeftJoin(fmt.Sprintf("%s m ON d.last_message_id = m.id", messageTable)).
		Where(sq.And{
			sq.Or{sq.Eq{"d.user1_id": userId}, sq.Eq{"d.user2_id": userId}},
			sq.NotEq{"u.id": userId},
		}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get dialog for userId %d. err: %w", userId, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get dialog for userId %d. err: %w", userId, err)
	}
	defer rows.Close()

	dialogs, err := scanDialogs(rows)
	if err != nil {
		return nil, err
	}

	return dialogs, nil
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
