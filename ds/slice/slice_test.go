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

func TestIter(t *testing.T) {
	a := make([]float32, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, float32(i))
	}
	t.Logf("a: %v", a)
	sliceA := Float32Slice(a)

	i := float32(0)
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		if iter.Value().(float32) != float32(i) {
			t.Fatalf("expect %v, but get %v",  float32(i), iter.Value().(float32))
		}
		i++
		iter.SetValue(i * 10)
	}

	for iter := sliceA.Last(); iter.IsValid();  iter.Prev() {
		if iter.Value().(float32) != float32(i  * 10) {
			t.Fatalf("expect %v, but get %v",  float32(i * 10), iter.Value().(float32))
		}
		i--
	}
}