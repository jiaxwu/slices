package slices

// FindIndex first index that meet the condition
func FindIndex[T any](slice []T, condition func(T) bool) int {
	for i, item := range slice {
		if condition(item) {
			return i
		}
	}
	return -1
}
