package iterator

// const iterator
type ConstIterator interface {
	IsValid() bool
	Next() ConstIterator
	Value() interface{}
	Clone() interface{}
} 

// mutable iterator
type Iterator interface {
	ConstIterator
	SetValue(value interface{}) error
}

//const key-value type iterator
type ConstKvIterator interface {
	ConstIterator
	Key() interface{}
}

// mutable key-value type iterator
type KvIterator interface {
	ConstKvIterator
	SetValue(value interface{}) error
}

//const bidirectional iterator
type ConstBidIterator interface {
	ConstIterator
	Prev() ConstBidIterator
}

//mutable bidirectional iterator
type BidIterator interface {
	ConstBidIterator
	SetValue(value interface{}) error
}

//const key-value type bidirectional iterator
type ConstKvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
}

//mutable key-value type bidirectional iterator
type KvBidIterator interface {
	ConstKvIterator
	Prev() ConstBidIterator
	SetValue(value interface{}) error
}
