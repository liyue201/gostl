package skiplist

import (
	"github.com/liyue201/gostl/comparator"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	list := New(5, comparator.BuiltinTypeComparator)

	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := rand.Int() % 100
		list.Insert(key, i)
		m[key] = i
	}
	for key, v := range m {
		ret := list.Get(key)
		//t.Logf("%v = %v", key, ret)
		if ret != v {
			t.Fatalf("get value of %v error, expect %v but get %v", key, v, ret)
		}
	}
	if len(m) != list.Len() {
		t.Fatalf("get list len error, expect %v but get %v", len(m), list.Len())
	}
}

func TestRemove(t *testing.T) {
	list := New(10, comparator.BuiltinTypeComparator)

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		key := rand.Int() % 1000
		list.Insert(key, i)
		m[key] = i
	}

	t.Logf("len = %v %v", len(m), list.Len())
	if len(m) != list.Len() {
		t.Fatalf("11 get list len error, expect %v but get %v", len(m), list.Len())
	}

	for i := 0; i < 300; i++ {
		key := rand.Int() % 1000
		list.Remove(key)
		delete(m, key)

		key2 := rand.Int() % 10440

		list.Insert(key2, key)
		m[key2] = key
	}

	for key, v := range m {
		ret := list.Get(key)
		//t.Logf("%v = %v", key, ret)
		if ret != v {
			t.Fatalf("get value of %v error, expect %v but get %v", key, v, ret)
		}
	}
	t.Logf("len = %v %v", len(m), list.Len())
	if len(m) != list.Len() {
		t.Fatalf("get list len error, expect %v but get %v", len(m), list.Len())
	}
}
