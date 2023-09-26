package handler

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"io"
	"net/http"
)

const (
	salt = "LKJksdfbdkjhgk213234dfhLKJnlkj"
)

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: попробовать достать сессию из редиса
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		var requestData struct {
			Mail     string `json:"mail"`
			Password string `json:"password"`
		}

		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		user, err := h.Repositories.GetUser(requestData.Mail, generatePasswordHash(requestData.Password))
		if err == nil {
			jsonResponse, _ := json.Marshal(user)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
			return
		}
	}
	response := map[string]string{
		"error": "Authentication failed",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}

func (h *Handler) logoutHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: очистить сессию
	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: попробовать достать сессию из редиса, вдруг уже зареган пользователь
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		user := model.User{}

		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}
		user.PasswordHash = generatePasswordHash(user.PasswordHash)
		id, err := h.Repositories.CreateUser(user)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response := map[string]int{
				"id": id,
			}
			jsonResponse, _ := json.Marshal(response)
			w.Write(jsonResponse)
			return
		}
	}
	response := map[string]string{
		"error": "Registration failed",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
