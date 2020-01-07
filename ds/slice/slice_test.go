package slice

import (
	"fmt"
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
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
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

func TestUIntSlice(t *testing.T) {
	a := make([]uint, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, uint(i))
	}
	sliceA := UIntSlice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, uint(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestInt8Slice(t *testing.T) {
	a := make([]int8, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, int8(i))
	}
	sliceA := Int8Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, int8(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestUInt8Slice(t *testing.T) {
	a := make([]uint8, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, uint8(i))
	}
	sliceA := UInt8Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, uint8(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestInt16Slice(t *testing.T) {
	a := make([]int16, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, int16(i))
	}
	sliceA := Int16Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, int16(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestUInt16Slice(t *testing.T) {
	a := make([]uint16, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, uint16(i))
	}
	sliceA := UInt16Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, uint16(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestInt32Slice(t *testing.T) {
	a := make([]int32, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, int32(i))
	}
	sliceA := Int32Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, int32(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestUInt32Slice(t *testing.T) {
	a := make([]uint32, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, uint32(i))
	}
	sliceA := UInt32Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, uint32(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestInt64Slice(t *testing.T) {
	a := make([]int64, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, int64(i))
	}
	sliceA := Int64Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, int64(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestUInt64Slice(t *testing.T) {
	a := make([]uint64, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, uint64(i))
	}
	sliceA := UInt64Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, uint64(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestFloat32Slice(t *testing.T) {
	a := make([]float32, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, float32(i))
	}
	sliceA := Float32Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, float32(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestFloat64Slice(t *testing.T) {
	a := make([]float64, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, float64(i))
	}
	sliceA := Float64Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, float64(100))
	assert.EqualValues(t, 100, sliceA.At(8))
}

func TestStringSlice(t *testing.T) {
	a := make([]string, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	sliceA := StringSlice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, "5", sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, fmt.Sprintf("%v", i), iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, fmt.Sprintf("%v", i), iter.Value())
		i--
	}
	sliceA.Set(8, "100")
	assert.EqualValues(t, "100", sliceA.At(8))
}

func TestSlice(t *testing.T) {
	a := make([]interface{}, 0, 10)

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	sliceA := Slice(a)

	assert.Equal(t, 10, sliceA.Len())
	assert.EqualValues(t, 5, sliceA.At(5))

	i := 0
	for iter := sliceA.Begin(); !iter.Equal(sliceA.End()); iter.Next() {
		assert.EqualValues(t, i, iter.Value())
		i++
	}

	i = 9
	for iter := sliceA.Last(); iter.IsValid(); iter.Prev() {
		assert.EqualValues(t, i, iter.Value())
		i--
	}
	sliceA.Set(8, 100)
	assert.EqualValues(t, 100, sliceA.At(8))
}
