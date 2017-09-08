package utils

func InSlice(key string, slice_str []string) (bool) {
	for _, value := range slice_str {
		if value == key {
			return true
		}
	}
	return false
}
