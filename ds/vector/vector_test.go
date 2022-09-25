package vector

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVectorBase(t *testing.T) {
	v := New[int](WithCapacity(10))
	assert.True(t, v.Empty())
	assert.Equal(t, 10, v.Capacity())
	v.PushBack(1)
	v.PushBack(2)

	assert.False(t, v.Empty())
	assert.Equal(t, 2, v.Size())

	assert.Equal(t, 1, v.Front())
	assert.Equal(t, 2, v.Back())
}

func TestVectorResize(t *testing.T) {
	v := New[int](WithCapacity(10))
	v.PushBack(1)
	v.PushBack(2)
	v.ShrinkToFit()
	assert.Equal(t, 2, v.Capacity())

	assert.Equal(t, 1, v.At(0))
	assert.Equal(t, 1, v.Front())

	v.Reserve(20)

	assert.Equal(t, 20, v.Capacity())
	assert.Equal(t, 2, v.Size())
	assert.Equal(t, 2, v.At(1))
	v.Clear()
	assert.Equal(t, 0, v.Size())
	assert.True(t, v.Empty())

	for i := 0; i < 10; i++ {
		v.PushBack(i)
	}
	assert.Equal(t, 10, v.Size())
	v.Resize(20)
	assert.Equal(t, 10, v.Size())
	v.Resize(4)
	assert.Equal(t, 4, v.Size())

	b := NewFromVector(v)
	assert.Equal(t, 4, b.Size())
	assert.Equal(t, "[0 1 2 3]", fmt.Sprintf("%v", b.Data()))
}

func TestModifyVector(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	//[1,2,3]
	assert.Equal(t, 3, v.PopBack())
	//[1 2]
	v.PushBack(4)
	//[1 2 4]

	v.SetAt(1, 9)
	assert.Equal(t, 9, v.At(1))
	//[1 9 4]

	v.InsertAt(0, 8)
	//[8 1 9 4]

	assert.Equal(t, "[8 1 9 4]", v.String())
	v.Clear()
}

func TestVectorIter(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)
	//[1 2 3 4]

	i := 0
	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i+1, v.At(i))
		i++
	}

	i = 3
	for iter := v.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i+1, v.At(i))
		i--
	}
	iter := v.Erase(v.Begin())
	t.Logf("v: %v", v.String())

	assert.Equal(t, 2, iter.Value())
	//[2 3 4]

	v.PushBack(5)
	v.PushBack(6)
	//[2 3 4 5 6]

	iter = v.EraseRange(v.Begin().Next(), v.Begin().Next().Next().Next())
	//[2 5 6]

	assert.Equal(t, 5, iter.Value())
	assert.Equal(t, "[2 5 6]", v.String())

	iter = v.Begin()
	iter = v.Insert(iter, 7)
	//[7 2 5 6]

	assert.Equal(t, 7, iter.Value())
	assert.Equal(t, "[7 2 5 6]", v.String())

	assert.True(t, v.Begin().Equal(v.Begin().Clone()))
	assert.False(t, v.Begin().Equal(v.Last()))
}

func TestSort(t *testing.T) {
	v := New[int]()
	for i := 10; i >= 0; i-- {
		v.PushBack(i)
	}
	sort.Sort[int](v.Begin(), v.End(), comparator.IntComparator)
	for i := 0; i < v.Size(); i++ {
		assert.Equal(t, i, v.At(i))
	}

	sort.Sort[int](v.Begin(), v.End(), comparator.Reverse(comparator.IntComparator))
	for i := 0; i < v.Size(); i++ {
		assert.Equal(t, v.Size()-i-1, v.At(i))
	}
}
