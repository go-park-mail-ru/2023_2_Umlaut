package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
	if dialog.User1Id > dialog.User2Id {
		tmp := dialog.User1Id
		dialog.User1Id = dialog.User2Id
		dialog.User2Id = tmp
	}

	var id int
	query, args, err := psql.Insert(dialogTable).
		Columns("user1_id", "user2_id").
		Values(dialog.User1Id, dialog.User2Id).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create dialog. err: %w", err)
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, model.AlreadyExists
		}
	}
	return id, err
}

func (r *DialogPostgres) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	query, args, err := psql.Select(dialogTable+".id", "user1_id", "user2_id", userTable+".name").From(dialogTable).
		InnerJoin(userTable + " ON " +
			dialogTable + ".user1_id =" + userTable + ".id OR " +
			dialogTable + ".user2_id =" + userTable + ".id").
		Where(sq.Or{
			sq.And{
				sq.Eq{"user1_id": userId},
				sq.NotEq{userTable + ".id": userId},
			},
			sq.And{
				sq.Eq{"user2_id": userId},
				sq.NotEq{userTable + ".id": userId},
			},
		}).
		ToSql()

	if err != nil {
		return []model.Dialog{}, fmt.Errorf("failed to get dialog for userId %d. err: %w", userId, err)
	}

	rows, err := r.db.Query(ctx, query, args...)

	if err != nil {
		return []model.Dialog{}, fmt.Errorf("failed to get dialog for userId %d. err: %w", userId, err)
	}
	defer rows.Close()
	var dialogs []model.Dialog
	for rows.Next() {
		var dialog model.Dialog
		err = scanDialog(rows, &dialog)
		if errors.Is(err, pgx.ErrNoRows) {
			return dialogs, fmt.Errorf("dialogs doesn't exists for userId %d", userId)
		}
		if dialog.User2Id == userId {
			dialog.User2Id = dialog.User1Id
			dialog.User1Id = userId
		}
		dialogs = append(dialogs, dialog)
	}
	if err = rows.Err(); err != nil {
		return dialogs, fmt.Errorf("failed to get dialog for userId %d. err: %w", userId, err)
	}

	return dialogs, nil
}

func scanDialog(row pgx.Row, dialog *model.Dialog) error {
	err := row.Scan(
		&dialog.Id,
		&dialog.User1Id,
		&dialog.User2Id,
		&dialog.Сompanion,
	)
	return err
}
