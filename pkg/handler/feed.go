package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary feedHandler
// @Tags feed
// @Description feed
// @ID feed
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/feed [get]
func (h *Handler) feedHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, "no session", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		ctx := r.Context()
		id, err := h.Repositories.GetSession(ctx, session.Value)
		if err != nil {
			http.Error(w, "Redis server is unavailable", http.StatusInternalServerError)
			return
		}
		userId := id
		user, _ := h.Repositories.GetUserById(userId)
		nextUser, err := h.Repositories.GetNextUser(user)
		if err == nil {
			jsonResponse, _ := json.Marshal(nextUser)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
			return
		}
	}
	response := map[string]string{
		"error": "Failed",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}
