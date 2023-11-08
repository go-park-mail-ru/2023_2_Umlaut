package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtToken struct {
	Ctx    *context.Context
	Secret []byte
}

func NewJwtToken(ctx *context.Context, secret string) *JwtToken {
	return &JwtToken{Ctx: ctx, Secret: []byte(secret)}
}

type JwtCsrfClaims struct {
	SessionID string `json:"sid"`
	UserID    int    `json:"uid"`
	jwt.StandardClaims
}

func (tk *JwtToken) Create(SID string, id int, tokenExpTime int64) (string, error) {
	data := JwtCsrfClaims{
		SessionID: SID,
		UserID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpTime,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) parseSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, fmt.Errorf("bad sign method")
	}
	return tk.Secret, nil
}

func (tk *JwtToken) Check(SID string, inputToken string) (bool, error) {
	payload := &JwtCsrfClaims{}
	_, err := jwt.ParseWithClaims(inputToken, payload, tk.parseSecretGetter)
	if err != nil {
		return false, fmt.Errorf("cant parse jwt token: %v", err)
	}
	if payload.Valid() != nil {
		return false, fmt.Errorf("invalid jwt token: %v", err)
	}
	return payload.SessionID == SID && payload.UserID == (*tk.Ctx).Value(keyUserID).(int), nil
}
