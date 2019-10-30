package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/utils/comparator"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	q := deque.New()
	q.PushBack(2)
	q.PushFront(1)
	q.PushBack(3)
	q.PushFront(4)
	fmt.Printf("%v\n", q)

	sort.Sort(q.Begin(), q.End(), comparator.BuiltinTypeComparator)
	fmt.Printf("%v\n", q)
}
