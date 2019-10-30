package vector

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"testing"
)

func TestVectorBase(t *testing.T) {
	v := New(WithCapacity(10))

	if !v.Empty() {
		t.Fatalf("is not emtpy")
	}
	if v.Capacity() != 10 {
		t.Fatalf("capacity error")
	}
	v.PushBack(1)
	v.PushBack(2)
	if v.Empty() {
		t.Fatalf("v is emtpy")
	}
	if v.Size() != 2 {
		t.Fatalf("size error")
	}

	if val := v.Front(); val == nil || val.(int) != 1 {
		t.Fatalf("val error: %v", val)
	}
	if val := v.Back(); val == nil || val.(int) != 2 {
		t.Fatalf("val error: %v ", val)
	}

	v.ShrinkToFit()
	if v.Capacity() != 2 {
		t.Fatalf("capacity error")
	}
	if val := v.At(0); val == nil || val.(int) != 1 {
		t.Fatalf("val error: %v ", val)
	}
	v.Reserve(20)
	if v.Capacity() != 20 || v.Size() != 2 {
		t.Fatalf("capacity or size error")
	}
	if val := v.At(1); val.(int) != 2 {
		t.Fatalf("val error: %v", val)
	}
	v.Clear()
	if v.Size() != 0 || !v.Empty() {
		t.Fatalf("size error")
	}
}

func TestModifyVector(t *testing.T) {
	v := New()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	//[1,2,3]
	if val := v.PopBack(); val == nil || val.(int) != 3 {
		t.Fatalf("val error: %v", val)
	}
	//[1 2]
	v.PushBack(4)
	//[1 2 4]

	v.SetAt(1, 9)
	if val := v.At(1); val == nil || val.(int) != 9 {
		t.Fatalf("val error: %v", val)
	}
	//[1 9 4]

	v.InsertAt(0, 8)
	//[8 1 9 4]

	t.Logf("data: %v", v)
	if v.String() != "[8 1 9 4]" {
		t.Fatalf("data error: %v", v.String())
	}
}

func TestVectorIter(t *testing.T) {
	v := New()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)
	//[1 2 3 4]

	i := 0
	t.Logf("v: %v", v.String())
	for iter := v.Begin(); iter.IsValid(); iter.Next() {
		if val := v.At(i); val.(int) != iter.Value().(int) {
			t.Fatalf("value error: expect %v, but get %v", val, iter.Value().(int))
		}
		i++
	}

	i = 3
	for iter := v.Last(); iter.IsValid(); iter.Prev() {
		if val := v.At(i); val.(int) != iter.Value().(int) {
			t.Fatalf("traversal value error: expect %v, but get %v", val, iter.Value().(int))
		}
		i--
	}
	iter := v.Erase(v.Begin())
	t.Logf("v: %v", v.String())
	if iter.Value().(int) != 2 {
		t.Fatalf("erase error: expect %v, but get %v", 2, iter.Value().(int))
	}
	//[2 3 4]

	v.PushBack(5)
	v.PushBack(6)
	//[2 3 4 5 6]

	iter = v.EraseRange(v.Begin().Next(), v.Begin().Next().Next().Next())
	//[2 5 6]
	t.Logf("v: %v", v.String())
	if iter.Value().(int) != 5 {
		t.Fatalf("erase error: expect 5, but get %v", iter.Value().(int))
	}
	if v.String() != "[2 5 6]" {
		t.Fatalf("erase error: expect [2 5 6], but get %v", v.String())
	}
	//
	iter = v.Begin()
	iter = v.Insert(iter, 7)
	//[7 2 5 6]
	if iter.Value().(int) != 7 {
		t.Fatalf("erase error: expect 7, but get %v", iter.Value().(int))
	}
	if v.String() != "[7 2 5 6]" {
		t.Fatalf("erase error: expect [7 2 5 6], but get %v", v.String())
	}
}

func TestSort(t *testing.T) {
	v := New()
	for i := 10; i >= 0; i-- {
		v.PushBack(i)
	}
	sort.Sort(v.Begin(), v.End(), comparator.BuiltinTypeComparator)
	for i := 0; i < v.Size(); i++ {
		t.Logf("%v", v.At(i))
		if i != v.At(i).(int) {
			t.Fatalf("sort error")
		}
	}

	sort.Sort(v.Begin(), v.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	for i := 0; i < v.Size(); i++ {
		t.Logf("%v", v.At(i))
		if v.Size()-i-1 != v.At(i).(int) {
			t.Fatalf("sort error")
		}
	}
}
