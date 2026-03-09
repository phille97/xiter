package xiter

import "iter"

func Head[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		seq(func(v T) bool {
			if count < n {
				count++
				return yield(v)
			}
			return false
		})
	}
}
