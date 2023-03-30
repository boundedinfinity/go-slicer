package idiomatic

func Filter[E comparable](match E, elems ...E) []E {
	output := make([]E, 0)

	for _, elem := range elems {
		if elem == match {
			output = append(output, elem)
		}
	}

	return output
}

func FilterFn[E any](fn func(E) bool, elems ...E) []E {
	output := make([]E, 0)

	for _, elem := range elems {
		if fn(elem) {
			output = append(output, elem)
		}
	}

	return output
}

func FilterFnI[T any](fn func(int, T) bool, elems ...T) []T {
	output := make([]T, 0)

	for i, elem := range elems {
		if fn(i, elem) {
			output = append(output, elem)
		}
	}

	return output
}

func FilterFnErr[T any](fn func(T) (bool, error), elems ...T) ([]T, error) {
	output := make([]T, 0)

	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return output, err
		}

		if ok {
			output = append(output, elem)
		}
	}

	return output, nil
}

func FilterFnErrI[T any](fn func(int, T) (bool, error), elems ...T) ([]T, error) {
	output := make([]T, 0)

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return output, err
		}

		if ok {
			output = append(output, elem)
		}
	}

	return output, nil
}
