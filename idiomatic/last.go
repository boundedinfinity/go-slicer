package idiomatic

func Last[T any](items ...T) (T, bool) {
	l := len(items)

	if l > 0 {
		return items[l-1], true
	}

	var zero T
	return zero, false
}
