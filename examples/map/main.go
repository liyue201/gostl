package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/map"
)

func main() {
	m := treemap.New()

	m.Insert("a", "aaa")
	m.Insert("b", "bbb")

	fmt.Printf("a = %v\n", m.Get("a"))
	fmt.Printf("b = %v\n", m.Get("b"))

}
