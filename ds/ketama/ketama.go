package ketama

import (
	"github.com/liyue201/gostl/algorithm/hash"
	"github.com/liyue201/gostl/ds/map"
	"github.com/liyue201/gostl/utils/sync"
	gosync "sync"
)

var (
	defaultReplicas = 10
	defaultLocker   sync.FakeLocker
)

const Salt = "ni9fkh72hgh1g"

type Option struct {
	replicas int
	locker   sync.Locker
}

type Options func(option *Option)

func WithThreadSave() Options {
	return func(option *Option) {
		option.locker = &gosync.RWMutex{}
	}
}

func WithReplicas(replicas int) Options {
	return func(option *Option) {
		option.replicas = replicas
	}
}

type Ketama struct {
	locker   sync.Locker
	replicas int
	m        *treemap.Map
}

// New new a ketama ring
// Ketama is a thread-safe implementation of consistent hash
func New(opts ...Options) *Ketama {
	option := Option{
		replicas: defaultReplicas,
		locker:   defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	this := &Ketama{
		replicas: option.replicas,
		locker:   option.locker,
		m:        treemap.New(),
	}
	return this
}

// Empty returns true if  Ketama is empty, or false if not empty
func (this *Ketama) Empty() bool {
	this.locker.RLock()
	defer this.locker.RUnlock()

	return this.m.Size() == 0
}

// Add add nodes to ketama ring
func (this *Ketama) Add(nodes ...string) {
	this.locker.Lock()
	defer this.locker.Unlock()

	for _, node := range nodes {
		hashs := hash.GenHashInts([]byte(Salt+node), this.replicas)
		for i := 0; i < this.replicas; i++ {
			key := hashs[i]
			if !this.m.Contains(key) {
				this.m.Insert(key, node)
			}
		}
	}
}

// Get remove nodes from ketama ring
func (this *Ketama) Remove(nodes ...string) {
	this.locker.Lock()
	defer this.locker.Unlock()

	for _, node := range nodes {
		hashs := hash.GenHashInts([]byte(Salt+node), this.replicas)
		for i := 0; i < this.replicas; i++ {
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

	this.locker.Lock()
	defer this.locker.Unlock()

	iter := this.m.LowerBound(hash)
	if iter.IsValid() {
		return iter.Value().(string), true
	}
	return this.m.First().Value().(string), true
}
