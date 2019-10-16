package set

import (
	. "github.com/liyue201/gostl/container"
	"github.com/liyue201/gostl/container/rbtree"
)

const (
	Empty = 0
)

type Set struct {
	tree *rbtree.RbTree
}

func New(cmp Comparator) *Set {
	return &Set{tree: rbtree.New(cmp)}
}

//Insert inserts value to the set
func (this *Set) Insert(value interface{}) {
	if this.tree.Find(value) != nil {
		return
	}
	this.tree.Insert(value, Empty)
}

//Erase erases value in the set
func (this *Set) Erase(value interface{}) {
	node := this.tree.FindItr(value)
	if node != nil {
		this.tree.Delete(node)
	}
}

//Begin returns the ConstIterator related to value in the set, return nil if not exist.
func (this *Set) Find(value interface{}) ConstIterator {
	node := this.tree.FindItr(value)
	if node == nil {
		return nil
	}
	return &SetIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than value in the set
func (this *Set) LowerBound(value interface{}) ConstIterator {
	node := this.tree.FindLowerBoundNode(value)
	return &SetIterator{node: node}
}

//Begin returns the ConstIterator with the min value in the set, return nil if empty.
func (this *Set) Begin() ConstIterator {
	return &SetIterator{node: this.tree.Begin()}
}

//End returns ConstIterator with nil value in the set
func (this *Set) End() ConstIterator {
	return &SetIterator{node: nil}
}

//Begin returns the ConstIterator with the max value in the set, return nil if empty.
func (this *Set) RBegin() ConstIterator {
	//todo:
	return nil
}

//REnd returns ConstIterator with nil value in the set
func (this *Set) REnd() ConstIterator {
	//return &SetIterator{node: nil}
	return nil
}

func (this *Set) EqualRange(value interface{}) ConstIterator {
	//todo:
	return nil
}

//Clear clears the set
func (this *Set) Clear() {
	this.tree.Clear()
}

// Contains returns true if value in the set. otherwise returns false.
func (this *Set) Contains(value interface{}) bool {
	if this.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of set
func (this *Set) Size() int {
	return this.tree.Size()
}
