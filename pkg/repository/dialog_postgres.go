package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
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
			return 0, static.ErrAlreadyExists
		}
	}
	return id, err
}

func (r *DialogPostgres) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	query, args, err := psql.
		Select("d.id", "d.user1_id", "d.user2_id", "u.name", "m.id", "m.sender_id", "m.dialog_id", "m.message_text", "m.timestamp").
		From(dialogTable + " d").
		InnerJoin(fmt.Sprintf("%s u ON d.user1_id = u.id OR d.user2_id = u.id", userTable)).
		InnerJoin(fmt.Sprintf("%s m ON d.last_message_id = m.id", messageTable)).
		Where(sq.Or{sq.Eq{"user1_id": userId}, sq.Eq{"user2_id": userId}}).
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

func (r *DialogPostgres) GetDialogMessages(ctx context.Context, dialogId int) ([]model.Message, error) {
	queryBuilder := psql.
		Select("*").
		From(messageTable).
		Where(sq.Eq{"dialog_id": dialogId}).
		OrderBy("timestamp desc")
	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get message for dialogId %d. err: %w", dialogId, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get message for dialogId %d. err: %w", dialogId, err)
	}
	defer rows.Close()

	messages, err := scanMessages(rows)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func scanDialogs(rows pgx.Rows) ([]model.Dialog, error) {
	var dialogs []model.Dialog
	var err error
	for rows.Next() {
		var dialog model.Dialog
		//var lastMessage model.Message
		err = rows.Scan(
			&dialog.Id,
			&dialog.User1Id,
			&dialog.User2Id,
			&dialog.Сompanion,
			&dialog.LastMessage,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		//TODO: check scan last msg
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

func scanMessages(rows pgx.Rows) ([]model.Message, error) {
	var messages []model.Message
	var err error
	for rows.Next() {
		var message model.Message
		err = rows.Scan(
			&message.Id,
			&message.SenderId,
			&message.DialogId,
			&message.MessageText,
			&message.TimeStamp,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		messages = append(messages, message)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetDialogMessages: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetDialogMessages: %v", rows.Err())
	}

	return messages, nil
}
