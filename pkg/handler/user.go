package handler

import (
	"encoding/json"
	"errors"
	"net/http"
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
	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusBadRequest, "Failed")
	}

	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}

	currentUser, err := h.services.GetCurrentUser(r.Context(), session.Value)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse, _ := json.Marshal(currentUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
