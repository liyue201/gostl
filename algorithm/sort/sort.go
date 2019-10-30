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
func Sort(begin, end iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	randSeed := getSeed()
	if len(cmp) == 0 {
		quickSort(begin, end, comparator.BuiltinTypeComparator, randSeed)
	} else {
		quickSort(begin, end, cmp[0], randSeed)
	}
}

func quickSort(begin, end iterator.RandomAccessIterator, cmp comparator.Comparator, randSeed uint64) {
	if begin.Position()+1 >= end.Position() {
		return
	}
	randNum := randInt64(randSeed)
	len := end.Position() - begin.Position()
	pos := int(randNum%uint64(len)) + begin.Position()
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
		quickSort(begin, leftIter.Clone().(iterator.RandomAccessIterator), cmp, randNum+uint64(pos))
		quickSort(leftIter.Clone().(iterator.RandomAccessIterator).Next().(iterator.RandomAccessIterator), end, cmp, randNum+uint64(pos+1))
	} else {
		quickSort(begin, leftIter.Clone().(iterator.RandomAccessIterator), cmp, randNum+uint64(pos))
		quickSort(leftIter.Clone().(iterator.RandomAccessIterator).(iterator.RandomAccessIterator), end, cmp, randNum+uint64(pos+1))
	}
}

func swapValue(a, b iterator.RandomAccessIterator) {
	valA := a.Value()
	valB := b.Value()
	a.SetValue(valB)
	b.SetValue(valA)
}
