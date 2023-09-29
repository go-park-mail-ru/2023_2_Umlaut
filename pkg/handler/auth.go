package handler

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type signInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary signIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Sign-in input parameters"
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Router /auth/login [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrorResponse(w, http.StatusBadRequest, "Authentication failed")
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	var input signInInput
	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	user, err := h.Repositories.GetUser(input.Mail)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if generatePasswordHash(input.Password, user.Salt) != user.PasswordHash {
		newErrorResponse(w, http.StatusUnauthorized, "invalid mail or password")
		return
	}

	SID := generateCookie()
	ctx := r.Context()
	if err = h.Repositories.SetSession(ctx, SID, user.Id, 10*time.Hour); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)

	jsonResponse, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// @Summary logout
// @Tags auth
// @Description logout
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} errorResponse
// @Router /auth/logout [get]
func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}
	ctx := r.Context()
	if err = h.Repositories.DeleteSession(ctx, session.Value); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Invalid cookie deletion")
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

// @Summary signUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "Sign-up input user"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrorResponse(w, http.StatusBadRequest, "Registration failed")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	user := model.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	user.Salt = generateSalt()
	user.PasswordHash = generatePasswordHash(user.PasswordHash, user.Salt)
	id, err := h.Repositories.CreateUser(user)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	ctx := r.Context()
	SID := generateCookie()
	if err = h.Repositories.SetSession(ctx, SID, id, 10*time.Hour); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]int{
		"id": id,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
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
