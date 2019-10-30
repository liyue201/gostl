package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/ds/slice"
)

func main() {
	a := slice.IntSlice(make([]int, 0))
	for i := 5; i >= 1; i-- {
		a = append(a, i)
	}
	fmt.Printf("%v\n", a)
	sort.Sort(a.Begin(), a.End(), comparator.BuiltinTypeComparator)
	fmt.Printf("%v\n", a)
}
