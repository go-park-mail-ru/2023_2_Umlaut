package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/google/uuid"
)

type UserService struct {
	repoUser  repository.User
	repoStore repository.Store
	repoMinio repository.FileServer
}

func NewUserService(repoUser repository.User, repoStore repository.Store, repoMinio repository.FileServer) *UserService {
	return &UserService{repoUser: repoUser, repoStore: repoStore, repoMinio: repoMinio}
}

func (s *UserService) GetCurrentUser(ctx context.Context, userId int) (model.User, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("GetCurrentUser error: %v", err)
	}
	user.Sanitize()

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	if user.PasswordHash != "" {
		if !user.IsValid() {
			return user, static.ErrInvalidUser
		}
		user.Salt = generateUuid()
		user.PasswordHash = generatePasswordHash(user.PasswordHash, user.Salt)
		err := s.repoUser.UpdateUserPassword(ctx, user)
		if err != nil {
			return model.User{}, err
		}
	}
	correctUser, err := s.repoUser.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	correctUser.Sanitize()

	return correctUser, nil
}

func (s *UserService) CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error) {
	fileName := generateImageName()
	bucketName := getBucketName(userId)
	err := s.repoMinio.CreateBucket(ctx, bucketName)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}

	buffer := make([]byte, size)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	err = s.repoMinio.UploadFile(ctx, bucketName, fileName, contentType, file, size)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}
	currentUser, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}
	if currentUser.ImagePaths == nil {
		currentUser.ImagePaths = &[]string{fileName}
	} else {
		*currentUser.ImagePaths = append(*currentUser.ImagePaths, fileName)
	}
	_, err = s.repoUser.UpdateUser(ctx, currentUser)
	if err != nil {
		return fileName, fmt.Errorf("CreateFile error: %v", err)
	}

	return fileName, err
}

func (s *UserService) DeleteFile(ctx context.Context, userId int, link string) error {
	currentUser, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}
	if currentUser.ImagePaths == nil {
		return static.ErrNoFiles
	}

	err = s.repoMinio.DeleteFile(ctx, getBucketName(userId), getFileName(link))
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}
	*currentUser.ImagePaths = remove(*currentUser.ImagePaths, link)

	_, err = s.repoUser.UpdateUser(ctx, currentUser)
	if err != nil {
		return fmt.Errorf("DeleteFile error: %v", err)
	}

	return err
}

func generateImageName() string {
	return uuid.NewString()
}

func remove(data []string, value string) []string {
	for i, item := range data {
		if item != value {
			continue
		}
		data[i] = data[len(data)-1]
		data = data[:len(data)-1]
		break
	}
	return data
}

func getFileName(link string) string {
	i := strings.LastIndex(link, "/")
	if i == -1 {
		return link
	}
	return link[i+1:]
}

func getBucketName(userId int) string {
	return fmt.Sprintf("user-id-%d", userId)
}
