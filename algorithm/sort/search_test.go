package sort

import (
	"github.com/liyue201/gostl/containers/deque"
	"github.com/liyue201/gostl/comparator"
	"testing"
)  
     
func TestSearch(t *testing.T) {
	a := deque.New(10)
	a.PushBack(1)
	a.PushBack(3)
	a.PushFront(5)
	a.PushFront(4)
	a.PushBack(7)
	a.PushFront(3)
	a.PushBack(15)

	Sort(a.Begin(), a.End(), comparator.BuiltinTypeComparator)
  
	t.Logf("a: %v", a)
	if !BinarySearch(a.Begin(), a.End(), 5, comparator.BuiltinTypeComparator) {
		t.Fatalf("BinarySearch 5 error: expect true,but get flase")
	}
	if BinarySearch(a.Begin(), a.End(), 10, comparator.BuiltinTypeComparator) {
		t.Fatalf("BinarySearch 10 error: expect false,but get ture")
	}

	iter := LowerBound(a.Begin(), a.End(), 3, comparator.BuiltinTypeComparator)
	if !iter.IsValid() {
		t.Fatalf("LowerBound 3 error, not found")
	}
	if iter.Value() != 3 && iter.Clone().Next().Value() != 3 {
		t.Fatalf("LowerBound 3 error:")
	}

	iter = UpperBound(a.Begin(), a.End(), 4, comparator.BuiltinTypeComparator)
	if !iter.IsValid() {
		t.Fatalf("UpperBound 4 error")
	}
	if iter.Value() != 5 {
		t.Fatalf("UpperBound 4 error: expect %v, but get %v", 5, iter.Value())
	}
}