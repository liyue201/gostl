package ketama

import (
	"github.com/liyue201/gostl/algorithm/hash"
	"github.com/liyue201/gostl/ds/map"
	"sync"
)

var (
	defaultReplicas = 10
)

const Salt = "ni9fkh72hgh1g"

type Ketama struct {
	sync.RWMutex
	option Option
	m      *treemap.Map
}

type Option struct {
	replicas int
}

type Options func(option *Option)

func WithReplicas(replicas int) Options {
	return func(option *Option) {
		option.replicas = replicas
	}
}

// New new a ketama ring
// Ketama is a thread-safe implementation of consistent hash
func New(opts ...Options) *Ketama {
	option := Option{
		replicas: defaultReplicas,
	}
	this := &Ketama{
		option: option,
		m:      treemap.New(),
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
		hashs := hash.GenHashInts([]byte(Salt+node), this.option.replicas)
		for i := 0; i < this.option.replicas; i++ {
			key := hashs[i]
			if !this.m.Contains(key) {
				this.m.Insert(key, node)
			}
		}
	}
}

// Get remove nodes from ketama ring
func (this *Ketama) Remove(nodes ...string) {
	this.Lock()
	defer this.Unlock()
	for _, node := range nodes {
		hashs := hash.GenHashInts([]byte(Salt+node), this.option.replicas)
		for i := 0; i < this.option.replicas; i++ {
			key := hashs[i]
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

	hashs := hash.GenHashInts([]byte(Salt+key), 1)
	hash := hashs[0]

	this.Lock()
	defer this.Unlock()

	iter := this.m.LowerBound(hash)
	if iter.IsValid() {
		return iter.Value().(string), true
	}
	return this.m.First().Value().(string), true
}
