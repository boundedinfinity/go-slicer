package idiomatic

func ignoreIndex0[I any](fn func(I)) func(int, I) {
	return func(_ int, i I) {
		fn(i)
	}
}

func ignoreIndex1[I any, O any](fn func(I) O) func(int, I) O {
	return func(_ int, i I) O {
		return fn(i)
	}
}

func ignoreIndex1Err[I any, O any](fn func(I) (O, error)) func(int, I) (O, error) {
	return func(_ int, i I) (O, error) {
		return fn(i)
	}
}

func stripErr1[I any](i I, err error) I {
	return i
}

func stripErr2[I1 any, I2 any](i1 I1, i2 I2, err error) (I1, I2) {
	return i1, i2
}
