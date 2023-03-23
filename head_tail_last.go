package slicer

func Head[T any](items ...T) (T, bool) {
	if len(items) > 0 {
		return items[0], true
	}

	var zero T
	return zero, false
}

func Tail[T any](items ...T) ([]T, bool) {
	tail := make([]T, 0)

	for i, item := range items {
		if i == 0 {
			continue
		}

		tail = append(tail, item)
	}

	return tail, len(tail) > 0
}

func Last[T any](items ...T) (T, bool) {
	l := len(items)

	if l > 0 {
		return items[l-1], true
	}

	var zero T
	return zero, false
}
