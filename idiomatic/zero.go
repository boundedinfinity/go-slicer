package idiomatic

import "reflect"

// Zero returns an empty slice of capacity 0.  The slice zero value.
func Zero[T any]() []T {
	return Empty[T](0)
}

// IsZero returns true if t is equal to the zero value, false otherwise.
func IsZero[T comparable](t []T) bool {
	return reflect.ValueOf(t).IsZero()
}
