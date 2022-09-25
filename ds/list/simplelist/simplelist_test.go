package simplelist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := New[int]()
	assert.Equal(t, 0, list.Len())
	list.PushBack(1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.FrontNode().Value)
	assert.Equal(t, 1, list.BackNode().Value)
	list.PushFront(2)

	assert.Equal(t, 2, list.Len())
	assert.Equal(t, "[2 1]", list.String())

	list.PushBack(3)
	list.PushBack(4)
	assert.Equal(t, "[2 1 3 4]", list.String())

	list.MoveToFront(list.FrontNode(), list.FrontNode().Next())

	assert.Equal(t, "[1 2 3 4]", list.String())

	list.MoveToBack(list.FrontNode(), list.FrontNode().Next())

	assert.Equal(t, "[1 3 4 2]", list.String())

	ret := make([]int, 0)
	list.Traversal(func(value int) bool {
		ret = append(ret, value)
		return true
	})
	assert.Equal(t, "[1 3 4 2]", fmt.Sprintf("%v", ret))
}

func TestList_InsertAfter(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	list.InsertAfter(6, list.FrontNode())
	assert.Equal(t, "[1 6 2 3 4 5]", list.String())

	list.InsertAfter(7, list.FrontNode().Next())
	assert.Equal(t, "[1 6 7 2 3 4 5]", list.String())

	list.InsertAfter(8, list.BackNode())
	assert.Equal(t, "[1 6 7 2 3 4 5 8]", list.String())
}

func TestList_Remove(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	list.Remove(nil, list.FrontNode())
	assert.Equal(t, "[2 3 4 5]", list.String())

	list.Remove(list.FrontNode(), list.FrontNode().Next())
	assert.Equal(t, "[2 4 5]", list.String())

	list.PushFront(6)
	assert.Equal(t, "[6 2 4 5]", list.String())
}

func TestListIterator(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	i := 1
	for iter := NewIterator(list.FrontNode()); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		iter.SetValue(i * 2)
		i++
	}
	iter := NewIterator(list.FrontNode())
	assert.Equal(t, 2, iter.Value())
	assert.True(t, iter.Equal(iter.Clone()))
}
