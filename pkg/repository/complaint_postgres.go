package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
)

type ComplaintPostgres struct {
	db *pgxpool.Pool
}

func NewComplaintPostgres(db *pgxpool.Pool) *ComplaintPostgres {
	return &ComplaintPostgres{db: db}
}

func (r *ComplaintPostgres) GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error) {
	query, args, err := psql.
		Select(static.ComplaintTypeDbFiend).
		From(complaintTypeTable).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to get complaint types. err: %w", err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get complaint types. err: %w", err)
	}
	defer rows.Close()

	complaintTypes, err := scanComplaintTypes(rows)
	if err != nil {
		return nil, err
	}

	return complaintTypes, nil
}

// func (r *ComplaintPostgres) CreateComplaint(ctx context.Context, complaint model.Complaint) (int, error) {
// 	var id int
// 	query, args, err := psql.Insert(complaintTable).
// 		Columns("reporter_user_id", "reported_user_id", "complaint_type_id").
// 		Values(complaint.ReporterUserId, complaint.ReportedUserId, complaint.ComplaintTypeId).
// 		ToSql()

// 	if err != nil {
// 		return 0, fmt.Errorf("failed to create complaint. err: %w", err)
// 	}

// 	query += " RETURNING id"
// 	row := r.db.QueryRow(ctx, query, args...)
// 	err = row.Scan(&id)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
// 			return 0, static.ErrAlreadyExists
// 		}
// 	}
// 	return id, err
// }

func scanComplaintTypes(rows pgx.Rows) ([]model.ComplaintType, error) {
	var complaintTypes []model.ComplaintType
	var err error
	for rows.Next() {
		var complaintType model.ComplaintType
		err = rows.Scan(
			&complaintType.Id,
			&complaintType.TypeName,
		)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		complaintTypes = append(complaintTypes, complaintType)
	}
	if err != nil {
		return nil, fmt.Errorf("scan error in GetComplaintTypes: %v", err)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error in GetComplaintTypes: %v", rows.Err())
	}

	return complaintTypes, nil
}
