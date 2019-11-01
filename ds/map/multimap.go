package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/comparator"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
)

type MultiMap struct {
	tree   *rbtree.RbTree
	keyCmp Comparator
	locker sync.Locker
}

func NewMultiMap(opts ...Options) *MultiMap {
	option := Option{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &MultiMap{tree: rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

//Insert inserts key-value to the set
func (this *MultiMap) Insert(key, value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.tree.Insert(key, value)
}

//Get returns the first node's value by key if found, or nil if not found
func (this *MultiMap) Get(key interface{}) interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases key in the Map
func (this *MultiMap) Erase(key interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(key)
	for node != nil && this.keyCmp(node.Key(), key) == 0 {
		nextNode := node.Next()
		this.tree.Delete(node)
		node = nextNode
	}
}

//Begin returns the ConstIterator related to key in the set, or an invalid iterator if not exist.
func (this *MultiMap) Find(key interface{}) ConstKvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than key in the Map
func (this *MultiMap) LowerBound(key interface{}) ConstKvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *MultiMap) Begin() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.First()
}

//First returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *MultiMap) First() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &MapIterator{node: this.tree.First()}
}

//Last returns the ConstIterator with the maximum key in the Map, return nil if empty.
func (this *MultiMap) Last() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &MapIterator{node: this.tree.Last()}
}

//Clear clears the Map
func (this *MultiMap) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.tree.Clear()
}

// Contains returns true if value in the Map. otherwise returns false.
func (this *MultiMap) Contains(value interface{}) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	if this.tree.Find(value) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (this *MultiMap) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.tree.Size()
}

// Traversal traversals elements in the map, it will not stop until to the end or visitor returns false
func (this *MultiMap) Traversal(visitor visitor.KvVisitor) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	this.tree.Traversal(visitor)
}
