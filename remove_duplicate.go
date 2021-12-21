package slices

// RemoveDuplicate remove duplicate items
func RemoveDuplicate[T comparable](slice []T) []T {
	var filtered []T
	contains := make(map[T]bool, len(slice))
	for _, item := range slice {
		if !contains[item] {
			contains[item] = true
			filtered = append(filtered, item)
		}
	}
	return filtered
}
