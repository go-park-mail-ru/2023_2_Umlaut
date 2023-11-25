package utils

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
