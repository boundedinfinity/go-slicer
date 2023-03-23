package slicer_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/go-slicer"
	"github.com/stretchr/testify/assert"
)

type joinFaker struct {
	i int
}

func (t joinFaker) String() string {
	return fmt.Sprintf("%v", t.i)
}

func Test_Join(t *testing.T) {
	tests := []struct {
		input    []fmt.Stringer
		sep      string
		expected string
	}{
		{
			input:    []fmt.Stringer{joinFaker{1}, joinFaker{2}},
			sep:      ",",
			expected: "1,2",
		},
		{
			input:    []fmt.Stringer{joinFaker{1}},
			sep:      ",",
			expected: "1",
		},
		{
			input:    nil,
			sep:      ",",
			expected: "",
		},
	}

	for _, tc := range tests {
		actual := slicer.Join(tc.sep, tc.input...)

		assert.NotNil(t, actual)
		assert.Equal(t, tc.expected, actual)
	}
}

func Test_JoinFn(t *testing.T) {
	tests := []struct {
		input    []int
		sep      string
		expected string
	}{
		{
			input:    []int{1, 2},
			sep:      ",",
			expected: "1,2",
		},
		{
			input:    []int{1},
			sep:      ",",
			expected: "1",
		},
		{
			input:    nil,
			sep:      ",",
			expected: "",
		},
	}

	for _, tc := range tests {
		fn := func(i int) fmt.Stringer {
			return joinFaker{i}
		}

		actual := slicer.JoinFn(tc.sep, fn, tc.input...)

		assert.NotNil(t, actual)
		assert.Equal(t, tc.expected, actual)
	}
}
