package simplelist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := New()
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
	list.Traversal(func(value interface{}) bool {
		ret = append(ret, value.(int))
		return true
	})
	assert.Equal(t, "[1 3 4 2]", fmt.Sprintf("%v", ret))
}

func TestListIterator(t *testing.T) {
	list := New()
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
