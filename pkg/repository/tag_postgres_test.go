package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

func TestComplaintPostgres_GetAllTags(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	tagRepo := NewTagPostgres(mock)

	testTags := []string{"type1", "type2"}

	mock.ExpectQuery(fmt.Sprintf(`SELECT %s FROM "tag"`, "name")).
		WillReturnRows(pgxmock.NewRows([]string{"name"}).
			AddRow("type1").AddRow("type2"))

	tags, err := tagRepo.GetAllTags(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, tags, testTags)

	mock.ExpectQuery(fmt.Sprintf(`SELECT %s FROM "tag"`, "name")).
		WillReturnError(errors.New("some other error"))

	_, err = tagRepo.GetAllTags(context.Background())

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
