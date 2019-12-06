package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	a := make([]string, 0)
	a = append(a, "bbbb")
	a = append(a, "ccc")
	a = append(a, "aaaa")
	a = append(a, "bbbb")
	a = append(a, "bb")

	sliceA := slice.StringSlice(a)

	////Sort in ascending order
	sort.Sort(sliceA.Begin(), sliceA.End())
	//sort.Stable(sliceA.Begin(), sliceA.End())
	fmt.Printf("%v\n", a)

	if sort.BinarySearch(sliceA.Begin(), sliceA.End(), "bbbb") {
		fmt.Printf("BinarySearch: found bbbb\n")
	}

	iter := sort.LowerBound(sliceA.Begin(), sliceA.End(), "bbbb")
	if iter.IsValid() {
		fmt.Printf("LowerBound bbbb: %v\n", iter.Value())
	}
	iter = sort.UpperBound(sliceA.Begin(), sliceA.End(), "bbbb")
	if iter.IsValid() {
		fmt.Printf("UpperBound bbbb: %v\n", iter.Value())
	}
	//Sort in descending order
	sort.Sort(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	//sort.Stable(sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v\n", a)

}
