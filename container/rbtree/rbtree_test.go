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
	for i := 0; i < 10; i++ {
		iter := tree.FindLowerBoundNode(i)
		if iter == nil {
			t.Fatalf("findIt %v nil", i)
		}
		if iter.Value != i+10000 {
			t.Fatalf("findIt %v: found %v, %v ", i, iter.Key, iter.Value)
		}
	}
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+20000)
	}

	for i := 0; i < 10; i++ {
		iter := tree.FindLowerBoundNode(i)
		if iter == nil {
			t.Fatalf("findIt %v nil", i)
		}
		for n := iter; n != nil; n = n.Next() {
			if n.Key != i {
				break
			}
			t.Logf("travesal: %v = %v ", n.Key, n.Value)
		}
	}

	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.Key, n.Value)
	}

	t.Logf("==============")
	tree.Delete(tree.findFirstNode(7))
	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.Key, n.Value)
	}

	t.Logf("==============")
	tree.Delete(tree.findFirstNode(5))
	tree.Delete(tree.findFirstNode(7))
	tree.Delete(tree.findFirstNode(1))
	tree.Delete(tree.findFirstNode(8))
	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.Key, n.Value)
	}
}
