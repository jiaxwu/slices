package slices

// LastIndexOf slice[i] == item
// you can set start to len(slice), if you want to start from last
func LastIndexOf[T comparable](slice []T, item T, start int) int {
	length := len(slice)
	if start < -length {
		return -1
	}
	i := length - 1
	if start-length < 0 {
		i = (start + length) % length
	}
	for ; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
