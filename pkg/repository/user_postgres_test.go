package repository

//
//import (
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
//	"testing"
//)
//
//func TestUserPostgres_CreateUser(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewUserPostgres(db)
//	tests := []struct {
//		name    string
//		mock    func(user model.User)
//		input   model.User
//		want    int
//		wantErr bool
//	}{
//		{
//			name:  "Ok",
//			input: model.User{Mail: "m@m.ru", Name: "n", PasswordHash: "pass", Salt: "salt"},
//			want:  1,
//			mock: func(user model.User) {
//				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
//				mock.ExpectQuery("INSERT INTO users").
//					WithArgs(user.Name, user.Mail, user.PasswordHash, user.Salt).WillReturnRows(rows)
//			},
//			wantErr: false,
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			test.mock(test.input)
//
//		})
//	}
//}
