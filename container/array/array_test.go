package array

import "testing"
  
func TestArray(t *testing.T) {
	a := New(10) 
	if a.Size() != 10 {
		t.Fatalf("array size error")
	}
	a.Fill(88)
	for i := 0; i < a.Size(); i++ {
		if val, _ := a.At(i); val != 88 {
			t.Fatalf("expect 88, but get %v", val)
		}
	} 

	b := New(10)
	b.Fill(66)
	a.Swap(b) 
 
	for i := 0; i < a.Size(); i++ {
		if val, _ := a.At(i); val != 66 {
			t.Fatalf("expect 66, but get %v", val)
		}
		if val, _ := b.At(i); val != 88 {
			t.Fatalf("expect 88, but get %v", val)
		}
	} 

	for i := 0; i < a.Size(); i++ {
		a.Set(i, i)
	}

	t.Logf("Traversal a:")
	i := 0
	for iter := a.Begin(); !iter.Equal(a.End()); iter = iter.Next() {
		t.Logf("%v ",  iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i, iter.Value().(int))
		}
		i++
	}

	t.Logf(" Reverse traversal a:")
	i = a.Size() - 1
	for iter := a.RBegin(); !iter.Equal(a.REnd()); iter = iter.Next() {
		t.Logf("%v ",  iter.Value().(int))
		if iter.Value().(int) != i {
			t.Fatalf("expect %v, but get %v", i,  iter.Value().(int))
		}
		i--
	}
}
