package array

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	a := New[int](10)
	assert.Equal(t, 10, a.Size())
	assert.Equal(t, false, a.Empty())

	va := 10
	a.Fill(va)
	for i := 0; i < a.Size(); i++ {
		val := a.At(i)
		assert.Equal(t, va, val)
	}

	b := New[int](10)
	vb := 66
	b.Fill(vb)
	a.SwapArray(b)

	for i := 0; i < a.Size(); i++ {
		assert.Equal(t, vb, a.At(i))
		assert.Equal(t, va, b.At(i))
	}

	for i := 0; i < a.Size(); i++ {
		a.Set(i, i)
	}

	t.Logf("Traversal a:")
	i := 0
	for iter := a.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v ", iter.Value())
		assert.Equal(t, i, iter.Value())
		i++
	}

	t.Logf(" Reverse traversal a:")
	i = a.Size() - 1
	for iter := a.Last(); iter.IsValid(); iter.Next() {
		t.Logf("%v ", iter.Value())
		assert.Equal(t, i, iter.Value())
		i--
	}
}

func TestNewFromArray(t *testing.T) {
	a := New[int](10)
	for i := 0; i < 10; i++ {
		a.Set(i, i*10)
	}
	b := NewFromArray(a)

	assert.Equal(t, a.Size(), b.Size())
	for i := 0; i < 10; i++ {
		assert.Equal(t, a.At(i), b.At(i))
	}
}

func TestSort(t *testing.T) {
	a := New[int](10)
	for i := 0; i < 10; i++ {
		a.Set(i, 10-i)
	}
	sort.Stable[int](a.Begin(), a.End(), comparator.IntComparator)
	t.Logf("a: %v", a.String())
	for i := 0; i < a.Size()-1; i++ {
		assert.LessOrEqual(t, a.At(i), a.At(i))
	}
}

func TestIterator(t *testing.T) {
	a := New[int](10)
	n := 0
	for iter := a.Begin(); iter.IsValid(); iter.Next() {
		iter.SetValue(n)
		n++
	}
	assert.Equal(t, 2, a.At(2))
	assert.Equal(t, 0, a.Front())
	assert.Equal(t, 9, a.Back())

	n = 0
	for iter := a.Begin(); !iter.Equal(a.End()); iter.Next() {
		assert.Equal(t, n, iter.Value())
		n++
	}

	n = 9
	for iter := a.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, n, iter.Value())
		n--
	}

}
