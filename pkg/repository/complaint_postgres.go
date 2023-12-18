package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type ComplaintPostgres struct {
	db PgxPoolInterface
}

func NewComplaintPostgres(db PgxPoolInterface) *ComplaintPostgres {
	return &ComplaintPostgres{db: db}
}

func (r *ComplaintPostgres) GetComplaintTypes(ctx context.Context) ([]core.ComplaintType, error) {
	query, args, err := psql.
		Select(constants.ComplaintTypeDbFiend).
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

func (r *ComplaintPostgres) CreateComplaint(ctx context.Context, complaint core.Complaint) (int, error) {
	var id int
	query, args, err := psql.Insert(complaintTable).
		Columns("reporter_user_id",
			"reported_user_id",
			"complaint_type_id",
			"complaint_text").
		Values(complaint.ReporterUserId,
			complaint.ReportedUserId,
			complaint.ComplaintTypeId,
			complaint.ComplaintText).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create complaint. err: %w", err)
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

func (r *ComplaintPostgres) GetNextComplaint(ctx context.Context) (core.Complaint, error) {
	query, args, err := psql.Select(constants.ComplaintDbFiend).
		From(complaintTable).
		Where(sq.Eq{"report_status": 0}).
		Limit(1).ToSql()
	if err != nil {
		return core.Complaint{}, fmt.Errorf("failed to get next complaint. err: %w", err)
	}
	var nextComplaint core.Complaint
	row := r.db.QueryRow(ctx, query, args...)
	err = scanComplaint(row, &nextComplaint)

	if errors.Is(err, pgx.ErrNoRows) || nextComplaint.Id == 0 {
		return core.Complaint{}, constants.ErrNoData
	}

	return nextComplaint, err
}

func (r *ComplaintPostgres) DeleteComplaint(ctx context.Context, complaintId int) error {
	query, args, err := psql.Delete(complaintTable).
		Where(sq.Eq{"id": complaintId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to delete complaint. err: %w", err)
	}

	r.db.QueryRow(ctx, query, args...)

	return nil
}

func (r *ComplaintPostgres) AcceptComplaint(ctx context.Context, complaintId int) (core.Complaint, error) {
	query, args, err := psql.Update(complaintTable).
		Set("report_status", 1).
		Where(sq.Eq{"id": complaintId}).
		ToSql()

	if err != nil {
		return core.Complaint{}, err
	}

	query += fmt.Sprintf(" RETURNING %s", constants.ComplaintDbFiend)
	var updatedComplaint core.Complaint
	row := r.db.QueryRow(ctx, query, args...)
	err = scanComplaint(row, &updatedComplaint)

	return updatedComplaint, err
}

func scanComplaint(row pgx.Row, complaint *core.Complaint) error {
	err := row.Scan(
		&complaint.Id,
		&complaint.ReporterUserId,
		&complaint.ReportedUserId,
		&complaint.ComplaintTypeId,
		&complaint.ComplaintText,
		&complaint.CreatedAt,
	)

	return err
}

func scanComplaintTypes(rows pgx.Rows) ([]core.ComplaintType, error) {
	var complaintTypes []core.ComplaintType
	var err error
	for rows.Next() {
		var complaintType core.ComplaintType
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
