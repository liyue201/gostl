package slice

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"testing"
)

func TestIntSlice(t *testing.T) {
	a := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	t.Logf("a: %v", a)
	sliceA := IntSlice(a)

	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))

	t.Logf("a: %v", a)

	for i, v := range a {
		if v != 9-i {
			t.Fatalf("sort error: v expect %v, but get %v", 9-i, v)
		}
	}
}
