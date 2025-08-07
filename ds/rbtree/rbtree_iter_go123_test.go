//go:build go1.23

package rbtree

import (
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+100)
	}

	i := 0
	for k, v := range tree.Iter2() {
		assert.Equal(t, i, k)
		assert.Equal(t, i+100, v)
		i++
	}

	for _, vPtr := range tree.IterMut2() {
		*vPtr *= 2
	}
	i = 0
	for k, v := range tree.Iter2() {
		assert.Equal(t, i, k)
		assert.Equal(t, (i+100)*2, v)
		i++
	}
}
