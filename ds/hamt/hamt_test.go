package hamt

import (
	"fmt"
	"testing"
)

func TestHamt(t *testing.T) {
	h := New(WithThreadSafe())

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		h.Insert(Key(key), i)

		retVal := h.Get(Key(key))
		if retVal != i {
			t.Fatalf("Get %v error, value: %v", key, retVal)
		}
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		retVal := h.Get(Key(key))

		//t.Logf("%v = %v", key, retVal)

		if retVal != i {
			t.Fatalf("Get %v error", key)
		}
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%07d", i)
		if !h.Erase(Key(key)) {
			t.Fatalf("Erase %v error", key)
		}
		retVal := h.Get(Key(key))
		if retVal != nil {
			t.Fatalf("Get %v error, expect nil, but get %v", key, retVal)
		}
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("%08d", i)
		h.Insert([]byte(key), i)

		retVal := h.Get(Key(key))
		if retVal != i {
			t.Fatalf("Get %v error", key)
		}
	}
}
