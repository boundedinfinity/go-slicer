package idiomatic

// TODO: https://www.scala-lang.org/api/3.2.2/scala/collection/Iterable.html#empty-0

// Empty returns an empty slice of capacity with size.
func Empty[T any](size int) []T {
	return make([]T, size)
}

func IsEmpty[T any](items ...T) bool {
	if items == nil {
		return true
	}

	return len(items) == 0
}

func IsEmptyFn[T any](fn func(T) bool, items ...T) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}

	return true
}
