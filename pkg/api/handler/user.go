package handler

import (
	"errors"
	static2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/utils"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	id := r.Context().Value(static2.KeyUserID).(int)
	currentUser, err := h.services.User.GetCurrentUser(r.Context(), id)
	if errors.Is(err, static2.ErrBannedUser) {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	session, _ := r.Cookie("session_id")
	jwtToken := utils.NewJwtToken(static2.Secret)
	token, err := jwtToken.Create(session.Value, id, time.Now().Add(12*time.Hour).Unix())
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "csrf token creation error")
		return
	}
	w.Header().Set("X-Csrf-Token", token)
	w.Header().Set("Access-Control-Expose-Headers", "X-Csrf-Token")

	dto.NewSuccessClientResponseDto(r.Context(), w, currentUser)
}

// @Summary get user information by id
// @Tags user
// @ID userById
// @Accept  json
// @Param id path integer true "user ID"
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 404,500 {object} ClientResponseDto[string]
// @Router /api/v1/user/{id} [get]
func (h *Handler) userById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}

	currentUser, err := h.services.User.GetCurrentUser(r.Context(), id)
	currentUser.Mail = ""
	if errors.Is(err, static2.ErrBannedUser) {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, currentUser)
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var user core.User
	if err := user.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	user.Id = r.Context().Value(static2.KeyUserID).(int)
	currentUser, err := h.services.User.UpdateUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, static2.ErrAlreadyExists) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "account with this email already exists")
			return
		}
		if errors.Is(err, static2.ErrInvalidUser) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid field for update")
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, currentUser)
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
	id := r.Context().Value(static2.KeyUserID).(int)
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, head, err := r.FormFile("file")
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	defer file.Close()

	link, err := h.services.User.CreateFile(r.Context(), id, file, head.Size)
	if errors.Is(err, static2.ErrBannedUser) {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusForbidden, err.Error())
		return
	}
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, link)
}

// @Summary delete user photo
// @Tags user
// @Accept  json
// @Param input body deleteLink true "link for deleting file"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/user/photo [delete]
func (h *Handler) deleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var link dto.DeleteLink
	if err := link.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	id := r.Context().Value(static2.KeyUserID).(int)

	err = h.services.User.DeleteFile(r.Context(), id, link.Link)
	if errors.Is(err, static2.ErrBannedUser) {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusForbidden, err.Error())
		return
	}
	if errors.Is(err, static2.ErrNoFiles) {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusNotFound, "This user has no photos")
		return
	}
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary get user share link
// @Tags user
// @Accept  json
// @Success 200 {object} ClientResponseDto[shareCridentialsOutput]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/user/share [get]
func (h *Handler) getUserShareCridentials(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(static2.KeyUserID).(int)
	invitesCount, link, err := h.services.User.GetUserShareCridentials(r.Context(), id)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	share := dto.ShareCridentialsOutput{InvitesCount: invitesCount, ShareLink: link}

	dto.NewSuccessClientResponseDto(r.Context(), w, share)
}
