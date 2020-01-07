package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	for i := 10; i >= 1; i-- {
		assert.False(t, s.Contains(i))
		s.Insert(i)
		assert.True(t, s.Contains(i))
	}

	i := 1
	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		assert.Equal(t, i, iter.Value())
		i++
	}

	i = 10
	for iter := s.Last(); iter.IsValid(); iter.Prev() {
		assert.Equal(t, i, iter.Value())
		i--
	}
	s.Erase(5)
	assert.False(t, s.Contains(5))
	assert.Equal(t, 9, s.Size())

	iter := s.LowerBound(3)
	assert.Equal(t, 3, iter.Value())

	s.Erase(3)
	iter = s.LowerBound(3)
	assert.Equal(t, 4, iter.Value())
}

func TestSet_Cal(t *testing.T) {
	s1 := New()
	s2 := New()
	for i := 1; i <= 5; i++ {
		s1.Insert(i)     // [1 2 3 4 5]
		s2.Insert(i + 3) // [4 5 6 7 8]
	}

	assert.Equal(t, "[4 5]", s1.Intersect(s2).String())
	assert.Equal(t, "[1 2 3 4 5 6 7 8]", s1.Union(s2).String())
	assert.Equal(t, "[1 2 3]", s1.Diff(s2).String())
	assert.Equal(t, "[6 7 8]", s2.Diff(s1).String())
}

func TestSet_Erase(t *testing.T) {
	s := New()
	for i := 0; i < 1000; i++ {
		s.Insert(i)
	}
	for i := 0; i < 1000; i++ {
		s.Erase(i)
	}
	assert.Equal(t, 0, s.Size())
}
