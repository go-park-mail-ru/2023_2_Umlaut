package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/stretchr/testify/assert"
)

func initPostgres() (*pgx.Conn, error) {
	return NewPostgresDB(PostgresConfig{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "1474",
	})
}

func TestUserPostgres_CreateUser(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()

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
			mail:         "testuser100@example.com",
			passwordHash: "hashedpassword1",
			salt:         "salt1",
			expectedID:   1,
			expectedErr:  false,
		},
		{
			name:         "Recurring User",
			mail:         "testuser100@example.com",
			passwordHash: "hashedpassword2",
			salt:         "salt2",
			expectedID:   2,
			expectedErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user := model.User{
				Name:         test.name,
				Mail:         test.mail,
				PasswordHash: test.passwordHash,
				Salt:         test.salt,
			}
			id, err := repo.CreateUser(ctx, user)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedID, id)
			}
		})
	}

}

func TestUserPostgres_GetUser(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()

	tests := []struct {
		name         string
		mail         string
		expectedUser model.User
		expectedErr  bool
	}{
		{
			name: "Ok",
			mail: "testuser1@example.com",
			expectedUser: model.User{
				Id:           1,
				Name:         "testName1",
				Mail:         "testuser1@example.com",
				PasswordHash: "testHash1",
				Salt:         "testSalt1",
			},
			expectedErr: false,
		},
		{
			name: "Incorrect mail",
			mail: "notcorrecttestuser2@example.com",
			expectedUser: model.User{
				Id:           2,
				Name:         "testName2",
				Mail:         "notcorrecttestuser2@example.com",
				PasswordHash: "testHash2",
				Salt:         "testSalt2",
			},
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := repo.GetUser(ctx, test.mail)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedUser, user)
			}
		})
	}
}

func TestUserPostgres_GetUserById(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()

	tests := []struct {
		name         string
		id           int
		expectedUser model.User
		expectedErr  bool
	}{
		{
			name: "Ok",
			id:   1,
			expectedUser: model.User{
				Id:           1,
				Name:         "testName1",
				Mail:         "testuser1@example.com",
				PasswordHash: "testHash1",
				Salt:         "testSalt1",
			},
			expectedErr: false,
		},
		{
			name: "Incorrect id",
			id:   -2,
			expectedUser: model.User{
				Id:           -2,
				Name:         "testName2",
				Mail:         "testuser2@example.com",
				PasswordHash: "testHash2",
				Salt:         "testSalt2",
			},
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := repo.GetUserById(ctx, test.id)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedUser, user)
			}
		})
	}
}

func TestUserPostgres_UpdateUser(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()

	tests := []struct {
		name        string
		user        model.User
		expectedErr bool
	}{
		{
			name: "Ok",
			user: model.User{
				Id:           1,
				Name:         "testName1",
				Mail:         "testuser1@example.com",
				PasswordHash: "testHash1",
				Salt:         "testSalt1",
			},
			expectedErr: false,
		},
		{
			name: "Incorrect user",
			user: model.User{
				Id:           -2,
				Name:         "testName2",
				Mail:         "testuser2@example.com",
				PasswordHash: "testHash2",
				Salt:         "testSalt2",
			},
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := repo.UpdateUser(ctx, test.user)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.user, user)
			}
		})
	}
}

func TestUserPostgres_UpdateUserPhoto(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()
	image_path1 := "testimagepath1"

	tests := []struct {
		name        string
		id          int
		imagePath   *string
		expectedErr bool
	}{
		{
			name:        "Ok",
			id:          1,
			imagePath:   &image_path1,
			expectedErr: false,
		},
		{
			name:        "Ok: nil",
			id:          2,
			imagePath:   nil,
			expectedErr: false,
		},
		{
			name:        "Incorrect id",
			id:          -2,
			imagePath:   nil,
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newImgPath, err := repo.UpdateUserPhoto(ctx, test.id, test.imagePath)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.imagePath, newImgPath)
			}
		})
	}
}

func TestUserPostgres_GetNextUsers(t *testing.T) {
	pool, err := initPostgres()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewUserPostgres(pool)

	ctx := context.Background()

	tests := []struct {
		name        string
		user        model.User
		usedUsers   []int
		expectedErr bool
	}{
		{
			name: "Ok",
			user: model.User{
				Id:           1,
				Name:         "testName1",
				Mail:         "testuser1@example.com",
				PasswordHash: "testHash1",
				Salt:         "testSalt1",
			},
			usedUsers:   []int{1},
			expectedErr: false,
		},
		{
			name: "Ok",
			user: model.User{
				Id:           2,
				Name:         "testName2",
				Mail:         "testuser2@example.com",
				PasswordHash: "testHash2",
				Salt:         "testSalt2",
			},
			usedUsers:   []int{2},
			expectedErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			users, err := repo.GetNextUsers(ctx, test.user, test.usedUsers)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(users), 5)
			}
		})
	}
}
