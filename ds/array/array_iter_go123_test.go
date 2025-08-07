//go:build go1.23

package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	a := New[int](10)
	n := 0
	for ptr := range a.IterMut() {
		*ptr = n
		n++
	}
	assert.Equal(t, 2, a.At(2))
	assert.Equal(t, 0, a.Front())
	assert.Equal(t, 9, a.Back())

	n = 0
	for val := range a.Iter() {
		assert.Equal(t, n, val)
		n++
	}
}
