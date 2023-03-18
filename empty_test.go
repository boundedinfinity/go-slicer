package slicer_test

import (
	"testing"

	"github.com/boundedinfinity/go-slicer"
	"github.com/stretchr/testify/assert"
)

func Test_Empty(t *testing.T) {
	actual := slicer.Empty[string](5)

	assert.NotNil(t, actual)
	assert.IsType(t, []string{}, actual)
	assert.Equal(t, 5, cap(actual))
}

func Test_Zero(t *testing.T) {
	actual := slicer.Zero[string]()

	assert.NotNil(t, actual)
	assert.IsType(t, []string{}, actual)
	assert.Equal(t, 0, cap(actual))
}
