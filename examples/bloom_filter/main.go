package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/bloomfilter"
)

func main() {
	filter := bloom.New(100, 4, bloom.WithThreadSave())
	filter.Add("hhhh")
	filter.Add("gggg")

	fmt.Printf("%v\n", filter.Contains("aaaa"))
	fmt.Printf("%v\n", filter.Contains("gggg"))
}
