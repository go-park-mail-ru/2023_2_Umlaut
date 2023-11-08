package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jackc/pgx/v5"
)

type TagPostgres struct {
	db *pgxpool.Pool
}

func NewTagPostgres(db *pgxpool.Pool) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) GetAllTags(ctx context.Context) ([]string, error) {
	var tags []string

	query, args, err := psql.Select("name").From(tagTable).ToSql()
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

func scanTags(rows pgx.Rows, tags *[]string) error {
	var err error
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		*tags = append(*tags, name)
	}
	if err != nil {
		return fmt.Errorf("scan error in GetAllTags: %v", err)
	}
	if rows.Err() != nil {
		return fmt.Errorf("rows error in GetAllTags: %v", rows.Err())
	}

	return nil
}
