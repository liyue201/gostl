package bloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoomfilter(t *testing.T) {
	b := New(10000, 7, WithGoroutineSafe())

	assert.False(t, b.Contains("aa"))
	b.Add("aa")
	assert.True(t, b.Contains("aa"))

	other := NewFromData(b.Data(), WithGoroutineSafe())

	assert.True(t, other.Contains("aa"))

	b = NewWithEstimates(100000, 0.0001, WithGoroutineSafe())
	b.Add("bbbbb")
	assert.True(t, b.Contains("bbbbb"))
}
