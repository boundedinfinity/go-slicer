package idiomatic

func Any[E any](fn func(E) bool, elems ...E) bool {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem), nil
	}
	output, _ := AnyErrI(fn2, elems...)
	return output
}

func AnyI[E any](fn func(int, E) bool, elems ...E) bool {
	fn2 := func(i int, elem E) (bool, error) {
		return fn(i, elem), nil
	}
	output, _ := AnyErrI(fn2, elems...)
	return output
}

func AnyErr[E any](fn func(E) (bool, error), elems ...E) (bool, error) {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem)
	}
	return AnyErrI(fn2, elems...)
}

func AnyErrI[E any](fn func(int, E) (bool, error), elems ...E) (bool, error) {
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
