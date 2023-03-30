package idiomatic_test

import (
	"fmt"
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	tests := []struct {
		input    []int
		expected []string
	}{
		{
			input:    []int{1, 2, 3},
			expected: []string{"1", "2", "3"},
		},
		{
			input:    []int{1},
			expected: []string{"1"},
		},
		{
			input:    make([]int, 0),
			expected: []string{},
		},
		{
			input:    nil,
			expected: []string{},
		},
	}

	fn := func(i int) string {
		return fmt.Sprintf("%v", i)
	}

	for i, test := range tests {
		actual := slicer.Map(fn, test.input...)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("test[%v]", i))
	}
}

func Test_MapErr(t *testing.T) {
	fn1 := func(i int) (string, error) {
		return fmt.Sprintf("%v", i), nil
	}

	fnErr := fmt.Errorf("map err")
	fn2 := func(i int) (string, error) {
		if i == 1 {
			return "", fnErr
		}

		return fmt.Sprintf("%v", i), nil
	}

	tests := []struct {
		input    []int
		expected []string
		fn       func(int) (string, error)
		err      error
	}{
		{
			input:    []int{1, 2, 3},
			expected: []string{"1", "2", "3"},
			fn:       fn1,
			err:      nil,
		},
		{
			input:    []int{1},
			expected: []string{"1"},
			fn:       fn1,
			err:      nil,
		},
		{
			input:    make([]int, 0),
			expected: []string{},
			fn:       fn1,
			err:      nil,
		},
		{
			input:    nil,
			expected: []string{},
			fn:       fn1,
			err:      nil,
		},
		{
			input:    []int{1, 2, 3},
			expected: []string{},
			fn:       fn2,
			err:      fnErr,
		},
		{
			input:    []int{1},
			expected: []string{},
			fn:       fn2,
			err:      fnErr,
		},
	}

	for i, test := range tests {
		actual, err := slicer.MapErr(test.fn, test.input...)
		assert.Equal(t, test.err, err, fmt.Sprintf("%v", i))
		assert.Equal(t, test.expected, actual, fmt.Sprintf("%v", i))
	}
}
