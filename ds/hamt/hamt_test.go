package hamt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHamt(t *testing.T) {
	h := New[int](WithGoroutineSafe())

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		h.Insert(Key(key), i)
		v, _ := h.Get(Key(key))
		assert.Equal(t, i, v)
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		v, _ := h.Get(Key(key))
		assert.Equal(t, i, v)
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		assert.True(t, h.Erase(Key(key)))
		_, err := h.Get(Key(key))
		assert.Equal(t, err, ErrorNotFound)
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%08d", i)
		h.Insert([]byte(key), i)
		v, _ := h.Get(Key(key))
		assert.Equal(t, i, v)
	}
}

func TestTraversal(t *testing.T) {
	h := New[string]()
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
