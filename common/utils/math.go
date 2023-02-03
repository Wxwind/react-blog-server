package utils

func IsContain[T string | int](items []T, item T) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
