package slices

// Map []T1 to []T2 by mapper
func Map[T1 any, T2 any](slice []T1, mapper func(T1) T2) []T2 {
	mapped := make([]T2, 0, len(slice))
	for _, item := range slice {
		mapped = append(mapped, mapper(item))
	}
	return mapped
}
