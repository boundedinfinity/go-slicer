package slicer_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/go-slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Fold(t *testing.T) {
	tests := []struct {
		input    []int
		initial  int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			initial:  0,
			expected: 6,
		},
		{
			input:    []int{1},
			initial:  0,
			expected: 1,
		},
		{
			input:    make([]int, 0),
			initial:  0,
			expected: 0,
		},
		{
			input:    nil,
			initial:  0,
			expected: 0,
		},
	}

	fn := func(a, b int) int {
		return a + b
	}

	for i, tc := range tests {
		actual := slicer.Fold(tc.initial, fn, tc.input...)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("test[%v]", i))
	}
}

func Test_FoldErr(t *testing.T) {
	fn1 := func(a, b int) (int, error) {
		return a + b, nil
	}

	ferr := fmt.Errorf("fold err")

	fn2 := func(a, b int) (int, error) {
		if b == 2 {
			return a, ferr
		}

		return a + b, nil
	}

	tests := []struct {
		input    []int
		fn       func(a, b int) (int, error)
		initial  int
		expected int
		err      error
	}{
		{
			input:    []int{1, 2, 3},
			initial:  0,
			expected: 6,
			fn:       fn1,
			err:      nil,
		},
		{
			input:    []int{1},
			initial:  0,
			expected: 1,
			fn:       fn1,
			err:      nil,
		},
		{
			input:    make([]int, 0),
			initial:  0,
			expected: 0,
			fn:       fn1,
			err:      nil,
		},
		{
			input:    nil,
			initial:  0,
			expected: 0,
			fn:       fn1,
			err:      nil,
		},
		{
			input:    []int{1, 2, 3},
			initial:  0,
			expected: 1,
			fn:       fn2,
			err:      ferr,
		},
	}

	for i, tc := range tests {
		actual, err := slicer.FoldErr(tc.initial, tc.fn, tc.input...)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("test[%v]", i))
	}
}
