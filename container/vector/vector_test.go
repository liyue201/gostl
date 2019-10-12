package vector

import "testing"

func TestVectorBase(t *testing.T) {
	v := New(10)

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

	if val, err := v.Front(); err != nil || val.(int) != 1 {
		t.Fatalf("val error: %v , %v", val, err.Error())
	}
	if val, err := v.Back(); err != nil || val.(int) != 2 {
		t.Fatalf("val error: %v , %v", val, err.Error())
	}

	v.ShrinkToFit()
	if v.Capacity() != 2 {
		t.Fatalf("capacity error")
	}
	if val, err := v.At(0); err != nil || val.(int) != 1 {
		t.Fatalf("val error: %v , %v", val, err.Error())
	}
	v.Reserve(20)
	if v.Capacity() != 20 || v.Size() != 2 {
		t.Fatalf("capacity or size error")
	}
	if val, err := v.At(1); err != nil || val.(int) != 2 {
		t.Fatalf("val error: %v , %v", val, err.Error())
	}
	v.Clear()
	if v.Size() != 0 || !v.Empty() {
		t.Fatalf("size error")
	}
}

func TestModifyVector(t *testing.T) {
	v := New(10)
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	//[1,2,3]
	if val, err := v.PopBack(); err != nil || val.(int) != 3 {
		t.Fatalf("val error: %v , %v", val, err.Error())
	}
	//[1 2]
	v.PushBack(4)
	//[1 2 4]

	v.SetAt(1, 9)
	if val, err := v.At(1); err != nil || val.(int) != 9 {
		t.Fatalf("val error: %v , %v", val, err.Error())
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
	v := New(10)
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)

	i := 0
	t.Logf("v: %v", v.String())
	for iter := v.Begin(); !iter.Equal(v.End()); iter = iter.Next() {
		if val, _ := v.At(i); val.(int) != iter.Value().(int) {
			t.Fatalf("value error: expect %v, but get %v", val, iter.Value().(int))
		}
		i++
	}

	i = 3
	for iter := v.RBegin(); !iter.Equal(v.REnd()); iter = iter.Next() {
		if val, _ := v.At(i); val.(int) != iter.Value().(int) {
			t.Fatalf("traversal value error: expect %v, but get %v", val, iter.Value().(int))
		}
		i--
	}
	iter := v.Erase(v.Begin())
	t.Logf("v: %v", v.String())
	if iter.Value().(int) != 2 {
		t.Fatalf("erase error: expect %v, but get %v", 2, iter.Value().(int))
	}
	next := iter.Next()
	v.EraseRange(iter, next)

	t.Logf("v: %v", v.String())
	if iter.Value().(int) != 2 {
		t.Fatalf("erase error: expect %v, but get %v", 2, iter.Value().(int))
	}
}
