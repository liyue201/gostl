package set

import (
	"testing"
)

func TestMultiSet(t *testing.T) {
	mset := NewMultiSet()

	mset.Insert(1)
	mset.Insert(5)
	mset.Insert(1)

	if mset.Size() != 3 {
		t.Errorf("size error: expect %v, but get %v", 3, mset.Size())
	}
	for iter := mset.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v\n", iter.Value())
	}
	t.Logf("=======================")

	mset.Erase(1)

	for iter := mset.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v\n", iter.Value())
	}
}
