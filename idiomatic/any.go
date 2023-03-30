package idiomatic

func Any[T any](fn func(T) bool, elems ...T) bool {
	for _, elem := range elems {
		if fn(elem) {
			return true
		}
	}

	return false
}

func AnyI[T any](fn func(int, T) bool, elems ...T) bool {
	for i, elem := range elems {
		if fn(i, elem) {
			return true
		}
	}

	return false
}

func AnyErr[T any](fn func(T) (bool, error), elems ...T) (bool, error) {
	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}

func AnyErrI[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}
