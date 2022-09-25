package visitor

// Visitor is a function use to visit a data structure
type Visitor[V any] func(value V) bool

// KvVisitor is a function use to visit a key-value type data structure
type KvVisitor[K, V any] func(key K, value V) bool
