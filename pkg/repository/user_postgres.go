package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, mail, password_hash, salt) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(ctx, query, user.Name, user.Mail, user.PasswordHash, user.Salt)
	err := row.Scan(&id)

	return id, err
}

func (r *UserPostgres) GetUser(ctx context.Context, mail string) (model.User, error) {
	var user model.User

	queryBuilder := sq.Select("*").From(usersTable).Where(sq.Eq{"mail": mail}).Limit(1)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return user, err
	}
	row := r.db.QueryRow(ctx, query, args...)
	err = ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(ctx, query, id)
	err := ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetNextUser(ctx context.Context, user model.User) (model.User, error) {
	var nextUser model.User
	queryBuilder := sq.Select("*").From(usersTable).Where(sq.NotEq{"id": user.Id})
	if user.PreferGender != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"prefer_gender": user.PreferGender})
	}
	queryBuilder = queryBuilder.OrderBy("RANDOM()").Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nextUser, err
	}
	row := r.db.QueryRow(ctx, query, args...)
	err = ScanUser(row, &nextUser)

	return nextUser, err
}

func (r *UserPostgres) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	query := fmt.Sprintf(`
		UPDATE %s
		SET name = $2, mail = $3, user_gender = $4, prefer_gender = $5, description = $6, age = $7, looking = $8, education = $9, hobbies = $10, tags = $11
		WHERE id = $1
		RETURNING *`, usersTable)

	var updatedUser model.User
	row := r.db.QueryRow(
		ctx,
		query,
		user.Id,
		user.Name,
		user.Mail,
		user.UserGender,
		user.PreferGender,
		user.Description,
		user.Age,
		user.Looking,
		user.Education,
		user.Hobbies,
		user.Tags,
	)
	err := ScanUser(row, &updatedUser)

	return updatedUser, err
}

func ScanUser(row pgx.Row, user *model.User) error {
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
