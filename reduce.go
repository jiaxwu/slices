package slices

// Reduce slice to a value
func Reduce[T any](slice []T, reduce func(cur T, item T) T, init T) T {
	for i := 0; i < len(slice); i++ {
		init = reduce(init, slice[i])
	}
	return init
}
