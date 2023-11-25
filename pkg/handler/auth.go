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

// @Summary log in to admin
// @Tags auth
// @ID adminLogin
// @Accept  json
// @Produce  json
// @Param input body signInInput true "Sign-in input parameters"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,404 {object} ClientResponseDto[string]
// @Router /api/v1/auth/admin [post]
func (h *Handler) LogInAdmin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input signInInput
	if err := decoder.Decode(&input); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	cookie, err := h.authMicroservice.LogInAdmin(
		r.Context(),
		&proto.SignInInput{Mail: input.Mail, Password: input.Password},
	)

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	http.SetCookie(w, createCookie("admin_session_id", cookie.Cookie))
	NewSuccessClientResponseDto(r.Context(), w, "")
}

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

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	http.SetCookie(w, createCookie("session_id", cookie.Cookie))
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

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
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

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
	}

	http.SetCookie(w, createCookie("session_id", userId.Cookie.Cookie))

	NewSuccessClientResponseDto(r.Context(), w, idResponse{Id: int(userId.Id)})
}

func createCookie(name, SID string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		//SameSite: http.SameSiteNoneMode,
		//Secure:   true,
	}
}

func parseError(err error) (int, string) {
	status, ok := status.FromError(err)
	if ok {
		switch status.Code() {
		case codes.InvalidArgument:
			return http.StatusBadRequest, status.Message()
		case codes.Unauthenticated:
			return http.StatusUnauthorized, status.Message()
		case codes.Internal:
			return http.StatusInternalServerError, status.Message()
		}
	}
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return 200, ""
}
