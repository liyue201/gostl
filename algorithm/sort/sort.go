package sort

import (
	"fmt"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

//Sort sorts the container by using quick sort
func Sort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	fmt.Printf("============================\n")
	n := last.Position() - first.Position()
	if len(cmp) == 0 {
		quickSort(first, first.IteratorAt(first.Position()+n-1), comparator.BuiltinTypeComparator)
	} else {
		quickSort(first, first.IteratorAt(first.Position()+n-1), cmp[0])
	}
}

func quickSort(first, last iterator.RandomAccessIterator, cmp comparator.Comparator) {
	if first.Position() >= last.Position() {
		return
	}
	len := last.Position() - first.Position() + 1
	if len < 3 {
		if cmp(first.Value(), last.Value()) > 0 {
			swapValue(first, last)
		}
		return
	}

	mid := first.IteratorAt(first.Position() + len/2)
	doPivot(first, mid, last, cmp)
	swapValue(mid, first.IteratorAt(last.Position()-1))
	if len == 3 {
		return
	}
	baseItem := first.IteratorAt(last.Position() - 1)

	leftIter := first.IteratorAt(first.Position() + 1)
	rightIter := first.IteratorAt(last.Position() - 2)
	for leftIter.Position() <= rightIter.Position() {
		leftCmp := cmp(leftIter.Value(), baseItem.Value())
		if leftCmp <= 0 {
			leftIter.Next()
		} else {
			rightCmp := cmp(rightIter.Value(), baseItem.Value())
			if rightCmp > 0 {
				rightIter.Prev()
			} else {
				swapValue(leftIter, rightIter)
				leftIter.Next()
				rightIter.Prev()
			}
		}
	}
	rightIter.Next()
	m := rightIter.Position()
	swapValue(baseItem, rightIter)

	quickSort(first, first.IteratorAt(m-1), cmp)
	quickSort(first.IteratorAt(m).Next().(iterator.RandomAccessIterator), last, cmp)
}

func doPivot(first, mid, last iterator.RandomAccessIterator, cmp comparator.Comparator) {
	if cmp(first.Value(), mid.Value()) > 0 {
		swapValue(first, mid)
	}
	if cmp(first.Value(), last.Value()) > 0 {
		swapValue(first, last)
	}
	if cmp(mid.Value(), last.Value()) > 0 {
		swapValue(mid, last)
	}
}

func swapValue(a, b iterator.RandomAccessIterator) {
	valA := a.Value()
	valB := b.Value()
	a.SetValue(valB)
	b.SetValue(valA)
}
