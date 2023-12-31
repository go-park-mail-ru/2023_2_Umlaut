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

type MessagePostgres struct {
	db PgxPoolInterface
}

func NewMessagePostgres(db PgxPoolInterface) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) CreateMessage(ctx context.Context, message core.Message) (core.Message, error) {
	query, args, err := psql.Insert(messageTable).
		Columns("sender_id", "recipient_id", "dialog_id", "message_text").
		Values(message.SenderId, message.RecipientId, message.DialogId, message.Text).
		ToSql()

	if err != nil {
		return core.Message{}, err
	}

	query += fmt.Sprintf(" RETURNING %s", constants.MessageDbField)
	row := r.db.QueryRow(ctx, query, args...)
	newMessage, err := scanMessage(row)
	if err != nil {
		return core.Message{}, err
	}

	return newMessage, err
}

func (r *MessagePostgres) UpdateMessage(ctx context.Context, message core.Message) (core.Message, error) {
	query, args, err := psql.Update(messageTable).
		Set("message_text", message.Text).
		Set("is_read", message.IsRead).
		Where(sq.Eq{"id": message.Id}).
		ToSql()

	if err != nil {
		return core.Message{}, err
	}

	query += fmt.Sprintf(" RETURNING %s", constants.MessageDbField)
	row := r.db.QueryRow(ctx, query, args...)
	newMessage, err := scanMessage(row)
	if err != nil {
		return core.Message{}, err
	}

	return newMessage, err
}

func (r *MessagePostgres) GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core.Message, error) {
	queryBuilder := psql.
		Select(constants.MessageDbField).
		From(messageTable).
		Where(sq.And{
			sq.Or{sq.Eq{"sender_id": userId}, sq.Eq{"sender_id": recipientId}},
			sq.Or{sq.Eq{"recipient_id": userId}, sq.Eq{"recipient_id": recipientId}},
		}).
		OrderBy("created_at")
	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get message for users: %d, %d err: %w", userId, recipientId, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get message for users: %d, %d err: %w", userId, recipientId, err)
	}
	defer rows.Close()

	messages, err := scanMessages(rows)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func scanMessages(rows pgx.Rows) ([]core.Message, error) {
	var messages []core.Message
	var err error
	for rows.Next() {
		var message core.Message
		err = rows.Scan(
			&message.Id,
			&message.DialogId,
			&message.SenderId,
			&message.RecipientId,
			&message.Text,
			&message.IsRead,
			&message.CreatedAt,
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

func scanMessage(rows pgx.Row) (core.Message, error) {
	var message core.Message
	err := rows.Scan(
		&message.Id,
		&message.DialogId,
		&message.SenderId,
		&message.RecipientId,
		&message.Text,
		&message.IsRead,
		&message.CreatedAt,
	)
	return message, err
}
