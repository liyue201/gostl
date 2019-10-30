package deque

import "testing"

func TestPushPop(t *testing.T) {
	q := New()

	q.PushBack(1)  //[1]
	q.PushFront(2) //[2 1]
	q.PushBack(3)  //[2 1 3] 

	t.Logf("q: %v, %v", q, q.data)

	if q.Size() != 3 {
		t.Fatalf("size error: %v", q.Size())
	}
	if q.Front() != 2 {
		t.Fatalf("front error: %v", q.Front())
	}
	if q.Back() != 3 {
		t.Fatalf("back error: %v", q.Front())
	}
	if q.At(1) != 1 {
		t.Fatalf("at 1 error: %v", q.At(1))
	}
	q.Insert(0, 5) //[5 2 1 3]  
	q.Insert(3, 6) //[5 2 1 6 3]
	q.Insert(2, 7) //[5 2 7 1 6 3]  
	t.Logf("q: %v, %v", q, q.data)

	if q.String() != "[5 2 7 1 6 3]" {
		t.Fatalf("Insert error: %v", q.String())
	}

	val := q.PopBack()
	if val != 3 || q.String() != "[5 2 7 1 6]" {
		t.Fatalf("PopBack error: %v %v", val, q.String())
	}
	val = q.PopFront()
	if val != 5 || q.String() != "[2 7 1 6]" {
		t.Fatalf("PopBack error: %v %v", val, q.String())
	}
	t.Logf("q: %v, %v", q, q.data)
}

func TestErase(t *testing.T) {
	q := New()
	for i := 0; i < 5; i++ {
		q.PushBack(i + 1)
	}
   //[1 2 3 4 5]

	t.Logf("capacity: %v", q.Capacity())
	t.Logf("q: %v, %v", q, q.data)

	q.Erase(1) //[1 3 4 5]
	if q.String() != "[1 3 4 5]" {
		t.Fatalf("Erase pos=1 error: %v", q.String())
	}
	t.Logf("q: %v, %v", q, q.data)

	q.Erase(0) //[3 4 5]
	t.Logf("q: %v, %v", q, q.data)
	if q.String() != "[3 4 5]" {
		t.Fatalf("Erase  error: %v", q.String())
	}

	q.PushFront(6)
	q.PushBack(7)
	q.PushFront(8)
	t.Logf("q: %v, %v", q, q.data)
	if q.String() != "[8 6 3 4 5 7]" {
		t.Fatalf("Push error: %v", q.String())
	}
	q.EraseRange(3, 5)
	t.Logf("q: %v, %v", q, q.data)
	if q.String() != "[8 6 3 7]" {
		t.Fatalf("Push error: %v", q.String())
	}
}
