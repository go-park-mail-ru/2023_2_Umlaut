package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"math/rand"
	"net/http"
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
	return user, nil
}

func (s *AuthService) GenerateCookie(ctx context.Context, id int) (*http.Cookie, error) {
	SID := generateCookie()
	if err := s.repoStore.SetSession(ctx, SID, id, 10*time.Hour); err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	return cookie, nil
}

func (s *AuthService) DeleteCookie(ctx context.Context, session *http.Cookie) error {
	if err := s.repoStore.DeleteSession(ctx, session.Value); err != nil {
		return err
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
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
