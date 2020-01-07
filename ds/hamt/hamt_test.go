package hamt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHamt(t *testing.T) {
	h := New(WithThreadSafe())

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		h.Insert(Key(key), i)
		assert.Equal(t, i, h.Get(Key(key)))
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		assert.Equal(t, i, h.Get(Key(key)))
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		assert.True(t, h.Erase(Key(key)))
		assert.Equal(t, nil, h.Get(Key(key)))
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%08d", i)
		h.Insert([]byte(key), i)
		assert.Equal(t, i, h.Get(Key(key)))
	}
}

func TestTraversal(t *testing.T) {
	h := New()
	m := make(map[string]string)

	h.Insert(Key("222"), "bbb")
	h.Insert(Key("111"), "aaa")
	h.Insert(Key("333"), "ccc")
	m["111"] = "aaa"
	m["222"] = "bbb"
	m["333"] = "ccc"
	keys := h.Keys()

	strKeys := h.StringKeys()

	for i := 0; i < len(keys); i++ {
		assert.Equal(t, string(keys[i]), strKeys[i])
	}

	sort.Strings(strKeys)

	assert.Equal(t, "111", strKeys[0])
	assert.Equal(t, "222", strKeys[1])
	assert.Equal(t, "333", strKeys[2])

	h.Traversal(func(key, value interface{}) bool {
		val := m[string(key.(Key))]
		assert.Equal(t, val, value.(string))
		return true
	})
}
