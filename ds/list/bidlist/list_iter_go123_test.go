//go:build go1.23

package bidlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	i := 1
	for v := range list.Iter() {
		assert.Equal(t, i, v)
		i++
	}

	for v := range list.IterMut() {
		*v *= 2
	}

	i = 1
	for v := range list.Iter() {
		assert.Equal(t, i*2, v)
		i++
	}
}