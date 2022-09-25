package treemap

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMulitMap(t *testing.T) {
	m := NewMultiMap[int, int](comparator.IntComparator)

	assert.Equal(t, 0, m.Size())
	assert.False(t, m.Contains(5))
	v, err := m.Get(3)
	assert.Equal(t, err, ErrorNotFound)

	for i := 9; i >= 0; i-- {
		m.Insert(i, i+1000)
	}

	assert.Equal(t, 10, m.Size())
	assert.True(t, m.Contains(5))
	v, err = m.Get(3)
	assert.Equal(t, 3+1000, v)
	m.Erase(3)
	v, err = m.Get(3)
	assert.Equal(t, ErrorNotFound, err)
	m.Clear()
	assert.False(t, m.Contains(50))
	assert.Equal(t, 0, m.Size())
}

func TestMultiMapIterator(t *testing.T) {
	m := NewMultiMap[int, int](comparator.IntComparator, WithGoroutineSafe())

	for i := 1; i <= 10; i++ {
		m.Insert(i, i)
	}

	i := 1
	for iter := m.First(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		i++
	}

	i = 10
	for iter := m.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value())
		i--
	}

	assert.True(t, m.Begin().Equal(m.First()))

	iter := m.Find(8)
	assert.Equal(t, 8, iter.Value())

	iter = m.LowerBound(8)
	assert.Equal(t, 8, iter.Value())

	iter = m.UpperBound(5)
	assert.Equal(t, 6, iter.Value())
}

func TestMultiMap_Traversal(t *testing.T) {
	m := NewMultiMap[int, int](comparator.IntComparator)

	for i := 1; i <= 5; i++ {
		m.Insert(i, i+1000)
	}

	for i := 1; i <= 5; i++ {
		m.Insert(i, i+1000)
	}

	i := 1
	count := 0
	m.Traversal(func(key, value int) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i+1000, value)
		count++
		if count%2 == 0 {
			i++
		}
		return true
	})
}
