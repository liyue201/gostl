package rbtree

import (
	"testing"
)

func TestRbTee(t *testing.T) {
	cmp := func(a, b interface{}) int {
		if a.(int) == b.(int) {
			return 0
		}
		if a.(int) < b.(int) {
			return -1
		}
		return 1
	}
	tree := New(cmp)
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+10000)
	}
	for i := 0; i < 10; i++ {
		val := tree.Find(i)
		if val == nil || val.(int) != i+10000 {
			t.Fatalf("find %d not found.", i)
		}
	}
	tree.Delete(7)
	if tree.Find(7) != nil || tree.size != 9 {
		t.Fatalf("delete key:7 error")
	}

	tree.Delete(0)
	if tree.Find(0) != nil || tree.size != 8 {
		t.Fatalf("delete key:7 error")
	}

	count := 0
	for i := 0; i < 10; i++ {
		val := tree.Find(i)
		if val != nil {
			count++
			if val.(int) != i+10000 {
				t.Fatalf("get %d error.", i)
			}
		}
	}
	if count != 8 {
		t.Fatalf("count != 8")
	}
}
