package idiomatic

import (
	"fmt"
	"strings"
)

func Join[S ~string](sep S, elems ...fmt.Stringer) S {
	stringElems := make([]string, len(elems))

	for i, elem := range elems {
		stringElems[i] = elem.String()
	}

	s := strings.Join(stringElems, string(sep))

	return S(s)
}

func JoinFn[T any, S ~string](sep S, fn func(T) fmt.Stringer, elems ...T) S {
	stringElems := make([]string, len(elems))

	for i, elem := range elems {
		stringElems[i] = fn(elem).String()
	}

	s := strings.Join(stringElems, string(sep))

	return S(s)
}

func JoinFnErr[T any, S ~string](sep S, fn func(T) (fmt.Stringer, error), elems ...T) (S, error) {
	stringElems := make([]string, len(elems))

	for i, elem := range elems {
		strer, err := fn(elem)

		if err != nil {
			return S(""), err
		}

		stringElems[i] = strer.String()
	}

	s := strings.Join(stringElems, string(sep))

	return S(s), nil
}
