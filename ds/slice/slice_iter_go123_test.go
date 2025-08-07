//go:build go1.23

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIter(t *testing.T) {
	a := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	sliceA := NewSliceWrapper(a)

	l := make([]int, 0)
	for v := range sliceA.Iter() {
		l = append(l, v)
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, l)

	// mut
	for vPtr := range sliceA.IterMut() {
		*vPtr *= 2
	}

	l = make([]int, 0)
	for v := range sliceA.Iter() {
		l = append(l, v)
	}
	assert.Equal(t, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, l)
}
