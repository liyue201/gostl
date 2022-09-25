package slice

// ISlice is an interface of Slice for iterator
type ISlice[T any] interface {
	Len() int
	At(position int) T
	Set(position int, val T)
}
