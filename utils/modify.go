package utils

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
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

func ModifyArray(data *[]string) []string {
	if data == nil {
		return []string{}
	}
	var result []string
	for _, path := range *data {
		result = append(result, path)
	}
	return result
}

func ModifyLikedMap(likedMap []*proto.LikedMap) map[string]int32 {
	result := make(map[string]int32)
	for _, item := range likedMap {
		result[item.Liked] = item.Count
	}
	return result
}

func ModifyNeedFixMap(needFixMap []*proto.NeedFixMap) map[string]model.NeedFixObject {
	result := make(map[string]model.NeedFixObject)
	for _, item := range needFixMap {
		result[item.NeedFix] = model.NeedFixObject{
			Count:      item.NeedFixObject.Count,
			CommentFix: item.NeedFixObject.CommentFix,
		}
	}
	return result
}
