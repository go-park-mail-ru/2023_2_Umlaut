package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

// @Summary feed
// @Tags feed
// @Description Next user for feed
// @ID feed
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		newErrorResponse(w, http.StatusBadRequest, "Failed")
	}

	session, err := r.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		newErrorResponse(w, http.StatusUnauthorized, "no session")
		return
	}

	ctx := r.Context()
	id, err := h.Repositories.GetSession(ctx, session.Value)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Redis server is unavailable")
		return
	}

	user, _ := h.Repositories.GetUserById(id)
	nextUser, err := h.Repositories.GetNextUser(user)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse, _ := json.Marshal(nextUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
