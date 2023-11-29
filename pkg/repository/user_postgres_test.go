package repository

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPostgres_CreateUser(t *testing.T) {
	mock, mockErr := pgxmock.NewPool()
	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}
	defer mock.Close()

	userRepo := NewUserPostgres(mock)

	testUser := model.User{
		Name:         "John Doe",
		Mail:         "john@example.com",
		PasswordHash: "hashed_password",
		Salt:         "salt",
	}

	mock.ExpectQuery(`INSERT INTO "user"`).
		WithArgs(testUser.Name, testUser.Mail, testUser.PasswordHash, testUser.Salt).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))

	createdID, err := userRepo.CreateUser(context.Background(), testUser)

	assert.NoError(t, err)
	assert.Equal(t, 1, createdID)
}
