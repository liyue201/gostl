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

	sliceA := slice.NewSliceWrapper(a)

	////Sort in ascending order
	sort.Sort[string](sliceA.Begin(), sliceA.End(), comparator.OrderedTypeCmp[string])

	sort.Stable[string](sliceA.Begin(), sliceA.End(), comparator.StringComparator)
	fmt.Printf("%v\n", a)

	if sort.BinarySearch[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator) {
		fmt.Printf("BinarySearch: found bbbb\n")
	}

	iter := sort.LowerBound[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator)
	if iter.IsValid() {
		fmt.Printf("LowerBound bbbb: %v\n", iter.Value())
	}
	iter = sort.UpperBound[string](sliceA.Begin(), sliceA.End(), "bbbb", comparator.StringComparator)
	if iter.IsValid() {
		fmt.Printf("UpperBound bbbb: %v\n", iter.Value())
	}
	//Sort in descending order
	sort.Sort[string](sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.StringComparator))
	//sort.Stable[string](sliceA.Begin(), sliceA.End(), comparator.Reverse(comparator.StringComparator))
	fmt.Printf("%v\n", a)
}
