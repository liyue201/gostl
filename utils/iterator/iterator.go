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
	SetValue(value T)
}

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator[K, V any] interface {
	ConstIterator[V]
	Key() K
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator[K, V any] interface {
	ConstKvIterator[K, V]
	SetValue(value V)
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
type ConstKvBidIterator[K, V any] interface {
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator[K, V any] interface {
	KvIterator[K, V]
	BidIterator[V]
}

// RandomAccessIterator is an interface of mutable random access iterator
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	//IteratorAt returns a new iterator at position
	IteratorAt(position int) RandomAccessIterator[T]
	Position() int
}
