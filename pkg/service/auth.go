package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/google/uuid"
)

type AuthService struct {
	repoUser  repository.User
	repoStore repository.Store
}

func NewAuthService(repoUser repository.User, repoStore repository.Store) *AuthService {
	return &AuthService{repoUser: repoUser, repoStore: repoStore}
}

func (s *AuthService) CreateUser(ctx context.Context, user model.User) (int, error) {
	if !user.IsValid() {
		return 0, errors.New("invalid fields")
	}
	user.Salt = generateUuid()
	user.PasswordHash = generatePasswordHash(user.PasswordHash, user.Salt)
	id, err := s.repoUser.CreateUser(ctx, user)
	if errors.Is(err, static.ErrAlreadyExists) {
		fmt.Println("account with this email already exists")
	}
	return id, err
}

func (s *AuthService) GetUser(ctx context.Context, mail, password string) (model.User, error) {
	user, err := s.repoUser.GetUser(ctx, mail)
	if err != nil {
		return user, err
	}
	if generatePasswordHash(password, user.Salt) != user.PasswordHash {
		return user, errors.New("invalid")
	}
	user.Sanitize()

	return user, nil
}

func (s *AuthService) GenerateCookie(ctx context.Context, id int) (string, error) {
	SID := generateUuid()
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

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateUuid() string {
	return uuid.NewString()
}
