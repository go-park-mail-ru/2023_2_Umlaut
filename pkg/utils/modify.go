package utils

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
)

func ModifyString(data *string) string {
	if data == nil {
		return ""
	}
	return *data
}

func ModifyInt(data *int) int32 {
	if data == nil {
		return 0
	}
	return int32(*data)
}

func ToPtrInt(number int) *int {
	return &number
}

func ToPtrString(str string) *string {
	return &str
}

//func ToPtrTime(timing time.Time) *time.Time {
//	return &timing
//}

func ModifyArray(data *[]string) []string {
	if data == nil {
		return []string{}
	}
	var result []string
	result = append(result, *data...)
	return result
}

func ModifyLikedMap(likedMap []*proto.LikedMap) map[string]int32 {
	result := make(map[string]int32)
	for _, item := range likedMap {
		result[item.Liked] = item.Count
	}
	return result
}

func ModifyNeedFixMap(needFixMap []*proto.NeedFixMap) map[string]core.NeedFixObject {
	result := make(map[string]core.NeedFixObject)
	for _, item := range needFixMap {
		result[item.NeedFix] = core.NeedFixObject{
			Count:      item.NeedFixObject.Count,
			CommentFix: item.NeedFixObject.CommentFix,
		}
	}
	return result
}
