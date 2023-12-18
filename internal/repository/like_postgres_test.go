package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLikePostgres_CreateLike(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	likeRepo := NewLikePostgres(mock)

	testLike := core.Like{
		LikedByUserId: 1,
		LikedToUserId: 2,
		IsLike:        true,
	}

	mock.ExpectQuery(`INSERT INTO "like"`).
		WithArgs(testLike.LikedByUserId, testLike.LikedToUserId, testLike.IsLike).
		WillReturnError(constants.ErrAlreadyExists)

	_, err = likeRepo.CreateLike(context.Background(), testLike)

	assert.ErrorIs(t, err, constants.ErrAlreadyExists)

	mock.ExpectQuery(`INSERT INTO "like"`).
		WithArgs(testLike.LikedByUserId, testLike.LikedToUserId, testLike.IsLike).
		WillReturnError(errors.New("some other error"))

	_, err = likeRepo.CreateLike(context.Background(), testLike)

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLikePostgres_IsMutualLike(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	likeRepo := NewLikePostgres(mock)

	testLike := core.Like{
		LikedByUserId: 1,
		LikedToUserId: 2,
		IsLike:        true,
	}

	mock.ExpectQuery(`SELECT`).
		WithArgs(testLike.IsLike, testLike.LikedToUserId, testLike.LikedByUserId).
		WillReturnRows(pgxmock.NewRows([]string{"liked_by_user_id", "liked_to_user_id"}))

	isMutual, err := likeRepo.IsMutualLike(context.Background(), testLike)

	assert.NoError(t, err)
	assert.False(t, isMutual)

	mock.ExpectQuery(`SELECT`).
		WithArgs(testLike.IsLike, testLike.LikedToUserId, testLike.LikedByUserId).
		WillReturnRows(pgxmock.NewRows([]string{"liked_by_user_id", "liked_to_user_id"}).
			AddRow(1, 2))

	isMutual, err = likeRepo.IsMutualLike(context.Background(), testLike)

	assert.NoError(t, err)
	assert.True(t, isMutual)

	assert.NoError(t, mock.ExpectationsWereMet())
}
