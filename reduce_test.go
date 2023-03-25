package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Reduce(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
		ok       bool
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: "a",
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
		actual, ok := slicer.Min(tc.input...)

		assert.Equal(t, tc.ok, ok)
		assert.Equal(t, tc.expected, actual)
	}
}

func Test_ReduceErr(t *testing.T) {
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
			expected: fake{"c", 1},
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
		actual, ok := slicer.MinBy(by, tc.input...)

		assert.Equal(t, tc.ok, ok)
		assert.Equal(t, tc.expected, actual)
	}
}
