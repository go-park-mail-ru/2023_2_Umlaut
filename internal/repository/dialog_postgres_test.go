package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
)

func TestDialogPostgres_CreateDialog(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	dialogRepo := NewDialogPostgres(mock)

	testDialog := core.Dialog{
		User1Id: 1,
		User2Id: 2,
	}

	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

	createdID, err := dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdID)

	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnError(constants.ErrAlreadyExists)

	_, err = dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.ErrorIs(t, err, constants.ErrAlreadyExists)

	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnError(errors.New("some other error"))

	_, err = dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

//func TestDialogPostgres_GetDialogs(t *testing.T) {
//	mock, err := pgxmock.NewPool()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer mock.Close()
//
//	dialogRepo := NewDialogPostgres(mock)
//
//	id := 1
//	m := "Test message"
//	b := false
//	time := time2.Now()
//
//	mock.ExpectQuery(`SELECT`).WithArgs(id, id, id).
//		WillReturnRows(pgxmock.NewRows([]string{
//			"id", "user1_id", "user2_id", "banned", "name", "image_paths",
//			"message_id", "sender_id", "dialog_id", "message_text", "is_read", "created_at",
//		}).AddRow(1, 1, 2, false, "User2", nil, &id, &id, &id, &m, &b, &time))
//
//	dialogs, err := dialogRepo.GetDialogs(context.Background(), id)
//
//	assert.NoError(t, err)
//	assert.Len(t, dialogs, 1)
//
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
