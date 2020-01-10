package treemap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	m := New()

	assert.Equal(t, 0, m.Size())
	assert.False(t, m.Contains(5))
	assert.Equal(t, nil, m.Get(3))

	for i := 9; i >= 0; i-- {
		m.Insert(i, i+1000)
	}

	assert.Equal(t, 10, m.Size())
	assert.True(t, m.Contains(5))
	assert.Equal(t, 3+1000, m.Get(3))
	m.Erase(3)
	assert.Equal(t, nil, m.Get(3))
	m.Clear()
	assert.False(t, m.Contains(50))
	assert.Equal(t, 0, m.Size())
}

func TestMapIterator(t *testing.T) {
	m := New(WithGoroutineSafe())

	for i := 1; i <= 10; i++ {
		m.Insert(i, i)
	}

	i := 1
	for iter := m.First(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value().(int))
		i++
	}

	i = 10
	for iter := m.Last().Clone().(*MapIterator); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value().(int))
		i--
	}

	assert.True(t, m.Begin().Equal(m.First()))

	iter := m.Find(8)
	assert.Equal(t, 8, iter.Key().(int))
	assert.Equal(t, 8, iter.Value().(int))

	iter = m.LowerBound(8)
	assert.Equal(t, 8, iter.Value().(int))

	m.EraseIter(iter)
	assert.False(t, m.Contains(8))

}

func TestMapIteratorSetValue(t *testing.T) {
	m := New(WithGoroutineSafe())
	m.Insert(1, "aaa")
	m.Insert(2, "bbb")
	m.Insert(3, "hhh")

	assert.Equal(t, "aaa", m.Get(1))

	iter := m.Find(1)

	assert.Equal(t, "aaa", iter.Value())

	iter.SetValue("ccc")
	assert.Equal(t, "ccc", m.Get(1))
}

func TestMap_Traversal(t *testing.T) {
	m := New()

	for i := 1; i <= 5; i++ {
		m.Insert(i, i+1000)
	}

	i := 1
	m.Traversal(func(key, value interface{}) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i+1000, value)
		i++
		return true
	})
}
