package container

type Container interface {
	PushBack(value interface{})
	PushFront(value interface{})
	Front() interface{}
	Back() interface{}
	PopFront() interface{}
	PopBack() interface{}
	Empty() bool
	Size() int
	String() string
	Clear()
}
