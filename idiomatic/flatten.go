package idiomatic

func Flatten[T any](listOfList ...[]T) []T {
	outputs := make([]T, 0)

	for _, list := range listOfList {
		outputs = append(outputs, list...)
	}

	return outputs
}

func FlattenFn[T any](fn func([]T), listOfList ...[]T) []T {
	outputs := make([]T, 0)

	for _, list := range listOfList {
		fn(list)
		outputs = append(outputs, list...)
	}

	return outputs
}

func FlattenFnI[T any](fn func(int, []T), listOfList ...[]T) []T {
	outputs := make([]T, 0)

	for i, list := range listOfList {
		fn(i, list)
		outputs = append(outputs, list...)
	}

	return outputs
}
