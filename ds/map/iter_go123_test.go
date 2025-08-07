//go:build go1.23

package treemap

import (
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
)

func TestMapIter(t *testing.T) {
	m := New[int, int](comparator.IntComparator, WithGoroutineSafe())

	for i := 1; i <= 10; i++ {
		m.Insert(i, i)
	}

	i := 1
	for k, v := range m.Iter2() {
		assert.Equal(t, i, k)
		assert.Equal(t, i, v)
		i++
	}
}

func TestMultiMapIter(t *testing.T) {
	m := NewMultiMap[int, int](comparator.IntComparator, WithGoroutineSafe())

	for i := 1; i <= 10; i++ {
		m.Insert(i, i)
	}

	i := 1
	for k, v := range m.Iter2() {
		assert.Equal(t, i, k)
		assert.Equal(t, i, v)
		i++
	}
}
