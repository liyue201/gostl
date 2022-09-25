package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//Stable sorts the container by using merge sort
func Stable[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T]) {
	tempSlice := make([]T, last.Position()-first.Position(), last.Position()-first.Position())
	mergeSort(first, last, cmp, tempSlice)

}

func mergeSort[T any](first, last iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	if first.Position()+1 == last.Position() {
		return
	}
	mid := first.IteratorAt((first.Position() + last.Position()) >> 1)
	mergeSort(first, mid, cmp, tempSlice)
	mergeSort(mid, last, cmp, tempSlice)
	merge(first, mid, last, cmp, tempSlice)
}

func merge[T any](first, mid, end iterator.RandomAccessIterator[T], cmp comparator.Comparator[T], tempSlice []T) {
	firstIter := (first.Clone()).(iterator.RandomAccessIterator[T])
	secondIter := (mid.Clone()).(iterator.RandomAccessIterator[T])
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

	iter := first.Clone().(iterator.RandomAccessIterator[T])
	for idx := 0; idx < pos; idx++ {
		iter.SetValue(tempSlice[idx])
		iter.Next()
	}
}
