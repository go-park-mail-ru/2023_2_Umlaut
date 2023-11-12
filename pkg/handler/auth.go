package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"net/http"
	"time"
)

// @Summary log in to account
// @Tags auth
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Sign-in input parameters"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/auth/login [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input signInInput
	if err := decoder.Decode(&input); err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusBadRequest, "invalid input body")
		return
	}

	if input.Mail == "" || input.Password == "" {
		newErrorClientResponseDto(&h.ctx, w, http.StatusBadRequest, "missing required fields")
		return
	}

	user, err := h.services.Authorization.GetUser(r.Context(), input.Mail, input.Password)
	if err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusUnauthorized, "invalid mail or password")
		return
	}
	SID, err := h.services.Authorization.GenerateCookie(r.Context(), user.Id)
	if err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	http.SetCookie(w, createCookie(SID))
	NewSuccessClientResponseDto(&h.ctx, w, "")
}

// @Summary log out of account
// @Tags auth
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/auth/logout [get]
func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorClientResponseDto(&h.ctx, w, http.StatusUnauthorized, "no session")
		return
	}
	if err = h.services.Authorization.DeleteCookie(r.Context(), session.Value); err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusInternalServerError, "Invalid cookie deletion")
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	session.Path = "/"

	http.SetCookie(w, session)
	NewSuccessClientResponseDto(&h.ctx, w, "")
}

// @Summary sign up account
// @Tags auth
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "Sign-up input user"
// @Success 200 {object} ClientResponseDto[idResponse]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input signUpInput
	if err := decoder.Decode(&input); err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusBadRequest, "invalid input body")
		return
	}

	if input.Name == "" || input.Mail == "" || input.Password == "" {
		newErrorClientResponseDto(&h.ctx, w, http.StatusBadRequest, "missing required fields")
		return
	}

	user := model.User{Name: input.Name, Mail: input.Mail, PasswordHash: input.Password}

	id, err := h.services.Authorization.CreateUser(r.Context(), user)
	if err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusBadRequest, err.Error())
		return
	}

	SID, err := h.services.Authorization.GenerateCookie(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(&h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	http.SetCookie(w, createCookie(SID))

	NewSuccessClientResponseDto(&h.ctx, w, idResponse{Id: id})
}

func createCookie(SID string) *http.Cookie {
	return &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		//SameSite: http.SameSiteNoneMode,
		//Secure:   true,
	}
}
