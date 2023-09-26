package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) feedHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: проверить сессию, если нет, то разлогинить (наверное)
	if r.Method == http.MethodGet {
		userId := 0 //TODO: id из ссессии достать надо по логике
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
