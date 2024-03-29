package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Zero(t *testing.T) {
	actual := slicer.Zero[string]()

	assert.NotNil(t, actual)
	assert.IsType(t, []string{}, actual)
	assert.Equal(t, 0, cap(actual))
}

func Test_IsZero(t *testing.T) {
	actual := slicer.Zero[string]()

	assert.NotNil(t, actual)
	assert.IsType(t, []string{}, actual)
	assert.Equal(t, 0, cap(actual))
}
