package main

import (
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/ds/map"
	"fmt"
)

func main()  {
	m := treemap.New(comparator.BuiltinTypeComparator)

	m.Insert("a", "aaa")
	m.Insert("b", "bbb")

	fmt.Printf("a = %v\n", m.Get("a"))
	fmt.Printf("b = %v\n", m.Get("b"))

}
