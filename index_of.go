package slices

// IndexOf slice[i] == item
// you can set start to 0, if you want to start from beginning
func IndexOf[T comparable](slice []T, item T, start int) int {
	length := len(slice)
	if start >= length {
		return -1
	}
	var i int
	if start+length > 0 {
		i = (start + length) % length
	}
	for ; i < length; i++ {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
