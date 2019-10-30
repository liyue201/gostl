package sort

import (
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/comparator"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	a := make([]int, 0)
	for i := 0; i < 3; i++ {
		a = append(a, i+1)
	}
	sliceA := slice.IntSlice(a)
	t.Logf("a : %v", a)
	for {
		ok := NextPermutation(sliceA.Begin(), sliceA.End(), comparator.BuiltinTypeComparator)
		if !ok {
			break
		}
		t.Logf("a : %v", a)
	}
}

func TestPrePermutation(t *testing.T) {
	a := make([]int, 0)
	for i := 0; i < 3; i++ {
		a = append(a, 3-i)
	}
	sliceA := slice.IntSlice(a)
	t.Logf("a : %v", a)
	for {
		ok := NextPermutation(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
		if !ok {
			break
		}
		t.Logf("a : %v", a)
	}
}
