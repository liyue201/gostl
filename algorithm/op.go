package algorithm

import (
	"github.com/liyue201/gostl/utils/iterator"
)

// Swap swaps the value of two iterator
func Swap(a, b iterator.Iterator) {
	va := a.Value()
	vb := b.Value()
	a.SetValue(vb)
	b.SetValue(va)
}

// Reverse reverse the elements in the range [first, last]
func Reverse(first, last iterator.BidIterator) {
	left := first.Clone().(iterator.BidIterator)
	right := last.Clone().(iterator.BidIterator).Prev().(iterator.BidIterator)
	for !left.Equal(right) {
		Swap(left, right)
		left.Next()
		if left.Equal(right) {
			break
		}
		right.Prev()
	}
}
