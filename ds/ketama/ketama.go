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

const salt = "ni9fkh72hgh1g"

// Options hold Ketama's options
type Options struct {
	replicas int
	locker   sync.Locker
}

// Options configures Ketama's options
type Option func(option *Options)

// WithGoroutineSafe configures goroutine-safe
func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &gosync.RWMutex{}
	}
}

// WithReplicas configures replicas
func WithReplicas(replicas int) Option {
	return func(option *Options) {
		option.replicas = replicas
	}
}

// Ketama is an implementation of consistent-hash
type Ketama struct {
	locker   sync.Locker
	replicas int
	m        *treemap.Map
}

// New new a ketama ring
// Ketama is a goroutine-safe implementation of consistent hash
func New(opts ...Option) *Ketama {
	option := Options{
		replicas: defaultReplicas,
		locker:   defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}
	k := &Ketama{
		replicas: option.replicas,
		locker:   option.locker,
		m:        treemap.New(),
	}
	return k
}

// Empty returns true if  Ketama is empty, or false if not empty
func (k *Ketama) Empty() bool {
	k.locker.RLock()
	defer k.locker.RUnlock()

	return k.m.Size() == 0
}

// Add add nodes to ketama ring
func (k *Ketama) Add(nodes ...string) {
	k.locker.Lock()
	defer k.locker.Unlock()

	for _, node := range nodes {
		hashs := hash.GenHashInts([]byte(salt+node), k.replicas)
		for i := 0; i < k.replicas; i++ {
			key := hashs[i]
			if !k.m.Contains(key) {
				k.m.Insert(key, node)
			}
		}
	}
}

// Remove remove nodes from ketama ring
func (k *Ketama) Remove(nodes ...string) {
	k.locker.Lock()
	defer k.locker.Unlock()

	for _, node := range nodes {
		hashs := hash.GenHashInts([]byte(salt+node), k.replicas)
		for i := 0; i < k.replicas; i++ {
			key := hashs[i]
			iter := k.m.Find(key)
			if iter.IsValid() && iter.Value() == node {
				k.m.EraseIter(iter)
			}
		}
	}
}

// Get returns the node which closest to key in the clockwise direction
func (k *Ketama) Get(key string) (string, bool) {
	if k.Empty() {
		return "", false
	}

	hashs := hash.GenHashInts([]byte(salt+key), 1)
	hash := hashs[0]

	k.locker.Lock()
	defer k.locker.Unlock()

	iter := k.m.LowerBound(hash)
	if iter.IsValid() {
		return iter.Value().(string), true
	}
	return k.m.First().Value().(string), true
}
