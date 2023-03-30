package idiomatic

func Reverse[E any](elems ...E) []E {
	l := len(elems)
	out := make([]E, l)

	for i := range elems {
		out[l-1-i] = elems[i]
	}

	return out
}
