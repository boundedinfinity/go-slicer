package slicer

func Filter[T comparable](match T, items ...T) []T {
	output := make([]T, 0)

	for _, item := range items {
		if item == match {
			output = append(output, item)
		}
	}

	return output
}

func FilterFn[T any](fn func(T) bool, items ...T) []T {
	output := make([]T, 0)

	for _, item := range items {
		if fn(item) {
			output = append(output, item)
		}
	}

	return output
}

func FilterFnErr[T any](fn func(T) (bool, error), items ...T) ([]T, error) {
	output := make([]T, 0)

	for _, item := range items {
		ok, err := fn(item)

		if err != nil {
			return output, err
		}

		if ok {
			output = append(output, item)
		}
	}

	return output, nil
}

func All[T any](fn func(T) bool, items ...T) bool {
	for _, item := range items {
		if !fn(item) {
			return false
		}
	}

	return true
}

func AllErr[T any](fn func(T) (bool, error), items ...T) (bool, error) {
	for _, item := range items {
		ok, err := fn(item)

		if err != nil {
			return false, err
		}

		if !ok {
			return false, nil
		}
	}

	return true, nil
}

func Any[T any](fn func(T) bool, items ...T) bool {
	for _, item := range items {
		if fn(item) {
			return true
		}
	}

	return false
}

func AnyErr[T any](fn func(T) (bool, error), items ...T) (bool, error) {
	for _, item := range items {
		ok, err := fn(item)

		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}
