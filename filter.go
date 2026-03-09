package xiter

import "iter"

func Filter[T any](seq iter.Seq[T], fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(v T) bool {
			if fn(v) {
				return yield(v)
			}
			return true
		})
	}
}
