package deque

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
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

func TestRandom(t *testing.T) {
	q := New()
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 0)
	for i := 0; i < 11000; i++ {
		k := rand.Int()
		if k%6 == 0 {
			q.PushBack(i)
			a = append(a, i)
		} else if k%6 == 1 {
			q.PushFront(i)
			a = append([]int{i}, a...)
		} else if k%6 == 2 {
			q.PopFront()
			if len(a) > 0 {
				a = a[1:]
			}
		} else if k%6 == 3 {
			q.PopBack()
			if len(a) > 0 {
				a = a[:len(a)-1]
			}
		} else if k%6 == 4 {
			if q.Size() == 0 {
				continue
			}
			s1 := fmt.Sprintf("%v", a)
			s2 := q.String()

			k = k % q.Size()
			b := make([]int, 0)
			b = append(b, a[k:]...)
			a = a[:k]
			a = append(a, i)
			a = append(a, b...)
			q.Insert(k, i)
			if !assert.Equal(t, fmt.Sprintf("%v", a), q.String()) {
				fmt.Printf("%v %v\n", s1, s2)
				fmt.Printf("%v %v\n", k, i)
				return
			}
		} else {
			if q.Size() == 0 {
				continue
			}
			k = k % q.Size()
			q.EraseAt(k)
			a = append(a[:k], a[k+1:]...)
			if !assert.Equal(t, fmt.Sprintf("%v", a), q.String()) {
				return
			}
		}
		if !assert.Equal(t, fmt.Sprintf("%v", a), q.String()) {
			return
		}
		//t.Logf("%v", q.String())
		//t.Logf("%v", a)
	}
}

func TestInsert1(t *testing.T) {
	q := New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4]", q.String())

	q.Insert(0, 5)
	assert.Equal(t, "[5 0 1 2 3 4]", q.String())

	q = New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4]", q.String())

	q.Insert(1, 5)
	assert.Equal(t, "[0 5 1 2 3 4]", q.String())

	q = New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4]", q.String())

	q.Insert(2, 5)
	assert.Equal(t, "[0 1 5 2 3 4]", q.String())

	q = New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	q.Insert(3, 5)
	assert.Equal(t, "[0 1 2 5 3 4]", q.String())

	q = New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	q.Insert(4, 5)
	assert.Equal(t, "[0 1 2 3 5 4]", q.String())

	q = New()
	for i := 0; i < 6; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4 5]", q.String())

	q.Insert(5, 6)
	assert.Equal(t, "[0 1 2 3 4 6 5]", q.String())
}

func TestInsert2(t *testing.T) {
	q := New()
	for i := 0; i < 4; i++ {
		q.PushBack(i)
	}
	q.PushFront(4) //[4 | 0 1 2 3]
	assert.Equal(t, "[4 0 1 2 3]", q.String())

	q.Insert(0, 5)
	assert.Equal(t, "[5 4 0 1 2 3]", q.String())

	q = New()
	for i := 0; i < 4; i++ {
		q.PushBack(i)
	}
	q.PushFront(4) //[4 | 0 1 2 3]
	q.Insert(1, 5)
	assert.Equal(t, "[4 5 0 1 2 3]", q.String())

	q = New()
	for i := 0; i < 4; i++ {
		q.PushBack(i)
	}
	q.PushFront(4) //[4 | 0 1 2 3]
	q.Insert(2, 5)
	assert.Equal(t, "[4 0 5 1 2 3]", q.String())

	q = New()
	for i := 0; i < 4; i++ {
		q.PushBack(i)
	}
	q.PushFront(4) //[4 | 0 1 2 3]
	q.Insert(3, 5)
	assert.Equal(t, "[4 0 1 5 2 3]", q.String())
}
