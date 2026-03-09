package xiter

import "iter"

func Map11[Tin any, Tout any](seq iter.Seq[Tin], fn func(Tin) Tout) iter.Seq[Tout] {
	return func(yield func(Tout) bool) {
		seq(func(v Tin) bool {
			return yield(fn(v))
		})
	}
}

func Map12[Tin, Tout1, Tout2 any](seq iter.Seq[Tin], fn func(Tin) (Tout1, Tout2)) iter.Seq2[Tout1, Tout2] {
	return func(yield func(Tout1, Tout2) bool) {
		seq(func(v Tin) bool {
			return yield(fn(v))
		})
	}
}

func Map22[Tin1, Tin2, Tout1, Tout2 any](seq iter.Seq2[Tin1, Tin2], fn func(Tin1, Tin2) (Tout1, Tout2)) iter.Seq2[Tout1, Tout2] {
	return func(yield func(Tout1, Tout2) bool) {
		seq(func(v1 Tin1, v2 Tin2) bool {
			return yield(fn(v1, v2))
		})
	}
}

func Map21[Tin1, Tin2, Tout any](seq iter.Seq2[Tin1, Tin2], fn func(Tin1, Tin2) Tout) iter.Seq[Tout] {
	return func(yield func(Tout) bool) {
		seq(func(v1 Tin1, v2 Tin2) bool {
			return yield(fn(v1, v2))
		})
	}
}
