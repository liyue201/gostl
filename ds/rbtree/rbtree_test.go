package rbtree

import (
	"testing"
)

func TestRbTeeFind(t *testing.T) {
	tree := New()
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
		if iter.value != i+10000 {
			t.Fatalf("findIt %v: found %v, %v ", i, iter.key, iter.value)
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
			if n.key != i {
				break
			}
			t.Logf("travesal: %v = %v ", n.key, n.value)
		}
	}
}

func TestRbTeeDelete(t *testing.T) {
	tree := New()
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+10000)
	}
	for i := 0; i < 10; i++ {
		tree.Insert(i, i+20000)
	}

	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.key, n.value)
	}

	t.Logf("==============")
	tree.Delete(tree.findFirstNode(7))
	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.key, n.value)
	}

	t.Logf("==============")
	tree.Delete(tree.findFirstNode(5))
	tree.Delete(tree.findFirstNode(7))
	tree.Delete(tree.findFirstNode(1))
	tree.Delete(tree.findFirstNode(8))
	for n := tree.Begin(); n != nil; n = n.Next() {
		t.Logf("found: %v = %v ", n.key, n.value)
	}
}

func TestTravesal(t *testing.T) {
	tree := New()
	for i := 20; i >= 1; i-- {
		tree.Insert(i, 0)
	}
	for n := tree.Begin(); n != nil; n = n.Next() {
		k := n.key
		var p interface{}
		d := 0
		if n.parent != nil {
			p = n.parent.key
			if n.parent.left == n {
				d = 0
			} else {
				d = 1
			}
		}
		t.Logf("found: %v, %v, %v ", k, p, d)
	}
}

func TestIterator(t *testing.T) {
	tree := New()
	for i := 10; i >= 1; i-- {
		tree.Insert(i, i+100)
	}
	for iter := tree.IterFirst(); iter.IsValid(); iter.Next() {
		t.Logf("found: %v, %v", iter.Key(), iter.Value())
	}
}
