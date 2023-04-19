package pair

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakePair(t *testing.T) {
	p := MakePair(1.1, 's')
	assert.Equal(t, 1.1, p.Front)
	assert.Equal(t, 's', p.Back)
	p = MakePair("ABC", 123)
	assert.Equal(t, "ABC", p.Front)
	assert.Equal(t, 123, p.Back)
}

func TestPair_New(t *testing.T) {
	var p Pair
	p.New(1.1, 's')
	assert.Equal(t, 1.1, p.Front)
	assert.Equal(t, 's', p.Back)
	p.New("ABC", 123)
	assert.Equal(t, "ABC", p.Front)
	assert.Equal(t, 123, p.Back)
}

func TestPair_Equal(t *testing.T) {
	p := MakePair(1.1, 's')
	q := MakePair("ABC", 123)
	assert.False(t, p.Equal(*q))
	q.Front = 1.1
	q.Back = 's'
	assert.True(t, p.Equal(*q))
}

func TestPair_Fronts(t *testing.T) {
	p := MakePair(1.1, 's')
	q := MakePair("ABC", 123)
	assert.Equal(t, 1.1, p.Fronts())
	assert.Equal(t, "ABC", q.Fronts())
}

func TestPair_Backs(t *testing.T) {
	p := MakePair(1.1, 's')
	q := MakePair("ABC", 123)
	assert.Equal(t, 's', p.Backs())
	assert.Equal(t, 123, q.Backs())
}
