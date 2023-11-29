package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/stretchr/testify/assert"

	"github.com/pashagolub/pgxmock/v3"
)

func TestDialogPostgres_CreateDialog(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	dialogRepo := NewDialogPostgres(mock)

	testDialog := model.Dialog{
		User1Id: 1,
		User2Id: 2,
	}

	// Ожидаем успешное создание диалога
	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

	createdID, err := dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdID)

	// Проверяем ситуацию с дубликатом
	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnError(static.ErrAlreadyExists)

	_, err = dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.ErrorIs(t, err, static.ErrAlreadyExists)

	// Проверка других случаев ошибок
	mock.ExpectQuery(`INSERT INTO "dialog"`).
		WithArgs(testDialog.User1Id, testDialog.User2Id).
		WillReturnError(errors.New("some other error"))

	_, err = dialogRepo.CreateDialog(context.Background(), testDialog)

	assert.Error(t, err)

	// Проверяем, что не остались ожидающие запросы
	assert.NoError(t, mock.ExpectationsWereMet())
}
