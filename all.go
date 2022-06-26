package slices

// All item is meet the condition
func All[T any](slice []T, condition func(item T, index int, slice []T) bool) bool {
	for index, item := range slice {
		if !condition(item, index, slice) {
			return false
		}
	}
	return true
}
