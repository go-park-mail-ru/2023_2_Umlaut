package constants

import "time"

var (
	Host         = "https://umlaut-bmstu.me"
	CookieExpire = 30 * 24 * time.Hour
	Message      = "message"
	Match        = "match"
	Banned       = 3
	ManGender    = 1
	WomanGender  = 0
	VkFields     = "id,photo_max,email,sex,bdate"
	ClientId     = "CLIENT_ID"
	ClientSecret = "CLIENT_SECRET"
	RedirectUrl  = "REDIRECT_URL"
)
