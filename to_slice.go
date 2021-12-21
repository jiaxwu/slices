package slices

// ToSlice iterates map's key and value to generate a new slice
func ToSlice[Key comparable, Value any, T any](dict map[Key]Value, transform func(key Key, value Value) T) []T {
	list := make([]T, len(dict))
	var i = 0
	for key, value := range dict {
		list[i] = transform(key, value)
		i += 1
	}
	return list
}
