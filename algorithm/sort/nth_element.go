package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// NthElement Rearranges the elements in the range [first,last), in such a way that the element at the nth position is the element that would be in that position in a sorted sequence
func NthElement[T any](first, last iterator.RandomAccessIterator[T], n int, cmps ...comparator.Comparator) {
	if first.Position() < 0 || last.Position()-first.Position() < n {
		return
	}
	cmp := comparator.BuiltinTypeComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	len := last.Position() - first.Position()
	nthElement(first, last.IteratorAt(first.Position()+len-1), n, cmp)
}

func nthElement[T any](first, last iterator.RandomAccessIterator[T], n int, cmp comparator.Comparator) {
	if first.Position()+1 >= last.Position() {
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

	if n <= m-first.Position() {
		nthElement(first, first.IteratorAt(m), n, cmp)
	} else {
		nthElement(first.IteratorAt(m).Next().(iterator.RandomAccessIterator[T]), last, n-(m-first.Position()+1), cmp)
	}
}
