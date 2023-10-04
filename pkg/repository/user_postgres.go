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
	err := row.Scan(&id)

	return id, err
}

func (r *UserPostgres) GetUser(mail string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE mail=$1", usersTable)
	row := r.db.QueryRow(query, mail)
	err := ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetUserById(id int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query, id)
	err := ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetNextUser(user model.User) (model.User, error) {
	var nextUser model.User
	var query string
	var err error
	if user.PreferGender != nil {
		query = fmt.Sprintf("SELECT * FROM %s WHERE id != $1 and user_gender = $2 ORDER BY RANDOM() LIMIT 1", usersTable)
		row := r.db.QueryRow(query, user.Id, user.PreferGender)
		err = ScanUser(row, &nextUser)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE id != $1 ORDER BY RANDOM() LIMIT 1", usersTable)
		row := r.db.QueryRow(query, user.Id)
		err = ScanUser(row, &nextUser)
	}

	return nextUser, err
}

func ScanUser(row *sql.Row, user *model.User) error {
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Mail,
		&user.PasswordHash,
		&user.Salt,
		&user.UserGender,
		&user.PreferGender,
		&user.Description,
		&user.Age,
		&user.Looking,
		&user.Education,
		&user.Hobbies,
		&user.Tags,
	)
	return err
}
