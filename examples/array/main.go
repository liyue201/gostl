package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/array"
)

func main() {
	a := array.New(5)
	for i := 0; i < a.Size(); i++ {
		a.Set(i, i + 1)
	}
	for i := 0; i < a.Size(); i++ {
		fmt.Printf("%v ", a.At(i))
	}

	fmt.Printf("\n")
	for iter := a.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
}
