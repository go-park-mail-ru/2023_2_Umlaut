package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/static"

	"go.uber.org/zap"
)

type signInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpInput struct {
	Name      string  `json:"name" binding:"required"`
	Mail      string  `json:"mail" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	InvitedBy *string `json:"invited_by"`
}

type shareCridentialsOutput struct {
	InvitesCount int    `json:"invites_count" binding:"required"`
	ShareLink    string `json:"share_link" binding:"required"`
}

type deleteLink struct {
	Link string `json:"link"`
}

type idResponse struct {
	Id int `json:"id"`
}

type ClientResponseDto struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func NewClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string, payload interface{}) {
	response := ClientResponseDto{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	}
	sendData(ctx, w, response, statusCode, message)
}

func NewSuccessClientResponseDto(ctx context.Context, w http.ResponseWriter, payload interface{}) {
	NewClientResponseDto(ctx, w, 200, "success", payload)
}

func newErrorClientResponseDto(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	NewClientResponseDto(ctx, w, statusCode, message, "")
}

func sendData(ctx context.Context, w http.ResponseWriter, response ClientResponseDto, statusCode int, message string) {
	responseJSON, err := response.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	logger, ok := ctx.Value(static.KeyLogger).(*zap.Logger)
	if !ok {
		log.Println("Logger not found in context")
	} else {
		*logger = *logger.With(zap.Any("Status", statusCode), zap.Any("Message", message))
	}

	requestInfo, ok := ctx.Value(static.KeyRequestInfo).(*RequestInfo)
	if !ok {
		log.Println("Request info not found in context")
	} else {
		requestInfo.Status = statusCode
		requestInfo.Message = message
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
