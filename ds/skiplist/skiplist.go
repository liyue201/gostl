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

// Options holds Skiplist's options
type Options struct {
	keyCmp   comparator.Comparator
	maxLevel int
	locker   sync.Locker
}

// Option is a function used to set Options
type Option func(option *Options)

// WithKeyComparator sets Key comparator option
func WithKeyComparator(cmp comparator.Comparator) Option {
	return func(option *Options) {
		option.keyCmp = cmp
	}
}

// WithGoroutineSafe sets Skiplist goroutine-safety,
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithMaxLevel sets max level of Skiplist
func WithMaxLevel(maxLevel int) Option {
	return func(option *Options) {
		option.maxLevel = maxLevel
	}
}

// Node is a list node
type Node struct {
	next []*Element
}

// Element is a kind of node with key-value data
type Element struct {
	Node
	key   any
	value any
}

// Skiplist is a kind of data structure which can search quickly by exchanging space for time
type Skiplist struct {
	locker         sync.Locker
	head           Node
	maxLevel       int
	keyCmp         comparator.Comparator
	len            int
	prevNodesCache []*Node
	rander         *rand.Rand
}

// New news a Skiplist
func New(opts ...Option) *Skiplist {
	option := Options{
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

// Insert inserts a key-value pair into the skiplist
func (sl *Skiplist) Insert(key, value any) {
	sl.locker.Lock()
	defer sl.locker.Unlock()
	prevs := sl.findPrevNodes(key)

	if prevs[0].next[0] != nil && sl.keyCmp(prevs[0].next[0].key, key) == 0 {
		//same key, update value
		prevs[0].next[0].value = value
		return
	}

	level := sl.randomLevel()

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

	sl.len++
}

// Get returns the value associated with the passed key if the key is in the skiplist, otherwise returns nil
func (sl *Skiplist) Get(key any) any {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	var pre = &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := sl.keyCmp(cur.key, key)
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

// Remove removes the key-value pair associated with the passed key and returns true if the key is in the skiplist, otherwise returns false
func (sl *Skiplist) Remove(key any) bool {
	sl.locker.Lock()
	defer sl.locker.Unlock()

	prevs := sl.findPrevNodes(key)
	element := prevs[0].next[0]
	if element == nil {
		return false
	}
	if element != nil && sl.keyCmp(element.key, key) != 0 {
		return false
	}

	for i, v := range element.next {
		prevs[i].next[i] = v
	}
	sl.len--
	return true
}

// Len returns the amount of key-value pair in the skiplist
func (sl *Skiplist) Len() int {
	sl.locker.RLock()
	defer sl.locker.RUnlock()
	return sl.len
}

func (sl *Skiplist) randomLevel() int {
	total := uint64(1)<<uint64(sl.maxLevel) - 1 // 2^n-1
	k := sl.rander.Uint64() % total
	levelN := uint64(1) << (uint64(sl.maxLevel) - 1)

	level := 1
	for total -= levelN; total > k; level++ {
		levelN >>= 1
		total -= levelN
	}
	return level
}

func (sl *Skiplist) findPrevNodes(key any) []*Node {
	prevs := sl.prevNodesCache
	prev := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		if sl.head.next[i] != nil {
			for next := prev.next[i]; next != nil; next = next.next[i] {
				if sl.keyCmp(next.key, key) >= 0 {
					break
				}
				prev = &next.Node
			}
		}
		prevs[i] = prev
	}
	return prevs
}

// Traversal traversals elements in the skiplist, it will stop until to the end or the visitor returns false
func (sl *Skiplist) Traversal(visitor visitor.KvVisitor) {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		if !visitor(e.key, e.value) {
			return
		}
	}
}

// Keys returns all keys in the skiplist
func (sl *Skiplist) Keys() []any {
	var keys []any
	sl.Traversal(func(key, value any) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}
