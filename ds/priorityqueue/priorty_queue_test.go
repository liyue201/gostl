package priorityqueue

import (
	. "github.com/liyue201/gostl/utils/comparator"
	"testing"
)

func TestMinPriorityQueue(t *testing.T) {
	pq := New()
	pq.Push(4)
	pq.Push(8)
	pq.Push(1)
	pq.Push(6)
	pq.Push(3)
	for !pq.Empty() {
		t.Logf("%v, %v", pq.Top(), pq.Pop())
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	pq := New(WithComparator(Reverse(BuiltinTypeComparator)), WithThreadSave())
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
	pq := New(WithComparator(Reverse(BuiltinTypeComparator)))
	pq.Push("fdsf")
	pq.Push("aavdsav")
	pq.Push("hrh42y5")
	pq.Push("u2ffaf")
	pq.Push("gqgfeg")
	for !pq.Empty() {
		t.Logf("%v, %v", pq.Top(), pq.Pop())
	}
}
