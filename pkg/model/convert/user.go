package convert

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"time"
)

func IntoCoreVkUser(vkUser dto.VkUser) core.User {

	var gender int
	if vkUser.Sex > 0 {
		gender = vkUser.Sex - 1
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
		Name:       vkUser.Name,
		Mail:       vkUser.Email,
		UserGender: &gender,
		Birthday:   data,
		ImagePaths: &photo,
		OauthId:    &vkUser.Id,
	}
}
