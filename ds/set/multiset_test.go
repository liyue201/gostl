package set

import (
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
	"github.com/stretchr/testify/assert"
)

func TestMultiSet(t *testing.T) {
	mset := NewMultiSet(comparator.IntComparator, WithGoroutineSafe())

	mset.Insert(1)
	mset.Insert(5)
	mset.Insert(1)

	assert.Equal(t, 3, mset.Size())
	assert.True(t, mset.Contains(1))
	assert.Equal(t, 1, mset.Begin().Value())
	assert.Equal(t, 5, mset.Last().Value())

	for iter := mset.First(); iter.IsValid(); iter.Next() {
		t.Logf("%v\n", iter.Value())
	}
	assert.Equal(t, "[1 1 5]", mset.String())
	mset.EraseAll(1)
	assert.Equal(t, "[5]", mset.String())

	mset.Insert(5)
	assert.Equal(t, 2, mset.Count(5))
	assert.Equal(t, "[5 5]", mset.String())
	mset.Erase(5)
	assert.Equal(t, 1, mset.Count(5))
	assert.Equal(t, "[5]", mset.String())

	mset.Clear()
	assert.Equal(t, 0, mset.Size())

	mset.Insert(2)
	mset.Insert(7)
	mset.Insert(8)
	mset.Insert(3)
	assert.True(t, mset.Contains(7))

	iter := mset.LowerBound(5)
	assert.Equal(t, 7, iter.Value())

	iter = mset.UpperBound(3)
	assert.Equal(t, 7, iter.Value())

	iter = mset.Find(3)
	assert.Equal(t, 3, iter.Value())
}
