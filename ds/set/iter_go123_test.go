//go:build go1.23

package set

import (
	"sort"
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
)

func TestSetIter(t *testing.T) {
	s := New(comparator.IntComparator, WithGoroutineSafe())
	for i := 0; i < 10; i++ {
		s.Insert(i)
	}

	l := make([]int, 0)
	for v := range s.Iter() {
		l = append(l, v)
	}
	sort.Ints(l)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, l)
}

func TestMultiSetIter(t *testing.T) {
	mset := NewMultiSet(comparator.IntComparator, WithGoroutineSafe())

	mset.Insert(1)
	mset.Insert(5)
	mset.Insert(1)

	l := make([]int, 0)
	for v := range mset.Iter() {
		l = append(l, v)
	}
	sort.Ints(l)
	assert.Equal(t, []int{1, 1, 5}, l)
}
