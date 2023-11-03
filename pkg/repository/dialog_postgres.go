package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DialogPostgres struct {
	db *pgxpool.Pool
}

func NewDialogPostgres(db *pgxpool.Pool) *DialogPostgres {
	return &DialogPostgres{db: db}
}

func (r *DialogPostgres) CreateDialog(ctx context.Context, dialog model.Dialog) (int, error) {
	var id int
	query, args, err := psql.Insert(dialogTable).
		Columns("user1_id", "user2_id").
		Values(dialog.User1Id, dialog.User2Id).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *DialogPostgres) Exists(ctx context.Context, dialog model.Dialog) (bool, error) {
	query, args, err := psql.Select("*").
		From(dialogTable).
		Where(sq.Or{
			sq.Eq{"user1_id": dialog.User1Id, "user2_id": dialog.User2Id},
			sq.Eq{"user1_id": dialog.User2Id, "user2_id": dialog.User1Id}}).
		ToSql()
	if err != nil {
		return false, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanDialog(row, &dialog)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *DialogPostgres) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	query, args, err := psql.Select("*").
		From(dialogTable).
		Where(sq.Or{
			sq.Eq{"user1_id": userId},
			sq.Eq{"user2_id": userId}}).
		ToSql()
	if err != nil {
		return []model.Dialog{}, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return []model.Dialog{}, err
	}
	defer rows.Close()
	var dialogs []model.Dialog
	for rows.Next() {
		var dialog model.Dialog
		err = scanDialog(rows, &dialog)
		if errors.Is(err, pgx.ErrNoRows) {
			return dialogs, nil
		}
		dialogs = append(dialogs, dialog)
	}
	if err = rows.Err(); err != nil {
		return dialogs, err
	}

	return dialogs, nil
}

func scanDialog(row pgx.Row, dialog *model.Dialog) error {
	err := row.Scan(
		&dialog.Id,
		&dialog.User1Id,
		&dialog.User2Id,
	)
	return err
}
