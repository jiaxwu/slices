package slices

import "math/rand"

// Shuffle slice
func Shuffle[T any](slice []T) {
	for i := 0; i < len(slice); i++ {
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}
}
