package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/ketama"
)

func main() {
	k := ketama.New()
	k.Add("1.2.3.3")
	k.Add("2.4.5.6")
	k.Add("5.5.5.1")

	for i := 0; i < 10; i++ {
		node, _ := k.Get(fmt.Sprintf("%d", i))
		fmt.Printf("%v\n", node)
	}
	k.Remove("2.4.5.6")
}
