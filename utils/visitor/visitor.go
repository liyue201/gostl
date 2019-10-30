package visitor

type Visitor func(value interface{}) bool

type KvVisitor func(key, value interface{}) bool
