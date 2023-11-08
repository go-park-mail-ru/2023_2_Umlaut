package repository

import (
	"context"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/stretchr/testify/assert"
)

func TestDialogPostgres_CreateDialog(t *testing.T) {
	ctx := context.Background()

	pool, err := initPostgres(ctx)
	
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a test database connection", err)
	}

	repo := NewDialogPostgres(pool)

	tests := []struct {
		name        string
		dialog      model.Dialog
		expectedID  int
		expectedErr bool
	}{
		{
			name: "Ok",
			dialog: model.Dialog{
				User1Id: 1,
				User2Id: 2,
			},
			expectedID:  1,
			expectedErr: false,
		},
		{
			name: "Repeat dialog",
			dialog: model.Dialog{
				User1Id: 2,
				User2Id: 1,
			},
			expectedID:  2,
			expectedErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			id, err := repo.CreateDialog(ctx, test.dialog)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedID, id)
			}
		})
	}

}
