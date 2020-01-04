package visitor

// Visitor is a function use to visit a data structure
type Visitor func(value interface{}) bool

// Visitor is a function use to visit a key-value type data structure
type KvVisitor func(key, value interface{}) bool
