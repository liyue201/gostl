package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := New(100)

	t.Logf("size : %v", bm.Size())
	for i := 0; i < 100; i++ {
		k := uint64(i)
		assert.Equal(t, false, bm.IsSet(k))
		bm.Set(k)
		assert.Equal(t, true, bm.IsSet(k))
	}

	for i := 0; i < 100; i++ {
		k := uint64(i)
		assert.Equal(t, true, bm.IsSet(k))
		bm.Unset(k)
		assert.Equal(t, false, bm.IsSet(k))
	}
	bm.Set(10)
	bm.Clear()
	assert.Equal(t, false, bm.IsSet(10))
}

func TestNewFromData(t *testing.T) {
	bm := New(100)

	bm.Set(6)
	bm.Set(20)
	bm.Set(77)

	bm2 := NewFromData(bm.Data())

	assert.Equal(t, bm.Size(), bm2.Size())

	for i := uint64(0); i < 100; i++ {
		assert.Equal(t, bm.IsSet(i), bm2.IsSet(i))
	}
}

func TestResize(t *testing.T) {
	bm := New(100)

	bm.Set(6)
	bm.Set(20)
	bm.Set(77)

	bm.Resize(1000)

	assert.Equal(t, true, bm.IsSet(6))
	assert.Equal(t, true, bm.IsSet(20))
	assert.Equal(t, true, bm.IsSet(77))

	bm.Resize(10)

	assert.Equal(t, true, bm.IsSet(6))
	assert.Equal(t, false, bm.IsSet(20))
	assert.Equal(t, false, bm.IsSet(77))
}
