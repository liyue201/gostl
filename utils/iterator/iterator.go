package iterator

// const iterator
type ConstIterator interface {
	IsValid() bool
	Next() ConstIterator
	Value() interface{}
	Clone() ConstIterator
	Equal(other ConstIterator) bool
}

// mutable iterator
type Iterator interface {
	ConstIterator
	SetValue(value interface{}) error
}

// const key-value type iterator
type ConstKvIterator interface {
	ConstIterator
	Key() interface{}
}

// mutable key-value type iterator
type KvIterator interface {
	ConstKvIterator
	SetValue(value interface{}) error
}

// const bidirectional iterator
type ConstBidIterator interface {
	ConstIterator
	Prev() ConstBidIterator
}

// mutable bidirectional iterator
type BidIterator interface {
	ConstBidIterator
	SetValue(value interface{}) error
}

// const key-value type bidirectional iterator
type ConstKvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
}

// mutable key-value type bidirectional iterator
type KvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
	SetValue(value interface{}) error
}

// random access iterator
type RandomAccessIterator interface {
	BidIterator
	//IteratorAt returns a new iterator at position
	IteratorAt(position int) RandomAccessIterator
	Position() int
}
