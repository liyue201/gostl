package main

import (
	"fmt"
	"github.com/liyue201/gostl/containers/bitmap"
)

func main() {
	bm := bitmap.New(1000)
	bm.Set(6)
	bm.Set(10)

	fmt.Printf("%v\n", bm.IsSet(5))
	fmt.Printf("%v\n", bm.IsSet(6))
	bm.Unset(6)
	fmt.Printf("%v\n", bm.IsSet(6))
}
