package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSlice(t *testing.T) {
	a := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	sliceA := IntSlice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); !iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, 100)
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestIter(t *testing.T) {
	a := make([]float32, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, float32(i))
	}
	sliceA := Float32Slice(a)

	assert.EqualValues(t, 5, sliceA.Begin().IteratorAt(5).Value())

	i := float32(0)
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		assert.EqualValues(t, i, iter.Position())
		i++
		iter.SetValue(i * 10)
	}

	for iter := sliceA.Last().Clone().(*SliceIterator); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i*10, iter.Value())
		i--
	}
}
