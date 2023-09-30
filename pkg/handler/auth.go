package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"io"
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
		return
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

	user, err := h.services.GetUser(input.Mail, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, "invalid mail or password")
		return
	}
	cookie, err := h.services.GenerateCookie(r.Context(), user.Id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
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
	if err = h.services.DeleteCookie(r.Context(), session); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Invalid cookie deletion")

	}
	http.SetCookie(w, session)
	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

// @Summary signUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "Sign-up input user"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrorResponse(w, http.StatusBadRequest, "Registration failed")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	var input signUpInput
	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}
	user := model.User{Name: input.Name, Mail: input.Mail, PasswordHash: input.Password}

	id, err := h.services.CreateUser(user)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Account with this email already exists")
		return
	}

	cookie, err := h.services.GenerateCookie(r.Context(), user.Id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
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
