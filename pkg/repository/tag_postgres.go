package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/jackc/pgx/v5"
)

type TagPostgres struct {
	db mock_repository.PgxPoolInterface
}

func NewTagPostgres(db mock_repository.PgxPoolInterface) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) GetAllTags(ctx context.Context) ([]string, error) {
	query, args, err := psql.Select("name").From(tagTable).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error in GetAllTags: %v", err)
	}
	defer rows.Close()

	tags, err := scanTags(rows)
	if err != nil {
		return tags, err
	}

	return tags, err
}

func scanTags(rows pgx.Rows) ([]string, error) {
	var tags []string
	var err error
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		tags = append(tags, name)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetAllTags: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetAllTags: %v", rows.Err())
	}

	return tags, nil
}
