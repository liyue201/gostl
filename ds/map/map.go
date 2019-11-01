package treemap

import (
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/comparator"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	gosync "sync"
)

var (
	defaultKeyComparator = BuiltinTypeComparator
	defaultLocker        sync.FakeLocker
)

type Option struct {
	keyCmp Comparator
	locker sync.Locker
}

type Options func(option *Option)

func WithKeyComparator(cmp Comparator) Options {
	return func(option *Option) {
		option.keyCmp = cmp
	}
}

// Note that iterators are not thread safe, and it is useless to turn on the setting option here.
// so don't use iterators in multi goroutines
func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

type Map struct {
	tree   *rbtree.RbTree
	locker sync.Locker
}

func New(opts ...Options) *Map {
	option := Option{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &Map{tree: rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		locker: option.locker,
	}
}

//Insert inserts key-value to the map
func (this *Map) Insert(key, value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(key)
	if node != nil {
		node.SetValue(value)
		return
	}
	this.tree.Insert(key, value)
}

//Get returns the value by key if found, or nil if not found
func (this *Map) Get(key interface{}) interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(key)
	if node != nil {
		return node.Value()
	}
	return nil
}

//Erase erases node by key in the Map
func (this *Map) Erase(key interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(key)
	if node != nil {
		this.tree.Delete(node)
	}
}

//Erase erases node by iter in the Map
func (this *Map) EraseIter(iter ConstKvIterator) {
	this.locker.Lock()
	defer this.locker.Unlock()

	mpIter, ok := iter.(*MapIterator)
	if ok {
		this.tree.Delete(mpIter.node)
	}
}

//Begin returns the ConstIterator related to value in the map, or an invalid iterator if not exist.
func (this *Map) Find(key interface{}) ConstKvIterator {
	this.locker.RUnlock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(key)
	return &MapIterator{node: node}
}

//LowerBound returns the first ConstIterator that equal or greater than key in the Map
func (this *Map) LowerBound(key interface{}) ConstKvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindLowerBoundNode(key)
	return &MapIterator{node: node}
}

//Begin returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *Map) Begin() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.First()
}

//First returns the ConstIterator with the minimum key in the Map, return nil if empty.
func (this *Map) First() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &MapIterator{node: this.tree.First()}
}

//Last returns the ConstIterator with the maximum key in the Map, return nil if empty.
func (this *Map) Last() KvIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &MapIterator{node: this.tree.Last()}
}

//Clear clears the Map
func (this *Map) Clear() {
	this.locker.Lock()
	defer this.locker.Unlock()

	this.tree.Clear()
}

// Contains returns true if key in the Map. otherwise returns false.
func (this *Map) Contains(key interface{}) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	if this.tree.Find(key) != nil {
		return true
	}
	return false
}

// Contains returns the size of Map
func (this *Map) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.tree.Size()
}

// Traversal traversals elements in map, it will not stop until to the end or visitor returns false
func (this *Map) Traversal(visitor visitor.KvVisitor) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	this.tree.Traversal(visitor)
}
