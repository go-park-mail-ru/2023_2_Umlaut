package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/jackc/pgx/v5"
)

type AdminPostgres struct {
	db PgxPoolInterface
}

func NewAdminPostgres(db PgxPoolInterface) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) GetAdmin(ctx context.Context, mail string) (core.Admin, error) {
	var admin core.Admin

	query, args, err := psql.Select(constants.AdminDbField).From(adminTable).Where(sq.Eq{"mail": mail}).ToSql()

	if err != nil {
		return admin, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanAdmin(row, &admin)

	if errors.Is(err, pgx.ErrNoRows) {
		return core.Admin{}, fmt.Errorf("admin with mail: %s not found", mail)
	}

	return admin, err
}

func (r *AdminPostgres) ShowFeedback(ctx context.Context, userId int) (bool, error) {
	var id int

	query, args, err := psql.Select("id").From(feedbackTable).Where(
		sq.And{
			sq.Lt{"EXTRACT(DAY FROM NOW()-created_at)": "7"},
			sq.Eq{"user_id": userId},
		}).ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to check can show feedback. err: %w", err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	if errors.Is(err, pgx.ErrNoRows) {
		return true, nil
	}

	return false, err
}

func (r *AdminPostgres) ShowRecommendation(ctx context.Context, userId int) (bool, error) {
	var id int

	query, args, err := psql.Select("id").From(recommendationTable).Where(
		sq.And{
			sq.Lt{"EXTRACT(DAY FROM NOW()-created_at)": "7"},
			sq.Eq{"user_id": userId},
		}).ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to check can show recommendation. err: %w", err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	if errors.Is(err, pgx.ErrNoRows) {
		return true, nil
	}

	return false, err
}

func (r *AdminPostgres) CreateFeedback(ctx context.Context, stat core.Feedback) (int, error) {
	var id int
	query, args, err := psql.Insert(feedbackTable).
		Columns("user_id", "rating", "liked", "need_fix", "comment").
		Values(stat.UserId, stat.Rating, stat.Liked, stat.NeedFix, stat.Comment).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create statistic. err: %w", err)
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) CreateRecommendation(ctx context.Context, rec core.Recommendation) (int, error) {
	var id int
	query, args, err := psql.Insert(recommendationTable).
		Columns("user_id", "rating").
		Values(rec.UserId, rec.Rating).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) CreateFeedFeedback(ctx context.Context, rec core.Recommendation) (int, error) {
	var id int
	query, args, err := psql.Insert(feedFeedbackTable).
		Columns("user_id", "rating").
		Values(rec.UserId, rec.Rating).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *AdminPostgres) GetFeedbacks(ctx context.Context) ([]core.Feedback, error) {
	query, args, err := psql.
		Select(constants.FeedbackDbField).
		From(feedbackTable).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}
	defer rows.Close()

	feedbacks, err := scanFeedbacks(rows)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (r *AdminPostgres) GetRecommendations(ctx context.Context) ([]core.Recommendation, error) {
	query, args, err := psql.
		Select("id", "user_id", "rating").
		From(recommendationTable).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get feedback. err: %w", err)
	}
	defer rows.Close()

	recommendations, err := scanRecommendations(rows)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func scanAdmin(row pgx.Row, admin *core.Admin) error {
	err := row.Scan(
		&admin.Id,
		&admin.Mail,
		&admin.PasswordHash,
		&admin.Salt,
	)

	return err
}

func scanFeedbacks(rows pgx.Rows) ([]core.Feedback, error) {
	var feedbacks []core.Feedback
	var err error
	for rows.Next() {
		var feedback core.Feedback
		err = rows.Scan(
			&feedback.Id,
			&feedback.UserId,
			&feedback.Rating,
			&feedback.Liked,
			&feedback.NeedFix,
			&feedback.Comment,
			&feedback.CreatedAt,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		feedbacks = append(feedbacks, feedback)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetFeedbacks: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetFeedbacks: %v", rows.Err())
	}

	return feedbacks, nil
}
func scanRecommendations(rows pgx.Rows) ([]core.Recommendation, error) {
	var recommendations []core.Recommendation
	var err error
	for rows.Next() {
		var feedback core.Recommendation
		err = rows.Scan(
			&feedback.Id,
			&feedback.UserId,
			&feedback.Rating,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		recommendations = append(recommendations, feedback)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetFeedbacks: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetFeedbacks: %v", rows.Err())
	}

	return recommendations, nil
}
