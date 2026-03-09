package xiter

import "iter"

func Find[T comparable](seq iter.Seq[T], x T) bool {
	for v := range seq {
		if v == x {
			return true
		}
	}
	return false
}
