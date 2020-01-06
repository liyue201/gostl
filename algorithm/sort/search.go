package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//BinarySearch returns true if exist an element witch value is val in the range [first, last), or false if not exist
func BinarySearch(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) bool {
	if len(cmp) == 0 {
		return binarySearch(first, last, val, comparator.BuiltinTypeComparator)
	}
	return binarySearch(first, last, val, cmp[0])
}

func binarySearch(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) bool {
	if !first.IsValid() || first.Position() >= last.Position() {
		return false
	}
	left := first.Clone().(iterator.RandomAccessIterator)
	right := last.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)

	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := right.IteratorAt(midPos)
		cmpRet := cmp(val, midIter.Value())
		if cmpRet == 0 {
			return true
		} else if cmpRet < 0 {
			right = midIter.Prev().(iterator.RandomAccessIterator)
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator)
		}
	}
	return false
}

//LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [first, last), or iterator last if not exist.
func LowerBound(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	if len(cmp) == 0 {
		return lowerBound(first, last, val, comparator.BuiltinTypeComparator)
	}
	return lowerBound(first, last, val, cmp[0])
}

func lowerBound(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	if !first.IsValid() || first.Position() >= last.Position() {
		return last.Clone().(iterator.RandomAccessIterator)
	}

	left := first.Clone().(iterator.RandomAccessIterator)
	right := last.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)
	if cmp(val, right.Value()) > 0 {
		return last.Clone().(iterator.RandomAccessIterator)
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) <= 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.RandomAccessIterator)
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator)
		}
	}
	return first.IteratorAt(pos)
}

//UpperBound returns the iterator pointing to the first element greater than val in the range [first, last), or iterator last if not exist.
func UpperBound(first, last iterator.RandomAccessIterator, val interface{}, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	if len(cmp) == 0 {
		return upperBound(first, last, val, comparator.BuiltinTypeComparator)
	}
	return upperBound(first, last, val, cmp[0])
}

func upperBound(first, last iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	if !first.IsValid() || first.Position() >= last.Position() {
		return last.Clone().(iterator.RandomAccessIterator)
	}

	left := first.Clone().(iterator.RandomAccessIterator)
	right := last.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)
	if cmp(val, right.Value()) >= 0 {
		return last.Clone().(iterator.RandomAccessIterator)
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) < 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.RandomAccessIterator)
		} else {
			left = midIter.Next().(iterator.RandomAccessIterator)
		}
	}
	return first.IteratorAt(pos)
}
