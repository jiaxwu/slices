package slices

// Some item is meet the condition
func Some[T any](slice []T, condition func(T) bool) bool {
	for _, item := range slice {
		if condition(item) {
			return true
		}
	}
	return false
}
