package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/convert"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"golang.org/x/oauth2"
	"net/http"
)

var (
	vkOauthConfig = &oauth2.Config{
		ClientID:     "51820172",
		ClientSecret: "BWzMJDcxBMQOkcZapM6V",
		RedirectURL:  "https://umlaut-bmstu.me/feed",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.vk.com/authorize",
			TokenURL: "https://oauth.vk.com/access_token",
		},
		Scopes: []string{"email"},
	}
)

// @Summary log out of account
// @Tags vk-auth
// @ID vk-login
// @Accept  json
// @Produce  json
// @Param invite_by query string false "invite_by value"
// @Router /api/v1/auth/vk-login [get]
func (h *Handler) vkLogin(w http.ResponseWriter, r *http.Request) {
	invite := r.URL.Query().Get("invite_by")
	url := vkOauthConfig.AuthCodeURL(invite)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// @Summary log out of account
// @Tags vk-auth
// @ID vk-sign-up
// @Accept  json
// @Produce  json
// @Param code query string true "code from oauth"
// @Param invite_by query string false "invite_by param"
// @Success 200 {object} ClientResponseDto[dto.IdResponse]
// @Failure 400,404,414 {object} ClientResponseDto[string]
// @Router /api/v1/auth/vk-sign-up [get]
func (h *Handler) vkSignUp(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	invite := r.URL.Query().Get("invite_by")
	token, err := vkOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "code error")
		return
	}
	vkUser, err := fetchVkUserData(token)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "fetch data error")
		return
	}
	user := convert.IntoCoreVkUser(vkUser)
	id, err := h.services.Authorization.OAuth(r.Context(), user, invite)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	cookie, err := h.services.Authorization.GenerateCookie(r.Context(), id)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "cookie error")
		return
	}

	http.SetCookie(w, createCookie("session_id", cookie))
	dto.NewSuccessClientResponseDto(r.Context(), w, dto.IdResponse{Id: id})
}

func fetchVkUserData(token *oauth2.Token) (dto.VkUser, error) {
	httpClient := vkOauthConfig.Client(context.Background(), token)

	resp, err := httpClient.Get(
		fmt.Sprintf("https://api.vk.com/method/users.get?fields=%s&v=5.131", constants.VkFields),
	)
	if err != nil {
		return dto.VkUser{}, fmt.Errorf("error fetching user data: %v", err)
	}
	defer resp.Body.Close()

	var userData struct {
		Response []dto.VkUser `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return dto.VkUser{}, fmt.Errorf("error decoding user data: %v", err)
	}

	if len(userData.Response) == 0 {
		return dto.VkUser{}, errors.New("no user data found")
	}

	return userData.Response[0], nil
}
