package idiomatic

func Each[E any](fn func(E), elems ...E) {
	for _, elem := range elems {
		fn(elem)
	}
}

func EachI[E any](fn func(int, E), elems ...E) {
	for i, elem := range elems {
		fn(i, elem)
	}
}

func EachErr[E any](fn func(E) error, elems ...E) error {
	for _, elem := range elems {
		if err := fn(elem); err != nil {
			return err
		}
	}

	return nil
}

func EachErrI[E any](fn func(int, E) error, elems ...E) error {
	for i, elem := range elems {
		if err := fn(i, elem); err != nil {
			return err
		}
	}

	return nil
}
