package skiplist

import (
	"github.com/liyue201/gostl/comparator"
	"math/rand"
	"sync"
	"time"
)

var (
	defaultKeyComparator = comparator.BuiltinTypeComparator
	defaultMaxLevel      = 10
)

type Option struct {
	keyCmp   comparator.Comparator
	maxLevel int
}

type Options func(option *Option)

func WithKeyComparator(cmp comparator.Comparator) Options {
	return func(option *Option) {
		option.keyCmp = cmp
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
	sync.RWMutex
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
	}
	for _, opt := range opts {
		opt(&option)
	}
	l := &Skiplist{
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
	this.Lock()
	defer this.Unlock()
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
	this.RLock()
	defer this.RUnlock()

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
	this.Lock()
	defer this.Unlock()

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
	this.RLock()
	defer this.RUnlock()
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
