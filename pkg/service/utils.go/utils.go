package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/google/uuid"
)

func GeneratePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateUuid() string {
	return uuid.NewString()
}

func Remove(data []string, value string) []string {
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

func GetBucketName(userId int) string {
	return fmt.Sprintf("user-id-%d", userId)
}
