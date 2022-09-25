package rbtree

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestRbTeeFind(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+10000)
	}
	assert.False(t, tree.Empty())
	assert.Equal(t, 10, tree.Size())

	for i := 0; i < 10; i++ {
		val, _ := tree.Find(i)
		assert.Equal(t, i+10000, val)
	}
	for i := 0; i < 10; i++ {
		iter := tree.FindLowerBoundNode(i)
		assert.Equal(t, i+10000, iter.Value())

		iter2 := tree.FindUpperBoundNode(i - 1)
		assert.Equal(t, i+10000, iter2.Value())
	}
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+20000)
	}

	for i := 0; i < 10; i++ {
		iter := tree.FindLowerBoundNode(i)
		count := 0
		for n := iter; n != nil; n = n.Next() {
			if n.key != i {
				break
			}
			count++
			//t.Logf("travesal: %v = %v ", n.key, n.value)
		}
		assert.Equal(t, 2, count)
	}

	for i := 0; i < 10; i++ {
		iter := tree.FindUpperBoundNode(i - 1)
		count := 0
		for n := iter; n != nil; n = n.Next() {
			if n.key != i {
				break
			}
			count++
			//t.Logf("travesal: %v = %v ", n.key, n.value)
		}
		assert.Equal(t, 2, count)
	}
}

func TestRbTeeDelete(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		tree.Insert(i, i)
		m[i] = i
	}
	count := 1000
	for k, v := range m {
		//t.Logf("%v", k)
		node := tree.FindNode(k)
		assert.Equal(t, v, node.Value())
		tree.Delete(node)
		assert.Nil(t, tree.FindNode(k))
		count--
		assert.Equal(t, count, tree.Size())
	}
}

func TestTraversal(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+100)
	}
	i := 0
	tree.Traversal(func(key, value int) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i+100, value)
		i++
		return true
	})
}

func TestInsertDelete(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	m := make(map[int]int)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10000; i++ {
		key := rand.Int() % 1000
		val := rand.Int()
		if v, ok := m[key]; ok {
			n := tree.findNode(key)
			assert.Equal(t, v, n.Value())
			delete(m, key)
			tree.Delete(n)
		} else {
			n := tree.findNode(key)
			assert.Nil(t, n)

			m[key] = val
			tree.Insert(key, val)
		}
		assert.Equal(t, len(m), tree.Size())
		b, _ := tree.IsRbTree()
		assert.True(t, b)
	}
	tree.Clear()
	assert.Equal(t, 0, tree.Size())
}

func TestIterator(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+100)
	}

	i := 0
	for iter := tree.IterFirst().Clone().(*RbTreeIterator[int, int]); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Key())
		assert.Equal(t, i+100, iter.Value())
		i++
	}

	i = 9
	for iter := tree.IterLast(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Key())
		assert.Equal(t, i+100, iter.Value())
		iter.SetValue(i * 2)
		i--
	}

	i = 0
	for iter := tree.IterFirst(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Key())
		assert.Equal(t, i*2, iter.Value())
		i++
	}
	assert.True(t, tree.IterFirst().Equal(tree.IterFirst().Clone()))
	assert.False(t, tree.IterFirst().Equal(nil))
	assert.False(t, tree.IterFirst().Equal(tree.IterLast()))
}

func TestNode(t *testing.T) {
	tree := New[int, int](comparator.IntComparator)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+100)
	}

	i := 0
	for n := tree.Begin(); n != nil; n = n.Next() {
		assert.Equal(t, i, n.Key())
		assert.Equal(t, i+100, n.Value())
		i++
	}

	i = 9
	for n := tree.RBegin(); n != nil; n = n.Prev() {
		assert.Equal(t, i, n.Key())
		assert.Equal(t, i+100, n.Value())
		n.SetValue(i * 2)
		i--
	}

	i = 0
	for n := tree.Begin(); n != nil; n = n.Next() {
		assert.Equal(t, i, n.Key())
		assert.Equal(t, i*2, n.Value())
		i++
	}
}
