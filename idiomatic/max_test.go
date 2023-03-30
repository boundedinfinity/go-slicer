package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Max(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
		ok       bool
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: "c",
			ok:       true,
		},
		{
			input:    []string{"a"},
			expected: "a",
			ok:       true,
		},
		{
			input:    make([]string, 0),
			expected: "",
			ok:       false,
		},
		{
			input:    nil,
			expected: "",
			ok:       false,
		},
	}

	for _, tc := range tests {
		actual, ok := slicer.Max(tc.input...)

		assert.Equal(t, tc.ok, ok)
		assert.Equal(t, tc.expected, actual)
	}
}

func Test_MaxBy(t *testing.T) {
	type fake struct {
		a string
		i int
	}

	by := func(f fake) int {
		return f.i
	}

	tests := []struct {
		input    []fake
		expected fake
		ok       bool
	}{
		{
			input:    []fake{{"a", 3}, {"b", 2}, {"c", 1}},
			expected: fake{"a", 3},
			ok:       true,
		},
		{
			input:    []fake{{"a", 3}},
			expected: fake{"a", 3},
			ok:       true,
		},
		{
			input:    make([]fake, 0),
			expected: fake{},
			ok:       false,
		},
		{
			input:    nil,
			expected: fake{},
			ok:       false,
		},
	}

	for _, tc := range tests {
		actual, ok := slicer.MaxBy(by, tc.input...)

		assert.Equal(t, tc.ok, ok)
		assert.Equal(t, tc.expected, actual)
	}
}
