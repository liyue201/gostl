package array

import "testing"

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
	a.Swap(b)

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
	for iter := a.Begin(); !iter.Equal(a.End()); iter = iter.Next() {
		t.Logf("%v ", iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i, iter.Value().(int))
		}
		i++
	}

	t.Logf(" Reverse traversal a:")
	i = a.Size() - 1
	for iter := a.RBegin(); !iter.Equal(a.REnd()); iter = iter.Next() {
		t.Logf("%v ", iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i, iter.Value().(int))
		}
		i--
	}
}
