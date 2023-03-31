package idiomatic

func Find[E any](fn func(E) bool, elems ...E) (E, bool) {
	ignore := func(_ int, elem E) bool {
		return fn(elem)
	}

	return FindI(ignore, elems...)
}

func FindRight[E any](fn func(E) bool, elems ...E) (E, bool) {
	return FindRight(fn, Reverse(elems...)...)
}

func FindI[E any](fn func(int, E) bool, elems ...E) (E, bool) {
	fn2 := func(i int, elem E) (bool, error) {
		return fn(i, elem), nil
	}
	output, ok, _ := FindErrI(fn2, elems...)
	return output, ok
}

func FindRightI[E any](fn func(int, E) bool, elems ...E) (E, bool) {
	return FindI(fn, Reverse(elems...)...)
}

func FindErr[E any](fn func(E) (bool, error), elems ...E) (E, bool, error) {
	fn2 := func(_ int, elem E) (bool, error) {
		return fn(elem)
	}
	return FindErrI(fn2, elems...)
}

func FindRightErr[E any](fn func(E) (bool, error), elems ...E) (E, bool, error) {
	return FindErr(fn, Reverse(elems...)...)
}

func FindErrI[E any](fn func(int, E) (bool, error), elems ...E) (E, bool, error) {
	var ok bool
	var found E
	var err error

	for i, elem := range elems {
		ok, err = fn(i, elem)

		if err != nil {
			break
		}

		if ok {
			ok = true
			found = elem
			break
		}
	}

	return found, ok, err
}

func FindRightErrI[E any](fn func(int, E) (bool, error), elems ...E) (E, bool, error) {
	return FindErrI(fn, Reverse(elems...)...)
}
