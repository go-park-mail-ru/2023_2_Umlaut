package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

// @Summary get user information
// @Tags user
// @ID user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Failure 401,404 {object} errorResponse
// @Router /api/user [get]
func (h *Handler) user(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}

	id, err := h.services.GetSessionValue(r.Context(), session.Value)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	currentUser, err := h.services.GetCurrentUser(r.Context(), id)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse, _ := json.Marshal(currentUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// @Summary update user
// @Tags user
// @ID user
// @Accept  json
// @Produce  json
// @Param input body model.User true "User data to update"
// @Success 200 {object} model.User
// @Failure 401,404 {object} errorResponse
// @Router /api/user [post]
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}

	id, err := h.services.GetSessionValue(r.Context(), session.Value)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	user.Id = id
	currentUser, err := h.services.UpdateUser(r.Context(), user)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse, _ := json.Marshal(currentUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// @Summary update user photo
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200
// @Failure 401,404 {object} errorResponse
// @Router /api/user/photo [post]
func (h *Handler) updateUserPhoto(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}

	id, err := h.services.GetSessionValue(r.Context(), session.Value)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	r.ParseMultipartForm(5 * 1024 * 1025)
	file, head, err := r.FormFile("file")
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}
	defer file.Close()

	_, err = h.services.UpdatePhoto(r.Context(), id, model.ImageUnit{
		FileLoad:    file,
		Size:        head.Size,
		ContentType: r.Header.Get("Content-Type"),
	})
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	//TODO:: обновление бд с ссылкой на фото
}
