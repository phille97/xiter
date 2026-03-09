package xiter

import "iter"

func Buffer[T any](seq iter.Seq[T], bufferSize int) iter.Seq[T] {
	ch := make(chan T, bufferSize)
	done := make(chan struct{})
	go func() {
		defer close(ch)
		seq(func(v T) bool {
			select {
			case ch <- v:
				return true
			case <-done:
				return false
			}
		})
	}()

	return func(yield func(T) bool) {
		for v := range ch {
			if !yield(v) {
				close(done)
				break
			}
		}
	}
}

func Buffer2[T1, T2 any](seq iter.Seq2[T1, T2], bufferSize int) iter.Seq2[T1, T2] {
	type pair struct {
		v1 T1
		v2 T2
	}
	// converts the Seq2 to Seq of pairs, buffers it, then converts back to Seq2.
	// This way we can reuse the same Buffer function we use for Seq
	return Map12(
		Buffer(Map21(seq, func(v1 T1, v2 T2) pair {
			return pair{v1, v2}
		}), bufferSize),
		func(p pair) (T1, T2) {
			return p.v1, p.v2
		},
	)
}
