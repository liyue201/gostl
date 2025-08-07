//go:build go1.23

package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	n := 0
	for v := range q.Iter() {
		assert.Equal(t, n, v)
		n++
	}

	for v := range q.IterMut() {
		*v *= 2
	}

	n = 0
	for v := range q.Iter() {
		assert.Equal(t, n*2, v)
		n++
	}
}
