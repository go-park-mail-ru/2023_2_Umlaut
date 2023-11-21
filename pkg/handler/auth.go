package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Summary log in to account
// @Tags auth
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Sign-in input parameters"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/v1/auth/login [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input signInInput
	if err := decoder.Decode(&input); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	cookie, err := h.authMicroservice.SignIn(
		r.Context(),
		&proto.SignInInput{Mail: input.Mail, Password: input.Password},
	)
	status, ok := status.FromError(err)
	if ok {
		if status.Code() == codes.InvalidArgument {
			newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, status.Message())
			return
		}
		if status.Code() == codes.Unauthenticated {
			newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, status.Message())
			return
		}
		if status.Code() == codes.Internal {
			newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, status.Message())
			return
		}
	}
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, createCookie(cookie.Cookie))
	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary log out of account
// @Tags auth
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/v1/auth/logout [get]
func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "no session")
		return
	}
	_, err = h.authMicroservice.LogOut(
		r.Context(),
		&proto.Cookie{Cookie: session.Value},
	)

	status, ok := status.FromError(err)
	if ok {
		if status.Code() == codes.Internal {
			newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, status.Message())
			return
		}
	}
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	session.Path = "/"

	http.SetCookie(w, session)
	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary sign up account
// @Tags auth
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "Sign-up input user"
// @Success 200 {object} ClientResponseDto[idResponse]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/v1/auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input signUpInput
	if err := decoder.Decode(&input); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.authMicroservice.SignUp(
		r.Context(),
		&proto.SignUpInput{Mail: input.Mail, Password: input.Password, Name: input.Name},
	)

	status, ok := status.FromError(err)
	if ok {
		if status.Code() == codes.InvalidArgument {
			newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, status.Message())
			return
		}
		if status.Code() == codes.Internal {
			newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, status.Message())
			return
		}
	}
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, createCookie(userId.Cookie.Cookie))

	NewSuccessClientResponseDto(r.Context(), w, idResponse{Id: int(userId.Id)})
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
