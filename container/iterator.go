package container

// const iterator
type ConstIterator interface {
	Next() ConstIterator
	Value() interface{}
	Equal(ConstIterator) bool
} 

// iterator
type Iterator interface {
	Next() Iterator
	Value() interface{}
	Equal(Iterator) bool
	Set(value interface{}) error
}

//const key-value type iterator
type ConstKvIterator interface {
	Next() ConstKvIterator
	Key() interface{}
	Value() interface{}
	Equal(ConstKvIterator) bool
}

// key-value type iterator
type KvIterator interface {
	Next() KvIterator
	Key() interface{}
	Value() interface{}
	SetValue(value interface{}) error
	Equal(KvIterator) bool
}
