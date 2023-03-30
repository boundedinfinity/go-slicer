package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Tail(t *testing.T) {
	tests := []struct {
		input         []string
		expectedValue []string
		expectedOk    bool
	}{
		{
			input:         []string{"a", "b", "c"},
			expectedValue: []string{"b", "c"},
			expectedOk:    true,
		},
		{
			input:         []string{"a"},
			expectedValue: make([]string, 0),
			expectedOk:    false,
		},
		{
			input:         make([]string, 0),
			expectedValue: make([]string, 0),
			expectedOk:    false,
		},
		{
			input:         nil,
			expectedValue: make([]string, 0),
			expectedOk:    false,
		},
	}

	for _, tc := range tests {
		actualValue, actualOk := slicer.Tail(tc.input...)

		assert.Equal(t, tc.expectedOk, actualOk)
		assert.Equal(t, tc.expectedValue, actualValue)
	}
}
