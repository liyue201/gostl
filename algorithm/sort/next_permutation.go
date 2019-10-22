package sort

import (
	"github.com/liyue201/gostl/iterator"
	"github.com/liyue201/gostl/uitls/comparator"
)

func NextPermutation(begin, end iterator.SortableIterator, cmp comparator.Comparator) bool {
	len := end.Position() - begin.Position()
	endPos := begin.Position() + len - 1
	cur := endPos
	pre := cur - 1

	endIter := begin.IteratorAt(endPos)
	for cur > begin.Position() && cmp(begin.IteratorAt(pre).Value(), begin.IteratorAt(cur).Value()) >= 0 {
		cur--
		pre--
	}
	if cur <= begin.Position() {
		return false
	}
	for cur = endPos; cmp(begin.IteratorAt(cur).Value(),
		begin.IteratorAt(pre).Value()) <= 0; cur-- {
	}
	swapValue(begin.IteratorAt(cur), begin.IteratorAt(pre))
	reverse(begin.IteratorAt(pre+1), endIter)
	return true
}

func reverse(s, e iterator.SortableIterator) {
	for s.Position() < e.Position() {
		swapValue(s, e)
		s.Next()
		e.Prev()
	}
}
