package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/containers/slice"
)

func main() {
	a := make([]int, 0)
	for i := 5; i >= 1; i-- {
		a = append(a, i)
	}
	//Sort in ascending order
	sliceA := slice.IntSlice(a)
	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.BuiltinTypeComparator)
	fmt.Printf("%v\n", a)
}
