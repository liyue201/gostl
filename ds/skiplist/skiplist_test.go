package skiplist

import (
	"github.com/bmizerany/assert"
	"github.com/liyue201/gostl/utils/comparator"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	list := New(WithMaxLevel(5))

	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := rand.Int() % 100
		list.Insert(key, i)
		m[key] = i
	}
	for key, v := range m {
		ret := list.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), list.Len())
}

func TestRemove(t *testing.T) {
	list := New(WithGoroutineSafe(), WithKeyComparator(comparator.IntComparator))

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		key := rand.Int() % 1000
		list.Insert(key, i)
		m[key] = i
	}
	assert.Equal(t, len(m), list.Len())

	for i := 0; i < 300; i++ {
		key := rand.Int() % 1000
		list.Remove(key)
		delete(m, key)
		key2 := rand.Int() % 10440
		list.Insert(key2, key)
		m[key2] = key
	}

	for key, v := range m {
		ret := list.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), list.Len())
}

func TestSkiplist_Traversal(t *testing.T) {
	list := New()
	for i := 0; i < 10; i++ {
		list.Insert(i, i*10)
	}
	keys := list.Keys()
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, keys[i])
	}
	i := 0
	list.Traversal(func(key, value interface{}) bool {
		assert.Equal(t, i, key.(int))
		assert.Equal(t, i*10, value.(int))
		i++
		return true
	})
}
