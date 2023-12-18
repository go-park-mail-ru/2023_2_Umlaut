package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"io"

	"github.com/google/uuid"
)

var encodingKey = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func DecryptString(cryptoText string) (string, error) {
	newKeyString, _ := hashTo32Bytes(encodingKey)
	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(newKeyString))
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func EncryptString(plainText string) (string, error) {
	newKeyString, err := hashTo32Bytes(encodingKey)
	if err != nil {
		return "", err
	}

	key := []byte(newKeyString)
	value := []byte(plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(value))

	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], value)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func hashTo32Bytes(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("no input supplied")
	}

	hasher := sha256.New()
	hasher.Write([]byte(input))

	stringToSHA256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return stringToSHA256[:32], nil
}

func GenerateUserShakeLink(edcodePart string) string {
	return constants.Host + "/signup/" + edcodePart
}
