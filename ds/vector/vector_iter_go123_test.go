//go:build go1.23

package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorIter(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	v.PushBack(4)
	//[1 2 3 4]

	i := 0
	for v := range v.Iter() {
		assert.Equal(t, i+1, v)
		i++
	}

	// mut
	for vPtr := range v.IterMut() {
		*vPtr *= 2
	}

	i = 0
	for v := range v.Iter() {
		assert.Equal(t, (i+1)*2, v)
		i++
	}
}
