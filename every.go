package slices

// Every item is meet the condition
func Every[T any](slice []T, condition func(T) bool) bool {
	for _, item := range slice {
		if !condition(item) {
			return false
		}
	}
	return true
}
