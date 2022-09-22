package bidlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := New[int]()
	assert.True(t, list.Empty())
	list.PushBack(1)

	assert.Equal(t, 1, list.FrontNode().Value)
	assert.Equal(t, 1, list.BackNode().Value)
	list.PushFront(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Len())

	list.PushBack(3)
	list.PushFront(4)

	assert.Equal(t, 4, list.FrontNode().Value)
	assert.Equal(t, 4, list.Front())
	assert.Equal(t, 3, list.BackNode().Value)
	assert.Equal(t, 3, list.Back())

	t.Logf("list: %v", list)
	list.Remove(list.FrontNode())

	t.Logf("list: %v", list)
	assert.Equal(t, "[2 1 3]", list.String())

	list.Remove(list.BackNode())
	assert.Equal(t, "[2 1]", list.String())

	list.PushBack(5)
	list.PushBack(6)
	list.InsertAfter(7, list.FrontNode())
	t.Logf("list: %v", list)
	assert.Equal(t, "[2 7 1 5 6]", list.String())

	list.InsertBefore(8, list.BackNode().Prev())

	assert.Equal(t, "[2 7 1 8 5 6]", list.String())

	list.Remove(list.FrontNode().Next().Next())
	assert.Equal(t, "[2 7 8 5 6]", list.String())

	list.PopFront()
	assert.Equal(t, "[7 8 5 6]", list.String())

	list.PopBack()
	assert.Equal(t, "[7 8 5]", list.String())

	list.PushBack(2)
	list.PushBack(1)
	assert.Equal(t, "[7 8 5 2 1]", list.String())

	list.MoveToBack(list.FrontNode().Next())
	assert.Equal(t, "[7 5 2 1 8]", list.String())

	list.MoveToFront(list.FrontNode().Next())
	assert.Equal(t, "[5 7 2 1 8]", list.String())

	list.MoveToFront(list.BackNode())
	assert.Equal(t, "[8 5 7 2 1]", list.String())

	list.MoveToBack(list.FrontNode())
	assert.Equal(t, "[5 7 2 1 8]", list.String())

	list.MoveToFront(list.FrontNode())
	assert.Equal(t, "[5 7 2 1 8]", list.String())

	list.MoveToBack(list.BackNode())
	assert.Equal(t, "[5 7 2 1 8]", list.String())

	list.moveToAfter(list.FrontNode().Next(), list.BackNode().Prev())
	assert.Equal(t, "[5 2 1 7 8]", list.String())

	ret := make([]int, 0)
	list.Traversal(func(value any) bool {
		ret = append(ret, value.(int))
		if value == 1 {
			return false
		}
		return true
	})
	assert.Equal(t, "[5 2 1]", fmt.Sprintf("%v", ret))

	ret = make([]int, 0)
	list.Traversal(func(value any) bool {
		ret = append(ret, value.(int))
		return true
	})
	assert.Equal(t, "[5 2 1 7 8]", fmt.Sprintf("%v", ret))
}

func TestPushBackList(t *testing.T) {
	list := New[int]()
	list.PushBack(7)
	list.PushBack(8)
	list.PushBack(5)
	list.PushBackList(list)
	t.Logf("list: %v", list)
	assert.Equal(t, "[7 8 5 7 8 5]", list.String())

	list2 := New[int]()
	list2.PushBack(1)
	list2.PushBack(2)
	list2.PushBack(3)
	list2.PushFrontList(list2)
	t.Logf("list: %v", list2)
	assert.Equal(t, "[1 2 3 1 2 3]", list2.String())

	list.PushBackList(list2)
	t.Logf("list: %v", list)
	assert.Equal(t, "[7 8 5 7 8 5 1 2 3 1 2 3]", list.String())
}

func TestListIterator(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	i := 1
	for iter := NewIterator(list.FrontNode()); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		i++
	}

	i = 5
	for iter := NewIterator(list.BackNode()); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value())
		iter.SetValue(i * 2)
		i--
	}
	iter := NewIterator(list.FrontNode())
	assert.Equal(t, 2, iter.Value())
	assert.True(t, iter.Equal(iter.Clone()))
}
