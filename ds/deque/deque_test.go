package deque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushPop(t *testing.T) {
	q := New()

	q.PushBack(1)  //[1]
	q.PushFront(2) //[2 1]
	q.PushBack(3)  //[2 1 3]

	t.Logf("q: %v", q)

	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 2, q.Front())
	assert.Equal(t, 3, q.Back())
	assert.Equal(t, 1, q.At(1))

	q.Insert(0, 5) //[5 2 1 3]
	q.Insert(3, 6) //[5 2 1 6 3]
	q.Insert(2, 7) //[5 2 7 1 6 3]
	t.Logf("q: %v", q)

	assert.Equal(t, "[5 2 7 1 6 3]", q.String())

	assert.Equal(t, 3, q.PopBack())
	assert.Equal(t, "[5 2 7 1 6]", q.String())

	assert.Equal(t, 5, q.PopFront())
	assert.Equal(t, "[2 7 1 6]", q.String())

	t.Logf("q: %v", q)
}

func TestErase(t *testing.T) {
	q := New()
	assert.True(t, q.Empty())
	for i := 0; i < 5; i++ {
		q.PushBack(i + 1)
	}
	assert.False(t, q.Empty())

	//[1 2 3 4 5]
	t.Logf("q: %v", q)
	q.EraseAt(1) //[1 3 4 5]
	assert.Equal(t, "[1 3 4 5]", q.String())

	t.Logf("q: %v", q)
	q.EraseAt(0) //[3 4 5]

	t.Logf("q: %v", q)
	assert.Equal(t, "[3 4 5]", q.String())

	q.PushFront(6)
	q.PushBack(7)
	q.PushFront(8)
	t.Logf("q: %v", q)

	assert.Equal(t, "[8 6 3 4 5 7]", q.String())

	q.EraseRange(3, 5)
	t.Logf("q: %v", q)
	assert.Equal(t, "[8 6 3 7]", q.String())

	q.Clear()
	assert.True(t, q.Empty())
}

func TestIterator(t *testing.T) {
	q := New()
	for i := 0; i < 10; i++ {
		q.PushBack(i)
	}
	n := 0
	for iter := q.Begin(); !iter.Equal(q.End()); iter.Next() {
		assert.Equal(t, n, iter.Value())
		n++
	}

	n = 9
	for iter := q.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, n, iter.Value())
		n--
	}

	iter := q.First().IteratorAt(5).Clone().(*DequeIterator)
	assert.Equal(t, 5, iter.Position())
	assert.Equal(t, 5, iter.Value())

	iter.SetValue(555)
	assert.Equal(t, 555, iter.Value())
}
