package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)
 
//NextPermutation transform range [first last) to next permutation,return true if success, or false if failure
func NextPermutation(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) bool {
	if len(cmp) == 0 {
		return nextPermutation(first, last, comparator.BuiltinTypeComparator)
	} else {
		return nextPermutation(first, last, cmp[0])
	}
}

func nextPermutation(first, last iterator.RandomAccessIterator, cmp comparator.Comparator) bool {
	len := last.Position() - first.Position()
	endPos := first.Position() + len - 1
	cur := endPos
	pre := cur - 1

	endIter := first.IteratorAt(endPos)
	for cur > first.Position() && cmp(first.IteratorAt(pre).Value(), first.IteratorAt(cur).Value()) >= 0 {
		cur--
		pre--
	}
	if cur <= first.Position() {
		return false
	}
	for cur = endPos; cmp(first.IteratorAt(cur).Value(),
		first.IteratorAt(pre).Value()) <= 0; cur-- {
	}
	swapValue(first.IteratorAt(cur), first.IteratorAt(pre))
	reverse(first.IteratorAt(pre+1), endIter)
	return true
}

func reverse(s, e iterator.RandomAccessIterator) {
	for s.Position() < e.Position() {
		swapValue(s, e)
		s.Next()
		e.Prev()
	}
}
