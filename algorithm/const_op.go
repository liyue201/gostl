package algorithm

import (
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/utils/iterator"
)

// Count returns the number of elements that their value is equal to value in range [first, last)
func Count(first, last iterator.ConstIterator, value interface{}, cmps ...comparator.Comparator) int {
	var count int
	cmp := comparator.BuiltinTypeComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if cmp(iter.Value(), value) == 0 {
			count++
		}
	}
	return count
}

// CountIf returns the number of elements are satisfied the function f in range [first, last)
func CountIf(first, last iterator.ConstIterator, f func(iterator.ConstIterator) bool) int {
	var count int
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if f(iter) {
			count++
		}
	}
	return count
}

// Find finds the first element that its value is equal to value in range [first, last), and returns its iterator, or last if not found
func Find(first, last iterator.ConstIterator, value interface{}, cmps ...comparator.Comparator) iterator.ConstIterator {
	cmp := comparator.BuiltinTypeComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if cmp(iter.Value(), value) == 0 {
			return iter
		}
	}
	return last
}

// FindIf finds the first element that is satisfied the function f, and returns its iterator, or last if not found
func FindIf(first, last iterator.ConstIterator, f func(iterator.ConstIterator) bool) iterator.ConstIterator {
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if f(iter) {
			return iter
		}
	}
	return last
}
