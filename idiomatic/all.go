package idiomatic

func All[T any](fn func(T) bool, elems ...T) bool {
	for _, elem := range elems {
		if !fn(elem) {
			return false
		}
	}

	return true
}

func AllI[T any](fn func(int, T) bool, elems ...T) bool {
	for i, elem := range elems {
		if !fn(i, elem) {
			return false
		}
	}

	return true
}

func AllErr[T any](fn func(T) (bool, error), elems ...T) (bool, error) {
	for _, elem := range elems {
		ok, err := fn(elem)

		if err != nil {
			return false, err
		}

		if !ok {
			return false, nil
		}
	}

	return true, nil
}

func AllErrI[T any](fn func(int, T) (bool, error), elems ...T) (bool, error) {
	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return false, err
		}

		if !ok {
			return false, nil
		}
	}

	return true, nil
}
