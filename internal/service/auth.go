package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/utils"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

type AuthService struct {
	repoUser  repository.User
	repoStore repository.Store
	repoAdmin repository.Admin
}

func NewAuthService(repoUser repository.User, repoStore repository.Store, repoAdmin repository.Admin) *AuthService {
	return &AuthService{repoUser: repoUser, repoStore: repoStore, repoAdmin: repoAdmin}
}

func (s *AuthService) CreateUser(ctx context.Context, user core.User) (int, error) {
	if !user.IsValid() {
		return 0, errors.New("invalid fields")
	}
	if user.InvitedBy != nil && *user.InvitedBy == 0 {
		user.InvitedBy = nil
	}
	user.Salt = utils.GenerateUuid()
	user.PasswordHash = utils.GeneratePasswordHash(user.PasswordHash, user.Salt)
	id, err := s.repoUser.CreateUser(ctx, user)
	if errors.Is(err, constants.ErrAlreadyExists) {
		fmt.Println("account with this email already exists")
	}
	return id, err
}

func (s *AuthService) GetUser(ctx context.Context, mail, password string) (core.User, error) {
	user, err := s.repoUser.GetUser(ctx, mail)
	if errors.Is(err, constants.ErrBannedUser) {
		return core.User{}, err
	}
	if err != nil {
		return user, err
	}
	if utils.GeneratePasswordHash(password, user.Salt) != user.PasswordHash {
		return user, errors.New("invalid")
	}
	user.Sanitize()

	return user, nil
}

func (s *AuthService) GetAdmin(ctx context.Context, mail, password string) (core.Admin, error) {
	admin, err := s.repoAdmin.GetAdmin(ctx, mail)
	if err != nil {
		return admin, err
	}
	if utils.GeneratePasswordHash(password, admin.Salt) != admin.PasswordHash {
		return admin, errors.New("invalid")
	}
	admin.Sanitize()

	return admin, nil
}

func (s *AuthService) GenerateCookie(ctx context.Context, id int) (string, error) {
	SID := utils.GenerateUuid()
	if err := s.repoStore.SetSession(ctx, SID, id, 10*time.Hour); err != nil {
		return SID, err
	}

	return SID, nil
}

func (s *AuthService) DeleteCookie(ctx context.Context, session string) error {
	if err := s.repoStore.DeleteSession(ctx, session); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetSessionValue(ctx context.Context, session string) (int, error) {
	id, err := s.repoStore.GetSession(ctx, session)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AuthService) GetDecodeUserId(ctx context.Context, message string) (int, error) {
	newMessage, err := utils.DecryptString(message)
	if err != nil {
		return 0, fmt.Errorf("GetDecodeUserId error: %v", err)
	}

	return strconv.Atoi(newMessage)
}
