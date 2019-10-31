package set

import (
	"math/rand"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	s := New()
	for i := 10; i >= 1; i-- {
		s.Insert(i)
	}

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		t.Logf("%v\n", iter.Value())
	}
	t.Logf("=======================")

	for iter := s.Last(); iter.IsValid(); iter.Prev() {
		t.Logf("%v\n", iter.Value())
	}
	t.Logf("=======================")

	if !s.Contains(5) {
		t.Logf("Contains(10) error\n")
	}
	s.Erase(5)
	if s.Contains(5) {
		t.Logf("s.Erase(10) Contains(10) error\n")
	}

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		t.Logf("%v\n", iter.Value())
	}

	if s.Size() != 9 {
		t.Logf("s.Size() error: %v\n", s.Size())
	}
}

func TestSet_Cal(t *testing.T) {
	s1 := New()
	s2 := New()
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		s1.Insert(rand.Int() % 20)
		s2.Insert(rand.Int() % 20)
	}
	t.Logf("s1: %v", s1)
	t.Logf("s2: %v", s2)
	t.Logf("s1 & s2: %v", s1.Intersect(s2))
	t.Logf("s1 | s2: %v", s1.Union(s2))
	t.Logf("s1 - s2: %v", s1.Diff(s2))
	t.Logf("s2 - s1: %v", s2.Diff(s1))
}
