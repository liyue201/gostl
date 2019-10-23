package main

import (
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/containers/slice"
	"fmt"
)

func main()  {
	a := make([]string, 0)
	a = append(a, "bbbb")
	a = append(a, "ccc")
	a = append(a, "aaaa")
	a = append(a, "bb")

	sliceA := slice.StringSlice(a)

	////Sort in ascending order
	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.BuiltinTypeComparator)
	fmt.Printf("%v\n", a)

	//Sort in descending order
	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v\n", a)
}