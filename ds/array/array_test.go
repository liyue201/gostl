package array

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"testing"
)

func TestArray(t *testing.T) {
	a := New(10)
	if a.Size() != 10 {
		t.Fatalf("array size error")
	}

	va := 10
	a.Fill(va)
	for i := 0; i < a.Size(); i++ {
		if val := a.At(i); val.(int) != va {
			t.Fatalf("expect %v, but get %v", va, val)
		}
	}

	b := New(10)
	vb := 66
	b.Fill(vb)
	a.SwapArray(b)

	for i := 0; i < a.Size(); i++ {
		if val := a.At(i); val.(int) != vb {
			t.Fatalf("a: expect %v, but get %v", vb, val)
		}
		if val := b.At(i); val.(int) != va {
			t.Fatalf("b: expect %v, but get %v", va, val)
		}
	}

	for i := 0; i < a.Size(); i++ {
		a.Set(i, i)
	}

	t.Logf("Traversal a:")
	i := 0
	for iter := a.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v ", iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i, iter.Value().(int))
		}
		i++
	}

	t.Logf(" Reverse traversal a:")
	i = a.Size() - 1
	for iter := a.Last(); iter.IsValid(); iter.Next() {
		t.Logf("%v ", iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i, iter.Value().(int))
		}
		i--
	}
}

func TestSort(t *testing.T) {
	a := New(10)
	if a.Size() != 10 {
		t.Fatalf("array size error")
	}
	for i := 0; i < 10; i++ {
		a.Set(i, 10-i)
	}
	sort.Stable(a.Begin(), a.End(), comparator.BuiltinTypeComparator)
	t.Logf("a: %v", a.String())
}
