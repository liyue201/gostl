package set

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/comparator"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

type MultiSet struct {
	tree   *rbtree.RbTree
	keyCmp Comparator
	locker sync.Locker
}

func NewMultiSet(opts ...Options) *MultiSet {
	option := Option{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiSet{
		tree:   rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

// Insert inserts element to the MultiSet
func (this *MultiSet) Insert(element interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.tree.Insert(element, Empty)
}

// Erase erases all node with element in this MultiSet
func (this *MultiSet) Erase(element interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(element)
	for node != nil && this.keyCmp(node.Key(), element) == 0 {
		nextNode := node.Next()
		this.tree.Delete(node)
		node = nextNode
	}
}

// Begin returns the ConstIterator related to element in the MultiSet,or an invalid iterator if not exist.
func (this *MultiSet) Find(element interface{}) ConstIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(element)
	return &SetIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than element in the MultiSet
func (this *MultiSet) LowerBound(element interface{}) ConstIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *MultiSet) Begin() ConstIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.First()
}

// First returns the ConstIterator with the minimum element in the MultiSet, return nil if empty.
func (this *MultiSet) First() ConstBidIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &SetIterator{node: this.tree.First()}
}

//Last returns the ConstIterator with the maximum element in the MultiSet, return nil if empty.
func (this *MultiSet) Last() ConstBidIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &SetIterator{node: this.tree.Last()}
}

// Clear clears the MultiSet
func (this *MultiSet) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.tree.Clear()
}

// Contains returns true if element in the MultiSet. otherwise returns false.
func (this *MultiSet) Contains(element interface{}) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	if this.tree.Find(element) != nil {
		return true
	}
	return false
}

// Contains returns the size of MultiSet
func (this *MultiSet) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.tree.Size()
}

// Traversal traversals elements in MultiSet, it will not stop until to the end or visitor returns false
func (this *MultiSet) Traversal(visitor visitor.Visitor) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	for node := this.tree.First(); node != nil; node = node.Next() {
		if !visitor(node.Key()) {
			break
		}
	}
}

// String returns the set's elements in string format
func (this *MultiSet) String() string {
	str := "["
	this.Traversal(func(value interface{}) bool {
		if str != "[" {
			str += " "
		}
		str += fmt.Sprintf("%v", value)
		return true
	})
	str += "]"
	return str
}
