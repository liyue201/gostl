package sort
  
import (
	"github.com/liyue201/gostl/utils/iterator"
	"github.com/liyue201/gostl/utils/comparator"
)

//LowerBound returns true if exist an element witch value is val in the range [begin, end), or false if not exist
func BinarySearch(begin, end iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) bool {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return false
	}
	left := begin.Clone().(iterator.RandomAccessIterator)
	right := end.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)

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

//LowerBound returns the iterator pointing to the first element greater than or equal to value passed in the range [begin, end), or iterator end if not exist.
func LowerBound(begin, end iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return end.Clone().(iterator.RandomAccessIterator)
	}

	left := begin.Clone().(iterator.RandomAccessIterator)
	right := end.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)
	if cmp(val, right.Value()) > 0 {
		return end.Clone().(iterator.RandomAccessIterator)
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
	return begin.IteratorAt(pos)
}

//LowerBound returns the iterator pointing to the first element greater than val in the range [begin, end), or iterator end if not exist.
func UpperBound(begin, end iterator.RandomAccessIterator, val interface{}, cmp comparator.Comparator) iterator.RandomAccessIterator {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return end.Clone().(iterator.RandomAccessIterator)
	}

	left := begin.Clone().(iterator.RandomAccessIterator)
	right := end.Clone().(iterator.RandomAccessIterator).Prev().(iterator.RandomAccessIterator)
	if cmp(val, right.Value()) >= 0 {
		return end.Clone().(iterator.RandomAccessIterator)
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
	return begin.IteratorAt(pos)
}
