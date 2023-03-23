package slicer

func Flatten[T any](listOfList ...[]T) []T {
	outputs := make([]T, 0)

	for _, list := range listOfList {
		outputs = append(outputs, list...)
	}

	return outputs
}
