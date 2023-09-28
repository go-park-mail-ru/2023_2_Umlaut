package handler

import (
	"crypto/sha1"
	"math/rand"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"io"
	"net/http"
	"time"
)

const (
	salt = "LKJksdfbdkjhgk213234dfhLKJnlkj"
)

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ctx := r.Context()
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

		SID := generateCookie()
		if err := h.Repositories.SetSession(ctx, SID, user.Id, 10 * time.Hour); err == nil {
			cookie := &http.Cookie{
				Name:    "session_id",
				Value:   SID,
				Expires: time.Now().Add(10 * time.Hour),
			}
			http.SetCookie(w, cookie)
		}

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
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, "no session", http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	if err := h.Repositories.DeleteSession(ctx, session.Value); err != nil {
		http.Error(w, "Invalid cookie deletion", http.StatusInternalServerError)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ctx := r.Context()
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

		SID := generateCookie()
		if err := h.Repositories.SetSession(ctx, SID, id, 10 * time.Hour); err == nil {
			cookie := &http.Cookie{
				Name:    "session_id",
				Value:   SID,
				Expires: time.Now().Add(10 * time.Hour),
			}
			http.SetCookie(w, cookie)
		}

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

func generateCookie() string {
	return randStringRunes(32)
}

func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
