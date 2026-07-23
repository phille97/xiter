package xiter

import (
	"iter"
	"time"
)

func SleepBetween[T any](seq iter.Seq[T], duration time.Duration) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(v T) bool {
			if !yield(v) {
				return false
			}
			time.Sleep(duration)
			return true
		})
	}
}

func SleepBetween2[T1, T2 any](seq iter.Seq2[T1, T2], duration time.Duration) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		seq(func(v1 T1, v2 T2) bool {
			if !yield(v1, v2) {
				return false
			}
			time.Sleep(duration)
			return true
		})
	}
}
