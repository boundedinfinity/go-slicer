package idiomatic

func Fold[I any, O any](initial O, fn func(O, I) O, items ...I) O {
	acc := initial

	for _, item := range items {
		acc = fn(acc, item)
	}

	return acc
}

func FoldI[I any, O any](initial O, fn func(int, O, I) O, items ...I) O {
	acc := initial

	for i, item := range items {
		acc = fn(i, acc, item)
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

func FoldErrI[I any, O any](initial O, fn func(int, O, I) (O, error), items ...I) (O, error) {
	var err error
	acc := initial

	for i, item := range items {
		acc, err = fn(i, acc, item)

		if err != nil {
			break
		}
	}

	return acc, err
}
