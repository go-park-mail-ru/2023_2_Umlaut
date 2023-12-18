package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type DialogPostgres struct {
	db PgxPoolInterface
}

func NewDialogPostgres(db PgxPoolInterface) *DialogPostgres {
	return &DialogPostgres{db: db}
}

func (r *DialogPostgres) CreateDialog(ctx context.Context, dialog core.Dialog) (int, error) {
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
			return 0, constants.ErrAlreadyExists
		}
	}
	return id, err
}

func (r *DialogPostgres) GetDialogs(ctx context.Context, userId int) ([]core.Dialog, error) {
	query, args, err := psql.
		Select("d.id", "d.user1_id", "d.user2_id", "d.banned", "u.name", "u.image_paths", "m.id", "m.sender_id", "m.recipient_id", "m.dialog_id", "m.message_text", "m.is_read", "m.created_at").
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

	dialogs, err := scanDialogs(rows, userId)
	if err != nil {
		return nil, err
	}

	return dialogs, nil
}

func (r *DialogPostgres) GetDialogById(ctx context.Context, id int) (core.Dialog, error) {
	query, args, err := psql.
		Select("d.id", "d.user1_id", "d.user2_id", "d.banned", "u.name", "u.image_paths", "m.id", "m.sender_id", "m.recipient_id", "m.dialog_id", "m.message_text", "m.is_read", "m.created_at").
		From(dialogTable + " d").
		LeftJoin(fmt.Sprintf("%s u on d.user1_id = u.id or d.user2_id = u.id", userTable)).
		LeftJoin(fmt.Sprintf("%s m ON d.last_message_id = m.id", messageTable)).
		Where(sq.Eq{"d.id": id}).
		ToSql()

	if err != nil {
		return core.Dialog{}, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return core.Dialog{}, fmt.Errorf("failed to get dialog with id %d. err: %w", id, err)
	}
	defer rows.Close()
	dialog, err := scanDialogs(rows, ctx.Value(constants.KeyUserID).(int))

	if errors.Is(err, pgx.ErrNoRows) {
		return core.Dialog{}, fmt.Errorf("dialog with id: %d not found", id)
	}

	return dialog[0], err
}

func scanDialogs(rows pgx.Rows, userId int) ([]core.Dialog, error) {
	var dialogs []core.Dialog
	var err error
	for rows.Next() {
		var dialog core.Dialog
		var lastMessage core.Message
		err = rows.Scan(
			&dialog.Id,
			&dialog.User1Id,
			&dialog.User2Id,
			&dialog.Banned,
			&dialog.Companion,
			&dialog.CompanionImagePaths,
			&lastMessage.Id,
			&lastMessage.SenderId,
			&lastMessage.RecipientId,
			&lastMessage.DialogId,
			&lastMessage.Text,
			&lastMessage.IsRead,
			&lastMessage.CreatedAt,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		if lastMessage.Id == nil {
			dialog.LastMessage = nil
		} else {
			dialog.LastMessage = &lastMessage
		}

		if dialog.User1Id == userId {
			tmp := dialog.User1Id
			dialog.User1Id = dialog.User2Id
			dialog.User2Id = tmp
		}

		dialogs = append(dialogs, dialog)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetDialogs: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetDialogs: %v", rows.Err())
	}

	return dialogs, nil
}
