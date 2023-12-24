package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/convert"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func getVkOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv(constants.ClientId),
		ClientSecret: os.Getenv(constants.ClientSecret),
		RedirectURL:  os.Getenv(constants.RedirectUrl),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.vk.com/authorize",
			TokenURL: "https://oauth.vk.com/access_token",
		},
		Scopes: []string{"email"},
	}
}

// @Summary redirect to VK
// @Tags vk-auth
// @ID vk-login
// @Accept  json
// @Produce  json
// @Param invite_by query string false "invite_by value"
// @Router /api/v1/auth/vk-login [get]
func (h *Handler) vkLogin(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Request vkLogin",
		zap.String("ClientId", os.Getenv(constants.ClientId)),
	)
	vkOauthConfig := getVkOauthConfig()
	invite := r.URL.Query().Get("invite_by")
	url := vkOauthConfig.AuthCodeURL(invite)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// @Summary need call after redirect VK
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
	vkOauthConfig := getVkOauthConfig()
	code := r.URL.Query().Get("code")
	invite := r.URL.Query().Get("invite_by")

	h.logger.Info("Request vkSignUp",
		zap.String("code", code),
		zap.String("invite_by", invite),
	)

	token, err := vkOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "code error")
		return
	}
	vkUser, err := fetchVkUserData(token, vkOauthConfig)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "fetch data error")
		return
	}

	h.logger.Info("Request vkSignUp",
		zap.String("vkUser", fmt.Sprintf("%v", vkUser)),
	)

	user := convert.IntoCoreVkUser(vkUser)

	h.logger.Info("Request vkSignUp",
		zap.String("user", fmt.Sprintf("%v", user)),
	)

	id, err := h.services.Authorization.OAuth(r.Context(), user, invite)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("Request vkSignUp",
		zap.String("id", fmt.Sprintf("%v", id)),
	)

	cookie, err := h.services.Authorization.GenerateCookie(r.Context(), id)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "cookie error")
		return
	}

	h.logger.Info("Request vkSignUp",
		zap.String("cookie", fmt.Sprintf("%v", cookie)),
	)

	http.SetCookie(w, createCookie("session_id", cookie))
	dto.NewSuccessClientResponseDto(r.Context(), w, dto.IdResponse{Id: id})
}

func fetchVkUserData(token *oauth2.Token, vkOauthConfig *oauth2.Config) (dto.VkUser, error) {
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
