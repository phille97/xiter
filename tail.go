package xiter

import "iter"

func Tail[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		buffer := make([]T, 0, n)
		seq(func(v T) bool {
			if len(buffer) < n {
				buffer = append(buffer, v)
			} else {
				copy(buffer, buffer[1:])
				buffer[len(buffer)-1] = v
			}
			return true
		})
		for _, v := range buffer {
			if !yield(v) {
				break
			}
		}
	}
}
