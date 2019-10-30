package treemap

import (
	"testing"
) 

func TestMap(t *testing.T) {
	m := New()

	for i := 9; i >= 0; i-- {
		m.Insert(i, i+1000)
	}  
	if m.Size() != 10 {
		t.Errorf("size error: %v", m.Size())
	}
	i := 0
	t.Log("===========================")
	for iter := m.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v : %v", i, iter.Value().(int))
		if iter.Value().(int) != i+1000 {
			t.Errorf("get %v error: %v", i, iter.Value().(int))
		} 
		i++
	}
	if iter := m.Find(8); iter.Value().(int) != 8 + 1000 {
		t.Errorf("get key %v, value %v", 8, iter.Value().(int))
	}
}
