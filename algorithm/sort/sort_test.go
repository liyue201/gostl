package sort

import (
	"github.com/liyue201/gostl/ds/vector"
	"github.com/liyue201/gostl/utils/comparator"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	// test size = 0
	rand.Seed(time.Now().UnixNano())
	v := vector.New[int]()
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)
	t.Logf("v: %v", v.String())

	// test size = 1
	v = vector.New[int]()
	for i := 0; i < 1; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}

	// test size = 2
	v = vector.New[int]()
	for i := 0; i < 2; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}

	// test size = 10
	v = vector.New[int]()
	for i := 0; i < 10; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}

func TestSort2(t *testing.T) {
	// test size = 31
	v := vector.New[int]()
	for i := 0; i < 31; i++ {
		v.PushBack(rand.Int() % 10)
	}
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}

	// test size = 50
	v = vector.New[int]()
	for i := 0; i < 50; i++ {
		v.PushBack(rand.Int() % 100)
	}
	Sort[int](v.Begin(), v.End(), comparator.IntComparator)

	t.Logf("v: %v", v.String())
	for i := 0; i < v.Size()-1; i++ {
		if v.At(i) > v.At(i+1) {
			t.Fatalf("sort vector error")
		}
	}
}
