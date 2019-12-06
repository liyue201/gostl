package sort

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
	"math/rand"
	"time"
)

var randTable [256]uint64
var seedChan chan uint64

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seedChan = make(chan uint64, 256)

	for i := 0; i < len(randTable); i++ {
		randTable[i] = r.Uint64()
		seedChan <- r.Uint64()
	}
}

func getSeed() uint64 {
	id := <-seedChan
	seedChan <- id + 78431
	return id
}

func randInt64(seed uint64) uint64 {
	return (randTable[seed&0xff] & 0xffff) | (randTable[seed>>16&0xff] & 0xffff0000) | (randTable[seed>>32&0xff] & 0xffff00000000) | (randTable[seed>>48&0xff] & 0xffff000000000000)
}

//sort the container by using quick sort
func Sort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	randSeed := getSeed()
	if len(cmp) == 0 {
		quickSort(first, last, comparator.BuiltinTypeComparator, randSeed)
	} else {
		quickSort(first, last, cmp[0], randSeed)
	}
}

func quickSort(first, last iterator.RandomAccessIterator, cmp comparator.Comparator, randSeed uint64) {
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

	quickSort(first, first.IteratorAt(m), cmp, randNum+uint64(pos))
	quickSort(first.IteratorAt(m).Next().(iterator.RandomAccessIterator), last, cmp, randNum+uint64(pos+1))
}

func swapValue(a, b iterator.RandomAccessIterator) {
	valA := a.Value()
	valB := b.Value()
	a.SetValue(valB)
	b.SetValue(valA)
}
