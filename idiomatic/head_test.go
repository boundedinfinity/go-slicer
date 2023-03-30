package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Head(t *testing.T) {
	tests := []struct {
		input         []string
		expectedValue string
		expectedOk    bool
	}{
		{
			input:         []string{"a", "b", "c"},
			expectedValue: "a",
			expectedOk:    true,
		},
		{
			input:         []string{"a"},
			expectedValue: "a",
			expectedOk:    true,
		},
		{
			input:         make([]string, 0),
			expectedValue: "",
			expectedOk:    false,
		},
		{
			input:         nil,
			expectedValue: "",
			expectedOk:    false,
		},
	}

	for _, tc := range tests {
		actualValue, actualOk := slicer.Head(tc.input...)

		assert.Equal(t, tc.expectedOk, actualOk)
		assert.Equal(t, tc.expectedValue, actualValue)
	}

}
