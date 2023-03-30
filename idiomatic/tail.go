package idiomatic

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
