package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	tests := []struct {
		input    []string
		match    string
		expected []string
	}{
		{
			input:    []string{"a", "b", "b", "c"},
			match:    "b",
			expected: []string{"b", "b"},
		},
		{
			input:    []string{"a"},
			match:    "b",
			expected: []string{},
		},
		{
			input:    nil,
			match:    "b",
			expected: []string{},
		},
	}

	for _, tc := range tests {
		actual := slicer.Filter(tc.match, tc.input...)

		assert.NotNil(t, actual)
		assert.Len(t, actual, len(tc.expected))
		assert.Equal(t, tc.expected, actual)
	}

}

func Test_FilterFn(t *testing.T) {
	tests := []struct {
		input    []string
		fn       func(string) bool
		expected []string
	}{
		{
			input:    []string{"a", "b", "b", "c"},
			fn:       func(x string) bool { return x == "b" },
			expected: []string{"b", "b"},
		},
		{
			input:    []string{"a"},
			fn:       func(x string) bool { return x == "b" },
			expected: []string{},
		},
		{
			input:    nil,
			fn:       func(x string) bool { return x == "b" },
			expected: []string{},
		},
	}

	for _, tc := range tests {
		actual := slicer.FilterFn(tc.fn, tc.input...)

		assert.NotNil(t, actual)
		assert.Len(t, actual, len(tc.expected))
		assert.Equal(t, tc.expected, actual)
	}
}
