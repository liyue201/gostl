package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/deque"
)

func main() {
	a := deque.New[int]()
	a.PushBack(9)
	a.PushBack(8)
	a.PushBack(7)
	a.PushBack(6)
	a.PushBack(5)
	a.PushBack(4)
	a.PushBack(3)
	a.PushBack(2)
	a.PushBack(1)
	fmt.Printf("%v\n", a)
	sort.NthElement[int](a.Begin(), a.End(), 3)
	fmt.Printf("%v\n", a.At(3))
	fmt.Printf("%v\n", a)
}
