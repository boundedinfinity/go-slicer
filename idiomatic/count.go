package idiomatic

func Count[E any](fn func(E) bool, elems ...E) int {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem), nil
	}
	output, _ := CountErrI(fn2, elems...)
	return output
}

func CountI[E any](fn func(int, E) bool, elems ...E) int {
	fn2 := func(i int, elem E) (bool, error) {
		return fn(i, elem), nil
	}
	output, _ := CountErrI(fn2, elems...)
	return output
}

func CountErr[E any](fn func(E) (bool, error), elems ...E) (int, error) {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem)
	}
	return CountErrI(fn2, elems...)
}

func CountErrI[E any](fn func(int, E) (bool, error), elems ...E) (int, error) {
	var count int
	var err error
	var ok bool

	for i, elem := range elems {
		ok, err = fn(i, elem)

		if err != nil {
			break
		}

		if ok {
			count++
		}
	}

	return count, err
}
