package slices

// ToMap iterates the slice to map following the keyFunc and valueFunc
func ToMap[T any, Key comparable, Value any](slice []T, keyFunc func(T) Key, valueFunc func(T) Value) map[Key]Value {
	dict := make(map[Key]Value)
	for _, value := range slice {
		dict[keyFunc(value)] = valueFunc(value)
	}
	return dict
}
