package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-slicer"
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

func Test_Last(t *testing.T) {
	tests := []struct {
		input         []string
		expectedValue string
		expectedOk    bool
	}{
		{
			input:         []string{"a", "b", "c"},
			expectedValue: "c",
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
		actualValue, actualOk := slicer.Last(tc.input...)

		assert.Equal(t, tc.expectedOk, actualOk)
		assert.Equal(t, tc.expectedValue, actualValue)
	}
}

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
