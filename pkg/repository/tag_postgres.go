package repository

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TagPostgres struct {
	db *pgxpool.Pool
}

func NewTagPostgres(db *pgxpool.Pool) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) GetAllTags(ctx context.Context) ([]model.Tag, error) {
	var tags []model.Tag

	query, args, err := psql.Select("*").From(tagTable).ToSql()
	if err != nil {
		return tags, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return tags, fmt.Errorf("query error in GetAllTags: %v", err)
	}
	defer rows.Close()

	err = scanTags(rows, &tags)
	if err != nil {
		return tags, err
	}

	return tags, err
}

func (r *TagPostgres) GetUserTags(ctx context.Context, userId int) ([]model.Tag, error) {
	var tags []model.Tag

	query, args, err := psql.Select("id", "name").
		From(tagTable).
		Join(fmt.Sprintf("%s on %s.id = %s.tag_id", userTagTable, tagTable, userTagTable)).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return tags, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return tags, fmt.Errorf("query error in GetAllTags: %v", err)
	}
	defer rows.Close()

	err = scanTags(rows, &tags)
	if err != nil {
		return tags, err
	}

	return tags, err
}

func scanTags(rows pgx.Rows, tags *[]model.Tag) error {
	var err error
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		*tags = append(*tags, model.Tag{Id: id, Name: name})
	}
	if err != nil {
		return fmt.Errorf("scan error in GetAllTags: %v", err)
	}
	if rows.Err() != nil {
		return fmt.Errorf("rows error in GetAllTags: %v", rows.Err())
	}

	return nil
}
