package treemap

import (
	"github.com/liyue201/gostl/containers/rbtree"
	. "github.com/liyue201/gostl/comparator"
	. "github.com/liyue201/gostl/iterator"
)

type MultiMap struct {
	tree    *rbtree.RbTree
	cmpFunc Comparator
}

func NewMultiMap(cmp Comparator) *MultiMap {
	return &MultiMap{tree: rbtree.New(cmp),
		cmpFunc: cmp,
	}
}

//Insert inserts key-value to the set
func (this *MultiMap) Insert(key, value interface{}) {
	this.tree.Insert(key, value)
}

//Erase erases key in the Map
func (this *MultiMap) Erase(key interface{}) {
	node := this.tree.FindNode(key)
	for node != nil && this.cmpFunc(node.Key(), key) == 0 {
		nextNode := node.Next()
		this.tree.Delete(node)
		node = nextNode
	}
}

//Begin returns the ConstIterator related to key in the set, return nil if not exist.
func (this *MultiMap) Find(key interface{}) ConstKvIterator {
	node := this.tree.FindNode(key)
	if node == nil {
		return nil
	}
	return &MapIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than key in the Map
func (this *MultiMap) LowerBound(key interface{}) ConstKvIterator {
	node := this.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *MultiMap) Begin() KvIterator {
	return this.First()
}

//First returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *MultiMap) First() KvIterator {
	return &MapIterator{node: this.tree.First()}
}

//Last returns the ConstIterator with the maximum key in the Map, return nil if empty.
func (this *MultiMap) Last() KvIterator {
	return &MapIterator{node: this.tree.Last()}
}

//Clear clears the Map
func (this *MultiMap) Clear() {
	this.tree.Clear()
}

// Contains returns true if value in the Map. otherwise returns false.
func (this *MultiMap) Contains(value interface{}) bool {
	if this.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (this *MultiMap) Size() int {
	return this.tree.Size()
}
