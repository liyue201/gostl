package skiplist

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/sync"
	"github.com/liyue201/gostl/utils/visitor"
	"math/rand"
	gosync "sync"
	"time"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultMaxLevel      = 10
	defaultLocker        sync.FakeLocker
)

type Option struct {
	keyCmp   comparator.Comparator
	maxLevel int
	locker   sync.Locker
}

type Options func(option *Option)

func WithKeyComparator(cmp comparator.Comparator) Options {
	return func(option *Option) {
		option.keyCmp = cmp
	}
}

func WithThreadSafe() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

func WithMaxLevel(maxLevel int) Options {
	return func(option *Option) {
		option.maxLevel = maxLevel
	}
}

type Node struct {
	next []*Element
}

type Element struct {
	Node
	key   interface{}
	value interface{}
}

type Skiplist struct {
	locker         sync.Locker
	head           Node
	maxLevel       int
	keyCmp         comparator.Comparator
	len            int
	prevNodesCache []*Node
	rander         *rand.Rand
}

func New(opts ...Options) *Skiplist {
	option := Option{
		keyCmp:   defaultKeyComparator,
		maxLevel: defaultMaxLevel,
		locker:   defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	l := &Skiplist{
		locker:   option.locker,
		maxLevel: option.maxLevel,
		keyCmp:   option.keyCmp,
		rander:   rand.New(rand.NewSource(time.Now().Unix())),
	}
	l.head.next = make([]*Element, l.maxLevel)
	l.prevNodesCache = make([]*Node, l.maxLevel)
	return l
}

// Insert inserts a key-value pair into skiplist
func (this *Skiplist) Insert(key, value interface{}) {
	this.locker.Lock()
	defer this.locker.Unlock()
	prevs := this.findPrevNodes(key)

	if prevs[0].next[0] != nil && this.keyCmp(prevs[0].next[0].key, key) == 0 {
		//same key, update value
		prevs[0].next[0].value = value
		return
	}

	level := this.randomLevel()

	e := &Element{
		key:   key,
		value: value,
		Node: Node{
			next: make([]*Element, level),
		},
	}

	for i := range e.next {
		e.next[i] = prevs[i].next[i]
		prevs[i].next[i] = e
	}

	this.len++
}

// Get gets the value associated with the key passed if exist, or nil if not exist
func (this *Skiplist) Get(key interface{}) interface{} {
	this.locker.RLock()
	defer this.locker.RUnlock()

	var pre = &this.head
	for i := this.maxLevel - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := this.keyCmp(cur.key, key)
			if cmpRet == 0 {
				return cur.value
			}
			if cmpRet > 0 {
				break
			}
			pre = &cur.Node
		}
	}
	return nil
}

// Remove removes the element associated with the key passed and returns true if exist,or false if not exist
func (this *Skiplist) Remove(key interface{}) bool {
	this.locker.Lock()
	defer this.locker.Unlock()

	prevs := this.findPrevNodes(key)
	element := prevs[0].next[0]
	if element == nil {
		return false
	}
	if element != nil && this.keyCmp(element.key, key) != 0 {
		return false
	}

	for i, v := range element.next {
		prevs[i].next[i] = v
	}
	this.len--
	return true
}

// Len returns the number of elements in the skiplist
func (this *Skiplist) Len() int {
	this.locker.RLock()
	defer this.locker.RUnlock()
	return this.len
}

func (this *Skiplist) randomLevel() int {
	total := uint64(1)<<uint64(this.maxLevel) - 1 // 2^n-1
	k := this.rander.Uint64() % total
	levelN := uint64(1) << (uint64(this.maxLevel) - 1)

	level := 1
	for total -= levelN; total > k; level++ {
		levelN >>= 1
		total -= levelN
	}
	return level
}

func (this *Skiplist) findPrevNodes(key interface{}) []*Node {
	prevs := this.prevNodesCache
	prev := &this.head
	for i := this.maxLevel - 1; i >= 0; i-- {
		if this.head.next[i] != nil {
			for next := prev.next[i]; next != nil; next = next.next[i] {
				if this.keyCmp(next.key, key) >= 0 {
					break
				}
				prev = &next.Node
			}
		}
		prevs[i] = prev
	}
	return prevs
}

// Traversal traversals elements in Skiplist, it will stop until to the end or visitor returns false
func (this *Skiplist) Traversal(visitor visitor.KvVisitor) {
	this.locker.RLock()
	defer this.locker.RUnlock()

	for e := this.head.next[0]; e != nil; e = e.next[0] {
		if !visitor(e.key, e.value) {
			return
		}
	}
}

// Keys returns all keys in the Skiplist
func (this *Skiplist) Keys() []interface{} {
	var keys []interface{}
	this.Traversal(func(key, value interface{}) bool {
		keys = append(keys, key)
		return false
	})
	return keys
}
