package handler

import (
	"context"
	"encoding/json"
	"net/http"
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

type idResponse struct {
	Id int `json:"id"`
}

type ClientResponseDto[K comparable] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload K      `json:"payload"`
}

func NewClientResponseDto[K comparable](ctx *context.Context, w http.ResponseWriter, statusCode int, message string, payload K) {
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
	*ctx = context.WithValue(*ctx, keyStatus, statusCode)
	*ctx = context.WithValue(*ctx, keyMessage, message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

func NewSuccessClientResponseDto[K comparable](ctx *context.Context, w http.ResponseWriter, payload K) {
	NewClientResponseDto[K](ctx, w, 200, "success", payload)
}

func newErrorClientResponseDto(ctx *context.Context, w http.ResponseWriter, statusCode int, message string) {
	NewClientResponseDto[string](ctx, w, statusCode, message, "")
}
