package slices

// ReduceRight slice to a value
func ReduceRight[T any](slice []T, reduce func(total T, item T, index int, slice []T) T, init T) T {
	for index := len(slice) - 1; index >= 0; index-- {
		init = reduce(init, slice[index], index, slice)
	}
	return init
}
