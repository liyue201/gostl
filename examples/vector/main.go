package main

import (
	"fmt"
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/comparator"
	"github.com/liyue201/gostl/ds/vector"
)

func main() {
	v := vector.New(1)
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	for i := 0; i < v.Size(); i++ {
		fmt.Printf("%v\n", v.At(i))
	}
	//reverse sort
	sort.Sort(v.Begin(), v.End(), comparator.Reverse(comparator.BuiltinTypeComparator))
	fmt.Printf("%v", v)
}
