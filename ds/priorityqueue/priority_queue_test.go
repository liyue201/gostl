package priorityqueue

import (
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
)

func TestMinPriorityQueue(t *testing.T) {
	pq := New(comparator.IntComparator)
	pq.Push(4)
	pq.Push(8)
	pq.Push(1)
	pq.Push(6)
	pq.Push(3)
	for !pq.Empty() {
		t.Logf("%v, %v", pq.Top(), pq.Pop())
	}
}

func TestMinPriorityQueueWithGoroutineSafe(t *testing.T) {
	pq := New(comparator.IntComparator, WithGoroutineSafe())
	pq.Push(4)
	pq.Push(8)
	pq.Push(1)
	pq.Push(6)
	pq.Push(3)
	for !pq.Empty() {
		t.Logf("%v,%v", pq.Top(), pq.Pop())
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	pq := New(comparator.Reverse(comparator.IntComparator))
	pq.Push(4)
	pq.Push(8)
	pq.Push(1)
	pq.Push(6)
	pq.Push(3)
	for !pq.Empty() {
		t.Logf("%v, %v", pq.Top(), pq.Pop())
	}
}

func TestMaxPriorityQueueWithGoroutineSafe(t *testing.T) {
	pq := New(comparator.Reverse(comparator.IntComparator), WithGoroutineSafe())
	pq.Push(4)
	pq.Push(8)
	pq.Push(1)
	pq.Push(6)
	pq.Push(3)
	for !pq.Empty() {
		t.Logf("%v,%v", pq.Top(), pq.Pop())
	}
}

func TestStringPriorityQueue(t *testing.T) {
	pq := New(comparator.Reverse(comparator.StringComparator))
	pq.Push("fdsf")
	pq.Push("aavdsav")
	pq.Push("hrh42y5")
	pq.Push("u2ffaf")
	pq.Push("gqgfeg")
	for !pq.Empty() {
		t.Logf("%v, %v", pq.Top(), pq.Pop())
	}
}
