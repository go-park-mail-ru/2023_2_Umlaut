package repository

// import (
// 	"context"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// func TestNewRedisClient(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		config        RedisConfig
// 		expectedError bool
// 	}{
// 		{
// 			name: "success test: create Redis client",
// 			config: RedisConfig{
// 				Addr:     "localhost:6379",
// 				Password: "password",
// 				DB:       1,
// 			},
// 			expectedError: false,
// 		},
// 		{
// 			name: "error test: invalid address",
// 			config: RedisConfig{
// 				Addr:     "invalid_address",
// 				Password: "password",
// 				DB:       1,
// 			},
// 			expectedError: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			client, err := NewRedisClient(test.config)

// 			if test.expectedError {
// 				assert.Error(t, err)
// 				assert.Nil(t, client)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, client)

// 				err = client.Ping(context.Background()).Err()
// 				assert.NoError(t, err)

// 				defer client.Close()
// 			}
// 		})
// 	}
// }
