package set

import (
	. "github.com/liyue201/gostl/comparator"
	"testing"
)
 
func TestSet(t *testing.T) {
	s := New(BuiltinTypeComparator)
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
