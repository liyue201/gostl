package main

import (
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/containers/set"
	"fmt"
)

func main()  {
	s := set.New(comparator.BuiltinTypeComparator)
	s.Insert(1)
	s.Insert(5)
	s.Insert(3)
	s.Insert(4)
	s.Insert(2)

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v\n", iter.Value())
	}
}
