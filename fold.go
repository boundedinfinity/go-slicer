package slicer

func Fold[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	acc := initial

	for _, item := range items {
		acc = fn(acc, item)
	}

	return acc
}

func FoldErr[I any, O any](initial O, fn func(O, I) (O, error), items ...I) (O, error) {
	var err error
	acc := initial

	for _, item := range items {
		acc, err = fn(acc, item)

		if err != nil {
			break
		}
	}

	return acc, err
}
