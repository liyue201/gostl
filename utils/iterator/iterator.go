package iterator

// ConstIterator is an interface of const iterator
type ConstIterator interface {
	IsValid() bool
	Next() ConstIterator
	Value() interface{}
	Clone() ConstIterator
	Equal(other ConstIterator) bool
}

// Iterator is an interface of mutable iterator
type Iterator interface {
	ConstIterator
	SetValue(value interface{}) error
}

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator interface {
	ConstIterator
	Key() interface{}
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator interface {
	ConstKvIterator
	SetValue(value interface{}) error
}

// ConstBidIterator is an interface of const bidirectional iterator
type ConstBidIterator interface {
	ConstIterator
	Prev() ConstBidIterator
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator interface {
	ConstBidIterator
	SetValue(value interface{}) error
}

// ConstKvBidIterator is an interface of const key-value type bidirectional iterator
type ConstKvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
	SetValue(value interface{}) error
}

// RandomAccessIterator is an interface of mutable random access iterator
type RandomAccessIterator interface {
	BidIterator
	//IteratorAt returns a new iterator at position
	IteratorAt(position int) RandomAccessIterator
	Position() int
}
