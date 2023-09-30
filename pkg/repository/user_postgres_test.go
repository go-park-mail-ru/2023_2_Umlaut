package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserPostgres(db)
	tests := []struct {
		name         string
		mail         string
		passwordHash string
		salt         string
		expectedID   int
		expectedErr  bool
	}{
		{
			name:         "Ok",
			mail:         "testuser1@example.com",
			passwordHash: "hashedpassword1",
			salt:         "salt1",
			expectedID:   1,
			expectedErr:  false,
		},
		{
			name:         "Ok",
			mail:         "testuser2@example.com",
			passwordHash: "hashedpassword2",
			salt:         "salt2",
			expectedID:   2,
			expectedErr:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(test.expectedID)

			mock.ExpectQuery("INSERT INTO users").
				WithArgs(test.name, test.mail, test.passwordHash, test.salt).
				WillReturnRows(rows)

			user := model.User{
				Name:         test.name,
				Mail:         test.mail,
				PasswordHash: test.passwordHash,
				Salt:         test.salt,
			}

			id, err := repo.CreateUser(user)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedID, id)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
