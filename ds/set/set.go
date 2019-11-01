package set

import (
	"fmt"
	"github.com/liyue201/gostl/ds/rbtree"
	. "github.com/liyue201/gostl/utils/comparator"
	. "github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	gosync "sync"
)

const (
	Empty = 0
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

type Set struct {
	tree   *rbtree.RbTree
	keyCmp Comparator
	locker sync.Locker
}

func New(opts ...Options) *Set {
	option := Option{
		keyCmp: defaultKeyComparator,
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	return &Set{
		tree:   rbtree.New(rbtree.WithKeyComparator(option.keyCmp)),
		keyCmp: option.keyCmp,
		locker: option.locker,
	}
}

// Insert inserts element to the Set
func (this *Set) Insert(element interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(element)
	if node != nil {
		return
	}
	this.tree.Insert(element, Empty)
}

// Erase erases element in the Set
func (this *Set) Erase(element interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()

	node := this.tree.FindNode(element)
	if node != nil {
		this.tree.Delete(node)
	}
}

// Begin returns the ConstIterator related to element in the Set,or an invalid iterator if not exist.
func (this *Set) Find(element interface{}) ConstIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindNode(element)
	return &SetIterator{node: node}
}

// LowerBound returns the first ConstIterator that equal or greater than element in the Set
func (this *Set) LowerBound(element interface{}) ConstIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	node := this.tree.FindLowerBoundNode(element)
	return &SetIterator{node: node}
}

// Begin returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *Set) Begin() ConstIterator {
	return this.First()
}

// First returns the ConstIterator with the minimum element in the Set, return nil if empty.
func (this *Set) First() ConstBidIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &SetIterator{node: this.tree.First()}
}

// Last returns the ConstIterator with the maximum element in the Set, return nil if empty.
func (this *Set) Last() ConstBidIterator {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return &SetIterator{node: this.tree.Last()}
}

// Clear clears the Set
func (this *Set) Clear() {
	this.locker.Lock()
	defer this.locker.Lock()

	this.tree.Clear()
}

// Contains returns true if element in the Set. otherwise returns false.
func (this *Set) Contains(element interface{}) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	if this.tree.Find(element) != nil {
		return true
	}
	return false
}

// Contains returns the size of Set
func (this *Set) Size() int {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.tree.Size()
}

// Traversal traversals elements in set, it will not stop until to the end or visitor returns false
func (this *Set) Traversal(visitor visitor.Visitor) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	for node := this.tree.First(); node != nil; node = node.Next() {
		if !visitor(node.Key()) {
			break
		}
	}
}

// String returns the set's elements in string format
func (this *Set) String() string {
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

// Intersect returns a set with the common elements in this set and the other set
// Please ensure this set and other set uses the same keyCmp
func (this *Set) Intersect(other *Set) *Set {
	this.locker.RLock()
	defer this.locker.RUnlock()

	set := New(WithKeyComparator(this.keyCmp))
	thisIter := this.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for thisIter.IsValid() && otherIter.IsValid() {
		cmp := this.keyCmp(thisIter.Key(), otherIter.Key())
		if cmp == 0 {
			set.tree.Insert(thisIter.Key(), Empty)
			thisIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			thisIter.Next()
		} else {
			otherIter.Next()
		}
	}
	return set
}

// Union returns  a set with the all elements in this set and the other set
// Please ensure this set and other set uses the same keyCmp
func (this *Set) Union(other *Set) *Set {
	this.locker.RLock()
	defer this.locker.RUnlock()

	set := New(WithKeyComparator(this.keyCmp))
	thisIter := this.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for thisIter.IsValid() && otherIter.IsValid() {
		cmp := this.keyCmp(thisIter.Key(), otherIter.Key())
		if cmp == 0 {
			set.tree.Insert(thisIter.Key(), Empty)
			thisIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			set.tree.Insert(thisIter.Key(), Empty)
			thisIter.Next()
		} else {
			set.tree.Insert(otherIter.Key(), Empty)
			otherIter.Next()
		}
	}
	for ; thisIter.IsValid(); thisIter.Next() {
		set.tree.Insert(thisIter.Key(), Empty)
	}
	for ; otherIter.IsValid(); otherIter.Next() {
		set.tree.Insert(otherIter.Key(), Empty)
	}
	return set
}

// Diff returns a set with the elements in this set but not in the other set
// Please ensure this set and other set uses the same keyCmp
func (this *Set) Diff(other *Set) *Set {
	this.locker.RLock()
	defer this.locker.RUnlock()

	set := New(WithKeyComparator(this.keyCmp))
	thisIter := this.tree.IterFirst()
	otherIter := other.tree.IterFirst()
	for thisIter.IsValid() && otherIter.IsValid() {
		cmp := this.keyCmp(thisIter.Key(), otherIter.Key())
		if cmp == 0 {
			thisIter.Next()
			otherIter.Next()
		} else if cmp < 0 {
			set.tree.Insert(thisIter.Key(), Empty)
			thisIter.Next()
		} else {
			otherIter.Next()
		}
	}
	for ; thisIter.IsValid(); thisIter.Next() {
		set.tree.Insert(thisIter.Key(), Empty)
	}
	return set
}
