package convert

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"log"
	"time"
)

func IntoCoreVkUser(vkUser dto.VkUser) core.User {

	var userGender, preferGender *int
	if vkUser.Sex > 0 {
		gender := vkUser.Sex - 1
		userGender = &gender
		log.Printf("[VK] %d gender == constants.ManGender: %v", gender, gender == constants.ManGender)
		log.Printf("[VK] %d gender == constants.WomanGender: %v", gender, gender == constants.WomanGender)
		if gender == constants.ManGender {
			preferGender = &constants.WomanGender
			log.Printf("[VK] %d", *preferGender)
		} else {
			preferGender = &constants.ManGender
			log.Printf("[VK] %d", *preferGender)
		}
	}

	var data *time.Time
	birthDate, err := time.Parse("2.1.2006", vkUser.BirthDate)
	if err == nil {
		data = &birthDate
	}

	var photo []string
	if len(vkUser.Photo) > 0 {
		photo = append(photo, vkUser.Photo)
	}

	return core.User{
		Name:         vkUser.Name,
		Mail:         vkUser.Email,
		UserGender:   userGender,
		PreferGender: preferGender,
		Birthday:     data,
		ImagePaths:   &photo,
		OauthId:      &vkUser.Id,
	}
}
