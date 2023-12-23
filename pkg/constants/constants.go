package constants

import "time"

var (
	Host = "https://umlaut-bmstu.me"
	// Host    = "http://localhost:8000"
	Message      = "message"
	Match        = "match"
	Banned       = 3
	Empty        = "empty"
	VkFields     = "id,photo_max,email,sex,bdate"
	CookieExpire = 30 * 24 * time.Hour
)
