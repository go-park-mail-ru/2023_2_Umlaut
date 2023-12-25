package repository

import (
	"context"
	"errors"
	"fmt"
	static2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/stretchr/testify/assert"

	"github.com/pashagolub/pgxmock/v3"
)

func TestComplaintPostgres_CreateComplaint(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	complaintRepo := NewComplaintPostgres(mock)

	complaintText := "complaint"
	testComplaint := core2.Complaint{
		ReporterUserId:  1,
		ReportedUserId:  2,
		ComplaintTypeId: 1,
		ComplaintText:   &complaintText,
	}

	mock.ExpectQuery(`INSERT INTO "complaint"`).
		WithArgs(testComplaint.ReporterUserId, testComplaint.ReportedUserId, testComplaint.ComplaintTypeId, testComplaint.ComplaintText).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

	createdID, err := complaintRepo.CreateComplaint(context.Background(), testComplaint)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdID)

	mock.ExpectQuery(`INSERT INTO "complaint"`).
		WithArgs(testComplaint.ReporterUserId, testComplaint.ReportedUserId, testComplaint.ComplaintTypeId, testComplaint.ComplaintText).
		WillReturnError(static2.ErrAlreadyExists)

	_, err = complaintRepo.CreateComplaint(context.Background(), testComplaint)

	assert.ErrorIs(t, err, static2.ErrAlreadyExists)

	mock.ExpectQuery(`INSERT INTO "complaint"`).
		WithArgs(testComplaint.ReporterUserId, testComplaint.ReportedUserId, testComplaint.ComplaintTypeId, testComplaint.ComplaintText).
		WillReturnError(errors.New("some other error"))

	_, err = complaintRepo.CreateComplaint(context.Background(), testComplaint)

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestComplaintPostgres_GetComplaintTypes(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	complaintRepo := NewComplaintPostgres(mock)

	testComplaint := []core2.ComplaintType{
		{
			Id:       1,
			TypeName: "type1",
		},
		{
			Id:       2,
			TypeName: "type2",
		},
	}

	mock.ExpectQuery(fmt.Sprintf(`SELECT %s FROM "complaint_type"`, static2.ComplaintTypeDbFiend)).
		WillReturnRows(pgxmock.NewRows([]string{"id", "type_name"}).
			AddRow(1, "type1").AddRow(2, "type2"))

	complaintTypes, err := complaintRepo.GetComplaintTypes(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, complaintTypes, testComplaint)

	mock.ExpectQuery(fmt.Sprintf(`SELECT %s FROM "complaint_type"`, static2.ComplaintTypeDbFiend)).
		WillReturnError(errors.New("some other error"))

	_, err = complaintRepo.GetComplaintTypes(context.Background())

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestComplaintPostgres_DeleteComplaint(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	complaintRepo := NewComplaintPostgres(mock)

	mock.ExpectQuery(`DELETE FROM "complaint"`).
		WithArgs(1).
		WillReturnRows(pgxmock.NewRows([]string{"DELETE", "id"}).
			AddRow("DELETE", 1))

	err = complaintRepo.DeleteComplaint(context.Background(), 1)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestComplaintPostgres_AcceptComplaint(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	complaintRepo := NewComplaintPostgres(mock)

	testComplaintText := "complaint"

	mock.ExpectQuery(`UPDATE "complaint"`).
		WithArgs(1, 1).
		WillReturnRows(pgxmock.NewRows([]string{"id", "reporter_user_id", "reported_user_id", "complaint_type_id", "complaint_text", "created_at"}).
			AddRow(1, 1, 2, 1, &testComplaintText, &time.Time{}))

	_, err = complaintRepo.AcceptComplaint(context.Background(), 1)

	assert.NoError(t, err)

	mock.ExpectQuery(fmt.Sprintf(`SELECT %s FROM "complaint_type"`, static2.ComplaintTypeDbFiend)).
		WillReturnError(errors.New("some other error"))

	_, err = complaintRepo.GetComplaintTypes(context.Background())

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestComplaintPostgres_GetNextComplaint(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	complaintRepo := NewComplaintPostgres(mock)

	testComplaintText := "complaint"

	testComplaint := core2.Complaint{
		Id:              1,
		ReporterUserId:  1,
		ReportedUserId:  2,
		ComplaintTypeId: 1,
		ComplaintText:   &testComplaintText,
		CreatedAt:       &time.Time{},
	}

	mock.ExpectQuery(`SELECT`).
		WithArgs(0).
		WillReturnRows(pgxmock.NewRows([]string{"id", "reporter_user_id", "reported_user_id", "complaint_type_id", "complaint_text", "created_at"}).
			AddRow(1, 1, 2, 1, &testComplaintText, &time.Time{}))

	complaint, _ := complaintRepo.GetNextComplaint(context.Background())
	assert.Equal(t, testComplaint, complaint)

	mock.ExpectQuery(`SELECT`).
		WithArgs(0).
		WillReturnError(pgx.ErrNoRows)

	_, err = complaintRepo.GetNextComplaint(context.Background())

	assert.ErrorIs(t, err, static2.ErrNoData)

	mock.ExpectQuery(`SELECT`).
		WithArgs(0).
		WillReturnError(errors.New("some error"))

	_, err = complaintRepo.GetNextComplaint(context.Background())

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
