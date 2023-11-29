package repository

//func TestLikePostgres_CreateLike(t *testing.T) {
//	mock, err := pgxmock.NewPool()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer mock.Close()
//
//	likeRepo := NewLikePostgres(mock)
//
//	testLike := model.Like{
//		LikedByUserId: 1,
//		LikedToUserId: 2,
//		IsLike:        true,
//	}
//
//	mock.ExpectQuery(`INSERT INTO "like"`).
//		WithArgs(testLike.LikedByUserId, testLike.LikedToUserId).
//		WillReturnRows(pgxmock.NewRows([]string{"liked_by_user_id", "liked_to_user_id", "is_like"}).
//			AddRow(testLike.LikedByUserId, testLike.LikedToUserId, testLike.IsLike))
//
//	newLike, err := likeRepo.CreateLike(context.Background(), testLike)
//
//	assert.NoError(t, err)
//	assert.Equal(t, true, newLike.IsLike)
//
//	mock.ExpectQuery(`INSERT INTO "like"`).
//		WithArgs(testLike.LikedByUserId, testLike.LikedToUserId, testLike.IsLike).
//		WillReturnError(static.ErrAlreadyExists)
//
//	_, err = likeRepo.CreateLike(context.Background(), testLike)
//
//	assert.ErrorIs(t, err, static.ErrAlreadyExists)
//
//	mock.ExpectQuery(`INSERT INTO "like"`).
//		WithArgs(testLike.LikedByUserId, testLike.LikedToUserId, testLike.IsLike).
//		WillReturnError(errors.New("some other error"))
//
//	_, err = likeRepo.CreateLike(context.Background(), testLike)
//
//	assert.Error(t, err)
//
//	assert.NoError(t, mock.ExpectationsWereMet())
//}

//func TestLikePostgres_IsMutualLike(t *testing.T) {
//	mock, err := pgxmock.NewPool()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer mock.Close()
//
//	likeRepo := NewLikePostgres(mock)
//
//	testLike := model.Like{
//		LikedByUserId: 1,
//		LikedToUserId: 2,
//		IsLike:        true,
//	}
//
//	mock.ExpectQuery(`SELECT`).
//		WithArgs(testLike.LikedToUserId, testLike.LikedByUserId, testLike.IsLike).
//		WillReturnRows(pgxmock.NewRows([]string{static.LikeDbField}))
//
//	isMutual, err := likeRepo.IsMutualLike(context.Background(), testLike)
//
//	assert.NoError(t, err)
//	assert.False(t, isMutual)
//
//	mock.ExpectQuery(`SELECT`).
//		WithArgs(testLike.LikedToUserId, testLike.LikedByUserId, testLike.IsLike).
//		WillReturnRows(pgxmock.NewRows([]string{static.LikeDbField}).AddRow(1))
//
//	isMutual, err = likeRepo.IsMutualLike(context.Background(), testLike)
//
//	assert.NoError(t, err)
//	assert.True(t, isMutual)
//
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
