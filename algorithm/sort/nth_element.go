package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

func NthElement(first, last iterator.RandomAccessIterator, n int, cmps ...comparator.Comparator) {
	if first.Position() < 0 || last.Position()-first.Position() < n {
		return
	}
	cmp := comparator.BuiltinTypeComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	randSeed := getSeed()
	nthElement(first, last, n, cmp, randSeed)
}

func nthElement(first, last iterator.RandomAccessIterator, n int, cmp comparator.Comparator, randSeed uint64) {
	if first.Position()+1 >= last.Position() {
		return
	}
	randNum := randInt64(randSeed)
	len := last.Position() - first.Position()
	pos := int(randNum%uint64(len)) + first.Position()
	baseItem := first.IteratorAt(pos)
	baseValue := baseItem.Value()
	if baseItem.Position() != first.Position() {
		swapValue(baseItem, first)
	}

	leftIter := (first.Clone().(iterator.RandomAccessIterator).Next()).(iterator.RandomAccessIterator)
	rightIter := (first.Clone().(iterator.RandomAccessIterator)).IteratorAt(first.Position() + len - 1)
	for leftIter.Position() <= rightIter.Position() {
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
				rightIter.Prev()
			}
		}
	}
	leftIter.Prev()
	m := leftIter.Position()

	if cmp(leftIter.Value(), first.Value()) < 0 {
		swapValue(first, leftIter)
	}

	if n <= m-first.Position() {
		nthElement(first, first.IteratorAt(m), n, cmp, randNum+uint64(pos))
	} else {
		nthElement(first.IteratorAt(m).Next().(iterator.RandomAccessIterator), last, n-(m-first.Position()+1), cmp, randNum+uint64(pos+1))
	}
}
