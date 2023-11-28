package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessagePostgres struct {
	db *pgxpool.Pool
}

func NewMessagePostgres(db *pgxpool.Pool) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) CreateMessage(ctx context.Context, message model.Message) (int, error) {
	var id int
	query, args, err := psql.Insert(messageTable).
		Columns("sender_id", "dialog_id", "message_text").
		Values(message.SenderId, message.DialogId, message.Text).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *MessagePostgres) UpdateMessage(ctx context.Context, message model.Message) (int, error) {
	var id int
	query, args, err := psql.Update(messageTable).
		Set("message_text", message.Text).
		Set("is_read", message.IsRead).
		Where(sq.Eq{"id": message.Id}).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	return id, err
}

func (r *MessagePostgres) GetDialogMessages(ctx context.Context, dialogId int) ([]model.Message, error) {
	queryBuilder := psql.
		Select(static.MessageDbField).
		From(messageTable).
		Where(sq.Eq{"dialog_id": dialogId}).
		OrderBy("created_at")
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

func scanMessages(rows pgx.Rows) ([]model.Message, error) {
	var messages []model.Message
	var err error
	for rows.Next() {
		var message model.Message
		err = rows.Scan(
			&message.Id,
			&message.DialogId,
			&message.SenderId,
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
