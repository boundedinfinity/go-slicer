package idiomatic

func Drop[E any](count int, elems ...E) ([]E, bool) {
	var ok bool
	l := len(elems)
	out := make([]E, 0)

	if l < count+1 {
		ok = false
	} else {
		out = append(out, elems[count+1:l-1]...)
	}

	return out, ok
}

func DropRight[E any](count int, elems ...E) ([]E, bool) {
	var ok bool
	var out []E
	l := len(elems)

	if l < count+1 {
		out := Reverse(elems...)
		out, ok = Drop(count, out...)
		out = Reverse(out...)
	}

	return out, ok
}

func DropWhile[E any](fn func(E) bool, elems ...E) []E {
	out := make([]E, 0)
	start := -1

	for i, elem := range elems {
		if fn(elem) {
			continue
		}

		start = i
	}

	if start >= 0 {
		out = elems[start:]
	}

	return out
}

func DropRightWhile[E any](fn func(E) bool, elems ...E) []E {
	out := Reverse(elems...)
	out = DropWhile(fn, out...)
	out = Reverse(out...)

	return out
}
