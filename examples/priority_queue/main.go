package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/priority_queue"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	q := priority_queue.New(priority_queue.WithComparator(comparator.Reverse(comparator.BuiltinTypeComparator)),
		priority_queue.WithThreadSave())
	q.Push(5)
	q.Push(13)
	q.Push(7)
	q.Push(9)
	q.Push(0)
	q.Push(88)
	for !q.Empty() {
		fmt.Printf("%v\n", q.Pop())
	}
}
