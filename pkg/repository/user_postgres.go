package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
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

	return id, err
}

func (r *UserPostgres) GetUser(ctx context.Context, mail string) (model.User, error) {
	var user model.User

	query, args, err := psql.Select("*").From(userTable).Where(sq.Eq{"mail": mail}).ToSql()
	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetUserById(ctx context.Context, id int) (model.User, error) {
	var user model.User

	query, args, err := psql.Select("*").From(userTable).Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = ScanUser(row, &user)

	return user, err
}

func (r *UserPostgres) GetNextUser(ctx context.Context, user model.User) (model.User, error) {
	var nextUser model.User
	queryBuilder := psql.Select("*").From(userTable).Where(sq.NotEq{"id": user.Id})
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
	user.CalculateAge()

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
		Set("education", user.Education).
		Set("hobbies", user.Hobbies).
		//Set("tags", user.Tags).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return model.User{}, err
	}

	query += " RETURNING *"
	var updatedUser model.User
	row := r.db.QueryRow(ctx, query, args...)
	err = ScanUser(row, &updatedUser)

	return updatedUser, err
}

func (r *UserPostgres) UpdateUserPhoto(ctx context.Context, userId int, imagePath string) (string, error) {
	queryBuilder := squirrel.Update(userTable).
		Set("image_path", imagePath).
		Where(squirrel.Eq{"id": userId})
	
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return "", err
	}

	query += " RETURNING image_path"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&imagePath)

	return imagePath, err
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
		&user.Looking,
		&user.Education,
		&user.Hobbies,
		//&user.Tags,
		&user.Birthday,
		&user.Online,
	)
	return err
}
