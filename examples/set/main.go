package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/set"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	s := set.New[int](comparator.IntComparator, set.WithGoroutineSafe())
	s.Insert(1)
	s.Insert(5)
	s.Insert(3)
	s.Insert(4)
	s.Insert(2)

	s.Erase(4)

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v\n", iter.Value())
	}

	fmt.Printf("%v\n", s.Contains(3))
	fmt.Printf("%v\n", s.Contains(10))
}
