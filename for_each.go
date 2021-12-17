package slices

// ForEach item execute action
func ForEach[T any](slice []T, action func(T)) {
	for _, item := range slice {
		action(item)
	}
}
