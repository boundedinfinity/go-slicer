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
	fn2 := func(i, count int, elem E) (int, error) {
		ok, err := fn(i, elem)

		if err != nil {
			return count, err
		}

		if ok {
			count++
		}

		return count, nil
	}

	return FoldErrI(0, fn2, elems...)
}
