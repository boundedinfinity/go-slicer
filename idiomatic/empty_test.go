package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func Test_Empty(t *testing.T) {
	actual := slicer.Empty[string](5)

	assert.NotNil(t, actual)
	assert.IsType(t, []string{}, actual)
	assert.Equal(t, 5, cap(actual))
}
