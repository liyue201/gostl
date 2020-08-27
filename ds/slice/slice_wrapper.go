package slice

import "reflect"

//SliceWrapper wraps a slice in order to provide functions related to iterators
type SliceWrapper struct {
	sliceValue reflect.Value
}

// NewSliceWrapper creates a SliceWrapper
func NewSliceWrapper(slice interface{}) *SliceWrapper {
	return &SliceWrapper{
		sliceValue: reflect.ValueOf(slice),
	}
}

// Attach update the internal slice to newSlice
func (s *SliceWrapper) Attach(newSlice interface{}) {
	if reflect.ValueOf(newSlice).Kind() == s.sliceValue.Kind() {
		s.sliceValue = reflect.ValueOf(newSlice)
	}
}

// Len returns the length of s
func (s *SliceWrapper) Len() int {
	return s.sliceValue.Len()
}

// At returns the value at position
func (s *SliceWrapper) At(position int) interface{} {
	if position < 0 || position >= s.Len() {
		return nil
	}
	return s.sliceValue.Index(position).Interface()
}

// Set sets value at position
func (s *SliceWrapper) Set(position int, val interface{}) {
	if position < 0 || position >= s.Len() {
		return
	}
	s.sliceValue.Index(position).Set(reflect.ValueOf(val))
}

// Begin returns the first iterator of s
func (s *SliceWrapper) Begin() *SliceIterator {
	return s.First()
}

// End returns the end iterator of s
func (s *SliceWrapper) End() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}

// First returns the first iterator of s
func (s *SliceWrapper) First() *SliceIterator {
	return &SliceIterator{s: s,
		position: 0,
	}
}

// Last returns the last iterator of s
func (s *SliceWrapper) Last() *SliceIterator {
	return &SliceIterator{s: s,
		position: s.Len(),
	}
}
