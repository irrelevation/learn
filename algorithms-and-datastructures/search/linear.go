package search

func Linear[T comparable](slice []T, value T) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return -1
}