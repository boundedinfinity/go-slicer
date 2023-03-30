package idiomatic

func Head[T any](items ...T) (T, bool) {
	if len(items) > 0 {
		return items[0], true
	}

	var zero T
	return zero, false
}
