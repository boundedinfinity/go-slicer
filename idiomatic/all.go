package idiomatic

func All[E any](fn func(E) bool, elems ...E) bool {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem), nil
	}
	output, _ := AllErrI(fn2, elems...)
	return output
}

func AllI[E any](fn func(int, E) bool, elems ...E) bool {
	fn2 := func(i int, elem E) (bool, error) {
		return fn(i, elem), nil
	}
	output, _ := AllErrI(fn2, elems...)
	return output
}

func AllErr[E any](fn func(E) (bool, error), elems ...E) (bool, error) {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem)
	}
	return AllErrI(fn2, elems...)
}

func AllErrI[E any](fn func(int, E) (bool, error), elems ...E) (bool, error) {
	if len(elems) < 1 {
		return false, nil
	}

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
