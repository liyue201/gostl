package treemap

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	m := New[int, int](comparator.IntComparator)

	assert.Equal(t, 0, m.Size())
	assert.False(t, m.Contains(5))
	_, err := m.Get(3)
	assert.Equal(t, err, ErrorNotFound)

	for i := 9; i >= 0; i-- {
		m.Insert(i, i+1000)
	}

	assert.Equal(t, 10, m.Size())
	assert.True(t, m.Contains(5))
	v, _ := m.Get(3)
	assert.Equal(t, 3+1000, v)
	m.Erase(3)
	_, err = m.Get(3)
	assert.Equal(t, err, ErrorNotFound)
	m.Clear()
	assert.False(t, m.Contains(50))
	assert.Equal(t, 0, m.Size())
}

func TestMapIterator(t *testing.T) {
	m := New[int, int](comparator.IntComparator, WithGoroutineSafe())

	for i := 1; i <= 10; i++ {
		m.Insert(i, i)
	}

	i := 1
	for iter := m.First(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		i++
	}

	i = 10
	for iter := m.Last().Clone().(*MapIterator[int, int]); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value())
		i--
	}

	assert.True(t, m.Begin().Equal(m.First()))

	iter := m.Find(8)
	assert.Equal(t, 8, iter.Key())
	assert.Equal(t, 8, iter.Value())

	iter = m.LowerBound(8)
	assert.Equal(t, 8, iter.Value())

	iter = m.UpperBound(6)
	assert.Equal(t, 7, iter.Value())

	m.EraseIter(iter)
	assert.False(t, m.Contains(7))
}

func TestMapIteratorSetValue(t *testing.T) {
	m := New[int, string](comparator.IntComparator, WithGoroutineSafe())
	m.Insert(1, "aaa")
	m.Insert(2, "bbb")
	m.Insert(3, "hhh")

	v, _ := m.Get(3)
	assert.Equal(t, "hhh", v)

	iter := m.Find(1)
	assert.Equal(t, "aaa", iter.Value())

	iter.SetValue("ccc")
	v, _ = m.Get(1)
	assert.Equal(t, "ccc", v)
}

func TestMap_Traversal(t *testing.T) {
	m := New[int, int](comparator.IntComparator)

	for i := 1; i <= 5; i++ {
		m.Insert(i, i+1000)
	}

	i := 1
	m.Traversal(func(key, value int) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i+1000, value)
		i++
		return true
	})
}
