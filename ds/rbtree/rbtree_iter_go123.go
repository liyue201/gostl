//go:build go1.23

package rbtree

import "iter"

// Iter() provides a range-based immutable for loop over keys. Example:
//
//	for k := range t.Iter() {
//		fmt.Println(k)
//	}
func (t *RbTree[K, V]) Iter() iter.Seq[K] {
	return func(yield func(K) bool) {
		for n := t.First(); n != nil; n = n.Next() {
			if !yield(n.Key()) {
				break
			}
		}
	}
}

// Iter2() provides a range-based immutable for loop over key-value pairs. Example:
//
//	for k, v := range t.Iter2() {
//		fmt.Println(k, v)
//	}
func (t *RbTree[K, V]) Iter2() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for n := t.First(); n != nil; n = n.Next() {
			if !yield(n.Key(), n.Value()) {
				break
			}
		}
	}
}

// IterMut2() provides a range-based mutable for loop over key-value pairs. Example:
//
//	for k, vPtr := range t.IterMut2() {
//		*vPtr = newValue
//	}
func (t *RbTree[K, V]) IterMut2() iter.Seq2[K, *V] {
	return func(yield func(K, *V) bool) {
		for n := t.First(); n != nil; n = n.Next() {
			if !yield(n.Key(), &n.value) {
				break
			}
		}
	}
}
