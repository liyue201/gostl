package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
	"math/rand"
)

func main() {
	q := deque.New[int]()
	for i := 0; i < 100; i++ {
		r := rand.Int() % 100
		q.PushBack(r)
		q.PushFront(r)
	}
	fmt.Printf("%v\n", q)

	sort.Sort[int](q.Begin(), q.End())
	fmt.Printf("%v\n", q)

	for !q.Empty() {
		r := rand.Int() % q.Size()
		q.EraseAt(r)
	}
	fmt.Printf("%v\n", q)
}
