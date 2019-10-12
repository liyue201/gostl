package gostl

type ConstIterator interface {
	Next() ConstIterator
	Data() interface{}
	Equal(ConstIterator) bool
}

type Iterator interface {
	Next() Iterator
	Data() interface{}
	Equal(Iterator) bool
	Assign(interface{}) error
}

type ReverseIterator interface {
	Next() ReverseIterator
	Data() interface{}
	Equal(ReverseIterator) bool
	Assign(interface{}) error
}
