package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//Stable sort the container by using merge sort
func Stable(begin, end iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	tempSlice := make([]interface{}, end.Position()-begin.Position(), end.Position()-begin.Position())
	if len(cmp) == 0 {
		mergeSort(begin, end, comparator.BuiltinTypeComparator, tempSlice)
	}else{
		mergeSort(begin, end, cmp[0], tempSlice)
	}
}

func mergeSort(begin, end iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	if begin.Position()+1 == end.Position() {
		return
	}
	mid := begin.IteratorAt((begin.Position() + end.Position()) >> 1)
	mergeSort(begin, mid, cmp, tempSlice)
	mergeSort(mid, end, cmp, tempSlice)
	merge(begin, mid, end, cmp, tempSlice)
}

func merge(begin, mid, end iterator.RandomAccessIterator, cmp comparator.Comparator, tempSlice []interface{}) {
	firstIter := (begin.Clone()).(iterator.RandomAccessIterator)
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

	iter := begin.Clone().(iterator.RandomAccessIterator)
	for idx := 0; idx < pos; idx++ {
		iter.SetValue(tempSlice[idx])
		iter.Next()
	}
}
