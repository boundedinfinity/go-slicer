package idiomatic

import (
	c "golang.org/x/exp/constraints"
)

func Min[O c.Ordered](items ...O) (O, bool) {
	fn := func(acc, curr O) O {
		if acc < curr {
			return acc
		}

		return curr
	}

	return Reduce(fn, items...)
}

func MinBy[I any, O c.Ordered](fn func(I) O, items ...I) (I, bool) {
	by := func(acc, curr I) I {
		if fn(acc) < fn(curr) {
			return acc
		}

		return curr
	}

	return Reduce(by, items...)
}
