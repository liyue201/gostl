package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"

	"math/rand"
)   
               
//sort the container by using quick sort
func Sort(begin, end iterator.RandomAccessIterator, cmp comparator.Comparator) {
	quickSort(begin, end, cmp)
}        
          
func quickSort(begin, end iterator.RandomAccessIterator, cmp comparator.Comparator) {
	if begin.Position()+1 >= end.Position() {
		return
	}     

	len := end.Position() - begin.Position()
	pos := rand.Int()%len + begin.Position()
	baseItem := begin.IteratorAt(pos)
	baseValue := baseItem.Value()
	if baseItem.Position() != begin.Position() {
		swapValue(baseItem, begin)
	}  

	leftIter := (begin.Clone().(iterator.RandomAccessIterator).Next()).(iterator.RandomAccessIterator)
	rightIter := (begin.Clone().(iterator.RandomAccessIterator)).IteratorAt(begin.Position() + len - 1)
	for leftIter.Position() < rightIter.Position() {
		leftCmp := cmp(leftIter.Value(), baseValue)
		if leftCmp <= 0 {
			leftIter.Next()
		} else {
			rightCmp := cmp(rightIter.Value(), baseValue)
			if rightCmp >= 0 {
				rightIter.Prev()
			} else {
				swapValue(leftIter, rightIter)
				leftIter.Next()
				if leftIter.Position() < rightIter.Position() {
					rightIter.Prev()
				}
			}
		}
	}

	if cmp(leftIter.Value(), begin.Value()) < 0 {
		swapValue(begin, leftIter)
		quickSort(begin, leftIter.Clone().(iterator.RandomAccessIterator), cmp)
		quickSort(leftIter.Clone().(iterator.RandomAccessIterator).Next().(iterator.RandomAccessIterator), end, cmp)
	}else {
		quickSort(begin, leftIter.Clone().(iterator.RandomAccessIterator), cmp)
		quickSort(leftIter.Clone().(iterator.RandomAccessIterator).(iterator.RandomAccessIterator), end, cmp)
	}
}

func swapValue(a, b iterator.RandomAccessIterator) {
	valA := a.Value()
	valB := b.Value()
	a.SetValue(valB)
	b.SetValue(valA)
}
