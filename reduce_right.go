package slices

// ReduceRight slice to a value
func ReduceRight[T any](slice []T, reduce func(cur T, item T) T, init T) T {
	for i := len(slice) - 1; i >= 0; i-- {
		init = reduce(init, slice[i])
	}
	return init
}
