package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//Stable sorts the container by using merge sort
func Stable(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	tempSlice := make([]interface{}, last.Position()-first.Position(), last.Position()-first.Position())
	if len(cmp) == 0 {
		mergeSort(first, last, comparator.BuiltinTypeComparator, tempSlice)
	} else {
		mergeSort(first, last, cmp[0], tempSlice)
	}
}

func mergeSort(first, last iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	if first.Position()+1 == last.Position() {
		return
	}
	mid := first.IteratorAt((first.Position() + last.Position()) >> 1)
	mergeSort(first, mid, cmp, tempSlice)
	mergeSort(mid, last, cmp, tempSlice)
	merge(first, mid, last, cmp, tempSlice)
}

func merge(first, mid, end iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	firstIter := (first.Clone()).(iterator.RandomAccessIterator)
	secondIter := (mid.Clone()).(iterator.RandomAccessIterator)
	pos := 0

	for firstIter.Position() < mid.Position() && secondIter.Position() < end.Position() {
		if cmp(firstIter.Value(), secondIter.Value()) <= 0 {
			tempSlice[pos] = firstIter.Value()
			pos++
			firstIter.Next()
		} else {
			tempSlice[pos] = secondIter.Value()
			pos++
			secondIter.Next()
		}
	}
	for ; firstIter.Position() < mid.Position(); firstIter.Next() {
		tempSlice[pos] = firstIter.Value()
		pos++
	}
	for ; secondIter.Position() < end.Position(); secondIter.Next() {
		tempSlice[pos] = secondIter.Value()
		pos++
	}

	iter := first.Clone().(iterator.RandomAccessIterator)
	for idx := 0; idx < pos; idx++ {
		iter.SetValue(tempSlice[idx])
		iter.Next()
	}
}
