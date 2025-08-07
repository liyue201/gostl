//go:build go1.23

package treemap

import "iter"

// Iter() provides a range-based immutable for loop over keys. Example:
//
//	for k := range m.Iter() {
//		fmt.Println(k)
//	}
func (m *Map[K, V]) Iter() iter.Seq[K] {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.tree.Iter()
}

// Iter2() provides a range-based immutable for loop over key-value pairs. Example:
//
//	for k, v := range m.Iter2() {
//		fmt.Println(k, v)
//	}
func (m *Map[K, V]) Iter2() iter.Seq2[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.tree.Iter2()
}

// IterMut2() provides a range-based mutable for loop over key-value pairs. Example:
//
//	for k, vPtr := range m.IterMut2() {
//		*vPtr = newValue
//	}
func (m *Map[K, V]) IterMut2() iter.Seq2[K, *V] {
	m.locker.Lock()
	defer m.locker.Unlock()
	return m.tree.IterMut2()
}

// Iter() provides a range-based immutable for loop over keys. Example:
//
//	for k := range m.Iter() {
//		fmt.Println(k)
//	}
func (m *MultiMap[K, V]) Iter() iter.Seq[K] {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.tree.Iter()
}

// Iter2() provides a range-based immutable for loop over key-value pairs. Example:
//
//	for k, v := range m.Iter2() {
//		fmt.Println(k, v)
//	}
func (m *MultiMap[K, V]) Iter2() iter.Seq2[K, V] {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.tree.Iter2()
}

// IterMut2() provides a range-based mutable for loop over key-value pairs. Example:
//
//	for k, vPtr := range m.IterMut2() {
//		*vPtr = newValue
//	}
func (m *MultiMap[K, V]) IterMut2() iter.Seq2[K, *V] {
	m.locker.Lock()
	defer m.locker.Unlock()
	return m.tree.IterMut2()
}
