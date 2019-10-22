package sort

import (
	"github.com/liyue201/gostl/iterator"
	"github.com/liyue201/gostl/uitls/comparator"
)   

//LowerBound returns true if exist an element witch value is val in the range [begin, end), or false if not exist
func BinarySearch(begin, end iterator.SortableIterator, val interface{}, cmp comparator.Comparator) bool {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return false
	}
	left := begin.Clone().(iterator.SortableIterator)
	right := end.Clone().(iterator.SortableIterator).Prev().(iterator.SortableIterator)

	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := right.IteratorAt(midPos)
		cmpRet := cmp(val, midIter.Value())
		if cmpRet == 0 {
			return true
		} else if cmpRet < 0 {
			right = midIter.Prev().(iterator.SortableIterator)
		} else {
			left = midIter.Next().(iterator.SortableIterator)
		}
	}
	return false
}

//LowerBound returns the iterator of the first element greater than or equal to value in the range [begin, end), or iterator end if not exist.
func LowerBound(begin, end iterator.SortableIterator, val interface{}, cmp comparator.Comparator) iterator.SortableIterator {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return end.Clone().(iterator.SortableIterator)
	}

	left := begin.Clone().(iterator.SortableIterator)
	right := end.Clone().(iterator.SortableIterator).Prev().(iterator.SortableIterator)
	if cmp(val, right.Value()) > 0 {
		return end.Clone().(iterator.SortableIterator)
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) <= 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.SortableIterator)
		} else {
			left = midIter.Next().(iterator.SortableIterator)
		}
	}
	return begin.IteratorAt(pos)
}

//LowerBound returns the iterator of the first element greater than val in the range [begin, end), or iterator end if not exist.
func UpperBound(begin, end iterator.SortableIterator, val interface{}, cmp comparator.Comparator) iterator.SortableIterator {
	if !begin.IsValid() || begin.Position() >= end.Position() {
		return end.Clone().(iterator.SortableIterator)
	}

	left := begin.Clone().(iterator.SortableIterator)
	right := end.Clone().(iterator.SortableIterator).Prev().(iterator.SortableIterator)
	if cmp(val, right.Value()) >= 0 {
		return end.Clone().(iterator.SortableIterator)
	}
	var pos int
	for left.Position() <= right.Position() {
		midPos := (left.Position() + right.Position()) >> 1
		midIter := left.IteratorAt(midPos)
		if cmp(val, midIter.Value()) < 0 {
			pos = midIter.Position()
			right = midIter.Prev().(iterator.SortableIterator)
		} else {
			left = midIter.Next().(iterator.SortableIterator)
		}
	}
	return begin.IteratorAt(pos)
}
