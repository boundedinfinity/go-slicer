package idiomatic_test

import (
	"testing"

	slicer "github.com/boundedinfinity/go-slicer/idiomatic"
	"github.com/stretchr/testify/assert"
)

func allOk(s string) bool {
	return s == "A"
}

func allOkI(_ int, s string) bool {
	return s == "A"
}

func Test_All(t *testing.T) {
	assert.False(t, slicer.All(allOk))
	assert.True(t, slicer.All(allOk, "A", "A"))
	assert.True(t, slicer.All(allOk, "A"))
	assert.False(t, slicer.All(allOk, "A", "B"))
}

func Test_AllI(t *testing.T) {
	assert.False(t, slicer.AllI(allOkI))
	assert.True(t, slicer.AllI(allOkI, "A", "A"))
	assert.True(t, slicer.AllI(allOkI, "A"))
	assert.False(t, slicer.AllI(allOkI, "A", "B"))
}
