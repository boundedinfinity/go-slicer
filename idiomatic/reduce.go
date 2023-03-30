package idiomatic

func Reduce[I any](fn func(I, I) I, items ...I) (I, bool) {
	var output I
	var ok bool

	if len(items) > 0 {
		ok = true
		output = Fold(items[0], fn, items...)
	}

	return output, ok
}

func ReduceI[I any](fn func(int, I, I) I, items ...I) (I, bool) {
	var output I
	var ok bool

	if len(items) > 0 {
		ok = true
		output = FoldI(items[0], fn, items...)
	}

	return output, ok
}

func ReduceErr[I any](fn func(I, I) (I, error), items ...I) (I, bool, error) {
	var output I
	var ok bool
	var err error

	if len(items) > 0 {

		output, err = FoldErr(items[0], fn, items...)

		if err == nil {
			ok = true
		}
	}

	return output, ok, err
}

func ReduceErrI[I any](fn func(int, I, I) (I, error), items ...I) (I, bool, error) {
	var output I
	var ok bool
	var err error

	if len(items) > 0 {

		output, err = FoldErrI(items[0], fn, items...)

		if err == nil {
			ok = true
		}
	}

	return output, ok, err
}
