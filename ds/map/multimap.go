package treemap

import (
	. "github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/visitor"
)

type MultiMap struct {
	tree   *rbtree.RbTree
	keyCmp Comparator
}

func NewMultiMap(opts ...Options) *MultiMap {
	option := Option{
		keyCmp: defaultKeyComparator,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiMap{tree: rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
	}
}

//Insert inserts key-value to the set
func (this *MultiMap) Insert(key, value interface{}) {
	this.tree.Insert(key, value)
}

//Get returns the first node's value by key if found, or nil if not found
func (this *MultiMap) Get(key interface{}) interface{} {
	node := this.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases key in the Map
func (this *MultiMap) Erase(key interface{}) {
	node := this.tree.FindNode(key)
	for node != nil && this.keyCmp(node.Key(), key) == 0 {
		nextNode := node.Next()
		this.tree.Delete(node)
		node = nextNode
	}
}

//Begin returns the ConstIterator related to key in the set, or an invalid iterator if not exist.
func (this *MultiMap) Find(key interface{}) ConstKvIterator {
	node := this.tree.FindNode(key)
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

// Traversal traversals elements in the map, it will not stop until to the end or visitor returns false
func (this *MultiMap) Traversal(visitor visitor.KvVisitor) {
	this.tree.Traversal(visitor)
}