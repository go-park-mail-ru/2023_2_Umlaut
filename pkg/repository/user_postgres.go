package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type UserPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user model.User) (int, error) {
	var id int
	query, args, err := psql.Insert(userTable).
		Columns("name", "mail", "password_hash", "salt").
		Values(user.Name, user.Mail, user.PasswordHash, user.Salt).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, static.ErrAlreadyExists
		}
		return 0, err
	}
	return id, err
}

func (r *UserPostgres) GetUser(ctx context.Context, mail string) (model.User, error) {
	var user model.User

	query, args, err := psql.Select(static.UserDbField).From(userTable).Where(sq.Eq{"mail": mail}).ToSql()

	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &user)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, fmt.Errorf("user with mail: %s not found", mail)
	}

	return user, err
}

func (r *UserPostgres) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query, args, err := psql.Select(static.UserDbField).From(userTable).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &user)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, fmt.Errorf("user with id: %d not found", id)
	}

	return user, err
}

func (r *UserPostgres) GetNextUser(ctx context.Context, user model.User) (model.User, error) {
	var nextUser model.User
	queryBuilder := psql.Select(static.UserDbField).From(userTable).Where(sq.NotEq{"id": user.Id})
	if user.PreferGender != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"user_gender": user.PreferGender})
	}
	queryBuilder = queryBuilder.OrderBy("RANDOM()").Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nextUser, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &nextUser)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, fmt.Errorf("user for: %s not found", user.Mail)
	}

	return nextUser, err
}

func (r *UserPostgres) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	query, args, err := psql.Update(userTable).
		Set("name", user.Name).
		Set("mail", user.Mail).
		Set("user_gender", user.UserGender).
		Set("prefer_gender", user.PreferGender).
		Set("description", user.Description).
		Set("birthday", user.Birthday).
		Set("looking", user.Looking).
		Set("image_paths", user.ImagePaths).
		Set("education", user.Education).
		Set("hobbies", user.Hobbies).
		Set("tags", user.Tags).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return model.User{}, err
	}

	query += fmt.Sprintf(" RETURNING %s", static.UserDbField)
	var updatedUser model.User
	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &updatedUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return updatedUser, static.ErrAlreadyExists
		}
	}

	return updatedUser, err
}

func (r *UserPostgres) UpdateUserPassword(ctx context.Context, user model.User) error {
	query, args, err := psql.Update(userTable).
		Set("password_hash", user.PasswordHash).
		Set("salt", user.Salt).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return err
	}

	query += fmt.Sprintf(" RETURNING %s", static.UserDbField)
	var updatedUser model.User
	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &updatedUser)

	return err
}

func (r *UserPostgres) GetNextUsers(ctx context.Context, user model.User, usedUsersId []int) ([]model.User, error) {
	var exp sq.And
	exp = append(exp, sq.NotEq{"id": user.Id})
	for _, id := range usedUsersId {
		exp = append(exp, sq.NotEq{"id": id})
	}

	queryBuilder := psql.Select(static.UserDbField).From(userTable).Where(exp)
	if user.PreferGender != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"user_gender": user.PreferGender})
	}

	query, args, err := queryBuilder.OrderBy("RANDOM()").Limit(5).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get next usersfor userId: %d. err: %w", user.Id, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get next usersfor userId: %d. err: %w", user.Id, err)
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err = scanUser(rows, &user)
		if errors.Is(err, pgx.ErrNoRows) {
			return users, fmt.Errorf("there are no suitable users for userId %d", user.Id)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, fmt.Errorf("failed to get next usersfor userId: %d. err: %w", user.Id, err)
	}

	return users, nil
}

func scanUser(row pgx.Row, user *model.User) error {
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Mail,
		&user.PasswordHash,
		&user.Salt,
		&user.UserGender,
		&user.PreferGender,
		&user.Description,
		&user.Looking,
		&user.ImagePaths,
		&user.Education,
		&user.Hobbies,
		&user.Birthday,
		&user.Online,
		&user.Tags,
	)

	user.CalculateAge()
	return err
}
