package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type signInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpInput struct {
	Name     string `json:"name" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type likeDto struct {
	LikedUserId int      `json:"liked_user_id" binding:"required"`
	CommittedAt JsonTime `json:"committed_at" binding:"required"`
}

type idResponse struct {
	Id int `json:"id"`
}

type ClientResponseDto[K comparable] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload K      `json:"payload"`
}

type JsonTime time.Time

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonTime(t)
	return nil
}

func NewClientResponseDto[K comparable](ctx context.Context, w http.ResponseWriter, statusCode int, message string, payload K) {
	response := ClientResponseDto[K]{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	ctx = context.WithValue(ctx, "Status", statusCode)
	ctx = context.WithValue(ctx, "Message", message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

func NewSuccessClientResponseDto[K comparable](ctx context.Context, w http.ResponseWriter, payload K) {
	NewClientResponseDto[K](ctx, w, 200, "success", payload)
}

func newErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	NewClientResponseDto[string](ctx, w, statusCode, message, "")
}
