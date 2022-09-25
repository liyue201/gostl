package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//BinarySearch returns true if exist an element witch value is val in the range [first, last), or false if not exist
func BinarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	return binarySearch(first, last, val, cmp)
}

func binarySearch[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) bool {
	if !first.IsValid() || first.Position() >= last.Position() {
		return false
	}
	left := first.Clone().(iterator.RandomAccessIterator[T])
	right := last.Clone().(iterator.RandomAccessIterator[T]).Prev().(iterator.RandomAccessIterator[T])

	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := right.IteratorAt(midPos)
		cmpRet := cmp(val, midIter.Value())
		if cmpRet == 0 {
			return true
		} else if cmpRet < 0 {
			right = midIter.Prev().(iterator.RandomAccessIterator[T])
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator[T])
		}
	}
	return false
}

//LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [first, last), or iterator last if not exist.
func LowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	return lowerBound(first, last, val, cmp)
}

func lowerBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	if !first.IsValid() || first.Position() >= last.Position() {
		return last.Clone().(iterator.RandomAccessIterator[T])
	}

	left := first.Clone().(iterator.RandomAccessIterator[T])
	right := last.Clone().(iterator.RandomAccessIterator[T]).Prev().(iterator.RandomAccessIterator[T])
	if cmp(val, right.Value()) > 0 {
		return last.Clone().(iterator.RandomAccessIterator[T])
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) <= 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.RandomAccessIterator[T])
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator[T])
		}
	}
	return first.IteratorAt(pos)
}

//UpperBound returns the iterator pointing to the first element greater than val in the range [first, last), or iterator last if not exist.
func UpperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	return upperBound(first, last, val, cmp)
}

func upperBound[T any](first, last iterator.RandomAccessIterator[T], val T, cmp comparator.Comparator[T]) iterator.RandomAccessIterator[T] {
	if !first.IsValid() || first.Position() >= last.Position() {
		return last.Clone().(iterator.RandomAccessIterator[T])
	}

	left := first.Clone().(iterator.RandomAccessIterator[T])
	right := last.Clone().(iterator.RandomAccessIterator[T]).Prev().(iterator.RandomAccessIterator[T])
	if cmp(val, right.Value()) >= 0 {
		return last.Clone().(iterator.RandomAccessIterator[T])
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) < 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.RandomAccessIterator[T])
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator[T])
		}
	}
	return first.IteratorAt(pos)
}
