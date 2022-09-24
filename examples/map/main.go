package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/map"
	"github.com/liyue201/gostl/utils/comparator"
)

func main() {
	m := treemap.New[string, string](comparator.StringComparator, treemap.WithGoroutineSafe())

	m.Insert("a", "aaa")
	m.Insert("b", "bbb")

	a, _ := m.Get("a")
	b, _ := m.Get("b")
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

	m.Erase("b")
}
