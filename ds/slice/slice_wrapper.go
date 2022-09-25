package slice

//SliceWrapper wraps a slice in order to provide functions related to iterators
type SliceWrapper[T any] struct {
	slice []T
}

// NewSliceWrapper creates a SliceWrapper
func NewSliceWrapper[T any](slice []T) *SliceWrapper[T] {
	return &SliceWrapper[T]{
		slice: slice,
	}
}

// Attach update the internal slice to newSlice
func (s *SliceWrapper[T]) Attach(newSlice []T) {
	s.slice = newSlice
}

// Len returns the length of s
func (s *SliceWrapper[T]) Len() int {
	return len(s.slice)
}

// At returns the value at position
func (s *SliceWrapper[T]) At(position int) T {
	if position < 0 || position >= s.Len() {
		panic("Out off range")
	}
	return s.slice[position]
}

// Set sets value at position
func (s *SliceWrapper[T]) Set(position int, val T) {
	if position < 0 || position >= s.Len() {
		return
	}
	s.slice[position] = val
}

// Begin returns the first iterator of s
func (s *SliceWrapper[T]) Begin() *SliceIterator[T] {
	return s.First()
}

// End returns the end iterator of s
func (s *SliceWrapper[T]) End() *SliceIterator[T] {
	return &SliceIterator[T]{s: s,
		position: s.Len(),
	}
}

// First returns the first iterator of s
func (s *SliceWrapper[T]) First() *SliceIterator[T] {
	return &SliceIterator[T]{s: s,
		position: 0,
	}
}

// Last returns the last iterator of s
func (s *SliceWrapper[T]) Last() *SliceIterator[T] {
	return &SliceIterator[T]{s: s,
		position: s.Len(),
	}
}
