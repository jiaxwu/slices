package slices

// Find first item that meet the condition
func Find[T any](slice []T, condition func(T) bool) (T, bool) {
	for _, item := range slice {
		if condition(item) {
			return item, true
		}
	}
	var t T
	return t, false
}
