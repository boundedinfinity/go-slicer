package idiomatic

func Map[A any, B any](fn func(A) B, items ...A) []B {
	bs := make([]B, len(items))

	for i, a := range items {
		bs[i] = fn(a)
	}

	return bs
}

func MapI[A any, B any](fn func(int, A) B, items ...A) []B {
	bs := make([]B, len(items))

	for i, a := range items {
		bs[i] = fn(i, a)
	}

	return bs
}

func MapErr[A any, B any](fn func(A) (B, error), items ...A) ([]B, error) {
	var err error
	var b B
	bs := make([]B, 0)

	for _, a := range items {
		b, err = fn(a)

		if err != nil {
			break
		}

		bs = append(bs, b)
	}

	return bs, err
}

func MapErrI[A any, B any](fn func(int, A) (B, error), items ...A) ([]B, error) {
	var err error
	var b B
	bs := make([]B, 0)

	for i, a := range items {
		b, err = fn(i, a)

		if err != nil {
			break
		}

		bs = append(bs, b)
	}

	return bs, err
}
