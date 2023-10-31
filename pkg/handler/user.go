package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/gorilla/mux"
)

// @Summary get user information
// @Tags user
// @ID user
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user [get]
func (h *Handler) user(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyUserID).(int)
	currentUser, err := h.services.GetCurrentUser(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	session, _ := r.Cookie("session_id")
	jwtToken := NewJwtToken(h.ctx, secret)
	token, err := jwtToken.Create(session.Value, id, time.Now().Add(12*time.Hour).Unix())
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, "csrf token creation error")
		return
	}
	w.Header().Set("X-CSRF-Token", token)

	NewSuccessClientResponseDto(h.ctx, w, currentUser)
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
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "invalid input body")
		return
	}

	user.Id = r.Context().Value(keyUserID).(int)
	currentUser, err := h.services.UpdateUser(r.Context(), user)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseDto(h.ctx, w, currentUser)
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
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "invalid input body")
		return
	}
	defer file.Close()
	fileName, err := h.services.CreateFile(r.Context(), id, file, head.Size)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.UpdateUserPhoto(r.Context(), id, fileName)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(h.ctx, w, "")
}

// @Summary get user photo
// @Tags user
// @Accept  json
// @Param id path integer true "User ID"
// @Success 200
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user/{id}/photo [get]
func (h *Handler) getUserPhoto(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "invalid params")
		return
	}

	currentUser, err := h.services.GetCurrentUser(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	if currentUser.ImagePath == nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "This user has no photos")
		return
	}
	buffer, contentType, err := h.services.GetFile(r.Context(), id, *currentUser.ImagePath)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(buffer)
}

// @Summary delete user photo
// @Tags user
// @Accept  json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user/photo [delete]
func (h *Handler) deleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(keyUserID).(int)
	currentUser, err := h.services.GetCurrentUser(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	if currentUser.ImagePath == nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "This user has no photos")
		return
	}
	err = h.services.DeleteFile(r.Context(), id, *currentUser.ImagePath)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(h.ctx, w, "")
}
