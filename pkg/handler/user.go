package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
)

// @Summary get user information
// @Tags user
// @ID user
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 404,500 {object} ClientResponseDto[string]
// @Router /api/v1/user [get]
func (h *Handler) user(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyUserID).(int)
	currentUser, err := h.services.User.GetCurrentUser(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	session, _ := r.Cookie("session_id")
	jwtToken := NewJwtToken(secret)
	token, err := jwtToken.Create(session.Value, id, time.Now().Add(12*time.Hour).Unix())
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "csrf token creation error")
		return
	}
	w.Header().Set("X-Csrf-Token", token)
	w.Header().Set("Access-Control-Expose-Headers", "X-Csrf-Token")

	NewSuccessClientResponseDto(r.Context(), w, currentUser)
}

// @Summary update user
// @Tags user
// @ID user
// @Accept  json
// @Produce  json
// @Param input body model.User true "User data to update"
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user [post]
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	user.Id = r.Context().Value(keyUserID).(int)
	currentUser, err := h.services.User.UpdateUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, static.ErrAlreadyExists) {
			newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "account with this email already exists")
			return
		}
		if errors.Is(err, static.ErrInvalidUser) {
			newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid field for update")
			return
		}
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseDto(r.Context(), w, currentUser)
}

// @Summary update user photo
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user/photo [post]
func (h *Handler) updateUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyUserID).(int)
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, head, err := r.FormFile("file")
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	defer file.Close()

	_, err = h.services.User.CreateFile(r.Context(), id, file, head.Size)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary delete user photo
// @Tags user
// @Accept  json
// @Param input body deleteLink true "link for deleting file"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user/photo [delete]
func (h *Handler) deleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var link deleteLink
	if err := decoder.Decode(&link); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	id := r.Context().Value(keyUserID).(int)

	err := h.services.User.DeleteFile(r.Context(), id, link.Link)
	if err == static.ErrNoFiles {
		newErrorClientResponseDto(r.Context(), w, http.StatusNotFound, "This user has no photos")
		return
	}
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}
