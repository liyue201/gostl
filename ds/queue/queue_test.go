package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New(WithThreadSafe())
	for i := 0; i < 10; i++ {
		q.Push(i)
		if q.Back() != i {
			t.Fatalf("expect %v, but get %v", i, q.Back())
		}
	}
	t.Logf("%v", q.String())
	if q.Size() != 10 {
		t.Fatalf("size error, expect %v, but get %v", 10, q.Size())
	}
	i := 0
	for !q.Empty() {
		if i != q.Front() {
			t.Fatalf(" expect %v, but get %v", i, q.Front())
		}
		k := q.Pop()
		if k != i {
			t.Fatalf(" expect %v, but get %v", i, k)
		}
		i++
	}
}

func TestQueueWithListContainer(t *testing.T) {
	q := New(WithListContainer())
	for i := 0; i < 10; i++ {
		q.Push(i)
		if q.Back() != i {
			t.Fatalf("expect %v, but get %v", i, q.Back())
		}
	}
	t.Logf("%v", q.String())
	if q.Size() != 10 {
		t.Fatalf("size error, expect %v, but get %v", 10, q.Size())
	}
	i := 0
	for !q.Empty() {
		if i != q.Front() {
			t.Fatalf(" expect %v, but get %v", i, q.Front())
		}
		k := q.Pop()
		if k != i {
			t.Fatalf(" expect %v, but get %v", i, k)
		}
		i++
	}
	q.Push(10)
	q.Clear()
	if q.Size() != 0 {
		t.Fatalf("size error, expect %v, but get %v", 0, q.Size())
	}
}
