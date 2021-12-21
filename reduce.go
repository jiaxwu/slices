package slices

// Reduce slice to a value
func Reduce[T any](slice []T, reduce func(total T, item T, index int, slice []T) T, init T) T {
	for index := 0; index < len(slice); index++ {
		init = reduce(init, slice[index], index, slice)
	}
	return init
}
