package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, mail, password_hash, salt) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Mail, user.PasswordHash, user.Salt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUser(mail string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE mail=$1", usersTable)
	row := r.db.QueryRow(query, mail)
	if err := row.Scan(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) GetUserById(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) GetNextUser(user model.User) (model.User, error) {
	var nextUser model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id != $1 and user_gender = $2 ORDER BY RANDOM() LIMIT 1", usersTable)
	row := r.db.QueryRow(query, user.Id, user.PreferGender)
	if err := row.Scan(&nextUser); err != nil {
		return model.User{}, err
	}
	return nextUser, nil
}
