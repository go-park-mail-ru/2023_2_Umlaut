package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type UserPostgres struct {
	db PgxPoolInterface
}

func NewUserPostgres(db PgxPoolInterface) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user core.User) (int, error) {
	var id int

	var mail *string
	if len(user.Mail) > 1 {
		mail = &user.Mail
	}
	query, args, err := psql.Insert(userTable).
		Columns("name", "mail", "password_hash", "salt", "user_gender", "prefer_gender", "birthday", "image_paths", "invited_by", "oauth_id").
		Values(user.Name, mail, user.PasswordHash, user.Salt, user.UserGender, user.PreferGender, user.Birthday, user.ImagePaths, user.InvitedBy, user.OauthId).
		ToSql()

	if err != nil {
		return 0, err
	}

	query += " RETURNING id"
	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, constants.ErrAlreadyExists
		}
		return 0, err
	}
	return id, err
}

func (r *UserPostgres) InsertOrUpdateUser(ctx context.Context, user core.User) (int, error) {
	var id int
	query, args, err := psql.Select("id").From(userTable).
		Where(sq.Eq{"oauth_id": user.OauthId}).ToSql()

	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return r.CreateUser(ctx, user)
	} else if err != nil || id > 0 {
		return id, err
	}

	return r.CreateUser(ctx, user)
}

func (r *UserPostgres) GetUser(ctx context.Context, mail string) (core.User, error) {
	var user core.User

	query, args, err := psql.Select(constants.UserDbField).From(userTable).Where(sq.Eq{"mail": mail}).ToSql()

	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &user)

	if errors.Is(err, pgx.ErrNoRows) {
		return core.User{}, fmt.Errorf("user with mail: %s not found", mail)
	}
	if user.Role == constants.Banned {
		return core.User{}, constants.ErrBannedUser
	}

	return user, err
}

func (r *UserPostgres) GetUserById(ctx context.Context, id int) (core.User, error) {
	var user core.User

	query, args, err := psql.Select(constants.UserDbField).From(userTable).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return user, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &user)

	if errors.Is(err, pgx.ErrNoRows) {
		return core.User{}, fmt.Errorf("user with id: %d not found", id)
	}
	if user.Role == constants.Banned {
		return core.User{}, constants.ErrBannedUser
	}

	return user, err
}

func (r *UserPostgres) GetNextUser(ctx context.Context, user core.User, params dto.FilterParams) (core.User, error) {
	var nextUser core.User

	queryBuilder := psql.Select(constants.UserDbField).
		From(userTable).
		Where(sq.NotEq{"id": user.Id}).
		Where(sq.NotEq{"role": constants.Banned}).
		Where(fmt.Sprintf("id NOT IN (SELECT reported_user_id FROM %s WHERE reporter_user_id = %d)", complaintTable, user.Id)).
		Where(fmt.Sprintf("id NOT IN (SELECT liked_to_user_id FROM %s WHERE liked_by_user_id = %d)", likeTable, user.Id))

	if user.PreferGender != nil && user.UserGender != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"user_gender": user.PreferGender, "prefer_gender": user.UserGender})
	}
	if params.MinAge != 0 {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"age": params.MinAge})
	}
	if params.MaxAge != 0 {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"age": params.MaxAge})
	}
	if len(params.Tags) > 0 && len(params.Tags[0]) > 1 {
		queryBuilder = queryBuilder.Where("ARRAY[" + buildTagArray(params.Tags) + "]::TEXT[] <@ tags")
	}
	queryBuilder = queryBuilder.OrderBy("RANDOM()").Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nextUser, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &nextUser)

	if errors.Is(err, pgx.ErrNoRows) || nextUser.Id == 0 {
		return core.User{}, fmt.Errorf("user for: %s not found", user.Mail)
	}

	return nextUser, err
}

func buildTagArray(tags []string) string {
	arrayString := ""
	for i, tag := range tags {
		if i > 0 {
			arrayString += ", "
		}
		arrayString += "'" + tag + "'"
	}
	return arrayString
}

func (r *UserPostgres) UpdateUser(ctx context.Context, user core.User) (core.User, error) {
	var mail *string
	if len(user.Mail) > 1 {
		mail = &user.Mail
	}
	query, args, err := psql.Update(userTable).
		Set("name", user.Name).
		Set("mail", mail).
		Set("user_gender", user.UserGender).
		Set("prefer_gender", user.PreferGender).
		Set("description", user.Description).
		Set("birthday", user.Birthday).
		Set("looking", user.Looking).
		Set("education", user.Education).
		Set("hobbies", user.Hobbies).
		Set("tags", user.Tags).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return core.User{}, err
	}

	query += fmt.Sprintf(" RETURNING %s", constants.UserDbField)
	var updatedUser core.User
	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &updatedUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return updatedUser, constants.ErrAlreadyExists
		}
	}

	return updatedUser, err
}

func (r *UserPostgres) UpdateUserPhoto(ctx context.Context, user core.User) error {
	query, args, err := psql.Update(userTable).
		Set("image_paths", user.ImagePaths).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *UserPostgres) UpdateUserPassword(ctx context.Context, user core.User) error {
	query, args, err := psql.Update(userTable).
		Set("password_hash", user.PasswordHash).
		Set("salt", user.Salt).
		Where(sq.Eq{"id": user.Id}).
		ToSql()

	if err != nil {
		return err
	}

	query += fmt.Sprintf(" RETURNING %s", constants.UserDbField)
	var updatedUser core.User
	row := r.db.QueryRow(ctx, query, args...)
	err = scanUser(row, &updatedUser)

	return err
}

func (r *UserPostgres) ShowCSAT(ctx context.Context, userId int) (bool, error) {
	var id int

	query, args, err := psql.Select("id").From(userTable).Where(
		sq.And{
			sq.Lt{"EXTRACT(DAY FROM NOW()-created_at)": "1"},
			sq.Eq{"id": userId},
		}).ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to check can show csat. err: %w", err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&id)

	if errors.Is(err, pgx.ErrNoRows) {
		return true, nil
	}

	return false, err
}

func (r *UserPostgres) GetUserInvites(ctx context.Context, userId int) (int, error) {
	var count int

	query, args, err := psql.Select("count(id)").From(userTable).
		Where(fmt.Sprintf("invited_by = %d AND description IS NOT NULL", userId)).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to get user invites. err: %w", err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	err = row.Scan(&count)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}

	return count, err
}

func (r *UserPostgres) ResetLikeCounter(ctx context.Context) error {
	_, err := r.db.Exec(ctx, fmt.Sprintf("update %s set like_counter = default where role != 2;", userTable))
	return err
}

func scanUser(row pgx.Row, user *core.User) error {
	var mail *string
	err := row.Scan(
		&user.Id,
		&user.Name,
		&mail,
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
		&user.Role,
		&user.LikeCounter,
		&user.Online,
		&user.Tags,
		&user.OauthId,
	)
	if mail != nil {
		user.Mail = *mail
	}

	user.CalculateAge()
	return err
}
