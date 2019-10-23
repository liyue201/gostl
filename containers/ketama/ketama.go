package ketama

import (
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/containers/map"
	"hash/crc32"
	"strconv"
	"sync"
)

type HashFunc func(data []byte) uint32

var (
	defaultReplicas = 10
	defaultHashFunc = crc32.ChecksumIEEE
)

type Ketama struct {
	sync.RWMutex
	option Option
	m      *treemap.Map
}

type Option struct {
	replicas int
	hash     HashFunc
}

type Options func(option *Option)

func WithReplicas(replicas int) Options {
	return func(option *Option) {
		option.replicas = replicas
	}
}

func WithHashFunc(hash HashFunc) Options {
	return func(option *Option) {
		option.hash = hash
	}
}

// New new a ketama ring
// Ketama is a thread-safe implementation of consistent hash
func New(opts ...Options) *Ketama {
	option := Option{
		replicas: defaultReplicas,
		hash:     defaultHashFunc,
	}
	this := &Ketama{
		option: option,
		m:      treemap.New(comparator.BuiltinTypeComparator),
	}
	for _, opt := range opts {
		opt(&this.option)
	}
	return this
}

// Empty returns true if  Ketama is empty, or false if not empty
func (this *Ketama) Empty() bool {
	this.RLock()
	defer this.RUnlock()

	return this.m.Size() == 0
}

// Add add nodes to ketama ring
func (this *Ketama) Add(nodes ...string) {
	this.Lock()
	defer this.Unlock()

	for _, node := range nodes {
		for i := 0; i < this.option.replicas; i++ {
			key := int(this.option.hash([]byte(makeRingKey(node, i))))
			if !this.m.Contains(key) {
				this.m.Insert(key, node)
			}
		}
	}
}

func makeRingKey(node string, index int) string {
	return node + "-" + strconv.Itoa(index)
}

// Get remove nodes from ketama ring
func (this *Ketama) Remove(nodes ...string) {
	this.Lock()
	defer this.Unlock()
	for _, node := range nodes {
		for i := 0; i < this.option.replicas; i++ {
			key := int(this.option.hash([]byte(makeRingKey(node, i))))
			iter := this.m.Find(key)
			if iter.IsValid() && iter.Value() == node {
				this.m.EraseIter(iter)
			}
		}
	}
}

// Get returns the node which closest to key in the clockwise direction
func (this *Ketama) Get(key string) (string, bool) {
	if this.Empty() {
		return "", false
	}

	hash := int(this.option.hash([]byte(key)))

	this.Lock()
	defer this.Unlock()

	iter := this.m.LowerBound(hash)
	if iter.IsValid() {
		return iter.Value().(string), true
	}
	return this.m.First().Value().(string), true
}
