package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	a := make([]int, 0)
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	wa := slice.NewSliceWrapper(a)
	fmt.Println("NextPermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation[int](wa.Begin(), wa.End()) {
			break
		}
	}
	fmt.Println("PrePermutation")
	for {
		fmt.Printf("%v\n", a)
		if !sort.NextPermutation[int](wa.Begin(), wa.End(), comparator.Reverse(comparator.BuiltinTypeComparator)) {
			break
		}
	}
}
