package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm"
	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/utils/iterator"
)

func isEven(iter iterator.ConstIterator) bool {
	return iter.Value().(int)%2 == 0
}

func greaterThan5(iter iterator.ConstIterator) bool {
	return iter.Value().(int) > 5
}

func main() {
	a := deque.New()
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	for i := 0; i < 5; i++ {
		a.PushBack(i)
	}
	fmt.Printf("%v\n", a)

	fmt.Printf("Count 2: %v\n", algorithm.Count(a.Begin(), a.End(), 2))
	fmt.Printf("Count 2: %v\n", algorithm.CountIf(a.Begin(), a.End(), isEven))

	iter := algorithm.Find(a.Begin(), a.End(), 2)
	if !iter.Equal(a.End()) {
		fmt.Printf("Fund %v\n", iter.Value())
	}
	iter = algorithm.FindIf(a.Begin(), a.End(), greaterThan5)
	if !iter.Equal(a.End()) {
		fmt.Printf("FindIf greaterThan5 : %v\n", iter.Value())
	}
	iter = algorithm.MaxElement(a.Begin(), a.End())
	if !iter.Equal(a.End()) {
		fmt.Printf("largest value : %v\n", iter.Value())
	}
	iter = algorithm.MinElement(a.Begin(), a.End())
	if !iter.Equal(a.End()) {
		fmt.Printf("largest value : %v\n", iter.Value())
	}
}
