package bitmap

import (
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := New(100)

	t.Logf("size : %v", bm.Size())
	for i := 0; i < 100; i++ {
		k := uint64(i)
		if bm.IsSet(k) {
			t.Fatalf("%d is not set", k)
		}
		bm.Set(k)
		if !bm.IsSet(k) {
			t.Fatalf("%d is set", k)
		}
	}

	for i := 0; i < 100; i++ {
		k := uint64(i)
		if !bm.IsSet(k) {
			t.Fatalf("%d is set", k)
		}
		bm.Unset(k)
		if bm.IsSet(k) {
			t.Fatalf("%d is unset", k)
		}
	}
}
