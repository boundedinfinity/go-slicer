package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Flatten(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected []string
	}{
		{
			input:    [][]string{{"a"}, {"b"}},
			expected: []string{"a", "b"},
		},
		{
			input:    [][]string{{"a"}},
			expected: []string{"a"},
		},
		{
			input:    nil,
			expected: []string{},
		},
	}

	for _, tc := range tests {
		actual := slicer.Flatten(tc.input...)

		assert.NotNil(t, actual)
		assert.Len(t, actual, len(tc.expected))
		assert.Equal(t, tc.expected, actual)
	}

}
