package idiomatic

func Uniq[E comparable](elems ...E) []E {
	m := make(map[E]bool)
	out := make([]E, 0)

	for _, elem := range elems {
		if _, ok := m[elem]; !ok {
			m[elem] = true
		}
	}

	for k := range m {
		out = append(out, k)
	}

	return out
}

func UniqBy[E any, C comparable](fn func(E) C, elems ...E) []E {
	return UniqByI(ignoreIndex1(fn), elems...)
}

func UniqByI[E any, C comparable](fn func(int, E) C, elems ...E) []E {
	wrapper := func(i int, elem E) (C, error) {
		return fn(i, elem), nil
	}

	return stripErr1(UniqByErrI(wrapper, elems...))
}

func UniqByErr[E any, C comparable](fn func(E) (C, error), elems ...E) ([]E, error) {
	return UniqByErrI(ignoreIndex1Err(fn), elems...)
}

func UniqByErrI[E any, C comparable](fn func(int, E) (C, error), elems ...E) ([]E, error) {
	m := make(map[C]E)
	out := make([]E, 0)

	for i, elem := range elems {
		comp, err := fn(i, elem)

		if err != nil {
			return out, err
		}

		if _, ok := m[comp]; !ok {
			m[comp] = elem
		}
	}

	for _, v := range m {
		out = append(out, v)
	}

	return out, nil
}
