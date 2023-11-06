package repository

//import (
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
//	"github.com/stretchr/testify/assert"
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
//		name         string
//		mail         string
//		passwordHash string
//		salt         string
//		expectedID   int
//		expectedErr  bool
//	}{
//		{
//			name:         "Ok",
//			mail:         "testuser1@example.com",
//			passwordHash: "hashedpassword1",
//			salt:         "salt1",
//			expectedID:   1,
//			expectedErr:  false,
//		},
//		{
//			name:         "Ok",
//			mail:         "testuser2@example.com",
//			passwordHash: "hashedpassword2",
//			salt:         "salt2",
//			expectedID:   2,
//			expectedErr:  false,
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			rows := sqlmock.NewRows([]string{"id"}).AddRow(test.expectedID)
//
//			mock.ExpectQuery("INSERT INTO users").
//				WithArgs(test.name, test.mail, test.passwordHash, test.salt).
//				WillReturnRows(rows)
//
//			user := model.User{
//				Name:         test.name,
//				Mail:         test.mail,
//				PasswordHash: test.passwordHash,
//				Salt:         test.salt,
//			}
//
//			id, err := repo.CreateUser(user)
//
//			if test.expectedErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, test.expectedID, id)
//			}
//
//			if err := mock.ExpectationsWereMet(); err != nil {
//				t.Errorf("There were unfulfilled expectations: %s", err)
//			}
//		})
//	}
//}
//
//func TestUserPostgres_GetUser(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewUserPostgres(db)
//	tests := []struct {
//		name         string
//		mail         string
//		expectedUser model.User
//		expectedErr  bool
//	}{
//		{
//			name: "Ok",
//			mail: "testuser1@example.com",
//			expectedUser: model.User{
//				Id:           1,
//				Name:         "testName1",
//				Mail:         "testuser1@example.com",
//				PasswordHash: "testHash1",
//				Salt:         "testSalt1",
//			},
//			expectedErr: false,
//		},
//		{
//			name: "Ok",
//			mail: "testuser2@example.com",
//			expectedUser: model.User{
//				Id:           2,
//				Name:         "testName2",
//				Mail:         "testuser2@example.com",
//				PasswordHash: "testHash2",
//				Salt:         "testSalt2",
//			},
//			expectedErr: false,
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			rows := sqlmock.NewRows([]string{"id", "name", "mail", "passwordHash", "salt", "userGender",
//				"preferGender", "description", "age", "looking", "education", "hobbies", "tags"}).
//				AddRow(test.expectedUser.Id, test.expectedUser.Name, test.expectedUser.Mail, test.expectedUser.PasswordHash,
//					test.expectedUser.Salt, test.expectedUser.UserGender, test.expectedUser.PreferGender,
//					test.expectedUser.Description, test.expectedUser.Age, test.expectedUser.Looking,
//					test.expectedUser.Education, test.expectedUser.Hobbies, test.expectedUser.Tags)
//
//			mock.ExpectQuery("SELECT (.+) FROM users").
//				WithArgs(test.mail).
//				WillReturnRows(rows)
//
//			user, err := repo.GetUser(test.mail)
//
//			if test.expectedErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, test.expectedUser, user)
//			}
//
//			if err := mock.ExpectationsWereMet(); err != nil {
//				t.Errorf("There were unfulfilled expectations: %s", err)
//			}
//		})
//	}
//}
//
//func TestUserPostgres_GetUserById(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewUserPostgres(db)
//	tests := []struct {
//		name         string
//		id           int
//		expectedUser model.User
//		expectedErr  bool
//	}{
//		{
//			name: "Ok",
//			id:   1,
//			expectedUser: model.User{
//				Id:           1,
//				Name:         "testName1",
//				Mail:         "testuser1@example.com",
//				PasswordHash: "testHash1",
//				Salt:         "testSalt1",
//			},
//			expectedErr: false,
//		},
//		{
//			name: "Ok",
//			id:   2,
//			expectedUser: model.User{
//				Id:           2,
//				Name:         "testName2",
//				Mail:         "testuser2@example.com",
//				PasswordHash: "testHash2",
//				Salt:         "testSalt2",
//			},
//			expectedErr: false,
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			rows := sqlmock.NewRows([]string{"id", "name", "mail", "passwordHash", "salt", "userGender",
//				"preferGender", "description", "age", "looking", "education", "hobbies", "tags"}).
//				AddRow(test.expectedUser.Id, test.expectedUser.Name, test.expectedUser.Mail, test.expectedUser.PasswordHash,
//					test.expectedUser.Salt, test.expectedUser.UserGender, test.expectedUser.PreferGender,
//					test.expectedUser.Description, test.expectedUser.Age, test.expectedUser.Looking,
//					test.expectedUser.Education, test.expectedUser.Hobbies, test.expectedUser.Tags)
//
//			mock.ExpectQuery("SELECT (.+) FROM users").
//				WithArgs(test.id).
//				WillReturnRows(rows)
//
//			user, err := repo.GetUserById(test.id)
//
//			if test.expectedErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, test.expectedUser, user)
//			}
//
//			if err := mock.ExpectationsWereMet(); err != nil {
//				t.Errorf("There were unfulfilled expectations: %s", err)
//			}
//		})
//	}
//}
//
//func TestUserPostgres_GetNextUser(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewUserPostgres(db)
//	tests := []struct {
//		name         string
//		currentUser  model.User
//		expectedUser model.User
//		expectedErr  bool
//	}{
//		{
//			name: "Ok",
//			currentUser: model.User{
//				Id:           2,
//				Name:         "testName2",
//				Mail:         "testuser2@example.com",
//				PasswordHash: "testHash2",
//				Salt:         "testSalt2",
//			},
//			expectedUser: model.User{
//				Id:           1,
//				Name:         "testName1",
//				Mail:         "testuser1@example.com",
//				PasswordHash: "testHash1",
//				Salt:         "testSalt1",
//			},
//			expectedErr: false,
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			rows := sqlmock.NewRows([]string{"id", "name", "mail", "passwordHash", "salt", "userGender",
//				"preferGender", "description", "age", "looking", "education", "hobbies", "tags"}).
//				AddRow(test.expectedUser.Id, test.expectedUser.Name, test.expectedUser.Mail, test.expectedUser.PasswordHash,
//					test.expectedUser.Salt, test.expectedUser.UserGender, test.expectedUser.PreferGender,
//					test.expectedUser.Description, test.expectedUser.Age, test.expectedUser.Looking,
//					test.expectedUser.Education, test.expectedUser.Hobbies, test.expectedUser.Tags).
//				AddRow(test.currentUser.Id, test.currentUser.Name, test.currentUser.Mail, test.currentUser.PasswordHash,
//					test.currentUser.Salt, test.currentUser.UserGender, test.currentUser.PreferGender,
//					test.currentUser.Description, test.currentUser.Age, test.currentUser.Looking,
//					test.currentUser.Education, test.currentUser.Hobbies, test.currentUser.Tags)
//			mock.ExpectQuery("^SELECT (.+) FROM users*").
//				WithArgs(test.currentUser.Id).
//				WillReturnRows(rows)
//
//			user, err := repo.GetNextUser(test.currentUser)
//
//			if test.expectedErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, test.expectedUser, user)
//			}
//
//			if err := mock.ExpectationsWereMet(); err != nil {
//				t.Errorf("There were unfulfilled expectations: %s", err)
//			}
//		})
//	}
//}
