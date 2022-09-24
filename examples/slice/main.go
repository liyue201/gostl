package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	a := make([]int, 0)
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 3)
	fmt.Printf("%v\n", a)

	wa := slice.NewSliceWrapper(a)

	// sort in ascending
	sort.Sort[int](wa.Begin(), wa.End(), comparator.IntComparator)
	fmt.Printf("%v\n", a)

	// sort in descending
	sort.Sort[int](wa.Begin(), wa.End(), comparator.Reverse(comparator.IntComparator))
	fmt.Printf("%v\n", a)
}
