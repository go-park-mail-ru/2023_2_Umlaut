package repository

import (
	"context"
	"github.com/go-redis/redismock/v9"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()

func TestRedisStore_SetSession(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		value       int
		expectValue int
		time        time.Duration
		expectedErr bool
	}{
		{
			name:        "success test: redis set session",
			key:         "examplekey",
			value:       123,
			time:        30 * time.Minute,
			expectedErr: false,
		},
	}

	client, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("PONG")
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := NewRedisStore(client)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectSet(test.key, test.value, test.time).SetVal("OK")
			err = repo.SetSession(ctx, test.key, test.value, test.time)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestRedisStore_GetSession(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		expectedVal int
		expectedErr bool
	}{
		{
			name:        "success test: redis get session",
			key:         "examplekey",
			expectedVal: 123,
			expectedErr: false,
		},
	}

	client, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("PONG")
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := NewRedisStore(client)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectGet(test.key).SetVal(strconv.Itoa(test.expectedVal))
			val, err := repo.GetSession(ctx, test.key)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedVal, val)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestRedisStore_DeleteSession(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		expectedErr bool
	}{
		{
			name:        "success test: redis delete session",
			key:         "examplekey",
			expectedErr: false,
		},
	}

	client, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("PONG")
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	repo := NewRedisStore(client)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectDel(test.key).SetVal(int64(1))
			err := repo.DeleteSession(ctx, test.key)

			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
