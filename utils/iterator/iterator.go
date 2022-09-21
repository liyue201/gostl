package iterator

// ConstIterator is an interface of const iterator
type ConstIterator[T any] interface {
	IsValid() bool
	Next() ConstIterator[T]
	Value() T
	Clone() ConstIterator[T]
	Equal(other ConstIterator[T]) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any] interface {
	ConstIterator[T]
	SetValue(value interface{})
}

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator[T any] interface {
	ConstIterator[T]
	Key() T
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator[T any] interface {
	ConstKvIterator[T]
	SetValue(value T)
}

// ConstBidIterator is an interface of const bidirectional iterator
type ConstBidIterator[T any] interface {
	ConstIterator[T]
	Prev() ConstBidIterator[T]
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator[T any] interface {
	ConstBidIterator[T]
	SetValue(value T)
}

// ConstKvBidIterator is an interface of const key-value type bidirectional iterator
type ConstKvBidIterator[T any] interface {
	ConstKvIterator[T]
	BidIterator[T]
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator[T any] interface {
	KvIterator[T]
	BidIterator[T]
}

// RandomAccessIterator is an interface of mutable random access iterator
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	//IteratorAt returns a new iterator at position
	IteratorAt(position int) RandomAccessIterator[T]
	Position() int
}
