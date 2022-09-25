package container

//Container is an interface for a base linear container
type Container[T any] interface {
	PushBack(value T)
	PushFront(value T)
	Front() T
	Back() T
	PopFront() T
	PopBack() T
	Empty() bool
	Size() int
	String() string
	Clear()
}
