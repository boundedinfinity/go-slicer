package idiomatic

func Count[E any](fn func(E) bool, elems ...E) int {
	wrapper := func(_ int, elem E) bool {
		return fn(elem)
	}

	return CountI(wrapper, elems...)
}

func CountI[E any](fn func(int, E) bool, elems ...E) int {
	var count int

	for i, elem := range elems {
		if fn(i, elem) {
			count++
		}
	}

	return count
}

func CountErr[E any](fn func(E) (bool, error), elems ...E) (int, error) {
	wrapper := func(_ int, elem E) (bool, error) {
		return fn(elem)
	}

	return CountErrI(wrapper, elems...)
}

func CountErrI[E any](fn func(int, E) (bool, error), elems ...E) (int, error) {
	var count int
	var err error

	for i, elem := range elems {
		ok, err := fn(i, elem)

		if err != nil {
			return count, err
		}

		if ok {
			count++
		}
	}

	return count, err
}
