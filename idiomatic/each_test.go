package idiomatic_test

import (
	"strings"
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Each(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b"},
			expected: []string{"A", "B"},
		},
		{
			input:    []string{"a"},
			expected: []string{"A"},
		},
		{
			input:    []string{},
			expected: make([]string, 0),
		},
		{
			input:    nil,
			expected: make([]string, 0),
		},
	}

	for _, test := range tests {
		actual := make([]string, 0)

		fn := func(s string) {
			actual = append(actual, strings.ToUpper(s))
		}

		slicer.Each(fn, test.input...)

		assert.Len(t, actual, len(test.expected))
		assert.Equal(t, test.expected, actual)
	}

}

// func Test_EachErr(t *testing.T) {
// 	tests := []struct {
// 		input    []string
// 		expected []string
// 	}{
// 		{
// 			input:    []string{"a", "b"},
// 			expected: []string{"A", "B"},
// 		},
// 		{
// 			input:    []string{"a"},
// 			expected: []string{"A"},
// 		},
// 		{
// 			input:    []string{},
// 			expected: make([]string, 0),
// 		},
// 		{
// 			input:    nil,
// 			expected: make([]string, 0),
// 		},
// 	}

// 	fakeErr := fmt.Errorf("each err")

// 	for _, tc := range tests {
// 		actual := make([]string, 0)

// 		slicer.EachErr(tc.input...)

// 		assert.NotNil(t, actual)
// 		assert.Len(t, actual, len(tc.expected))
// 		assert.Equal(t, tc.expected, actual)
// 	}

// }
