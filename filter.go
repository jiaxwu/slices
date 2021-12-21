package slices

// Filter items that are not meet the condition
func Filter[T any](slice []T, condition func(T) bool) []T {
	var filtered []T
	for _, item := range slice {
		if condition(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func FilterWithIndex[T any](slice []T, condition func(T, int) bool) []T {
	var filtered []T
	for index, item := range slice {
		if condition(item, index) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
