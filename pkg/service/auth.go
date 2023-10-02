package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"math/rand"
	"time"
)

type AuthService struct {
	repoUser  repository.User
	repoStore repository.Store
}

func NewAuthService(repoUser repository.User, repoStore repository.Store) *AuthService {
	return &AuthService{repoUser: repoUser, repoStore: repoStore}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	if !user.IsValid() {
		return 0, errors.New("invalid fields")
	}
	user.Salt = generateSalt()
	user.PasswordHash = generatePasswordHash(user.PasswordHash, user.Salt)
	id, err := s.repoUser.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *AuthService) GetUser(mail, password string) (model.User, error) {
	user, err := s.repoUser.GetUser(mail)
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
	SID := generateCookie()
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

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateCookie() string {
	return randStringRunes(32)
}

func generateSalt() string {
	return randStringRunes(22)
}

func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
