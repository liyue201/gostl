package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New(WithThreadSave())
	for i := 0; i < 10; i++ {
		s.Push(i)
		if s.Top() != i {
			t.Fatalf("expect %v, but get %v", i, s.Top())
		}
	}
	t.Logf("%v", s.String())
	if s.Size() != 10 {
		t.Fatalf("expect %v, but get %v", 10, s.Size())
	}
	i := 9
	for !s.Empty() {
		k := s.Pop()
		if k != i {
			t.Fatalf("expect %v, but get %v", i, k)
		}
		i--
	}
}

func TestStackWithListContainer(t *testing.T) {
	s := New(WithListContainer())
	for i := 0; i < 10; i++ {
		s.Push(i)
		if s.Top() != i {
			t.Fatalf("expect %v, but get %v", i, s.Top())
		}
	}
	if s.Size() != 10 {
		t.Fatalf("expect %v, but get %v", 10, s.Size())
	}
	i := 9
	for !s.Empty() {
		k := s.Pop()
		if k != i {
			t.Fatalf("expect %v, but get %v", i, k)
		}
		i--
	}
	s.Push(10)
	s.Clear()
	if !s.Empty() {
		t.Fatalf("expect true, but get false")
	}
}
