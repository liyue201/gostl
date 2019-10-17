package treemap

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
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

//Begin returns the ConstIterator with the min value in the Map, return nil if empty.
func (this *Map) Begin() ConstKvIterator {
	return &MapIterator{node: this.tree.Begin()}
}

//End returns ConstIterator with nil value in the Map
func (this *Map) End() ConstKvIterator {
	return &MapIterator{node: nil}
}

//Begin returns the ConstIterator with the max value in the Map, return nil if empty.
func (this *Map) RBegin() ConstKvIterator {
	return &MapReverseIterator{node: this.tree.RBegin()}
}

//REnd returns ConstIterator with nil value in the set
func (this *Map) REnd() ConstKvIterator {
	return &MapReverseIterator{node: nil}
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
