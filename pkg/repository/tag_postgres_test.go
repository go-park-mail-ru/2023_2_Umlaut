package repository

// import (
// 	"context"
// 	"testing"

// 	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
// 	"github.com/stretchr/testify/assert"
// )

// func TestTagPostgres_GetAllTags(t *testing.T) {
// 	pool, err := initPostgres()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
// 	}

// 	repo := NewTagPostgres(pool)

//	ctx := context.Background()

// 	tests := []struct {
// 		name         string
// 		expectedTags []model.Tag
// 		expectedErr  bool
// 	}{
// 		{
// 			name: "Ok",
// 			expectedTags: []model.Tag{
// 				{
// 					Id:   1,
// 					Name: "testtag1",
// 				},
// 				{
// 					Id:   2,
// 					Name: "testtag2",
// 				},
// 			},
// 			expectedErr: false,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			tags, err := repo.GetAllTags(ctx)

// 			if test.expectedErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, test.expectedTags, tags)
// 			}
// 		})
// 	}

// }
