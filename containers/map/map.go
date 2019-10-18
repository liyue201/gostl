package treemap

import (
	"github.com/liyue201/gostl/containers/rbtree"
	. "github.com/liyue201/gostl/uitls/comparator"
	. "github.com/liyue201/gostl/uitls/iterator"
)

type Map struct {
	tree *rbtree.RbTree
}

func New(cmp Comparator) *Map {
	return &Map{tree: rbtree.New(cmp)}
}

//Insert inserts key-value to the set
func (this *Map) Insert(key, value interface{}) {
	node := this.tree.FindNode(key)
	if node != nil {
		node.SetValue(value)
		return
	}
	this.tree.Insert(key, value)
}

//Erase erases value in the Map
func (this *Map) Erase(key interface{}) {
	node := this.tree.FindNode(key)
	if node != nil {
		this.tree.Delete(node)
	}
}

//Begin returns the ConstIterator related to value in the set, return nil if not exist.
func (this *Map) Find(key interface{}) ConstKvIterator {
	node := this.tree.FindNode(key)
	if node == nil {
		return nil
	}
	return &MapIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than key in the Map
func (this *Map) LowerBound(key interface{}) ConstKvIterator {
	node := this.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *Map) Begin() KvIterator {
	return this.First()
}

//First returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *Map) First() KvIterator {
	return &MapIterator{node: this.tree.Begin()}
}

//Last returns the ConstIterator with the maximum key in the Map, return nil if empty.
func (this *Map) Last() KvIterator {
	return &MapIterator{node: this.tree.Last()}
}

//Clear clears the Map
func (this *Map) Clear() {
	this.tree.Clear()
}

// Contains returns true if value in the Map. otherwise returns false.
func (this *Map) Contains(value interface{}) bool {
	if this.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (this *Map) Size() int {
	return this.tree.Size()
}
