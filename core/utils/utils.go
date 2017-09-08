package utils

import "strings"

func InSlice(key string, slice_str []string) (bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return false
	}

	for _, value := range slice_str {
		if value == key {
			return true
		}
	}
	return false
}
